- Kubernetes 之所以需要 Service，一方面是因为 Pod 的 IP 不是固定的，另一方面则是因为一组 Pod 实例之间总会有负载均衡的需求。
- 被 selector 选中的 Pod，就称为 Service 的 Endpoints，可以使用 kubectl get ep 命令看到它们，需要注意的是，只有处于 Running
  状态，且 readinessProbe 检查通过的 Pod，才会出现在 Service 的 Endpoints 列表里。并且，当某一个 Pod 出现问题时，Kubernetes
  会自动把它从 Service 里摘除掉。

- Service 是由 kube-proxy 组件，加上 iptables 来共同实现的。

- iptables 模式的工作原理：
-
- 举个例子，对于我们前面创建的名叫 hostnames 的 Service 来说，一旦它被提交给 Kubernetes，那么 kube-proxy 就可以通过 Service
  的 Informer 感知到这样一个 Service 对象的添加。而作为对这个事件的响应，它就会在宿主机上创建这样一条 iptables 规则（你可以通过
  iptables-save 看到它），这条 iptables 规则的含义是：凡是目的地址是 10.0.1.175、目的端口是 80 的 IP 包，都应该跳转到另外一条名叫
  KUBE-SVC-NWV5X2332I4OT4T3的 iptables 链【它是一组规则的集合：指向的最终目的地pod的DNAT规则链，用于Service
  实现负载均衡。】进行处理。如下所示：

```shell
-A KUBE-SERVICES -d 10.0.1.175/32 -p tcp -m comment --comment "default/hostnames: cluster IP" -m tcp --dport 80 -j KUBE-SVC-NWV5X2332I4OT4T3
```

- DNAT 规则的作用，就是在 PREROUTING 检查点之前，也就是在路由之前，将流入 IP 包的目的地址和端口，改成–to-destination
  所指定的新的目的地址和端口。可以看到，这个目的地址和端口，正是被代理 Pod 的 IP 地址和端口。

- 访问 Service VIP 的 IP 包经过上述 iptables 处理之后，就已经变成了访问具体某一个后端 Pod 的 IP 包。这些 Endpoints 对应的
  iptables 规则，正是 kube-proxy 通过监听 Pod 的变化事件，在宿主机上生成并维护的。

- Kubernetes 的 kube-proxy 还支持一种叫作 IPVS 的模式。

- kube-proxy 通过 iptables 处理 Service 的过程，其实需要在宿主机上设置相当多的 iptables 规则。而且，kube-proxy
  还需要在控制循环里不断地刷新这些规则来确保它们始终是正确的。所以当宿主机上有大量 Pod 的时候，成百上千条 iptables
  规则不断地被刷新，会大量占用该宿主机的 CPU 资源，甚至会让宿主机“卡”在这个过程中。
- 因此，一直以来，基于 iptables 的 Service 实现，都是制约 Kubernetes 项目承载更多量级的 Pod 的主要障碍。

- IPVS 模式的工作原理：
- 首先，创建 Service 之后，kube-proxy 首先会在宿主机上创建一个虚拟网卡（叫作：kube-ipvs0），并为它分配 Service VIP 作为 IP 地址。
- 接下来，kube-proxy 就会通过 Linux 的 IPVS 模块，为这个 IP 地址设置三个 IPVS
  虚拟主机，并设置这三个虚拟主机之间使用轮询模式 (rr) 来作为负载均衡策略。
- 这时，任何发往 service 的请求，就都会被 IPVS 模块转发到某一个后端 Pod 上。

```shell
# ip addr
  ...
  73:kube-ipvs:<BROADCAST,NOARP>  mtu 1500 qdisc noop state DOWN qlen 1000
  link/ether  1a:ce:f5:5f:c1:4d brd ff:ff:ff:ff:ff:ff
  inet 10.0.1.175/32  scope global kube-ipvs0
  valid_lft forever  preferred_lft forever
# ipvsadm -ln
 IP Virtual Server version 1.2.1 (size=4096)
  Prot LocalAddress:Port Scheduler Flags
    ->  RemoteAddress:Port           Forward  Weight ActiveConn InActConn     
  TCP  10.102.128.4:80 rr
    ->  10.244.3.6:9376    Masq    1       0          0         
    ->  10.244.1.7:9376    Masq    1       0          0
    ->  10.244.2.3:9376    Masq    1       0          0
```

- IPVS高性能原理【将重要操作放入内核态】：相比于 iptables，IPVS 在内核中的实现其实也是基于 Netfilter 的 NAT 模式，
  所以在转发这一层上，理论上IPVS并没有显著的性能提升。但是，IPVS 并不需要在宿主机上为每个 Pod 设置
  iptables规则，而是把对这些“规则”的处理放到了内核态，从而极大地降低了维护这些规则的代价。
- 在大规模集群里，非常建议为 kube-proxy 设置–proxy-mode=ipvs 来开启这个功能。它为 Kubernetes 集群规模带来的提升，还是非常巨大的。


- 在 Kubernetes 中，Service 和 Pod 都会被分配对应的 DNS A 记录（从域名解析 IP 的记录）。
- 对于 ClusterIP 模式【ClusterIP 模式的 Service 提供的，就是一个 Pod 的稳定的 IP 地址，即 VIP】的 Service 来说，它代理的
  Pod 被自动分配的 A 记录的格式是：..pod.cluster.local。这条记录指向 Pod 的 IP
  地址。
- 对 Headless 模式【Headless Service 为你提供的，则是一个 Pod 的稳定的 DNS 名字】来说，它代理的 Pod 被自动分配的 A
  记录的格式是：...svc.cluster.local。这条记录也指向 Pod 的 IP 地址。
