- 监听相关的参数主要有以下几个：
  - listeners 
  - advertised.listeners
  - listener.security.protocol.map
  - inter.broker.listener.name
  - security.inter.broker.protocol
  - control.plane.listener.name



- 有时候我们会碰到网络是通畅的，但是却连不上Kafka，特别是在多网卡环境或者云环境上很容易出现，这个其实和Kafka的监听配置有关系。
- 其中最重要的配置就是 listeners 和 advertised.listeners。
  - listeners ：集群启动时监听 listeners 配置的地址
  - advertised.listeners：并将 advertised.listeners配置的地址写到 Zookeeper里面，作为集群元数据的一部分。
- 客户端【生产者/消费者】连接Kafka集群进行操作的流程分为两大步骤：
  - 通过listeners配置的连接信息连接到 Broker（Broker会定期获取并缓存zk中的元数据信息），获取到集群元数据advertised.listeners的连接信息。
  - 通过获取到的集群元数据advertised.listeners信息和Kafka集群进行通信。