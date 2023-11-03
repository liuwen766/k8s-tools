# 一、Kafka启动

##  启停kafka

```shell
# 先启动Zookeeper
bin/zookeeper-server-start.sh config/zookeeper.properties &
# 启动kafka
bin/kafka-server-start.sh config/server.properties &
# 停止kafka
bin/kafka-server-stop.sh config/server.properties &

# 生产者发送消息
bin/kafka-console-producer.sh --broker-list localhost:9092 --topic message-topic

# 消费者消费消息
bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic message-topic --from-beginning
```


# 二、Topic操作

## 1、Topic创建

```shell
# 集群创建topic【推荐】
bin/kafka-topics.sh --create --bootstrap-server localhost:9092 --replication-factor 3 --partitions 2 --topic message-topic
```
> 参数说明：
> - 1、--bootstrap-server：用来指定 Kafka 服务的连接，如有这个参数，那么 --zookeeper 参数可以不需要。--bootstrap-server localhost:9092
> - 2、--zookeeper：废弃，通过zk的方式连接到 Kafka 集群。 --zookeeper localhost:2181
> - 3、--replication-factor：用来指定副本数量，**这里需要注意的是不能大于 Broker 的数量，不提供的话用集群默认值**。
> - 4、--partitions：用来指定分区数量。当创建或修改 Topic 时, ⽤它指定分区数。如果创建时没提供该参数,则⽤集群中默认值。如果修改的时候，这里需要注意不能比之前的小，不然会报错。


## 2、Topic扩容

```shell
# kafka 老版本使用zk方式，这里不推荐使用
bin/kafka-topics.sh --zookeeper localhost:2181 --alter --topic message-topic  --partitions 4
# kafka 2.2 版本以后推荐以下方式，这是推荐使用的方式。 broker_host:port
# 注意：topic一旦创建，partition 只能增加，不能减少。
bin/kafka-topics.sh --bootstrap-server localhost:9092 --alter --topic message-topic --partitions 4
# 也支持正则匹配批量 topic 扩容
bin/kafka-topics.sh --bootstrap-server localhost:9092 --alter --topic ".*?" --replication-factor 3 --partitions 4
```

## 3、Topic删除

```shell
# 删除指定topic
bin/kafka-topics.sh --bootstrap-server localhost:9092  --delete --topic message-topic

# 删除正则匹配的topic
bin/kafka-topics.sh --bootstrap-server localhost:9092  --delete --topic "message-*"

# 删除所有topic【谨慎使用】
bin/kafka-topics.sh --bootstrap-server localhost:9092  --delete --topic ".*?"
```

## 4、Topic列表查看

```shell
bin/kafka-topics.sh --bootstrap-server localhost:9092 --list
# --exclude-internal 排除kafka内部的topic
bin/kafka-topics.sh --bootstrap-server localhost:9092 --list --exclude-internal
# 查询正则匹配的topic
bin/kafka-topics.sh --bootstrap-server localhost:9092 --list --exclude-internal --topic "message-*"
```

## 5、Topic详细信息查看

```shell
bin/kafka-topics.sh --bootstrap-server localhost:9092 --topic message-topic --describe --exclude-internal
bin/kafka-topics.sh --bootstrap-server localhost:9092 --topic "message-*" --describe --exclude-internal
```

## 6、Topic Message数量

```shell
# -time -1 表示获取所有分区的最大位移
bin/kafka-run-class.sh kafka.tools.GetOffsetShell --broker-list localhost:9092 --topic message-topic -time -1
# -time -1 表示获取所有分区的最早位移
bin/kafka-run-class.sh kafka.tools.GetOffsetShell --broker-list localhost:9092 --topic message-topic -time -2
```


# 三、Consumer管理命令

## 1、查看Consumer Group列表

```shell
bin/kafka-consumer-groups.sh --list --bootstrap-server localhost:9092
```
## 2、查看指定Group.id的消费情况

```shell
bin/kafka-consumer-groups.sh --bootstrap-server localhost:9092 --group console-consumer-91424 --describe
```
## 3、删除Group

```shell
bin/kafka-consumer-groups.sh --bootstrap-server localhost:9092 --group console-consumer-91424 --delete
```
## 4、重置Offset

```shell
# 当前group必须处于active状态【即消费者必须在消费】
# 开启消费，自定义消费者组 my-test-group-id-001
bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic message-topic --from-beginning --group my-test-group-id-001

# 查看消费者组 my-test-group-id-001 的Offset信息
[root@mylinux kafka_2.13-2.5.0]# bin/kafka-consumer-groups.sh --bootstrap-server localhost:9092 --group my-test-group-id-001 --describe

GROUP                TOPIC           PARTITION  CURRENT-OFFSET  LOG-END-OFFSET  LAG             CONSUMER-ID                                                          HOST            CLIENT-ID
my-test-group-id-001 message-topic   3          2               2               0               consumer-my-test-group-id-001-1-23da858b-583e-4bf5-a367-8ccef15854df /127.0.0.1      consumer-my-test-group-id-001-1
my-test-group-id-001 message-topic   2          1               1               0               consumer-my-test-group-id-001-1-23da858b-583e-4bf5-a367-8ccef15854df /127.0.0.1      consumer-my-test-group-id-001-1
my-test-group-id-001 message-topic   4          1               1               0               consumer-my-test-group-id-001-1-23da858b-583e-4bf5-a367-8ccef15854df /127.0.0.1      consumer-my-test-group-id-001-1
my-test-group-id-001 message-topic   0          3               3               0               consumer-my-test-group-id-001-1-23da858b-583e-4bf5-a367-8ccef15854df /127.0.0.1      consumer-my-test-group-id-001-1
my-test-group-id-001 message-topic   1          2               2               0               consumer-my-test-group-id-001-1-23da858b-583e-4bf5-a367-8ccef15854df /127.0.0.1      consumer-my-test-group-id-001-1
my-test-group-id-001 message-topic   5          2               2               0               consumer-my-test-group-id-001-1-23da858b-583e-4bf5-a367-8ccef15854df /127.0.0.1      consumer-my-test-group-id-001-1

# 重置 消费者组 my-test-group-id-001 指定 topic 的 Offset 之后，消费者又可以重新消费了
bin/kafka-consumer-groups.sh --bootstrap-server localhost:9092 --group my-test-group-id-001  --reset-offsets --to-earliest  --topic message-topic --execute

# 重置 消费者组 my-test-group-id-001 所有 topic 的Offset
bin/kafka-consumer-groups.sh --bootstrap-server localhost:9092 --group my-test-group-id-001  --reset-offsets --all-topics  --to-earliest --execute
```
