- Pod 扮演的是传统部署环境里“虚拟机”的角色。这样的设计，是为了使用户从传统环境（虚拟机环境）向 Kubernetes（容器环境）的迁移，更加平滑。
- Pod 的设计，就是要让它里面的容器尽可能多地共享 Linux Namespace，仅保留必要的隔离和限制能力。这样，Pod
  模拟出的效果，就跟虚拟机里程序间的关系非常类似了。


- Pod 中几个重要字段的含义和用法：凡是调度、网络、存储，以及安全相关的属性，基本上是 Pod 级别的。
- NodeSelector：是一个供用户将 Pod 与 Node 进行绑定的字段。
- HostAliases：定义了 Pod 的 hosts 文件（比如 /etc/hosts）里的内容。
- 凡是跟容器的 Linux Namespace 相关的属性，一定是 Pod 级别的。eg：example-pod-pid.yaml
- 凡是 Pod 中的容器要共享宿主机的 Namespace，也一定是 Pod 级别的定义。eg：example-pod-ns.yaml

- 上面是Pod级别的字段定义，下面是Container级别的字段定义：
- ImagePullPolicy：定义了镜像拉取的策略。
- Lifecycle：定义的是 Container Lifecycle Hooks，在容器状态发生变化时触发一系列“钩子”。eg:example-lifecycle.yaml


- Pod 对象在 Kubernetes 中的生命周期Status——Pending、Running、Failed、Succeeded、Unknown。
- 应用：比如，Pod 当前的 Status 是 Pending，对应的 Condition 是 Unschedulable，这就意味着它的调度出现了问题。
