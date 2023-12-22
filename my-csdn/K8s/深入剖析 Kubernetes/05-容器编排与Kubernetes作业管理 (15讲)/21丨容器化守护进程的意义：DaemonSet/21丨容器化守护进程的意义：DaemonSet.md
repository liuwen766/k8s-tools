
- StatefulSet 其实就是对现有典型运维业务的容器化抽象【参考上文好好理解这句话】。
- StatefulSet 可以说是 Kubernetes 项目中最为复杂的编排对象。
```shell
# 比如只更新一部分pod。 只有序号大于或者等于 2 的 Pod 会被更新到这个版本
kubectl patch statefulset mysql -p '{"spec":{"updateStrategy":{"type":"RollingUpdate","rollingUpdate":{"partition":2}}}}'
statefulset.apps/mysql patched
```


- **DaemonSet 又是如何保证每个 Node 上有且只有一个被管理的 Pod 呢**？
- DaemonSet Controller，首先从 Etcd 里获取所有的 Node 列表，然后遍历所有的 Node。这时，它就可以很容易地去检查，当前这个 Node 上是不是有一个携带了 name=fluentd-elasticsearch 标签的 Pod 在运行。
而检查的结果，可能有这么三种情况：
- 1、没有这种 Pod，那么就意味着要在这个 Node 上创建这样一个 Pod；【用 nodeSelector 或者 nodeAffinity】
- 2、有这种 Pod，但是数量大于 1，那就说明要把多余的 Pod 从这个 Node 上删除掉；【直接调用 Kubernetes API进行删除】
- 3、正好只有一个这种 Pod，那说明这个节点是正常的。
- DaemonSet 只管理 Pod 对象，然后通过 nodeAffinity 和 Toleration 这两个调度器的小功能，保证了每个节点上有且只有一个 Pod。
- 当然，DaemonSet 并不需要修改用户提交的 YAML 文件里的 Pod 模板里的nodeAffinity，而是在向 Kubernetes 发起请求之前，直接修改根据模板生成的 Pod 对象。
- 此外，DaemonSet 自动地给被管理的 Pod 加上了一些特殊的 “容忍” Toleration，来容忍一些 “污点” Taints，以保证每个节点上都能被调度一个 Pod。


- Deployment资源通过控制ReplicaSet对象来进行版本控制。那么DaemonSet[它直接控制pod对象]如何控制版本？→ ControllerRevision对象。
```shell
#查看 fluentd-elasticsearch 对应的 ControllerRevision
kubectl get controllerrevision -n kube-system -l name=fluentd-elasticsearch
#将 DaemonSet 回滚到 某个版本 eg：Revision=1 时的状态
kubectl rollout undo daemonset fluentd-elasticsearch --to-revision=1 -n kube-system
```
