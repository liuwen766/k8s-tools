# RabbitMQ高级特性

•消息可靠性投递：生产者到MQ中间件

•Consumer ACK：MQ中间件到消费端

•消费端限流

•TTL

•死信队列

•延迟队列

•日志与监控

•消息可靠性分析与追踪

•管理

**RabbitMQ应用问题**

•消息可靠性保障

•消息幂等性处理

**RabbitMQ集群搭建**

•RabbitMQ高可用集群

## 1. RabbitMQ 高级特性

#### 1.1 生产端确认

在使用 RabbitMQ 的时候，作为消息发送方希望杜绝任何消息丢失或者投递失败场景。RabbitMQ 为我们提供了两种方式用来控制消息的投递可靠性模式。

- ==confirm 确认模式==
- ==return 退回模式==



rabbitmq 整个消息投递的路径为：

producer--->rabbitmq-broker--->exchange--->queue--->consumer

- 生产者到交换机：则会返回一个异步 `confirmCallback() `。都会执行，返回false就失败，可以在生产者端处理
- 交换机到队列：投递失败则会返回一个 `returnCallback() `。

我们将利用这两个 callback 控制消息的可靠性投递

##### 生产确认ConfirmCallback

打开`spring.rabbitmq.publisher-confirms="true"` 或者改xml

```xml
<!-- 定义rabbitmq connectionFactory -->
<rabbit:connection-factory id="connectionFactory" host="${rabbitmq.host}"
                           port="${rabbitmq.port}"
                           username="${rabbitmq.username}"
                           password="${rabbitmq.password}"
                           virtual-host="${rabbitmq.virtual-host}"
                           publisher-confirms="true"
                           />
```

在rabbitTemplate中设置回调函数

消息只要被broker接收就会执行ConfirmCallback()，如果是cluster模型，需要所有broker收到才会调用ConfirmCallback

被broker接收到只能表示message已经到达服务器，并不能保证消息一定被投递到目标queue里。所以需要用到后面的returnCallback()

> 怎么设置呢？可以在配置类中集合`@PostConstruct`注解方法

```java
@Autowired
private RabbitTemplate rabbitTemplate;

/**
     * 确认模式：
     * 步骤：
     * 1. 确认模式开启：ConnectionFactory中开启publisher-confirms="true"
     * 2. 在rabbitTemplate定义ConfirmCallBack回调函数
     */
@Test
public void testConfirm() {

    //2. 生产者端定义回调
    rabbitTemplate.setConfirmCallback(
        new RabbitTemplate.ConfirmCallback() {

            @Override
            public void confirm(CorrelationData correlationData, // 相关配置信息
                                boolean ack, // 交换机 是否成功收到了消息。true 成功，false代表失败
                                String cause) { // 失败原因
                if (ack) { //接收成功
                    System.out.println("接收成功消息" + cause);
                } else {  //接收失败。如何失败：交换机等配置错误
                    System.out.println("接收失败消息" + cause);
                    //做一些处理，让消息再次发送。
                }
            }
        });

    //3. 发送消息
    rabbitTemplate.convertAndSend("test_exchange_confirm111", // 交换机
                                  "confirm", // 路由key
                                  "message confirm...."); // 消息
}
```



##### 队列回退：ReturnCallback

- spring.rabbitmq.publister-returns=true
- spring.rabbitmq.template.mandatory=true  抵达队列之后优先执行异步回调。则会将消息退回给producer。并执行回调函数returnedMessage。

```xml
<!-- 定义rabbitmq connectionFactory -->
<rabbit:connection-factory id="connectionFactory"
                           host="${rabbitmq.host}"
                           port="${rabbitmq.port}"
                           username="${rabbitmq.username}"
                           password="${rabbitmq.password}"
                           virtual-host="${rabbitmq.virtual-host}"
                           publisher-confirms="true"
                           publisher-returns="true"
                           />
```

先设置rabbitTemplate.setMandatory(true);  抵达队列之后优先执行异步回调

然后处理消息回调的`replyText`

交换器无法根据自身类型和路由键找到符合条件队列时：

- 设置mandatory = true，代表返回消息给生产者；
- 设置mandatory = false，代表直接丢弃

```java
@Autowired
private RabbitTemplate rabbitTemplate;

/**
     * 回退模式： 当消息发送给Exchange后，Exchange路由到Queue失败是 才会执行 ReturnCallBack
     * 步骤：
     * 1. 开启回退模式:publisher-returns="true"
     * 2. 设置ReturnCallBack
     * 3. 设置Exchange处理消息的模式：
     * 1. 如果消息没有路由到Queue，则丢弃消息（默认）
     * 2. 如果消息没有路由到Queue，返回给消息发送方ReturnCallBack
     */
@Test
public void testReturn() {

    // 设置交换机处理失败消息的模式 // true代表返回给生产者，false代表直接丢弃
    rabbitTemplate.setMandatory(true);

    // 2.设置ReturnCallBack ，用于交换机到队列
    rabbitTemplate.setReturnCallback(
            new RabbitTemplate.ReturnCallback() {
            /** 则会将消息退回给producer。并执行回调函数returnedMessage。 */
            @Override
            public void returnedMessage(Message message, // 消息对象
                                        int replyCode, // 错误码
                                        String replyText, // 错误信息
                                        String exchange, // 交换机
                                        String routingKey) { // 路由键
                System.out.println("return 执行了....");

                System.out.println(message);
                System.out.println(replyCode);
                System.out.println(replyText);
                System.out.println(exchange);
                System.out.println(routingKey);
                //处理
            }
        });

    //3. 发送消息
    rabbitTemplate.convertAndSend("test_exchange_confirm", "confirm", "message confirm....");
}
```





**消息的可靠投递小结**

- 设置ConnectionFactory的`publisher-confirms="true" `开启 确认模式。
- 使用`rabbitTemplate.setConfirmCallback`设置回调函数。当消息发送到exchange后回调confirm方法。在方法中判断ack，如果为true，则发送成功，如果为false，则发送失败，需要处理。
- 
- 设置ConnectionFactory的`publisher-returns="true" `开启 退回模式。
- 使用`rabbitTemplate.setReturnCallback`设置退回函数，当消息从exchange路由到queue失败后，如果设置了rabbitTemplate.setMandatory(true)参数，则会将消息退回给producer。并执行回调函数returnedMessage。
- 在RabbitMQ中也提供了事务机制，但是性能较差，此处不做讲解。

使用channel下列方法，完成事务控制：

- `txSelect()` ：用于将当前channel设置成transaction模式
- `txCommit()`：用于提交事务
- `txRollback()`：用于回滚事务

#### 1.2 消费端确认

ack指Acknowledge，确认。 表示消费端收到消息后的确认方式。

消费者从消息队列消费到消息后有三种确认方式：

