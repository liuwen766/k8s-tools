- 在 Kubernetes 项目中，负责完成授权（Authorization）工作的机制，就是 RBAC：基于角色的访问控制（Role-Based Access Control）。
- 明确三个最基本的概念。
  - 1、Role：角色，它其实是一组规则，定义了一组对 Kubernetes API 对象的操作权限。
  - 2、Subject：被作用者，既可以是“人”，也可以是“机器”，也可以使你在 Kubernetes 里定义的“用户”。
  - 3、RoleBinding：定义了“被作用者”和“角色”的绑定关系。


- Role 和 RoleBinding 对象都是 Namespaced 对象（Namespaced Object），它们对权限的限制规则仅在它们自己的 Namespace 内有效，roleRef 也只能引用当前 Namespace 里的 Role 对象。
- ClusterRole 和 ClusterRoleBinding 对象都是 Non-Namespaced 对象（比如：Node）。

- 一个所有权限的 verbs 字段的全集: ["get", "list", "watch", "create", "update", "patch", "delete"]【这是
  Kubernetes（v1.11）里能够对 API 对象进行的所有操作】


- ServiceAccount：由 Kubernetes 负责管理的“内置用户”。k8s中最普遍的用法还是 ServiceAccount。 它分配权限的过程：
  - 首先，我们要定义一个 ServiceAccount。如 example-serviceaccount.yaml
  - 然后，我们通过编写 RoleBinding 的 YAML 文件，来为这个 ServiceAccount 分配权限。如 serviceaccount-rolebinding.yaml
  - 接着，我们用 kubectl 命令创建这三个对象：

```shell
$ kubectl create -f example-serviceaccount.yaml
$ kubectl create -f serviceaccount-rolebinding.yaml
$ kubectl create -f example-role.yaml

# 查看
$ kubectl get sa -n mynamespace -o yaml
- apiVersion: v1
  kind: ServiceAccount
  metadata:
    creationTimestamp: 2018-09-08T12:59:17Z
    name: example-sa
    namespace: mynamespace
    resourceVersion: "409327"
    ...
  secrets:         # k8s 会为一个 ServiceAccount 自动创建并分配一个 Secret 对象   这个Secret，就是这个 ServiceAccount 对应的、用来跟 APIServer 进行交互的授权文件
  - name: example-sa-token-vmfg6
```

- 启动一个跑pod使用这个ServiceAccount，如 pod-with-sa.yaml 文件。
- 等这个 Pod 运行起来之后，我们就可以看到，该 ServiceAccount 的 token，也就是一个 Secret 对象，被 Kubernetes 自动挂载到了容器的 /var/run/secrets/kubernetes.io/serviceaccount 目录下。
- 在生产环境中，强烈建议为所有 Namespace 下的默认 ServiceAccount，绑定一个只读权限的 Role。
- 用户 system:serviceaccount:<ServiceAccount 名字 > 和用户组 system:serviceaccounts:<Namespace 名字>如：user-and-usergroup.yaml



- 在 Kubernetes 中已经内置了很多个为系统保留的 ClusterRole，它们的名字都以 system: 开头。可以通过 kubectl getclusterroles 查看到它们。一般来说，这些系统 ClusterRole，是绑定给 Kubernetes 系统组件对应的ServiceAccount 使用的。
- Kubernetes 还提供了四个预先定义好的 ClusterRole 来供用户直接使用：
  - 1、cluster-admin——Kubernetes 项目中的最高权限（verbs=*）
  - 2、admin
  - 3、edit
  - 4、view——只有 Kubernetes API 的只读权限
- eg：如下：这个 system:kube-scheduler 的 ClusterRole，就会被绑定给 kube-system Namesapce 下名叫 kube-scheduler 的ServiceAccount，它正是 Kubernetes 调度器的 Pod 声明使用的 ServiceAccount。

```shell
$ kubectl describe clusterrole system:kube-scheduler
Name:         system:kube-scheduler
...
PolicyRule:
  Resources                    Non-Resource URLs Resource Names    Verbs
  ---------                    -----------------  --------------    -----
...
  services                     []                 []                [get list watch]
  replicasets.apps             []                 []                [get list watch]
  statefulsets.apps            []                 []                [get list watch]
  replicasets.extensions       []                 []                [get list watch]
  poddisruptionbudgets.policy  []                 []                [get list watch]
  pods/status                  []                 []                [patch update]
```

