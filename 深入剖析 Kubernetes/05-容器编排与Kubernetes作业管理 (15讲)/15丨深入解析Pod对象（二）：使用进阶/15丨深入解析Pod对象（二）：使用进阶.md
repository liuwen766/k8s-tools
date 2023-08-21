- 在 Kubernetes 中，有几种特殊的 Volume，它们存在的意义不是为了存放容器里的数据，也不是用来进行容器和宿主机之间的数据交换。这些特殊
  Volume 的作用，是为容器提供预先定义好的数据。所以，从容器的角度来看，这些 Volume 里的信息就是仿佛是被
  Kubernetes“投射”（Project）进入容器当中的。这正是 Projected Volume 的含义。
- 目前为止，Kubernetes 支持的 Projected Volume 一共有四种： 1、Secret；2、ConfigMap；3、Downward API；4、ServiceAccountToken。

- Secret的作用，是把 Pod 想要访问的加密数据，存放到 Etcd 中。然后，就可以通过在 Pod 的容器里挂载 Volume 的方式，访问到这些
  Secret 里保存的信息。eg：example-secret.yaml
- 这里创建的 Secret 对象，它里面的内容仅仅是经过了base64转码，而并没有被加密。在真正的生产环境中，需要在 Kubernetes 中开启
  Secret 的加密插件，增强数据的安全性。
- 像这样通过挂载方式进入到容器里的 Secret，一旦其对应的 Etcd 里的数据被更新，这些 Volume 里的文件内容，同样也会被更新。其实，这是
  kubelet 组件在定时维护这些 Volume。需要注意的是，这个更新可能会有一定的延时。所以在编写应用程序时，在发起数据库连接的代码处写好
  重试和超时的逻辑，绝对是个好习惯。

```shell
# 创建secret 方式一：
$ cat ./username.txt
admin
$ cat ./password.txt
c1oudc0w!
 
$ kubectl create secret generic user --from-file=./username.txt
$ kubectl create secret generic pass --from-file=./password.txt

# 创建secret 方式二：
$ kubectl apply -f example-secret.yaml

# 通过base64加密
$ echo -n 'admin' | base64
YWRtaW4=
$ echo -n '1f2d1e2e67df' | base64
MWYyZDFlMmU2N2Rm
```

- ConfigMap：它与 Secret 的区别在于，ConfigMap 保存的是不需要加密的、应用所需的配置信息。
- Downward API：它的作用是让 Pod 里的容器能够直接获取到这个 Pod API 对象本身的信息。eg：example-download-api.yaml
- Downward API 能够获取到的信息，一定是 Pod 里的容器进程启动之前就能够确定下来的信息。

```shell
$ kubectl create -f dapi-volume.yaml
$ kubectl logs test-downwardapi-volume
cluster="test-cluster1"
rack="rack-22"
zone="us-est-coast"
```

- Service Account 对象的作用，就是 Kubernetes 系统内置的一种“服务账户”，它是 Kubernetes 进行权限分配的对象。
- 像这样的 Service Account 的授权信息和文件，实际上保存在它所绑定的一个特殊的 Secret 对象里的。这个特殊的 Secret
  对象，就叫作ServiceAccountToken。任何运行在 Kubernetes 集群上的应用，都必须使用这个 ServiceAccountToken 里保存的授权信息，也就是
  Token，才可以合法地访问 API Server。


- Pod 的另一个重要的配置：容器健康检查和恢复机制。eg：example-liveness.yaml
- Kubernetes 里的Pod 恢复机制，也叫 restartPolicy。它是 Pod 的 Spec 部分的一个标准字段（pod.spec.restartPolicy），默认值是
  Always，即：任何时候这个容器发生了异常，它一定会被重新创建。
- Always：在任何情况下，只要容器不在运行状态，就自动重启容器；
  OnFailure: 只在容器 异常时才自动重启容器；
  Never: 从来不重启容器。
- Pod 的恢复过程，永远都是发生在当前节点上，而不会跑到别的节点上去。事实上，一旦一个 Pod
  与一个节点（Node）绑定，除非这个绑定发生了变化（pod.spec.node 字段被修改），否则它永远都不会离开这个节点。这也就意味着，如果这个宿主机宕机了，这个
  Pod 也不会主动迁移到其他节点上去。

- 只要 Pod 的 restartPolicy 指定的策略允许重启异常的容器（比如：Always），那么这个 Pod 就会保持 Running 状态，并进行容器重启。否则，Pod
  就会进入 Failed 状态 。
- 对于包含多个容器的 Pod，只有它里面所有的容器都进入异常状态后，Pod 才会进入 Failed 状态。在此之前，Pod 都是 Running
  状态。此时，Pod的 READY 字段会显示正常容器的个数。


- 可以定义一个 PodPreset 对象。在这个对象中，凡是想在开发人员编写的 Pod 里追加的字段，都可以预先定义好。 PodPreset 是专门用来对
  Pod 进行批量化、自动化修改的工具对象。
  PodPreset里定义的内容，只会在 Pod API 对象被创建之前追加在这个对象本身上，而不会影响任何 Pod 的控制器的定义。
- 比如，现在提交的是一个 nginx-deployment，那么这个 Deployment 对象本身是永远不会被 PodPreset 改变的，被修改的只是这个
  Deployment 创建出来的所有 Pod。
- 如果定义了同时作用于一个 Pod 对象的多个 PodPreset，会发生什么呢？ 实际上，Kubernetes 项目会合并（Merge）这两个
  PodPreset 要做的修改。而如果它们要做的修改有冲突的话，这些冲突字段就不会被修改。