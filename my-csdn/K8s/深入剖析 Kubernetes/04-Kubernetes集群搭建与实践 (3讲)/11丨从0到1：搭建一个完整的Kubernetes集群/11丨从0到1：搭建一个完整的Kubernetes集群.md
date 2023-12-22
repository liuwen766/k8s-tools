-  kubeadm 能够用于生产环境吗？ ——不能。 因为 kubeadm 目前最欠缺的是，一键部署一个高可用的 Kubernetes 集群，即：Etcd、Master 组件都应该是多节点集群，而不是现在这样的单点。 


 

- 从 0 到 1 搭建一个完整的 Kubernetes 集群。 
  - 在所有节点上安装 Docker 和 kubeadm；
  - 部署 Kubernetes Master；
  - 部署容器网络插件；
  - 部署 Kubernetes Worker；
  - 部署 Dashboard 可视化插件；
  - 部署容器存储插件。



- 参考 从0开始安装k8s1.25：https://blog.csdn.net/qq_41822345/article/details/126679925



-  默认情况下 Master 节点是不允许运行用户 Pod 的。而 Kubernetes 做到这一点，依靠的是 Kubernetes 的 Taint/Toleration 机制——它的原理非常简单：一旦某个节点被加上了一个 Taint，即被“打上了污点”，那么所有 Pod 就都不能在这个节点上运行，因为 Kubernetes 的 Pod 都有“洁癖”。  除非，有个别的 Pod 声明自己能“容忍”这个“污点”，即声明了 Toleration，它才可以在这个节点上运行 。

  ```shell
  # 为节点打上“污点”（Taint）的命令是:
  $ kubectl taint nodes node1 foo=bar:NoSchedule
  # Pod 声明 Toleration
  apiVersion: v1
  kind: Pod
  ...
  spec:
    tolerations:
    - key: "foo"
      operator: "Equal"
      value: "bar"
      effect: "NoSchedule"
  # 如果是想要一个单节点的 Kubernetes，删除这个 Taint 才是正确的选择：
  $ kubectl taint nodes --all node-role.kubernetes.io/master-
  ```

  

- Kubernetes 集群的部署过程并不像传说中那么繁琐，这主要得益于： 
  - kubeadm 项目大大简化了部署 Kubernetes 的准备工作，尤其是配置文件、证书、二进制文件的准备和制作，以及集群版本管理等操作，都被 kubeadm 接管了。
  - Kubernetes 本身“一切皆容器”的设计思想，加上良好的可扩展机制，使得插件的部署非常简便。

- 上述思想，也是开发和使用 Kubernetes 的重要指导思想，即：基于 Kubernetes 开展工作时，一定要优先考虑这两个问题：
  - 我的工作是不是可以容器化？
  - 我的工作是不是可以借助 Kubernetes API 和可扩展机制来完成？

-  一旦工作能够基于 Kubernetes 实现容器化，就很有可能像 Kubernetes 部署过程一样，大幅简化原本复杂的运维工作。对于时间宝贵的技术人员来说，这个变化的重要性是不言而喻的。 