# Kafka简介



- kafka是一个**分布式的基于发布/订阅模式**的消息队列，主要应用于大数据实时处理。
- 为什么用消息队列——解耦、数据冗余、峰值处理、异步通信。
- kafka基本概念：Producer、Consumer Group、Consumer、Broker、Topic、Partition、Replica、Leader、Follower、Offset、Zookeeper、......
- 
- 一个topic【逻辑概念】对应多个Partition【物理概念】、一个Partition对应多个segment【分片和索引机制】、一个segment对应4个文件【.log文件、.index文件、.snapshot文件、.timeindex文件】
- 消费进度：由于一个partition只能由一个消费者组下的某个consumer消费，所以consumer可以根据自己的<Group ID,Topic,Partition Id>到Broker集群中获取到自己offset，再根据offset值到稀疏索引.index文件中索引到自己在消息.log文件中位置，即获取到了自己的消费进度。

