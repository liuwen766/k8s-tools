# Go

Demo代码参考：https://github.com/AliwareMQ/aliware-kafka-demos

- Apache Kafka 的 Golang 客户端库使用的最多的就是 Sarama。



- Go Sarama 生产者：分为 同步生产者SyncProducer 和 异步生产者AsyncProducer。
- 生产者发送消息，返回分区（partition）、消息偏移量（offset）和错误（err）。



- Go Sarama 消费者：分为 消费者Consumer 和 消费者组ConsumerGroup。它们消费的是某个分区（partition）

> 注意：为了防止内存泄露，生产者和消费者都必须调用Close()来关闭，防止它超出范围的时候不能进行垃圾自动回收，从而造成内存泄露，久而久之可能会引起OOM

# Java

Demo代码参考：https://github.com/AliwareMQ/aliware-kafka-demos

- Java 生产者客户端需要以下几个步骤：
  - 1、配置生产者客户端参数；
  - 2、构造KafkaProducer客户端实例；
  - 3、构建待发送消息；
  - 4、发送消息。它分为三种模式：
    - 发后即忘【fire-and-forget】
    - 同步【sync】
    - 异步【async】



- Java 消费者客户端需要以下几个步骤：
  - 1、配置消费者客户端参数；
  - 2、构造KafkaConsumer客户端实例；
  - 3、订阅相应Topic；
  - 4、拉取消息，最后提交消费位移数据Offset。