- 自动确认：`acknowledge="none"` 。不管处理成功与否，业务处理异常也不管，进入channel通道就算接收到
- 手动确认：`acknowledge="manual"`  。可以解决业务异常的情况
- 根据异常情况确认：`acknowledge="auto"`，（这种方式使用麻烦，不作讲解）

其中自动确认是指，当消息一旦被Consumer接收到，则自动确认收到，并将相应 message 从 RabbitMQ 的消息缓存中移除。但是在实际业务处理中，很可能消息接收到，业务处理出现异常，那么该消息就会丢失。如果设置了手动确认方式，则需要在业务处理成功后，调用`channel.basicAck()`，手动签收，如果出现异常，则调用`channel.basicNack()`方法，让其自动重新发送消息。

小结：

- 在`<rabbit:listener-container>`标签中设置acknowledge属性，设置ack方式 `none：自动确认`，`manual：手动确认`
- 如果在消费端正常消费完，则调用`channel.basicAck(deliveryTag他是通道内按序自增的,false批量确认);`方法确认签收消息
- 如果出现异常，则在catch中调用 `basicNack()`或 `basicReject()`，拒绝消息，让MQ重新发送消息。
- 如果没有调用ack/nack方法，broker认为此消息正在被处理，不会投递给别人，此时客户端断开，消息不会被broker移除，会投递给别人

> 下面演示的是spring的方式，在springboot中用注解消费也是一样的

##### AckListener

什么都不设置就是自动签收。下面讲解的是手动签收

```xml
<!--定义监听器容器-->
<rabbit:listener-container connection-factory="connectionFactory"
                           acknowledge="manual" >
    <rabbit:listener ref="ackListener" queue-names="test_queue_confirm"></rabbit:listener>
</rabbit:listener-container>
```



```java
/**
 * Consumer ACK机制：
 *  1. 设置手动签收。acknowledge="manual"
 *  2. 让监听器类实现ChannelAwareMessageListener接口
 *  3. 如果消息成功处理，则调用channel的 basicAck()签收
 *  4. 如果消息处理失败，则调用channel的basicNack()拒绝签收，broker重新发送给consumer
 */
@Component
public class AckListener implements ChannelAwareMessageListener {

    @Override
    public void onMessage(Message message, Channel channel) throws Exception {
        long deliveryTag = message.getMessageProperties().getDeliveryTag();

        try {
            //1.接收转换消息
            System.out.println(new String(message.getBody()));
            //2. 处理业务逻辑
            System.out.println("处理业务逻辑...");
            
            int i = 3/0;//出现错误
            
            //3. 手动签收
            channel.basicAck(deliveryTag,true);
            
        } catch (Exception e) {
            //e.printStackTrace();

            //4.拒绝签收
            channel.basicNack(deliveryTag,
                              true,
                              true); // ：requeue：重回队列。如果设置为true，则消息重新回到queue，broker会重新发送该消息给消费端
            
            //channel.basicReject(deliveryTag,true);
        }
    }
}
```

```xml

<!--定义监听器容器-->
<rabbit:listener-container connection-factory="connectionFactory" acknowledge="manual" prefetch="1" >
    <!-- <rabbit:listener ref="ackListener" queue-names="test_queue_confirm"></rabbit:listener>-->
    <!-- <rabbit:listener ref="qosListener" queue-names="test_queue_confirm"></rabbit:listener>-->
    <!--定义监听器，监听正常队列-->
    <!--<rabbit:listener ref="dlxListener" queue-names="test_queue_dlx"></rabbit:listener>-->

    <!--延迟队列效果实现：  一定要监听的是 死信队列！！！-->
    <rabbit:listener ref="orderListener" queue-names="order_queue_dlx"></rabbit:listener>
</rabbit:listener-container>
```



````java
@Component
public class DlxListener implements ChannelAwareMessageListener {

    @Override
    public void onMessage(Message message, Channel channel) throws Exception {
        long deliveryTag = message.getMessageProperties().getDeliveryTag();

        try {
            //1.接收转换消息
            System.out.println(new String(message.getBody()));
            //2. 处理业务逻辑
            System.out.println("处理业务逻辑...");
            int i = 3/0;//出现错误
            //3. 手动签收
            channel.basicAck(deliveryTag,true);
        } catch (Exception e) {
            //e.printStackTrace();
            System.out.println("出现异常，拒绝接受");
            //4.拒绝签收，不重回队列 requeue=false
            channel.basicNack(deliveryTag,true,false);
        }
    }
}

````

```java
@Component
public class OrderListener implements ChannelAwareMessageListener {

    @Override
    public void onMessage(Message message, Channel channel) throws Exception {
        long deliveryTag = message.getMessageProperties().getDeliveryTag();

        try {
            //1.接收转换消息
            System.out.println(new String(message.getBody()));

            //2. 处理业务逻辑
            System.out.println("处理业务逻辑...");
            System.out.println("根据订单id查询其状态...");
            System.out.println("判断状态是否为支付成功");
            System.out.println("取消订单，回滚库存....");
            //3. 手动签收
            channel.basicAck(deliveryTag,true);
        } catch (Exception e) {
            //e.printStackTrace();
            System.out.println("出现异常，拒绝接受");
            //4.拒绝签收，不重回队列 requeue=false
            channel.basicNack(deliveryTag,true,false);
        }
    }
}
```

##### 消息可靠性总结：

1.持久化

- exchange要持久化
- queue要持久化
- message要持久化

2.生产方确认Confirm

3.消费方确认Ack

4.Broker高可用



#### **1.3** **消费端限流**

- 在`<rabbit:listener-container> `中配置 prefetch属性设置消费端一次拉取多少消息
- 消费端的确认模式一定为手动确认。acknowledge="manual"

```xml
<!--定义监听器容器-->
<rabbit:listener-container connection-factory="connectionFactory" 
                           acknowledge="manual"   手动确认
                           prefetch="1" >  一次拉取1条
    <rabbit:listener ref="qosListener" queue-names="test_queue_confirm"></rabbit:listener>
</rabbit:listener-container>
/**
 * Consumer 限流机制
 *  1. 确保ack机制为手动确认。
 *  2. listener-container配置属性
 *      perfetch = 1,表示消费端每次从mq拉去一条消息来消费，直到手动确认消费完毕后，才会继续拉去下一条消息。
 */
```



```java
/**
 * Consumer 限流机制
 *  1. 确保ack机制为手动确认。
 *  2. listener-container配置属性
 *      perfetch = 1,表示消费端每次从mq拉去一条消息来消费，直到手动确认消费完毕后，才会继续拉去下一条消息。
 */
@Component
public class QosListener implements ChannelAwareMessageListener {

    @Override
    public void onMessage(Message message, Channel channel) throws Exception {

        Thread.sleep(1000);
        //1.获取消息
        System.out.println(new String(message.getBody()));

        //2. 处理业务逻辑

        //3. 签收
        channel.basicAck(message.getMessageProperties().getDeliveryTag(),true);
    }
}
```



