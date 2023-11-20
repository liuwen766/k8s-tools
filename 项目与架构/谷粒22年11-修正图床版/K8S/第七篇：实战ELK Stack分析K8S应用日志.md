# 前情回顾

- 新一代容器云监控系统P+G
  - Exporters（可以自定义开发）
    - HTTP接口
    - 定义监控项和监控项的标签（维度）
    - 按一定的数据结构组织监控数据
    - 以时间序列被收集
  - Prometheus Server
    - Retrieve（数据收集器）
    - TSDB（时间序列数据库）
    - Configure（static_config、kubernetes_sd、file_sd）
    - HTTP Server
  - Grafana
    - 多种多样的插件
    - 数据源（Prometheus）
    - Dashboard（PromQL）
  - Alertmanager
    - rules.yml（PromQL）

# 第一章：ELK Stack概述

- 日志，对于任何系统来说都是及其重要的组成部分。在计算机系统里面，更是如此。但是由于现在的计算机系统大多比较复杂，很多系统都不是在一个地方，甚至都是跨国界的；即使是在一个地方的系统，也有不同的来源，比如，操作系统，应用服务，业务逻辑等等。他们都在不停的产生各种各样的日志数据。根据不完全统计，我们全球每天大约要产生2EB的数据。
- K8S系统里的业务应用是高度“动态化”的，随着容器编排的进行，业务容器在不断的被创建、被摧毁，被迁移（漂）、被扩缩容….
- 面对如此海量的数据，又是分布在各个不同地方，如果我们需要去查找一些重要的信息，难道还是使用传统的方法，去登录到一台台机器上查看？看来传统的工具和方法已经显得非常笨拙和低效了。于是，一些聪明人就提出了建立一套中式的方法，把不同来源的数据集中整合到一个地方。

- 我们需要这样一套日志收集、分析的系统：
  - 收集 – 能够采集多种来源的日志数据（流式日志收集器）
  - 传输 – 能够稳定的把日志数据传输到中央系统（消息队列）
  - 存储 – 可以将日志以结构化数据的形式存储起来（搜索引擎）
  - 分析 – 支持方便的分析、检索方法。最好有GUI管理系统（前端）
  - 警告 – 能够提供错误报告，监控机制（监控工具）
- 优秀的社区开源解决方案 – ELK Stack
  - E – ElasticSearch
  - L – LogStash
  - K – Kibana
- 传统ELK模型

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_3f187af3da8dccbfb52c7702bf22ea85_r.png)

- 缺点
  - Logstach使用Jruby语言考开发，吃资源，大量部署消耗极高
  - 业务程序与logstash耦合过松，不利于业务迁移
  - 日志收集与ES耦合过紧，易打爆、丢数据
  - 在容器云环境下，传统ELK模型难以完成工作
- 容器化ELK模型

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_47c2a2feae8680b870643b3afac1a163_r.png)

# 第二章：将dubbo-demo-consumer项目改造为tomcat启动的web项目

