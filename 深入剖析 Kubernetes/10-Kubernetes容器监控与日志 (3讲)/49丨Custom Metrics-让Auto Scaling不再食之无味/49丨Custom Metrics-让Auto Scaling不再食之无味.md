- Kubernetes 里以Prometheus 为核心的监控体系的架构可以提供一种非常有用的能力，那就是 Custom Metrics，自定义监控指标。
- Kubernetes 里的 Custom Metrics 机制，也是借助 Aggregator APIServer 扩展机制来实现的。这里的具体原理是，当把 Custom
  Metrics APIServer 启动之后，Kubernetes 里就会出现一个叫作custom.metrics.k8s.io的 API。而当访问这个 URL 时，Aggregator
  就会把请求转发给 Custom Metrics APIServer 。
- 凭借强大的 API 扩展机制，Custom Metrics 已经成为了 Kubernetes 的一项标准能力。并且，Kubernetes 的自动扩展器组件
  Horizontal Pod Autoscaler （HPA）， 也可以直接使用 Custom Metrics 来执行用户指定的扩展策略，这里的整个过程都是非常灵活和可定制的。


- Custom Metrics 具体的使用方式：
- 首先，是先部署 Prometheus 项目。这一步，会使用 Prometheus Operator 来完成。
- 第二步，需要把 Custom Metrics APIServer 部署起来。
- 第三步，需要为 Custom Metrics APIServer 创建对应的ClusterRoleBinding，以便能够使用curl来直接访问Custom Metrics的API。
- 第四步，把待监控的应用和 HPA 部署起来。HPA 的配置，就是设置 Auto Scaling 规则的地方。


- Kubernetes 的 Aggregator APIServer，是一个非常行之有效的 API 扩展机制。
- Kubernetes 社区已经为你提供了一套叫作 KubeBuilder 的工具库，帮助你生成一个 API Server 的完整代码框架，你只需要在里面添加自定义
  API，以及对应的业务逻辑即可。
