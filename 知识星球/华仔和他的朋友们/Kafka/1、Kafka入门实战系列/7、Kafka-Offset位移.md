
## _consumer_offsets


- Kafka的旧版本重度依赖Zookeeper来实现各种各样的协调管理，包括offset位移的保存。鉴于zk不适合高频写更新，在新版Kafka【0.8版本之后】中，重新设计了大名鼎鼎的consumer_offsets。

- _consumer_offsets：用来再Kafka集群内部保存Kafka Consumer提交的位移信息。它是Kafka自动创建的，类似于普通的Topic。它的消息格式也是Kafka自定义的，人为无法修改。



- consumer_offsets消息结构：它是一个Key-Value键值对。其中 key:<Group ID,Topic,Partition Id>, Value:< Offset >

![img.png](img.png)

- _consumer_offsets的创建：当Kafka集群中的第一个Consumer启动时， Kafka会自动创建consumer_offsets。分区依赖Broker端参数offsets.topic.num.partitions(默认值为50)，因此Kafka会自动创建一个有50个分区的_consumer_offsets。
- 如果consumer_offsets由Kafka自动创建的，那么该Topic的分区数是50，副本数是3，而具体Group的消费情况要存储到哪个Partition，根据abs(Groupld.hashCode()) % NumPartitions来计算的，这样就可以保证Consumer Offset信息与Consumer Group对应的 Coordinator处于同一个Broker节点上。



- 如何指定Kafka Offset位移值，重新消费数据？
  - 修改偏移量Offset
  - 通过consumer.subscribe()指定偏移量Offset
  - 通过auto.offset.reset 指定偏移量Offset
  - 通过指定时间的方式来消费

