- Deployment、StatefulSet，以及 DaemonSet 这三个编排概念,它们主要编排的对象，都是“在线业务”，即：Long Running
  Task（长作业）。比如，我在前面举例时常用的 Nginx、Tomcat，以及 MySQL 等等。这些应用一旦运行起来，除非出错或者停止，它的容器进程会一直保持在
  Running 状态。
- 有一类作业是“离线业务”，或者叫作 Batch Job（计算业务）。
  如example-job.yaml
- Job Controller 的工作原理：
  并行度（parallelism）——Job 最大的并行数
  任务总数（completions）——Job 最小的完成数
  首先，Job Controller 控制的对象，直接就是 Pod。
  其次，Job Controller 在控制循环中进行的调谐（Reconcile）操作，是根据实际在 Running 状态 Pod 的数目、已经成功退出的 Pod
  的数目，以及 parallelism、completions 参数的值共同计算出在这个周期里，应该创建或者删除的 Pod 数目，然后调用 Kubernetes API
  来执行这个操作。


- 三种常用的、使用 Job 对象的方法：
1、默认并行度和任务总数——外部管理器 +Job 模板。 如job-tmpl.yaml
```shell
$ mkdir ./jobs
# 外部管理器[外部工具]
$ for i in apple banana cherry
do
  cat job-tmpl.yaml | sed "s/\$ITEM/$i/" > ./jobs/job-$i.yaml
done

$ kubectl create -f ./jobs
$ kubectl get pods -l jobgroup=jobexample
```
2、指定任务总数——拥有固定任务数目的并行 Job。如job-tmpl2.yaml
3、指定并行度——但不设置固定的 completions 的值。如job-tmpl3.yaml

- 一个非常有用的 Job 对象，叫作：CronJob——CronJob 是一个 Job 对象的控制器（Controller）。
  如example-cj.yaml
- 由于定时任务的特殊性，很可能某个 Job 还没有执行完，另外一个新 Job 就产生了。这时候，你可以通过 spec.concurrencyPolicy 字段来定义具体的处理策略。比如：
concurrencyPolicy=Allow，这也是默认情况，这意味着这些 Job 可以同时存在；
concurrencyPolicy=Forbid，这意味着不会创建新的 Pod，该创建周期被跳过；
concurrencyPolicy=Replace，这意味着新产生的 Job 会替换旧的、没有执行完的 Job。

- spec.startingDeadlineSeconds=200 可以表示如果某一次 Job 创建失败，这次创建就会被标记为“miss”。当在指定的时间窗口内【这里是200s】，miss 的数目达到 100 时，那么 CronJob 会停止再创建这个 Job。



