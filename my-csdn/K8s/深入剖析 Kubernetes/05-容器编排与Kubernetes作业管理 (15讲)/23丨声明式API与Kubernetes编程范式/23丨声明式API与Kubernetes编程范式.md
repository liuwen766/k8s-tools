- 什么事声明式api？对比一下docker swarm 和 k8s 创建 两个Nginx容器并更新nginx镜像的命令，如下：

```shell
# 如果是 docker swarm —— 命令式命令行操作
docker service create --name nginx --replicas 2  nginx
docker service update --image nginx:1.7.9 nginx
# 如果是 k8s —— 命令式配置文件操作
kubectl create -f nginx.yaml
kubectl replace -f nginx.yaml
# 如果是 k8s —— 声明式 API操作  
# 执行这个命令，Kubernetes 就会立即触发这个 Deployment 的“滚动更新”
kubectl apply -f nginx.yaml
```

- **kubectl replace 的执行过程，是使用新的 YAML 文件中的 API 对象，替换原有的 API 对象；而 kubectl apply，则是执行了一个对原有API 对象的 PATCH 操作**。
- kubectl set image 和 kubectl edit 也是对已有 API 对象的修改。
- **声明式 API 在实际使用时的重要意义**：kube-apiserver 在响应命令式请求（比如，kubectl
  replace）的时候，一次只能处理一个写请求，否则会有产生冲突的可能。而对于声明式请求（比如，kubectl apply），一次能处理多个写操作，并且具备Merge 能力。
- 以 Istio 项目为例，来讲解一下声明式 API 在实际使用时的重要意义。
- Istio 项目，实际上就是一个基于 Kubernetes 项目的微服务治理框架。
- Istio 最根本的组件，是运行在每一个应用 Pod 里的 Envoy 容器【一个网络代理服务】。 Istio 项目，则把这个代理服务以 sidecar容器的方式，运行在了每一个被治理的应用 Pod 中。
  Pod 里的所有容器都共享同一个 Network Namespace。所以，Envoy 容器就能够通过配置 Pod 里的 iptables 规则，把整个 Pod的进出流量接管下来。
- 因此，Istio 的控制层（Control Plane）里的 Pilot 组件，就能够通过调用每个 Envoy 容器的 API，对这个 Envoy 代理进行配置，从而实现微服务治理。
- 在整个微服务治理的过程中，无论是对 Envoy 容器的部署，还是像上面这样对 Envoy 代理的配置，用户和应用都是完全“无感”的。Istio 项目使用的，是 Kubernetes 中的一个非常重要的功能，叫作 Dynamic Admission Control。也叫作：Initializer。如**pod-with-istio.yaml**。
- Istio原理：Istio 要做的就是编写一个用来为 Pod“自动注入”Envoy 容器的 Initializer
  - 1、首先，Istio 会将这个 Envoy 容器本身的定义，以 ConfigMap 的方式保存在 Kubernetes 当中。
    如**envoy-initializer.yaml**。在 Initializer 更新用户的 Pod 对象的时候，必须使用 PATCH API 来完成。而这种 PATCH API，正是声明式 API 最主要的能力。
  - 2、接下来，Istio 将一个编写好的 Initializer，作为一个 Pod 部署在 Kubernetes 中。
    如 **envoy-pod.yaml**。
  - 3、TwoWayMergePatch方法：使得我们可以直接使用新旧两个 Pod 对象。Initializer 的代码就可以使用这个 patch 的数据，调用Kubernetes 的 Client，发起一个 PATCH 请求。这样，一个用户提交的 Pod 对象里，就会被自动加上 Envoy 容器相关的字段。



- Kubernetes 还允许通过配置，来指定要对什么样的资源进行这个 Initialize 操作。如**InitializerConfiguration.yaml**，表示k8s 要对所有的 Pod 进行这个 Initialize 操作。
- 每一个新创建的 Pod，都会自动携带了 metadata.initializers.pending 的 Metadata 信息。它也正是接下来 Initializer 的控制器判断这个Pod 有没有执行过自己所负责的初始化操作的重要依据。
  Demo样例地址：https://github.com/resouer/kubernetes-initializer-tutorial
- **Istio 项目的核心，就是由无数个运行在应用 Pod 中的 Envoy 容器组成的服务代理网格。这也正是 Service Mesh 的含义**。
- **无论是对 sidecar 容器的巧妙设计，还是对 Initializer 的合理利用，Istio 项目的设计与实现，其实都依托于 Kubernetes 的声明式API 和它所提供的各种编排能力**。
- 这个机制得以实现的原理，正是借助了 Kubernetes 能够对 API 对象进行在线更新的能力，这也正是Kubernetes“声明式 API”的独特之处：
  - 首先，所谓“声明式”，指的就是只需要提交一个定义好的 API 对象来“声明”，所期望的状态是什么样子。
  - 其次，“声明式 API”允许有多个 API 写端，以 PATCH 的方式对 API 对象进行修改，而无需关心本地原始 YAML 文件的内容。
  - 最后，也是最重要的，有了上述两个能力，Kubernetes 项目才可以基于对 API 对象的增、删、改、查，在完全无需外界干预的情况下，完成对“实际状态”和“期望状态”的调谐（Reconcile）过程。
- 所以说，声明式 API，才是 Kubernetes 项目编排能力“赖以生存”的核心所在。
