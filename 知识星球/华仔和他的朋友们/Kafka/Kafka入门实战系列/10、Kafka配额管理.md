- 在Kafka中，配额是对某种资源的限制。它所能管理和配置的对象有以下三种：
    - 1、用户级别：user 【需要Kafka集群开启身份认证】
    - 2、客户端级别：clientId 【接入 Kafka 集群的生产者/消费者都是一个clientId】
    - 3、用户级别+客户端级别：user + clientId




- 配额管理可以配置的选项：
    - producer_byte_rate：这是生产者单位时间内[每秒]可以发送到Kafka集群中的单台Broker的字节数。
    - consumer_byte_rate：这是消费之单位时间内[每秒]可以发送到Kafka集群中的单台Broker的字节数。



- 客户端级别设置
```shell
# 设置client：consumer-my-test-group-id-001-1
bin/kafka-configs.sh --zookeeper localhost:2181 --alter --add-config 'producer_byte_rate=1024,consumer_byte_rate=2048' --entity-type clients --entity-name consumer-my-test-group-id-001-1
# 查看client：consumer-my-test-group-id-001-1
bin/kafka-configs.sh --zookeeper localhost:2181 --describe --entity-type clients --entity-name consumer-my-test-group-id-001-1
```
- 用户级别设置
```shell
# 设置user：consumer-my-test-user-id-001-1
bin/kafka-configs.sh --zookeeper localhost:2181 --alter --add-config 'producer_byte_rate=1024,consumer_byte_rate=2048' --entity-type users --entity-name consumer-my-test-user-id-001-1
# 查看user：consumer-my-test-user-id-001-1
bin/kafka-configs.sh --zookeeper localhost:2181 --describe --entity-type users --entity-name consumer-my-test-user-id-001-1
```
- 客户端级别+用户级别设置
```shell
bin/kafka-configs.sh --zookeeper localhost:2181 --alter --add-config 'producer_byte_rate=1024,consumer_byte_rate=2048' --entity-type users --entity-name consumer-my-test-user-id-001-1 --entity-type clients --entity-name consumer-my-test-group-id-001-1
```