[Tomcat官网](http://tomcat.apache.org/)

## 1.准备Tomcat的镜像底包

### 1.准备tomcat二进制包

运维主机shkf6-245.host.com上：

[Tomcat8下载链接](https://mirrors.tuna.tsinghua.edu.cn/apache/tomcat/tomcat-8/v8.5.50/bin/)

```shell
[root@shkf6-245 ~]# cd /opt/src
[root@shkf6-245 src]# wget http://mirrors.tuna.tsinghua.edu.cn/apache/tomcat/tomcat-8/v8.5.50/bin/apache-tomcat-8.5.50.tar.gz
[root@shkf6-245 src]# ls -l|grep tomcat
-rw-r--r-- 1 root root 10305939 Dec  8 03:42 apache-tomcat-8.5.50.tar.gz
[root@shkf6-245 src]# mkdir -p /data/dockerfile/tomcat8 && tar xf apache-tomcat-8.5.50.tar.gz -C /data/dockerfile/tomcat8
```

### 2.简单配置tomcat

1.关闭AJP端口

```shell
[root@shkf6-245 src]# vi /data/dockerfile/tomcat8/apache-tomcat-8.5.50/conf/server.xml

  <!--  <Connector port="8009" protocol="AJP/1.3" redirectPort="8443" /> -->
```

2.配置日志

- 删除3manager，4host-manager的handlers

```shell
[root@shkf6-245 src]# vi /data/dockerfile/tomcat8/apache-tomcat-8.5.50/conf/logging.properties

handlers = 1catalina.org.apache.juli.AsyncFileHandler, 2localhost.org.apache.juli.AsyncFileHandler, java.util.logging.ConsoleHandler
```

- 日志级别改为INFO

```shell
[root@shkf6-245 src]# vi /data/dockerfile/tomcat8/apache-tomcat-8.5.50/conf/logging.properties

1catalina.org.apache.juli.AsyncFileHandler.level = INFO
2localhost.org.apache.juli.AsyncFileHandler.level = INFO
java.util.logging.ConsoleHandler.level = INFO
```

- 注释掉所有关于3manager，4host-manager日志的配置

```shell
[root@shkf6-245 src]# vi /data/dockerfile/tomcat8/apache-tomcat-8.5.50/conf/logging.properties

#3manager.org.apache.juli.AsyncFileHandler.level = FINE
#3manager.org.apache.juli.AsyncFileHandler.directory = ${catalina.base}/logs
#3manager.org.apache.juli.AsyncFileHandler.prefix = manager.
#3manager.org.apache.juli.AsyncFileHandler.encoding = UTF-8

#4host-manager.org.apache.juli.AsyncFileHandler.level = FINE
#4host-manager.org.apache.juli.AsyncFileHandler.directory = ${catalina.base}/logs
#4host-manager.org.apache.juli.AsyncFileHandler.prefix = host-manager.
#4host-manager.org.apache.juli.AsyncFileHandler.encoding = UTF-8
```

### 3.准备Dockerfile

- Dockerfile

```shell
[root@shkf6-245 src]# cat /data/dockerfile/tomcat8/Dockerfile
From harbor.od.com/public/jre:8u112
RUN /bin/cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime &&\ 
    echo 'Asia/Shanghai' >/etc/timezone
ENV CATALINA_HOME /opt/tomcat
ENV LANG zh_CN.UTF-8
ADD apache-tomcat-8.5.50/ /opt/tomcat
ADD config.yml /opt/prom/config.yml
ADD jmx_javaagent-0.3.1.jar /opt/prom/jmx_javaagent-0.3.1.jar
WORKDIR /opt/tomcat
ADD entrypoint.sh /entrypoint.sh
CMD ["/entrypoint.sh"]
```

- config.yml

```shell
[root@shkf6-245 src]# cat /data/dockerfile/tomcat8/config.yml 
---
rules:
  - pattern: '.*'
```

- jmx_javaagent-0.3.1.jar

```shell
[root@shkf6-245 src]# wget -O /data/dockerfile/tomcat8/jmx_javaagent-0.3.1.jar https://repo1.maven.org/maven2/io/prometheus/jmx/jmx_prometheus_javaagent/0.3.1/jmx_prometheus_javaagent-0.3.1.jar -O jmx_javaagent-0.3.1.jar
```

- entrypoint.sh

```shell
[root@shkf6-245 src]# cat /data/dockerfile/tomcat8/entrypoint.sh 
#!/bin/bash
M_OPTS="-Duser.timezone=Asia/Shanghai -javaagent:/opt/prom/jmx_javaagent-0.3.1.jar=$(hostname -i):${M_PORT:-"12346"}:/opt/prom/config.yml"
C_OPTS=${C_OPTS}
MIN_HEAP=${MIN_HEAP:-"128m"}
MAX_HEAP=${MAX_HEAP:-"128m"}
JAVA_OPTS=${JAVA_OPTS:-"-Xmn384m -Xss256k -Duser.timezone=GMT+08  -XX:+DisableExplicitGC -XX:+UseConcMarkSweepGC -XX:+UseParNewGC -XX:+CMSParallelRemarkEnabled -XX:+UseCMSCompactAtFullCollection -XX:CMSFullGCsBeforeCompaction=0 -XX:+CMSClassUnloadingEnabled -XX:LargePageSizeInBytes=128m -XX:+UseFastAccessorMethods -XX:+UseCMSInitiatingOccupancyOnly -XX:CMSInitiatingOccupancyFraction=80 -XX:SoftRefLRUPolicyMSPerMB=0 -XX:+PrintClassHistogram  -Dfile.encoding=UTF8 -Dsun.jnu.encoding=UTF8"}
CATALINA_OPTS="${CATALINA_OPTS}"
JAVA_OPTS="${M_OPTS} ${C_OPTS} -Xms${MIN_HEAP} -Xmx${MAX_HEAP} ${JAVA_OPTS}"
sed -i -e "1a\JAVA_OPTS=\"$JAVA_OPTS\"" -e "1a\CATALINA_OPTS=\"$CATALINA_OPTS\"" /opt/tomcat/bin/catalina.sh

cd /opt/tomcat && /opt/tomcat/bin/catalina.sh run 2>&1 >> /opt/tomcat/logs/stdout.log

[root@shkf6-245 src]# chmod +x /data/dockerfile/tomcat8/entrypoint.sh 
```

### 4.制作镜像并推送

```shell
[root@shkf6-245 tomcat8]# ll
total 372
drwxr-xr-x 9 root root    220 Dec 27 10:36 apache-tomcat-8.5.50
-rw-r--r-- 1 root root     29 Dec 27 10:55 config.yml
-rw-r--r-- 1 root root    395 Dec 27 10:54 Dockerfile
-rwxr-xr-x 1 root root    988 Dec 27 10:57 entrypoint.sh
-rw-r--r-- 1 root root 367417 May 10  2018 jmx_javaagent-0.3.1.jar
[root@shkf6-245 tomcat8]# docker build . -t harbor.od.com/base/tomcat:v8.5.50
[root@shkf6-245 tomcat8]# docker push harbor.od.com/base/tomcat:v8.5.50
```

## 2.改造dubbo-demo-web项目

### 1.修改dubbo-client/pom.xml

```java
/d/workspace/dubbo-demo-web/dubbo-client/pom.xml

<packaging>war</packaging>

<dependency>
  <groupId>org.springframework.boot</groupId>
  <artifactId>spring-boot-starter-web</artifactId>
  <exclusions>
    <exclusion>
      <groupId>org.springframework.boot</groupId>
      <artifactId>spring-boot-starter-tomcat</artifactId>
    </exclusion>
  </exclusions>
</dependency>

<dependency>
  <groupId>org.apache.tomcat</groupId> 
  <artifactId>tomcat-servlet-api</artifactId>
  <version>8.0.36</version>
  <scope>provided</scope>
</dependency>
```

### 2.修改Application.java

```java
/d/workspace/dubbo-demo-web/dubbo-client/src/main/java/com/od/dubbotest/Application.java

import org.springframework.boot.autoconfigure.EnableAutoConfiguration;
import org.springframework.boot.autoconfigure.jdbc.DataSourceAutoConfiguration;
import org.springframework.context.annotation.ImportResource;

@ImportResource(value={"classpath*:spring-config.xml"})
@EnableAutoConfiguration(exclude={DataSourceAutoConfiguration.class})
```

### 3.创建ServletInitializer.java

```shell
/d/workspace/dubbo-demo-web/dubbo-client/src/main/java/com/od/dubbotest/ServletInitializer.java

package com.od.dubbotest;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.builder.SpringApplicationBuilder;
import org.springframework.boot.context.web.SpringBootServletInitializer;
import com.od.dubbotest.Application;

public class ServletInitializer extends SpringBootServletInitializer {

    @Override
    protected SpringApplicationBuilder configure(SpringApplicationBuilder builder) {
        return builder.sources(Application.class);
    }
}
```

## 3.新建Jenkins的pipeline

### 1.配置New job

- 使用admin登录
- New Item
- create new jobs
- Enter an item name

> tomcat-demo

- Pipeline -> OK
- Discard old builds

> Days to keep builds : 3
> Max # of builds to keep : 30

- This project is parameterized

1.Add Parameter -> String Parameter

> Name : app_name
> Default Value :
> Description : project name. e.g: dubbo-demo-web

2.Add Parameter -> String Parameter

> Name : image_name
> Default Value :
> Description : project docker image name. e.g: app/dubbo-demo-web

3.Add Parameter -> String Parameter

> Name : git_repo
> Default Value :
> Description : project git repository. e.g: [git@gitee.com](mailto:git@gitee.com):stanleywang/dubbo-demo-web.git

4.Add Parameter -> String Parameter

> Name : git_ver
> Default Value : tomcat
> Description : git commit id of the project.

5.Add Parameter -> String Parameter

> Name : add_tag
> Default Value :
> Description : project docker image tag, date_timestamp recommended. e.g: 190117_1920

6.Add Parameter -> String Parameter

> Name : mvn_dir
> Default Value : ./
> Description : project maven directory. e.g: ./

7.Add Parameter -> String Parameter

> Name : target_dir
> Default Value : ./dubbo-client/target
> Description : the relative path of target file such as .jar or .war package. e.g: ./dubbo-client/target

8.Add Parameter -> String Parameter

> Name : mvn_cmd
> Default Value : mvn clean package -Dmaven.test.skip=true
> Description : maven command. e.g: mvn clean package -e -q -Dmaven.test.skip=true

9.Add Parameter -> Choice Parameter

> Name : base_image
> Choices :
>
> - base/tomcat:v7.0.94
> - base/tomcat:v8.5.50
> - base/tomcat:v9.0.17
>
> Description : project base image list in harbor.od.com.

10.Add Parameter -> Choice Parameter

> Name : maven
> Choices :
>
> - 3.6.0-8u181
> - 3.2.5-6u025
> - 2.2.1-6u025
>
> Description : different maven edition.

11.Add Parameter -> String Parameter

> Name : root_url
> Default Value : ROOT
> Description : webapp dir.

### 2.Pipeline Script

```shell
pipeline {
  agent any 
    stages {
    stage('pull') { //get project code from repo 
      steps {
        sh "git clone ${params.git_repo} ${params.app_name}/${env.BUILD_NUMBER} && cd ${params.app_name}/${env.BUILD_NUMBER} && git checkout ${params.git_ver}"
        }
    }
    stage('build') { //exec mvn cmd
      steps {
        sh "cd ${params.app_name}/${env.BUILD_NUMBER}  && /var/jenkins_home/maven-${params.maven}/bin/${params.mvn_cmd}"
      }
    }
    stage('unzip') { //unzip  target/*.war -c target/project_dir
      steps {
        sh "cd ${params.app_name}/${env.BUILD_NUMBER} && cd ${params.target_dir} && mkdir project_dir && unzip *.war -d ./project_dir"
      }
    }
    stage('image') { //build image and push to registry
      steps {
        writeFile file: "${params.app_name}/${env.BUILD_NUMBER}/Dockerfile", text: """FROM harbor.od.com/${params.base_image}
ADD ${params.target_dir}/project_dir /opt/tomcat/webapps/${params.root_url}"""
        sh "cd  ${params.app_name}/${env.BUILD_NUMBER} && docker build -t harbor.od.com/${params.image_name}:${params.git_ver}_${params.add_tag} . && docker push harbor.od.com/${params.image_name}:${params.git_ver}_${params.add_tag}"
      }
    }
  }
}
```

## 4.构建应用镜像

使用Jenkins进行CI，并查看harbor仓库

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_beac99142dc6517d34db5bc6a78b2819_r.png)

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_b2e444afba57e56bade90d3975eb3b53_r.png)

