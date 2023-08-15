- Kubernetes资源模型——Kubernetes 的资源管理与调度部分的基础，我们要从它的资源模型开始说起。
- 在 Kubernetes 里，Pod 是最小的原子调度单位。这也就意味着，所有跟调度和资源管理相关的属性都应该是属于 Pod
  对象的字段。而这其中最重要的部分，就是 Pod 的 CPU 和内存配置。eg: example-resource.yaml
- CPU 这样的资源被称作“可压缩资源”（compressible resources）。它的典型特点是，当可压缩资源不足时，Pod 只会“饥饿”，但不会退出。
- 内存这样的资源，则被称作“不可压缩资源（incompressible resources）。当不可压缩资源不足时，Pod 就会因为
  OOM（Out-Of-Memory）被内核杀掉。
- request和limit的区别：在调度的时候，kube-scheduler 只会按照 requests 的值进行计算。而在真正设置 Cgroups 限制的时候，kubelet
  则会按照 limits 的值来进行设置。
- Kubernetes 的 requests+limits 的配合：用户在提交 Pod 时，可以声明一个相对较小的 requests 值供调度器使用，而 Kubernetes
  真正设置给容器 Cgroups 的，则是相对较大的 limits 值。


- Kubernetes的QoS模型
- Guaranteed类别：当 Pod 里的每一个 Container 都同时设置了 requests 和 limits，并且 requests 和 limits 值相等。 当 Pod 仅设置了
  limits 没有设置 requests 的时候，Kubernetes 会自动为它设置与 limits 相同的 requests 值。
- BurstAble类别：当 Pod 不满足 Guaranteed 的条件，但至少有一个 Container 设置了 requests。
- BestEffort类别：如果一个 Pod 既没有设置 requests，也没有设置 limits，那么它的 QoS 类别就是 BestEffort
- QoS 划分的主要应用场景，是当宿主机资源紧张的时候，kubelet 对 Pod 进行 Eviction（即资源回收）时需要用到的。当 Kubernetes
  所管理的宿主机上不可压缩资源短缺时，就有可能触发 Eviction。比如，可用内存（memory.available）、可用的宿主机磁盘空间（nodefs.available），
  以及容器运行时镜像存储空间（imagefs.available）等等

```shell
# Kubernetes 为你设置的 Eviction 的默认阈值如下所示：
memory.available<100Mi
nodefs.available<10%
nodefs.inodesFree<5%
imagefs.available<15%
# 上述各个触发条件在 kubelet 里都是可配置的，如下：
kubelet --eviction-hard=imagefs.available<10%,memory.available<500Mi,nodefs.available<5%,nodefs.inodesFree<5% --eviction-soft=imagefs.available<30%,nodefs.available<10% --eviction-soft-grace-period=imagefs.available=2m,nodefs.available=2m --eviction-max-pod-grace-period=600
```

- Eviction 在 Kubernetes 里其实分为 Soft 和 Hard 两种模式。
- Hard Eviction 模式下，Eviction 过程就会在阈值达到之后立刻开始。Soft意味着不足的阈值达到多少分钟【可设置】之后，kubelet 才会开始
  Eviction 的过程。
- 当宿主机的 Eviction 阈值达到后，就会进入 MemoryPressure 或者 DiskPressure 状态，从而避免新的 Pod 被调度到这台宿主机上。

- 当 Eviction 发生的时候，kubelet 具体会挑选哪些 Pod 进行删除操作，就需要参考这些 Pod 的 QoS 类别。删除顺序是：先BestEffort
  → 其次BurstAble → 最后Guaranteed。


- Kubernetes 里一个非常有用的特性：cpuset 的设置。
- 在使用容器的时候，你可以通过设置 cpuset 把容器绑定到某个 CPU 的核上，而不是像 cpushare 那样共享 CPU 的计算能力。
- 由于操作系统在 CPU 之间进行上下文切换的次数大大减少，容器里应用的性能会得到大幅提升。事实上，cpuset 方式，是生产环境里部署在线应用类型的
  Pod 时，非常常用的一种方式。
- 在 Kubernetes 里应该如何实现：首先，你的 Pod 必须是 Guaranteed 的 QoS 类型；
  然后，你只需要将 Pod 的 CPU 资源的 requests 和 limits 设置为同一个相等的整数值即可。
