- Service 的访问信息在 Kubernetes 集群之外，其实是无效的。这其实也容易理解：所谓 Service 的访问入口，其实就是每台宿主机上由
  kube-proxy 生成的 iptables 规则，以及 kube-dns 生成的 DNS 记录。而一旦离开了这个集群，这些信息对用户来说，也就自然没有作用了。

- 实际上，在理解了 Kubernetes Service 机制的工作原理之后，很多与 Service 相关的问题，其实都可以通过分析 Service 在宿主机上对应的
  iptables 规则（或者 IPVS 配置）得到解决。

- 如何从外部（Kubernetes 集群之外），访问到 Kubernetes 里创建的 Service？

- 从外部访问 Service 的第一种方式——最常用的一种方式就是：NodePort。【example-nodeport.yaml】

- Service NodePort 模式的工作原理：显然，kube-proxy 要做的就是在每台宿主机上生成这样一条 iptables 规则：

```shell
-A KUBE-NODEPORTS -p tcp -m comment --comment "default/my-nginx: nodePort" -m tcp --dport 8080 -j KUBE-SVC-67RL4FN6JRUPOJYM
```

- 在 NodePort 方式下，Kubernetes 会在 IP 包离开宿主机发往目的 Pod 时，对这个 IP 包做一次 SNAT 操作。
- 因为如果没有做 SNAT 操作的话，这时候，被转发来的 IP 包的源地址就是 client 的 IP 地址。所以此时，Pod 就会直接将回复发给client。对于
  client 来说，它的请求明明发给了 node 2，收到的回复却来自 node 1，这个 client 很可能会报错。

```shell
           client
             \ ^
              \ \
               v \
   node 1 <--- node 2
    | ^   SNAT
    | |   --->
    v |
 endpoint
```

- 将 Service 的 spec.externalTrafficPolicy 字段设置为 local，这就保证了所有 Pod 通过 Service 收到请求之后，一定可以看到真正的、外部
  client 的源地址。这个机制的实现原理也非常简单：这时候，一台宿主机上的 iptables 规则，会设置为只将 IP 包转发给运行在这台宿主机上的Pod。

- 从外部访问 Service 的第二种方式，适用于公有云上的 Kubernetes 服务。可以指定一个 LoadBalancer 类型的
  Service。【example-loadbalancer.yaml】
- LoadBalancer原理：在公有云提供的 Kubernetes 服务里，都使用了一个叫作 CloudProvider 的转接层，来跟公有云本身的 API
  进行对接。所以，在上述 LoadBalancer 类型的 Service 被提交后，Kubernetes 就会调用 CloudProvider 在公有云上为你创建一个负载均衡服务，
  并且把被代理的 Pod 的 IP 地址配置给负载均衡服务做后端。


- 从外部访问 Service 的第三种方式，是 Kubernetes 在 1.7 之后支持的一个新特性，叫作ExternalName。【example-externalname.yaml】
- ExternalName工作原理：当通过 Service 的 DNS 名字访问它的时候，比如访问：my-service.default.svc.cluster.local。那么，Kubernetes
  为你返回的就是my.database.example.com。所以说，ExternalName 类型的 Service，其实是在 kube-dns 里为你添加了一条 CNAME
  记录。这时，访问 my-service.default.svc.cluster.local 就和访问 my.database.example.com 这个域名是一个效果。

- 重点——SVC访问不通排查思路：【区分到底是 Service 本身的配置问题，还是集群的 DNS 出了问题】
- 1、检查 Kubernetes 自己的 Master 节点的 Service DNS 是否正常。【如果访问kubernetes.default返回的值都有问题，那就需要检查
  kube-dns 的运行状态和日志】
- 2、如果 Service 没办法通过 ClusterIP 访问到的时候，首先应该检查的是这个 Service 是否有 Endpoints。
- 3、如果 Endpoints 正常，那么就需要确认 kube-proxy 是否在正确运行。
- 4、如果 kube-proxy 一切正常，就应该仔细查看宿主机上的 Service 对应的iptables规则。
- 5、还有一种典型问题，就是 Pod 没办法通过 Service 访问到自己。【需要确保将 kubelet 的 hairpin-mode 设置为 hairpin-veth 或者
  promiscuous-bridge 即可】

```shell
# 在一个 Pod 里执行
$ nslookup kubernetes.default
Server:    10.0.0.10
Address 1: 10.0.0.10 kube-dns.kube-system.svc.cluster.local
 
Name:      kubernetes.default
Address 1: 10.0.0.1 kubernetes.default.svc.cluster.local



$ kubectl get endpoints hostnames
NAME        ENDPOINTS
hostnames   10.244.0.5:9376,10.244.0.6:9376,10.244.0.7:9376
```

- 一个 iptables 模式的 Service 对应的规则，它们包括:
- 1、KUBE-SERVICES 或者 KUBE-NODEPORTS 规则对应的 Service 的入口链，这个规则应该与 VIP 和 Service 端口一一对应； 
- 2、KUBE-SEP-(hash) 规则对应的 DNAT 链，这些规则应该与 Endpoints 一一对应； 
- 3、KUBE-SVC-(hash) 规则对应的负载均衡链，这些规则的数目应该与 Endpoints 数目一致； 
- 4、如果是 NodePort 模式的话，还有 POSTROUTING 处的 SNAT 链。

- 所谓 Service，其实就是 Kubernetes 为 Pod 分配的、固定的、基于 iptables（或者 IPVS）的访问入口。而这些访问入口代理的 Pod
  信息，则来自于 Etcd，由 kube-proxy 通过控制循环来维护。