## 5.准备资源配置清单

不再需要单独准备资源配置清单

## 6.应用资源配置清单

k8s的dashboard上直接修改image的值为jenkins打包出来的镜像
文档里的例子是：harbor.od.com:/app/dubbo-demo-web:tomcat_191227_1300

## 7.浏览器访问

http://demo-prod.od.com/hello?name=sunrise

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_f2fdb05aacf41fb0a47b60731a448055_r.png)

## 8.检查tomcat运行情况

任意一台运算节点主机上：

```shell
[root@shkf6-243 ~]# kubectl exec  dubbo-demo-consumer-65f9db9c8c-nhfvf ls logs/ -n prod 
catalina.2019-12-30.log
localhost.2019-12-30.log
localhost_access_log.2019-12-30.txt
stdout.log
```

# 第三章：实战安装部署ElasticSearch搜索引擎

## 1.部署ElasticSearch

[官网](https://www.elastic.co/)
[官方github地址](https://github.com/elastic/elasticsearch)
[下载地址](https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-6.8.6.tar.gz)

### 1.安装

```shell
[root@shkf6-242 opt]# cd src/
[root@shkf6-242 src]# wget https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-6.8.6.tar.gz
[root@shkf6-242 src]# ls -l|grep elasticsearch-6.8.6.tar.gz 
-rw-r--r-- 1 root root 149510158 Dec 19 00:05 elasticsearch-6.8.6.tar.gz
[root@shkf6-242 src]# tar xf elasticsearch-6.8.6.tar.gz  -C /opt
[root@shkf6-242 src]# ln -s /opt/elasticsearch-6.8.6/ /opt/elasticsearch
```

### 2.配置

#### 1.elasticsearch.yml

```shell
[root@shkf6-242 src]# mkdir -p /data/elasticsearch/{data,logs}

[root@shkf6-242 src]# vi /opt/elasticsearch/config/elasticsearch.yml
[root@shkf6-242 src]# grep -Ev "^#|^$" /opt/elasticsearch/config/elasticsearch.yml
cluster.name: es.od.com
node.name: shkf6-242.host.com
path.data: /data/elasticsearch/data
path.logs: /data/elasticsearch/logs
bootstrap.memory_lock: true
network.host: 192.168.6.242
http.port: 9200
```

#### 2.jvm.options

```shell
[root@shkf6-242 src]# grep -Ev "^#|^$" /opt/elasticsearch/config/jvm.options |grep Xm
-Xms512m
-Xmx512m
```

#### 3.创建普通用户

```shell
[root@shkf6-242 src]# cd /opt/elasticsearch
[root@shkf6-242 elasticsearch]# useradd -s /bin/bash -M es
[root@shkf6-242 elasticsearch]# chown -R es.es /opt/elasticsearch-6.8.6/
[root@shkf6-242 elasticsearch]# chown -R es.es /data/elasticsearch/
```

#### 4.文件描述符

```shell
[root@shkf6-242 elasticsearch]# vim /etc/security/limits.d/es.conf
[root@shkf6-242 elasticsearch]# cat /etc/security/limits.d/es.conf
es hard nofile 65536
es soft fsize unlimited
es hard memlock unlimited
es soft memlock unlimited
```

#### 5.调整内核参数

```shell
[root@shkf6-242 elasticsearch]# sysctl -w vm.max_map_count=262144
vm.max_map_count = 262144
[root@shkf6-242 elasticsearch]# echo "vm.max_map_count=262144" >> /etc/sysctl.conf
[root@shkf6-242 elasticsearch]# sysctl -p
vm.max_map_count = 262144
```

### 3.启动

方法一：

```shell
[root@shkf6-242 elasticsearch]# su -c "/opt/elasticsearch/bin/elasticsearch -d" es
```

方法二：

```shell
[root@shkf6-242 elasticsearch]# sudo -ues "/opt/elasticsearch/bin/elasticsearch -d"
```

检查：

```shell
[root@shkf6-242 elasticsearch]# netstat -lntup |grep 9200
tcp6       0      0 192.168.6.242:9200      :::*                    LISTEN      10726/java 
```

### 4.调整ES日志模板

```shell
[root@shkf6-242 elasticsearch]# curl -H "Content-Type:application/json" -XPUT http://192.168.6.242:9200/_template/k8s -d '{
  "template" : "k8s*",
  "index_patterns": ["k8s*"],  
  "settings": {
    "number_of_shards": 5,
    "number_of_replicas": 0
  }
}'
```

# 第四章：实战安装部署kafka消息队列及Kafka-manager

## 1.部署kafka

[官网](http://kafka.apache.org/)

[官方github地址](https://github.com/apache/kafka)

[下载地址](https://archive.apache.org/dist/kafka/2.2.0/kafka_2.12-2.2.0.tgz)

shkf6-241.host.com上：

### 1.安装

```shell
[root@shkf6-241 ~]# cd /opt/src/
[root@shkf6-241 src]# wget https://archive.apache.org/dist/kafka/2.2.0/kafka_2.12-2.2.0.tgz
[root@shkf6-241 src]# ls -l|grep kafka
-rw-r--r-- 1 root root  57028557 Mar 23  2019 kafka_2.12-2.2.0.tgz
[root@shkf6-241 src]# tar xf kafka_2.12-2.2.0.tgz -C /opt
[root@shkf6-241 src]# ln -s /opt/kafka_2.12-2.2.0/ /opt/kafka
```

### 2.配置

```shell
[root@shkf6-241 src]# vi /opt/kafka/config/server.properties
log.dirs=/data/kafka/logs
zookeeper.connect=localhost:2181
log.flush.interval.messages=10000
log.flush.interval.ms=1000
delete.topic.enable=true    # 行尾追加
host.name=shkf6-241.host.com  # 行尾追加
```

### 3.启动

```shell
[root@shkf6-241 src]# mkdir -p /data/kafka/logs
[root@shkf6-241 src]# cd /opt/kafka
[root@shkf6-241 kafka]# bin/kafka-server-start.sh -daemon config/server.properties
[root@shkf6-241 kafka]# netstat -luntp|grep 9092
tcp6       0      0 192.168.6.241:9092      :::*                    LISTEN      29424/java  
```

## 2.部署kafka-manager

[官方github地址](https://github.com/yahoo/kafka-manager)

[源码下载地址](https://github.com/yahoo/kafka-manager/archive/2.0.0.2.tar.gz)

运维主机shkf6-245.host.com

### 1.方法一：1、准备Dockerfile

```shell
[root@shkf6-245 ~]# mkdir /data/dockerfile/kafka-manager
[root@shkf6-245 ~]# vi /data/dockerfile/kafka-manager/Dockerfile
[root@shkf6-245 ~]# cat /data/dockerfile/kafka-manager/Dockerfile
FROM hseeberger/scala-sbt

ENV ZK_HOSTS=192.168.6.241:2181 \
     KM_VERSION=2.0.0.2

RUN mkdir -p /tmp && \
    cd /tmp && \
    wget https://github.com/yahoo/kafka-manager/archive/${KM_VERSION}.tar.gz && \
    tar xxf ${KM_VERSION}.tar.gz && \
    cd /tmp/kafka-manager-${KM_VERSION} && \
    sbt clean dist && \
    unzip  -d / ./target/universal/kafka-manager-${KM_VERSION}.zip && \
    rm -fr /tmp/${KM_VERSION} /tmp/kafka-manager-${KM_VERSION}

WORKDIR /kafka-manager-${KM_VERSION}

EXPOSE 9000
ENTRYPOINT ["./bin/kafka-manager","-Dconfig.file=conf/application.conf"]
```

### 2.方法一：2、制作docker镜像

```shell
[root@shkf6-245 ~]# cd /data/dockerfile/kafka-manager
[root@shkf6-245 kafka-manager]# docker build . -t harbor.od.com/infra/kafka-manager:v2.0.0.2

[root@shkf6-245 kafka-manager]# docker push harbor.od.com/infra/kafka-manager:v2.0.0.2
```

### 3.方法二：直接下载docker镜像

[镜像下载地址](https://hub.docker.com/r/sheepkiller/kafka-manager/tags)

```shell
[root@shkf6-245 ~]# docker pull sheepkiller/kafka-manager:stable
[root@shkf6-245 ~]# docker tag 34627743836f harbor.od.com/infra/kafka-manager:stable
[root@shkf6-245 ~]# docker push harbor.od.com/infra/kafka-manager:stable
docker push harbor.od.com/infra/kafka-manager:stable
The push refers to a repository [harbor.od.com/infra/kafka]
ef97dbc2670b: Pushed 
ec01aa005e59: Pushed 
de05a1bdf878: Pushed 
9c553f3feafd: Pushed 
581533427a4f: Pushed 
f6b229974fdd: Pushed 
ae150883d6e2: Pushed 
3df7be729841: Pushed 
f231cc200afe: Pushed 
9752c15164a8: Pushed 
9ab7eda5c826: Pushed 
402964b3d72e: Pushed 
6b3f8ebf864c: Pushed 
stable: digest: sha256:57fd46a3751284818f1bc6c0fdf097250bc0feed03e77135fb8b0a93aa8c6cc7 size: 3056
```

### 4.准备资源配置清单

- Deployment

```shell
[root@shkf6-245 ~]# vi /data/k8s-yaml/kafka-manager/dp.yaml
[root@shkf6-245 ~]# cat /data/k8s-yaml/kafka-manager/dp.yaml
kind: Deployment
apiVersion: extensions/v1beta1
metadata:
  name: kafka-manager
  namespace: infra
  labels: 
    name: kafka-manager
spec:
  replicas: 1
  selector:
    matchLabels: 
      app: kafka-manager
  strategy:
    type: RollingUpdate
    rollingUpdate: 
      maxUnavailable: 1
      maxSurge: 1
  revisionHistoryLimit: 7
  progressDeadlineSeconds: 600
  template:
    metadata:
      labels: 
        app: kafka-manager
    spec:
      containers:
      - name: kafka-manager
        image: harbor.od.com/infra/kafka-manager:v2.0.0.2
        imagePullPolicy: IfNotPresent
        ports:
        - containerPort: 9000
          protocol: TCP
        env:
        - name: ZK_HOSTS
          value: zk1.od.com:2181
        - name: APPLICATION_SECRET
          value: letmein
      imagePullSecrets:
      - name: harbor
      terminationGracePeriodSeconds: 30
      securityContext: 
        runAsUser: 0
```

- service

```shell
[root@shkf6-245 ~]# vi /data/k8s-yaml/kafka-manager/svc.yaml
[root@shkf6-245 ~]# cat /data/k8s-yaml/kafka-manager/svc.yaml
kind: Service
apiVersion: v1
metadata: 
  name: kafka-manager
  namespace: infra
spec:
  ports:
  - protocol: TCP
    port: 9000
    targetPort: 9000
  selector: 
    app: kafka-manager
```

- ingress

```shell
[root@shkf6-245 ~]# vi /data/k8s-yaml/kafka-manager/ingress.yaml
[root@shkf6-245 ~]# cat /data/k8s-yaml/kafka-manager/ingress.yaml
kind: Ingress
apiVersion: extensions/v1beta1
metadata: 
  name: kafka-manager
  namespace: infra
spec:
  rules:
  - host: km.od.com
    http:
      paths:
      - path: /
        backend: 
          serviceName: kafka-manager
          servicePort: 9000
```

### 5.应用资源配置清单

```shell
[root@shkf6-243 ~]# kubectl apply -f http://k8s-yaml.od.com/kafka-manager/dp.yaml
deployment.extensions/kafka-manager created
[root@shkf6-243 ~]# kubectl apply -f http://k8s-yaml.od.com/kafka-manager/svc.yaml
service/kafka-manager created
[root@shkf6-243 ~]# kubectl apply -f http://k8s-yaml.od.com/kafka-manager/ingress.yaml
ingress.extensions/kafka-manager created
```

### 6.解析域名

```shell
[root@shkf6-241 ~]# tail -1 /var/named/od.com.zone 
km                 A    192.168.6.66
```

### 7.浏览器访问

[http://km.od.com](http://km.od.com/)

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_a35ec6d3d6f3c237e38d99eeee9b8250_r.png)

- cluster –> Add Cluster

  > Cluster Name:kafka-od
  > Cluster Zookeepker Hosts:zk1.od.com:2181
  > kafka version:2.2.0

其他默认

- save

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_5e4f366b3437b647a0f8e794aa6a483e_r.png)

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_fe4d5282ebc64e05b14d52238c1e545c_r.png)

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_e1faffc62ca338b57812491c2992bb4c_r.png)

# 第五章：制作filebeat流式日志收集器docker镜像

[官方下载地址](https://www.elastic.co/downloads/beats/filebeat)

运维主机shkf6-245.host.com

## 1.制作docker镜像

### 1.准备dockerfile

```shell
[root@shkf6-245 ~]# mkdir /data/dockerfile/filebeat/
[root@shkf6-245 ~]# cd /data/dockerfile/filebeat/
[root@shkf6-245 filebeat]# cat Dockerfile 
FROM debian:jessie

ENV FILEBEAT_VERSION=7.4.0 \
    FILEBEAT_SHA1=c63bb1e16f7f85f71568041c78f11b57de58d497ba733e398fa4b2d071270a86dbab19d5cb35da5d3579f35cb5b5f3c46e6e08cdf840afb7c34

RUN set -x && \
  apt-get update && \
  apt-get install -y wget && \
  wget https://artifacts.elastic.co/downloads/beats/filebeat/filebeat-${FILEBEAT_VERSION}-linux-x86_64.tar.gz -O /opt/filebeat.tar.gz
  cd /opt && \
  echo "${FILEBEAT_SHA1}  filebeat.tar.gz" | sha512sum -c - && \
  tar xzvf filebeat.tar.gz && \
  cd filebeat-* && \
  cp filebeat /bin && \
  cd /opt && \
  rm -rf filebeat* && \
  apt-get purge -y wget && \
  apt-get autoremove -y && \
  apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

COPY docker-entrypoint.sh /
ENTRYPOINT ["/docker-entrypoint.sh"]
```

### 2.entrypoint.sh

```shell
[root@shkf6-245 filebeat]# cat docker-entrypoint.sh 
#!/bin/bash

ENV=${ENV:-"test"}
PROJ_NAME=${PROJ_NAME:-"no-define"}
MULTILINE=${MULTILINE:-"^\d{2}"}

cat > /etc/filebeat.yaml << EOF
filebeat.inputs:
- type: log
  fields_under_root: true
  fields:
    topic: logm-${PROJ_NAME}
  paths:
    - /logm/*.log
    - /logm/*/*.log
    - /logm/*/*/*.log
    - /logm/*/*/*/*.log
    - /logm/*/*/*/*/*.log
  scan_frequency: 120s
  max_bytes: 10485760
  multiline.pattern: '$MULTILINE'
  multiline.negate: true
  multiline.match: after
  multiline.max_lines: 100
- type: log
  fields_under_root: true
  fields:
    topic: logu-${PROJ_NAME}
  paths:
    - /logu/*.log
    - /logu/*/*.log
    - /logu/*/*/*.log
    - /logu/*/*/*/*.log
    - /logu/*/*/*/*/*.log
    - /logu/*/*/*/*/*/*.log
output.kafka:
  hosts: ["192.168.6.241:9092"]
  topic: k8s-fb-$ENV-%{[topic]}
  version: 2.0.0
  required_acks: 0
  max_message_bytes: 10485760
EOF

set -xe

# If user don't provide any command
# Run filebeat
if [[ "$1" == "" ]]; then
     exec filebeat  -c /etc/filebeat.yaml 
else
    # Else allow the user to run arbitrarily commands like bash
    exec "$@"
fi
[root@shkf6-245 filebeat]# chmod +x docker-entrypoint.sh
```

### 2.制作镜像

```shell
[root@shkf6-245 filebeat]# docker build . -t harbor.od.com/infra/filebeat:v7.4.0

[root@shkf6-245 filebeat]# docker push harbor.od.com/infra/filebeat:v7.4.0
```

# 第六章：实战K8S内应用接入filebeat并收集日志

## 1.修改资源配置清单

- 使用dubbo-demo-consumer的Tomcat版镜像

```shell
[root@shkf6-245 dubbo-demo-consumer]# cat /data/k8s-yaml/prod/dubbo-demo-consumer/dp-elk.yaml 
kind: Deployment
apiVersion: extensions/v1beta1
metadata:
  name: dubbo-demo-consumer
  namespace: prod
  labels: 
    name: dubbo-demo-consumer
spec:
  replicas: 1
  selector:
    matchLabels: 
      name: dubbo-demo-consumer
  template:
    metadata:
      labels: 
        app: dubbo-demo-consumer
        name: dubbo-demo-consumer
    spec:
      containers:
      - name: dubbo-demo-consumer
        image: harbor.od.com/app/dubbo-demo-web:tomcat_191227_1300
        ports:
        - containerPort: 8080
          protocol: TCP
        env:
        - name: C_OPTS
          value: -Denv=prod -Dapollo.meta=http://apollo-configservice:8080
        imagePullPolicy: IfNotPresent
        volumeMounts:
        - mountPath: /opt/tomcat/logs
          name: logm
      - name: filebeat
        image: harbor.od.com/infra/filebeat:v7.4.0
        imagePullPolicy: IfNotPresent
        env:
        - name: ENV
          value: prod
        - name: PROJ_NAME
          value: dubbo-demo-web
        volumeMounts:
        - mountPath: /logm
          name: logm
      volumes:
      - emptyDir: {}
        name: logm
      imagePullSecrets:
      - name: harbor
      restartPolicy: Always
      terminationGracePeriodSeconds: 30
      securityContext: 
        runAsUser: 0
      schedulerName: default-scheduler
  strategy:
    type: RollingUpdate
    rollingUpdate: 
      maxUnavailable: 1
      maxSurge: 1
  revisionHistoryLimit: 7
  progressDeadlineSeconds: 600
```

- 应用资源配置清单

```shell
[root@shkf6-243 ~]# kubectl apply -f http://k8s-yaml.od.com/prod/dubbo-demo-consumer/dp-elk.yaml
deployment.extensions/dubbo-demo-consumer configured
```

- 插看日志

```shell
[root@shkf6-243 ~]# kubectl exec -it dubbo-demo-consumer-5698bbd44d-rkpr6 bash -n prod 
Defaulting container name to dubbo-demo-consumer.
Use 'kubectl describe pod/dubbo-demo-consumer-5698bbd44d-rkpr6 -n prod' to see all of the containers in this pod.
root@dubbo-demo-consumer-5698bbd44d-rkpr6:/opt/tomcat# tailf logs/stdout.log 
2020-01-02 11:24:42 HelloAction接收到请求:sunrise
2020-01-02 11:24:42 HelloService返回到结果:2020-01-02 11:24:42 <h1>这是Dubbo 消费者端(Tomcat服务)</h1><h2>欢迎来到老男孩教育K8S容器云架构师专题课培训班1期。</h2>hello sunrise
2020-01-02 11:24:42 HelloAction接收到请求:sunrise
```

## 2.浏览器访问[http://km.od.com](http://km.od.com/)

看到kafaka-manager里，topic打进来，即为成功。

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_fd1d3d0af4047413ed874cff5730aa4a_r.png)

# 第七章：实战部署Logstash及Kibana

## 1.部署logstash

运维主机shkf6-245.host.com上：

### 1.选版本

[logstash选型](https://www.elastic.co/support/matrix)

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_86389684f1b91d0da799d981fcb186ac_r.png)

## 2.准备docker镜像

- 下载官方镜像

```shell
[root@shkf6-245 ~]# docker pull logstash:6.8.6
[root@shkf6-245 ~]# docker images|grep logstash
logstash                                          6.8.6                      d0a2dac51fcb        2 weeks ago         827MB
[root@shkf6-245 ~]# docker tag d0a2dac51fcb harbor.od.com/infra/logstash:v6.8.6
[root@shkf6-245 ~]# docker push harbor.od.com/infra/logstash:v6.8.6
```

### 2.启动docker镜像

在运维主机shkf6-245.host.com

- 创建配置

```shell
[root@shkf6-245 ~]# mkdir /etc/logstash/
[root@shkf6-245 ~]# vi /etc/logstash/logstash-prod.conf
[root@shkf6-245 ~]# cat /etc/logstash/logstash-prod.conf
input {
  kafka {
    bootstrap_servers => "192.168.6.241:9092"
    client_id => "192.168.6.245"
    consumer_threads => 4
    group_id => "k8s_prod"
    topics_pattern => "k8s-fb-prod-.*"
  }
}

filter {
  json {
    source => "message"
  }
}

output {
  elasticsearch {
    hosts => ["192.168.6.242:9200"]
    index => "k8s-prod-%{+YYYY.MM.DD}"
  }
}
```

- 启动logstash镜像

```shell
[root@shkf6-245 ~]# docker run -d --name logstash-test -v /etc/logstash:/etc/logstash harbor.od.com/infra/logstash:v6.8.6 -f /etc/logstash/logstash-test.conf

[root@shkf6-245 ~]# docker run -d --name logstash-prod -v /etc/logstash:/etc/logstash harbor.od.com/infra/logstash:v6.8.6 -f /etc/logstash/logstash-prod.conf
9e07358ed7b536d5874c85e0c2ed5d9e4004382473fe695155b7183a532539b2
[root@shkf6-245 ~]# docker ps -a|grep logstash
9e07358ed7b5        harbor.od.com/infra/logstash:v6.8.6                 "/usr/local/bin/dock…"   11 seconds ago      Up 8 seconds           5044/tcp, 9600/tcp          logstash-prod
```

- 验证ElasticSearch里的索引（等一分钟）

```shell
[root@shkf6-245 ~]# curl http://192.168.6.242:9200/_cat/indices?v
health status index               uuid                   pri rep docs.count docs.deleted store.size pri.store.size
green  open   k8s-prod-2020.01.02 5atOPCdhSDa5b7ovVTyfgA   5   0         70            0    173.1kb        173.1kb
```

## 2.部署Kibana

运维主机shkf6-245.host.com上：

### 1.准备docker镜像

[kibana官方镜像下载地址](https://hub.docker.com/_/kibana?tab=tags)

```shell
[root@shkf6-245 ~]# docker pull kibana:6.8.6
[root@shkf6-245 ~]# docker images|grep kibana
kibana                                            6.8.6                      adfab5632ef4        2 weeks ago         739MB
[root@shkf6-245 ~]# docker tag adfab5632ef4  harbor.od.com/infra/kibana:v6.8.6
[root@shkf6-245 ~]# docker push harbor.od.com/infra/kibana:v6.8.6
```

### 2.准备资源配置清单

- Deployment

```shell
[root@shkf6-245 ~]# cat /data/k8s-yaml/kibana/dp.yaml
kind: Deployment
apiVersion: extensions/v1beta1
metadata:
  name: kibana
  namespace: infra
  labels: 
    name: kibana
spec:
  replicas: 1
  selector:
    matchLabels: 
      name: kibana
  template:
    metadata:
      labels: 
        app: kibana
        name: kibana
    spec:
      containers:
      - name: kibana
        image: harbor.od.com/infra/kibana:v6.8.6
        imagePullPolicy: IfNotPresent
        ports:
        - containerPort: 5601
          protocol: TCP
        env:
        - name: ELASTICSEARCH_URL
          value: http://192.168.6.242:9200
      imagePullSecrets:
      - name: harbor
      securityContext: 
        runAsUser: 0
  strategy:
    type: RollingUpdate
    rollingUpdate: 
      maxUnavailable: 1
      maxSurge: 1
  revisionHistoryLimit: 7
  progressDeadlineSeconds: 600
```

- Service

```shell
[root@shkf6-245 ~]# cat /data/k8s-yaml/kibana/svc.yaml
kind: Service
apiVersion: v1
metadata: 
  name: kibana
  namespace: infra
spec:
  ports:
  - protocol: TCP
    port: 5601
    targetPort: 5601
  selector: 
    app: kibana
```

- Ingress

```shell
[root@shkf6-245 ~]# cat /data/k8s-yaml/kibana/ingress.yaml 
kind: Ingress
apiVersion: extensions/v1beta1
metadata: 
  name: kibana
  namespace: infra
spec:
  rules:
  - host: kibana.od.com
    http:
      paths:
      - path: /
        backend: 
          serviceName: kibana
          servicePort: 5601
```

### 3.应用资源配置清单

```shell
[root@shkf6-243 ~]# kubectl apply -f http://k8s-yaml.od.com/kibana/dp.yaml
deployment.extensions/kibana created
[root@shkf6-243 ~]# kubectl apply -f http://k8s-yaml.od.com/kibana/svc.yaml
service/kibana created
[root@shkf6-243 ~]# kubectl apply -f http://k8s-yaml.od.com/kibana/ingress.yaml
ingress.extensions/kibana created
```

### 4.解析域名

```shell
[root@shkf6-241 ~]# tail -1 /var/named/od.com.zone 
kibana             A    192.168.6.66
```

### 5.浏览器访问

http://kibana.od.com/

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_b0525168bc3a2b282bb9a7215c7ab43b_r.png)

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_130c237ac262918aa7758e1fe5a862de_r.png)

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_09396dae4287b642781368b550a1bfcd_r.png)

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_cee7c1626c99c4e05df64d549156e283_r.png)

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_aa39245323c07b9eab2547568bf124cc_r.png)

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_0d0b3bf1d0e1862cd65d4e802290213f_r.png)

