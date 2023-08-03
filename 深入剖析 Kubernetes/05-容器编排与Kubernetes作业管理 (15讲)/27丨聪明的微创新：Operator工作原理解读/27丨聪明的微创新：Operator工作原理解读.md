- 以 Etcd Operator 为例，来为你讲解一下 Operator 的工作原理和编写方法。
- 第一步，将这个 Operator 的代码 Clone 到本地：

```shell
$ git clone https://github.com/coreos/etcd-operator
```

- 第二步，将这个 Etcd Operator 部署在 Kubernetes 集群里。

1、先为 Etcd Operator 创建 RBAC 规则：

```shell
$ example/rbac/create_role.sh
```

2、创建Etcd Operator，即一个CRD

```shell
$ kubectl create -f example/deployment.yaml
```

3、创建Etcd集群，即CRD的一个具体实例CR

```shell
$ kubectl apply -f example/example-etcd-cluster.yaml
```

- Operator 的工作原理，实际上是利用了 Kubernetes 的自定义 API 资源（CRD），来描述我们想要部署的“有状态应用”；然后在自定义控制器里，根据自定义
  API 对象的变化，来完成具体的部署和运维工作。
