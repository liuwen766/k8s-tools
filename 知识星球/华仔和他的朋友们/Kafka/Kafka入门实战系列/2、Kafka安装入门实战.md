

- 安装一个单机kafka



- 官网链接：https://kafka.apache.org/downloads
- 下载安装包：https://archive.apache.org/dist/kafka/2.5.0/kafka_2.13-2.5.0.tgz



```shell
# 安装zookeeper
bin/zookeeper-server-start.sh config/zookeeper.properties &

# 启停kafka
bin/kafka-server-start.sh config/server.properties &
bin/kafka-server-stop.sh config/server.properties &

# 创建topic
# 单机创建topic
bin/kaftopics.sh --create --zookeeper localhost:2181 --replication-factor 1 --partitions 1 --topic message
# 集群创建topic【推荐】
bin/kafka-topics.sh --create --bootstrap-server localhost:9092 --replication-factor 3 --partitions 2 --topic message

# 查看topic
bin/kafka-topics.sh --list --zookeeper localhost:2181 
bin/kafka-topics.sh --list --bootstrap-server localhost:9092

# 生产者
bin/kafka-console-producer.sh --broker-list localhost:9092 --topic message

# 消费者
bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic message
bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic message --from-beginning
```


- kafka配置：
```shell
[root@mylinux kafka_2.13-2.5.0]# cat config/server.properties |grep '^[a-z]'
broker.id=0
num.network.threads=3
num.io.threads=8
socket.send.buffer.bytes=102400
socket.receive.buffer.bytes=102400
socket.request.max.bytes=104857600
log.dirs=/tmp/kafka-logs
num.partitions=1
num.recovery.threads.per.data.dir=1
offsets.topic.replication.factor=1
transaction.state.log.replication.factor=1
transaction.state.log.min.isr=1
log.retention.hours=168
log.segment.bytes=1073741824
log.retention.check.interval.ms=300000
zookeeper.connect=localhost:2181
zookeeper.connection.timeout.ms=18000
group.initial.rebalance.delay.ms=0
```
