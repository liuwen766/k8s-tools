payment8001 写在前面

- 官方代码地址：https://github.com/zzyybs/atguigu_spirngcloud2020
- 本文地址：https://blog.csdn.net/hancoder/article/details/109063671
- 本文下载方式：见博客中的谷粒商城



集中什么是微服务架构：

![1597213385700](https://i0.hdslb.com/bfs/album/6594a134e8fb60fb43ac99f4c7d5c4704f854302.png)

SpringCloud 是微服务一站式服务解决方案，微服务全家桶。它是微服务开发的主流技术栈。它采用了名称，而非数字版本号。

SpringCloud  和 springCloud Alibaba 目前是最主流的微服务框架组合。

版本选择：

> 选用 springboot 和 springCloud 版本有约束，不按照它的约束会有冲突。

### 版本问题

本次学习的各种软件的版本：

- boot使用的是数字作为版本。官网强烈建议升级到2.0以上
- cloud使用的是字母作为版本，伦敦地铁站站名

| Cloud Release Train                                          | Boot Version                     |
| :----------------------------------------------------------- | :------------------------------- |
| [Hoxton](https://github.com/spring-projects/spring-cloud/wiki/Spring-Cloud-Hoxton-Release-Notes) | 2.2.x, 2.3.x (Starting with SR5) |
| [Greenwich](https://github.com/spring-projects/spring-cloud/wiki/Spring-Cloud-Greenwich-Release-Notes) | 2.1.x                            |
| [Finchley](https://github.com/spring-projects/spring-cloud/wiki/Spring-Cloud-Finchley-Release-Notes) | 2.0.x                            |
| [Edgware](https://github.com/spring-projects/spring-cloud/wiki/Spring-Cloud-Edgware-Release-Notes) | 1.5.x                            |
| [Dalston](https://github.com/spring-projects/spring-cloud/wiki/Spring-Cloud-Dalston-Release-Notes) | 1.5.x                            |

查看版本对应关系：https://start.spring.io/actuator/info

```json
"spring-cloud": {
    "Finchley.M2": "Spring Boot >=2.0.0.M3 and <2.0.0.M5",
    "Finchley.M3": "Spring Boot >=2.0.0.M5 and <=2.0.0.M5",
    "Finchley.M4": "Spring Boot >=2.0.0.M6 and <=2.0.0.M6",
    "Finchley.M5": "Spring Boot >=2.0.0.M7 and <=2.0.0.M7",
    "Finchley.M6": "Spring Boot >=2.0.0.RC1 and <=2.0.0.RC1",
    "Finchley.M7": "Spring Boot >=2.0.0.RC2 and <=2.0.0.RC2",
    "Finchley.M9": "Spring Boot >=2.0.0.RELEASE and <=2.0.0.RELEASE",
    "Finchley.RC1": "Spring Boot >=2.0.1.RELEASE and <2.0.2.RELEASE",
    "Finchley.RC2": "Spring Boot >=2.0.2.RELEASE and <2.0.3.RELEASE",
    "Finchley.SR4": "Spring Boot >=2.0.3.RELEASE and <2.0.999.BUILD-SNAPSHOT",
    "Finchley.BUILD-SNAPSHOT": "Spring Boot >=2.0.999.BUILD-SNAPSHOT and <2.1.0.M3",
    "Greenwich.M1": "Spring Boot >=2.1.0.M3 and <2.1.0.RELEASE",
    "Greenwich.SR6": "Spring Boot >=2.1.0.RELEASE and <2.1.18.BUILD-SNAPSHOT",
    "Greenwich.BUILD-SNAPSHOT": "Spring Boot >=2.1.18.BUILD-SNAPSHOT and <2.2.0.M4",
    "Hoxton.SR8": "Spring Boot >=2.2.0.M4 and <2.3.5.BUILD-SNAPSHOT",
    "Hoxton.BUILD-SNAPSHOT": "Spring Boot >=2.3.5.BUILD-SNAPSHOT and <2.4.0.M1",
    "2020.0.0-M3": "Spring Boot >=2.4.0.M1 and <=2.4.0.M1",
    "2020.0.0-SNAPSHOT": "Spring Boot >=2.4.0.M2"
},
```

尚硅谷阳哥教程版本：

|               |               |
| ------------- | ------------- |
| cloud         | Hoxton.SR1    |
| boot          | 2.2.2.RELEASE |
| cloud alibaba | 2.1.0.RELEASE |

cloud版本决定了boot版本

#### 微服务停更说明

1,Eureka停用，可以使用zk作为服务注册中心

2,服务调用，Ribbon准备停更，代替为LoadBalance

3,Feign改为OpenFeign

4,Hystrix停更，改为resilence4j

​		或者阿里巴巴的sentienl

5.Zuul改为gateway

6,服务配置Config改为  Nacos

7,服务总线Bus改为Nacos

![1597213783265](https://i0.hdslb.com/bfs/album/8ba6cf17998c1de765f45bd28c8f0ffff21a1023.png)

## Cloud简介



参考资料，尽量去官网

https://cloud.spring.io/spring-cloud-static/Hoxton.SR1/reference/htmlsingle/

# 工程建造

写一个下图的Hello World

![1597225177689](https://i0.hdslb.com/bfs/album/5cc1e259da8a7f2a15afcbc507995e71110e8a7f.png)

构建父工程，后面的项目模块都在此工程中：

![1597214225785](https://i0.hdslb.com/bfs/album/94e45c7354d26b18481e8eb515cb11e21ab23db3.png)

设置编码：Settings -> File Encodings

注解激活：

![1597214602636](https://i0.hdslb.com/bfs/album/42da7b5e33d97d1387719879e3e86ffdfc1d3d09.png)

Java版本确定：

![1597214699619](https://i0.hdslb.com/bfs/album/59190bd501b84a21d4755db198b46efd42517bba.png)

## 父工程pom配置



```xml
<?xml version="1.0" encoding="UTF-8"?>

<project xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
    <modelVersion>4.0.0</modelVersion>

    <groupId>com.dkf.cloud</groupId>
    <artifactId>cloud2020</artifactId>
    <version>1.0-SNAPSHOT</version>
    <!-- 第一步 打包方式pom-->
    <packaging>pom</packaging>

    <!-- 统一管理 jar 包版本 -->
    <properties>
        <project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>
        <maven.compiler.source>1.8</maven.compiler.source>
        <maven.compiler.target>1.8</maven.compiler.target>
        <junit.version>4.12</junit.version>
        <log4j.version>1.2.17</log4j.version>
        <lombok.version>1.16.18</lombok.version>
        <mysql.version>5.1.47</mysql.version>
        <druid.version>1.1.16</druid.version>
        <mybatis.spring.boot.version>1.3.0</mybatis.spring.boot.version>
    </properties>

    <!-- 子块基础之后，提供作用：锁定版本 + 子module不用写 groupId 和 version -->
    <dependencyManagement>
        <dependencies>
            <dependency>
                <groupId>org.apache.maven.plugins</groupId>
                <artifactId>maven-project-info-reports-plugin</artifactId>
                <version>3.0.0</version>
            </dependency>

            <!-- 下面三个基本是微服务架构的标配 -->
            <!--spring boot 2.2.2-->
            <dependency>
                <groupId>org.springframework.boot</groupId>
                <artifactId>spring-boot-dependencies</artifactId>
                <version>2.2.2.RELEASE</version>
                <type>pom</type>
                <scope>import</scope>
            </dependency>
            <!--spring cloud Hoxton.SR1-->
            <dependency>
                <groupId>org.springframework.cloud</groupId>
                <artifactId>spring-cloud-dependencies</artifactId>
                <version>Hoxton.SR1</version>
                <type>pom</type>
                <scope>import</scope>
            </dependency>
            <!--spring cloud 阿里巴巴-->
            <dependency>
                <groupId>com.alibaba.cloud</groupId>
                <artifactId>spring-cloud-alibaba-dependencies</artifactId>
                <version>2.1.0.RELEASE</version>
                <type>pom</type>
                <scope>import</scope>
            </dependency>

            <!--mysql-->
            <dependency>
                <groupId>mysql</groupId>
                <artifactId>mysql-connector-java</artifactId>
                <version>${mysql.version}</version>
                <scope>runtime</scope>
            </dependency>
            <!-- druid-->
            <dependency>
                <groupId>com.alibaba</groupId>
                <artifactId>druid</artifactId>
                <version>${druid.version}</version>
            </dependency>
            <dependency>
                <groupId>org.mybatis.spring.boot</groupId>
                <artifactId>mybatis-spring-boot-starter</artifactId>
                <version>${mybatis.spring.boot.version}</version>
            </dependency>
            <!--junit-->
            <dependency>
                <groupId>junit</groupId>
                <artifactId>junit</artifactId>
                <version>${junit.version}</version>
            </dependency>
            <!--log4j-->
            <dependency>
                <groupId>log4j</groupId>
                <artifactId>log4j</artifactId>
                <version>${log4j.version}</version>
            </dependency>
        </dependencies>
    </dependencyManagement>

    <build>
        <plugins>
            <plugin>
                <groupId>org.springframework.boot</groupId>
                <artifactId>spring-boot-maven-plugin</artifactId>
                <configuration>
                    <fork>true</fork>
                    <addResources>true</addResources>
                </configuration>
            </plugin>
        </plugins>
    </build>
</project>
```

上面配置的解释：

> 首先要加  `<packaging>pom</packaging> ` 这个。
>
> 为了让项目顺利的运行，我们必须使用统一的版本号；
> **1、dependencyManagement**
> (1)在我们项目中，我们会发现在父模块的pom文件中常常会出现dependencyManagement元素，这是因为我们可以通过其来管理子模块的版本号，也就是说我们在父模块中**声明号依赖的版本，但是并不实现引入**；
> **2、dependencies**
> (1)上面说到dependencyManagement只是声明一个依赖，而不实现引入，故我们在子模块中也需要对依赖进行声明，倘若不声明子模块自己的依赖，是不会从父模块中继承的；只有子模块中也声明了依赖。并且没有写对应的版本号它才会从父类中继承；并且version和scope都是取自父类；此外要是子模块中自己定义了自己的版本号，是不会继承自父类的。
> **3、总结**
> dependencyManagement只是用来管理依赖，规定未添加版本号的子模块依赖继承自它，dependencies是用来声明子模块自己的依赖，可以在其中来写自己需要的版本号
>
> 聚合版本依赖，dependencyManagement只声明依赖，并不实现引入，所以子项目还需要写要引入的依赖。
>

可以统一版本

父工程创建完成执行mvn:install将父工程发布到仓库方便子工程继承

## 第一个微服务架构

创建一个module后(只能改a)，在父工程的pom里多了个`<modules>`，

在模块的pom里没有gv，只有a。

模块里的`<dependencies>`里的依赖只有ga，没有v

> payment8001 是提供者，order是消费者

### ==提供者==

cloud-provider-payment8001 子工程的pom文件：

> 这里面的 lombok 这个包，引入以后，实体类不用再写set 和 get
>
> 可以如下写实体类：
>
> ```java
> import lombok.AllArgsConstructor;
> import lombok.Data;
> import lombok.NoArgsConstructor;
> import java.io.Serializable;
> 
> @Data
> @AllArgsConstructor
> @NoArgsConstructor
> public class Payment implements Serializable {
>      private Integer id;
>      private String serial;
> }
> ```

```xml
<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0"
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
    <parent>
        <artifactId>cloud2020</artifactId>
        <groupId>com.dkf.cloud</groupId>
        <version>1.0-SNAPSHOT</version>
    </parent>
    <modelVersion>4.0.0</modelVersion>

    <artifactId>cloud-provider-payment8001</artifactId>
    <dependencies>
        <!--eureka-client-->
        <dependency>
            <groupId>org.springframework.cloud</groupId>
            <artifactId>spring-cloud-starter-netflix-eureka-client</artifactId>
        </dependency>
        <dependency>
            <!-- 引入自己定义的api通用包，可以使用Payment支付Entity -->
            <groupId>com.atguigu.springcloud</groupId>
            <artifactId>cloud-api-commons</artifactId>
            <version>${project.version}</version>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-web</artifactId>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <!--这个和web要写到一块-->
            <artifactId>spring-boot-starter-actuator</artifactId>
        </dependency>
        
        <dependency>
            <groupId>org.mybatis.spring.boot</groupId>
            <artifactId>mybatis-spring-boot-starter</artifactId>
        </dependency>
        <dependency>
            <groupId>com.alibaba</groupId>
            <artifactId>druid-spring-boot-starter</artifactId>
            <version>1.1.10</version>
        </dependency>
        <!--mysql-connector-java-->
        <dependency>
            <groupId>mysql</groupId>
            <artifactId>mysql-connector-java</artifactId>
        </dependency>
        <!--jdbc-->
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-jdbc</artifactId>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-devtools</artifactId>
            <scope>runtime</scope>
            <optional>true</optional>
        </dependency>
        <dependency>
            <groupId>org.projectlombok</groupId>
            <artifactId>lombok</artifactId>
            <optional>true</optional>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-test</artifactId>
            <scope>test</scope>
        </dependency>
    </dependencies>

</project>
```

cloud-provider-payment8001 子工程的yml文件：

```yml
server:
  port: 8001

spring:
  application:
    name: cloud-provider-payment8001
  datasource:
    type: com.alibaba.druid.pool.DruidDataSource
    driver-class-name: org.gjt.mm.mysql.Driver
    url: jdbc:mysql://localhost:3306/cloud2020?useUnicode=true&characterEncoding=utf-8&useSSL=false
    username: root
    password: 123456
    
mybatis:
  mapper-locations: classpath:mapper/*.xml
  type-aliases-package: com.dkf.springcloud.entities  # 所有Entity 别名类所在包
```

cloud-provider-payment8001 子工程的主启动类：

```java
package com.dkf.springcloud;

@SpringBootApplication
public class PaymentMain8001 {
    public static void main(String[] args){
        SpringApplication.run(PaymentMain8001.class, args);
    }
}
```

下面的常规操作：

- ①建表SQL：建立支付表

  ```mysql
  create table `payment`(
      `id` bigint(20) not null auto_increment comment 'ID',
      `serial` varchar(200) default '',
      PRIMARY KEY (`id`)
  
  )ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8
  
  select * from payment;
  ```

##### entities

建立支付实体和共用结果类

  ```java
@Data   //set/get方法
@AllArgsConstructor //有参构造器
@NoArgsConstructor  //无参构造器
public class Payment implements Serializable {
    private long id;//数据库是bigint
    private String serial;
}
  ```

  ```java
//返回给前端的通用json数据串
@Data   //set/get方法
@AllArgsConstructor //有参构造器
@NoArgsConstructor  //无参构造器
public class CommonResult<T> {
    private Integer code;
    private String message;
    private T data; //泛型，对应类型的json数据

    //自定义两个参数的构造方法
    public CommonResult(Integer code, String message){
        this(code, message, null);
    }
}
  ```

##### dao

创建支付的dao层：创建支付和查询支付

  ```java
@Mapper // 是ibatis下面的注解 //@Repositoty有时候会有问题
public interface PaymentDao {
    
    int create(Payment payment);

    Payment getPaymentById(@Param("id") Long id);
}
  ```

resource下创建mapper文件夹，新建PaymentMapper.xml。在yml里有所有entity别名类所在包，所有payment不用写全类名

  ```xml
<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE mapper PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN" "http://mybatis.org/dtd/mybatis-3-mapper.dtd" >
<mapper namespace="com.xzq.springcloud.dao.PaymentDao">
    
    <resultMap id="BaseResultMap" type="com.xzq.springcloud.entities.Payment">
        <id column="id" property="id" jdbcType="BIGINT"/>
        <id column="serial" property="serial" jdbcType="VARCHAR"/>
    </resultMap>

    <insert id="create" parameterType="Payment" useGeneratedKeys="true" keyProperty="id">
        insert into payment(serial) values (#{serial})
    </insert>

    <select id="getPaymentById" parameterType="Long" resultMap="BaseResultMap">
        select * from payment where id = #{id}
    </select>
</mapper>
  ```

@Param注解：https://blog.csdn.net/qq_39505065/article/details/90550705

##### service

  ```java
public interface PaymentService {
    int create(Payment payment);

    Payment getPaymentById(@Param("id") Long id);
}
  ```

  ```java
@Service
public class PaymentServiceImpl implements PaymentService {

    @Autowired
    private PaymentDao paymentDao;

    @Override
    public int create(Payment payment) {
        return paymentDao.create(payment);
    }

    @Override
    public Payment getPaymentById(Long id) {
        return paymentDao.getPaymentById(id);
    }
}
  ```

##### controller

  ```java
@RestController   //必须是这个注解，因为是模拟前后端分离的restful风格的请求，要求每个方法返回 json
@Slf4j
public class PaymentController {

    @Resource
    private PaymentService paymentService;

    @PostMapping(value = "/payment/create")
    //	    注意这里的 @RequestBody  是必须要写的，虽然 MVC可以自动封装参数成为对象，
    //      但是当消费者项目调用，它传参是 payment 整个实例对象传过来的， 即Json数据，因此需要写这个注解
    // https://blog.csdn.net/weixin_38004638/article/details/99655322
    public CommonResult create(@RequestBody Payment payment){
        int result = paymentService.create(payment);
        log.info("****插入结果：" + result);
        if(result > 0){
            return new CommonResult(200, "插入数据库成功", result);
        }
        return new CommonResult(444, "插入数据库失败", null);
    }

    @GetMapping(value = "/payment/{id}")
    public CommonResult getPaymentById(@PathVariable("id")Long id){
        Payment result = paymentService.getPaymentById(id);
        log.info("****查询结果：" + result);
        if(result != null){
            return new CommonResult(200, "查询成功", result);
        }
        return new CommonResult(444, "没有对应id的记录", null);
    }
}
  ```

对应POST方式的请求，要学会用`POSTMAN`工具

微服务多了之后就使用run dashboard

不但编译有个别地方会报错，启动也会报错，但是测试两个接口都是没问题的，推测启动报错是因为引入了下面才会引入的jar包，目前不影响。

### 热部署配置devtools

代码改动后希望自动生效

1. 具体模块里添加Jar包到工程中，上面的pom文件已经添加上了

```xml
<dependency>    
    <groupId>org.springframework.boot</groupId>    
    <artifactId>spring-boot-devtools</artifactId>    
    <scope>runtime</scope>    
    <optional>true</optional>
</dependency>
```

2. 添加plugin到父工程的pom文件中：上面也已经添加好了

```xml
<build>
    <plugins>
        <plugin>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-maven-plugin</artifactId>
            <configuration>
                <fork>true</fork>
                <addResources>true</addResources>
            </configuration>
        </plugin>
    </plugins>
</build>
```

3. ![1597222225568](https://i0.hdslb.com/bfs/album/a31101cf1f85773b32aef4632f736b95dbac4601.png)

4. shift + ctrl + alt + / 四个按键一块按，选择Reg项：

![1597222341651](https://i0.hdslb.com/bfs/album/e9783d3ed620b0bb7a3e4f2da36d6c91a7c29f9b.png)

重启IDEA

热部署只允许在开发阶段使用

### ==消费者==

新建模块cloud-consumer-order80

> 消费者现在只模拟调用提供者的Controller方法，没有持久层配置，只有Controller和实体类
>
> 当然也要配置主启动类和启动端口

pom文件：

```xml
<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0"
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
    <parent>
        <artifactId>cloud2020</artifactId>
        <groupId>com.dkf.cloud</groupId>
        <version>1.0-SNAPSHOT</version>
    </parent>
    <modelVersion>4.0.0</modelVersion>

    <artifactId>cloud-customer-order80</artifactId>

    <dependencies>
        <dependency>
            <groupId>org.springframework.cloud</groupId>
            <artifactId>spring-cloud-starter-netflix-eureka-client</artifactId>
        </dependency>
        <!--<dependency>&lt;!&ndash; 引入自己定义的api通用包，可以使用Payment支付Entity &ndash;&gt;-->
        <!--<groupId>com.atguigu.springcloud</groupId>-->
        <!--<artifactId>cloud-api-commons</artifactId>-->
        <!--<version>${project.version}</version>-->
        <!--</dependency>-->
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-web</artifactId>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-actuator</artifactId>
        </dependency>

        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-devtools</artifactId>
            <scope>runtime</scope>
            <optional>true</optional>
        </dependency>
        <dependency>
            <groupId>org.projectlombok</groupId>
            <artifactId>lombok</artifactId>
            <optional>true</optional>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-test</artifactId>
            <scope>test</scope>
        </dependency>
    </dependencies>
</project>
```

把CommonResult 和 Payment 两个 实体类也创建出来

- config/ApplicationContextCOnfig.java
- controller/OrderController.java
- entities/CommonResult.java
- entities/Payment.java
- OrderMain80.java

application.yml

```yaml
server:
	port:80 # 80端口就可以省略了
```

OrderMain80.java

```java
@SpringbootApplication
public class OrderMain80{
    public static void main(String[] args){
        SpringApplication.run(OrderMain80.class,args);
    }
}
```

entites包中的类也拷贝到本项目中

- entities/CommonResult.java
- entities/Payment.java

##### 配置RestTemplate

ApplicationContextConfig 内容：

```java
package com.dkf.springcloud.config;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.web.client.RestTemplate;//网络客户端

@Configuration
public class ApplicationContextConfig {

    @Bean // 远程调用工具
    public RestTemplate getRestTemplate(){
        return new RestTemplate();
        /*
        RestTemplate提供了多种便捷访问远程http服务的方法，
        是一种简单便捷的访问restful服务模板类，是spring提供的用于rest服务的客户端模板工具集
        */
    }
}
```

##### Controller

order消费payment

```java
@RestController
@Slf4j
public class OrderController {

    //远程调用的 地址
    public static final String PAYMENY_URL = "http://localhost:8001";

    @Resource
    private RestTemplate restTemplate;

    @PostMapping("customer/payment/create")
    public CommonResult<Payment> create (Payment payment){

        return restTemplate.postForObject(PAYMENY_URL + "/payment/create",//请求地址
                                          payment,//请求参数
                                          CommonResult.class);//返回类型
    }

    @GetMapping("customer/payment/{id}")
    public CommonResult<Payment> getPaymentById(@PathVariable("id")Long id){
        return restTemplate.getForObject(PAYMENY_URL + "/payment/" + id,//请求地址
                                         CommonResult.class);//返回类型
    }
}
```

如果 runDashboard 控制台没有出来，右上角搜索 即可

运用spring cloud框架基于spring boot构建微服务，一般需要启动多个应用程序，在idea开发工具中，多个同时启动的应用

需要在RunDashboard运行仪表盘中可以更好的管理，但有时候idea中的RunDashboard窗口没有显示出来，也找不到直接的开启按钮

idea中打开Run Dashboard的方法如下

　　　　view > Tool Windows > Run Dashboard

如果上述列表找不到Run Dashboard,则可以在工程目录下找到.idea文件夹下的workspace.xml，在其中相应位置加入以下代码（替换）即可：

```xml
<component name="RunDashboard">
    <option name="configurationTypes">
        <set>
            <option value="SpringBootApplicationConfigurationType"/>
        </set>
    </option>
    <option name="ruleStates">
        <list>
            <RuleState>
                <option name="name" value="ConfigurationTypeDashboardGroupingRule"/>
            </RuleState>
            <RuleState>
                <option name="name" value="StatusDashboardGroupingRule"/>
            </RuleState>
        </list>
    </option>
</component>
```





### 工程重构common

> 上面 两个子项目，有多次重复的 导入 jar，和重复的 Entity 实体类。可以把 多余的部分，加入到一个独立的模块中，将这个模块打包，并提供给需要使用的 module

1. 新建一个 cloud-api-commons 子模块
2. 将 entities 包里面的实体类放到这个子模块中，也将 pom 文件中，重复导入的 jar包放到这个新建的 模块的 pom 文件中。如下：

```xml
<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0"
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
    <parent>
        <artifactId>cloud2020</artifactId>
        <groupId>com.dkf.cloud</groupId>
        <version>1.0-SNAPSHOT</version>
    </parent>
    <modelVersion>4.0.0</modelVersion>

    <artifactId>cloud-api-commons</artifactId>

    <dependencies>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-devtools</artifactId>
            <scope>runtime</scope>
            <optional>true</optional>
        </dependency>
        <dependency>
            <groupId>org.projectlombok</groupId>
            <artifactId>lombok</artifactId>
            <optional>true</optional>
        </dependency>

        <!-- 这个是新添加的，之前没用到，后面会用到。关于这个hutool 是个功能强大的工具包，官网：https://hutool.cn/ -->
        <dependency>
            <groupId>cn.hutool</groupId>
            <artifactId>hutool-all</artifactId>
            <version>5.1.0</version>
        </dependency>
    </dependencies>

</project>
```

mvn跳过test，mvc clean，mvn install

将此项目打包 install 到 maven仓库。

3. 将 提供者 和 消费者 两个项目中的 entities 包删除，并删除掉加入到 cloud-api-commons 模块的 依赖配置。
4. 将 打包到 maven 仓库的 cloud-api-commons 模块，引入到 提供者 和 消费者的 pom 文件中，如下所示

```xml
<dependency><!-- 引入自己定义的api通用包，可以使用Payment支付Entity -->
    <groupId>com.dkf.cloud</groupId>
    <artifactId>cloud-api-commons</artifactId>
    <version>${project.version}</version>
</dependency>
```

# 二、服务注册中心

> 如果是上面只有两个微服务，通过 RestTemplate ，是可以相互调用的，但是当微服务项目的数量增大，就需要服务注册中心。目前没有学习服务调用相关技术，使用 SpringCloud 自带的 RestTemplate 来实现RPC

![1597213783265](https://i0.hdslb.com/bfs/album/8ba6cf17998c1de765f45bd28c8f0ffff21a1023.png)
## ==1、Eureka==

**什么是服务治理**：

SpringCloud封装了Netflix公司开发的Eureka模块来实现服务治理

在传统的rpc远程调用框架中，管理每个服务与服务之间依赖关系比较复杂，管理比较复杂，所以需要使用服务治理，管理服务于服务之间依赖关系，可以实现服务调用、负载均衡、容错等，实现服务发现与注册。

**什么是服务注册与发现**：

Eureka采用了CS的设计结构，Eureka Server服务注册功能的服务器，它是服务注册中心。而系统中的其他微服务，使用Eureka的客户端连接到Eureka Server并维持心跳连接。这样系统的维护人员就可以通过Eureka Server来监控系统中各个微服务是否正常运行。这点和zookeeper很相似

在服务注册与发现中，有一个注册中心。当服务器启动时候，会把当前自己服务器的信息 比如`服务地址` 通讯地址等以别名方式注册到注册中心上。另一方（消费者服务提供者），以该别名的方式去注册中心上获取到`实际的服务通讯地址`，然后再实现`本地RPC调用`。RPC远程调用框架核心设计思想：在于注册中心，因为便用注册中心管理每个服务与服务之间的一个依赖关系（服务治理概念)。在任何rpc远程框架中，都会有一个注册中心（存放服务地址相关信息（接口地址））

![1597291856582](https://i0.hdslb.com/bfs/album/d17cef49e815913d1d947b3c901305d737c8ede7.png)



> Eureka 官方停更不停用，以后可能用的越来越少。
>
> [Eureka](https://github.com/Netflix/Eureka) 是 [Netflix](https://github.com/Netflix) 开发的，一个基于 REST 服务的，服务注册与发现的组件，以实现中间层服务器的负载平衡和故障转移。
>
> Eureka 分为 Eureka Server 和 Eureka Client及服务端和客户端。**Eureka Server为注册中心，是服务端，而服务提供者和消费者即为客户端，消费者也可以是服务者，服务者也可以是消费者**。同时Eureka Server在启动时默认会注册自己，成为一个服务，所以Eureka Server也是一个客户端，这是搭建Eureka集群的基础。
>
> - Eureka Client：一个Java客户端，用于简化与 Eureka Server 的交互（通常就是微服务中的客户端和服务端）。通过注册中心进行访问。是一个Java客户端，用于简化Eureka Server的交互，客户端同时也具备一个内置的、使用轮询（round robin）负载算法的负载均衡器
>   在应用启动后，将会向Eureka Server发送心跳（默认周期为30秒）。如果Eureka Server在多个心跳周期内没有接收到某个节点的心跳，Eureka Server将会从服务注册表中把这个服务节点移除（默认90秒）
> - Eureka Server：提供服务注册服务，各个微服务节，通过配置启动后，会在Eureka Serverc中进行注册，这样Eureka Server中的服务注册表中将会存储所有可用服务节点信息，服务节点的信息可以在界面中直观看到。
>
> 服务在Eureka上注册，然后每隔30秒发送心跳来更新它们的租约。如果客户端不能多次续订租约，那么它将在大约90秒内从服务器注册表中剔除。注册信息和更新被复制到集群中的所有eureka节点。来自任何区域的客户端都可以查找注册表信息（每30秒发生一次）来定位它们的服务（可能在任何区域）并进行远程调用
>
> 服务**提供者**向注册中心注册服务，并**每隔30秒发送一次心跳**，就如同人还活着存在的信号一样，如果Eureka在90秒后还未收到服务提供者发来的心跳时，那么它就会认定该服务已经死亡就会注销这个服务。这里注销并不是立即注销，而是会在60秒以后对在这个之间段内“死亡”的服务集中注销，如果立即注销，势必会对Eureka造成极大的负担。这些时间参数都可以人为配置。
>
> Eureka还有`自我保护机制`，如果在15分钟内超过85%的节点都没有正常的心跳，那么Eureka就认为客户端与注册中心出现了网络故障，所以不会再接收心跳，也不会删除服务。
>
> 客户端消费者会向注册中心拉取服务列表，因为一个服务器的承载量是有限的，所以同一个服务会部署在多个服务器上，每个服务器上的服务都会去注册中心注册服务，他们会有相同的服务名称但有不同的实例id，所以**拉取的是服务列表**。我们最终通过负载均衡来获取一个服务，这样可以均衡各个服务器上的服务。



![在这里插入图片描述](https://img-blog.csdnimg.cn/20190816202635152.png)

### 单机版Eureka构建：

消费者端口80，提供者端口8001。

Eureka端口7001



#### 1) Server模块

pom

版本说明：

```xml
<!--新版本2020.02-->
<dependency>
    <groupId>org.springframework.cloud</groupId>
    <artifactId>spring-cloud-starter-netflix-eureka-server</artifactId>
</dependency>


<!--旧版本2018-->
<dependency>
    <groupId>org.springframework.cloud</groupId>
    <artifactId>spring-cloud-starter-eureka</artifactId>
</dependency>
```

```xml
	<artifactId>cloud-eureka-server7001</artifactId>

<dependencies>
  <dependency>
      <groupId>org.springframework.cloud</groupId>
      <artifactId>spring-cloud-starter-netflix-eureka-server</artifactId>
  </dependency>
    
  <dependency>
      <groupId>org.springframework.boot</groupId>
      <artifactId>spring-boot-starter-web</artifactId>
  </dependency>
  <dependency>
      <groupId>org.springframework.boot</groupId>
      <artifactId>spring-boot-starter-actuator</artifactId>
  </dependency>
  <dependency>
      <groupId>org.mybatis.spring.boot</groupId>
      <artifactId>mybatis-spring-boot-starter</artifactId>
  </dependency>
  <dependency>
      <groupId>org.springframework.boot</groupId>
      <artifactId>spring-boot-devtools</artifactId>
      <scope>runtime</scope>
      <optional>true</optional>
  </dependency>
  <dependency>
      <groupId>org.springframework.boot</groupId>
      <artifactId>spring-boot-starter-test</artifactId>
      <scope>test</scope>
  </dependency>
  <dependency>
      <groupId>com.dkf.cloud</groupId>
      <artifactId>cloud-api-commons</artifactId>
      <version>${project.version}</version>
  </dependency>
</dependencies>
```

application.yml

```yaml
server:
	port: 7001

eureka:
	instance:
		hostname: localhost  # eureka 服务端的实例名称

	client:
		# false 代表不向服务注册中心注册自己，因为它本身就是服务中心
		register-with-eureka: false
		# false 代表自己就是服务注册中心，自己的作用就是维护服务实例，并不需要去检索服务
		fetch-registry: false
		service-url:
			# 设置与 Eureka Server 交互的地址，查询服务 和 注册服务都依赖这个地址
			defaultZone: http://${eureka.instance.hostname}:${server.port}/eureka/
```

###### `@EnableEurekaServer`

最后写主启动类，如果启动报错，说没有配置 DataSource ，就在 主启动类的注解加上 这样的配置：

```java
// exclude ：启动时不启用 DataSource的自动配置检查
@SpringBootApplication(exclude = DataSourceAutoConfiguration.class)
@EnableEurekaServer   // 表示它是服务注册中心
public class EurekaServerMain7001 {
    public static void main(String[] args){
        SpringApplication.run(EurekaServerMain7001.class, args);
    }
}
```

启动测试，访问 7001 端口

#### 2) 提供者

> 这里的提供者，还是使用 上面的 cloud-provider-payment8001 模块，做如下修改：

1. 在 pom 文件的基础上引入 eureka 的client包，pom 的全部依赖如下所示：

```xml
<artifactId>cloud-provider-payment8001</artifactId>
<dependencies>
    <!--eureka-client-->
    <dependency>
        <groupId>org.springframework.cloud</groupId>
        <artifactId>spring-cloud-starter-netflix-eureka-client</artifactId>
    </dependency>

    <dependency>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-web</artifactId>
    </dependency>
    <dependency>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-actuator</artifactId>
    </dependency>
    <dependency>
        <groupId>org.mybatis.spring.boot</groupId>
        <artifactId>mybatis-spring-boot-starter</artifactId>
    </dependency>
    <dependency>
        <groupId>com.alibaba</groupId>
        <artifactId>druid-spring-boot-starter</artifactId>
        <version>1.1.10</version>
    </dependency>
    <!--mysql-connector-java-->
    <dependency>
        <groupId>mysql</groupId>
        <artifactId>mysql-connector-java</artifactId>
    </dependency>
    <!--jdbc-->
    <dependency>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-jdbc</artifactId>
    </dependency>
    <dependency>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-devtools</artifactId>
        <scope>runtime</scope>
        <optional>true</optional>
    </dependency>
    <dependency>
        <groupId>org.projectlombok</groupId>
        <artifactId>lombok</artifactId>
        <optional>true</optional>
    </dependency>
    <dependency>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-test</artifactId>
        <scope>test</scope>
    </dependency>
    <dependency>
        <groupId>com.dkf.cloud</groupId>
        <artifactId>cloud-api-commons</artifactId>
        <version>${project.version}</version>
    </dependency>
</dependencies>
```
###### `@EnableEurekaClient`
2. 主启动类 加上注解 ：` @EnableEurekaClient`
3. yml 文件添加关于 Eureka 的配置：

```yml
spring:
  application:
    name: cloud-payment-service # 项目名,也是注册的名字

eureka:
  client:
	# 注册进 Eureka 的服务中心
    register-with-eureka: true
    # 检索 服务中心 的其它服务
    fetch-registry: true
    service-url:
      # 设置与 Eureka Server 交互的地址
      defaultZone: http://localhost:7001/eureka/
```



#### 3) 消费者

> 这里的消费者 也是上面 的 cloud-customer-order80 模块

1. 修改 pom 文件，加入Eureka 的有关依赖， 全部 pom 依赖如下：

```xml
<artifactId>cloud-customer-order80</artifactId>

<dependencies>
    <dependency>
        <groupId>org.springframework.cloud</groupId>
        <artifactId>spring-cloud-starter-netflix-eureka-client</artifactId>
    </dependency>
    <dependency><!-- 引入自己定义的api通用包，可以使用Payment支付Entity -->
        <groupId>com.dkf.cloud</groupId>
        <artifactId>cloud-api-commons</artifactId>
        <version>${project.version}</version>
    </dependency>
    <dependency>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-web</artifactId>
    </dependency>
    <dependency>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-actuator</artifactId>
    </dependency>
    <dependency>
        <groupId>org.projectlombok</groupId>
        <artifactId>lombok</artifactId>
        <optional>true</optional>
    </dependency>
    <dependency>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-devtools</artifactId>
        <scope>runtime</scope>
        <optional>true</optional>
    </dependency>
    <dependency>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-test</artifactId>
        <scope>test</scope>
    </dependency>
</dependencies>
```

###### `@EnableEurekaClient`

2. 主启动类 加上注解 ： @EnableEurekaClient
3. yml 文件必须添加的内容：

```yml
eureka:
  client:
    register-with-eureka: true
    fetch-registry: true
    service-url:
      defaultZone: http://localhost:7001/eureka/
spring:
  application:
    name: cloud-order-service
```

- 1先启动eureka注册中心
- 2启动服务提供者payment支付服务
- 3支付服务启动后会把自身信息化 服务以别名方式注册进eureka
- 4消费者order服务在要调用接囗时，**使用服务别名去注册中心取实际的RPC远程调用地址**
- 5**消费者获得调用地址后，底层实际是利用`HttpClient`技术实现远程调用**
- 6**消费者获得服务地址后会存`jvm内存`中，默认每间隔30s更新一次服务调用地址**

Eureka Server在设计的时候就考虑了高可用设计，在Eureka服务治理设计中，**所有节点既是服务的提供方，也是服务的消费方**，服务注册中心也不例外。

Eureka Server的高可用实际上就是将自己做为服务向其他服务注册中心注册自己，这样就可以形成一组互相注册的服务注册中心，以实现服务清单的互相同步，达到高可用的效果。

 Eureka Server的同步遵循着一个非常简单的原则：只要有一条边将节点连接，就可以进行信息传播与同步。可以采用两两注册的方式实现集群中节点完全对等的效果，实现最高可用性集群，任何一台注册中心故障都不会影响服务的注册与发现。

### Eureka 集群700X



问题：微服务RPC远程服务调用最核心的是什么：

高可用，试想你的注册中心只有一个。onlyone,它出故障了那就呵呵了，会导致整个为服务环境不可用，所以要搭建Eureka注册中心集群，实现负载均衡+故障容错

Eureka 集群的原理：相互注册，互相守望。每台Eureka服务器都有集群里其他Eureka服务器地址的信息

开始构建Eureka集群：

现在创建 cloud-eureka-server7002 ，也就是第二个 Eureka 服务注册中心，pom 文件和 主启动类，与第一个Server一致。

> 模拟多个主机，修改`C:\Windows\System32\drivers\etc\hosts`  添加如下：
>
> 127.0.0.1 eureka7001.com
>
> 127.0.0.1 eureka7002.com

现在修改这两个 Server 的 yml 配置：

7001 端口的Server yml文件：

```yml
server:
  port: 7001

eureka:
  instance:
    hostname: eureka7001.com  # eureka 服务器的实例地址

  client:
    register-with-eureka: false
    fetch-registry: false
    service-url:
    ## 一定要注意这里的地址，这是搭建集群的关键。反过来写，写的是集群中其他Eureka服务器的地址
      defaultZone: http://eureka7002.com:7002/eureka/
```

7002 端口的Server yml文件：

```yml
server:
  port: 7002

eureka:
  instance:
    hostname: eureka7002.com  # eureka 服务器的实例地址

  client:
    register-with-eureka: false
    fetch-registry: false
    service-url:
    ## 一定要注意这里的地址 这是搭建集群的关键
      defaultZone: http://eureka7001.com:7001/eureka/
```

> eureka.instance.hostname 才是启动以后 本 Server 的注册地址，而 service-url  是 map 类型，只要保证 key:value 格式就行，它代表 本Server 指向了那些 其它Server 。利用这个，就可以实现Eureka Server 相互之间的注册，从而实现集群的搭建。

### 提供者集群800X

上面配置了多个Eureka作为集群，接下来要配置的是提供者集群，让提供者高可用

为提供者`cloud-provider-payment8001` 模块创建集群，新建模块名为 `cloud-provider-payment8002`

即两个提供者8001和8002

其余配置都一致，需要配置集群的配置如下：

配置区别要点：

- 集群中多个提供者的`spring:application:name:`要一致
- 启动类添加`@EnableDiscoveryClient`或者`@EnableEurekaClient`
   1，`@EnableDiscoveryClient`注解是基于`spring-cloud-commons`依赖，并且在classpath中实现；（推荐）
   2，`@EnableEurekaClient`注解是基于`spring-cloud-netflix`依赖，只能为eureka作用；
   如果你的classpath中添加了eureka，则它们的作用是一样的。

> 消费者（一般需要连接其他微服务的服务或者gateway/zuul）

```yml
# 提供者
server:
  port: 8001  # 端口号不一样

spring:
  application:
    name: cloud-provider-service  # 这次重点是这里，两个要写的一样，这是这个集群的关键
  datasource:
    type: com.alibaba.druid.pool.DruidDataSource
    driver-class-name: org.gjt.mm.mysql.Driver
    url: jdbc:mysql://localhost:3306/cloud2020?useUnicode=true&characterEncoding=utf-8&useSSL=false
    username: root
    password: 123456

mybatis:
  mapper-locations: classpath:mapper/*.xml
  type-aliases-package: com.dkf.springcloud.entities  

eureka:
  client:
    register-with-eureka: true
    fetch-registry: true
    service-url: # 提供者注册到多个eureka中
      defaultZone: http://eureka7001.com:7001/eureka/,http://eureka7002.com:7002/eureka/
```

注意在 Controller 返回不同的消息，从而区分者两个提供者的工作状态。（只是为了学习测试才这么做，生产环境直接复制即可）

在提供者的controller中

```java
@Value("${server.port}")
private String serverPort;
```



##### 消费者80

此时消费者一旦消费完之后，他以后访问的还是那台提供者。明显不对，原因在于消费者并没有去Rureka里找服务，而是自己找的

> 就是消费者如何访问 由这两个提供者组成的集群？

Eureka Server 上的提供者的服务名称如下：

```java
@RestController
@Slf4j
public class OrderController {
    
    // 重点是这里，改成 提供者在Eureka 上的名称，而且无需写端口号	
    public static final String PAYMENY_URL = "http://CLOUD-PROVIDER-SERVICE";//取决于我们在提供者出配置的name，CLOUD-PAYMENY-SERVICE，//同时要注意使用@LoadBalanced注解赋予RestTemplate负载均衡能力

    @Resource
    private RestTemplate restTemplate;

    @PostMapping("customer/payment/create")
    public CommonResult<Payment> create (Payment payment){
        // 需要自己rest远程调用
        return restTemplate.postForObject(PAYMENY_URL + "/payment/create", payment, CommonResult.class);
    }

    @GetMapping("customer/payment/{id}")
    public CommonResult<Payment> getPaymentById(@PathVariable("id")Long id){
        return restTemplate.getForObject(PAYMENY_URL + "/payment/" + id, CommonResult.class);
    }
}
```

```yaml
server:
  port: 80

spring:
    application:
        name: cloud-order-service
    zipkin:
      base-url: http://localhost:9411
    sleuth:
      sampler:
        probability: 1

eureka:
  client:
    #表示是否将自己注册进EurekaServer默认为true。
    register-with-eureka: false
    #是否从EurekaServer抓取已有的注册信息，默认为true。单节点无所谓，集群必须设置为true才能配合ribbon使用负载均衡
    fetchRegistry: true
    service-url:
      #单机
      #defaultZone: http://localhost:7001/eureka
      # 集群
      defaultZone: http://eureka7001.com:7001/eureka,http://eureka7002.com:7002/eureka  # 集群版
```

### 负载均衡

`@LoadBalanced`

还有，消费者里面对RestTemplate配置的config文件，需要更改成如下：（就是加一个注解 `@LoadBalanced`）

```java
@Configuration
public class ApplicationContextConfig {

    @Bean
    @LoadBalanced  //这个注解，就赋予了RestTemplate 负载均衡的能力
    public RestTemplate getRestTemplate(){
        return new RestTemplate();
    }
}
```

这时候，消费者消费的提供者多次访问就会变化了（这就是Ribbon的负载平衡功能）

### actuator让Eureka显示ip

为了在微服务Eureka控制台能看到我们的某个具体服务是在哪台服务器上部署的，我们需要配置一些内容。

修改 提供者在Eureka 注册中心显示的 主机名：即修改`eureka:instance:instance-id:`和`eureka:instance:prefer-ip-address:`

```yaml
# 提供者
server:
  port: 8001  # 端口号不一样

eureka:
  client:
    register-with-eureka: true
    fetch-registry: true
    service-url:
      defaultZone: http://eureka7001.com:7001/eureka/,http://eureka7002.com:7002/eureka/
  instance: #重点，和client平行
    instance-id: payment8001 # 每个提供者的id不同，显示的不再是默认的项目名
    prefer-ip-address: true # 可以显示ip地址
```

![1597301299700](https://i0.hdslb.com/bfs/album/39445ef7e3793b4ef3c5762b46bbd65d70c7e011.png)

![1597301396971](https://i0.hdslb.com/bfs/album/020f16542e99c807cde3b4a939fc3e6e2fe0c2be.png)

### 服务发现Discovery`@EnableDiscoveryClient`

对于注册进eureka里面的微服务，可以通过服务发现来获得该服务的信息。（即我们前面可视化页面的信息）

1. 在主启动类上添加注解：`@EnableDiscoveryClient`
2. 在 Controller 里面打印信息：

```java
@Resource // 自动注入
private DiscoveryClient discoveryClient;

@GetMapping("/customer/discovery")
public Object discovery(){
    //获得服务清单列表
    List<String> services = discoveryClient.getServices();
    for(String service: services){
        log.info("*****service: " + service);
    }
    
    // 根据具体服务进一步获得该微服务的信息
    List<ServiceInstance> instances = discoveryClient.getInstances("CLOUD-ORDER-SERVICE");
    for(ServiceInstance serviceInstance:instances){
        log.info(serviceInstance.getServiceId() + "\t" + serviceInstance.getHost()
                 + "\t" + serviceInstance.getPort() + "\t" + serviceInstance.getUri());
    }
    return this.discoveryClient;
}
```

### Eureka 自我保护机制

什么是自我保护模式？

默认情况下，如果Eureka Server在一定时间内没有接收到某个微服务实例的心跳，Eureka Server将会注销该实例（默认90秒）。但是当网络分区故障发生时、卡顿、拥挤）时，微服务与Eureka Server之间无法正常通信，以上行为可能变得非常危险了---因为微服务本身其实是健康的，此时本不应该注销这个微服务。Eureka通过"自我保护模式"来解决这个问题---==当Eureka Server节点在短时间内丢失过多客户端时（可能发生了网络分区故障），那么这个节点就会进入自我保护模式==。

> 属于CAP里的`AP`（高可用）分支

自我保护机制：默认情况下Eureka CIient定时向Eureka Server端发送心跳包。

如果Eureka在server端在一定时间内（默认90秒）没有收到Eureka Client发送心跳包，便会直接从服务注册列表中剔除该服务，但是在短时间（90秒内）内丢失了大量的服务实例心跳，这时Eureka Server会开启自我保护机制，不会剔除该服务（该现象可能出现如果网络不通但是Eureka Client出现宕机，此时如果别的注册中心如果一定时间内没有收到心跳会将剔除该服务这样就出现了严重失误，因为客户端还能正常发送心跳，只是网络延迟问题，而保护机制是为了解决此问题而产生的

在自我保护模式中，Eureka Server会保护服务注册表中的信息，不再注册任何服务实例。

它的设计哲学就是宁可保留错误的服务注册信息，也不盲目注销任何可能健康的服务实例。一句话讲解：好死不如赖活着

综上，自我保护模式是一种应对网络异常的安全保护措施。它的架构哲学是宁可同时保留所有微服务（健康的微服务和不健康的微服务都会保留）也不盲目注销任何健康的微服务。使用自我保护模式，可以让Eureka集群更加的健壮、稳定。



保护模式主要用于一组客户和Eureka Server之间存在网络分区场景下保护。一旦进入保护模式，Eureka Server将会尝试保护其服务注册表中的信息，不再删除服务注册表中固定信息，也就是不会注销任何微服务。

如果在Eureka Server的首页看到以下这段提示，则说明Eureka讲入了保护模式：

> EMERGENCY!EUREKA MAY BE INCORRECTLY CLAIMING INSTANCES ARE UP WHEN THEY'RE NOT.
> RENEWALS ARE LESSER THAN THRESHOLD AND HENCE THE INSTANCES ARE NOT BEING EXPIRED JUST TO BE SAFE

为什么会产生Eureka自我保护机制？

为了防止Eureka Client可以正常运行但是与Eureka Server网络不通情况下，Eureka Server不会立刻将Eureka Client服务剔除



禁止自我保护:（如果想）

> 在 Eureka Server 的模块中的 yml 文件进行配置：
>
> ```yml
> server:
>   port: 7001
> 
> eureka:
>   instance:
>     hostname: eureka7001.com
> 
>   client:
>     register-with-eureka: false
>     fetch-registry: false
>     service-url:
>       defaultZone: http://eureka7002.com:7002/eureka/
>   server: # 与client平行
>   	# 关闭自我保护机制，保证不可用该服务被及时剔除
>   	enable-self-preservation: false
>   	eviction-interval-timer-in-ms: 2000
> ```



修改 Eureka Client 模块的 心跳间隔时间：

```yaml
# 提供者
server:
  port: 8001  # 端口号不一样

eureka:
  client:
    register-with-eureka: true
    fetch-registry: true
    service-url: # 集群
      defaultZone: http://eureka7001.com:7001/eureka/,http://eureka7002.com:7002/eureka/
  instance: #重点，和client平行
    instance-id: payment8001 # 每个提供者的id不同，显示的不再是默认的项目名
    prefer-ip-address: true # 可以显示ip地址
    # Eureka客户端像服务端发送心跳的时间间隔，单位s，默认30s
    least-renewal-interval-in-seconds: 1
    # Rureka服务端在收到最后一次心跳后等待时间上线，单位为s，默认90s，超时将剔除服务
    least-expiration-duration-in-seconds: 2
```

##### eureka配置项解读：

在注册服务之后，服务提供者会维护一个心跳用来持续高速Eureka Server，“我还在持续提供服务”，否则Eureka Server的剔除任务会将该服务实例从服务列表中排除出去。我们称之为服务续约。
 面是服务续约的两个重要属性：

（1）eureka.instance.lease-expiration-duration-in-seconds
 leaseExpirationDurationInSeconds，表示eureka server至上一次收到client的心跳之后，等待下一次心跳的超时时间，在这个时间内若没收到下一次心跳，则将移除该instance。
 默认为90秒
 如果该值太大，则很可能将流量转发过去的时候，该instance已经不存活了。
 如果该值设置太小了，则instance则很可能因为临时的网络抖动而被摘除掉。
 该值至少应该大于leaseRenewalIntervalInSeconds

（2）eureka.instance.lease-renewal-interval-in-seconds
 leaseRenewalIntervalInSeconds，表示eureka client发送心跳给server端的频率。如果在leaseExpirationDurationInSeconds后，server端没有收到client的心跳，则将摘除该instance。除此之外，如果该instance实现了HealthCheckCallback，并决定让自己unavailable的话，则该instance也不会接收到流量。
 默认30秒

> ***eureka.client.registry-fetch-interval-seconds\*** :表示eureka client间隔多久去拉取服务注册信息，默认为30秒，对于api-gateway，如果要迅速获取服务注册状态，可以缩小该值，比如5秒
>  ***eureka.server.enable-self-preservation\***
>  是否开启自我保护模式，默认为true。
>  默认情况下，如果Eureka Server在一定时间内没有接收到某个微服务实例的心跳，Eureka Server将会注销该实例（默认90秒）。但是当网络分区故障发生时，微服务与Eureka Server之间无法正常通信，以上行为可能变得非常危险了——因为微服务本身其实是健康的，此时本不应该注销这个微服务。
>  Eureka通过“自我保护模式”来解决这个问题——当Eureka Server节点在短时间内丢失过多客户端时（可能发生了网络分区故障），那么这个节点就会进入自我保护模式。一旦进入该模式，Eureka Server就会保护服务注册表中的信息，不再删除服务注册表中的数据（也就是不会注销任何微服务）。当网络故障恢复后，该Eureka Server节点会自动退出自我保护模式。
>  综上，自我保护模式是一种应对网络异常的安全保护措施。它的架构哲学是宁可同时保留所有微服务（健康的微服务和不健康的微服务都会保留），也不盲目注销任何健康的微服务。使用自我保护模式，可以让Eureka集群更加的健壮、稳定。
>  ***eureka.server.eviction-interval-timer-in-ms\***
>  eureka server清理无效节点的时间间隔，默认60000毫秒，即60秒

Eureka停更说明：

2.0后停更了。



## ==2、Zookeeper==

springCloud 整合 zookeeper 

- zookeeper是一个分布式协调工具，可以实现注册中心功能
- 关闭Linux服务器防火墙后动zookeeper服务器
- zookeeper服务器取代Eureka服务器，zk作为服务注册中心

##### 提供者

新建模块cloud-provider-payment8004

pom文件如下：

```xml
<artifactId>cloud-provider-payment8004</artifactId>

<dependencies>
    <!--springcloud 整合 zookeeper 组件-->
    <dependency>
        <groupId>org.springframework.cloud</groupId>
        <!--zk发现-->
        <artifactId>spring-cloud-starter-zookeeper-discovery</artifactId>
        <exclusions>
            <exclusion>
                <groupId>org.apache.zookeeper</groupId>
                <artifactId>zookeeper</artifactId>
            </exclusion>
        </exclusions>
    </dependency>
    <dependency>
        <groupId>org.apache.zookeeper</groupId>
        <artifactId>zookeeper</artifactId>
        <version>3.4.9</version>
        <exclusions>
            <exclusion>
                <groupId>org.slf4j</groupId>
                <artifactId>slf4j-log4j12</artifactId>
            </exclusion>
        </exclusions>
    </dependency>

    <dependency>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-web</artifactId>
    </dependency>
    <dependency>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-actuator</artifactId>
    </dependency>
    <dependency>
        <groupId>org.mybatis.spring.boot</groupId>
        <artifactId>mybatis-spring-boot-starter</artifactId>
    </dependency>
    <dependency>
        <groupId>com.alibaba</groupId>
        <artifactId>druid-spring-boot-starter</artifactId>
        <version>1.1.10</version>
    </dependency>
    <!--mysql-connector-java-->
    <dependency>
        <groupId>mysql</groupId>
        <artifactId>mysql-connector-java</artifactId>
    </dependency>
    <!--jdbc-->
    <dependency>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-jdbc</artifactId>
    </dependency>
    <dependency>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-devtools</artifactId>
        <scope>runtime</scope>
        <optional>true</optional>
    </dependency>
    <dependency>
        <groupId>org.projectlombok</groupId>
        <artifactId>lombok</artifactId>
        <optional>true</optional>
    </dependency>
    <dependency>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-test</artifactId>
        <scope>test</scope>
    </dependency>
    <dependency><!-- 引入自己定义的api通用包，可以使用Payment支付Entity -->
        <groupId>com.dkf.cloud</groupId>
        <artifactId>cloud-api-commons</artifactId>
        <version>${project.version}</version>
    </dependency>
</dependencies>
```

yaml

```yml
server:
  port: 8004

spring:
  application:
    name: cloud-provider-service
  cloud:
    zookeeper:
      connect-string: 192.168.40.100:2181 # zk地址
```

主启动类：

```java
@SpringBootApplication
@EnableDiscoveryClient	 // 以后用这个就可以了，不用eureka了
public class PaymentMain8004 {
    public static void main(String[] args){
        SpringApplication.run(PaymentMain8004.class, args);
    }
}
```

Controller 打印信息：

```java
@RestController
@Slf4j
public class PaymentController {

    @Resource
    private PaymentService paymentService;

    @Value("${server.port}")
    private String serverPort;

    @RequestMapping("/payment/zk")
    public String paymentzk(){
        return "springcloud with zookeeper :" + serverPort + "\t" + UUID.randomUUID().toString();
    }
}
```

如果 zookeeper 的版本和导入的jar包版本不一致，启动就会报错，由zk-discovery和zk之间的jar包冲突的问题。

解决这种冲突，需要在 pom 文件中，排除掉引起冲突的jar包，添加和服务器zookeeper版本一致的 jar 包，

但是新导入的 zookeeper jar包 又有 slf4j 冲突问题，于是再次排除引起冲突的jar包

```xml
<!--springcloud 整合 zookeeper 组件-->
<dependency>
    <groupId>org.springframework.cloud</groupId>
    <artifactId>spring-cloud-starter-zookeeper-discovery</artifactId>
    <!-- 排除与zookeeper版本不一致到导致 冲突的 jar包 -->
    <exclusions>
        <exclusion>
            <groupId>org.apache.zookeeper</groupId>
            <artifactId>zookeeper</artifactId>
        </exclusion>
    </exclusions>
</dependency>
<!-- 添加对应版本的jar包 -->
<dependency>
    <groupId>org.apache.zookeeper</groupId>
    <artifactId>zookeeper</artifactId>
    <version>3.4.9</version>
    <!-- 排除和 slf4j 冲突的 jar包 -->
    <exclusions>
        <exclusion>
            <groupId>org.slf4j</groupId>
            <artifactId>slf4j-log4j12</artifactId>
        </exclusion>
    </exclusions>
</dependency>
```

启动测试：

```sh
# 在zk客户端
ls /services # 输入
[cloud-provider-service] #输出
# 继续向下查看
ls /services/cloud-provider-service
# 继续向下，然后get，返回了个json
```

我们在zk上注册的node是临时节点,当我们的服务一定时间内没有发送心跳，那么zk就会将这个服务的znode删除了。没有自我保护机制。重新建立连接后znode-id号也会变

##### 消费者

> 创建测试zookeeper作为服务注册中心的 消费者 模块 cloud-customerzk-order80
>
> 主启动类、pom文件、yml文件和提供者的类似

config类，注入 RestTemplate

```java
@SpringBootConfiguration
public class ApplicationContextConfig {
    @Bean
    @LoadBalanced // 继续加上这个
    public RestTemplate getTemplate(){
        return new RestTemplate();
    }
}

@SpringBootApplication
@EnableDiscoveryClient
public class OrderZKMain80{
    public static void main(String[] args) {
        SpringApplication.run(OrderZKMain80.class, args);
    }
}

```

controller层也是和之前类似：

```java
@RestController
@Slf4j
public class CustomerZkController {

    public static final String INVOKE_URL="http://cloud-provider-service"; //和原来一样

    @Resource
    private RestTemplate restTemplate;

    @RequestMapping("/customer/payment/zk")
    public String paymentInfo(){
        String result = restTemplate.getForObject(INVOKE_URL + "/payment/zk",String.class);
        return result;
    }
}
```

然后就在zk里查到consumer信息了。

关于 zookeeper 的集群搭建，目前使用较少，而且在 yml 文件中的配置也是类似，以列表形式写入 zookeeper 的多个地址即可，而且zookeeper 集群，在 hadoop的笔记中也有记录。总而言之，只要配合zookeeper集群，以及yml文件的配置就能完成集群搭建

后面会用ribbon代替RestTemplate



## ==3、Consul==

> consul也是服务注册中心的一个实现，是由go语言写的。官网地址： https://www.consul.io/intro   中文地址： https://www.springcloud.cc/spring-cloud-consul.html 
> Consul是一套开源的分布式服务发现和配置管理系统。
提供了微服务系统中的服务治理，配置中心，控制总线等功能。这些功能中的每一个都可以根据需要单独使用，也可以一起使用以构建全方位的服务网络。




- 服务发现：提供HTTP和DNS两种发现方式
- 健康监测：支持多种方式，HTTP、TCP、Docker、Shell脚本定制化
- KV存储：Key、Value的存储方式
- 多数据中心：Consul支持多数据中心
- 可视化Web界面

### 安装并运行

>  下载地址：https://www.consul.io/downloads.html
>
>  打开下载的压缩包，只有一个exe文件，实际上是不用安装的，在exe文件所在目录打开dos窗口使用即可。
>
>  使用开发模式启动：consul agent -dev
>
>  访问8500端口，即可访问首页

### 提供者

> 新建提供者模块：cloud-providerconsul-service8006 

pom 文件：

```xml
	<artifactId>cloud-providerconsul-service8006</artifactId>
    <dependencies>
        <!--springcloud consul server-->
        <dependency>
            <groupId>org.springframework.cloud</groupId>
            <artifactId>spring-cloud-starter-consul-discovery</artifactId>
        </dependency>

        <!-- springboot整合Web组件 -->
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-web</artifactId>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-actuator</artifactId>
        </dependency>

        <!-- 日常通用jar包 -->
        <dependency>
            <groupId>org.mybatis.spring.boot</groupId>
            <artifactId>mybatis-spring-boot-starter</artifactId>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-devtools</artifactId>
            <scope>runtime</scope>
            <optional>true</optional>
        </dependency>
        <dependency>
            <groupId>org.projectlombok</groupId>
            <artifactId>lombok</artifactId>
            <optional>true</optional>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-test</artifactId>
            <scope>test</scope>
        </dependency>
        <dependency><!-- 引入自己定义的api通用包，可以使用Payment支付Entity -->
            <groupId>com.dkf.cloud</groupId>
            <artifactId>cloud-api-commons</artifactId>
            <version>${project.version}</version>
        </dependency>
    </dependencies>
```

yml 文件：

```yml
server:
  port: 8006
spring:
  application:
    name: consul-provider-service
  cloud:
    consul:
      host: localhost
      port: 8500
      discovery:    # 指定注册对外暴露的服务名称
        service-name: ${spring.application.name}
```

主启动类：

```java
@SpringBootApplication
@EnableDiscoveryClient // 提供者
public class ConsulProviderMain8006 {
    public static void main(String[] args) {
        SpringApplication.run(ConsulProviderMain8006.class,args);
    }
}
```

controller也是简单的写一下就行。

### 消费者

> 新建 一个 在82端口的 消费者模块。pom和yml和提供者的类似，主启动类不用说，记得注入RestTemplate

```java
@SpringBootApplication
@EnableDiscoveryClient //该注解用于向使用consul或者zookeeper作为注册中心时注册服务
public class OrderConsulMain80{
    public static void main(String[] args) {
            SpringApplication.run(OrderConsulMain80.class, args);
    }
}


@Configuration
public class ApplicationContextConfig{
    @Bean
    @LoadBalanced
    public RestTemplate getRestTemplate()
    {
        return new RestTemplate();
    }
}
```

controller层：

```java
@RestController
public class CustomerConsulController {

    public static final String INVOKE_URL="http://consul-provider-service";

    @Resource
    private RestTemplate restTemplate;

    @RequestMapping("/customer/payment/consul")
    public String paymentInfo(){
        String result = restTemplate.getForObject(INVOKE_URL + "/payment/consul",String.class);
        return result;
    }
}
```

## 总结

| 组件名    | 语言 | CAP  | 服务健康检查 | 对外暴露接口 | SpringCloud集合 |
| --------- | ---- | ---- | ------------ | ------------ | --------------- |
| Eureka    | java | AP   | 可配支持     | HTTP         | 已集成          |
| Consul    | Go   | CP   | 支持         | HTTP/DNS     | 已集成          |
| Zookeeper | java | CP   | 支持         | 客户端       | 已集成          |

CAP：

- C：Consitency 强一致性
- A：Available 可用性
- P：Partition tolerance 分区容错性

CAP理论关注粒度是数据，而不是整体系统设计的



![ ](https://i0.hdslb.com/bfs/album/333d8e66cd0511b31f7026919ab516cc6d2cd667.png)



<img src="https://i0.hdslb.com/bfs/album/0704481274e95054e59d08d12da625b2e458fd20.png" alt="1597384508291" style="zoom:67%;" /><img src="https://i0.hdslb.com/bfs/album/c905344e142fb5533ddb09e1b25d5ccceedf19f1.png" alt="1597384554249" style="zoom:67%;" />

上面讲了服务注册中心，下面讲服务调用

# 三、服务调用


RestTemplate不是给我们提供远程调用了吗，那还要学其他的做什么。答案是**负载均衡**，在发送请求时通过负载均衡算法从提供者列表中选择一个。
> 都是使用在 client端，即有 ”消费者“ 需求的模块中。
![1597213783265](https://i0.hdslb.com/bfs/album/8ba6cf17998c1de765f45bd28c8f0ffff21a1023.png)
## ==1、Ribbon==负载均衡

我们这里提前启动好之前在搭建的 eureka Server 集群（5个模块）

### 简介

`SpringCloud Ribbon`是基于`NetfIix Ribbon`实现的一套**客户端**负载均衡的工具。

简单的说，Ribbon是Nefix发布的开源项目，主要功能是提供**客户端的软件负载均衡算法和服务调用**。Ribbon客户端组件提供一系列完善的配置项如**连接超时，重试**等。简单的说，就是在配置文件中列出LoadBalancer（简称LB)后面所有的机器，Ribbon会自动的帮助你基于某种规则（如简单轮询，随机连接等）去连接这些机器。我们很容易使用Ribbon实现自定义的负载均衡算法。

> LB负载均衡(LoadBalance)是什么？
>
> 简单的说就是将用户的请求平摊的分配到多个服务上，从而达到系统的HA（高可用）。
>
> 常见的负载均衡有软件Nginx，LVS，硬件F5等。

Ribbon本地负载均衡客户端 `VS` Nginx服务端负载均衡区别：

- Nginx是**服务器负载均衡**（集中式LB），客户端所有请求都会交给nginx，然后由nginx实现转发请求。即负载均衡是由服务端实现的。
- Ribbon是**本地负载均衡**（进程内LB），在调用微服务接口时候，会**在注册中心上获取注册信息服务列表之后缓存到JVM本地**，从而在本地实现RPC远程服务调用技术。

- 集中式LB：即在服务的消费方和提供方之间使用独立的LB设施（可以是硬件，如F5；也可以是软件，如nginx)，由该设施负责把访问请求通过某种策略转发至服务的提供方；
- 进程内LB：将LB逻辑集成到消费方，消费方从服务注册中心获知有哪些地址可用，然后自己再从这些地址中选择出一个合适的服务器！==Ribbon就属于进程内LB==，它只是一个类库，集成于消费方进程，==消费方通过它来获取到服务提供方的地址==。

<img src="https://i0.hdslb.com/bfs/album/95fb21c2c4033fa8697003cfcde30b85591cde82.png" alt="1597385250463" style="zoom:67%;" />

Ribbon在工作时分成两步：

- 第一步先选择Eureka Server，它优先选择在同一个区域内负载较少的server
- 第二步再根据用户指定的策略，在从server取到的服务注册列表中选择一个地址。
  - 其中Ribbon提供了多种策略：比如轮询、随相和根据响应时间加权。

上面在eureka时，确实实现了负载均衡机制，那是因为 **eureka-client包里面自带着ribbon**：

> 一句话，**Ribbon 就是 负载均衡 + RestTemplate 调用**。
>
> 实际上不止eureka的jar包有，zookeeper、consul、nacos的jar包都包含了他，就是上面使用的服务调用。
>
> ![1597385486515](https://i0.hdslb.com/bfs/album/633781c47860863861b079018f13d8c44457a425.png)
>
> 如果自己添加，在 模块的 pom 文件中引入：
>
> ```xml
> <dependency>
>     <groupId>org.springframework.cloud</groupId>
>     <artifactId>spring-cloud-starter-netflix-ribbon</artifactId>
> </dependency>
> ```
>

##### RestTemplate

对于RestTemplate 的一些说明：

> 有两种请求方式：post和get ,还有两种返回类型：object 和 Entity 

- getForObject()/getForEntity()
  - Object：返回对象响应体中数据转化成的对象，基本上可以理解成json
  - Entity：返回对象是ResponseEntity对象，包含了响应中的一些重要信息，比如响应头、响应状态码、响应体等
  - 返回的`entity.getBody()`即得到了Object
- postForObject()/postForEntity()

```java
package com.atguigu.springcloud.controller;

@RestController
@Slf4j
public class OrderController{
    //public static final String PAYMENT_URL = "http://localhost:8001";

    public static final String PAYMENT_URL = "http://CLOUD-PAYMENT-SERVICE";

    @Resource
    private RestTemplate restTemplate;

    @Resource // 负载均衡
    private LoadBalancer loadBalancer;
    @Resource
    private DiscoveryClient discoveryClient;

    @GetMapping("/consumer/payment/create")
    public CommonResult<Payment> create(Payment payment){
        return restTemplate.postForObject(PAYMENT_URL +"/payment/create",payment,CommonResult.class);
    }

    @GetMapping("/consumer/payment/get/{id}")
    public CommonResult<Payment> getPayment(@PathVariable("id") Long id){
        return restTemplate.getForObject(PAYMENT_URL+"/payment/get/"+id,CommonResult.class);
    }

    @GetMapping("/consumer/payment/getForEntity/{id}")
    public CommonResult<Payment> getPayment2(@PathVariable("id") Long id){
       //   ResponseEntity是spring中的类
        ResponseEntity<CommonResult> entity = restTemplate.getForEntity(PAYMENT_URL+"/payment/get/"+id,CommonResult.class);

        if(entity.getStatusCode().is2xxSuccessful()){
            return entity.getBody();//获取json
        }else{
            return new CommonResult<>(444,"操作失败");
        }
    }

    // ribbon用法
    @GetMapping(value = "/consumer/payment/lb")
    public String getPaymentLB(){
        // 找到服务列表
        List<ServiceInstance> instances = discoveryClient.getInstances("CLOUD-PAYMENT-SERVICE");

        if(instances == null || instances.size() <= 0){
            return null;
        }
        // 负载均衡中选择一个服务
        ServiceInstance serviceInstance = loadBalancer.instances(instances);
        URI uri = serviceInstance.getUri();

        return restTemplate.getForObject(uri+"/payment/lb",String.class);
    }

    // ====================> zipkin+sleuth
    @GetMapping("/consumer/payment/zipkin")
    public String paymentZipkin(){
        String result = restTemplate.getForObject("http://localhost:8001"+"/payment/zipkin/", String.class);
        return result;
    }
}
```

### 负载均衡

IRule：根据特定算法从服务列表中选择一个要访问的服务

Ribbon 负载均衡规则类型：

- com.netflix.loadbalancer.RoundRobinRule：轮询
- com.netflix.loadbalancer.RandomRule：随机
- com.netfIix.IoadbaIancer.RetryRuIe：先按照RoundRobinRule的策略获取服务，如果获取服务失败则在指定时间内会进行重试，获取可用的服务
- WeightedResponseTimeRule：对RoundRobinRule的扩展，响应速度越快的实例选择权重越大，越容易被选择
- BestAvailableRule：会先过滤掉由于多次访问故障而处于断路器跳闸状态的服务，然后选择一个并发量最小的服务
- AvailabilityFilteringRule：先过滤掉故障实例，再选择并发较小的实例
- ZoneAvoidanceRule：默认规则，复合判断server所在区域的性能和server的可用性选择服务器

![](https://i0.hdslb.com/bfs/album/583e87bfae6d2ceeb98801dbb96d3eb2c4c94c96.png)

配置负载均衡规则：

官方文档明确给出了警告：

> 这个自定义配置类不能放在@ComponentScan 所扫描的当前包下以及子包下，否则我们自定义的这个配置类就会被所有的Ribbon客户端所共享，达不到特殊化定制的目的了。

注意上面说的，而Springboot主启动类上的 @SpringBootApplication 注解，相当于加了@ComponentScan注解，会自动扫描当前包及子包，所以注意不要放在SpringBoot主启动类的包内。

创建包：

- java
  - com.hh
    - myrule
      - MySelfRule.java
    - springcloud
      - 主启动类

在这个包下新建 MySelfRule类：

```java
package com.dkf.myrule;

import com.netflix.loadbalancer.IRule;
import com.netflix.loadbalancer.RandomRule;

@Configuration
public class MySelfRule {
    @Bean // 定义负载均衡规则
    public IRule myrule(){
        return new RandomRule(); //负载均衡规则定义为随机
    }
}
```

然后在主启动类上添加如下注解 @RibbonClient：

```java
package com.dkf.springcloud;

import com.dkf.myrule.MySelfRule;

@SpringBootApplication
@EnableEurekaClient
@EnableDiscoveryClient
@RibbonClient(name="CLOUD-PROVIDER-SERVICE", 
              configuration = MySelfRule.class)//指定该负载均衡规则对哪个提供者服务使用 ， 加载自定义规则的配置类
public class OrderMain80 {

    public static void main(String[] args){
        SpringApplication.run(OrderMain80.class, args);
    }
}
```

### 轮询算法原理

**负载均衡轮询算法** ：

rest接口第几次请求次数 % 服务器集群总数量 = 实际调用服务器位置下标

每次服务器重启后，rest接口计数从1开始。

![1597387609476](https://i0.hdslb.com/bfs/album/57761cdb574b8b18581e4e99e0d9400f3e416ebb.png)

ribbon源码：

```java
private AtomicInteger nextServerCyclicCounter;

public Server choose(ILoadBalancer lb, Object key) {

    Server server = null;
    int count = 0;
    while (server == null && count++ < 10) {
        List<Server> reachableServers = lb.getReachableServers();
        List<Server> allServers = lb.getAllServers();
        int upCount = reachableServers.size();
        int serverCount = allServers.size();

        int nextServerIndex = incrementAndGetModulo(serverCount);
        server = allServers.get(nextServerIndex);

        if (server == null) {
            /* Transient. */
            Thread.yield();
            continue;
        }

        if (server.isAlive() && (server.isReadyToServe())) {
            return (server);
        }

        // Next.
        server = null;
    }

    if (count >= 10) {
        log.warn("No available alive servers after 10 tries from load balancer: "  + lb);
    }
    return server;
}
private int incrementAndGetModulo(int modulo) {
    for (;;) {
        int current = nextServerCyclicCounter.get();//获取原子的值
        int next = (current + 1) % modulo;
        if (nextServerCyclicCounter.compareAndSet(current, next)) //CAS
            return next;
    }
}
```

手写负载算法：cas+自旋

首先8001、8002服务controller层加上

```java
@GetMapping("/payment/lb")
public String getPaymentLB() {
    return SERVER_PORT;
}
```

LoadBalancer接口：

```java
import org.springframework.cloud.client.ServiceInstance;

public interface LoadBalancer {
    ServiceInstance instances(List<ServiceInstance> serviceInstances);
}
```

实现

```java
import org.springframework.cloud.client.ServiceInstance;
import java.sql.SQLOutput;

@Component
public class MyLB implements LoadBalancer {

    private AtomicInteger atomicInteger = new AtomicInteger(0);

    private final int getAndIncrement() {
        int current;
        int next;

        do {
            current = this.atomicInteger.get();
            next = current >= Integer.MAX_VALUE ? 0 : current + 1;
        } while (!atomicInteger.compareAndSet(current, next));
        System.out.println("第几次访问,次数next:" + next);
        return next;
    }

    @Override
    public ServiceInstance instances(List<ServiceInstance> serviceInstances) {
        int index = getAndIncrement() % serviceInstances.size();
        return serviceInstances.get(index);
    }
}
```

controller类中添加：

```java
@GetMapping("/consumer/payment/lb")
public String getPaymentLB() {
    List<ServiceInstance> instances = discoveryClient.getInstances("CLOUD-PAYMENT-SERVICE");//获得总的提供者数
    if (instances == null || instances.size() <= 0) {
        return null;
    }

    ServiceInstance serviceInstance = loadBalancer.instances(instances);//传入总的实例数
    URI uri = serviceInstance.getUri();

    return restTemplate.getForObject(uri + "/payment/lb", String.class);
}
```



## ==2、OpenFeign==

### 概述

> 这里和之前学的dubbo很像，例如消费者的controller 可以调用提供者的 service层方法，但是不一样，它只能调用提供者的 controller（ribbon直接调用的是service），即写一个提供者项目的controller的接口，消费者来调用这个接口方法，就还是相当于是调用提供者的 controller ，和RestTemplate 没有本质区别

Feign能干什么：

Feign旨在使编写JavaHttp客户端变得更容易。

前面在使用Ribbon+RestTemplate时，利用RestTemplate对http请求的封装处理，形成了一套模版化的调用方法。但是在实际开发中，由于对服务依赖的调用可能不止一处，往往一个接囗会被多处调用，所以通常都会针对每个微服务自行封装些客户端类来包装这些依赖服务的调用。所以，Feign在此基础上做了进一步封装，由他来帮助我们定义和实现依赖服务接口的定义。在Feign的实现下，我们**只需创建一个接口并使用注解的方式来配置它**（以前是Dao接口上面标注Mapper注解，现在是一个微服务接口上面标注一个Feign注解即可，即可完成对服务提供方的接口绑定，简化了使用Springcloud Ribbon时，自动封装服务调用客户端的开发量。

Feign集成了Ribbon：

利用Ribbon维护了Payment的服务列表信息，并目通过轮询实现了客户端的负载均衡。而与Ribbon不同的是，通过feign只需要定义服务绑定接口目以声明式的方法，优雅而简单的实现了服务调用

| Feign                                                        | OpenFeign                                                    |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| Feign是SpringCloud组件中的一个轻量级RESTful的HTTP服务客户端。Feign内置了Ribbon，用来做客户端负载均衡，去调用服务注册中心的服务。Feign的使用方式是：使用Feign的注解定义接口，调用这个接口，就可以调用服务注册中心的服务 | OpenFeign是SpringCloud在Feign的基础上支持了SpringMVC的注解，如@RequestMapping等等。OpenFeign的`@FeignClient`可以解析SpringMVC的下的接囗，并通过动态代理的方式产生实现类，实现类中做负载均衡并调用其他服务。 |
| <groupId>org.springframework.cloud</groupId>     <artifactId>spring-cloud-starter-feign</artifactId> | <groupId>org.springframework.cloud</groupId>     <artifactId>spring-cloud-starter-openfeign</artifactId> |

### 消费端

新建cloud-consumer-feign-order80模块

feign用在消费端，feign自带负载均衡配置，所以不用手动配置

pom ：

```xml
	<dependencies>
        <!-- Open Feign，他里面也有ribbon，所以有负载均衡 -->
         <dependency>
            <groupId>org.springframework.cloud</groupId>
            <artifactId>spring-cloud-starter-openfeign</artifactId>
        </dependency>
        <!-- eureka Client服务注册 -->
        <dependency>
            <groupId>org.springframework.cloud</groupId>
            <artifactId>spring-cloud-starter-netflix-eureka-client</artifactId>
        </dependency>
        <dependency><!-- 引入自己定义的api通用包，可以使用Payment支付Entity -->
            <groupId>com.dkf.cloud</groupId>
            <artifactId>cloud-api-commons</artifactId>
            <version>${project.version}</version>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-web</artifactId>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-actuator</artifactId>
        </dependency>
        <dependency>
            <groupId>org.projectlombok</groupId>
            <artifactId>lombok</artifactId>
            <optional>true</optional>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-devtools</artifactId>
            <scope>runtime</scope>
            <optional>true</optional>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-test</artifactId>
            <scope>test</scope>
        </dependency>
    </dependencies>
```

yaml

```yaml
server:
  port: 80

eureka:
  client:
    register-with-eureka: false
    service-url: # 配置服务中心，openFeign去里面找服务
      defaultZone: http://eureka7001.com:7001/eureka/,http://eureka7002.com:7002/eureka/
```

###### `@EnableFeignClients`

主启动类：

```java
@SpringBootApplication
@EnableFeignClients   //关键注解
public class CustomerFeignMain80 {
    public static void main(String[] args) {
        SpringApplication.run(CustomerFeignMain80.class, args);
    }
}
```

###### `@FeignClient`

新建一个service

> 这个service还是 customer 模块的接口，和提供者没有任何关系，不需要包类名一致。它使用起来就相当于是普通的service。
>
> 推测大致原理，对于这个service 接口，读取它某个方法的注解（GET或者POST注解不写报错），知道了请求方式和请求地址，而抽象方法，只是对于我们来讲，调用该方法时，可以进行传参等。

```java
@Component //别忘了添加这个
@FeignClient(value = "CLOUD-PROVIDER-SERVICE")  //服务名称，要和eureka上面的一致才行
public interface PaymentFeignService{
    //这个就是provider 的controller层的方法定义。
    @GetMapping(value = "/payment/get/{id}")
    public CommonResult<Payment> getPaymentById(@PathVariable("id") Long id);

    @GetMapping(value = "/payment/feign/timeout")
    public String paymentFeignTimeout();
}
/*
在提供端有：
	@GetMapping(value = "/payment/get/{id}")
    public CommonResult<Payment> getPaymentById(@PathVariable("id") Long id)
    {
        Payment payment = paymentService.getPaymentById(id);

        if(payment != null)
        {
            return new CommonResult(200,"查询成功,serverPort:  "+serverPort,payment);
        }else{
            return new CommonResult(444,"没有对应记录,查询ID: "+id,null);
        }
    }
*/
```

Controller层：

```java
//使用起来就相当于是普通的service。
@RestController
public class CustomerFeignController {

    @Resource
    private PaymentFeignService paymentFeignService;//动态代理

    @GetMapping("/customer/feign/payment/{id}")
    public CommonResult<Payment> getPaymentById(@PathVariable("id") Long id){
        return paymentFeignService.getPaymentById(id);
    }
}
```

### 超时控制

超时设置，故意设置超时演示出错情况：

- 服务提供方8001故意写暂停程序
- 服务消费方80添加超时方法PaymentFeignService
- 服务消费方80添加超时方法OrderFeignControIIer
- 测试：http://localhost/consumer/payment/feign/timeout
  错误页面

> Openfeign默认超时等待为一秒，在消费者里面配置超时时间

```java
//8001服务提供方 
@GetMapping(value = "/payment/feign/timeout")
public String paymentFeignTimeout(){
    // 业务逻辑处理正确，但是需要耗费3秒钟
    try { 
        TimeUnit.SECONDS.sleep(3); 
    } catch (InterruptedException e) { e.printStackTrace(); }
    return serverPort;
}
```

```java
//消费方80 
@GetMapping(value = "/consumer/payment/feign/timeout")
public String paymentFeignTimeout(){
    // OpenFeign客户端一般默认等待1秒钟
    return paymentFeignService.paymentFeignTimeout();
}
```



```yaml
eureka:
  client:
    register-with-eureka: false
    service-url:
      defaultZone: http://eureka7001.com:7001/eureka/,http://eureka7002.com:7002/eureka/
#设置feign客户端超时时间(OpenFeign默认支持ribbon)
ribbon:
  #指的是建立连接所用的时间，适用于网络状况正常的情况下,两端连接所用的时间
  ReadTimeout: 5000
  #指的是建立连接后从服务器读取到可用资源所用的时间
  ConnectTimeout: 5000
```

### 开启日志打印


Feign共了日志打印功能，我们可以诵过配置来调整日志级别，从而了解Feign中Tttp请求的细节。

说白了就是对Feign接口的调用情况进行监控和输出。

日志级别：

- NONE.默认的，不显示任何日志；
- BASIC，仅记录请求方法、URL、响应状态码及执行时间；
- HEADERS：除了BASIC中定义的信息之外，还有请求和响应的头信息
- FULL:除了HEADERS中定义的信息之外，还有请求和响应的正文及元数据。

首先写一个config配置类：

```java
package com.atguigu.springcloud.config;

import feign.Logger;

@Configuration
public class FeignConfig{
    @Bean
    Logger.Level feignLoggerLevel(){
        return Logger.Level.FULL;
    }
}
```

然后在yml文件中开启日志打印配置：

```yaml
server:
  port: 80

eureka:
  client:
    register-with-eureka: false
    service-url:
      defaultZone: http://eureka7001.com:7001/eureka/,http://eureka7002.com:7002/eureka/
  #设置feign客户端超时时间(OpenFeign默认支持ribbon)
ribbon:
  #指的是建立连接所用的时间，适用于网络状况正常的情况下,两端连接所用的时间
  ReadTimeout: 5000
  #指的是建立连接后从服务器读取到可用资源所用的时间
  ConnectTimeout: 5000

logging:
  level:
    # feign日志以什么级别监控哪个接口
    com.atguigu.springcloud.service.PaymentFeignService: debug
```

# 中级部分

> 主要是服务降级、服务熔断、服务限流的开发思想和框架实现

![1597213783265](https://i0.hdslb.com/bfs/album/8ba6cf17998c1de765f45bd28c8f0ffff21a1023.png)

# ==Hystrix== 断路器

> 官方地址：https://github.com/Netflix/Hystrix/wiki/How-To-Use
>
> "断路器“本身是一种开关装置，当某个服务单元发生故障之后，涌过断路器的故障监控（类似熔断保险丝），向调用方返回一个符合预期的、可处理的备选响应(FallBack)，而不是长时间的等待或者抛出调用方无法处理的异常，这样就保证了服务调用方的线程不会被长时间、不必要地占用，从而避免了故障在分布式系统中的蔓延，乃至雪崩。

## 概述

#### 1）服务雪崩：

多个微服务之间调用的时候，假设微服务A调用微服务B和微服务C，微服务B和微服务C又调用其它的微服务，这就是所谓的“**扇出**"

```
A-->B,C
BC-->D
```

如果扇出的链路上某个微服务的调用响应时间过长或者不可用，对微服务A的调用就会占用越来越多的系统资源，进而引起系统崩溃，即所谓的“`雪崩效应`”。

对于高流量的应用来说，单一的后端依赖可能会导致所有服务器上的所有资源都在几秒钟内饱和。比失败更糟糕的是，这些应用程序还可能导致服务之间的延迟增加，备份队列，线程和其他系统资源紧张，导致整个系统发生更多的级联故障。这些都表示需要对故障和延迟进行隔离和管理，以便单个依赖关系的失败，不能取消整个应用程序或系统。



Hystrix是一个用于处理分布式系统的延迟和容错的开源库，在分布式系统里，许多依赖不可避免的会调用失败，比如超时、异常等，Hystrix能够保证在一个依赖出问题的情况下，不会导致整体服务失败，避免级联故障，以提高分布式系统的弹性。



> Hystrix停止更新，进入维护阶段：https://github.com/Netflix/Hystrix
>
> https://github.com/Netflix/Hystrix/wiki/How-To-Use

#### 2）服务降级：

fallback

服务降级是指 当服务器压力剧增的情况下，根据实际业务情况及流量，对一些服务和页面有策略的不处理或换种简单的方式处理，从而释放服务器资源以保证核心业务正常运作或高效运作。说白了，就是**尽可能的把系统资源让给优先级高的服务**。
  资源有限，而请求是无限的。如果在并发高峰期，不做服务降级处理，一方面肯定会影响整体服务的性能，严重的话可能会导致宕机某些重要的服务不可用。所以，一般在高峰期，为了保证核心功能服务的可用性，都要对某些服务降级处理。比如当双11活动时，把交易无关的服务统统降级，如查看蚂蚁深林，查看历史订单等等。

> 服务器忙碌或者网络拥堵时，不让客户端等待并立刻返回一个友好提示，fallback。
>
> - 对方系统不可用了，你需要给我一个**兜底的方法**，不要耗死。
> - 向调用方返回一个符合预期的、可处理的**备选响应**(FallBack)，而不是长时间的等待或者抛出调用方无法处理的异常，这样就保证了服务调用方的线程不会被长时间、不必要地占用，从而避免了故障在分布式系统中的蔓延，乃至雪崩。
>
> 

降级、熔断：https://blog.csdn.net/www1056481167/article/details/81157171

降低发生的情况：

- 程序运行异常
- 超时
- 服务熔断触发服务降级
- 线程池/信号量打满也会导致服务降级

> 
>
>   服务降级主要用于什么场景呢？当整个微服务架构整体的负载超出了预设的上限阈值或即将到来的流量预计将会超过预设的阈值时，为了保证重要或基本的服务能正常运行，可以**将一些 不重要 或 不紧急 的服务或任务进行服务的 延迟使用 或 暂停使用**。
>   降级的方式可以根据业务来，可以延迟服务，比如延迟给用户增加积分，只是放到一个缓存中，等服务平稳之后再执行 ；或者在粒度范围内关闭服务，比如关闭相关文章的推荐。
>
>
>   实现服务降级需要考虑几个问题
>
> 1）那些服务是核心服务，哪些服务是非核心服务
> 2）那些服务可以支持降级，那些服务不能支持降级，降级策略是什么
> 3）除服务降级之外是否存在更复杂的业务放通场景，策略是什么？
>   
>   自动降级分类
>   1）超时降级：主要配置好超时时间和超时重试次数和机制，并使用异步机制探测回复情况
>   2）失败次数降级：主要是一些不稳定的api，当失败调用次数达到一定阀值自动降级，同样要使用异步机制探测回复情况
>   3）故障降级：比如要调用的远程服务挂掉了（网络故障、DNS故障、http服务返回错误的状态码、rpc服务抛出异常），则可以直接降级。降级后的处理方案有：默认值（比如库存服务挂了，返回默认现货）、兜底数据（比如广告挂了，返回提前准备好的一些静态页面）、缓存（之前暂存的一些缓存数据）
>   4）限流降级：秒杀或者抢购一些限购商品时，此时可能会因为访问量太大而导致系统崩溃，此时会使用限流来进行限制访问量，当达到限流阀值，后续请求会被降级；降级后的处理方案可以是：排队页面（将用户导流到排队页面等一会重试）、无货（直接告知用户没货了）、错误页（如活动太火爆了，稍后重试）。
>
> 

#### 3）服务熔断：

break

- 类比保险丝达到最大服务访问后，**直接拒绝访问**，拉闸限电，然后调用服务降级的方法并返回友好提示
- 就是保险丝：服务的降级->进熔断->恢复调用链路

服务熔断是应对**雪崩效应**的一种微服务链路保护机制。例如在高压电路中，如果**某个地方的电压过高**，熔断器就会熔断，对电路进行保护。同样，在微服务架构中，熔断机制也是起着类似的作用。当调用链路的某个微服务不可用或者响应时间太长时，会进行服务熔断，不再有该节点微服务的调用，快速返回错误的响应信息。当检测到该节点微服务调用响应正常后，恢复调用链路。

  在Spring Cloud框架里，熔断机制通过`Hystrix`实现。Hystrix会监控微服务间调用的状况，当失败的调用到一定阈值，缺省是5秒内20次调用失败，就会启动熔断机制。
  服务熔断解决如下问题： 1. 当所依赖的对象不稳定时，能够起到快速失败的目的；2. 快速失败后，能够根据一定的算法动态试探所依赖对象是否恢复。

> 服务熔断和服务降级的区别
>   触发原因不太一样，服务熔断一般是某个服务（下游服务）故障引起，而服务降级一般是从整体负荷考虑；
>   管理目标的层次不太一样，熔断其实是一个框架级的处理，每个微服务都需要（无层级之分），而降级一般需要对业务有层级之分（比如降级一般是从最外围服务开始）
>   实现方式不太一样，服务降级具有代码侵入性(由控制器完成/或自动降级)，熔断一般称为自我熔断。
>
> 限流：限制并发的请求访问量，超过阈值则拒绝；
> 降级：服务分优先级，牺牲非核心服务（不可用），保证核心服务稳定；从整体负荷考虑；
> 熔断：依赖的下游服务故障触发熔断，避免引发本系统崩溃；系统自动执行和恢复
>
> 

#### 4）服务限流：

flowlimit

秒杀高并发等操作，严禁一窝蜂的过来拥挤，大家排队，一秒钟几个，有序进行

> 可见，上面的技术不论是消费者还是提供者，根据真实环境都是可以加入配置的。

## 断路案例

> 首先构建一个eureka作为服务注册中心的单机版微服务架构 ，这里使用之前eureka Server 7001模块，作为服务中心

新建 提供者 ` cloud-provider-hystrix-payment8001` 模块：

pom 文件：

```xml
	<dependencies>
        <!-- hystrix -->
        <dependency>
            <groupId>org.springframework.cloud</groupId>
            <artifactId>spring-cloud-starter-netflix-hystrix</artifactId>
        </dependency>
        
        <!--eureka-client-->
        <dependency>
            <groupId>org.springframework.cloud</groupId>
            <artifactId>spring-cloud-starter-netflix-eureka-client</artifactId>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-web</artifactId>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-actuator</artifactId>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-devtools</artifactId>
            <scope>runtime</scope>
            <optional>true</optional>
        </dependency>
        <dependency>
            <groupId>org.projectlombok</groupId>
            <artifactId>lombok</artifactId>
            <optional>true</optional>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-test</artifactId>
            <scope>test</scope>
        </dependency>
        <dependency><!-- 引入自己定义的api通用包，可以使用Payment支付Entity -->
            <groupId>com.dkf.cloud</groupId>
            <artifactId>cloud-api-commons</artifactId>
            <version>${project.version}</version>
        </dependency>
    </dependencies>
```

下面主启动类、service、和controller代码都很简单普通。

主启动类：

```java
@SpringBootApplication(exclude = DataSourceAutoConfiguration.class)
@EnableEurekaClient
public class PaymentMain8001 {
    public static void main(String[] args) {
        SpringApplication.run(PaymentMain8001.class,args);
    }
}
```

service层：

```java
@Service
public class PaymentService {

    /* 可以正常访问的方法*/
    public String paymentinfo_Ok(Integer id){
        return "线程池：" + Thread.currentThread().getName() + "--paymentInfo_OK，id:" + id;
    }

    /* 超时访问的方法 */
    public String paymentinfo_Timeout(Integer id){
        int interTime = 3;
        try{
            TimeUnit.SECONDS.sleep(interTime);//模拟超时
        }catch (Exception e){
            e.printStackTrace();
        }
        return "线程池：" + Thread.currentThread().getName() + "--paymentInfo_Timeout，id:" + id + 
            "耗时" + interTime + "秒钟--";
    }
}
```

controller层：

```java
@RestController
@Slf4j
public class PaymentController {

    @Resource
    private PaymentService paymentService;

    @Value("${server.port}")
    private String serverPort;

    @GetMapping("/payment/hystrix/{id}")
    public String paymentInfo_OK(@PathVariable("id")Integer id){
        log.info("paymentInfo_OKKKKOKKK");
        return paymentService.paymentinfo_Ok(id);
    }

    @GetMapping("/payment/hystrix/timeout/{id}")
    public String paymentInfo_Timeout(@PathVariable("id")Integer id){
        log.info("paymentInfo_timeout");
        return paymentService.paymentinfo_Timeout(id);
    }
}
```

### `JMeter`模拟高并发

##### `JMeter`

这里使用一个新东西 JMeter 压力测试器，模拟多个请求

下载压缩包，解压，双击 /bin/ 下的 jmeter.bat 即可启动

<img src="https://i0.hdslb.com/bfs/album/9fb61438cc41ecc19a0b659f8361371f6f02ef16.png" alt="1597471683432" style="zoom: 80%;" /><img src="https://i0.hdslb.com/bfs/album/8a37eabc08a299f6f1b623391b4707926580b10f.png" alt="1597471750363" style="zoom:80%;" />

ctrl + S 保存后，输入请求地址开始压测。

![1597472029967](https://i0.hdslb.com/bfs/album/ca78c74f065ea175ba68449bfec24eef24f46563.png)



**从测试可以看出，当模拟的超长请求被高并发以后，访问普通的小请求速率也会被拉低。**

tomcat的默认工作线程数被打满了，没有多余的线程来分解压力和处理。

上面还是服务8001自己测试，加入此时外部的消费者80页来访问，那消费者只能干等，最终导致消费端80不满意，服务端8001直接被拖死。

测试可见，当启动高并发测试时，消费者访问也会变得很慢，甚至出现超时报错。

演示问题：8001端口自身已经被打满了，80还要访问8001，80也响应慢了。

- 超时导致服务器变慢（转圈）：超时不再等待
- 出错（宕机或程序运行出错）：出错要有兜底

解决思路：

- 对方服务（8001）`超时`了，调用者（80）不能一直卡死等待，必须有服务降级
- 对方服务（8001)  `宕机`了，调用者（80）不能一直卡死等待，必须有服务降级
- 对方服务（8001) OK，调用者（80）自己出故障或有自我要求（自己的等待时间小于服务提供者），自己处理降级

### 服务降级案例



新建消费者 `cloud-customer-feign-hystrix-order80` 模块：以feign为服务调用，eureka为服务中心的模块，

```xml
<!--openfeign-->
<dependency>
    <groupId>org.springframework.cloud</groupId>
    <artifactId>spring-cloud-starter-openfeign</artifactId>
</dependency>
<!--hystrix-->
<dependency>
    <groupId>org.springframework.cloud</groupId>
    <artifactId>spring-cloud-starter-netflix-hystrix</artifactId>
</dependency>
<!--eureka client-->
<dependency>
    <groupId>org.springframework.cloud</groupId>
    <artifactId>spring-cloud-starter-netflix-eureka-client</artifactId>
</dependency>
```

```yaml
server:
  port: 80

eureka:
  client:
    register-with-eureka: false
    service-url:
      defaultZone: http://eureka7001.com:7001/eureka/

feign:
  hystrix:
    enabled: true
```

主启动类

```java
@SpringBootApplication
@EnableFeignClients
@EnableHystrix // 豪猪
public class OrderHystrixMain80{//消费端
    public static void main(String[] args) {
        SpringApplication.run(OrderHystrixMain80.class,args);
    }
}
```

> 一般服务降级放在客户端，即 消费者端 ，但是提供者端一样能使用。
>
> 首先提供者，即8001 先从自身找问题，设置自身调用超时的峰值，峰值内正常运行，超出峰值需要有兜底的方法处理，作服务降级fallback

#### 服务端降级`@HystrixCommand`

首先 对 `8001`的service进行配置（对容易超时的方法进行配置) :

降级配置：`@HystrixCommand`，可以在里面指定超时/出错的回调方法，作为兜底方法



1）提供方 service方法：演示超时

```java
// 服务端service
@HystrixCommand(fallbackMethod = "paymentInfo_TimeOutHandler",//超时后回调方法
                commandProperties = {
                    @HystrixProperty(name="execution.isolation.thread.timeoutInMilliseconds",//时间单位
                                     value="3000")})//超时时间
public String paymentInfo_TimeOut(Integer id){
    try { TimeUnit.MILLISECONDS.sleep(5000); } catch (InterruptedException e) { e.printStackTrace(); }
    return "线程池:  "+Thread.currentThread().getName()+" id:  "+id+"\t"+"O(∩_∩)O哈哈~"+"  耗时(秒): ";
}

// 兜底方法
public String paymentInfo_TimeOutHandler(Integer id){ // 回调函数向调用方返回一个符合预期的、可处理的备选响应
    return "线程池:  "+Thread.currentThread().getName()+"  8001系统繁忙或者运行报错，请稍后再试,id:  "+id+"\t"+"o(╥﹏╥)o";
}
```

提供方 主启动类：

```java
@SpringBootApplication
@EnableEurekaClient
@EnableCircuitBreaker//加上这个
public class PaymentHystrixMain8001{//提供者
    public static void main(String[] args) {
        SpringApplication.run(PaymentHystrixMain8001.class, args);
    }
```

上面演示的是超时，下面演示上面service出错：同样走一样的回调方法

```java
// 服务端service
@HystrixCommand(fallbackMethod = "paymentInfo_TimeOutHandler",//超时后回调方法，出错后也走
                commandProperties = {
                    @HystrixProperty(name="execution.isolation.thread.timeoutInMilliseconds",//时间单位
                                     value="3000")})//超时时间
public String paymentInfo_TimeOut(Integer id){
    int age = 10/0;
    return "线程池:  "+Thread.currentThread().getName()+" id:  "+id+"\t"+"O(∩_∩)O哈哈~"+"  耗时(秒): ";
}

// 兜底方法
public String paymentInfo_TimeOutHandler(Integer id){ // 回调函数向调用方返回一个符合预期的、可处理的备选响应
    return "线程池:  "+Thread.currentThread().getName()+"  8001系统繁忙或者运行报错，请稍后再试,id:  "+id+"\t"+"o(╥﹏╥)o";
}
```

#### 消费端降级

上面的案例是服务端降级，现在我们服务端处理3s，然后返回。但是消费端等1s就等不住了，这时候就需要消费端也有降级方法

`80`的降级。原理是一样的，上面的`@HystrixCommand`降级可以放在服务端，也可以放在消费端。但一般放在客户端。

> 注意：我们自己配置过的热部署方式对java代码的改动明显，但对@HystrixCommand内属性的修改建议重启微服务。

```yaml
server:
  port: 80

eureka:
  client:
    register-with-eureka: false
    service-url:
      defaultZone: http://eureka7001.com:7001/eureka/

feign:
  hystrix:
    enabled: true
```

```java
@SpringBootApplication
@EnableFeignClients
@EnableHystrix // 加上这个  //注意区别我们在提供端添加的注解是@EnableCircuitBreaker
public class OrderHystrixMain80{
    public static void main(String[] args){
        SpringApplication.run(OrderHystrixMain80.class,args);
    }
}
```

然后对 80 进行服务降级：很明显 service 层是接口，所以我们对消费者，在它的 controller 层进行降级。继续使用`@HystrixCommand`注解指定方法超时后的回调方法

controller

```java
@RestController
@Slf4j
public class OrderHystirxController{
    @Resource
    private PaymentHystrixService paymentHystrixService;


    @GetMapping("/consumer/payment/hystrix/timeout/{id}")
    @HystrixCommand(fallbackMethod = "paymentTimeOutFallbackMethod",
                    commandProperties = {
            @HystrixProperty(name="execution.isolation.thread.timeoutInMilliseconds",
                             value="1500") })//只等1.5s
    //@HystrixCommand
    public String paymentInfo_TimeOut(@PathVariable("id") Integer id) {
        // int age = 10/0;
        String result = paymentHystrixService.paymentInfo_TimeOut(id);
        return result;
    }
    public String paymentTimeOutFallbackMethod(@PathVariable("id") Integer id) {
        return "我是消费者80,对方支付系统繁忙请10秒钟后再试 或者 自己运行出错请检查自己,o(╥﹏╥)o";
    }
}
//然后我们把上面的 int age=10/0;打开，出错后也会调用它的兜底方法
```

service

```java
@Component
@FeignClient(value = "CLOUD-PROVIDER-HYSTRIX-PAYMENT" ,fallback = PaymentFallbackService.class)
public interface PaymentHystrixService{
    @GetMapping("/payment/hystrix/ok/{id}")
    public String paymentInfo_OK(@PathVariable("id") Integer id);

    @GetMapping("/payment/hystrix/timeout/{id}")
    public String paymentInfo_TimeOut(@PathVariable("id") Integer id);
}
```

目前问题：

- 每个业务方法对应一个兜底的方法，代码膨胀
- 同样和自定义分开

我们定义一个全局的兜底方法，这样就不用每个方法都得写兜底方法了。

##### 全局兜底`@DefaultProperties`

```java
@RestController
@Slf4j
@DefaultProperties(defaultFallback = "payment_Global_FallbackMethod") //添加了这个
public class OrderHystirxController{
    @Resource
    private PaymentHystrixService paymentHystrixService;

    @GetMapping("/consumer/payment/hystrix/timeout/{id}")
    @HystrixCommand(fallbackMethod = "paymentTimeOutFallbackMethod",
                    commandProperties = {
            @HystrixProperty(name="execution.isolation.thread.timeoutInMilliseconds",
                             value="1500") })//只等1.5s
    public String paymentInfo_TimeOut(@PathVariable("id") Integer id) {
        // int age = 10/0;
        String result = paymentHystrixService.paymentInfo_TimeOut(id);
        return result;
    }
    public String paymentTimeOutFallbackMethod(@PathVariable("id") Integer id) {
        return "我是消费者80,对方支付系统繁忙请10秒钟后再试 或者 自己运行出错请检查自己,o(╥﹏╥)o";
    }

    
    // ------------------
    // 下面是全局fallback方法
    public String payment_Global_FallbackMethod(){
        return "Global异常处理信息，请稍后再试，/(ㄒoㄒ)/~~";
    }
}
// 自己指定了兜底方法的话就走自己的兜底方法，否则走全局的兜底方法
```

##### service降级`@FeignClient`

上面是controller端的降级，如果是调用service呢？

下面解决业务逻辑混在一起的问题（解耦）：我们改在service层进行服务降级

> 服务降级，客户端去调用服务端，碰上服务端宕机或关闭。
>
> 本次案例降级处理是在客户端80实现完成的，与服务端8001没有关系。只需要为Feign客户端定义的接口添加一个服务降级处理的实现类即可实现解耦。

在这种方式一般是在客户端，即消费者端，首先【去掉】上面在controller中添加的  @HystrixCommand 和 @DefaultProperties 两个注解。

```java
@Component 	// service
@FeignClient(value = "CLOUD-PROVIDER-HYSTRIX-PAYMENT" ,
             fallback = PaymentFallbackService.class)//指定的是回调的class类
public interface PaymentHystrixService{
    @GetMapping("/payment/hystrix/ok/{id}")
    public String paymentInfo_OK(@PathVariable("id") Integer id);

    @GetMapping("/payment/hystrix/timeout/{id}")
    public String paymentInfo_TimeOut(@PathVariable("id") Integer id);
}
```

service回调方法：是service接口的实现类。此时就不需要在controller上写fallback方法了

```java
@Component
public class PaymentFallbackService implements PaymentHystrixService{ // 统一为该接口异常处理
    @Override
    public String paymentInfo_OK(Integer id){
        return "-----PaymentFallbackService fall back-paymentInfo_OK ,o(╥﹏╥)o";
    }

    @Override //兜底方法，根据上述配置，程序内发生异常、或者运行超时，都会执行该兜底方法
    public String paymentInfo_TimeOut(Integer id){
        return "-----PaymentFallbackService fall back-paymentInfo_TimeOut ,o(╥﹏╥)o";
    }
}
```

此时的yml文件配置

```yml
server:
  port: 80
spring:
  application:
    name: cloud-customer-feign-hystrix-service
eureka:
  client:
    register-with-eureka: true
    fetch-registry: true
    service-url:
      defaultZone: http://eureka7001.com:7001/eureka/

# 用于服务降级 在注解@FeignClient 中添加 fallback 属性值
feign:
  hystrix:
    enabled: true  # 在feign中开启 hystrix
```

2. 修改service 接口：

```java
@Component											
@FeignClient(value = "CLOUD-PROVIDER-HYSTRIX-PAYMENT", 
             fallback = OrderFallbackService.class)// 这里是重点
public interface OrderService {

    @GetMapping("/payment/hystrix/{id}")
    public String paymentInfo_OK(@PathVariable("id")Integer id);

    @GetMapping("/payment/hystrix/timeout/{id}")
    public String paymentInfo_Timeout(@PathVariable("id")Integer id);

}
```

3. fallback 指向的类：

```java
package com.dkf.springcloud.service;

@Component		//注意这里，它实现了service接口
public class OrderFallbackService implements  OrderService{

    @Override
    public String paymentInfo_OK(Integer id) {
        return "OrderFallbackService --发生异常";
    }

    @Override
    public String paymentInfo_Timeout(Integer id) {
        return "OrderFallbackService --发生异常--paymentInfo_Timeout";
    }
}
```

新问题，这样配置如何设置超时时间？

> 首先要知道 下面两个 yml 配置项：
>
> ```properties
> hystrix.command.default.execution.timeout.enable=true    ## 默认值
> 
> ## 为false则超时控制有ribbon控制，为true则hystrix超时和ribbon超时都是用，但是谁小谁生效，默认为true
> 
> hystrix.command.default.execution.isolation.thread.timeoutInMilliseconds=1000  ## 默认值
> 
> ## 熔断器的超时时长默认1秒，最常修改的参数
> ```
>
> 看懂以后，所以：
>
> 只需要在yml配置里面配置 Ribbon 的 超时时长即可。注意：hystrix 默认自带 ribbon包。
>
> ```yml
> ribbon:
> 	ReadTimeout: xxxx
> 	ConnectTimeout: xxx
> ```

### 服务熔断案例

> 实际上服务熔断 和 服务降级 没有任何关系，就像 java 和 javaScript
>
> 服务熔断，有点自我恢复的味道
>
> 参考：https://blog.csdn.net/www1056481167/article/details/81157171
>

> **服务雪崩**
>
> 多个微服务之间调用的时候，假设微服务A调用微服务B和微服务C，微服务B和微服务C有调用其他的微服务，这就是所谓的”扇出”，如扇出的链路上某个微服务的调用响应式过长或者不可用，对微服务A的调用就会占用越来越多的系统资源，进而引起系统雪崩，所谓的”雪崩效应”
>
> **Hystrix**：
>
> Hystrix是一个用于分布式系统的延迟和容错的开源库。在分布式系统里，许多依赖不可避免的调用失败，比如超时、异常等，Hystrix能够保证在一个依赖出问题的情况下，不会导致整个服务失败，避免级联故障，以提高分布式系统的弹性。
>
> **断路器**：
>
> “断路器”本身是一种开关装置，当某个服务单元发生故障监控(类似熔断保险丝)，向调用方法返回一个符合预期的、可处理的备选响应(FallBack)，而不是长时间的等待或者抛出调用方法无法处理的异常，这样就保证了服务调用方的线程不会被长时间、不必要地占用，从而避免了故障在分布式系统中的蔓延。乃至雪崩。
>
> **服务熔断**：
>
> 熔断机制是应对雪崩效应的一种微服务链路保护机制，当扇出链路的某个微服务不可用或者响应时间太长时，会进行服务的降级，进而熔断该节点微服务的调用，快速返回”错误”的响应信息。
>
> 当检测到该节点微服务响应正常后恢复调用链路，在SpringCloud框架机制通过Hystrix实现，Hystrix会监控微服务见调用的状况，当失败的调用到一个阈值，==默认是5秒内20次调用失败就会启动熔断机制==，熔断机制的注解是`@HystrixCommand`



熔断的状态：

- 熔断打开：请求不再进行调用当前服务，内部设置时钟一般为MTTR（平均故障处理时间），当打开时长达到所设时钟则进入半熔断状态
- 熔断关闭：熔断关闭不会对服务进行熔断
- 熔断半开：**部分请求**根据规则调用当前服务，如果请求成功目符合规则，则认为当前服务恢复正常，关闭熔断。

修改cloud-provider-hystrix-payment8001模块

```java
package com.atguigu.springcloud.service;

import cn.hutool.core.util.IdUtil;
import com.netflix.hystrix.contrib.javanica.annotation.HystrixCommand;
import com.netflix.hystrix.contrib.javanica.annotation.HystrixProperty;

@Service
public class PaymentService{

    //=====服务熔断
    @HystrixCommand(fallbackMethod = "paymentCircuitBreaker_fallback",
                    commandProperties = {
            @HystrixProperty(name = "circuitBreaker.enabled",value = "true"),// 是否开启断路器
            @HystrixProperty(name = "circuitBreaker.requestVolumeThreshold",value = "10"),// 请求次数
            @HystrixProperty(name = "circuitBreaker.sleepWindowInMilliseconds",value = "10000"), // 时间窗口期
            @HystrixProperty(name = "circuitBreaker.errorThresholdPercentage",value = "60"),// 失败率达到多少后跳闸
    }) // 在10s内10次请求有60%失败 // 先看次数，再看百分比
    public String paymentCircuitBreaker(@PathVariable("id") Integer id){
        if(id < 0) {
            throw new RuntimeException("******id 不能负数");
        }
        String serialNumber = IdUtil.simpleUUID();// 等价于UUID.randomUUID().toString(); //pom中有hutool-all

        return Thread.currentThread().getName()+"\t"+"调用成功，流水号: " + serialNumber;
    }
    
    public String paymentCircuitBreaker_fallback(@PathVariable("id") Integer id){//服务降级
        return "id 不能负数，请稍后再试，/(ㄒoㄒ)/~~   id: " +id;
    }

}
```

涉及到断路器的三个重要参数快照时间窗、请求总数阀值、错误百分比阀值

- 1：快照时间窗：断路器确定是否打开需要统计一些请求和错误数据而统计的时间范围就是快照时间窗，默认为最近的10秒。
- 2：请求总数阀值：在快照时间窗内，必须满足请求总数阀值才有资格熔断。默认为20，意味着在10秒内，如果该hystrix命令的调用次数不足20次，即使所有的请求都超时或具他原因失败，断路器都不会打开。
- 3：错误百分比阀值：当请求总数在快照时间窗内超过了阀值，上日发生了30次调用，如果在这30次调用中，有15次发生了超时异常，也就是超过50％的错误百分比，在默认设定50％阀值情况，这时候就会将断路器打开。

> The precise way that the circuit opening and closing occurs is as follows:
>
> 1. Assuming the volume across a circuit meets a certain threshold (`HystrixCommandProperties.circuitBreakerRequestVolumeThreshold()`)…
> 2. And assuming that the error percentage exceeds the threshold error percentage (`HystrixCommandProperties.circuitBreakerErrorThresholdPercentage()`)…
> 3. Then the circuit-breaker transitions from `CLOSED` to `OPEN`.
> 4. While it is open, it short-circuits all requests made against that circuit-breaker.
> 5. After some amount of time (`HystrixCommandProperties.circuitBreakerSleepWindowInMilliseconds()`), the next single request is let through (this is the `HALF-OPEN` state). If the request fails, the circuit-breaker returns to the `OPEN` state for the duration of the sleep window. If the request succeeds, the circuit-breaker transitions to `CLOSED` and the logic in **1.** takes over again.经过一段时间后，如果有1个尝试成功了，就慢慢尝试恢复
>
> 参考：https://github.com/Netflix/Hystrix/wiki/How-it-Works#CircuitBreaker

controller:

```java
//====服务熔断
@GetMapping("/payment/circuit/{id}")
public String paymentCircuitBreaker(@PathVariable("id")Integer id){
    return paymentService.paymentCircuitBreaker(id);
}
```
实验效果为，多次出错调用fallback后，调用正常的也出错调用fallback。过了一会又自己恢复了。
展望：以后用的是Sentinel代替。

关于解耦以后的全局配置说明：

例如上面提到的全局服务降级，并且是feign+hystrix整合，即 service 实现类的方式，如何做全局配置？

> 上面有 做全局配置时，设置超时时间的方式，我们可以从中获得灵感，即在yml文件中 进行熔断配置：
>
> ```yml
> hystrix:
> command:
>  default:
>    circuitBreaker:
>      enabled: true
>      requestVolumeThreshold: 10
>      sleepWindowInMilliseconds: 10000
>      errorThresholdPercentage: 60
> ```

## Hystrix DashBoard



除了隔离依赖服务的调用以外，Hystrix还提供了准实时的调用监控(HystrixDashboard)，Hystrix会持续地记录所有通过Hystrix发起的请求的执行信息，并以统计报表和图形的形式展示给用户，包括每秒执行多少请求多少成功，多少失败等。Netflix通过hystrix-metrics-event-stream实现了对以上指标的监控。SpringCloud也提供了HystrixDashboard的整合，对监控内容转化成可视化界面。

新建模块 cloud-hystrix-dashboard9001 ：

pom 文件：

```xml
	<dependencies>
        <!-- hystrix Dashboard-->
        <dependency>
            <groupId>org.springframework.cloud</groupId>
            <artifactId>spring-cloud-starter-netflix-hystrix-dashboard</artifactId>
        </dependency>
        
        <!-- 常规 jar 包 -->
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-web</artifactId>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-actuator</artifactId>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-devtools</artifactId>
            <scope>runtime</scope>
            <optional>true</optional>
        </dependency>
        <dependency>
            <groupId>org.projectlombok</groupId>
            <artifactId>lombok</artifactId>
            <optional>true</optional>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-test</artifactId>
            <scope>test</scope>
        </dependency>
        <!-- 引入自己定义的api通用包，可以使用Payment支付Entity -->
        <dependency>
            <groupId>com.dkf.cloud</groupId>
            <artifactId>cloud-api-commons</artifactId>
            <version>${project.version}</version>
        </dependency>
    </dependencies>
```

yml文件只需要配置端口号，主启动类加上这样注解：@EnableHystrixDashboard

启动测试：访问  http://ocalhost:9001/hystrix

### 监控实战

> 下面使用上面 9001 Hystrix Dashboard 项目，来监控 8001 项目 Hystrix 的实时情况：

注意8001被监控的时候pom里要actuator依赖，此外还要有

- 主启动类上加`@EnableCircuitBreaker`，用于对豪猪熔断机制的支持
- 

```java
@Bean // 注入豪猪的servlet // 该servlet与服务容错本身无关 // springboot默认路径不是/hustrix.stream，只要在自己的项目里自己配置servlet
public ServletRegistrationBean getServlet(){
    HystrixMetricsStreamServlet streamServlet = new HystrixMetricsStreamServlet();
    ServletRegistrationBean servletRegistrationBean = new ServletRegistrationBean(streamServlet);
    servletRegistrationBean.setLoadOnStartup(1);
    servletRegistrationBean.addUrlMappings("/hystrix.stream");
    servletRegistrationBean.setName("HystrixMetricsStreamServlet");
    return servletRegistrationBean;
}
```

然后就可以取页面里看熔断情况了。输入localhost:8001/hystrix.stream

![1597558486510](https://i0.hdslb.com/bfs/album/bc44e0e580748903ca4f97df1b03be9e024e648c.png)

# 服务网关

<img src="https://i0.hdslb.com/bfs/album/8ba6cf17998c1de765f45bd28c8f0ffff21a1023.png" alt="1597213783265" style="zoom: 50%;" />

## ==Gateway==

> 内容过多，开发可参考 https://docs.spring.io/  官网文档

### 简介

SpringCloud Gateway是SpringCloud的一个全新项目，基于Spring5.O+Springboot 2.0和ProjectReactor等技术开发的网关，它旨在为微服务架构提供一种简单有效的统一的API路由管理方式。

SpringCloudGateway作为SpringCloud生态系统中的网关，目标是替代Zuul,在SpringCloud2.0以上版本中，没有对新版本的zuul2.0以上最新高性能版本进行集成，仍然还是使用的Zuul 1.x非Reactor模式的老版本。而为了提升网关的性能，**SpringCloud Gateway是基于WebFlux框架实现的，而webFlux框架底层则使用了高性能的Reactor模式通信框架Netty**。

springCloudGateway的目标提供统一的路由方式且基于Filter链的方式提供了网关基本的功能，例如：安全，监控/指标，和限流。



一方面因为Zuul 1.0已经进入了维护阶段，而且Gateway是SpringCloud团队研发的是亲儿子产品，值得信赖。而且很多功能比zull用起来都简单便捷。

Gateway是基于异步非阻塞模型上进行开发的，性能方面不需要担心。虽然Netflix早就发布了最新的Zuul2.x，但Spring Cloud貌似没有整合计划。而且Netflix相关组件都宣布进入维护期；不知前景如何？

多方面综合考虑Gateway是很理想的网关选择。

<img src="https://i0.hdslb.com/bfs/album/7251e8519a7ff93421fcc6159834ac2faf24cc8f.png" alt="1597559787064" style="zoom:67%;" />

> SpringCloudGateway与Zuul的区别：
>
> 在SpringCloudFinchley正式版之前，SpringCloud推荐的网关是Netflix提供的Zuul：
>
> - 1、Zuul 1.x是一个基于`阻塞I/O`的APIGateway
>   2、Zuul 1.x基于ServIet2.5使用阻塞架构，它**不支持任何长连接**（如WebSocket)，Zuul的设计模式和Nginx较像，每次I/O操作都是从
>   工作线程中选择一个执行，请求线程阻塞到工作线程完成，但是差别是Nginx用C++实现，Zuul用Java实现，而JVM本身会有第一次加载较慢的情况，使得Zuul的性能相对较差。
> - 3、Zuul 2.x理念更先进想基于Netty非阻塞和支持长连接，但SpringCloud目前还没有整合。Zuul2.x的性能较Zuul1.x有较大提升。在性能方面，根据官方提供的基准测试，SpringCloudGateway的RPS（每秒请求数）是Zuul的1.6倍。
> - 4、SpringCloudGateway建立在SpringFramework5、ProjectReactor和SpringB00t2．之上，使用**非阻塞**API
> - 5、SpringCloudGateway还支持WebSocket，并且与Spring紧密集成拥有更好的开发体验

Zuull.x
springcloud中所集成的zuul版本，采用的是tomcat容器，使用的是传统的servlet IO处理模型。

学过尚硅谷web中期课程都知道一个题目，Servlet的生命周期，servlet由servlet container进行生命周期管理

- container启动时构造servlet对象并调用servlet init()进行初始化，
- container运行时接受请求，并为每个请求分配一个线程（一般从线程池中获取空闲线程）然后调用service()
- container关闭时调用servlet destory()销毁servlet

<img src="https://i0.hdslb.com/bfs/album/4be2da7c4c3e451ab90b8f8f1957cc231a6520bd.png" style="zoom: 33%;" />

上述模式的缺点：

servlete—个简单的网络IO模型，当请求进入servlet container时，servlet container就会为其绑定一个线程在并发不高的场景下这种模型是适用的。但是一旦高并发（比如抽风用jemeter压),线程数量就会上涨，而线程资源代价是昂贵的（上线文切换，内存消耗大）严重影响请求的处理时间。

在一些简单业务场景下，不希望为每个request分配一个线程，只需要1个或几个线程就能应对极大并发的请求，这种业务场景下servlet模型没有优势

所以Zuul 1.x是基于servlet之上的一个阻塞式处理模型，即spring实现了处理所有request请求的一个servlet(DispatcherServlet)并由该servlet阻塞式处理处理。所以springcloudzuul无法摆脱servlet模型的弊端。



传统的Web框架比如说：struts2,springmvc等都是基于Servlet API与Servlet容器基础之上运行的。

但是，在**Servlet3.1之后有了异步非阻塞的支持**。而WebFlux是一个典型**非阻塞异步**的框架，它的核心是基于Reactor的相关API实现的。相对于传统的web框架来说，它可以运行在诸如Netty，Undertow及支持Servlet3.1的容器上。非阻塞式+函数式编程(Spring5必须让你使
用java8)

SpringWebFlux是Spring5.0引入的新的响应式框架区别于SpringMVC,它不需要依赖ServletAPI，它是完全异步非阻塞的，并且基于Reactor来实现响应式流规范。

##### Gateway是什么

Gateway特性：

- 基于SpringFramework5，ProjectReactor和SpringBoot 2.0进行构建；
- 动态路由：能够匹任何请求属性；
- 可以对路由指定Predicate（断言）和Filter（过滤器）·
- 集成Hystrix的断路器功能；
- 集成SpringCloud服务发现功能；
- 易于编写的Predicate（断言）和Filter（过滤器）·
- 请求限流功能；
- 支持路径重写。



GateWay的三大核心概念：

- Route(路由）：路由是构建网关的基本模块，它由ID+目标URI+一系列的断言和过滤器组成，如果断言为true则匹配该路由
- Predicate(断言）：参考的是Java8的java.util.function.predicate。开发人员可以匹配HTTP请求中的所有内容（例如请求头或请求参数），如果请求与断言相匹配则进行路由
  - 匹配条件
- Filter(过滤）：指的是spring框架中GatewayFilter的实例，使用过滤器，可以在请求被路由前或者之后对请求进行修改。
- 例子：通过`断言`虽然进来了，但老师要罚站10min（`过滤器`操作），然后才能正常坐下听课

web请求，通过一些匹配条件，定位到真正的服务节点。并在这个转发过程的前后，进行一些精细化控制。
而filter，就可以理解为一个无所不能的拦截器。有了这两个元素，再加上目标uri，就可以实现一个具体的路由了

<img src="https://i0.hdslb.com/bfs/album/4c0dc00cbe91184fadd4e0d061f051539477f971.png" style="zoom: 50%;" />

![Spring Cloud Gateway Diagram](https://docs.spring.io/spring-cloud-gateway/docs/2.2.5.RELEASE/reference/html/images/spring_cloud_gateway_diagram.png)

客户端向Spring Cloud Gateway发出请求。然后在Gateway HandlerMapping中找到与请求相匹配的路由，将其发送到Gateway Web Handler

Handler再通过指定的过滤器链来将请求发送到我们实际的服务执行业务逻辑，然后返回。

过滤器之间用虚线分开是因为过滤器可能会在发送代理请求之前（Pre）或之后（post）执行业务逻辑。

`Filter`在pre类型的过滤器可以做**参数校验，权限校验，流量监听，日志输出，协议转换**等，

在post类型的过滤器中可以做响应内容、响应头的修改，日志的输出，流量监控等有着非常重要的作用。

### 入门配置

新建模块 cloud-gateway-gateway9527

> 现在实现，通过Gateway (网关) 来访问其它项目，这里选择之前8001项目，要求注册进Eureka Server 。其它没要求。

pom文件：

```xml
<dependencies>
        <!--gateway-->
        <dependency>
            <groupId>org.springframework.cloud</groupId>
            <artifactId>spring-cloud-starter-gateway</artifactId>
        </dependency>
        <!--eureka-client gateWay网关作为一种微服务，也要注册进服务中心。哪个注册中心都可以，如zk-->
        <dependency>
            <groupId>org.springframework.cloud</groupId>
            <artifactId>spring-cloud-starter-netflix-eureka-client</artifactId>
        </dependency>
    <!-- gateway和spring web+actuator不能同时存在，即web相关jar包不能导入 -->
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-devtools</artifactId>
            <scope>runtime</scope>
            <optional>true</optional>
        </dependency>
        <dependency>
            <groupId>org.projectlombok</groupId>
            <artifactId>lombok</artifactId>
            <optional>true</optional>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-test</artifactId>
            <scope>test</scope>
        </dependency>
        <dependency><!-- 引入自己定义的api通用包，可以使用Payment支付Entity -->
            <groupId>com.dkf.cloud</groupId>
            <artifactId>cloud-api-commons</artifactId>
            <version>${project.version}</version>
        </dependency>
    </dependencies>
```

yml文件：

```yml
server:
  port: 9527
spring:
  application:
    name: cloud-gateway
  ## GateWay配置
  cloud:
    gateway: 
      routes: #多个路由
        - id: payment_routh  # 路由ID ， 没有固定的规则但要求唯一，建议配合服务名
          uri: http://localhost:8001  # 匹配后提供服务的路由地址 #uri+predicates  # 要访问这个路径得先经过9527处理
          predicates:
            - Path=/payment/get/**  # 断言，路径相匹配的进行路由

        - id: payment_routh2  # 路由ID ， 没有固定的规则但要求唯一，建议配合服务名
          uri: http://localhost:8001  # 匹配后提供服务的路由地址
          predicates:
            - Path=/payment/lb/**  # 断言，路径相匹配的进行路由

# 注册进 eureka Server # 网关他本身也是一个微服务，也要注册进注册主中心
eureka:
  client:
    service-url:
      defaultZone: http://eureka7001.com:7001/eureka/
    register-with-eureka: true
    fetch-registry: true
```

> 如果IDEA出现了bug，yml配置文件没有变为小叶子，就去structure里设置下模块的spring，选择该yml

主启动类

```java
@SpringBootApplication
@EnableEurekaClient
public class GatewayMain9527 {
    public static void main(String[] args) {
        SpringApplication.run(GatewayMain9527.class,args);
    }
}
```

8001看看controller的访问地址，我们目前不想暴露8001端口，希望在8001外面套一层9527。这样别人就攻击不了8001，有网关挡着

访问测试：

- 1 启动eureka Server，
- 2 启动 8001 项目，
- 3 启动 9527（Gateway项目）

> 可见，当我们访问 http://localhost:9527/payment/get/1 时，即访问网关地址时，会给我们转发到 8001 项目的请求地址，以此作出响应。
>
> 加入网关前：http://localhost:8001/payment/get/1
>
> 加入网关后：http://localhost:9527/payment/get/1

上面是以 yml 文件配置的路由，也有使用config类配置的方式：

```java
@Configuration
public class GateWayConfig{
    @Bean
    public RouteLocator customRouteLocator(RouteLocatorBuilder routeLocatorBuilder){
        RouteLocatorBuilder.Builder routes = routeLocatorBuilder.routes();

        // 分别是id，本地址，转发到的地址
        routes.route("path_route_atguigu",
                r -> r.path("/guonei").uri("http://news.baidu.com/guonei")
                    ).build();//JDK8新特性

        return routes.build();
    }
}
```

```java
@Component
@Slf4j
public class MyLogGateWayFilter implements GlobalFilter,Ordered {

    @Override
    public Mono<Void> filter(ServerWebExchange exchange, GatewayFilterChain chain) {
        log.info("***********come in MyLogGateWayFilter:  "+new Date());

        String uname = exchange.getRequest().getQueryParams().getFirst("uname");

        if(uname == null) {
            log.info("*******用户名为null，非法用户，o(╥﹏╥)o");
            exchange.getResponse().setStatusCode(HttpStatus.NOT_ACCEPTABLE);
            return exchange.getResponse().setComplete();
        }

        return chain.filter(exchange);
    }

    @Override
    public int getOrder() {
        return 0;
    }
}
```



### 动态配置

> 这里所谓的动态配置就是利用服务注册中心，来实现 负载均衡 的调用 多个微服务。
>
> 默认情况下gateway会根据注册中心注册的服务列表，以注册中心上微服务名为路径创建动态路由进行转发，从而实现动态路由的功能
>
> 注意，这是GateWay 的负载均衡

对yml进行配置：让其先通过gateway，再通过gateway去注册中心找提供者

```yml
server:
  port: 9527

spring:
  application:
    name: cloud-gateway
  cloud:
    gateway:
      discovery:
        locator:
          enabled: true # 开启从注册中心动态创建路由的功能，利用微服务名进行路由
      routes:
        - id: payment_routh  # 路由ID ， 没有固定的规则但要求唯一，建议配合服务名
         # uri: http://localhost:8001  # 匹配后提供服务的路由地址 
          uri: lb://CLOUD-PROVIDER-SERVICE # lb 属于GateWay 的关键字，代表是动态uri，即代表使用的是服务注册中心的微服务名，它默认开启使用负载均衡机制 
          predicates:
            - Path=/payment/get/**  # 断言，路径相匹配的进行路由

        - id: payment_routh2  # 路由ID ， 没有固定的规则但要求唯一，建议配合服务名
        #  uri: http://localhost:8001  # 匹配后提供服务的路由地址
          uri: lb://CLOUD-PROVIDER-SERVICE
          predicates:
            - Path=/payment/lb/**  # 断言，路径相匹配的进行路由

# uri: lb://CLOUD-PROVIDER-SERVICE  解释：lb 属于GateWay 的关键字，代表是动态uri，即代表使用的是服务注册中心的微服务名，它默认开启使用负载均衡机制 

eureka:
  instance:
    hostname: cloud-gateway-service
  client: #服务提供者provider注册进eureka服务列表内
    service-url:
      register-with-eureka: true
      fetch-registry: true
      defaultZone: http://eureka7001.com:7001/eureka
```

> 下面可以开启 8002 模块，并将它与8001同微服务名，注册到 Eureka Server 进行测试。

### Gateway:Predicate

> 注意到上面yml配置中，有个predicates 属性值。

- 1 After Route Predicate：只有在指定时间后,才可以路由到指定微服务
- 2 Before Route Predicate
- 3 Between Route Predicate
- 5 Header Route Predicate：只有包含指定请求头的请求,才可以路由
- 6 Host Route Predicate：只有指定主机的才可以访问,比如我们当前的网站的域名是www.aa.com,只有用户是www.aa.com的请求,才进行路由
- 7 Method Route Predicate：只有指定请求才可以路由,比如get请求...
- 8 Path Route Predicate：只有访问指定路径,才进行路由
- 9 Query Route Predicate：必须带有请求参数才可以访问



具体使用：

```yaml
spring:
  application:
    name: cloud-gateway
  cloud:
    gateway:
      discovery:
        locator:
          enabled: true #开启从注册中心动态创建路由的功能，利用微服务名进行路由
      routes:
        - id: payment_routh #payment_route    #路由的ID，没有固定规则但要求唯一，建议配合服务名
          #uri: http://localhost:8001          #匹配后提供服务的路由地址
          uri: lb://cloud-payment-service #匹配后提供服务的路由地址
          predicates:
            - Path=/payment/get/**         # 断言，路径相匹配的进行路由

        - id: payment_routh2 #payment_route    #路由的ID，没有固定规则但要求唯一，建议配合服务名
          #uri: http://localhost:8001          #匹配后提供服务的路由地址
          uri: lb://cloud-payment-service #匹配后提供服务的路由地址
          predicates:
            - Path=/payment/lb/**         # 断言，路径相匹配的进行路由
            #- After=2020-02-21T15:51:37.485+08:00[Asia/Shanghai]
            #- Cookie=username,zzyy
            #- Header=X-Request-Id, \d+  # 请求头要有X-Request-Id属性并且值为整数的正则表达式
```

predicates下面可以有多个属性，表示多个属性与操作为true这个路由才生效

predicates属性下的After属性

```yaml
		predicates:
            - Path=/payment/lb/** 
            - After=2020-02-21T15:51:37.485+08:00[Asia/Shanghai] # 会在这个时间之后这个路由才生效
```

predicates属性下的Cookie属性

```yaml
predicates:
	- Cookie=username,zzyy,ch.p # 需要有这个Cookie值才生效  最后一个是正则表达式
	# curl 地址 --cookie "a=b"
```

predicates属性下的Header属性

```yaml
spring:
  cloud:
    gateway:
      routes:
      - id: header_route
        uri: https://example.org
        predicates:
        - Header=X-Request-Id, \d+
```

```yaml
spring:
  cloud:
    gateway:
      routes:
      - id: host_route
        uri: https://example.org
        predicates:
        - Host=**.somehost.org,**.anotherhost.org  # Host Route Predicate Factory接收一组参数，一组匹配的域名列表，这个模板是一个ant分隔的模板，用.号作为分隔符。它通过参数中的主机地址作为匹配规则
        
        
        # 放爬虫思路，前后端分离的话，只限定前端项目主机访问，这样可以屏蔽大量爬虫。

例如我加上： - Host=localhost:**       ** 代表允许任何端口

就只能是主机来访
```

更多属性可以参考：https://docs.spring.io/spring-cloud-gateway/docs/2.2.5.RELEASE/reference/html/#configuring-route-predicate-factories-and-gateway-filter-factories

需要注意的是value部分写的是正则表达式

高并发工具：jmeter、postman、curl

配置错误页面:

> 注意，springboot默认/static/error/ 下错误代码命名的页面为错误页面，即 404.html
>
> 而且不需要导入额外的包，Gateway 里面都有。

### Gateway:Filter

生命周期：在请求进入路由之前,和处理请求完成,再次到达路由之前

- pre：
- post：

种类：

- GatewayFilter：https://docs.spring.io/spring-cloud-gateway/docs/2.2.5.RELEASE/reference/html/#gatewayfilter-factories
- GlobalFilter：https://docs.spring.io/spring-cloud-gateway/docs/2.2.5.RELEASE/reference/html/#global-filters

```yaml
spring:
  cloud:
    gateway:
      routes:
      - id: add_request_header_route
        uri: https://example.org
        filters:
        - AddRequestHeader=X-Request-red, blue # 添加了个请求头X-Request-red
```

主要是配置全局自定义过滤器，其它的小配置具体看官网吧

2个主要接口：`implements GlobalFilter,Ordered`

场景：

- 全局日志记录
- 统一网关鉴权
- 。。。

自定义全局过滤器配置类：

```java
@Component
@Slf4j
public class MyLogGateWayFilter
    implements GlobalFilter,Ordered {

    @Override
    public Mono<Void> filter(ServerWebExchange exchange,
        GatewayFilterChain chain) {

        log.info("********** come in MyLogGateWayFilter:  "+new Date());

        String uname = exchange.getRequest().getQueryParams().getFirst("uname");

        //合法性检验
        if(uname == null) {
            log.info("*******用户名为null，非法用户，o(╥﹏╥)o，请求不被接受");
            //设置 response 状态码   因为在请求之前过滤的，so就算是返回NOT_FOUND 也不会返回错误页面
            exchange.getResponse().setStatusCode(HttpStatus.NOT_ACCEPTABLE);
            //完成请求调用
            return exchange.getResponse().setComplete();
        }

        return chain.filter(exchange);//过滤链放行
    }

    // 返回值是加载顺序，一般全局的都是第一位加载
    @Override
    public int getOrder() {
        return 0;
    }
}
```



# 服务配置

<img src="https://i0.hdslb.com/bfs/album/8ba6cf17998c1de765f45bd28c8f0ffff21a1023.png" alt="1597213783265" style="zoom: 50%;" />

## ==Config==

> SpringCloud Config 分布式配置中心

### 概述

微服务意味着要将单应用中的业务拆分成一个个子服务，每个服务的粒度相对较小，因此系统中会出现大量的服务。由于每个服务都需要必要的配置信息才能运行，所以一套集中式的、动态的配置管理设施是必不可少的。

springCloud提供了ConfigServer来解决这个问题，我们每一个微服务自己带着一个application.yml，上百个配置文件的管理。比如数据库的信息，我们可以写到一个统一的地方。

- config+bus
- alibaba nacos
- 携程 阿波罗

<img src="https://i0.hdslb.com/bfs/album/b966d78a0557fa1121d4d04a5529100853d59356.png" style="zoom:50%;" />

SpringCloud Config是什么：Spring Cloud Config为微服务架构中的微服务提供集中化的外部配置支持，配置服务器为各个不同微服务应用的所有环境提供了一个**中心化的外部配置**。

怎么玩：
SpringCloud Config分为服务端和客户端两部分。

- 服务端也称为分布式配置中心，它是一个独立的微服务应用，用来连接配置服务器并为客户端提供获取配置信息，加密/解密信息等访问接口
- 客户端则是通过指定的配置中心来管理应用资源，以及与业务相关的配置内容，并在启动的时候从配置中心取和加载配置信息配置服务器默认采用git来存储配置信息，这样就有助于对环境配置进行版本管理，并且可以通过git客户端工具来方便的管理和访问配置内容

能干嘛：

- 集中管理配置文件
- 不同环境不同配置，**动态化的配置更新**，分环境部署比如dev/test/prod/beta/release
- 运行期间动态调整配置，不再需要在每个服务部署的机器上编写配置文件，服务会向配置中心统一拉取配置自己的信息
- 当配置发生变动时，服务不需要重启即可感知到配置的变化并应用新的配置
- 将配置信息以REST接囗的形式暴露



### config服务端配置

这个服务端指的是消费端与github之间的桥接

> 首先在github上新建一个仓库 springcloud-config
>
> git@github.com:名字/项目.git
>
> 然后使用git命令克隆到本地，命令：git clone https://github.com/LZXYF/springcloud-config
>
> 注意上面的操作不是必须的，只要github上有就可以，克隆到本地只是修改文件。
>
> 常用命令：
>
> - git add
> - git commit -m "标记"
> - git push origin master

在git根目录下创建

- 开发环境：config-dev.yml
- 生产环境：config-pro.yml
- 测试环境：config-test.tml

注意格式

新建 `cloud-config-center3344` 模块：

pom文件：

```xml
	<dependencies>
        <!-- config Server -->
        <dependency>
            <groupId>org.springframework.cloud</groupId>
            <artifactId>spring-cloud-config-server</artifactId>
        </dependency>

        <!--eureka-client config Server也要注册进服务中心-->
        <dependency>
            <groupId>org.springframework.cloud</groupId>
            <artifactId>spring-cloud-starter-netflix-eureka-client</artifactId>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-devtools</artifactId>
            <scope>runtime</scope>
            <optional>true</optional>
        </dependency>
        <dependency>
            <groupId>org.projectlombok</groupId>
            <artifactId>lombok</artifactId>
            <optional>true</optional>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-test</artifactId>
            <scope>test</scope>
        </dependency>
        <dependency><!-- 引入自己定义的api通用包，可以使用Payment支付Entity -->
            <groupId>com.dkf.cloud</groupId>
            <artifactId>cloud-api-commons</artifactId>
            <version>${project.version}</version>
        </dependency>
    </dependencies>
```

yml 配置：

```yml
server:
  port: 3344

spring:
  application:
    name:  cloud-config-center #注册进Eureka服务器的微服务名
  cloud:
    config:
      server:
        git: # 此处使用的是老师是配置中心
          uri: git@github.com:zzyybs/springcloud-config.git #GitHub上面的git仓库名字
          ####搜索目录
          search-paths:
            - springcloud-config
      ####读取分支
      label: master
      
#服务注册到eureka地址
eureka:
  client:
    service-url:
      defaultZone: http://localhost:7001/eureka
```

主启动类：

```java
@SpringBootApplication
@EnableConfigServer   //关键注解
public class ConfigCenterMain3344 { // 注意先去把Eureka启动起来
    public static void main(String[] args) {
        SpringApplication.run(ConfigCenterMain3344.class,args);
    }
}
```

添加模拟映射：【C:\Windows\System32\drivers\etc\hosts】文件中添加：  `127.0.0.1 config-3344.com`

启动微服务3344，访问 http://config-3344.com:3344/master/config-dev.yml 文件（注意，要提前在git上弄一个这文件）

老师的配置中心地址：https://github.com/zzyybs/springcloud-config

文件命名和访问的规则：

![1597646186970](https://i0.hdslb.com/bfs/album/ef2c5a899e003e8004a05e591fa2202018b5e168.png)

不加分支名默认是master:

![1597646308915](https://i0.hdslb.com/bfs/album/64d9399ebdc84caf4f4e372c372971cea7f9e4f3.png)

最后一个出的是json串

- label：分支branch
- name：服务名
- profiles：环境dev/test/prod

### config客户端配置

> 这里的客户端指的是，使用 Config Server 统一配置文件的项目。既有之前说的消费者，又有提供者

新建 `cloud-config-client-3355` 模块：

pom文件：

```xml
	<dependencies>
        <!-- config Client 和 服务端的依赖不一样 -->
        <dependency>
            <groupId>org.springframework.cloud</groupId>
            <artifactId>spring-cloud-starter-config</artifactId>
        </dependency>

        <!--eureka-client config Server也要注册进服务中心-->
        <dependency>
            <groupId>org.springframework.cloud</groupId>
            <artifactId>spring-cloud-starter-netflix-eureka-client</artifactId>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-web</artifactId>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-actuator</artifactId>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-devtools</artifactId>
            <scope>runtime</scope>
            <optional>true</optional>
        </dependency>
        <dependency>
            <groupId>org.projectlombok</groupId>
            <artifactId>lombok</artifactId>
            <optional>true</optional>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-test</artifactId>
            <scope>test</scope>
        </dependency>
        <dependency><!-- 引入自己定义的api通用包，可以使用Payment支付Entity -->
            <groupId>com.dkf.cloud</groupId>
            <artifactId>cloud-api-commons</artifactId>
            <version>${project.version}</version>
        </dependency>
    </dependencies>
```

要将Client模块下的application.yml文件改为bootstrap.yml，删掉application.yml这是很关键的

因为bootstrap.yml是比application.yml先加载的。bootstrap.yml优先级高于application.yml

appllication.yml是用户级的资源配置项

bootstrap.ym1是系统级的，优先级更加高

SpringCloud会创建一个"Bootstrap Context"作为Spring应用的ApplicationContext的`父上下文`。初始化的时候，Bootstrap Context负责从外部源加载配置属性并解析配置。这两个上下文共享一个从外部获取的Environment

Bootstrap 属性有高优先级，默认情况下，它们不会被本地配置覆盖。BootstrapContext和ApplicationContext、有着不同的约定，所以新增了一个bootstrap.yml文件，保证BootstrapContext和ApplicationContext配置的分离。

bootstrap.yml文件：

```yml
server:
  port: 3355

spring:
  application:
    name: config-client
  cloud:
    #Config客户端配置
    config:
      label: master #分支名称
      name: config #配置文件名称，文件也可以是client-config-dev.yml这种格式的，这里就写 client-config 
      profile: dev #读取后缀名称   上述3个综合：master分支上config-dev.yml的配置文件被读取http://config-3344.com:3344/master/config-dev.yml
      uri: http://localhost:3344 #配置中心地址
      # 综合上面四个 即读取配置文件地址为： http://config-3344.com:3344/master/config-dev.yml

#服务注册到eureka地址
eureka:
  client:
    service-url:
      defaultZone: http://localhost:7001/eureka
```

主启动类，极其普通：

```java
@SpringBootApplication
@EnableEurekaClient
public class ConfigClientMain3355 
    public static void main(String[] args) {
        SpringApplication.run(ConfigClientMain3355.class, args);
    }
}
```

controller层，测试读取配置信息

```java
package com.dkf.springcloud.controller;

@RestController
public class ConfigClientController {

    @Value("${config.info}") // 消费 //相当于配置了config后，就把config服务端里的变量引入进来了
    private String configInfo;

    @GetMapping("/configInfo")
    public String getConfigInfo(){
        return configInfo;
    }
}
```

启动测试完成！如果报错，注意github上的 yml 格式有没有写错！

启动Config配置中心3344微服务并自测，启动3355作为client访问 localhost:3355/configInfo

修改config-dev.yml配置文件并提交到github中，比如加个变量age或者版本号version。更改消费者端的配置看看其他环境能不能用

### 配置动态刷新

问题：

- Linux运维修改GitHub上的配置文件内容做调整：比如修改config-dev.yml提交
- 刷新3344，发现ConfigServer服务端配置中心立刻响应，得到最新值了
- 刷新3355，发现ConfigClient客户端没有任何响应，拿到的还是旧值
- 客户端3355没有变化除非自己重启或者重新加载，才能拿到最新值
- 难到每次运维修改配置文件，**客户端**都需要重启？？噩梦

> 就是github上面配置更新了，config Server 项目上是动态更新的，但是，client端的项目中的配置，目前还是之前的，它不能动态更新，必须重启才行。

动态刷新问题解决：

1. client端一定要有actuator依赖：

    ```xml
    <dependency>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-actuator</artifactId>
    </dependency>
    ```

2. client 端增加 yml 配置如下，即在 bootstrap.yml 文件中：

```yml
server:
  port: 3355

spring:
  application:
    name: config-client
  cloud:
    #Config客户端配置
    config:
      label: master #分支名称
      name: config #配置文件名称
      profile: dev #读取后缀名称   上述3个综合：master分支上config-dev.yml的配置文件被读取http://config-3344.com:3344/master/config-dev.yml
      uri: http://localhost:3344 #配置中心地址k

#服务注册到eureka地址
eureka:
  client:
    service-url:
      defaultZone: http://localhost:7001/eureka

# 暴露监控端点
management:
  endpoints:
    web:
      exposure:
        include: "*"
```

3. 在controller 上添加注解`@RefreshScope`：

   ```java
   @RestController
   @RefreshScope //这个注解
   public class ConfigClientController {
       @Value("${config.info}")
       private String configInfo;
   
       @GetMapping("/configInfo")
       public String getConfigInfo() {
           return configInfo;
       }
   }
   ```

> 到此为止，配置已经完成，但是测试客户端 localhost:3355/configInfo 仍然不能动态刷新，还是旧值（也就是说环境变量里的还是旧值），需要下一步。

4. 向 client 端发送一个 POST 请求

> 因为此时,还需要==外部==发送post请求通知3355
>
> 如 curl -X POST "http://localhost:3355/actuator/refresh"
>
> 两个必须：1.必须是 POST 请求，2.请求地址：http://localhost:3355/actuator/refresh
>
> 此时在刷新3355,发现可以获取到最新的配置文件了,这就实现了动态获取配置文件,因为3355并没有重启

成功获得到最新值

但是又有一个问题，就是要向每个微服务客户端发送一次POST请求，当微服务数量庞大，又是一个新的问题。

能否广播，一次通知，处处生效？（还要求不要全广播，差异化管理，定点清除，20台只有18台更新）

就有下面的消息总线！

<img src="https://i0.hdslb.com/bfs/album/8ba6cf17998c1de765f45bd28c8f0ffff21a1023.png" alt="1597213783265" style="zoom: 50%;" />

# 消息总线

消息总线的由来看上面一小节

## ==Bus==

spring cloud Bus配置spring cloud Config使用可以实现配置的动态刷新

spring cloud bus是用来将分布式系统的节点与轻量级消息系统链接起来的框架，它整合了java的事件处理机制和消息中间件的功能。

spring cloud bus目前支持RabbitMQ和Kafka（因为是主题订阅）

什么是总线

在微服务架构的系统中，通常会使用轻量级的消息代理来构建一个共用的消息主题，并让系统中所有微服务实例都连接上来。由于该主题中产生的消息会被所有实例**监听和消费**，所以称它为消息总线。在总线上的各个实例，都可以方便地广播一些需要让貝他连接在该主题上的实例都知道的消息。

基本原理：

`ConfigClient`实例都监听MQ中同一个topic主题(默认是`springCloud Bus`)。当一个服务刷新数据的时候，它会把这个信息放入到Topic中，这样其它监听同一topic的服务就能得到通知，然后去更新自身的配置

下图原盈利是就是给其中一台发送我们的刷新POST，他刷新完后给Bus发消息，然后Bus通过消息中间件发送给BC进行更新

<img src="https://i0.hdslb.com/bfs/album/f434848f7f35e02439843e66293fbdbf0752d35a.png" style="zoom:50%;" />

Bus能管理和传输分布式系统间的消息，就像一个分布式执行器，可用于广播状态更改、事件推送等，也可以当做微服务间的通信通道

<img src="https://i0.hdslb.com/bfs/album/9347357d6670b4561d85feddcf472d2235de8ed7.png" style="zoom: 25%;" />

### ==RabbitMQ==

> 在windows 上安装RabbitMQ

1. 安装RabbitMQ的依赖环境 Erlang  下载地址： http://erlang.org/download/otp_win64_21.3.exe 
2. 安装RabbitMQ   下载地址： http://dl.bintray.com/rabbitmq/all/rabbitmq-server/3.7.14/rabbitmq-server-3.7.14.exe 
3. 进入 rabbitMQ安装目录的sbin目录下，打开cmd窗口，执行 【`rabbitmq-plugins enable rabbitmq_management`】
4. 访问【http://localhost:15672/】，输入密码和账号：默认为guest

### 广播式刷新配置

- 必须先具有良好的RabbitMQ环境
- 演示广播效果，增加复杂度，再以3355为模板再制作一个3366
- 设计思想
  - 1）利用消息总线触发一个客户端/bus/refresh，而刷新所有客户端的配置
  - 2）利用消息总线触发一个服务端ConfigServer的/bus/refres端点，从而刷新所有客户端的配置
  - <img src="https://i0.hdslb.com/bfs/album/688c39eac9b309804981e82bf3397514afc3c7ad.png" style="zoom: 50%;" /><img src="https://i0.hdslb.com/bfs/album/d0e22b8d08186053865ab9bb3bc855d43f44f9c3.png" style="zoom:50%;" />
  - 图二的架构显然更加适合，图一不适合的原因如下
    - 打破了微服务的职责单一性，因为微服务本身是业务模块，它本不应该承担配置刷新的职责
    - 破坏了微服务各节点的对等性。
    - 有一定的局限性“例如，微服务在迁移时，它的网络地址常常会发生变化，此时如果想要做到自动刷新，那就会增加更多的
- 给cloud-config-center-3344配置中心服务端添加消息总线支持
- 给cloud-config-client-3355客户端添加消息总线支持
- 给cloud-config-client-3366客户端添加消息总线支持（以3355为模板）
- 测试
- 一次修改，广播通知，出处生效
- ==但还是得发一个POST请求，只不过只给config发而已==



首先给 config Server 和 config client 都添加如下依赖：

```xml
<!-- 添加rabbitMQ的消息总线支持包 -->
<dependency>
    <groupId>org.springframework.cloud</groupId>
    <artifactId>spring-cloud-starter-bus-amqp</artifactId>
</dependency>
```

config Server 的yml文件增加如下rabbitmq配置：

```yml
# rabbitMq的相关配置
rabbitmq:
  host: localhost
  port: 5672  # 这里没错，虽然rabbitMQ网页是 15672
  username: guest
  password: guest
# rabbitmq 的相关配置2 暴露bus刷新配置的端点
management:
  endpoints:
    web:
      exposure:
        include: 'bus-refresh'
```

config Client 的yml文件修改成如下配置：（注意对齐方式，和config Server不一样）

```yml
spring:
  cloud:
    # config 客户端配置
    config:
      label: master         # 分支名称
      name: client-config       # 配置文件名称
      profile: test      # 使用配置环境
      uri: http://config-3344.com:3344  # config Server 地址
```

3355客户端的bootstrap.yml

```yml
server:
  port: 3355 # client

spring:
  application:
    name: config-client
  cloud:
    #Config客户端配置
    config:
      label: master #分支名称
      name: config #配置文件名称
      profile: dev #读取后缀名称   上述3个综合：master分支上config-dev.yml的配置文件被读取http://config-3344.com:3344/master/config-dev.yml
      uri: http://localhost:3344 #配置中心地址k

#rabbitmq相关配置 15672是Web管理界面的端口；5672是MQ访问的端口
  rabbitmq:
    host: localhost
    port: 5672
    username: guest
    password: guest

#服务注册到eureka地址
eureka:
  client:
    service-url:
      defaultZone: http://localhost:7001/eureka

# 暴露监控端点
management:
  endpoints:
    web:
      exposure:
        include: "*"
```

3344注册中心的application.yml

```yml
server:
  port: 3344

spring:
  application:
    name:  cloud-config-center #注册进Eureka服务器的微服务名
  cloud:
    config:
      server:
        git:
          uri: git@github.com:zzyybs/springcloud-config.git #GitHub上面的git仓库名字
        ####搜索目录
          search-paths:
            - springcloud-config
      ####读取分支
      label: master
#rabbitmq相关配置
rabbitmq:
    host: localhost
    port: 5672
    username: guest
    password: guest

#服务注册到eureka地址
eureka:
  client:
    service-url:
      defaultZone: http://localhost:7001/eureka

##rabbitmq相关配置,暴露bus刷新配置的端点
management:
  endpoints: #暴露bus刷新配置的端点
    web:
      exposure:
        include: 'bus-refresh'
```



可在github上修改yml文件进行测试，修改完文件，向 config server 发送 请求：

给3344发就能全局同步了【curl -X POST "http://localhost:3344/actuator/bus-refresh"】

> 注意，之前是向config client 一个个发送请求，但是这次是向 config Server 发送请求，而所有的config client 的配置也都全部更新。

### 定点通知

新的需求：指定具体某一个实例（的参数）生效而不是全部，一些是最新值，一些是旧值

- 公式：`http://localhost:配置中心的端口号/actuator/bus-refresh/{destination}`
- 例子：curl -X POST "http//localhost:3344/actuator/bus-refresh/config-client:3355
- 即微服务名称+端囗号
- /bus/refresh请求不再发送到具体的服务实例上，而是发给configserver并通过destination参数类指定需要更新配置的服务或实例

- 我们这里以刷新运行在3355端口上的config-client为例
  - 只通知3355
  - 不通知3366
    

# 消息驱动

## ==Stream==

需求：消息中间件很多，希望向上抽象一个接口，我们不关心底层用的是什么消息中间件

屏蔽底层消息中间件的差异，降低切换成本，统一消息的编程模型

就像 JDBC 形成一种规范，统一不同数据库的接口

![1597725567239](https://i0.hdslb.com/bfs/album/da8e110a71af2f040bc3f4cb052b8701a7707ee2.png)



什么是SpringCloud Stream

官方定义SpringCloud Stream是一个构建消息驱动微服务的框架。https://spring.io/projects/spring-cloud-stream#overview

应用程序通过inputs或者outputs来与SpringCloud Stream中binder对象（绑定器）交互

涌过我们配置来binding（绑定）而SpringCloud Stream的binder对象负责与消息中间件交互。

所以，我们只需要搞清楚如何与springCloudstrearn交互就可以方便使用消息驱动的方式。



通过使用Spring Integration来连接消息代理中间件以实现消息事件驱动。

SpringCloud Stream为一些供应商的消息中间件产品提供了个性化的自动化配置实现，引用了发布-订阅、消费组、分区的三个核心概念



目前仅支持RabbitMQ、Kafka。

> 流程：
>
> pub生产者发送消息，BROKER接收消息放到队列中，订阅者接收到消息
>
> 选修必须走特定的通道：下嘻嘻通道MessageChannel
>
> 消息通道里的消息如何消费呢？谁负责收发处理：消息通道MessageChannel的子接口SubscribableChannel，由MessageHandler消息处理器所订阅
>
> 比如java里用的是RabbitMQ，大数据里用的是kafka，来回切换麻烦，链各个消息中间件的架构上不同
>
> 像RabbitMQ有exchange，kafka有Topic和Partitions分区
>
> 这些中间件的差异导致我们实际项目开发给我们造成了一定的困难，我们如果用了两个消息队列的其中一种，后面的业务需求，我们想往另外一种消息队列进行迁移，这时候无疑就是一个灾难性的，一大堆东西都要重新推倒重新做，因为它跟我们的系统耦合了，这时候SpringCloud Stream给我们提供了一种解耦合的方式。

Stream的消息通信方式遵循了**发布-订阅模式**

通过定义绑定器Binder作为中间层，实现了应用程序与消息中间件细节之间的隔离。INPUT对应于生产者，OUTPUT对应于消费者

![](https://i0.hdslb.com/bfs/album/a78fe3645bad6c7bc3576ecf40e7c0371ef6daaf.png)



Stream标准流程套路：

- binder：很方便的连接中间件，屏蔽差异
- Channel：通道，是队列Queue的一种抽象，在消息通讯系统中就是实现存储和转发的媒介，通过channel对队列进行配置
- Source（生产）和sink（消费）：简单地可理解为参照对象是spring cloud stream自身，从stream发布消息就是输出，接收消息就是输入

![1597730581088](https://i0.hdslb.com/bfs/album/85fb0b2bff7fd8d953f494bd3003bc30b28435db.png)

常用注解：

| 组成            | 说明                                                         |
| --------------- | ------------------------------------------------------------ |
| Middleware      | 中间件，目前只支FRabbitMQ和Kafka                             |
| Binder          | Binder是应用与消息中间件之间的封装，目前实行了KafKa和RabbitMQ的Binder,通过<br/>Binder可以很方便的连接中间件，可以动态的改变消息类型（对应kafka的topic，<br/>RabbitMQ的exchange)，这些都可以通过配置文件来实现 |
| @Input          | 注解标识输入通道，通过该输入通接收到的消息息进入应用程序     |
| @Output         | 注解标识输出通道，发布的消息将通过该通道离开应用程序         |
| @StreamListener | 监听队列，用于消费者的队列的消息接收                         |
| @EnableBinding  | 指信道channel和exchange绑定在一起                            |



### 消息生产者

要新建3个子模块

- cloud-stream-rabbitmq-provide8801：作为生产者进行发消息模块
- cloud-stream-rabbitmq-consumer8802：作为消息接收模块
- cloud-stream-rabbitmq-consumer8802：作为消息接收模块



新建模块 cloud-stream-rabbitmq-provider8801 

8801 pom依赖：

```xml
<!-- stream-rabbit -->
<dependency>
    <groupId>org.springframework.cloud</groupId>
    <artifactId>spring-cloud-starter-stream-rabbit</artifactId>
</dependency>

<!--eureka-client 目前，这个不是必须的-->
<dependency>
    <groupId>org.springframework.cloud</groupId>
    <artifactId>spring-cloud-starter-netflix-eureka-client</artifactId>
</dependency>
<dependency>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-starter-web</artifactId>
</dependency>
<dependency>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-starter-actuator</artifactId>
</dependency>
<dependency>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-devtools</artifactId>
    <scope>runtime</scope>
    <optional>true</optional>
</dependency>
<dependency>
    <groupId>org.projectlombok</groupId>
    <artifactId>lombok</artifactId>
    <optional>true</optional>
</dependency>
<dependency>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-starter-test</artifactId>
    <scope>test</scope>
</dependency>
<dependency><!-- 引入自己定义的api通用包，可以使用Payment支付Entity -->
    <groupId>com.dkf.cloud</groupId>
    <artifactId>cloud-api-commons</artifactId>
    <version>${project.version}</version>
</dependency>
```

yml 配置：

```yml
server:
  port: 8801
  
spring:
  application:
    name: cloud-stream-provider
  cloud:
    stream:
      binders: # 在此配置要绑定的rabbitMQ的服务信息
        defaultRabbit: # 表示定义的名称，用于和binding整合
          type: rabbit  # 消息组件类型
          environment:  # 设置rabbitmq的相关环境配置
            spring:
              rabbitmq:
                host: localhost
                port: 5672
                username: guest
                password: guest
      bindings:  # 服务的整合处理
        output:   # 表示是生产者，向rabbitMQ发送消息
          destination: studyExchange  # 表示要使用的Exchange名称
          content-type: application/json  # 设置消息类型，本次是json，文本是 "text/plain"
          binder: defaultRabbit  # 设置要绑定的消息服务的具体配置
          
eureka:
  client:
    service-url:
      defaultZone: http://eureka7001.com:7001/eureka/
  instance:
    lease-renewal-interval-in-seconds: 2 # 设置心跳时间，默认是30秒
    lease-expiration-duration-in-seconds: 5 # 最大心跳间隔不能超过5秒,默认90秒
    instance-id: send-8801.com # 在信息列表显示主机名称
    prefer-ip-address: true # 访问路径变为ip地址
```

主启动类没什么特殊的注解。

业务类：（此业务类不是以前的service，而实负责推送消息的服务类）

- 发送消息的接口类
- 发送消息接口类的实现类
- controller

<img src="https://i0.hdslb.com/bfs/album/85fb0b2bff7fd8d953f494bd3003bc30b28435db.png" alt="1597730581088" style="zoom:50%;" />

```java
package com.atguigu.springcloud.service;

public interface IMessageProvider {
    public String send();
}
```



```java
package com.dkf.springcloud.service;

import org.springframework.cloud.stream.annotation.EnableBinding;
import org.springframework.cloud.stream.messaging.Source;
import org.springframework.messaging.MessageChannel;
import org.springframework.messaging.support.MessageBuilder;

import javax.annotation.Resource;
import java.util.UUID;

@EnableBinding(Source.class)  // 定义消息的推送管道 output//不是和controller打交道的service,而是发送消息的推送服务类
public class IMessageProviderImpl implements IMessageProvider {
									     //上面是自定义的接口
    @Resource
    private MessageChannel output;//消息发送管道

    @Override
    public String send() {
        String serial = UUID.randomUUID().toString();
        output.send(MessageBuilder.withPayload(serial).build());// 绑定器
        System.out.println("******serial: " + serial);
        return null;
    }
}
```

controller:

```java
@RestController
public class SendMessageController {

    @Resource // 自己的类
    private IMessageProvider messageProvider;

    @GetMapping("/sendMessage")
    public String sendMessage(){
        return messageProvider.send(); // 自己定义的方法，但是里面调用了MessageChannel.send()方法
    }
}
```

> 启动Eureka Server 7001，再启动8801，进行测试，看是否rabbitMQ中有我们发送的消息。
>

### 消息消费者

新建模块 cloud-stream-rabbitmq-consumer8802

pom依赖和生产者一样。

yml配置: 在 stream的配置上，和生产者只有一处不同的地方，output 改成 input

```yml
server:
  port: 8802
spring:
  application:
    name: cloud-stream-provider
  cloud:
    stream:
      binders: # 在次配置要绑定的rabbitMQ的服务信息
        defaultRabbit: # 表示定义的名称，用于和binding整合
          type: rabbit  # 消息组件类型
          environment:  # 设置rabbitmq的相关环境配置
            spring:
              rabbitmq:
                host: localhost
                port: 5672
                username: guest
                password: guest
      bindings:  # 服务的整合处理
        input:   # 表示是消费者，这里是唯一和生产者不同的地方，向rabbitMQ发送消息
          destination: studyExchange  # 表示要使用的Exchange名称
          content-type: application/json  # 设置消息类型，本次是json，文本是 "text/plain"
          binder: defaultRabbit  # 设置要绑定的消息服务的具体配置
eureka:
  client:
    service-url:
      defaultZone: http://eureka7001.com:7001/eureka/
  instance:
    lease-renewal-interval-in-seconds: 2 # 设置心跳时间，默认是30秒
    lease-expiration-duration-in-seconds: 5 # 最大心跳间隔不能超过5秒,默认90秒
    instance-id: receive-8802.com # 在信息列表显示主机名称
    prefer-ip-address: true # 访问路径变为ip地址
```

接收消息的业务类：

```java
import org.springframework.cloud.stream.annotation.EnableBinding;
import org.springframework.cloud.stream.annotation.StreamListener;
import org.springframework.cloud.stream.messaging.Sink;
import org.springframework.messaging.Message;

@Component
@EnableBinding(Sink.class)
public class ConsumerController {

    @Value("${server.port}")
    private String serverPort;

    @StreamListener(Sink.INPUT) // 消费者
    public void input(Message<String> message){
        System.out.println("消费者1号，serverport: " + serverPort + "，接受到的消息：" + message.getPayload());
    }
}
```

### 配置分组消费

新建 cloud-stream-rabbitmq-consumer8802 模块：

> 8803 就是 8802 clone出来的。

当运行时，会有两个问题。

第一个问题，两个消费者都接收到了消息，这属于重复消费。例如，消费者进行订单创建，这样就创建了两份订单，会造成系统错误。

注意在stream中处同一个group中的多个消费者是竞争关系，就能保证消息只会被其中一个应用消费一次。

不同组是可以全面消费（重复消费）的

同一组内会发生竞争关系，只有其中一个可以消费。

> Stream默认不同的微服务是不同的组
>
> 



![1597731630685](https://i0.hdslb.com/bfs/album/b3971aee799a1cc48178ccf1baac7ed035fe2b26.png)

对于重复消费这种问题，导致的原因是默认每个微服务是不同的group，组流水号不一样，所以被认为是不同组，两个都可以消费。

解决的办法就是自定义配置分组：

消费者 yml 文件配置：

```yml
	# 8802 的消费者
	bindings:
        input:   
          destination: studyExchange  
          content-type: application/json  
          binder: defaultRabbit  
          group: dkfA  # 自定义分组配置
          
    # 8803 的消费者
	bindings:
        input:   
          destination: studyExchange  
          content-type: application/json  
          binder: defaultRabbit  
          group: dkfB  # 自定义分组配置
```

![1597732035990](https://i0.hdslb.com/bfs/album/2a3baed83ac8bd4984a72a3aa21a66b559ce4637.png)

当两个消费者配置的 group 都为 dkfA 时，就属于同一组，就不会被重复消费。（两个消费者消费同一队列）

![1597732238270](https://i0.hdslb.com/bfs/album/7965001dc159667a1aea59a1de829eb555ee3513.png)

### 消息持久化

> 加上group配置，就已经实现了消息的持久化。

# Sleuth

> 分布式请求链路跟踪，超大型系统。需要在微服务模块极其多的情况下，比如80调用8001的，8001调用8002的，这样就形成了一个链路，如果链路中某环节出现了故障，我们可以使用Sleuth进行链路跟踪，从而找到出现故障的环节。



在微服务框架中，一个由客户端发起的请求在后端系统中会经过多个不同的的服务节点调用来协同产生最后的请求结果，每一个前段请求都会形成一条杂的分布式服务调用链路，链路中的任何一环出现高延时或错误都会引起整个请求最后的失败。

> sleuth 负责跟踪，而zipkin负责展示。
>
> zipkin 下载地址： http://dl.bintray.com/openzipkin/maven/io/zipkin/java/zipkin-server/2.12.9/zipkin-server-2.12.9-exec.jar
>
> 使用 【java -jar】 命令运行下载的jar包，访问地址：【 http://localhost:9411/zipkin/ 】

## 案例

> 使用之前的 提供者8001 和 消费者80

分别给他们引入依赖：

```xml
<!-- 引入sleuth + zipkin -->
<dependency>
    <groupId>org.springframework.cloud</groupId>
    <artifactId>spring-cloud-starter-zipkin</artifactId>
</dependency>
```

yml增加配置：

```yml
spring:
  zipkin:
    base-url: http://localhost:9411  # zipkin 地址
  sleuth:
    sampler:
      # 采样率值 介于0-1之间 ，1表示全部采集
      probability: 1
```

# 高级部分

# SpringCloud Alibaba

> alibaba 的 github上有中文文档



spring netflix进入维护模式。

什么是维护模式：spring cloud团队将不会再向模块添加新功能，我们将修复block级别的bug以及安全问题，我们也会考虑并审查社区的小型pull request。我们打算继续支持这些模块，知道Greenwich版本被普遍采用至少一年

SpringCloud Netflix将不再开发新的组件

以下spring cloud netflix模块和响应的starter将进入维护模式：

- spring-cloud-netflix-archaius
- spring-cloud-netflix-hystrix-contract
- spring-cloud-netflix-hystrix-dashboard
- spring-cloud-netflix-hystrix-stream
- spring-cloud-netflix-hystrix
- spring-cloud-netflix-ribbon：代替Loadbalancer
- spring-cloud-netflix-turbine-stream
- spring-cloud-netflix-turbine
- spring-cloud-netflix-zuul

这不包括Eureka或并发限制模块。



我们都知道SpringCloud版本迭代是比较快的，因而出现了很多重大ISSUE都还来不及Flix就又推另一个RELEASE了。进入维护模式意思就是目前以及以后一段时间SpingCloud Netflix提供的报务和功能就这么多了，不在开发新的组件和功能了。以后将以雏护和Merge分支Full Request为主

新组件功能将以具他替代平代替的方式实现



历史：

- alibaba出了dubbo，停更
- spring结合netflix整出spring cloud。又停更了
- alibaba又出马了得出springcloud alibaba

spring cloud alibaba带来了什么？

2018.10.31，spring cloud Alibaba正式入驻了Spring Cloud官方孵化器，并在Maven中央库发布了第一个版本

主要功能：

- **服务限流降级**：默认支持 WebServlet、WebFlux, OpenFeign、RestTemplate、Spring Cloud Gateway, Zuul, Dubbo 和 RocketMQ 限流降级功能的接入，可以在运行时通过控制台实时修改限流降级规则，还支持查看限流降级 Metrics 监控。
- **服务注册与发现**：适配 Spring Cloud 服务注册与发现标准，默认集成了 Ribbon 的支持。
- **分布式配置管理**：支持分布式系统中的外部化配置，配置更改时自动刷新。
- **消息驱动能力**：基于 Spring Cloud Stream 为微服务应用构建消息驱动能力。
- **分布式事务**：使用 @GlobalTransactional 注解， 高效并且对业务零侵入地解决分布式事务问题。。
- **阿里云对象存储**：阿里云提供的海量、安全、低成本、高可靠的云存储服务。支持在任何应用、任何时间、任何地点存储和访问任意类型的数据。
- **分布式任务调度**：提供秒级、精准、高可靠、高可用的定时（基于 Cron 表达式）任务调度服务。同时提供分布式的任务执行模型，如网格任务。网格任务支持海量子任务均匀分配到所有 Worker（schedulerx-client）上执行。
- **阿里云短信服务**：覆盖全球的短信服务，友好、高效、智能的互联化通讯能力，帮助企业迅速搭建客户触达通道。

只需引入依赖：

```xml
<dependencyManagement>
    <dependencies>
        <dependency>
            <groupId>com.alibaba.cloud</groupId>
            <artifactId>spring-cloud-alibaba-dependencies</artifactId>
            <version>2.2.3.RELEASE</version>
            <type>pom</type>
            <scope>import</scope>
        </dependency>
    </dependencies>
</dependencyManagement>
```

组件：

- **[Sentinel](https://github.com/alibaba/Sentinel)**：把流量作为切入点，从流量控制、熔断降级、系统负载保护等多个维度保护服务的稳定性。
- **[Nacos](https://github.com/alibaba/Nacos)**：一个更易于构建云原生应用的动态服务发现、配置管理和服务管理平台。
- **[RocketMQ](https://rocketmq.apache.org/)**：一款开源的分布式消息系统，基于高可用分布式集群技术，提供低延时的、高可靠的消息发布与订阅服务。
- **[Dubbo](https://github.com/apache/dubbo)**：Apache Dubbo™ 是一款高性能 Java RPC 框架。
- **[Seata](https://github.com/seata/seata)**：阿里巴巴开源产品，一个易于使用的高性能微服务分布式事务解决方案。
- **[Alibaba Cloud OSS](https://www.aliyun.com/product/oss)**: 阿里云对象存储服务（Object Storage Service，简称 OSS），是阿里云提供的海量、安全、低成本、高可靠的云存储服务。您可以在任何应用、任何时间、任何地点存储和访问任意类型的数据。
- **[Alibaba Cloud SchedulerX](https://help.aliyun.com/document_detail/43136.html)**: 阿里中间件团队开发的一款分布式任务调度产品，提供秒级、精准、高可靠、高可用的定时（基于 Cron 表达式）任务调度服务。
- **[Alibaba Cloud SMS](https://www.aliyun.com/product/sms)**: 覆盖全球的短信服务，友好、高效、智能的互联化通讯能力，帮助企业迅速搭建客户触达通道。

https://github.com/alibaba/spring-cloud-alibaba/blob/master/README-zh.md

<img src="https://i0.hdslb.com/bfs/album/8ba6cf17998c1de765f45bd28c8f0ffff21a1023.png" alt="1597213783265" style="zoom: 50%;" />

##  ==Nacos==

nacos(NAming COnfiguration Service)：服务注册和配置中心

> Nacos = Eureka注册中心 + Config配置中心 + Bus消息总线
>
> 替代Eureka做服务注册中心
>
> 替代Config做服务配置中心

> github地址：  https://github.com/alibaba/Nacos 
>
> Nacos 地址：  https://nacos.io/zh-cn/ 

| 服务注册与服务框架 | CAP模型        | 控制台管理 | 社区活跃度      |
| ------------------ | -------------- | ---------- | --------------- |
| Eureka             | AP高可用       | 支持       | 低(2.x版本闭源) |
| Zookeeper          | CP一致         | 支持       | 中              |
| Consul             | CP             | 支持       | 高              |
| Nacos              | AP（可以切换） | 支持       | 高              |



nacos可以切换 AP 和 CP ,可使用如下命令切换成CP模式：

```
curl -X PUT '$NACOS_SERVER:8848/nacos/v1/ns/operator/switches?entry=serverMode&value=CP'
```

下载 ：

> 下载地址：  https://github.com/alibaba/nacos/releases/tag/1.1.4 
>
> 直接下载网址： https://github.com/alibaba/nacos/releases/download/1.1.4/nacos-server-1.1.4.zip
>
> 下载压缩包以后解压，进入bin目录，打开dos窗口，执行startup命令启动它。
>
> 端口号8848
>
> 可访问 ： 【 http://localhost:8848/nacos/index.html】地址，默认账号密码都是nacos

### nacos服务中心

https://nacos.io/zh-cn/docs/feature-list.html

https://spring-cloud-alibaba-group.github.io/github-pages/hoxton/en-us/index.html#_spring_cloud_alibaba_nacos_discovery

#### nacos提供者

新建模块 `cloudalibaba-provider-payment9001`

父pom中

```xml
<dependencyManagement>
    <dependencies>
        <dependency>
            <groupId>com.alibaba.cloud</groupId>
            <artifactId>spring-cloud-alibaba-dependencies</artifactId>
            <version>2.1.1.BUILD-SNAPSHOT</version>
            <type>pom</type>
            <scope>import</scope>
        </dependency>
    </dependencies>
</dependencyManagement>
```

子pom依赖：

```xml
	<dependencies>

        <!-- springcloud alibaba nacos 依赖 -->
        <dependency>
            <groupId>com.alibaba.cloud</groupId>
            <artifactId>spring-cloud-starter-alibaba-nacos-discovery</artifactId>
        </dependency>
        
        <!-- springboot整合Web组件 -->
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-web</artifactId>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-actuator</artifactId>
        </dependency>

        <!-- 日常通用jar包 -->
        <dependency>
            <groupId>org.mybatis.spring.boot</groupId>
            <artifactId>mybatis-spring-boot-starter</artifactId>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-devtools</artifactId>
            <scope>runtime</scope>
            <optional>true</optional>
        </dependency>
        <dependency>
            <groupId>org.projectlombok</groupId>
            <artifactId>lombok</artifactId>
            <optional>true</optional>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-test</artifactId>
            <scope>test</scope>
        </dependency>
        <dependency><!-- 引入自己定义的api通用包，可以使用Payment支付Entity -->
            <groupId>com.dkf.cloud</groupId>
            <artifactId>cloud-api-commons</artifactId>
            <version>${project.version}</version>
        </dependency>
    </dependencies>
```

yml 配置：

```yml
server:
  port: 9001
spring:
  application:
    name: nacos-provider
  cloud:
    nacos:
      discovery:
        server-addr: localhost:8848
management:
  endpoints:
    web:
      exposure:
        include: '*'
```

```java
@EnableDiscoveryClient // 注册
@SpringBootApplication
public class PaymentMain9001 {
    public static void main(String[] args) {
            SpringApplication.run(PaymentMain9001.class, args);
    }
}
```

```java
@RestController
public class PaymentController {
    @Value("${server.port}")
    private String serverPort;

    @GetMapping(value = "/payment/nacos/{id}")
    public String getPayment(@PathVariable("id") Integer id) {
        return "nacos registry, serverPort: "+ serverPort+"\t id"+id;
    }
}
```



<font color=red>Nacos 自带负载均衡机制</font>，下面创建第二个提供者9003。也可以`-Dserver.port=9011`

新建 cloudalibaba-provider-payment9003 提供者模块，clone 9001 就可以

> 如果要用restTemplate做负载均衡，@Bean注入restTemplate的同时在上面加@LoadBalance
>
> 因为Naocs要使用Ribbon进行负载均衡,那么就需要使用RestTemplate

#### nacos消费者

新建消费者 模块： cloudalibaba-customer-nacos-order83

```xml
<!--SpringCloud ailibaba nacos -->
<dependency>
    <groupId>com.alibaba.cloud</groupId>
    <artifactId>spring-cloud-starter-alibaba-nacos-discovery</artifactId>
</dependency>
```

```yaml
server:
  port: 83

spring:
  application:
    name: nacos-order-consumer
  cloud:
    nacos:
      discovery:
        server-addr: localhost:8848


#消费者将要去访问的微服务名称(注册成功进nacos的微服务提供者)
service-url:
  nacos-user-service: http://nacos-payment-provider
```

```java
@Configuration
public class ApplicationContextConfig { // nacos底层也是ribbon，注入RestTemplate
    @Bean
    @LoadBalanced // 负载均衡
    public RestTemplate getRestTemplate() {
        return new RestTemplate();
    }
}
```

controller :

```java
@RestController
@Slf4j
public class OrderNacosController {
    @Resource
    private RestTemplate restTemplate;

    @Value("${service-url.nacos-user-service}")
    private String serverURL; //在yml里面写的提供者服务路径，  值为：http://nacos-provider

    @GetMapping(value = "/consumer/payment/nacos/{id}")
    public String paymentInfo(@PathVariable("id") Long id) {
        return restTemplate.getForObject(serverURL+"/payment/nacos/"+id,String.class);
    }

}
```

#### 各种服务中心对比

| 服务注册与服务框架 | CAP模型 | 控制台管理 | 社区活跃度      |
| ------------------ | ------- | ---------- | --------------- |
| Eureka             | AP      | 支持       | 低(2.x版本闭源) |
| Zookeeper          | CP      | 支持       | 中              |
| Consul             | CP      | 支持       | 高              |
| Nacos              | AP/CP   | 支持       | 高              |

| 组件名    | 语言 | CAP  | 服务健康检查 | 对外暴露接口 | SpringCloud集合 |
| --------- | ---- | ---- | ------------ | ------------ | --------------- |
| Eureka    | java | AP   | 可配支持     | HTTP         | 已集成          |
| Consul    | Go   | CP   | 支持         | HTTP/DNS     | 已集成          |
| Zookeeper | java | CP   | 支持         | 客户端       | 已集成          |

NACOS支持CP和AP切换

C要求一致性，A要求可用性。

何时选择使用何种模式？

一般来说，

如果不需要存储服级的信息且服务实例是通过nacos-client注册，并能够保持心跳上报，那么就可以选择AP模式。当前主流的服务如Spring cloud和Dubbo服务，都适用于AP模式，AP模式为了服务的可用性而减弱了一致性，因此AP模式下只支持注册临时实例。

如果需要在服务级别编辑或者存储配置信息，那么CP是必须，K8S服务和DNS服务则适用于CP模式。

CP模式下则支持注册持久化实例，此时则是以Raft协议为集群运行模式，该模式下汪册实例之前须先注册服务，如果服务不存在，则会返回错误

切换命令：`curl -X PUT '$NACOS_SERVER:8848/nacos/v1/ns/operator/switches?entry=serverMode&value=CP'` 

![](https://i0.hdslb.com/bfs/album/0293faa9eee82f1de0063173efb8cdd63703a307.png)

![](https://i0.hdslb.com/bfs/album/6e098b50a86df17d80dd98d2d43083cfbb198866.png)

### nacos配置中心

##### 配置中心对比

| 对比项目      | Spring Cloud Config    | Apollo                   | Nacos                    |
| ------------- | ---------------------- | ------------------------ | ------------------------ |
| 配置实时推送  | 支持(Spring Cloud Bus) | 支持(HTTP长轮询1s内)     | 支持(HTTP长轮询1s内)     |
| 版本管理      | 支持(Git)              | 支持                     | 支持                     |
| 配置回滚      | 支持(Git)              | 支持                     | 支持                     |
| 灰度发布      | 支持                   | 支持                     | 不支持                   |
| 权限管理      | 支持(依赖Git)          | 支持                     | 不支持                   |
| 多集群        | 支持                   | 支持                     | 支持                     |
| 多环境        | 支持                   | 支持                     | 支持                     |
| 监听查询      | 支持                   | 支持                     | 支持                     |
| 多语言        | 只支持Java             | 主流语言，提供了Open API | 主流语言，提供了Open API |
| 配置格式校验  | 不支持                 | 支持                     | 支持                     |
| 单机读(QPS)   | 7(限流所致)            | 9000                     | 15000                    |
| 单击写(QPS)   | 5(限流所致)            | 1100                     | 1800                     |
| 3节点读 (QPS) | 21(限流所致)           | 27000                    | 45000                    |
| 3节点写 (QPS) | 5(限流所致)            | 3300                     | 5600                     |

从配置中心角度来看，性能方面Nacos的读写性能最高，Apollo次之，Spring Cloud Config依赖Git场景不适合开放的大规模自动化运维API。功能方面Apollo最为完善，nacos具有Apollo大部分配置管理功能，而Spring Cloud Config不带运维管理界面，需要自行开发。Nacos的一大优势是整合了注册中心、配置中心功能，部署和操作相比Apollo都要直观简单，因此它简化了架构复杂度，并减轻运维及部署工作。  

> nacos 还可以作为服务配置中心，下面是案例，创建一个模块，从nacos上读取配置信息。
>
> nacos 作为配置中心，不需要像springcloud config 一样做一个Server端模块。

新建模块 `cloudalibaba-config-nacos-client3377`

pom依赖：

```xml
 <dependencies>
     
        <!-- 以 nacos 做服务配置中心的依赖 -->
        <dependency>
            <groupId>com.alibaba.cloud</groupId>
            <artifactId>spring-cloud-starter-alibaba-nacos-config</artifactId>
        </dependency>
     
        <!-- springcloud alibaba nacos 依赖 -->
        <dependency>
            <groupId>com.alibaba.cloud</groupId>
            <artifactId>spring-cloud-starter-alibaba-nacos-discovery</artifactId>
        </dependency>
        <!-- springboot整合Web组件 -->
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-web</artifactId>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-actuator</artifactId>
        </dependency>

        <!-- 日常通用jar包 -->
        <dependency>
            <groupId>org.mybatis.spring.boot</groupId>
            <artifactId>mybatis-spring-boot-starter</artifactId>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-devtools</artifactId>
            <scope>runtime</scope>
            <optional>true</optional>
        </dependency>
        <dependency>
            <groupId>org.projectlombok</groupId>
            <artifactId>lombok</artifactId>
            <optional>true</optional>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-test</artifactId>
            <scope>test</scope>
        </dependency>
        <dependency><!-- 引入自己定义的api通用包，可以使用Payment支付Entity -->
            <groupId>com.dkf.cloud</groupId>
            <artifactId>cloud-api-commons</artifactId>
            <version>${project.version}</version>
        </dependency>
    </dependencies>
```

主启动类也是极其普通：

```java
@EnableDiscoveryClient // 消费
@SpringBootApplication
public class NacosConfigClientMain3377{
    public static void main(String[] args) {
        SpringApplication.run(NacosConfigClientMain3377.class, args);
    }
}
```

bootstrap.yml 配置：

```yml
# nacos配置
server:
  port: 3377

spring:
  application:
    name: nacos-config-client
  cloud:
    nacos:
      discovery:
        server-addr: localhost:8848 #Nacos服务注册中心地址
      config:
        server-addr: localhost:8848 #Nacos作为配置中心地址
        file-extension: yaml #指定yaml格式的配置
        group: DEV_GROUP
        namespace: 7d8f0f5a-6a53-4785-9686-dd460158e5d4


# ${spring.application.name}-${spring.profile.active}.${spring.cloud.nacos.config.file-extension}
# nacos-config-client-dev.yaml

# nacos-config-client-test.yaml   ----> config.info
```

application.yml

```yml
spring:
  profiles:
    active: dev # 表示开发环境
    #active: test # 表示测试环境
    #active: info
```

controller 层进行读取配置测试：

```java
@RestController
@RefreshScope  //支持Nacos的动态刷新
public class ConfigClientController {

    @Value("${config.info}") // 从nacos中取
    private String configInfo;

    @GetMapping("configclient/getconfiginfo")
    public String getConfigInfo(){
        return configInfo;
    }
}
```

nacos同springcloud-config一样，在项目初始化时，要先从配置中心进行配置拉取，拉取配置之后，才能保证项目的正常启动。

springboot的配置文件的加载是存在优先熟悉怒的，bootstrap优先级高于application。（bootstrap中放共性，application中放个性）

nacos中的dataid的组成格式及与springboot配置文件中的匹配规则：

##### nacos配置文件命名

在nacos中，消费端要的文件怎么和nacos中的文件匹配呢？

在 Nacos Spring Cloud 中，`dataId` 的完整格式如下：（就是说在nacos端我们怎么命名文件的）

```plain
${prefix}-${spring.profiles.active}.${file-extension}
```

- `prefix` 默认为 **服务名**`spring.application.name` 的值，也可以通过配置项 `spring.cloud.nacos.config.prefix`来配置。
- `spring.profiles.active` 即为当前环境对应的 profile，详情可以参考 [Spring Boot文档](https://docs.spring.io/spring-boot/docs/current/reference/html/boot-features-profiles.html#boot-features-profiles)。 **注意：当 `spring.profiles.active` 为空时，对应的连接符 `-` 也将不存在，dataId 的拼接格式变成 `${prefix}.${file-extension}`**
- `file-exetension` 为配置内容的数据格式，可以通过配置项 `spring.cloud.nacos.config.file-extension` 来配置。目前只支持 `properties` 和 `yaml` 类型。（注意nacos里必须使用yaml）

> 从上面可以看到重要的一点，配置文件的名称第二项，spring.profiles.active 是依据当前环境的profile属性值的，也就是这个值如果是 dev，即开发环境，它就会读取 dev 的配置信息，如果是test，测试环境，它就会读取test的配置信息，就是从 spring.profile.active 值获取当前应该读取哪个环境下的配置信息。

所以要配置spring.profiles.active，新建application.yml文件，添加如下配置：

```yaml
spring:
  profiles:
    active: dev # 表示开发环境
```

综合以上说明，和下面的截图，Nacos 的dataid（类似文件名）应为： nacos-config-client-dev.yaml  (必须是yaml)

<img src="https://fermhan.oss-cn-qingdao.aliyuncs.com/img/20201014233039.png" style="zoom:80%;" />

注意nacos里不要写成yml，要写成yaml：

![1597758227367](https://i0.hdslb.com/bfs/album/ad53ad9251aad3b2e97548c86cc51419f9ffa800.png)

当修改配置值，会发现 3377 上也已经修改，Nacos自带自动刷新功能！

nacos的优势在哪：

- 问题1：实际开发者，通常一个系统会准备dev/test/prod环境。如何保证环境启动时服务能正确读取nacos上相应环境的配置文件
  - 用namespace区分环境
- 问题2：一个大型分布式微服务系统有很多微服务子项目，每个微服务项目又都会有相应的开发环境、测试环境、预发环境、正式环境。那怎么对微服务配置进行管理呢？
  - 用group把不同的微服务划分到同一个分组里面去

<img src="https://i0.hdslb.com/bfs/album/d6a1c96467e57540e2d936c7fda7af0bc79cfe4a.png" alt="1597808385154" style="zoom:50%;" />

默认：Namespace=,Cluster=DEFAULT

- namespace环境
- group
- 

Service就是微服务，一个service可以包含多个cluster集群，nacos默认cluster是DEFAULT，Cluster是对指定微服务的一个虚拟划分。

比方说为了容灾，将service微服务分别部署在了杭州机房和广州机房，这是就可以给杭州机房的service微服务起一个集群名称HZ

给广州的service微服务起一个集群名称GZ，还可以尽量让同一个机房的微服务互相调用，以提升虚拟。

最后instance就是微服务的实例

##### dgn方案

`dataid`方案（就是nacos的文件名）：

- 指定spring.profile.active和配置文件的dataID来使不太环境下读取不同的配置
- 配置空间+配置分组+新建dev和test两个dataid：就是创建-后不同的两个文件名`nacos-config-client-dev.yaml`、`nacos-config-client-test.yaml`
- 通过IDEA里的spring.profile.active属性就能进行多环境下配置文件的读取

`Group`方案（默认DEFAULT_GROUP）：

- 在nacos创建配置文件时，给文件指定分组。
- 在IDEA中该group内容
- 实现的功能：当修改开发环境时，只会从同一group中进行切换。

`namespace`方案（默认public）：

- 这个是不允许删除的，可以创建一个新的命名空间，会自动给创建的命名空间一个流水号。
- 在nacos新建命名空间，自动出现7d8f0f5a-6a53-4785-9686-dd460158e5d4
- 在IDEA的yml中指定命名空间namespace: 7d8f0f5a-6a53-4785-9686-dd460158e5d4



最后，dataid、group、namespace 三者关系如下：（不同的dataid，是相互独立的，不同的group是相互隔离的，不同的namespace也是相互独立的）



> 上面只是小打小闹，下面才是真正的高级操作。

搭建集群必须持久化，不然多台机器上的nacos的配置信息不同，造成系统错乱。它不同于单个springcloud config，没有集群一说，而且数据保存在github上，也不同于eureka，配置集群就完事了，没有需要保存的配置信息。

### nacos集群/持久化

nacos挂了怎么办？

https://nacos.io/zh-cn/docs/cluster-mode-quick-start.html

> 一台linux虚拟机：nginx服务器（虚拟ip），3个nacos服务，一个mysql数据库。
>
> nginx的安装参考之前学，使用 ContOs7 至少需要安装gcc库，不然无法编译安装【yum install gcc】
>
> nacos下载linux版本的 tar.gz 包：https://github.com/alibaba/nacos/releases/download/1.1.4/nacos-server-1.1.4.tar.gz
>
> mysql root用户密码为 Dkf!!2020

<img src="https://cdn.nlark.com/yuque/0/2019/jpeg/338441/1561258986171-4ddec33c-a632-4ec3-bfff-7ef4ffc33fb9.jpeg" alt="deployDnsVipMode.jpg" style="zoom:50%;" />



VIP是虚拟IP，即nginx

nginx也该是集群

<img src="https://i0.hdslb.com/bfs/album/bda5729d154d1c6e573c87e461512bdbf35c96e4.png" alt="1597808678658" style="zoom:67%;" />

Nacos支持三种部署模式 https://nacos.io/zh-cn/docs/deployment.html

- 单机模式 - 用于测试和单机试用。
- 集群模式 - 用于生产环境，确保高可用。
- 多集群模式 - 用于多数据中心场景。

单机模式支持mysql：在0.7版本之前，在单机模式时nacos使用**嵌入式数据库**（derby，他的pom里有这个依赖）实现数据的存储，不方便观察数据存储的基本情况。0.7版本增加了支持mysql数据源能力，具体的操作步骤：

- 1.安装数据库，版本要求：5.6.5+

- 2.初始化mysql数据库，数据库初始化文件：nacos/conf/nacos-mysql.sql。创建个database数据库nacos_devtest

- 3.修改IDEA中nacos/conf/application.properties文件(切换数据库)，增加支持mysql数据源配置（目前只支持mysql），添加mysql数据源的url、用户名和密码。

- ```properties
  # 切换数据库
  spring.datasource.platform=mysql
  
  db.num=1
  db.url.0=jdbc:mysql://11.162.196.16:3306/nacos_devtest?characterEncoding=utf8&connectTimeout=1000&socketTimeout=3000&autoReconnect=true
  db.user=root
  db.password=123456
  ```

- 再以单机模式启动nacos(重启)，nacos所有写嵌入式数据库的数据都写到了mysql


单击的数据库都是独立的，我们得让他们共用一个数据库





Nacos集群配置

环境准备：

1. 64 bit OS Linux/Unix/Mac，推荐使用Linux系统。
2. 64 bit JDK 1.8+；[下载](http://www.oracle.com/technetwork/java/javase/downloads/jdk8-downloads-2133151.html).[配置](https://docs.oracle.com/cd/E19182-01/820-7851/inst_cli_jdk_javahome_t/)。
3. Maven 3.2.x+；[下载](https://maven.apache.org/download.cgi).[配置](https://maven.apache.org/settings.html)。
4. 3个或3个以上Nacos节点才能构成集群。

开始配置集群：

1. 首先对 nacos 进行持久化操作，操作如上面一致。

2. 修改 nacos/conf 下的cluster.conf文件，添加如下内容:

    ```sh
    # it is ip
    # 告诉这3个集群结点是一组的 # 不能写127.0.0.1，必须是linux hostname -i能够识别的ip
    192.168.1.2:3333
    192.168.1.2:4444
    192.168.1.2:5555
    ```

3. 修改nacos/conf/application.properties文件，添加设置我们的数据库信息

4. 模拟三台nacos服务，编辑nacos的startup.sh脚本，使他能够支持不同的端口启动多次。
   集群启动，我们希望可以类似其他软件的shell命令，传递不同的端口号启动不同的nacos实例。

   `vim startup.sh`

   ![1597812799242](https://i0.hdslb.com/bfs/album/c3cad088b12de298a83acbfc55041f877a9f71c4.png)

   ![1597813020494](https://i0.hdslb.com/bfs/album/8f8000fb341f560e98c40e24e44969db56c1a324.png)

   `nohup $JAVA -Dserver.port=${PORT} ${JAVA_POT} nacoas.nacos >> ${BASE_DIR}/logs/start.out 2>&1 & `

5. 依次执行命令启动3个nacos集群：
    `./startup.sh -p 3333` 表示启动端口号为3333的nacos服务器实例
    `./startup.sh -p 4444` 
    `./startup.sh -p 5555` 
    `ps -ef | grep nacos | grep -v grep | wc -l  `

6. 修改nginx配置，把他作为负载均衡：

    ```sh
    vim ./nginx/conf/nginx.conf
    ```

    ![1597813917440](https://i0.hdslb.com/bfs/album/e367e687524000360f5d16c978f4a38150d5aa18.png)

7. 启动nginx：`./nginx -c ../conf/nginx.conf`

8. 通过nginx访问：192.168.1.2:1111/nacos/#/login

9. 使用 9002 模块注册进Nacos集群，并获取它上面配置文件的信息application.yml中的`server-addr: 192.168.1.2:1111`，进行测试。



## ==Sentinel==

> sentinel在 springcloud Alibaba 中的作用是实现`熔断`和`限流`。类似于Hystrix豪猪

> 下载地址dashboard： https://github.com/alibaba/Sentinel/releases/download/1.7.1/sentinel-dashboard-1.7.1.jar
>
> 下载jar包以后，使用【java -jar】命令启动即可。
>
> 它使用 8080 端口，用户名和密码都为 ： sentinel

随着微服务的流行，服务和服务之间的稳定性变得越来越重要。Sentinel 以流量为切入点，从流量控制、熔断降级、系统负载保护等多个维度保护服务的稳定性。

Sentinel 具有以下特征:

- **丰富的应用场景**：Sentinel 承接了阿里巴巴近 10 年的双十一大促流量的核心场景，例如秒杀（即突发流量控制在系统容量可以承受的范围）、消息削峰填谷、集群流量控制、实时熔断下游不可用应用等。
- **完备的实时监控**：Sentinel 同时提供实时的监控功能。您可以在控制台中看到接入应用的单台机器秒级数据，甚至 500 台以下规模的集群的汇总运行情况。
- **广泛的开源生态**：Sentinel 提供开箱即用的与其它开源框架/库的整合模块，例如与 Spring Cloud、Dubbo、gRPC 的整合。您只需要引入相应的依赖并进行简单的配置即可快速地接入 Sentinel。
- **完善的 SPI 扩展点**：Sentinel 提供简单易用、完善的 SPI 扩展接口。您可以通过实现扩展接口来快速地定制逻辑。例如定制规则管理、适配动态数据源等。

![Sentinel-features-overview](https://user-images.githubusercontent.com/9434884/50505538-2c484880-0aaf-11e9-9ffc-cbaaef20be2b.png)

Sentinel 分为两个部分:

- 核心库（Java 客户端）不依赖任何框架/库，能够运行于所有 Java 运行时环境，同时对 Dubbo / Spring Cloud 等框架也有较好的支持。
- 控制台（Dashboard）基于 Spring Boot 开发，打包后可以直接运行，不需要额外的 Tomcat 等应用容器。

### Demo

> 先启动nacos
>
> 新建模块 `cloudalibaba-sentinel-service8401` ，使用nacos作为服务注册中心
>
> Sentinel可以对service进行监控、熔断、降级
>
> 没访问时再sentinel里是看不到监控的应用的，因为是懒加载，需要访问一次

pom依赖：

```xml
<dependencies>
    <!-- 后续做Sentinel的持久化会用到的依赖 -->
    <dependency>
        <groupId>com.alibaba.csp</groupId>
        <artifactId>sentinel-datasource-nacos</artifactId>
    </dependency>
    <!-- sentinel  -->
    <dependency>
        <groupId>com.alibaba.cloud</groupId>
        <artifactId>spring-cloud-starter-alibaba-sentinel</artifactId>
    </dependency>
    
    <!-- springcloud alibaba nacos 依赖,Nacos Server 服务注册中心 -->
    <dependency>
        <groupId>com.alibaba.cloud</groupId>
        <artifactId>spring-cloud-starter-alibaba-nacos-discovery</artifactId>
    </dependency>

    <!-- springboot整合Web组件 -->
    <dependency>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-web</artifactId>
    </dependency>
    <dependency>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-actuator</artifactId>
    </dependency>

    <!-- 日常通用jar包 -->
    <dependency>
        <groupId>org.mybatis.spring.boot</groupId>
        <artifactId>mybatis-spring-boot-starter</artifactId>
    </dependency>
    
    <dependency>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-devtools</artifactId>
        <scope>runtime</scope>
        <optional>true</optional>
    </dependency>
    <dependency>
        <groupId>org.projectlombok</groupId>
        <artifactId>lombok</artifactId>
        <optional>true</optional>
    </dependency>
    <dependency>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-test</artifactId>
        <scope>test</scope>
    </dependency>
    <dependency><!-- 引入自己定义的api通用包，可以使用Payment支付Entity -->
        <groupId>com.dkf.cloud</groupId>
        <artifactId>cloud-api-commons</artifactId>
        <version>${project.version}</version>
    </dependency>
</dependencies>
```

yml 配置：

```yml
server:
  port: 8401
spring:
  application:
    name: cloudalibaba-sentinel-service
  cloud:
    nacos:
      discovery:
        # 服务注册中心 # sentinel注册进nacos
        server-addr: localhost:8848
    sentinel:
      transport:
        # 配置 Sentinel Dashboard 的地址
        dashboard: localhost:8080
        # 默认8719 ，如果端口被占用，端口号会自动 +1，直到找到未被占用的端口，提供给 sentinel 的监控端口
        port: 8719
        
management:
  endpoints:
    web:
      exposure:
        include: '*'
```

写一个简单的主启动类，再写一个简单的controller测试sentinel的监控。

```java
@RestController
@Slf4j
public class FlowLimitController {
    @GetMapping("/testA")
    public String testA() {
        return "------testA";
    }

    @GetMapping("/testB")
    public String testB() {
        log.info(Thread.currentThread().getName()+"\t"+"...testB");
        return "------testB";
    }


    @GetMapping("/testD")
    public String testD() {
//        try { TimeUnit.SECONDS.sleep(1); } catch (InterruptedException e) { e.printStackTrace(); }
//        log.info("testD 测试RT");

        log.info("testD 异常比例");
        int age = 10/0;
        return "------testD";
    }

    @GetMapping("/testE")
    public String testE() {
        log.info("testE 测试异常数");
        int age = 10/0;
        return "------testE 测试异常数";
    }

    @GetMapping("/testHotKey")
    @SentinelResource(value = "testHotKey",blockHandler = "deal_testHotKey")
    public String testHotKey(@RequestParam(value = "p1",required = false) String p1,
                             @RequestParam(value = "p2",required = false) String p2) {
        //int age = 10/0;
        return "------testHotKey";
    }
    public String deal_testHotKey (String p1, String p2, BlockException exception) {
        return "------deal_testHotKey,o(╥﹏╥)o";  //sentinel系统默认的提示：Blocked by Sentinel (flow limiting)
    }

}
```

```java
@RestController
public class RateLimitController {
    @GetMapping("/byResource")
    @SentinelResource(value = "byResource",blockHandler = "handleException")
    public CommonResult byResource() {
        return new CommonResult(200,"按资源名称限流测试OK",new Payment(2020L,"serial001"));
    }
    public CommonResult handleException(BlockException exception) {
        return new CommonResult(444,exception.getClass().getCanonicalName()+"\t 服务不可用");
    }

    @GetMapping("/rateLimit/byUrl")
    @SentinelResource(value = "byUrl")
    public CommonResult byUrl() {
        return new CommonResult(200,"按url限流测试OK",new Payment(2020L,"serial002"));
    }


    @GetMapping("/rateLimit/customerBlockHandler")
    @SentinelResource(value = "customerBlockHandler",
            blockHandlerClass = CustomerBlockHandler.class,
            blockHandler = "handlerException2")
    public CommonResult customerBlockHandler() {
        return new CommonResult(200,"按客戶自定义",new Payment(2020L,"serial003"));
    }
}
```

```java
public class CustomerBlockHandler {

    public static CommonResult handlerException(BlockException exception) {
        return new CommonResult(4444,"按客戶自定义,global handlerException----1");
    }

    public static CommonResult handlerException2(BlockException exception) {
        return new CommonResult(4444,"按客戶自定义,global handlerException----2");
    }
}
```

```java
@EnableDiscoveryClient
@SpringBootApplication
public class MainApp8401 {
    public static void main(String[] args) {
        SpringApplication.run(MainApp8401.class, args);
    }
}
```



### 流控规则

- 资源名：唯一名称，默认请求路径
- 针对来源：sentinel可以针对调用者进行限流，填写微服务名，默认default（不区分来源）
- 阈值类型/单机值：
  - QPS（每秒钟的请求数量）：当调用该api就QPS达到阈值的时候，进行限流
  - 线程数．当调用该api的线程数达到阈值的时候，进行限流
- 单机/均摊阈值：和下面的选项有关
- 集群阈值模式：

  - 单机均摊：前面设置的阈值是`每台机器`的阈值
  - 总体阈值：前面设置的阈值是`集群总体`的阈值
- 流控模式：
  - 直接：api达到限流条件时，直接限流。分为QPS和线程数
  - 关联：当关联的资到阈值时，就限流自己。别人惹事，自己买单。当两个资源之间具有资源争抢或者依赖关系的时候，这两个资源便具有了关联。，举例来说，`read_db` 和 `write_db` 这两个资源分别代表数据库读写，我们可以给 `read_db` 设置限流规则来达到写优先的目的：设置 `strategy` 为 `RuleConstant.STRATEGY_RELATE` 同时设置 `refResource` 为 `write_db`。这样当写库操作过于频繁时，读数据的请求会被限流。
  - 链路：只记录指定链路上的流量（指定资源从入口资源进来的流量，如果达到阈值，就进行限流）【api级别的针对来源】
- 流控效果：
  - 快速失败：直接拒绝。当QPS超过任意规则的阈值后，新的请求就会被立即拒绝，拒绝方式为抛出`FlowException`
  - warm up：若干秒后才能达到阈值。当系统长期处于低水位的情况下，当流量突然增加时，直接把系统拉升到高水位可能瞬间把系统压垮。通过"冷启动"，让通过的流量缓慢增加，在一定时间内逐渐增加到阈值上限，给冷系统一个预热的时间，避免冷系统被压垮
  - 排队等待：让请求以均匀的速度通过
    ![image](https://i0.hdslb.com/bfs/album/cb949546aa3129de4eb10135f66e2af3176af112.png)





重要属性：

| Field           | 说明                                                         | 默认值                        |
| --------------- | ------------------------------------------------------------ | ----------------------------- |
| resource        | 资源名，资源名是限流规则的作用对象                           |                               |
| count           | 限流阈值                                                     |                               |
| grade           | 限流阈值类型，QPS 模式（1）或并发线程数模式（0）             | QPS 模式                      |
| limitApp        | 流控针对的调用来源                                           | `default`，代表不区分调用来源 |
| strategy        | 调用关系限流策略：直接、链路、关联                           | 根据资源本身（直接）          |
| controlBehavior | 流控效果（直接拒绝/WarmUp/匀速+排队等待），不支持按调用关系限流 | 直接拒绝                      |
| clusterMode     | 是否集群限流                                                 | 否                            |

我们先只针对/testA请求进行限制

##### **流控模式--直接**：

![1597819546992](https://i0.hdslb.com/bfs/album/67370f7e8ba87dd6c1bbb61da1c9859f600a7f6e.png)

> 限流表现：当超过阀值，就会被降级。
>
> 1s内多次刷新网页，localhost:8401/testA
>
> 返回Blocked by Sentienl(flow limiting)

##### **流控模式--关联**：

- 当与A关联的资源B达到阀值后，就限流A自己
- B惹事，A挂了。支付达到阈值，限流下单接口。B阈值达到1，A就挂
- 用post访问B让B忙，访问A发现挂了

![1597820015308](https://i0.hdslb.com/bfs/album/9e69cd5a9e96109a4c99112e430b606f5d1ccccf.png)

##### **流控效果--预热Warm up**：

访问数量慢慢升高

阈值初一coldFactor（默认3），经过预热时长后才会达到阈值。

<img src="https://i0.hdslb.com/bfs/album/4f4eb858a92fc02cf8344470090f07ffb32137e6.png" alt="1597820393038" style="zoom:67%;" /><img src="https://i0.hdslb.com/bfs/album/2e444559c0b25c7ecf0ba8cf7736cf5f9dff118f.png" alt="1597820534168" style="zoom:67%;" />

##### **流控效果--排队等待**：

匀速排队（Ru1eConstant.CONTROL_BEHAVIOR_RATE_LIMITER）方式会严格控制请求通过的间隔时间，即让请求以均匀的速度通过对应的是漏桶算法。详细文档可以参考流量控制-匀速器模式，具体的例子可以参见PaceFlowDemo

该方式的作用如下图所示

<img src="https://i0.hdslb.com/bfs/album/d955328e107384a16d142be71220f7c66b7432aa.png" style="zoom: 33%;" />

这种方式主要用于处理间隔性突发的流量，伊消息列。想象一下这样的场景，在某一秒有大量的请求到来，而接下来的月耖则处于空闲状态，我们希系统能够在接下来的空闲期间逐渐处理这些请求，而不是第一秒就拒绝多余的请求

### 熔断降级

新增降级规则：

##### 降低策略：RT

RT（平均响应时间，秒级）

**平均响应时间 超出阈值** 且 **在时间窗口内通过的请求>=5**，两个条件同时满足后触发降级

窗口期过后关闭断路器

RT最大4900（更大的需要通过-Dcsp.Sentinel.statistic.max.rt=XXXX才能生效）

<img src="https://i0.hdslb.com/bfs/album/d7486f1bdf3754e21f5584702bc579a65db424f7.png" alt="1597821156100" style="zoom:50%;" />

##### 异常比例（秒级）

QPS>=5且异常比例（秒级统计）超过阈值时，触发降级，时间窗口结束后，关闭降级。异常阈值范围是[0.0,l.0]，代表0％一100％。

sentinel熔断降级会在调用链路中某个资源出现不稳定状态时（例如调用超时或异常比例升高），对这个资源的调用进行限制，让请求快速失败，避免影响到其它的资源而导致级联错误。

当资源被降级后，在接下来的降级时间窗囗之内，对该资源的调用都自动熔断（默认行为是抛出DegradeException)。



##### 异常数：

异常数（DEGRADE-GRADE-EXCEPTION-COUNT）：当资源近1分钟的异常数目超过阈值之后会进行熔断。注意由于统计时间窗口是分钟级别的，若timeWindow小于60s,则结束熔断状态后仍可能再进入熔断状态。

时间窗口一定要大于等于60秒。

时间窗口结束后关闭降级



localhost:8401/testE  ,  第一次访问绝对报错，因为除数不能为零，
我们看到error窗口，但是达到5次报错后，进入熔断后降级。

降级策略里选择异常数

==如果没触发熔断,这正常抛出异常==:

如果触发熔断，显示的是blcoked by sentinel(flow limiting)

### 热点Key限流

何为热点？热点即经常访问的数据。很多时候我们希望统计某个热点数据中访问频次最高的TopK数据，并对其访问进行限制。比如：

- 商品ID为参数，统计一段时间内最常购买的商品ID并进行限制
- 用户ID为参数，针对一段时间内频繁访问的用户ID进行限制

参数限流会统计传入参数中的参数，并根据配置流阈值与模式，对包含热点参数的资源调用进行限流。热点参数限流可以看做是一种特殊的流量控制，仅对包含热点参数的资源调用生效。

比如:

​			localhost:8080/aa?name=aa

​			localhost:8080/aa?name=bb

​			加入两个请求中,带有参数aa的请求访问频次非常高,我们就现在name==aa的请求,但是bb的不限制

controller层写一个demo:

```java
	@GetMapping("/testhotkey")
    @SentinelResource(value = "testhotkey", blockHandler = "deal_testhotkey")
    //这个value是随意的值，并不和请求路径必须一致
    //在填写热点限流的 资源名 这一项时，可以填 /testhotkey 或者是 @SentinelResource的value的值
    public String testHotKey(
            @RequestParam(value="p1", required = false) String p1,
            @RequestParam(value = "p2", required = false) String p2
    ){
        return "testHotKey__success";
    }

	//类似Hystrix 的兜底方法
    public String deal_testhotkey(String p1, String p2, BlockException e){
        return "testhotkey__fail"; 
    }
```

![1597822501876](https://i0.hdslb.com/bfs/album/700ffb73a317dd906ccf21745f4062443250f67c.png)

![1597822772165](https://i0.hdslb.com/bfs/album/d311932c9e8fce4172cbb0a6329a6221a01cfde8.png)

说明：

@SentinelResource ：处理的是Sentine1控制台配置的违规情况，有blockHandler方法配置的兜底处理

@RuntimeException：int age=10/0，这个是java运行时报出的运行时异异常RunTimeException，@Sentine1Resource不管



### 系统规则

> 一般配置在网关或者入口应用中，但是这个东西有点危险，不但值不合适，就相当于系统瘫痪。

系统自适应限流

Sentinel 系统自适应限流从整体维度对应用入口流量进行控制，结合应用的 Load、CPU 使用率、总体平均 RT、入口 QPS 和并发线程数等几个维度的监控指标，通过自适应的流控策略，让系统的入口流量和系统的负载达到一个平衡，让系统尽可能跑在最大吞吐量的同时保证系统整体的稳定性。

系统规则包含下面几个重要的属性：

| Field             | 说明                                   | 默认值      |
| ----------------- | -------------------------------------- | ----------- |
| highestSystemLoad | `load1` 触发值，用于触发自适应控制阶段 | -1 (不生效) |
| avgRt             | 所有入口流量的平均响应时间             | -1 (不生效) |
| maxThread         | 入口流量的最大并发数                   | -1 (不生效) |
| qps               | 所有入口资源的 QPS                     | -1 (不生效) |
| highestCpuUsage   | 当前系统的 CPU 使用率（0.0-1.0）       | -1 (不生效) |



### @SentinelResource配置

> @SentinelResource 注解，主要是指定资源名（也可以用请求路径作为资源名），和指定降级处理方法的。
>
> **使用@SentinelResource直接实现降级方法,它等同Hystrix的@HystrixCommand**

例如： 

```java
package com.dkf.springcloud.controller;

import com.alibaba.csp.sentinel.annotation.SentinelResource;
import com.alibaba.csp.sentinel.slots.block.BlockException;
import com.dkf.springcloud.entities.CommonResult;
import com.dkf.springcloud.entities.Payment;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class RateLimitController {

    @GetMapping("/byResource")						//处理降级的方法名
    @SentinelResource(value = "byResource", blockHandler = "handleException")
    public CommonResult byResource(){
        return new CommonResult(200, "按照资源名限流测试0K", new Payment(2020L,"serial001"));
    }

    //降级方法
    public CommonResult handleException(BlockException e){
        return new CommonResult(444, e.getClass().getCanonicalName() + "\t 服务不可用");
    }
}
```

定义热点规则的时候可以指定参数索引，就是说带指定的参数时才会被限制（用了required=false）

![1597901945492](https://i0.hdslb.com/bfs/album/0108500b86f026bfc8e9a3038c0872459af08693.png)

很明显，上面虽然自定义了兜底方法，但是耦合度太高，下面要解决这个问题。

#### 自定义全局BlockHandler处理类

写一个 CustomerBlockHandler 自定义限流处理类：



![1597903188558](https://i0.hdslb.com/bfs/album/527c6d8568913e13771818b4456880f58b158f35.png)

### ==整合 openfeign 服务降级==

#### 前奏

> 之前有 open-feign 和 hystrix 的整合，现在来实现sentinel 整合 `ribbon + open-feign + fallback`进行服务熔断。
>
> 新建三个模块，两个提供者 9004、9005，和一个消费者 84
>
> 目的：
>
> fallback管运行异常
> blockHandIer管配置违规
>
> <font color=red>上面使用sentinel有一个很明显的问题，就是sentinel，对程序内部异常（各种异常，包括超时）这种捕捉，显得很乏力，它主要是针对流量控制，系统吞吐量，或者是异常比例这种，会发生降级或熔断，但是当程序内部发生异常，直接返回给用户错误页面，根本不会触发异常比例这种降级。所以才需要整合open-feign 来解决程序内部异常时，配置相应的兜底方法</font>

---------------两个提供者模块一致，如下：

pom依赖：

```xml
	<dependencies>
        <!-- springcloud alibaba nacos 依赖 -->
        <dependency>
            <groupId>com.alibaba.cloud</groupId>
            <artifactId>spring-cloud-starter-alibaba-nacos-discovery</artifactId>
        </dependency>
        
        <!-- springboot整合Web组件 -->
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-web</artifactId>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-actuator</artifactId>
        </dependency>

        <!-- 日常通用jar包 -->
        <dependency>
            <groupId>org.mybatis.spring.boot</groupId>
            <artifactId>mybatis-spring-boot-starter</artifactId>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-devtools</artifactId>
            <scope>runtime</scope>
            <optional>true</optional>
        </dependency>
        <dependency>
            <groupId>org.projectlombok</groupId>
            <artifactId>lombok</artifactId>
            <optional>true</optional>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-test</artifactId>
            <scope>test</scope>
        </dependency>
        <dependency><!-- 引入自己定义的api通用包，可以使用Payment支付Entity -->
            <groupId>com.dkf.cloud</groupId>
            <artifactId>cloud-api-commons</artifactId>
            <version>${project.version}</version>
        </dependency>
    </dependencies>
```

yml配置：

```yaml
server:
  port: 9005  # / 9004
spring:
  application:
    name: nacos-payment-provider
  cloud:
    nacos:
      discovery:
        server-addr: localhost:8848
management:
  endpoints:
    web:
      exposure:
        include: '*'
```

主启动类只是启动，没有其它注解。

controller : 

```java
package com.dkf.sprIngcloud.controller;

import com.dkf.springcloud.entities.CommonResult;
import com.dkf.springcloud.entities.Payment;

@RestController
public class PaymentController {

    @Value("${server.port}")
    private String serverPort;

    //模拟sql查询
    public static HashMap<Long, Payment> hashMap = new HashMap<>();
    static {
        hashMap.put(1L, new Payment(1L, "xcxcxcxcxcxcxcxcxcxcxcxcxc11111111"));
        hashMap.put(2L, new Payment(2L, "xcxcxcxcggggggggg2222222222222222"));
        hashMap.put(3L, new Payment(3L, "xcxcxcxccxxcxcfafdgdgdsgdsgds33333"));
    }


    @GetMapping("/payment/get/{id}")
    public CommonResult paymentSql(@PathVariable("id")Long id){
        Payment payment = hashMap.get(id);
        CommonResult result = new CommonResult(200, "from mysql, server port : " + serverPort + " ,查询成功", payment);
        return result;
    }
}
```

---消费者：

pom依赖：

```xml
	<dependencies>
        <!-- 后续做Sentinel的持久化会用到的依赖 -->
        <dependency>
            <groupId>com.alibaba.csp</groupId>
            <artifactId>sentinel-datasource-nacos</artifactId>
        </dependency>
        <!-- sentinel  -->
        <dependency>
            <groupId>com.alibaba.cloud</groupId>
            <artifactId>spring-cloud-starter-alibaba-sentinel</artifactId>
        </dependency>
        <!-- springcloud alibaba nacos 依赖 -->
        <dependency>
            <groupId>com.alibaba.cloud</groupId>
            <artifactId>spring-cloud-starter-alibaba-nacos-discovery</artifactId>
        </dependency>

        <!-- springboot整合Web组件 -->
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-web</artifactId>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-actuator</artifactId>
        </dependency>

        <!-- 日常通用jar包 -->
        <dependency>
            <groupId>org.mybatis.spring.boot</groupId>
            <artifactId>mybatis-spring-boot-starter</artifactId>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-devtools</artifactId>
            <scope>runtime</scope>
            <optional>true</optional>
        </dependency>
        <dependency>
            <groupId>org.projectlombok</groupId>
            <artifactId>lombok</artifactId>
            <optional>true</optional>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-test</artifactId>
            <scope>test</scope>
        </dependency>
        <dependency><!-- 引入自己定义的api通用包，可以使用Payment支付Entity -->
            <groupId>com.dkf.cloud</groupId>
            <artifactId>cloud-api-commons</artifactId>
            <version>${project.version}</version>
        </dependency>
    </dependencies>
```

yml配置：

```yaml
server:
  port: 84
spring:
  cloud:
    nacos:
      discovery:
        server-addr: localhost:8848
    sentinel:
      transport:
        dashboard: localhost:8080
        port: 8719
  application:
    name: nacos-order-consumer
```

主启动类不用说了。

config类里面注入 Resttemplate：

```java
@Configuration
public class ApplicationContextConfig {
    @Bean
    @LoadBalanced
    public RestTemplate getRestTemplate(){
        return new RestTemplate();
    }
}
```

controller 层：

```java
@RestController
public class OrderController {

    private static final String PAYMENT_URL="http://nacos-payment-provider";

    @Resource
    private RestTemplate restTemplate;

    @GetMapping("/consutomer/payment/get/{id}")
    public CommonResult getPayment(@PathVariable("id")Long id){
        if(id >= 4){
            throw new IllegalArgumentException("非法参数异常...");
        }else {
            return restTemplate.getForObject(PAYMENT_URL + "/payment/get/" + id, CommonResult.class);
        }
    }
}
```

上面只实现了 以nacos 作为服务注册中心，消费者使用ribbon 实现负载均衡调用提供者的效果。

#### 正式

只配置 fallback:

```java
@GetMapping("/consutomer/payment/get/{id}")
@SentinelResource(value = "fallback",
                  fallback = "handleFallback") //fallback只处理业务异常
public CommonResult getPayment(@PathVariable("id")Long id){
    if(id >= 4){
        throw new IllegalArgumentException("非法参数异常...");
    }else {
        return restTemplate.getForObject(PAYMENT_URL + "/payment/get/" + id, CommonResult.class);
    }
}

//兜底方法
public CommonResult handleFallback(@PathVariable("id")Long id,//参数一直
                                   Throwable e){ // 接收异常的参数
    return new CommonResult(414, "---非法参数异常--", e);
}
```

> 业务异常会被 fallback 处理，返回我们自定义的提示信息，而如果给它加上流控，并触发阈值，只能返回sentinel默认的提示信息。

只配置blockHandler:

```java
//@SentinelResource(value = "fallback", fallback = "handleFallback") //fallback只处理业务异常
@GetMapping("/consutomer/payment/get/{id}")
@SentinelResource(value = "fallback",
                  blockHandler = "handleblockHandler")
public CommonResult getPayment(@PathVariable("id")Long id){
    if(id >= 4){
        throw new IllegalArgumentException("非法参数异常...");
    }else {
        return restTemplate.getForObject(PAYMENT_URL + "/payment/get/" + id, CommonResult.class);
    }
}

//    //====fallback
//    public CommonResult handleFallback(@PathVariable("id")Long id, Throwable e){
//        return new CommonResult(414, "---非法参数异常--", e);
//    }

//====blockHandler                                       blockHandler的方法必须有这个参数
public CommonResult handleblockHandler(@PathVariable("id")Long id,
                                       BlockException e){
    return new CommonResult(414, "---非法参数异常--", e);
}
```

> 这时候的效果就是，运行异常直接报错错误页面。在sentinel上添加一个降级规则，设置2s内触发异常2次，触发阈值以后，返回的是我们自定义的 blockhanlder 方法返回的内容。

两者都配置：

```java
//@SentinelResource(value = "fallback", fallback = "handleFallback") //fallback只处理业务异常
@GetMapping("/consutomer/payment/get/{id}")
@SentinelResource(value = "fallback",
                  blockHandler = "handleblockHandler", 
                  fallback = "handleFallback")
public CommonResult getPayment(@PathVariable("id")Long id){
    if(id >= 4){
        throw new IllegalArgumentException("非法参数异常...");
    }else {
        return restTemplate.getForObject(PAYMENT_URL + "/payment/get/" + id, CommonResult.class);
    }
}
//====fallback
public CommonResult handleFallback(@PathVariable("id")Long id, Throwable e){
    return new CommonResult(414, "---非法参数异常--form fallback的提示", e);
}

//====blockHandler                                       blockHandler的方法必须有这个参数
public CommonResult handleblockHandler(@PathVariable("id")Long id, BlockException e){
    return new CommonResult(414, "---非法参数异常--", e);
}
```

> 明显两者都是有效的，可以同时配置。

#### 全局降级

> 上面是单个进行 fallback 和 blockhandler 的测试，下面是整合 openfeign 实现把降级方法解耦。和Hystrix 几乎一摸一样！

还是使用上面 84 这个消费者做测试：

1. 先添加open-feign依赖：

```xml
<dependency>
    <groupId>org.springframework.cloud</groupId>
    <artifactId>spring-cloud-starter-openfeign</artifactId>
</dependency>
```

2. yml 追加如下配置：

```yaml
# 激活Sentinel对Feign的支持
feign:
  sentinel:
    enabled: true
```

3. 主启动类添加注解 ： @EnableFeignClients  激活open-feign
4. service : 

```java
@FeignClient(value = "nacos-payment-provider", fallback = PaymentServiceImpl.class)
public interface PaymentService {

    @GetMapping("/payment/get/{id}")
    public CommonResult paymentSql(@PathVariable("id")Long id);
}
```

5. service 实现类：

```java
@Component
public class PaymentServiceImpl implements PaymentService {

    @Override
    public CommonResult paymentSql(Long id) {
        return new CommonResult(414, "open-feign 整合 sentinel 实现的全局服务降级策略",null);
    }
}
```

6. controller 层代码没什么特殊的，和普通调用service 一样即可。
7. 测试，关闭提供者的项目，会触发 service 实现类的方法。
8. 总结: 这种全局熔断，是针对 “访问提供者” 这个过程的，只有访问提供者过程中发生异常才会触发降级，也就是这些降级，是给service接口上这些提供者的方法加的，以保证在远程调用时能顺利进行。而且这明显是 fallback ，而不是 blockHandler，注意区分。

> fallback 和 blockHandler 肤浅的区别：
>
> F ： 不需要指定规则，程序内部异常均可触发（超时异常需要配置超时时间）
>
> B :  配上也没用，必须去 Sentinel 上指定规则才会被触发。 

### 异常忽略

> 这是 @SentinelResource 注解的一个值：`@SentinelResource(exceptionsToIgnore={IllegalArgumentException.class})`
>
> 报错指定异常类的话，不执行fallback兜底方法，没有降级效果
>
> 

### sentinel持久化

> 目前的sentinel 当重启以后，数据都会丢失，和 nacos 类似原理。需要持久化。它可以被持久化到 nacos 的数据库中。

1. pom依赖：

```xml
<dependency>
    <groupId>com.alibaba.csp</groupId>
    <artifactId>sentinel-datasource-nacos</artifactId>
</dependency>
```

2. yml配置：

```yaml
spring:
  cloud:
    sentinel:
      datasource:
        ds1:  
          nacos:
            server-addr: localhost:8848
            dataId: ${spring.application.name}
            group: DEFAULT_GROUP
            data-type: json
            rule-type: flow
```

3. 去nacos上创建一个dataid ,名字和yml配置的一致，json格式，内容如下：

```json
[
    {
        "resource": "/testA",
        "limitApp": "default",
        "grade": 1,
        "count": 1,
        "strategy": 0,
        "controlBehavior": 0,
        "clusterMode": false
    }
]
```

resource：资源名称
limitApp：来源应用
grade：阈值类型，0表示线程数，1表示QPS，
count：单机阈值，
strategy：流控模式，0表示直接，1表示关联，2表示链路

controlBehavior:流控效果，0表示快速失败，1表示Warm Up,2表示排队等待；

cIusterM0de是否集群。

4. 启动应用，发现存在 关于 /testA 请求路径的流控规则。
5. 总结: 就是在 sentinel 启动的时候，去 nacos 上读取相关规则配置信息，实际上它规则的持久化，就是第三步，粘贴到nacos上保存下来，就算以后在 sentinel 上面修改了，重启应用以后也是无效的。

## ==Seata==

> Seate 处理分布式事务。
>
> 微服务模块，连接多个数据库，多个数据源，而数据库之间的数据一致性需要被保证。
>
> 官网：  http://seata.io/zh-cn/ 

Seata术语： 一 + 三


Transaction ID XID：全局唯一的事务ID

3组件概念

- Transaction Coordinator(TC)：事务协调器，**维护全局事务**的运行状态，负责协调并驱动全局事务的提交或回滚，
- Transaction Manager(TM)：控制全局事务的边界，负责开启一个全局事务，并最终发起全局提交或全局回滚的决议，
- Resource Manager(RM)：控制分支事务，负责分支汪册、**状态汇报**，并接收事务协调器的指令，驱动分支（本地）事务提交或回滚

<img src="http://seata.io/img/solution.png" alt="img" style="zoom: 80%;" />

1．TM向TC申请开启一个全局事务，全局事务创建成功并生成一个全局唯一的XID;
2，XID在微服务调用链路的上下文中传播；
3，RM向TC汪册分支事务，将其纳入XID对应全局事务的管辖；
4，TM向TC发起针对XID的全局提交或回滚决议；
5，TC调度XID下管辖的全部分支事务完成提交或回请求。

### 下载安装

> 下载地址 ： https://github.com/seata/seata/releases/download/v1.0.0/seata-server-1.0.0.zip

我们只需要使用一个`@GlobalTransational`注解在业务方法上

### 初始化操作

1. 修改 conf/file.conf 文件：

> 主要修改自定义事务组名称 + 事务日志存储模式为db + 数据库连接信息
>
> ```ini
> service {
>   #transaction service group mapping
>   vgroup_mapping.dkf_tx_group = "default"   # 修改这里
>   #only support when registry.type=file, please don't set multiple addresses
>   default.grouplist = "127.0.0.1:8091"
>   #disable seata
>   disableGlobalTransaction = false
> }
> 
> ## transaction log store, only used in seata-server
> store {
>   ## store mode: file、db
>   mode = "db"   # 修改这里
> 
>   ## file store property
>   file {
>     ## store location dir
>     dir = "sessionStore"
>   }
> 
>   ## database store property
>   db {
>     ## the implement of javax.sql.DataSource, such as DruidDataSource(druid)/BasicDataSource(dbcp) etc.
>     datasource = "dbcp"
>     ## mysql/oracle/h2/oceanbase etc.
>     db-type = "mysql"
>     driver-class-name = "com.mysql.jdbc.Driver"
>     url = "jdbc:mysql://127.0.0.1:3306/seata"
>     user = "root"   # 修改对
>     password = "123456"
>   }
> }
> ```

2. 创建名和 file.conf 指定一致的数据库。

3. 在新建的数据库里面创建数据表，db_store.sql文件在 conf 目录下（1.0.0有坑，没有sql文件，下载0.9.0的，使用它的sql文件即可）

4. 修改 conf/registry.conf 文件内容：

   ```java
   registry {
     # file 、nacos 、eureka、redis、zk、consul、etcd3、sofa # 默认file
     type = "nacos"
   
     nacos {  # 修改nacos的端口8848
       serverAddr = "localhost:8848"
       namespace = ""
       cluster = "default"
     }
   ```

5. 先启动 nacos Server 服务，再启动seata Server 。

6. 启动 Seata Server 报错，在bin目录创建 /logs/seata_gc.log 文件。再次双击 bat文件启动。

### 案例

#### 数据库准备

这里我们会创建三个服务，一个订单服务，一个库存服务，一个账户服务。

当用户下单时，会在订单服务中创建一个订单，然后通过远程调用库存服务来扣减下单商品的库存，
再通过远程调用账户服务来扣减用户账户里面的余额，
最后在订单服务中修改订单状态为已完成。

该操作跨越三个数据库，有两次远程调用，很明显会有分布式事务问题。



创建三个数据库： `seata_account、seata_order、seata_storage`

每个数据库创建数据表：

order 库： 

```mysql
CREATE DATABASE seata_order;
USE seata_order;
CREATE TABLE t_order(
    id BIGINT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY ,
    user_id BIGINT(11) DEFAULT NULL COMMENT '用户id',
    product_id BIGINT(11) DEFAULT NULL COMMENT '产品id',
    count INT(11) DEFAULT NULL COMMENT '数量',
    money DECIMAL(11,0) DEFAULT NULL COMMENT '金额',
    status INT(1) DEFAULT NULL COMMENT '订单状态：0创建中，1已完结'
)ENGINE=InnoDB AUTO_INCREMENT=7 CHARSET=utf8;
```

account 库： 

```mysql
CREATE DATABASE seata_account;
USE seata_account;
CREATE TABLE t_account(
    id BIGINT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY ,
    user_id BIGINT(11) DEFAULT NULL COMMENT '用户id',
    total DECIMAL(10,0) DEFAULT NULL COMMENT '总额度',
    used DECIMAL(10,0) DEFAULT NULL COMMENT '已用额度',
    residue DECIMAL(10,0) DEFAULT 0 COMMENT '剩余可用额度'
)ENGINE=InnoDB AUTO_INCREMENT=7 CHARSET=utf8;
INSERT INTO t_account(id, user_id, total, used, residue) VALUES(1,1,1000,0,1000);
```

storage 库：

```mysql
CREATE DATABASE seata_storage;
USE seata_storage;
CREATE TABLE t_storage(
    id BIGINT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY ,
    product_id BIGINT(11) DEFAULT NULL COMMENT '产品id',
    total INT(11) DEFAULT NULL COMMENT '总库存',
    used INT(11) DEFAULT NULL COMMENT '已用库存',
    residue INT(11) DEFAULT NULL COMMENT '剩余库存'
)ENGINE=InnoDB AUTO_INCREMENT=7 CHARSET=utf8;
INSERT INTO t_storage(id, product_id, total, used, residue) VALUES(1,1,100,0,100);
```

三个数据库都创建一个回滚日志表，seata/conf/ 有相应的sql文件（1.0.0没有，依然使用0.9.0中的）。

最终效果：

- seata
  - branch_table
  - global_table
  - lock_table
- seata_account
  - t_account
  - undo_log
- seata_order
  - t_order
  - undo_log
- seata_storage
  - t_storage
  - undo_log

#### 开发

> 实现 下订单-> 减库存 -> 扣余额 -> 改（订单）状态
>
> 需要注意的是，下面做了 seata 与 mybatis 的整合，所以注意一下，和以往的mybatis的使用不太一样。

新建模块 cloudalibaba-seata-order2001 ：

pom依赖：

```xml
	<dependencies>
        <!-- seata -->
        <dependency>
            <groupId>com.alibaba.cloud</groupId>
            <artifactId>spring-cloud-starter-alibaba-seata</artifactId>
            <exclusions>
                <exclusion>
                    <artifactId>seata-all</artifactId>
                    <groupId>io.seata</groupId>
                </exclusion>
            </exclusions>
        </dependency>
        <dependency>
            <groupId>io.seata</groupId>
            <artifactId>seata-all</artifactId>
            <version>1.0.0</version>
        </dependency>
        <!-- springcloud alibaba nacos 依赖,Nacos Server 服务注册中心 -->
        <dependency>
            <groupId>com.alibaba.cloud</groupId>
            <artifactId>spring-cloud-starter-alibaba-nacos-discovery</artifactId>
        </dependency>

        <!-- open feign 服务调用 -->
        <dependency>
            <groupId>org.springframework.cloud</groupId>
            <artifactId>spring-cloud-starter-openfeign</artifactId>
        </dependency>

        <!-- springboot整合Web组件 -->
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-web</artifactId>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-actuator</artifactId>
        </dependency>

        <!-- 持久层支持 -->
        <dependency>
            <groupId>com.alibaba</groupId>
            <artifactId>druid-spring-boot-starter</artifactId>
            <version>1.1.10</version>
        </dependency>
        <!--mysql-connector-java-->
        <dependency>
            <groupId>mysql</groupId>
            <artifactId>mysql-connector-java</artifactId>
        </dependency>
        <!--jdbc-->
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-jdbc</artifactId>
        </dependency>
        <!-- mybatis -->
        <dependency>
            <groupId>org.mybatis.spring.boot</groupId>
            <artifactId>mybatis-spring-boot-starter</artifactId>
        </dependency>

        <!-- 日常通用jar包 -->
        <dependency>
            <groupId>org.mybatis.spring.boot</groupId>
            <artifactId>mybatis-spring-boot-starter</artifactId>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-devtools</artifactId>
            <scope>runtime</scope>
            <optional>true</optional>
        </dependency>
        <dependency>
            <groupId>org.projectlombok</groupId>
            <artifactId>lombok</artifactId>
            <optional>true</optional>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-test</artifactId>
            <scope>test</scope>
        </dependency>
        <dependency><!-- 引入自己定义的api通用包，可以使用Payment支付Entity -->
            <groupId>com.dkf.cloud</groupId>
            <artifactId>cloud-api-commons</artifactId>
            <version>${project.version}</version>
        </dependency>
    </dependencies>
```

yml配置：

```yaml
server:
  port: 2001
spring:
  application:
    name: seata-order-service
  cloud:
    alibaba:
      seata:
        # 自定义事务组，需要和当时在 seata/conf/file.conf 中的一致
        tx-service-group: dkf_tx_group
    nacos:
      discovery:
        server-addr: localhost:8848
  datasource:
    driver-class-name: com.mysql.jdbc.Driver
    url: jdbc:mysql://localhost:3306/seata_order
    username: root
    password: 123456


# 注意，这是自定义的，原来的是mapper_locations
mybatis:
  mapperLocations: classpath:mapper/*.xml

logging:
  level:
    io:
      seata: info
```

将 seata/conf/ 下的 file.conf 和 registry.cong 两个文件拷贝到 resource 目录下。

创建 domain 实体类 ： Order 和 CommonResult<T> 两个实体类。

dao :

```java
package com.dkf.springcloud.dao;

import org.apache.ibatis.annotations.Mapper;
import com.dkf.springcloud.domain.Order;
import org.apache.ibatis.annotations.Param;

@Mapper
public class OrderDao {

    //创建订单
    public void create(Order order);

    //修改订单状态
    public void update(@Param("userId") Long userId, @Param("status") Integer status);

}
```

Mapper文件：

```xml
<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE mapper PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN" "http://mybatis.org/dtd/mybatis-3-mapper.dtd" >

<mapper namespace="com.dkf.springcloud.dao.OrderDao">

    <!-- 以备后面会用到 -->
    <resultMap id="BaseResultMap" type="com.dkf.springcloud.domain.Order">
        <id column="id" property="id" jdbcType="BIGINT"></id>
        <result column="user_id" property="userId" jdbcType="BIGINT"></result>
        <result column="product_id" property="productId" jdbcType="BIGINT"></result>
        <result column="count" property="count" jdbcType="INTEGER"></result>
        <result column="money" property="money" jdbcType="DECIMAL"></result>
        <result column="status" property="status" jdbcType="INTEGER"></result>
    </resultMap>

    <insert id="create">
        insert into t_order(id, user_id, product_id, count, money, status)
        values (null, #{userId},#{productId},#{count},#{money},0)
    </insert>

    <update id="update">
        update t_order set status = 1 where user_id=#{userId} and status=#{status}
    </update>

</mapper>
```

创建service ： 

> 注意，红框标记的是通过 open-feign 远程调用微服务的service
>
> 远程服务接口有AccountService、StorageServer
>
> 而OrderServiceImpl实现我们的业务逻辑

serviceImpl : 

```java
@Service
@Slf4j
public class OrderServiceImpl  implements OrderService {

    @Resource
    private OrderDao orderDao;

    @Resource
    private StorageService storageService;

    @Resource
    private AccountService accountService;

    @Override
    public void create(Order order) {
        log.info("--------》 开始创建订单");
        orderDao.create(order);
        log.info("--------》 订单微服务开始调用库存，做扣减---Count-");
        storageService.decrease(order.getProductId(), order.getCount());
        log.info("--------》 订单微服务开始调用库存，库存扣减完成！！");
        log.info("--------》 订单微服务开始调用账户，账户扣减---money-");
        accountService.decrease(order.getUserId(),order.getMoney());
        log.info("--------》 订单微服务开始调用账户，账户扣减完成!!");
        //修改订单状态，从0到1
        log.info("--------》 订单微服务修改订单状态，start");
        orderDao.update(order.getUserId(),0);
        log.info("--------》 订单微服务修改订单状态，end");

        log.info("--订单结束--");
    }

    @Override
    public void update(Long userId, Integer status) {

    }
}
```

config （特殊点）:

```java
//下面是两个配置类，这个是和mybatis整合需要的配置
@Configuration
@MapperScan({"com.dkf.springcloud.alibaba.dao"})
public class MybatisConfig {
}


//这个是配置使用 seata 管理数据源，所以必须配置
package com.dkf.springcloud.config;

import com.alibaba.druid.pool.DruidDataSource;
import io.seata.rm.datasource.DataSourceProxy;
import org.apache.ibatis.session.SqlSessionFactory;
import org.mybatis.spring.SqlSessionFactoryBean;
import org.mybatis.spring.transaction.SpringManagedTransactionFactory;

import org.springframework.core.io.support.PathMatchingResourcePatternResolver;

import javax.sql.DataSource;

@Configuration
public class DataSourceProxyConfig {

    @Value("${mybatis.mapperLocations}")
    private String mapperLocations;

    @Bean
    @ConfigurationProperties(prefix = "spring.datasource")
    public DataSource druidDataSource(){
        return new DruidDataSource();
    }

    @Bean
    public DataSourceProxy dataSourceProxy(DataSource dataSource){
        return new DataSourceProxy(dataSource);
    }

    @Bean
    public SqlSessionFactory sqlSessionFactoryBean(DataSourceProxy dataSourceProxy) throws Exception {
        SqlSessionFactoryBean sqlSessionFactoryBean = new SqlSessionFactoryBean();
        sqlSessionFactoryBean.setDataSource(dataSourceProxy);
        sqlSessionFactoryBean.setMapperLocations(new PathMatchingResourcePatternResolver().getResources(mapperLocations));
        sqlSessionFactoryBean.setTransactionFactory(new SpringManagedTransactionFactory());
        return sqlSessionFactoryBean.getObject();
    }
}
```

主启动类：

```java
//这里必须排除数据源自动配置，因为写了配置类，让 seata 管理数据源
@SpringBootApplication(exclude = DataSourceAutoConfiguration.class)
@EnableFeignClients
@EnableDiscoveryClient
public class SeataOrderMain2001 {

    public static void main(String[] args) {
        SpringApplication.run(SeataOrderMain2001.class,args);
    }
}
```

controller 层调用 orderService 方法即可。

先启动 nacos  --》 再启动 seata --> 再启动此order服务，测试，可以启动。

仿照上面 创建  cloudalibaba-seata-storage2002 和  cloudalibaba-seata-account2003 两个模块，唯一大的区别就是这两个不需要导入 open-feign 远程调用其它模块。

操，累死老子啦，测试可以正常使用！

#### Seata使用

```java
	@Override
	//只需要在业务类的方法上加上该注解，name值自定义唯一即可。
    @GlobalTransactional(name = "dkf-create-order", rollbackFor = Exception.class)
    public void create(Order order) {
        log.info("--------》 开始创建订单");
        orderDao.create(order);
        log.info("--------》 订单微服务开始调用库存，做扣减---Count-");
        storageService.decrease(order.getProductId(), order.getCount());
        log.info("--------》 订单微服务开始调用库存，库存扣减完成！！");
        log.info("--------》 订单微服务开始调用账户，账户扣减---money-");
        accountService.decrease(order.getUserId(),order.getMoney());
        log.info("--------》 订单微服务开始调用账户，账户扣减完成!!");
        //修改订单状态，从0到1
        log.info("--------》 订单微服务修改订单状态，start");
        orderDao.update(order.getUserId(),0);
        log.info("--------》 订单微服务修改订单状态，end");

        log.info("--订单结束--");
    }
```

```
                      TC :seata 服务器
             
             
   TM                                     RM
  @GlobalTransational                事务的参与方
  事务的发起方
```

原理三个阶段：

![1597999156669](https://i0.hdslb.com/bfs/album/44d566c3a2785fe3cbbc9a61939f37438a6f3b85.png)

![1597999227092](https://i0.hdslb.com/bfs/album/8967825edcbf0a5f63a35600bd3d7943b43978b0.png)

二阶段如是顶利提交的话，因为"业务SQL"在一阶段已经提交至数据库，所以Seata框架只需将一阶段保存的快照数倨和行锁删掉，完成数倨清理即可。