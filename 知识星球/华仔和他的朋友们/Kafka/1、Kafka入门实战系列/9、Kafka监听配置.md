- 监听相关的参数主要有以下几个：
  - listeners 
    - 监听器列表，用于监听网络请求。
    - listeners地址是用于首次连接的，advertised.listeners的地址是会写到Zk里面，客户端通过 listeners地址建立连接获取该地址信息，然后通过该地址和集群交互。
  - advertised.listeners
    - 用于发布公开的监听器，通过zk发布。
    - 如果未配置则自动使用listeners属性。
    - 如果listeners属性配置为0.0.0.0，则advertised.listeners必须配置。
  - listener.security.protocol.map
    - 监听器名称和安全协议之间的映射关系集合。
    - 格式：监听名称1:安全协议1,监听名称2:安全协议2
    - 安全协议有：
      - plaintext：不需要授权，非加密通道
      - ssl：ssl加密通道
      - sasl_plaintext：使用sasl认证的非加密通道
      - sasl_ssl：使用sasl认证并且ssl加密的通道
  - inter.broker.listener.name
    - Broker集群内互相通信的listener名称。
  - security.inter.broker.protocol
    - 用在代理之间进行通信的安全协议。
  - control.plane.listener.name
    - 用在Controller和Broker之间进行通信的监听器名称。



- 有时候我们会碰到网络是通畅的，但是却连不上Kafka，特别是在多网卡环境或者云环境上很容易出现，这个其实和Kafka的监听配置有关系。
其中最重要的配置就是 listeners 和 advertised.listeners。
  - listeners ：集群启动时监听 listeners 配置的地址
  - advertised.listeners：并将 advertised.listeners配置的地址写到 Zookeeper里面，作为集群元数据的一部分。
- 客户端【生产者/消费者】连接Kafka集群进行操作的流程分为两大步骤：
  - 通过listeners配置的连接信息连接到 Broker（Broker会定期获取并缓存zk中的元数据信息），获取到集群元数据advertised.listeners的连接信息。
  - 通过获取到的集群元数据advertised.listeners信息和Kafka集群进行通信。
