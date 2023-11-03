- 在Kafka中，配额是对某种资源的限制。它所能管理和配置的对象有以下三种：
    - 1、用户级别：user 【需要Kafka集群开启身份认证】
    - 2、客户端级别：clientId 【接入 Kafka 集群的生产者/消费者都是一个clientId】
    - 3、用户级别+客户端级别：user + clientId




- 配额管理可以配置的选项：
    - producer_byte_rate
    - consumer_byte_rate



- 客户端级别设置
```shell

```
- 用户级别设置
```shell

```
- 客户端级别+用户级别设置
```shell

```