- Deployment 对应用做了一个简单化假设。 它认为，一个应用的所有 Pod，是完全一样的。所以，它们互相之间没有顺序，也无所谓运行在哪台宿主机上。
  需要的时候，Deployment 就可以通过Pod 模板创建新的 Pod；不需要的时候，Deployment 就可以“杀掉”任意一个 Pod。
- 在实际的场景中，并不是所有的应用都可以满足这样的要求——**尤其是分布式应用，它的多个实例之间，往往有依赖关系，比如：主从关系、主备关系。还有就是数据存储类应用，它的多个实例，往往都会在本地磁盘上保存一份数据。**而这些实例一旦被杀掉，即便重建出来，实例与数据之间的对应关系也已经丢失，从而导致应用失败。
- **这种实例之间有不对等关系，以及实例对外部数据有依赖关系的应用，就被称为“有状态应用”（Stateful Application）**。
- StatefulSet 的设计其实非常容易理解。它把真实世界里的应用状态，抽象为了两种情况：
  - 1、拓扑状态。这种应用的多个实例之间不是完全对等的关系。这些应用实例，必须按照某些顺序启动，比如应用的主节点A要先于从节点B启动。
  - 2、存储状态。这种应用的多个实例分别绑定了不同的存储数据。最典型的例子，就是一个数据库应用的多个存储实例。
- StatefulSet 的核心功能，就是通过某种方式记录这些状态，然后在 Pod 被重新创建时，能够为新 Pod 恢复这些状态。



- Service 的访问方式：
  - 第一种方式，是以 Service 的 VIP（Virtual IP，即：虚拟 IP）方式。它会把请求转发到该 Service 所代理的某一个 Pod 上。
  - 第二种方式，就是以 Service 的 DNS 方式。只要访问“my-svc.my-namespace.svc.cluster.local”这条 DNS 记录，就可以访问到名叫
    my-svc 的 Service 所代理的某一个 Pod。
  - 在第二种 Service DNS 的方式下，具体可以分为两种处理方法：
    - 第一种处理方法，是 Normal Service。这种情况下，访问“my-svc.my-namespace.svc.cluster.local”解析到的，正是 my-svc 这个Service 的 VIP，后面的流程就跟 VIP 方式一致了。
    - 第二种处理方法，正是 Headless Service。这种情况下，访问“my-svc.my-namespace.svc.cluster.local”解析到的，直接就是 my-svc代理的某一个 Pod 的 IP 地址。可以看到，这里的区别在于，Headless Service 不需要分配一个 VIP，而是可以直接以 DNS 记录的方式解析出被代理Pod 的 IP 地址。



- StatefulSet 是如何使用这个 DNS 记录来维持 Pod 的拓扑状态?
- StatefulSet 这个控制器的主要作用之一，就是使用 Pod 模板创建 Pod 的时候，对它们进行编号，并且按照编号顺序逐一完成创建工作。而当StatefulSet 的“控制循环”发现 Pod 的“实际状态”与“期望状态”不一致，需要新建或者删除 Pod 进行“调谐”的时候，它会严格按照这些Pod 编号的顺序，逐一完成这些操作。
- 与此同时，通过 Headless Service 的方式，StatefulSet 为每个 Pod 创建了一个固定并且稳定的 DNS 记录，来作为它的访问入口。
- 通过这种方法，Kubernetes 就成功地将 Pod 的拓扑状态（比如：哪个节点先启动，哪个节点后启动），按照 Pod 的“名字 +编号”的方式固定了下来。此外，Kubernetes 还为每一个 Pod 提供了一个固定并且唯一的访问入口，即：这个 Pod 对应的 DNS 记录。
- 通过这种严格的对应规则，StatefulSet 就保证了 Pod 网络标识的稳定性。