**选择区域**

- [@timestamp](https://github.com/timestamp)

> 对应日志的时间戳

- log.file.path

> 对应日志文件名

- message

> 对应日志内容

**时间选择器**

- 选择日志时间
  - 快速时间
  - 绝对时间
  - 相对时间

**环境选择器**

- 选择对应环境的日志

  > k8s-test-
  > k8s-prod-

**项目选择器**

- 对应filebeat的PROJ_NAME值

- Add a fillter

- topic is ${PROJ_NAME}

  > dubbo-demo-service
  > dubbo-demo-web

**关键字选择器**

- exception
- error

### 6.添加测试环境

这里把prod环境做好了，test环境也是类似的方法

1.修改test环境tomcat的dp

```shell
[root@shkf6-245 dubbo-demo-consumer]# pwd
/data/k8s-yaml/test/dubbo-demo-consumer
[root@shkf6-245 dubbo-demo-consumer]# cat dp-elk.yaml 
kind: Deployment
apiVersion: extensions/v1beta1
metadata:
  name: dubbo-demo-consumer
  namespace: test 
  labels: 
    name: dubbo-demo-consumer
spec:
  replicas: 1
  selector:
    matchLabels: 
      name: dubbo-demo-consumer
  template:
    metadata:
      labels: 
        app: dubbo-demo-consumer
        name: dubbo-demo-consumer
    spec:
      containers:
      - name: dubbo-demo-consumer
        image: harbor.od.com/app/dubbo-demo-web:tomcat_191227_1300
        ports:
        - containerPort: 8080
          protocol: TCP
        env:
        - name: C_OPTS
          value: -Denv=fat -Dapollo.meta=http://apollo-configservice:8080
        imagePullPolicy: IfNotPresent
        volumeMounts:
        - mountPath: /opt/tomcat/logs
          name: logm
      - name: filebeat
        image: harbor.od.com/infra/filebeat:v7.4.0
        imagePullPolicy: IfNotPresent
        env:
        - name: ENV
          value: test
        - name: PROJ_NAME
          value: dubbo-demo-web
        volumeMounts:
        - mountPath: /logm
          name: logm
      volumes:
      - emptyDir: {}
        name: logm
      imagePullSecrets:
      - name: harbor
      restartPolicy: Always
      terminationGracePeriodSeconds: 30
      securityContext: 
        runAsUser: 0
      schedulerName: default-scheduler
  strategy:
    type: RollingUpdate
    rollingUpdate: 
      maxUnavailable: 1
      maxSurge: 1
  revisionHistoryLimit: 7
  progressDeadlineSeconds: 600
```

2.logstash配置

```shell
[root@shkf6-245 ~]# cat /etc/logstash/logstash-test.conf 
input {
  kafka {
    bootstrap_servers => "192.168.6.241:9092"
    client_id => "192.168.6.245"
    consumer_threads => 4
    group_id => "k8s_test"
    topics_pattern => "k8s-fb-test-.*"
  }
}

filter {
  json {
    source => "message"
  }
}

output {
  elasticsearch {
    hosts => ["192.168.6.242:9200"]
    index => "k8s-test-%{+YYYY.MM.DD}"
  }
```

3.略