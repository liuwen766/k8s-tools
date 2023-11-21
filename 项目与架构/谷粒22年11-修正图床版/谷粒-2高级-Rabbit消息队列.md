

消息队列

作用：下订单后放到消息队列，库存订阅消息队列，库存自己减库存。再如秒杀服务

kafka笔记：[https://blog.csdn.net/hancoder/article/details/107446151](https://blog.csdn.net/hancoder/article/details/107446151)

可以参考之前学其他视频的rabbitMQ笔记：[https://blog.csdn.net/hancoder/article/details/114297652](https://blog.csdn.net/hancoder/article/details/114297652)

消息队列场景：分布式中，我们可以使用远程调用，也可以使用消息队列

## 1、MQ概述：

1．大多应用，可通过消息服务中间件来提升系统异步通信、扩展解耦能力
2．消息服务中两个重要概念：消息代理(messagebroker)和目的地(destination)

当消息发送者发送消息以后，将由消息代理接管，消息代理保证消息传递到指定目的地。

3，消息队列主要有两种形式的目的地

- 1．队列(queue)：点对点消息通信(point-to-point)
- 2．主题(topic)：发布(publish)/订阅(subscribe)消息通信

4．点对点式：

- 消息发送者发送消息，消息代理将亘放入一个队列中，消息接收者从队列中获取消息内容，消息读取后被移出队列
- 消息只有唯一的发送者和接受者，但并不是说只能有一个接收者

5、发布订阅式．
发送者（发布者〕发法消息到主题，多个接收者（订阅者）监听（订阅）这个主题，那么就会在消息到达时同时收到消息

6．JMS(Java Message Service)JAVA消息服务

- 基于JVM消息代理的规范。ActiveMQ、ActiveMQ、HornetMQ是JMS实现
- https://blog.csdn.net/jiuqiyuliang/article/details/46701559

7．AMQP(Advanced Message Queuing Protocol)

- 高级消息队列协议，也是一个消息代理的规范，兼容JMS
- RabbitMQ是AMQP的实现

8．Spring支持

- `spring-jms`提供了对`JMS`的支持
- `spring-rabbit`提供了对`AMQP`的支持
- 需要（onnectionFactory的实现来连接消息代理
- 提供`JmsTemplate`、`RabbitTemplate`来发送消息
- @JmsListener(JMS)、@RabbitListener(AMQP)注解在方法上监听消息代理发布的消息
- @EnabIeJms、@EnableRabbit开启支持

9．SpringBoot自动配置

- JmsAutoConfiguration
- RabbitAutoConfiguration

10、市面的MQ产品
ActiveMQ、RabbitMQ、RocketMQ、Kafka

## 2、核心概念

RabbitMQ简介：RabbitMQ是—个由erlang语言开发的AMQP(AdvanvedMessageQueueProtocoI)的开源实现。

核心概念：

- **Message**：消息，消息是不具名的，它由消息头和消息体组成。消息体是不透明的，而消息头则由一系列的可选属性组成，这些属性routing-key(路由键)、priority（相对于其他消息的优先权）、delivery-mode（指出该消息可能需要持久性存储
- **Publisher**：消息的生产者，也是一个向交换器发布消息的客户端应用程序
- **Exchange**：交唤器，用来接收生产者发送的消息并**将这些消息路由给服务器中的队列**。
  - Exchange有4种类型：direct默认fanout,topic,和headers,不同类型的Exchange转发消息的策略有所区别
- consumer：从消息队列中取得消息的客户端应用程序。
- VirtualHost：虚拟主机，表示一批交换器、消息队列和相关对象。虚拟主机是共享相同的身份认证和加密环境的独立服务器域。每个vhost本质上就是一个mini版的RabbitMQ服务器，拥有自己的队列、交换器、绑定和权限机制。vhost是AMQP概念的基础，必须在连接时指定，RabbitMQ默认的vhost是/。
- Broker：表示消息队列服务器实体

## 3、docker安装

```bash
docker run -d --name rabbitmq -p 5671:5671 -p 5672:5672 -p 4369:4369 -p 25672:25672 -p 15671:15671 -p 15672:15672 rabbitmq:management

docker update rabbitmq --restart=always
```

- 4369 25672 Ealang发现&集群端口
- 5672 5671 AMQP端口
- 15672 web管理后台端口guest
- 61613 61614 STOMP协议端口
- 1883 8883 MQTT协议端口

| Protocol        | Bound to | Port  |
| :-------------- | :------- | :---- |
| amqp            | ::       | 5672  |
| clustering      | ::       | 25672 |
| http            | ::       | 15672 |
| http/prometheus | ::       | 15692 |

默认的交换机：

|                                                              |         |          |                 |                  |      |
| :----------------------------------------------------------- | :------ | :------- | :-------------- | :--------------- | :--- |
| Name                                                         | Type    | Features | Message rate in | Message rate out | +/-  |
| [(AMQP default)](http://192.168.56.10:15672/#/exchanges/%2F/amq.default) | direct  | D        |                 |                  |      |
| [amq.direct](http://192.168.56.10:15672/#/exchanges/%2F/amq.direct) | direct  | D        |                 |                  |      |
| [amq.fanout](http://192.168.56.10:15672/#/exchanges/%2F/amq.fanout) | fanout  | D        |                 |                  |      |
| [amq.headers](http://192.168.56.10:15672/#/exchanges/%2F/amq.headers) | headers | D        |                 |                  |      |
| [amq.match](http://192.168.56.10:15672/#/exchanges/%2F/amq.match) | headers | D        |                 |                  |      |
| [amq.rabbitmq.trace](http://192.168.56.10:15672/#/exchanges/%2F/amq.rabbitmq.trace) | topic   | D I      |                 |                  |      |
| [amq.topic](http://192.168.56.10:15672/#/exchanges/%2F/amq.topic) | topic   | D        |                 |                  |      |

## 4、运行机制

AMQP中消息的路由过程和Java开发者熟悉的JMS存在一些差别，AMQP中增加了Exchange和Binding的角色。

生产者把消息发布到Exchange上，消息最终到达队列并消费者接收，而Binding决定交换器的消息应该发送哪个队列。

![img](https://upload-images.jianshu.io/upload_images/5015984-7fd73af768f28704.png)

## 交换机Exchange 类型

Exchange分发消息时根据类型的不同分发策略有区别，目前共四种类型：direct直接、fanout扇出、topic主题（发布订阅）、headers 。

- 14是点对点，23是发布订阅。4性能比较低

headers 匹配 AMQP 消息的 header 而不是路由键，此外 headers 交换器和 direct 交换器完全一致，但性能差很多，目前几乎用不到了，所以直接看另外三种类型：

#### 1) direct

![img](https:////upload-images.jianshu.io/upload_images/5015984-13db639d2c22f2aa.png)

direct 交换器

消息中的路由键（routing key）如果和 Binding 中的 binding key 一致， 交换器就将消息发到对应的队列中。路由键与队列名完全匹配，如果一个队列绑定到交换机要求路由键为“dog”，则只转发 routing key 标记为“dog”的消息，不会转发“dog.puppy”，也不会转发“dog.guard”等等。它是**完全匹配**、单播的模式。

#### 2) fanout

![img](https:////upload-images.jianshu.io/upload_images/5015984-2f509b7f34c47170.png?imageMogr2/auto-orient/strip|imageView2/2/w/463/format/webp)

fanout 交换器

每个发到 fanout 类型交换器的消息都会分到所有绑定的队列上去。fanout 交换器不处理路由键，只是简单的将队列绑定到交换器上，每个发送到交换器的消息都会被转发到与该交换器绑定的==所有==队列上。很像子网广播，每台子网内的主机都获得了一份复制的消息。fanout 类型转发消息是最快的。

#### 3) topic

![img](https:////upload-images.jianshu.io/upload_images/5015984-275ea009bdf806a0.png)

topic 交换器

topic 交换器通过模式匹配分配消息的路由键属性，将路由键和某个模式进行匹配，此时队列需要绑定到一个模式上。它将路由键和绑定键的字符串切分成单词，这些单词之间用点隔开。它同样也会识别两个通配符：符号“#”和符号“*”。

`#`匹配0个或多个单词，

`*`匹配不多不少一个单词。

接下来是一些可视化操作，没什么好记的。

# RabbitMQ 运行和管理

启动：找到安装后的 RabbitMQ 所在目录下的 sbin 目录，可以看到该目录下有6个以 rabbitmq 开头的可执行文件，直接执行 rabbitmq-server 即可，下面将 RabbitMQ 的安装位置以 . 代替，启动命令就是：

```undefined
./sbin/rabbitmq-server
```

启动正常的话会看到一些启动过程信息和最后的 completed with 7 plugins，这也说明启动的时候默认加载了7个插件。

![img](https:////upload-images.jianshu.io/upload_images/5015984-1392cdc83b0d8341.png)

正常启动

后台启动：如果想让 RabbitMQ 以守护程序的方式在后台运行，可以在启动的时候加上 -detached 参数：

```undefined
./sbin/rabbitmq-server -detached
```

查询服务器状态
sbin 目录下有个特别重要的文件叫 rabbitmqctl ，它提供了 RabbitMQ 管理需要的几乎一站式解决方案，绝大部分的运维命令它都可以提供。
查询 RabbitMQ 服务器的状态信息可以用参数 status ：

```undefined
./sbin/rabbitmqctl status
```

该命令将输出服务器的很多信息，比如 RabbitMQ 和 Erlang 的版本、OS 名称、内存等等

关闭 RabbitMQ 节点
我们知道 RabbitMQ 是用 Erlang 语言写的，在Erlang 中有两个概念：节点和应用程序。节点就是 Erlang 虚拟机的每个实例，而多个 Erlang 应用程序可以运行在同一个节点之上。节点之间可以进行本地通信（不管他们是不是运行在同一台服务器之上）。比如一个运行在节点A上的应用程序可以调用节点B上应用程序的方法，就好像调用本地函数一样。如果应用程序由于某些原因奔溃，Erlang 节点会自动尝试重启应用程序。
如果要关闭整个 RabbitMQ 节点可以用参数 stop ：

```undefined
./sbin/rabbitmqctl stop
```

它会和本地节点通信并指示其干净的关闭，也可以指定关闭不同的节点，包括远程节点，只需要传入参数 -n ：

```dart
./sbin/rabbitmqctl -n rabbit@server.example.com stop 
```

-n node 默认 node 名称是 rabbit@server ，如果你的主机名是 [server.example.com](https://link.jianshu.com?t=http://server.example.com) ，那么 node 名称就是 [rabbit@server.example.com](https://link.jianshu.com?t=mailto:rabbit@server.example.com) 。

关闭 RabbitMQ 应用程序
如果只想关闭应用程序，同时保持 Erlang 节点运行则可以用 stop_app：

```undefined
./sbin/rabbitmqctl stop_app
```

这个命令在后面要讲的集群模式中将会很有用。

启动 RabbitMQ 应用程序

```undefined
./sbin/rabbitmqctl start_app
```

重置 RabbitMQ 节点

```undefined
./sbin/rabbitmqctl reset
```

该命令将清除所有的队列。

查看已声明的队列

```undefined
./sbin/rabbitmqctl list_queues
```

查看交换器

```undefined
./sbin/rabbitmqctl list_exchanges
```

该命令还可以附加参数，比如列出交换器的名称、类型、是否持久化、是否自动删除：



```rust
./sbin/rabbitmqctl list_exchanges name type durable auto_delete
```

查看绑定

```undefined
./sbin/rabbitmqctl list_bindings
```

# Java 客户端访问

RabbitMQ 支持多种语言访问，以 Java 为例看下一般使用 RabbitMQ 的步骤。

pom中添加依赖

```xml
<dependency>
    <groupId>com.rabbitmq</groupId>
    <artifactId>amqp-client</artifactId>
    <version>4.1.0</version>
</dependency>
```

消息生产者

```dart
package org.study.rabbitmq;
import com.rabbitmq.client.Channel;
import com.rabbitmq.client.Connection;
import com.rabbitmq.client.ConnectionFactory;
import java.io.IOException;
import java.util.concurrent.TimeoutException;
public class Producer {

    public static void main(String[] args) throws IOException, TimeoutException {
        //创建连接工厂
        ConnectionFactory factory = new ConnectionFactory();
        factory.setUsername("guest");
        factory.setPassword("guest");
        //设置 RabbitMQ 地址
        factory.setHost("localhost");
        //建立到代理服务器到连接
        Connection conn = factory.newConnection();
        //获得信道
        Channel channel = conn.createChannel();
        //声明交换器
        String exchangeName = "hello-exchange";
        channel.exchangeDeclare(exchangeName, "direct", true);

        String routingKey = "hola";
        //发布消息
        byte[] messageBodyBytes = "quit".getBytes();
        channel.basicPublish(exchangeName, routingKey, null, messageBodyBytes);

        channel.close();
        conn.close();
    }
}
```

消息消费者

```dart
package org.study.rabbitmq;
import com.rabbitmq.client.*;
import java.io.IOException;
import java.util.concurrent.TimeoutException;
public class Consumer {

    public static void main(String[] args) throws IOException, TimeoutException {
        ConnectionFactory factory = new ConnectionFactory();
        factory.setUsername("guest");
        factory.setPassword("guest");
        factory.setHost("localhost");
        //建立到代理服务器到连接
        Connection conn = factory.newConnection();
        //获得信道
        final Channel channel = conn.createChannel();
        //声明交换器
        String exchangeName = "hello-exchange";
        channel.exchangeDeclare(exchangeName, "direct", true);
        //声明队列
        String queueName = channel.queueDeclare().getQueue();
        String routingKey = "hola";
        //绑定队列，通过键 hola 将队列和交换器绑定起来
        channel.queueBind(queueName, exchangeName, routingKey);

        while(true) {
            //消费消息
            boolean autoAck = false;
            String consumerTag = "";
            channel.basicConsume(queueName, autoAck, consumerTag, new DefaultConsumer(channel) {
                @Override
                public void handleDelivery(String consumerTag,
                                           Envelope envelope,
                                           AMQP.BasicProperties properties,
                                           byte[] body) throws IOException {
                    String routingKey = envelope.getRoutingKey();
                    String contentType = properties.getContentType();
                    System.out.println("消费的路由键：" + routingKey);
                    System.out.println("消费的内容类型：" + contentType);
                    long deliveryTag = envelope.getDeliveryTag();
                    //确认消息
                    channel.basicAck(deliveryTag, false);
                    System.out.println("消费的消息体内容：");
                    String bodyStr = new String(body, "UTF-8");
                    System.out.println(bodyStr);

                }
            });
        }
    }
}
```

启动 RabbitMQ 服务器

```undefined
./sbin/rabbitmq-server
```

先运行 Consumer ：这样当生产者发送消息的时候能在消费者后端看到消息记录。

运行 Producer：发布一条消息，在 Consumer 的控制台能看到接收的消息：

![img](https:////upload-images.jianshu.io/upload_images/5015984-6f2d0168cfc2878d.png?imageMogr2/auto-orient/strip|imageView2/2/w/1200/format/webp)

Consumer 控制台

# RabbitMQ 集群

RabbitMQ 最优秀的功能之一就是内建集群，这个功能设计的目的是允许消费者和生产者在节点崩溃的情况下继续运行，以及通过添加更多的节点来线性扩展消息通信吞吐量。RabbitMQ 内部利用 Erlang 提供的分布式通信框架 OTP 来满足上述需求，使客户端在失去一个 RabbitMQ 节点连接的情况下，还是能够重新连接到集群中的任何其他节点继续生产、消费消息。

##### RabbitMQ 集群中的一些概念

RabbitMQ 会始终记录以下四种类型的内部元数据：

1. 队列元数据
    包括队列名称和它们的属性，比如是否可持久化，是否自动删除
2. 交换器元数据
    交换器名称、类型、属性
3. 绑定元数据
    内部是一张表格记录如何将消息路由到队列
4. vhost 元数据
    为 vhost 内部的队列、交换器、绑定提供命名空间和安全属性

在单一节点中，RabbitMQ 会将所有这些信息存储在内存中，同时将标记为可持久化的队列、交换器、绑定存储到硬盘上。存到硬盘上可以确保队列和交换器在节点重启后能够重建。而在集群模式下同样也提供两种选择：存到硬盘上（独立节点的默认设置），存在内存中。

如果在集群中创建队列，集群只会在单个节点而不是所有节点上创建完整的队列信息（元数据、状态、内容）。结果是只有队列的所有者节点知道有关队列的所有信息，因此当集群节点崩溃时，该节点的队列和绑定就消失了，并且任何匹配该队列的绑定的新消息也丢失了。还好RabbitMQ 2.6.0之后提供了镜像队列以避免集群节点故障导致的队列内容不可用。

RabbitMQ 集群中可以共享 user、vhost、exchange等，所有的数据和状态都是必须在所有节点上复制的，例外就是上面所说的消息队列。RabbitMQ 节点可以动态的加入到集群中。

当在集群中声明队列、交换器、绑定的时候，这些操作会直到所有集群节点都成功提交元数据变更后才返回。集群中有内存节点和磁盘节点两种类型，内存节点虽然不写入磁盘，但是它的执行比磁盘节点要好。内存节点可以提供出色的性能，磁盘节点能保障配置信息在节点重启后仍然可用，那集群中如何平衡这两者呢？

RabbitMQ 只要求集群中至少有一个磁盘节点，所有其他节点可以是内存节点，当节点加入火离开集群时，它们必须要将该变更通知到至少一个磁盘节点。如果只有一个磁盘节点，刚好又是该节点崩溃了，那么集群可以继续路由消息，但不能创建队列、创建交换器、创建绑定、添加用户、更改权限、添加或删除集群节点。换句话说集群中的唯一磁盘节点崩溃的话，集群仍然可以运行，但知道该节点恢复，否则无法更改任何东西。

##### RabbitMQ 集群配置和启动

如果是在一台机器上同时启动多个 RabbitMQ 节点来组建集群的话，只用上面介绍的方式启动第二、第三个节点将会因为节点名称和端口冲突导致启动失败。所以在每次调用 rabbitmq-server 命令前，设置环境变量 RABBITMQ_NODENAME 和 RABBITMQ_NODE_PORT 来明确指定唯一的节点名称和端口。下面的例子端口号从5672开始，每个新启动的节点都加1，节点也分别命名为test_rabbit_1、test_rabbit_2、test_rabbit_3。

启动第1个节点：

```undefined
RABBITMQ_NODENAME=test_rabbit_1 RABBITMQ_NODE_PORT=5672 ./sbin/rabbitmq-server -detached
```

启动第2个节点：

```undefined
RABBITMQ_NODENAME=test_rabbit_2 RABBITMQ_NODE_PORT=5673 ./sbin/rabbitmq-server -detached
```

启动第2个节点前建议将 RabbitMQ 默认激活的插件关掉，否则会存在使用了某个插件的端口号冲突，导致节点启动不成功。

现在第2个节点和第1个节点都是独立节点，它们并不知道其他节点的存在。集群中除第一个节点外后加入的节点需要获取集群中的元数据，所以要先停止 Erlang 节点上运行的 RabbitMQ 应用程序，并重置该节点元数据，再加入并且获取集群的元数据，最后重新启动 RabbitMQ 应用程序。

停止第2个节点的应用程序：

```undefined
./sbin/rabbitmqctl -n test_rabbit_2 stop_app
```

重置第2个节点元数据：

```undefined
./sbin/rabbitmqctl -n test_rabbit_2 reset
```

第2节点加入第1个节点组成的集群：

```dart
./sbin/rabbitmqctl -n test_rabbit_2 join_cluster test_rabbit_1@localhost
```

启动第2个节点的应用程序

```undefined
./sbin/rabbitmqctl -n test_rabbit_2 start_app
```

第3个节点的配置过程和第2个节点类似：

```dart
RABBITMQ_NODENAME=test_rabbit_3 RABBITMQ_NODE_PORT=5674 ./sbin/rabbitmq-server -detached

./sbin/rabbitmqctl -n test_rabbit_3 stop_app

./sbin/rabbitmqctl -n test_rabbit_3 reset

./sbin/rabbitmqctl -n test_rabbit_3 join_cluster test_rabbit_1@localhost

./sbin/rabbitmqctl -n test_rabbit_3 start_app
```

##### RabbitMQ 集群运维

停止某个指定的节点，比如停止第2个节点：

```undefined
RABBITMQ_NODENAME=test_rabbit_2 ./sbin/rabbitmqctl stop
```

查看节点3的集群状态：

```undefined
./sbin/rabbitmqctl -n test_rabbit_3 cluster_status
```

离线笔记均为markdown格式，图片也是云图，10多篇笔记20W字，压缩包仅500k，推荐使用typora阅读。也可以自己导入有道云笔记等软件中

阿里云图床现在**每周得几十元充值**，都要自己往里搭了，麻烦不要散播与转发

![](https://i0.hdslb.com/bfs/album/ff3fb7e24f05c6a850ede4b1f3acc54312c3b0c6.png)

打赏后请主动发支付信息到邮箱  553736044@qq.com  ，上班期间很容易忽略收账信息，邮箱回邮基本秒回

禁止转载发布，禁止散播，若发现大量散播，将对本系统文章图床进行重置处理。

技术人就该干点技术人该干的事



如果帮到了你，留下赞吧，谢谢支持