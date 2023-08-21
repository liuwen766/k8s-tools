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


- Pod 的另一个重要的配置：容器健康检查和恢复机制。如