#### **1.4 TTL**

- TTL 全称 Time To Live（存活时间/过期时间）。
- 当消息到达存活时间后，还没有被消费，会被自动清除。
- RabbitMQ可以对消息设置过期时间，也可以对整个队列（Queue）设置过期时间。

可以在管理台新建队列、交换机，绑定



- 设置队列过期时间使用参数：`x-message-ttl`，单位：ms(毫秒)，会对整个队列消息统一过期。
- 设置消息过期时间使用参数：`expiration`。单位：ms(毫秒)，当该消息在队列头部时（消费时），会单独判断这一消息是否过期。
- 如果两者都进行了设置，以时间短的为准。

```xml
<!--ttl-->
<rabbit:queue name="test_queue_ttl" id="test_queue_ttl">
    <!--设置queue的参数-->
    <rabbit:queue-arguments>
        <!--x-message-ttl指队列的过期时间-->
        <entry key="x-message-ttl" value="100000" value-type="java.lang.Integer"></entry>
    </rabbit:queue-arguments>
</rabbit:queue>

<rabbit:topic-exchange name="test_exchange_ttl" >
    <rabbit:bindings>
        <rabbit:binding pattern="ttl.#" queue="test_queue_ttl"></rabbit:binding>
    </rabbit:bindings>
</rabbit:topic-exchange>
```

编写消息超时后的处理逻辑，即**消息后置处理器**，发送消息时作为参数传入

```JAVA
/**
     * TTL:过期时间
     *  1. 队列统一过期
     *  2. 消息单独过期
     *
     * 如果设置了消息的过期时间，也设置了队列的过期时间，它以时间短的为准。
     * 队列过期后，会将队列所有消息全部移除。
     * 消息过期后，只有消息在队列顶端，才会判断其是否过期(移除掉)
     */
@Test
public void testTtl() {
    /*  for (int i = 0; i < 10; i++) {
            // 发送消息
            rabbitTemplate.convertAndSend("test_exchange_ttl", "ttl.hehe", "message ttl....");
        }*/

    
    // 消息后处理对象，设置一些消息的参数信息
    MessagePostProcessor messagePostProcessor = new MessagePostProcessor() {
        @Override
        public Message postProcessMessage(Message message) throws AmqpException {
            //1.设置message的信息
            message.getMessageProperties().setExpiration("5000");//消息的过期时间
            //2.返回该消息
            return message;
        }
    };
    //消息单独过期
    for (int i = 0; i < 10; i++) {
        if(i == 5){
            //消息单独过期
            rabbitTemplate.convertAndSend("test_exchange_ttl", "ttl.hehe", "message ttl....",messagePostProcessor);
        }else{
            //不过期的消息
            rabbitTemplate.convertAndSend("test_exchange_ttl", "ttl.hehe", "message ttl....");
        }
    }
}

```



#### **1.5** **死信队列**

死信队列，英文缩写：DLX 。Dead Letter Exchange（死信交换机，因为其他MQ产品中没有交换机的概念），当消息成为Dead message后，可以被重新发送到另一个交换机，这个交换机就是DLX。

比如消息队列的消息过期，如果绑定了死信交换器，那么该消息将发送给死信交换机

