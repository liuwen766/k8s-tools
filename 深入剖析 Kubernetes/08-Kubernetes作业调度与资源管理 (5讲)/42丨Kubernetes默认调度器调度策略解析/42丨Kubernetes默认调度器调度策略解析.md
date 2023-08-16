- Predicates策略
- Predicates 在调度过程中的作用，可以理解为 Filter，即：它按照调度策略，从当前集群的所有节点中，“过滤”出一系列符合条件的节点。这些节点，
  都是可以运行待调度 Pod 的宿主机。
- 在 Kubernetes 中，默认的调度策略有如下三种：
- 1、GeneralPredicates：负责的是最基础的调度策略。比如，PodFitsResources 计算的就是宿主机的 CPU 和内存资源等是否够用。PodFitsHost
  检查的是，宿主机的名字是否跟 Pod 的 spec.nodeName 一致。PodFitsHostPorts 检查的是，Pod
  申请的宿主机端口（spec.nodePort）是不是跟已经被使用的端口有冲突。PodMatchNodeSelector 检查的是，Pod 的 nodeSelector 或者
  nodeAffinity 指定的节点，是否与待考察节点匹配，等等。 像这样一组 GeneralPredicates，正是 Kubernetes 考察一个 Pod
  能不能运行在一个 Node 上最基本的过滤条件。所以，GeneralPredicates
  也会被其他组件（比如 kubelet）直接调用。kubelet 在启动 Pod 前，会执行一个 Admit 操作来进行二次确认。这里二次确认的规则，就是执行一遍
  GeneralPredicates。
- 2、Volume 相关的过滤规则：比如，NoDiskConflict 检查多个 Pod 声明挂载的持久化 Volume 是否有冲突。MaxPDVolumeCountPredicate
  检查的条件，则是一个节点上某种类型的持久化 Volume 是不是已经超过了一定数目。VolumeZonePredicate，则是检查持久化 Volume 的
  Zone（高可用域）标签，是否与待考察节点的 Zone 标签相匹配。VolumeBindingPredicate 的规则。它负责检查的，是该 Pod 对应的 PV 的
  nodeAffinity 字段，是否跟某个节点的标签相匹配。
- 3、宿主机相关的过滤规则：主要考察待调度 Pod 是否满足 Node 本身的某些条件。比如，PodToleratesNodeTaints，负责检查 Pod
  的 Toleration 字段与 Node 的 Taint 字段是否匹配。NodeMemoryPressurePredicate，检查的是当前节点的内存是否充足。
- 4、Pod 相关的过滤规则：跟 GeneralPredicates 大多数是重合的。比较特殊的，是 PodAffinityPredicate。这个规则的作用，是检查待调度
  Pod 与 Node 上的已有 Pod 之间的亲密（affinity）和反亲密（anti-affinity）关系。

- 这四种类型的 Predicates，就构成了调度器确定一个 Node 可以运行待调度 Pod 的基本策略。在具体执行的时候， 当开始调度一个 Pod
  时，Kubernetes 调度器会同时启动 16 个 Goroutine，来并发地为集群里的所有 Node 计算 Predicates，最后返回可以运行这个 Pod
  的宿主机列表。

- Priorities策略
- 在 Predicates 阶段完成了节点的“过滤”之后，Priorities 阶段的工作就是为这些节点打分。这里打分的范围是 0-10 分，得分最高的节点就是最后被
  Pod 绑定的最佳节点。它依据一些计算公式进行打分。
- LeastRequestedPriority计算公式：【实际上就是在选择空闲资源（CPU 和 Memory）最多的宿主机】
  `score = (cpu((capacity-sum(requested))10/capacity) + memory((capacity-sum(requested))10/capacity))/2`
- BalancedResourceAllocation计算公式：【调度完成后，所有节点里各种资源分配最均衡的那个节点，从而避免一个节点上 CPU 被大量分配、而
  Memory 大量剩余的情况】
  `score = 10 - variance(cpuFraction,memoryFraction,volumeFraction)*10`
- NodeAffinityPriority、TaintTolerationPriority、InterPodAffinityPriority：一个 Node 满足这些规则的字段数目越多，它的得分就会越高。
- ImageLocalityPriority：如果待调度 Pod 需要使用的镜像很大，并且已经存在于某些 Node 上，那么这些 Node 的得分就会比较高。

- 在实际的执行过程中，调度器里关于集群和 Pod 的信息都已经缓存化，所以这些算法的执行过程还是比较快的。

- Kubernetes 调度器里其实还有一些默认不会开启的策略。可以通过为 kube-scheduler 指定一个配置文件或者创建一个
  ConfigMap，来配置哪些规则需要开启、哪些规则需要关闭。并且，可以通过为 Priorities 设置权重，来控制调度器的调度行为。
