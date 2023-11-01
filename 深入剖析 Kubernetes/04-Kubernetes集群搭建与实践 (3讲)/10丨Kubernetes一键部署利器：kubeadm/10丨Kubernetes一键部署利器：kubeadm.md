- **要真正发挥容器技术的实力，就不能仅仅局限于对 Linux 容器本身的钻研和使用。** 

- 更深入的学习容器技术的关键在于，**如何使用这些技术来“容器化”的应用。** 

-  Kubernetes项目盛行起来的主要原因：这个项目体现出来的容器化“表达能力”，具有独有的先进性和完备性。这就使得它不仅能运行 Java Web 与 MySQL 这样的常规组合，还能够处理 Cassandra 容器集群等复杂编排问题。



- 作为一个典型的分布式项目，Kubernetes 的部署一直以来都是挡在初学者前面的一只“拦路虎”。尤其是在 Kubernetes 项目发布初期，它的部署完全要依靠一堆由社区维护的脚本。其实在部署时，它的每一个组件都是一个需要被执行的、单独的二进制文件。但是，除了将各个组件编译成二进制文件外，用户还要负责为这些二进制文件编写对应的配置文件、配置自启动脚本，以及为 kube-apiserver 配置授权文件等等诸多运维工作。 

-  难道 Kubernetes 项目就没有简单的部署方法了吗？—— 直到 2017 年，社区才终于发起了一个独立的部署工具，名叫：kubeadm—— 这个项目的目的，就是要让用户能够通过这样两条指令完成一个 Kubernetes 集群的部署：

  ```shell
  # 创建一个 Master 节点
  $ kubeadm init
   
  # 将一个 Node 节点加入到当前集群中
  $ kubeadm join <Master 节点的 IP 和端口 >
  ```

- 如果需要指定 kube-apiserver 的启动参数。

  ```shell
  # 在kubeadm.yaml文件里填写各种自定义的部署参数
  $ kubeadm init --config kubeadm.yaml
  ```

  

- 为什么不用容器部署 Kubernetes 呢？—— **如何容器化 kubelet**？ kubelet 是 Kubernetes 项目用来操作 Docker 等容器运行时的核心组件。可是，除了跟容器运行时打交道外，kubelet 在配置容器网络、管理容器数据卷时，都需要直接操作宿主机。而如果现在 kubelet 本身就运行在一个容器里，那么直接操作宿主机就会变得很麻烦。 



### kubeadm工作原理：

- kubeadm 方案——把 kubelet 直接运行在宿主机上，然后使用容器部署其他的 Kubernetes 组件。 
- 使用kubeadm 的前提，是在机器上手动安装 kubeadm、kubelet 和 kubectl 这三个二进制文件。 

#### kubeadm init 的工作流程：

- 执行 **kubeadm init** 指令后，kubeadm 首先要做的，是一系列的检查工作，以确定这台机器可以用来部署 Kubernetes。——这称为 “Preflight Checks” 。它会检查以下项：
  - Linux 内核的版本必须是否是 3.10 以上？
  - Linux Cgroups 模块是否可用？
  - 机器的 hostname 是否标准？在 Kubernetes 项目里，机器的名字以及一切存储在 Etcd 中的 API 对象，都必须使用标准的 DNS 命名（RFC 1123）。
  - 用户安装的 kubeadm 和 kubelet 的版本是否匹配？
  - 机器上是不是已经安装了 Kubernetes 的二进制文件？
  - Kubernetes 的工作端口 10250/10251/10252 端口是不是已经被占用？
  - ip、mount 等 Linux 指令是否存在？
  - Docker 是否已经安装？
  - ……

- **通过了 Preflight Checks 之后，kubeadm 要为你做的，是生成 Kubernetes 对外提供服务所需的各种证书和对应的目录**。即 /etc/kubernetes/pki/ca.{crt,key} 【 Kubernetes 对外提供服务时，除非专门开启“不安全模式”，否则都要通过 HTTPS 才能访问 kube-apiserver。这就需要为 Kubernetes 集群配置好证书文件。】
- **证书生成后，kubeadm 接下来会为其他组件生成访问 kube-apiserver 所需的配置文件**。 即 /etc/kubernetes/xxx.conf
-  **接下来，kubeadm 会为 Master 组件生成 Pod 配置文件**。即 /etc/kubernetes/manifests —— 这三个 Master 组件 kube-apiserver、kube-controller-manager、kube-scheduler，而它们都会被使用 Pod 的方式部署起来 【 在 Kubernetes 中，有一种特殊的容器启动方法叫做**“Static Pod”**。它允许你把要部署的 Pod 的 YAML 文件放在一个指定的目录里。这样，当这台机器上的 kubelet 启动时，它会自动检查这个目录，加载所有的 Pod YAML 文件，然后在这台机器上启动它们】
- 在这一步完成后，kubeadm 还会再生成一个 Etcd 的 Pod YAML 文件，用来通过同样的 Static Pod 的方式启动 Etcd。 
- Master 容器启动后，kubeadm 会通过检查 localhost:6443/healthz 这个 Master 组件的健康检查 URL，等待 Master 组件完全运行起来。 
- **然后，kubeadm 就会为集群生成一个 bootstrap token**。在后面，只要持有这个 token，任何一个安装了 kubelet 和 kubadm 的节点，都可以通过 kubeadm join 加入到这个集群当中。 
- **在 token 生成之后，kubeadm 会将 ca.crt 等 Master 节点的重要信息，通过 ConfigMap 的方式保存在 Etcd 当中，供后续部署 Node 节点使用**。 
- **kubeadm init 的最后一步，就是安装默认插件**。Kubernetes 默认 kube-proxy 和 DNS 这两个插件是必须安装的。它们分别用来提供整个集群的服务发现和 DNS 功能—— kubeadm通过Kubernetes客户端创建这两个 Pod。 

#### kubeadm join 的工作流程：

- kubeadm init 生成 bootstrap token 之后，就可以在任意一台安装了 kubelet 和 kubeadm 的机器上执行 kubeadm join。
- 任何一台机器想要成为 Kubernetes 集群中的一个节点，就必须在集群的 kube-apiserver 上注册。—— 因此kubeadm 至少需要发起一次“不安全模式”的访问到 kube-apiserver 来获取 cluster-info。—— 这里bootstrap token，扮演的就是这个“不安全访问”过程中的安全验证的角色。 
- 获取到cluster-info 里的 kube-apiserver 的地址、端口、证书之后，kubelet 就可以以“安全模式”连接到 apiserver 上，至此一个新的节点就部署完成。