![](https://i0.hdslb.com/bfs/album/3209d574afbe6853ddaeff898c0b7734873392c4.png)

当消息成为死信后，如果该队列绑定了死信交换机，则消息会被死信交换机重新路由到死信队列

**消息成为死信的三种情况：**

- \1. 队列消息长度到达限制；
- \2. 消费者拒接消费消息（`basicNack/basicReject`），并且不把消息重新放入原目标队列（requeue=false；不重回队列）
- \3. 原队列存在消息过期设置，消息到达超时时间未被消费；

**队列绑定死信交换机：**

给队列设置参数： `x-dead-letter-exchange` 和 `x-dead-letter-routing-key`



```xml
<!--
        死信队列：
            1. 声明正常的队列(test_queue_dlx)和交换机(test_exchange_dlx)
            2. 声明死信队列(queue_dlx)和死信交换机(exchange_dlx)
            3. 正常队列绑定死信交换机
                设置两个参数：
                    * x-dead-letter-exchange：死信交换机名称
                    * x-dead-letter-routing-key：发送给死信交换机的routingkey
-->

<!--
     1. 声明正常的队列(test_queue_dlx)和交换机(test_exchange_dlx)
-->
<rabbit:queue name="test_queue_dlx" id="test_queue_dlx">
    <!--3. 正常队列绑定死信交换机-->
    <rabbit:queue-arguments>
        <!--3.1 x-dead-letter-exchange：死信交换机名称-->
        <entry key="x-dead-letter-exchange" value="exchange_dlx" />

        <!--3.2 x-dead-letter-routing-key：发送给死信交换机的routingkey-->
        <entry key="x-dead-letter-routing-key" value="dlx.hehe" />

        <!--4.1 设置队列的过期时间 ttl-->
        <entry key="x-message-ttl" value="10000" value-type="java.lang.Integer" />
        <!--4.2 设置队列的长度限制 max-length -->
        <entry key="x-max-length" value="10" value-type="java.lang.Integer" />
    </rabbit:queue-arguments>
</rabbit:queue>

<rabbit:topic-exchange name="test_exchange_dlx">
    <rabbit:bindings>
        <rabbit:binding pattern="test.dlx.#" queue="test_queue_dlx"></rabbit:binding>
    </rabbit:bindings>
</rabbit:topic-exchange>


<!--
       2. 声明死信队列(queue_dlx)和死信交换机(exchange_dlx)
   -->
<rabbit:queue name="queue_dlx" id="queue_dlx"></rabbit:queue>
<rabbit:topic-exchange name="exchange_dlx">
    <rabbit:bindings>
        <rabbit:binding pattern="dlx.#" queue="queue_dlx"></rabbit:binding>
    </rabbit:bindings>
</rabbit:topic-exchange>
```



```JAVA
/**
     * 发送测试死信消息：
     *  1. 过期时间
     *  2. 长度限制
     *  3. 消息拒收
     */
@Test
public void testDlx(){
    //1. 测试过期时间，死信消息
    //rabbitTemplate.convertAndSend("test_exchange_dlx","test.dlx.haha","我是一条消息，我会死吗？");

    //2. 测试长度限制后，消息死信
    /* for (int i = 0; i < 20; i++) {
            rabbitTemplate.convertAndSend("test_exchange_dlx","test.dlx.haha","我是一条消息，我会死吗？");
        }*/

    //3. 测试消息拒收
    rabbitTemplate.convertAndSend("test_exchange_dlx","test.dlx.haha","我是一条消息，我会死吗？");
}
```

#### **1.6** **延迟队列**

延迟队列，即消息进入队列后不会立即被消费，只有到达指定时间后，才会被消费。



需求：

- \1. 下单后，30分钟未支付，取消订单，回滚库存。
- \2. 新用户注册成功7天后，发送短信问候。

![](https://i0.hdslb.com/bfs/album/cfa9e6f4a92616d5e18eb62df87af96a8dba80b0.png)

实现方式：

- \1. 定时器
- \2. 延迟队列

很可惜，在RabbitMQ中并未提供延迟队列功能。

但是可以使用：`TTL+死信队列` 组合实现延迟队列的效果。

![](https://i0.hdslb.com/bfs/album/7d6292ac10322ce4eb1d42ad0c242784d51e468a.png)

小结：

- \1. 延迟队列 指消息进入队列后，可以被延迟一定时间，再进行消费。
- \2. RabbitMQ没有提供延迟队列功能，但是可以使用 ： TTL + DLX 来实现延迟队列效果。

```xml
<!--
        延迟队列：
            1. 定义正常交换机（order_exchange）和队列(order_queue)
            2. 定义死信交换机（order_exchange_dlx）和队列(order_queue_dlx)
            3. 绑定，设置正常队列过期时间为30分钟
    -->
<!-- 1. 定义正常交换机（order_exchange）和队列(order_queue)-->
<rabbit:queue id="order_queue" name="order_queue">
    <!-- 3. 绑定，设置正常队列过期时间为30分钟-->
    <rabbit:queue-arguments>
        <entry key="x-dead-letter-exchange" value="order_exchange_dlx" />
        <entry key="x-dead-letter-routing-key" value="dlx.order.cancel" />
        <entry key="x-message-ttl" value="10000" value-type="java.lang.Integer" />

    </rabbit:queue-arguments>
</rabbit:queue>
<rabbit:topic-exchange name="order_exchange">
    <rabbit:bindings>
        <rabbit:binding pattern="order.#" queue="order_queue"></rabbit:binding>
    </rabbit:bindings>
</rabbit:topic-exchange>

<!--  2. 定义死信交换机（order_exchange_dlx）和队列(order_queue_dlx)-->
<rabbit:queue id="order_queue_dlx" name="order_queue_dlx"></rabbit:queue>
<rabbit:topic-exchange name="order_exchange_dlx">
    <rabbit:bindings>
        <rabbit:binding pattern="dlx.order.#" queue="order_queue_dlx"></rabbit:binding>
    </rabbit:bindings>
</rabbit:topic-exchange>
```

```JAVA
@Test
public  void testDelay() throws InterruptedException {
    //1.发送订单消息。 将来是在订单系统中，下单成功后，发送消息
    rabbitTemplate.convertAndSend("order_exchange","order.msg","订单信息：id=1,time=2019年8月17日16:41:47");

    /*//2.打印倒计时10秒
        for (int i = 10; i > 0 ; i--) {
            System.out.println(i+"...");
            Thread.sleep(1000);
        }*/
}
```



#### 1.7 日志与监控

##### 1.7.1 RabbitMQ日志

RabbitMQ默认日志存放路径： `/var/log/rabbitmq/rabbit@xxx.log`

日志包含了RabbitMQ的版本号、Erlang的版本号、RabbitMQ服务节点名称、cookie的hash值、RabbitMQ配置文件地址、内存限制、磁盘限制、默认账户guest的创建以及权限配置等等。

##### 1.7.2  web管控台监控

##### 1.7.3  rabbitmqctl管理和监控

```bash
# 查看队列
rabbitmqctl list_queues

# 查看exchanges
rabbitmqctl list_exchanges

# 查看用户
rabbitmqctl list_users

# 查看连接
rabbitmqctl list_connections

# 查看消费者信息
rabbitmqctl list_consumers

# 查看环境变量
rabbitmqctl environment

# 查看未被确认的队列
rabbitmqctl list_queues  name messages_unacknowledged

# 查看单个队列的内存使用
rabbitmqctl list_queues name memory

# 查看准备就绪的队列
rabbitmqctl list_queues name messages_ready
```

#### 1.8 消息追踪

在使用任何消息中间件的过程中，难免会出现某条消息异常丢失的情况。对于RabbitMQ而言，可能是因为生产者或消费者与RabbitMQ断开了连接，而它们与RabbitMQ又采用了不同的确认机制；也有可能是因为交换器与队列之间不同的转发策略；甚至是交换器并没有与任何队列进行绑定，生产者又不感知或者没有采取相应的措施；另外RabbitMQ本身的集群策略也可能导致消息的丢失。这个时候就需要有一个较好的机制跟踪记录消息的投递过程，以此协助开发和运维人员进行问题的定位。

在RabbitMQ中可以使用`Firehose`和`rabbitmq_tracing`插件功能来实现消息追踪。

##### 1.8 消息追踪-Firehose

firehose的机制是将生产者投递给rabbitmq的消息，rabbitmq投递给消费者的消息按照指定的格式发送到`默认的exchange`上。这个默认的exchange的名称为`amq.rabbitmq.trace`，它是一个`topic`类型的exchange。发送到这个exchange上的消息的routing key为 publish.exchangename 和 deliver.queuename。其中exchangename和queuename为实际exchange和queue的名称，分别对应生产者投递到exchange的消息，和消费者从queue上获取的消息。



注意：打开 trace 会影响消息写入功能，适当打开后请关闭。

- rabbitmqctl trace_on：开启Firehose命令
- rabbitmqctl trace_off：关闭Firehose命令

##### 1.8 消息追踪-rabbitmq_tracing

rabbitmq_tracing和Firehose在实现上如出一辙，只不过rabbitmq_tracing的方式比Firehose多了一层GUI的包装，更容易使用和管理。

启用插件：rabbitmq-plugins enable rabbitmq_tracing

## 2. RabbitMQ 应用问题

RabbitMQ应用问题

\1. 消息可靠性保障

•消息补偿机制

\2. 消息幂等性保障

•乐观锁解决方案

### 2.1 消息可靠性

需求：100%确保消息发送成功

对于消息的可靠性传输，每种MQ都要从三个角度来分析：

- 生产者丢数据（生产到交换机）
  - 发送者确认模式--监听。`rabbitTemplate.setConfirmCallback(对象);`
- 消息队列丢数据（交换机到消息队列、消息队列持久化）
  - 交换机到消息队列：路由失败通知---监听，只针对失败，成功就没有了（默认不开启）`rabbitTemplate.setMandatory(true);`打开失败通知然后设置回调方法`rabbitTemplate.setReturnCallback(对象);`
  - 消息持久化：还没有存储系统宕机了怎么办。
    - 创建队列的时候默认已经是持久化的了https://blog.csdn.net/u013256816/article/details/60875666/
- 消费者丢数据（消息队列到消费者）
  - 根本没收到消息：消息到了消费者就被剔除了，可以选择手动ack模式
  - 在业务代码里出异常了：还是手动ack模式

#### 生产者丢消息

> （1）事务机制：
>
> 发送消息前，开启事务（channel.txSelect()），然后发送消息，如果发送过程中出现什么异常，事务就会回滚（channel.txRollback()），如果发送成功则提交事务（channel.txCommit()）
>
> 该方式的缺点是生产者发送消息会**同步阻塞等待发送结果是成功还是失败**，导致生产者发送消息的吞吐量降下降。
>
> （2）确认机制：
>
> 生产环境常用的是**confirm模式**。生产者将信道 channel 设置成 **confirm 模式**，一旦 channel 进入 confirm 模式，所有在该信道上发布的消息都将会被指派一个唯一的ID，一旦消息被投递到所有匹配的队列之后，rabbitMQ就会发送一个确认给生产者（包含消息的唯一ID），这样生产者就知道消息已经正确到达目的队列了。如果rabbitMQ没能处理该消息，也会发送一个Nack消息给你，这时就可以进行重试操作。
>
> Confirm模式最大的好处在于它是**异步**的（不阻塞），一旦发布消息，生产者就可以在等信道返回确认的同时继续发送下一条消息，当消息最终得到确认之后，生产者便可以通过回调方法来处理该确认消息。
>
> ```java
> channel.addConfirmListener(new ConfirmListener() {  
>     @Override  
>     public void handleNack(long deliveryTag, boolean multiple) throws IOException {  
>         System.out.println("nack: deliveryTag = "+deliveryTag+" multiple: "+multiple);  
>     }  
>     @Override  
>     public void handleAck(long deliveryTag, boolean multiple) throws IOException {  
>         System.out.println("ack: deliveryTag = "+deliveryTag+" multiple: "+multiple);  
>     }  
> }); 
> // 或者在rabbitTemplate.setConfirmCallback();//只能确保到了交换器
> ```

#### 消息队列丢数据：

处理消息队列丢数据的情况，一般是开启持久化磁盘。持久化配置可以和生产者的 confirm 机制配合使用，在消息持久化磁盘后，再给生产者发送一个Ack信号。这样的话，如果消息持久化磁盘之前，即使rabbitMQ挂掉了，生产者也会因为收不到Ack信号而再次重发消息。

> 持久化设置如下（必须同时设置以下 2 个配置）：
>
> （1）创建queue的时候，将queue的持久化标志durable在设置为true，代表是一个持久的队列，这样就可以保证 rabbitmq 持久化 queue 的元数据，但是不会持久化queue里的数据；
>
> （2）发送消息的时候将 deliveryMode 设置为 2，将消息设置为持久化的，此时 rabbitmq就会将消息持久化到磁盘上去。

#### 消费者丢数据：

消费者丢数据一般是因为采用了自动确认消息模式。该模式下，虽然消息还在处理中，但是消费中者会自动发送一个确认，通知rabbitmq已经收到消息了，这时rabbitMQ就会立即将消息删除。这种情况下，如果消费者出现异常而未能处理消息，那就会丢失该消息。

解决方案就是采用**手动确认消息**，等到消息被真正消费之后，再手动发送一个确认信号，即使中途消息没处理完，但是服务器宕机了，那rabbitmq就收不到发的ack，然后 rabbitmq 就会将这条消息重新分配给其他的消费者去处理。

需要注意的是：消息可靠性增强了，性能就下降了，因为写磁盘比写 RAM 慢的多，两者的吞吐量可能有 10 倍的差距。所以，是否要对消息进行持久化，需要综合考虑业务场景、性能需要，以及可能遇到的问题。若想达到单RabbitMQ服务器 10W 条/秒以上的消息吞吐量，则要么使用其他的方式来确保消息的可靠传输，要么使用非常快速的存储系统以支持全持久化，例如使用 SSD。或者仅对关键消息作持久化处理，且应该保证关键消息的量不会导致性能瓶颈。

思想：

- 同时发送两条消息：正常消息和延迟消息
- 正常消息被消费者 消费到后，消费者又给Q2发消息，

消息补偿：

![](https://i0.hdslb.com/bfs/album/cc9aab843999b0490ad8afd4bad5631fed5a3be3.png)

2发送正常消息，3过会再发一条相同的消息

2发送的消息在Q1中被正常消费到写入DB，发送ack给Q2。回调检查服务监听到Q2的消息，将消息写入MDB

如果1成功2失败，因为3也发送了消息放入Q3。此时回调检查服务也监听到了Q3，要去比对MDB是否一致，如果一致则代表消费过。如果MDB中不存在，就代表2失败了，就走8让生产者重新发。

如果2个都发送失败了，有MDB的定时检查服务，比对业务数据库DB与消息数据库MDB，就能发现差异

### 2.2 防止重复消费

如何保证不被重复消费

场景：正常情况下，消费者在消费消息后，会给消息队列发送一个确认，消息队列接收后就知道消息已经被成功消费了，然后就从队列中删除该消息，也就不会将该消息再发送给其他消费者了。不同消息队列发出的确认消息形式不同，RabbitMQ是通过发送一个ACK确认消息。**但是因为网络故障，消费者发出的确认并没有传到消息队列，导致消息队列不知道该消息已经被消费，然后就再次消息发送给了其他消费者，从而造成重复消费的情况**。

重复消费问题的解决思路是：保证消息的唯一性，即使多次传输，也不让消息的多次消费带来影响，也就是保证消息等幂性；幂等性指一个操作执行任意多次所产生的影响均与一次执行的影响相同。

具体解决方案如下：

（1）乐观锁：改造业务逻辑，使得在重复消费时也不影响最终的结果。例如对SQL语句： update t1 set money = 150 where id = 1 and money = 100; 做了个前置条件判断，即 money = 100 的情况下才会做更新，更通用的是做个 version 即版本号控制，对比消息中的版本号和数据库中的版本号。

（2）数据库的唯一主键约束：消费完消息之后，到数据库中做一个 insert 操作，如果出现重复消费的情况，就会导致主键冲突，避免数据库出现脏数据。

（3）**通过记录关键的key**，当重复消息过来时，先判断下这个key是否已经被处理过了，如果没处理再进行下一步。

- ① 通过数据库：比如处理订单时，记录订单ID，在消费前，去数据库中进行查询该记录是否存在，如果存在则直接返回。
- ② 使用全局唯一ID，再配合第三组主键做消费记录，比如使用 redis 的 set 结构，生产者发送消息时给消息分配一个全局ID，在每次消费者开始消费前，先去redis中查询有没有消费记录，如果消费过则不进行处理，如果没消费过，则进行处理，**消费完之后，就将这个ID以k-v的形式存入redis中(过期时间根据具体情况设置)**。

### 2.3 消息幂等性保障



幂等性指一次和多次请求某一个资源，对于资源本身应该具有同样的结果。也就是说，其任意多次执行对资源本身所产生的影响均与一次执行的影响相同。

> 

在MQ中指，消费多条相同的消息，得到与消费该消息一次相同的结果。

![](https://i0.hdslb.com/bfs/album/4e6da8efbbb953d55567fefe01ddb8b1e2b3e202.png)

##### 生产

假设一下发送10条一样的

```java
@RunWith(SpringJUnit4ClassRunner.class)
@ContextConfiguration(locations = "classpath:spring-rabbitmq-producer.xml")
public class ProducerTest {

    @Autowired
    private RabbitTemplate rabbitTemplate;

    @Test
    public void testSend() {
        for (int i = 0; i < 10; i++) {
            // 发送消息
            rabbitTemplate.convertAndSend("test_exchange_confirm", "confirm", "message confirm....");
        }
    }
}
```



##### 消费

```java
/**
 * 发送消息
 */
public class HelloWorld {
    public static void main(String[] args) throws IOException, TimeoutException {

        //1.创建连接工厂
        ConnectionFactory factory = new ConnectionFactory();
        //2. 设置参数
        factory.setHost("172.16.98.133");//ip  HaProxy的ip
        factory.setPort(5672); //端口 HaProxy的监听的端口
        //3. 创建连接 Connection
        Connection connection = factory.newConnection();
        //4. 创建Channel
        Channel channel = connection.createChannel();
        //5. 创建队列Queue
        channel.queueDeclare("hello_world",true,false,false,null);
        String body = "hello rabbitmq~~~";
        //6. 发送消息
        channel.basicPublish("","hello_world",null,body.getBytes());
        //7.释放资源
        channel.close();
        connection.close();

        System.out.println("send success....");

    }
}

```

### 消息有序性

针对保证消息有序性的问题，解决方法就是保证生产者入队的顺序是有序的，出队后的顺序消费则交给消费者去保证。

- （1）方法一：拆分queue，使得一个queue只对应一个消费者。由于MQ一般都能保证内部队列是先进先出的，所以把需要保持先后顺序的一组消息使用某种算法都分配到同一个消息队列中。然后只用一个消费者单线程去消费该队列，这样就能保证消费者是按照顺序进行消费的了。但是消费者的吞吐量会出现瓶颈。如果多个消费者同时消费一个队列，还是可能会出现顺序错乱的情况，这就相当于是多线程消费了
- （2）方法二：对于多线程的消费同一个队列的情况，可以使用重试机制：比如有一个微博业务场景的操作，发微博、写评论、删除微博，这三个异步操作。如果一个消费者先执行了写评论的操作，但是这时微博都还没发，写评论一定是失败的，等一段时间。等另一个消费者，先执行发微博的操作后，再执行，就可以成功。
- 

### 消息堆积

场景题：几千万条数据在MQ里积压了七八个小时。

1、出现该问题的原因：

消息堆积往往是生产者的生产速度与消费者的消费速度不匹配导致的。有可能就是消费者消费能力弱，渐渐地消息就积压了，也有可能是因为消息消费失败反复复重试造成的，也有可能是消费端出了问题，导致不消费了或者消费极其慢。比如，消费端每次消费之后要写mysql，结果mysql挂了，消费端hang住了不动了，或者消费者本地依赖的一个东西挂了，导致消费者挂了。

所以如果是 bug 则处理 bug；如果是因为本身消费能力较弱，则优化消费逻辑，比如优化前是一条一条消息消费处理的，那么就可以批量处理进行优化。

2、临时扩容，快速处理积压的消息：

（1）先修复 consumer 的问题，确保其恢复消费速度，然后将现有的 consumer 都停掉；

（2）临时创建原先 N 倍数量的 queue ，然后写一个**临时分发数据的消费者程序**，将该程序部署上去消费队列中积压的数据，消费之后不做任何耗时处理，直接均匀轮询写入临时建立好的 N 倍数量的 queue 中；

（3）接着，临时征用 N 倍的机器来部署 consumer，每个 consumer 消费一个临时 queue 的数据

（4）等快速消费完积压数据之后，恢复原先部署架构 ，重新用原先的 consumer 机器消费消息。

这种做法相当于临时将 queue 资源和 consumer 资源扩大 N 倍，以正常 N 倍速度消费。

3、恢复队列中丢失的数据：

如果使用的是 rabbitMQ，并且设置了过期时间，消息在 queue 里积压超过一定的时间会被 rabbitmq 清理掉，导致数据丢失。这种情况下，实际上队列中没有什么消息挤压，而是丢了大量的消息。所以就不能说增加 consumer 消费积压的数据了，这种情况可以采取 “批量重导” 的方案来进行解决。在流量低峰期，写一个程序，手动去查询丢失的那部分数据，然后将消息重新发送到mq里面，把丢失的数据重新补回来。

4、MQ长时间未处理导致MQ写满的情况如何处理：

如果消息积压在MQ里，并且长时间都没处理掉，导致MQ都快写满了，这种情况肯定是临时扩容方案执行太慢，这种时候只好采用 “丢弃+批量重导” 的方式来解决了。首先，临时写个程序，连接到mq里面消费数据，消费一个丢弃一个，快速消费掉积压的消息，降低MQ的压力，然后在流量低峰期时去手动查询重导丢失的这部分数据。



## 3.RabbitMQ集群搭建

摘要：实际生产应用中都会采用消息队列的集群方案，如果选择RabbitMQ那么有必要了解下它的集群方案原理

一般来说，如果只是为了学习RabbitMQ或者验证业务工程的正确性那么在本地环境或者测试环境上使用其单实例部署就可以了，但是出于MQ中间件本身的可靠性、并发性、吞吐量和消息堆积能力等问题的考虑，在生产环境上一般都会考虑使用RabbitMQ的集群方案。

##### 普通集群模式

就是在多台机器上启动多个 RabbitMQ 实例，每个机器启动一个。我们创建的 queue，只会放在其中一个 RabbitMQ 实例上，但是每个实例都同步 queue 的元数据（元数据是 queue 的一些配置信息，通过元数据，可以找到 queue 所在实例）。**消费的时候，如果连接到了另外一个实例，那么那个实例会从 queue 所在实例上拉取数据过来**。

（1）优点：普通集群模式主要用于提高系统的吞吐量，可以通过添加更加的节点来线性的扩展消息队列的吞吐量，就是说让集群中多个节点来服务某个 queue 的读写操作

（2）缺点：无高可用性，queue所在的节点宕机了，其他实例就无法从那个实例拉取数据；RabbitMQ 内部也会产生大量的数据传输。

### 3.1 集群方案的原理

> 镜像集群模式
>
> RabbitMQ 真正的高可用模式。镜像集群模式下，队列的元数据和消息会存在于多个实例上，每次写消息到 queue 时，会**自动将消息同步到各个实例的 queue** ，也就是说每个 RabbitMQ 节点都有这个 queue 的完整镜像，包含 queue 的全部数据。任何一个机器宕机了，其它机器节点还包含了这个 queue 的完整数据，其他 consumer 都可以到其它节点上去消费数据。
>
> 配置镜像队列的集群都包含一个主节点master和若干个从节点slave，slave会准确地按照master执行命令的顺序进行动作，故slave与master上维护的状态应该是相同的。如果master由于某种原因失效，那么按照slave加入的时间排序，"资历最老"的slave会被提升为新的master。
>
> **除发送消息外的所有动作都只会向master发送，然后再由master将命令执行的结果广播给各个slave**。如果消费者与slave建立连接并进行订阅消费，其实质上都是从master上获取消息，只不过看似是从slave上消费而已。比如消费者与slave建立了TCP连接之后执行一个Basic.Get的操作，那么首先是由slave将Basic.Get请求发往master，再由master准备好数据返回给slave，最后由slave投递给消费者。
>
> 

RabbitMQ这款消息队列中间件产品本身是基于Erlang编写，Erlang语言天生具备分布式特性（通过同步Erlang集群各节点的magic cookie来实现）。因此，RabbitMQ天然支持Clustering。这使得RabbitMQ本身不需要像ActiveMQ、Kafka那样通过ZooKeeper分别来实现HA方案和保存集群的元数据。集群是保证可靠性的一种方式，同时可以通过水平扩展以达到增加消息吞吐量能力的目的。

![1565245219265](https://i0.hdslb.com/bfs/album/d756499281fb54b24844eedbdb9264c69a7a8f6d.png)


### 3.2 单机多实例部署

由于某些因素的限制，有时候你不得不在一台机器上去搭建一个rabbitmq集群，这个有点类似zookeeper的单机版。真实生成环境还是要配成多机集群的。有关怎么配置多机集群的可以参考其他的资料，这里主要论述如何在单机中配置多个rabbitmq实例。

主要参考官方文档：https://www.rabbitmq.com/clustering.html

首先确保RabbitMQ运行没有问题

```shell
[root@super ~]# rabbitmqctl status
Status of node rabbit@super ...
[{pid,10232},
 {running_applications,
     [{rabbitmq_management,"RabbitMQ Management Console","3.6.5"},
      {rabbitmq_web_dispatch,"RabbitMQ Web Dispatcher","3.6.5"},
      {webmachine,"webmachine","1.10.3"},
      {mochiweb,"MochiMedia Web Server","2.13.1"},
      {rabbitmq_management_agent,"RabbitMQ Management Agent","3.6.5"},
      {rabbit,"RabbitMQ","3.6.5"},
      {os_mon,"CPO  CXC 138 46","2.4"},
      {syntax_tools,"Syntax tools","1.7"},
      {inets,"INETS  CXC 138 49","6.2"},
      {amqp_client,"RabbitMQ AMQP Client","3.6.5"},
      {rabbit_common,[],"3.6.5"},
      {ssl,"Erlang/OTP SSL application","7.3"},
      {public_key,"Public key infrastructure","1.1.1"},
      {asn1,"The Erlang ASN1 compiler version 4.0.2","4.0.2"},
      {ranch,"Socket acceptor pool for TCP protocols.","1.2.1"},
      {mnesia,"MNESIA  CXC 138 12","4.13.3"},
      {compiler,"ERTS  CXC 138 10","6.0.3"},
      {crypto,"CRYPTO","3.6.3"},
      {xmerl,"XML parser","1.3.10"},
      {sasl,"SASL  CXC 138 11","2.7"},
      {stdlib,"ERTS  CXC 138 10","2.8"},
      {kernel,"ERTS  CXC 138 10","4.2"}]},
 {os,{unix,linux}},
 {erlang_version,
     "Erlang/OTP 18 [erts-7.3] [source] [64-bit] [async-threads:64] [hipe] [kernel-poll:true]\n"},
 {memory,
     [{total,56066752},
      {connection_readers,0},
      {connection_writers,0},
      {connection_channels,0},
      {connection_other,2680},
      {queue_procs,268248},
      {queue_slave_procs,0},
      {plugins,1131936},
      {other_proc,18144280},
      {mnesia,125304},
      {mgmt_db,921312},
      {msg_index,69440},
      {other_ets,1413664},
      {binary,755736},
      {code,27824046},
      {atom,1000601},
      {other_system,4409505}]},
 {alarms,[]},
 {listeners,[{clustering,25672,"::"},{amqp,5672,"::"}]},
 {vm_memory_high_watermark,0.4},
 {vm_memory_limit,411294105},
 {disk_free_limit,50000000},
 {disk_free,13270233088},
 {file_descriptors,
     [{total_limit,924},{total_used,6},{sockets_limit,829},{sockets_used,0}]},
 {processes,[{limit,1048576},{used,262}]},
 {run_queue,0},
 {uptime,43651},
 {kernel,{net_ticktime,60}}]
```

停止rabbitmq服务

```shell
[root@super sbin]# service rabbitmq-server stop
Stopping rabbitmq-server: rabbitmq-server.

```



启动第一个节点：

```shell
[root@super sbin]# RABBITMQ_NODE_PORT=5673 RABBITMQ_NODENAME=rabbit1 rabbitmq-server start

              RabbitMQ 3.6.5. Copyright (C) 2007-2016 Pivotal Software, Inc.
  ##  ##      Licensed under the MPL.  See http://www.rabbitmq.com/
  ##  ##
  ##########  Logs: /var/log/rabbitmq/rabbit1.log
  ######  ##        /var/log/rabbitmq/rabbit1-sasl.log
  ##########
              Starting broker...
 completed with 6 plugins.
```

启动第二个节点：

> web管理插件端口占用,所以还要指定其web插件占用的端口号。

```shell
[root@super ~]# RABBITMQ_NODE_PORT=5674 RABBITMQ_SERVER_START_ARGS="-rabbitmq_management listener [{port,15674}]" RABBITMQ_NODENAME=rabbit2 rabbitmq-server start

              RabbitMQ 3.6.5. Copyright (C) 2007-2016 Pivotal Software, Inc.
  ##  ##      Licensed under the MPL.  See http://www.rabbitmq.com/
  ##  ##
  ##########  Logs: /var/log/rabbitmq/rabbit2.log
  ######  ##        /var/log/rabbitmq/rabbit2-sasl.log
  ##########
              Starting broker...
 completed with 6 plugins.

```

结束命令：

```shell
rabbitmqctl -n rabbit1 stop
rabbitmqctl -n rabbit2 stop
```



rabbit1操作作为主节点：

```shell
[root@super ~]# rabbitmqctl -n rabbit1 stop_app  
Stopping node rabbit1@super ...
[root@super ~]# rabbitmqctl -n rabbit1 reset	 
Resetting node rabbit1@super ...
[root@super ~]# rabbitmqctl -n rabbit1 start_app
Starting node rabbit1@super ...
[root@super ~]# 
```

rabbit2操作为从节点：

```shell
[root@super ~]# rabbitmqctl -n rabbit2 stop_app
Stopping node rabbit2@super ...
[root@super ~]# rabbitmqctl -n rabbit2 reset
Resetting node rabbit2@super ...
[root@super ~]# rabbitmqctl -n rabbit2 join_cluster rabbit1@'super' ###''内是主机名换成自己的
Clustering node rabbit2@super with rabbit1@super ...
[root@super ~]# rabbitmqctl -n rabbit2 start_app
Starting node rabbit2@super ...

```

查看集群状态：

```
[root@super ~]# rabbitmqctl cluster_status -n rabbit1
Cluster status of node rabbit1@super ...
[{nodes,[{disc,[rabbit1@super,rabbit2@super]}]},
 {running_nodes,[rabbit2@super,rabbit1@super]},
 {cluster_name,<<"rabbit1@super">>},
 {partitions,[]},
 {alarms,[{rabbit2@super,[]},{rabbit1@super,[]}]}]
```

web监控：

- rabbit1@super
- rabbit2@super





### 3.3 集群管理

**rabbitmqctl join_cluster {cluster_node} [–ram]**
将节点加入指定集群中。在这个命令执行前需要停止RabbitMQ应用并重置节点。

**rabbitmqctl cluster_status**
显示集群的状态。

**rabbitmqctl change_cluster_node_type {disc|ram}**
修改集群节点的类型。在这个命令执行前需要停止RabbitMQ应用。

**rabbitmqctl forget_cluster_node [–offline]**
将节点从集群中删除，允许离线执行。

**rabbitmqctl update_cluster_nodes {clusternode}**

在集群中的节点应用启动前咨询clusternode节点的最新信息，并更新相应的集群信息。这个和join_cluster不同，它不加入集群。考虑这样一种情况，节点A和节点B都在集群中，当节点A离线了，节点C又和节点B组成了一个集群，然后节点B又离开了集群，当A醒来的时候，它会尝试联系节点B，但是这样会失败，因为节点B已经不在集群中了。

**rabbitmqctl cancel_sync_queue [-p vhost] {queue}**
取消队列queue同步镜像的操作。

**rabbitmqctl set_cluster_name {name}**
设置集群名称。集群名称在客户端连接时会通报给客户端。Federation和Shovel插件也会有用到集群名称的地方。集群名称默认是集群中第一个节点的名称，通过这个命令可以重新设置。

### 3.4 RabbitMQ镜像集群配置

> 上面已经完成RabbitMQ默认集群模式，但并不保证队列的高可用性，尽管交换机、绑定这些可以复制到集群里的任何一个节点，但是队列内容不会复制。虽然该模式解决一项目组节点压力，但队列节点宕机直接导致该队列无法应用，只能等待重启，所以要想在队列节点宕机或故障也能正常应用，就要复制队列内容到集群里的每个节点，必须要创建镜像队列。
>
> 镜像队列是基于普通的集群模式的，然后再添加一些策略，所以你还是得先配置普通集群，然后才能设置镜像队列，我们就以上面的集群接着做。

**设置的镜像队列可以通过开启的网页的管理端Admin->Policies，也可以通过命令。**

> rabbitmqctl set_policy my_ha "^" '{"ha-mode":"all"}'

在管理台点击Add/update a policy

> - Name:策略名称  my_ha
> - Pattern：匹配的规则，如果是匹配所有的队列，是^
> - Definition:使用ha-mode模式中的all，也就是同步所有匹配的队列。问号链接帮助文档。
> - apply-to:Exchanges and queues

### 3.5 负载均衡-HAProxy

HAProxy提供高可用性、负载均衡以及基于TCP和HTTP应用的代理，支持虚拟主机，它是免费、快速并且可靠的一种解决方案,包括Twitter，Reddit，StackOverflow，GitHub在内的多家知名互联网公司在使用。HAProxy实现了一种事件驱动、单一进程模型，此模型支持非常大的并发连接数。

##### 3.5.1  安装HAProxy

```shell
//下载依赖包
yum install gcc vim wget
//上传haproxy源码包
//解压
tar -zxvf haproxy-1.6.5.tar.gz -C /usr/local
//进入目录、进行编译、安装
cd /usr/local/haproxy-1.6.5
make TARGET=linux31 PREFIX=/usr/local/haproxy
make install PREFIX=/usr/local/haproxy
mkdir /etc/haproxy
//赋权
groupadd -r -g 149 haproxy
useradd -g haproxy -r -s /sbin/nologin -u 149 haproxy
//创建haproxy配置文件
mkdir /etc/haproxy
vim /etc/haproxy/haproxy.cfg
```




##### 3.5.2 配置HAProxy

配置文件路径：/etc/haproxy/haproxy.cfg

```shell
#logging options
global
	log 127.0.0.1 local0 info
	maxconn 5120
	chroot /usr/local/haproxy
	uid 99
	gid 99
	daemon
	quiet
	nbproc 20
	pidfile /var/run/haproxy.pid

defaults
	log global
	
	mode tcp

	option tcplog
	option dontlognull
	retries 3
	option redispatch
	maxconn 2000
	contimeout 5s
   
     clitimeout 60s

     srvtimeout 15s	
#front-end IP for consumers and producters

listen rabbitmq_cluster
	bind 0.0.0.0:5672
	
	mode tcp
	#balance url_param userid
	#balance url_param session_id check_post 64
	#balance hdr(User-Agent)
	#balance hdr(host)
	#balance hdr(Host) use_domain_only
	#balance rdp-cookie
	#balance leastconn
	#balance source //ip
	
	balance roundrobin
	
        server node1 127.0.0.1:5673 check inter 5000 rise 2 fall 2
        server node2 127.0.0.1:5674 check inter 5000 rise 2 fall 2

listen stats
	bind 172.16.98.133:8100
	mode http
	option httplog
	stats enable
	stats uri /rabbitmq-stats
	stats refresh 5s
```

启动HAproxy负载

```shell
/usr/local/haproxy/sbin/haproxy -f /etc/haproxy/haproxy.cfg
//查看haproxy进程状态
ps -ef | grep haproxy

访问如下地址对mq节点进行监控
http://172.16.98.133:8100/rabbitmq-stats
```

代码中访问mq集群地址，则变为访问haproxy地址:5672