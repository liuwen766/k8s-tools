# 1.前情回顾

- Dubbo微服务
  - 注册中心zookeeper（集群）
  - 提供者（集群）
  - 消费者（集群）
  - 监控（dubbo-monitor/dubbo-admin）
- 在K8s内交付dubbo微服务的步骤：
  - step0：有可用的k8s集群
  - step1：部署zk集群（通常放在K8S集群外，有状态）
  - step2:部署jenkins（以容器的形式交付在K8S集群内）
    - root、时区、ssh-key、docker客户端、harbor连接配置
  - step3:部署maven软件
  - step4:制作dubbo微服务底包
  - step5：配置jenkins持续构建（CI）流水线
  - step6：使用流水线构建项目，查看harbor仓库
  - step7：使用资源配置清单，交付项目到K8S集群
- 交付dubbo-monitor
- 运维八荣八耻
  - 以可配置为荣，以硬编码为耻
  - 以互备为荣，以单点为耻
  - 以随时重启为荣，以不能迁移为耻
  - 以整体交付为荣，以部分交付为耻
  - 以无状态为荣，以有状态为耻
  - 以标准化为荣，以特殊化为耻
  - 以自动化工具为荣，以手工+人肉为耻
  - 以无人值守为荣，以人工介入为耻
- 考虑我们交付进K8S集群的两个dubbo微服务和一个monitor，最大的问题是什么？
  - 他们的配置写死在容器里了！

# 第一章：配置中心的概述

- 配置其实是独立于程序的可配置变量，同一份程序在不同配置下会有不同的行为，常见的配置有连接字符串，应用配置和业务配置等。
- 配置有多种形态，下面是一些常见的：
  - 程序内部hardcode，这种做法是反模式，一般我们不建议！
  - 配置文件，比如spring应用程序的配置一般放在application.properties文件中。
  - 环境变量，配置可以预置在操作系统的环境变量里头，程序运行时读取。
  - 启动参数，可以预置在操作系统的环境变量里头，程序运行时读取。
  - 启动参数，可以在程序启动时一次性提供参数，例如java程序启动时可以通过java -D 方式配启动参数。
  - 基于数据库，有经验的开发人员把易变配置放在数据库中，这样可以在运行期灵活调整配置。

- 配置管理的现状：
  - 配置散乱格式不标准（xml、ini、conf、yaml…）
  - 主要采用本地静态配置，应用多副本下配置修改麻烦
  - 易引发生产事故（测试环境。生产环境配置混用）
  - 配置缺乏安全审计和版本控制功能（config review）
  - 不同环境的应用，配置不同，造成多次打包，测试失效
- 配置中心是什么？
  - 顾名思义，就是集中管理应用程序配置的“中心”。
- 常见的配置中心有：
  - XDiamond：全局配置中心，存储应用的配置项，解决配置换乱分散的问题。名字来源于淘宝的开源项目diamond，前面加一个字母X以示区别。
  - Qconf：Qconf是一个分布式配置管理工具。用来替代传统的配置文件，使得配置信息和程序代码分离，同时配置变化能够实时同步到客户端，而且保证用户高效读取配置，这使得工程师从琐碎的配置修改、代码提交、配置上线流程中解放出来，极大地简化了配置管理工作。
  - Disconf：专注于各种【分布式系统配置管理】的【通用组件】和【通用平台】，提供统一的【配置管理服务】
  - SpringCloudConfig：Spring Cloud Config为分布式系统中的外部配置提供服务器和客户端支持。
  - K8S ConfigMap：K8S的一种标准资源，专门用来集中管理应用的配置。
  - Apollo：携程框架部门开源的，分布式配置中心。

# 第二章：实战K8S配置中心-ConfigMap

## 1.使用ConfigMap管理应用配置

### 1.拆分环境

| 主机名             | 角色                 | ip            |
| :----------------- | :------------------- | :------------ |
| shkf6-241.host.com | zk1.od.com(Test环境) | 192.168.6.241 |
| shkf6-242.host.com | zk2.od.com(Prod环境) | 192.168.6.242 |

### 2.重配zookeeper

在shkf6-241.host.com主机上：

```shell
[root@shkf6-241 ~]# vi /opt/zookeeper/conf/zoo.cfg 
[root@shkf6-241 ~]# cat /opt/zookeeper/conf/zoo.cfg
tickTime=2000
initLimit=10
syncLimit=5
dataDir=/data/zookeeper/data
dataLogDir=/data/zookeeper/logs
clientPort=2181
```

在shkf6-242.host.com主机上：

```shell
[root@shkf6-242 ~]# vi /opt/zookeeper/conf/zoo.cfg 
[root@shkf6-242 ~]# cat /opt/zookeeper/conf/zoo.cfg
tickTime=2000
initLimit=10
syncLimit=5
dataDir=/data/zookeeper/data
dataLogDir=/data/zookeeper/logs
clientPort=2181
```

重启zk(删除数据文件)

```shell
[root@shkf6-241 ~]# rm -rf /data/zookeeper/data/* && rm -rf /data/zookeeper/logs/*
[root@shkf6-242 ~]# rm -rf /data/zookeeper/data/* && rm -rf /data/zookeeper/logs/*

[root@shkf6-241 ~]# /opt/zookeeper/bin/zkServer.sh restart && /opt/zookeeper/bin/zkServer.sh status
[root@shkf6-242 ~]# /opt/zookeeper/bin/zkServer.sh restart && /opt/zookeeper/bin/zkServer.sh status
[root@shkf6-243 ~]# /opt/zookeeper/bin/zkServer.sh stop
```

### 3.准备资源配置清单(dubbo-monitor)

在运维主机shkf6-245.host.com上：

- configmap

```shell
[root@shkf6-245 dubbo-monitor-cm]# cat /data/k8s-yaml/dubbo-monitor-cm/cm.yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: dubbo-monitor-cm
  namespace: infra
data:
  dubbo.properties: |
    dubbo.container=log4j,spring,registry,jetty
    dubbo.application.name=simple-monitor
    dubbo.application.owner=
    dubbo.registry.address=zookeeper://zk1.od.com:2181
    dubbo.protocol.port=20880
    dubbo.jetty.port=8080
    dubbo.jetty.directory=/dubbo-monitor-simple/monitor
    dubbo.charts.directory=/dubbo-monitor-simple/charts
    dubbo.statistics.directory=/dubbo-monitor-simple/statistics
    dubbo.log4j.file=/dubbo-monitor-simple/logs/dubbo-monitor.log
    dubbo.log4j.level=WARN
```

- deployment

```shell
[root@shkf6-245 dubbo-monitor-cm]# cat /data/k8s-yaml/dubbo-monitor-cm/dp.yaml
kind: Deployment
apiVersion: extensions/v1beta1
metadata:
  name: dubbo-monitor
  namespace: infra
  labels: 
    name: dubbo-monitor
spec:
  replicas: 1
  selector:
    matchLabels: 
      name: dubbo-monitor
  template:
    metadata:
      labels: 
        app: dubbo-monitor
        name: dubbo-monitor
    spec:
      containers:
      - name: dubbo-monitor
        image: harbor.od.com/infra/dubbo-monitor:latest
        ports:
        - containerPort: 8080
          protocol: TCP
        - containerPort: 20880
          protocol: TCP
        imagePullPolicy: IfNotPresent
        volumeMounts:
          - name: configmap-volume
            mountPath: /dubbo-monitor-simple/conf
      volumes:
        - name: configmap-volume
          configMap:
            name: dubbo-monitor-cm
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

### 4.应用资源配置清单

在任意一台k8s运算节点执行：

```shell
[root@shkf6-243 ~]# kubectl apply -f http://k8s-yaml.od.com/dubbo-monitor-cm/cm.yaml
configmap/dubbo-monitor-cm created
[root@shkf6-243 ~]# kubectl apply -f http://k8s-yaml.od.com/dubbo-monitor-cm/dp.yaml
deployment.extensions/dubbo-monitor configured
```

### 5.重新发版，修改dubbo项目的配置文件

##### 1.修改项目源代码

- duboo-demo-service

```java
dubbo-server/src/main/java/config.properties   #代码中的路径

dubbo.registry=zookeeper://zk1.od.com:2181
dubbo.port=28080
```

- dubbo-demo-web

```java
dubbo-client/src/main/java/config.properties  #代码中的路径

dubbo.registry=zookeeper://zk1.od.com:2181
```

##### 2.使用Jenkins进行CI

略

##### 3.修改/应用资源配置清单

k8s的dashboard上，修改deployment使用的容器版本，提交应用

### 6.验证configmap的配置

在K8S的dashboard上，修改dubbo-monitor的configmap配置为不同的zk，重启POD，浏览器打开[http://dubbo-monitor.od.com](http://dubbo-monitor.od.com/) 观察效果

### 7.configmap 复杂配置文件处理方法

```shell
[root@shkf6-243 ~]# cd /opt/kubernetes/server/bin/conf/
[root@shkf6-243 conf]# ll
total 24
-rw-r--r-- 1 root root 2223 Nov 27 16:46 audit.yaml
-rw-r--r-- 1 root root  258 Nov 27 16:46 k8s-node.yaml
-rw------- 1 root root 6198 Nov 27 16:46 kubelet.kubeconfig   # 把这个配置做成configmap
-rw------- 1 root root 6218 Nov 27 16:46 kube-proxy.kubeconfig

[root@shkf6-243 conf]# kubectl create cm kubelet-configmap --from-file=./kubelet.kubeconfig 
configmap/kubelet-configmap created

[root@shkf6-243 conf]# kubectl get cm kubelet-configmap -o yaml
```

# 第三章：apollo配置中心介绍

**Apollo（阿波罗）是携程框架部门研发的分布式配置中心，能够集中化管理应用不同环境、不同集群的配置，配置修改后能够实时推送到应用端，并且具备规范的权限、流程治理等特性，适用于微服务配置管理场景。**

官方GitHub地址：
[Apollo官方地址](https://github.com/ctripcorp/apollo)

[官方release包](https://github.com/ctripcorp/apollo/releases)

- 基础架构

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_7019e15a1f6db2ff97430f6676e1ddfd_r.png)

> 1. Cconfig Service提供配置的读取、推送等功能，服务对象是Apollo客户端
> 2. Admin Service提供配置的修改、发布等功能，服务对象是Apollo Portal（管理界面）
> 3. Config Service和Admin Server都是多实例、无状态部署，所以需要将自己注册到Eureka中并保持心跳
> 4. 在Eureka之上我们架了一层Meta Server用于封装Eureka的服务发现接口
> 5. Client通过域名访问Meta Server获取Config Server服务列表（IP+Port），而后直接通过IP+Port访问服务，同时在Client测绘做load balance、错误重试
> 6. Portal通过域名访问Meta Server获取Admin Service服务列表（IP + port），而后直接通过IP+Port访问服务，同时在Portal测绘做load balance、错误重试

- 简化模型

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_df1ee811cd18d04e4a6acc9e3432fc47_r.png)

# 第四章：实战交付apollo配置中心组件–configservice到kubernetes集群

## 1.准备软件包

在运维主机shkf6-245.host.com上：

[下载官方release包](https://github.com/ctripcorp/apollo/releases/download/v1.5.1/apollo-configservice-1.5.1-github.zip)

```shell
[root@shkf6-245 ~]# wget -O /opt/src/apollo-configservice-1.5.1-github.zip https://github.com/ctripcorp/apollo/releases/download/v1.5.1/apollo-configservice-1.5.1-github.zip
[root@shkf6-245 ~]# mkdir /data/dockerfile/apollo-configservice && unzip -o /opt/src/apollo-configservice-1.5.1-github.zip -d /data/dockerfile/apollo-configservice
```

## 2.安装数据库

在数据库主机shkf6-241.host.com上：

注意：MySQL版本应为5.6或以上！

- 更新yum源

```shell
[root@shkf6-241 ~]# cat /etc/yum.repos.d/MariaDB.repo
[mariadb]
name = MariaDB
baseurl = https://mirrors.ustc.edu.cn/mariadb/yum/10.1/centos7-amd64/
gpgkey=https://mirrors.ustc.edu.cn/mariadb/yum/RPM-GPG-KEY-MariaDB
gpgcheck=1
```

- 导入GPG-KEY

```shell
[root@shkf6-241 ~]# rpm --import https://mirrors.ustc.edu.cn/mariadb/yum/RPM-GPG-KEY-MariaDB
```

- 安装数据库版本

```shell
[root@shkf6-241 ~]# yum list mariadb-server --show-duplicates
Loaded plugins: fastestmirror
Loading mirror speeds from cached hostfile
 * base: mirrors.163.com
 * epel: my.mirrors.thegigabit.com
 * extras: mirrors.huaweicloud.com
 * updates: mirrors.huaweicloud.com
Available Packages
MariaDB-server.x86_64                                          10.1.40-1.el7.centos                                           mariadb
MariaDB-server.x86_64                                          10.1.41-1.el7.centos                                           mariadb
MariaDB-server.x86_64                                          10.1.43-1.el7.centos                                           mariadb
mariadb-server.x86_64                                          1:5.5.64-1.el7                                                 base

[root@shkf6-241 ~]# yum install MariaDB-server -y
```

- 配置数据库字符集

```shell
[root@shkf6-241 ~]# grep utf8mb4 /etc/my.cnf.d/mysql-clients.cnf
default-character-set = utf8mb4
[root@shkf6-241 ~]# grep utf8mb4 /etc/my.cnf.d/server.cnf 
character_set_server = utf8mb4
collation_server = utf8mb4_general_ci
init_connect = "SET NAMES 'utf8mb4'"
```

示例：

```shell
[mysql]
default-character-set = utf8mb4
[mysqld]
character_set_server = utf8mb4
collation_server = utf8mb4_general_ci
init_connect = "SET NAMES 'utf8mb4'"
```

- 启动数据库

  ```shell
  [root@shkf6-241 ~]# systemctl start mariadb.service 
  [root@shkf6-241 ~]# netstat -lntup|grep 3306
  tcp6       0      0 :::3306                 :::*                    LISTEN      14887/mysqld 
  [root@shkf6-241 ~]# systemctl enable mariadb.service
  ```

- 设置数据管理员密码

```shell
[root@shkf6-241 ~]# mysqladmin -u root password
New password: （123456）
Confirm new password: （123456）
```

- 检查数据库配置

```shell
[root@shkf6-241 ~]# mysql -uroot -p
Enter password: 
Welcome to the MariaDB monitor.  Commands end with ; or \g.
Your MariaDB connection id is 3
Server version: 10.1.43-MariaDB MariaDB Server

Copyright (c) 2000, 2018, Oracle, MariaDB Corporation Ab and others.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

MariaDB [(none)]> \s
--------------
mysql  Ver 15.1 Distrib 10.1.43-MariaDB, for Linux (x86_64) using readline 5.1
......
Server characterset:    utf8mb4
Db     characterset:    utf8mb4
Client characterset:    utf8mb4
Conn.  characterset:    utf8mb4
UNIX socket:        /var/lib/mysql/mysql.sock
Uptime:            3 min 51 sec

MariaDB [(none)]> drop database test;
Query OK, 0 rows affected (0.00 sec)

MariaDB [(none)]> show databases;
+--------------------+
| Database           |
+--------------------+
| information_schema |
| mysql              |
| performance_schema |
+--------------------+
```

## 3.执行数据库脚本

[数据库脚本地址](https://raw.githubusercontent.com/ctripcorp/apollo/1.5.1/scripts/db/migration/configdb/V1.0.0__initialization.sql)

```shell
[root@shkf6-241 ~]# mysql -uroot -p < apolloconfig.sql 
Enter password: （123456）
```

## 4.数据库用户授权

```sql
MariaDB [ApolloConfigDB]> grant INSERT,DELETE,UPDATE,SELECT on ApolloConfigDB.* to "apolloconfig"@"192.168.6.%" identified by "123456";
```

## 5.修改初始数据

```sql
MariaDB [ApolloConfigDB]> update ApolloConfigDB.ServerConfig set ServerConfig.Value="http://config.od.com/eureka" where ServerConfig.Key="eureka.service.url";

MariaDB [ApolloConfigDB]> select * from ServerConfig\G
*************************** 1. row ***************************
                       Id: 1
                      Key: eureka.service.url
                  Cluster: default
                    Value: http://config.od.com/eureka
                  Comment: Eureka服务Url，多个service以英文逗号分隔
                IsDeleted:  
     DataChange_CreatedBy: default
   DataChange_CreatedTime: 2019-12-10 16:08:44
DataChange_LastModifiedBy: 
      DataChange_LastTime: 2019-12-10 16:13:53
```

## 6.制作docker镜像

在运维主机shkf6-245.host.com上：

- 删除无用的数据

```shell
[root@shkf6-245 ~]# cd /data/dockerfile/apollo-configservice/
[root@shkf6-245 apollo-configservice]# rm -f apollo-configservice-1.5.1-sources.jar
[root@shkf6-245 apollo-configservice]# rm -f scripts/shutdown.sh
```

- 配置数据库连接串（这里可以不改，可以用cm挂载）

```shell
[root@shkf6-245 apollo-configservice]# cat config/application-github.properties
# DataSource
spring.datasource.url = jdbc:mysql://mysql.od.com:3306/ApolloConfigDB?characterEncoding=utf8
spring.datasource.username = apolloconfig
spring.datasource.password = 123456


#apollo.eureka.server.enabled=true
#apollo.eureka.client.enabled=true
```

- 配置启动脚本

[官方提供的k8s启动脚本](https://raw.githubusercontent.com/ctripcorp/apollo/1.5.1/scripts/apollo-on-kubernetes/apollo-config-server/scripts/startup-kubernetes.sh)

```shell
[root@shkf6-245 apollo-configservice]# cat scripts/startup.sh
#!/bin/bash
SERVICE_NAME=apollo-configservice
## Adjust log dir if necessary
LOG_DIR=/opt/logs/apollo-config-server
## Adjust server port if necessary
SERVER_PORT=8080
APOLLO_CONFIG_SERVICE_NAME=$(hostname -i)

SERVER_URL="http://${APOLLO_CONFIG_SERVICE_NAME}:${SERVER_PORT}"

## Adjust memory settings if necessary
#export JAVA_OPTS="-Xms6144m -Xmx6144m -Xss256k -XX:MetaspaceSize=128m -XX:MaxMetaspaceSize=384m -XX:NewSize=4096m -XX:MaxNewSize=4096m -XX:SurvivorRatio=8"

## Only uncomment the following when you are using server jvm
#export JAVA_OPTS="$JAVA_OPTS -server -XX:-ReduceInitialCardMarks"

########### The following is the same for configservice, adminservice, portal ###########
export JAVA_OPTS="$JAVA_OPTS -XX:ParallelGCThreads=4 -XX:MaxTenuringThreshold=9 -XX:+DisableExplicitGC -XX:+ScavengeBeforeFullGC -XX:SoftRefLRUPolicyMSPerMB=0 -XX:+ExplicitGCInvokesConcurrent -XX:+HeapDumpOnOutOfMemoryError -XX:-OmitStackTraceInFastThrow -Duser.timezone=Asia/Shanghai -Dclient.encoding.override=UTF-8 -Dfile.encoding=UTF-8 -Djava.security.egd=file:/dev/./urandom"
export JAVA_OPTS="$JAVA_OPTS -Dserver.port=$SERVER_PORT -Dlogging.file=$LOG_DIR/$SERVICE_NAME.log -XX:HeapDumpPath=$LOG_DIR/HeapDumpOnOutOfMemoryError/"

# Find Java
if [[ -n "$JAVA_HOME" ]] && [[ -x "$JAVA_HOME/bin/java" ]]; then
    javaexe="$JAVA_HOME/bin/java"
elif type -p java > /dev/null 2>&1; then
    javaexe=$(type -p java)
elif [[ -x "/usr/bin/java" ]];  then
    javaexe="/usr/bin/java"
else
    echo "Unable to find Java"
    exit 1
fi

if [[ "$javaexe" ]]; then
    version=$("$javaexe" -version 2>&1 | awk -F '"' '/version/ {print $2}')
    version=$(echo "$version" | awk -F. '{printf("%03d%03d",$1,$2);}')
    # now version is of format 009003 (9.3.x)
    if [ $version -ge 011000 ]; then
        JAVA_OPTS="$JAVA_OPTS -Xlog:gc*:$LOG_DIR/gc.log:time,level,tags -Xlog:safepoint -Xlog:gc+heap=trace"
    elif [ $version -ge 010000 ]; then
        JAVA_OPTS="$JAVA_OPTS -Xlog:gc*:$LOG_DIR/gc.log:time,level,tags -Xlog:safepoint -Xlog:gc+heap=trace"
    elif [ $version -ge 009000 ]; then
        JAVA_OPTS="$JAVA_OPTS -Xlog:gc*:$LOG_DIR/gc.log:time,level,tags -Xlog:safepoint -Xlog:gc+heap=trace"
    else
        JAVA_OPTS="$JAVA_OPTS -XX:+UseParNewGC"
        JAVA_OPTS="$JAVA_OPTS -Xloggc:$LOG_DIR/gc.log -XX:+PrintGCDetails"
        JAVA_OPTS="$JAVA_OPTS -XX:+UseConcMarkSweepGC -XX:+UseCMSCompactAtFullCollection -XX:+UseCMSInitiatingOccupancyOnly -XX:CMSInitiatingOccupancyFraction=60 -XX:+CMSClassUnloadingEnabled -XX:+CMSParallelRemarkEnabled -XX:CMSFullGCsBeforeCompaction=9 -XX:+CMSClassUnloadingEnabled  -XX:+PrintGCDateStamps -XX:+PrintGCApplicationConcurrentTime -XX:+PrintHeapAtGC -XX:+UseGCLogFileRotation -XX:NumberOfGCLogFiles=5 -XX:GCLogFileSize=5M"
    fi
fi

printf "$(date) ==== Starting ==== \n"

cd `dirname $0`/..
chmod 755 $SERVICE_NAME".jar"
./$SERVICE_NAME".jar" start

rc=$?;

if [[ $rc != 0 ]];
then
    echo "$(date) Failed to start $SERVICE_NAME.jar, return code: $rc"
    exit $rc;
fi

tail -f /dev/null
```

- 写Dockerfile

```shell
[root@shkf6-245 apollo-configservice]# cat Dockerfile 
FROM sunrisenan/jre8:8u112

ENV VERSION 1.5.1

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime &&\
    echo "Asia/Shanghai" > /etc/timezone

ADD apollo-configservice-${VERSION}.jar /apollo-configservice/apollo-configservice.jar
ADD config/ /apollo-configservice/config
ADD scripts/ /apollo-configservice/scripts

CMD ["/apollo-configservice/scripts/startup.sh"]
```

- 制作镜像并推送

```shell
[root@shkf6-245 apollo-configservice]# docker build . -t harbor.od.com/infra/apollo-configservice:v1.5.1

[root@shkf6-245 apollo-configservice]# docker push harbor.od.com/infra/apollo-configservice:v1.5.1
```

## 7.解析域名

DNS主机shkf6-241.host.com上：

```shell
[root@shkf6-241 ~]# tail -2 /var/named/od.com.zone 
mysql              A    192.168.6.241
config             A    192.168.6.66
```

## 8.准备资源配置清单

在运维主机shkf6-245.host.com上

```shell
[root@shkf6-245 apollo-configservice]# mkdir /data/k8s-yaml/apollo-configservice && cd /data/k8s-yaml/apollo-configservice
```

- ConfigMap

```shell
[root@shkf6-245 apollo-configservice]# cat cm.yaml 
apiVersion: v1
kind: ConfigMap
metadata:
  name: apollo-configservice-cm
  namespace: infra
data:
  application-github.properties: |
    # DataSource
    spring.datasource.url = jdbc:mysql://mysql.od.com:3306/ApolloConfigDB?characterEncoding=utf8
    spring.datasource.username = apolloconfig
    spring.datasource.password = 123456
    eureka.service.url = http://config.od.com/eureka
  app.properties: |
    appId=100003171
```

- Deployment

```shell
[root@shkf6-245 apollo-configservice]# cat dp.yaml 
kind: Deployment
apiVersion: extensions/v1beta1
metadata:
  name: apollo-configservice
  namespace: infra
  labels: 
    name: apollo-configservice
spec:
  replicas: 1
  selector:
    matchLabels: 
      name: apollo-configservice
  template:
    metadata:
      labels: 
        app: apollo-configservice 
        name: apollo-configservice
    spec:
      volumes:
      - name: configmap-volume
        configMap:
          name: apollo-configservice-cm
      containers:
      - name: apollo-configservice
        image: harbor.od.com/infra/apollo-configservice:v1.5.1
        ports:
        - containerPort: 8080
          protocol: TCP
        volumeMounts:
        - name: configmap-volume
          mountPath: /apollo-configservice/config
        terminationMessagePath: /dev/termination-log
        terminationMessagePolicy: File
        imagePullPolicy: IfNotPresent
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

- Service

```shell
[root@shkf6-245 apollo-configservice]# cat svc.yaml 
kind: Service
apiVersion: v1
metadata: 
  name: apollo-configservice
  namespace: infra
spec:
  ports:
  - protocol: TCP
    port: 8080
    targetPort: 8080
  selector: 
    app: apollo-configservice
```

- Ingress

```shell
[root@shkf6-245 apollo-configservice]# cat ingress.yaml 
kind: Ingress
apiVersion: extensions/v1beta1
metadata: 
  name: apollo-configservice
  namespace: infra
spec:
  rules:
  - host: config.od.com
    http:
      paths:
      - path: /
        backend: 
          serviceName: apollo-configservice
          servicePort: 8080
```

## 9.应用资源配置清单

```shell
[root@shkf6-243 ~]# kubectl apply -f http://k8s-yaml.od.com/apollo-configservice/cm.yaml
configmap/apollo-configservice-cm created
[root@shkf6-243 ~]# kubectl apply -f http://k8s-yaml.od.com/apollo-configservice/dp.yaml
deployment.extensions/apollo-configservice created
[root@shkf6-243 ~]# kubectl apply -f http://k8s-yaml.od.com/apollo-configservice/svc.yaml
service/apollo-configservice created
[root@shkf6-243 ~]# kubectl apply -f http://k8s-yaml.od.com/apollo-configservice/ingress.yaml
ingress.extensions/apollo-configservice created
```

## 10.浏览器访问

[http://config.od.com](http://config.od.com/)

## 11.分析mysql连接ip

```shell
MariaDB [ApolloConfigDB]> show processlist;
+----+--------------+---------------------+----------------+---------+------+-------+------------------+----------+
| Id | User         | Host                | db             | Command | Time | State | Info             | Progress |
+----+--------------+---------------------+----------------+---------+------+-------+------------------+----------+
|  5 | root         | localhost           | ApolloConfigDB | Query   |    0 | init  | show processlist |    0.000 |
|  7 | apolloconfig | 192.168.6.243:46834 | ApolloConfigDB | Sleep   |    0 |       | NULL             |    0.000 |
|  8 | apolloconfig | 192.168.6.243:46836 | ApolloConfigDB | Sleep   |    0 |       | NULL             |    0.000 |
|  9 | apolloconfig | 192.168.6.243:46838 | ApolloConfigDB | Sleep   |    0 |       | NULL             |    0.000 |
| 10 | apolloconfig | 192.168.6.243:46840 | ApolloConfigDB | Sleep   | 1218 |       | NULL             |    0.000 |
| 11 | apolloconfig | 192.168.6.243:46842 | ApolloConfigDB | Sleep   | 1218 |       | NULL             |    0.000 |
| 12 | apolloconfig | 192.168.6.243:46844 | ApolloConfigDB | Sleep   | 1217 |       | NULL             |    0.000 |
| 13 | apolloconfig | 192.168.6.243:46846 | ApolloConfigDB | Sleep   | 1217 |       | NULL             |    0.000 |
| 14 | apolloconfig | 192.168.6.243:46848 | ApolloConfigDB | Sleep   | 1217 |       | NULL             |    0.000 |
| 15 | apolloconfig | 192.168.6.243:46850 | ApolloConfigDB | Sleep   | 1217 |       | NULL             |    0.000 |
| 16 | apolloconfig | 192.168.6.243:46852 | ApolloConfigDB | Sleep   | 1217 |       | NULL             |    0.000 |
+----+--------------+---------------------+----------------+---------+------+-------+------------------+----------+
```

# 第五章：实战交付apollo配置中心组件–adminservice到kubernetes集群

## 1.准备软件包

在运维主机shkf6-245.host.com上：

[下载官方release包](https://github.com/ctripcorp/apollo/releases/download/v1.5.1/apollo-adminservice-1.5.1-github.zip)

- 下载并解压

```shell
[root@shkf6-245 ~]# wget -O /opt/src/apollo-adminservice-1.5.1-github.zip https://github.com/ctripcorp/apollo/releases/download/v1.5.1/apollo-adminservice-1.5.1-github.zip
[root@shkf6-245 ~]# mkdir /data/dockerfile/apollo-adminservice && unzip -o /opt/src/apollo-adminservice-1.5.1-github.zip -d /data/dockerfile/apollo-adminservice
```

- 删除无用的文件

```shell
[root@shkf6-245 ~]# rm -f /data/dockerfile/apollo-adminservice/apollo-adminservice-1.5.1-sources.jar 
[root@shkf6-245 ~]# rm -f /data/dockerfile/apollo-adminservice/apollo-adminservice.conf
[root@shkf6-245 ~]# rm -f /data/dockerfile/apollo-adminservice/scripts/shutdown.sh
[root@shkf6-245 ~]# cat /data/dockerfile/apollo-adminservice/config/app.properties 
appId=100003172
jdkVersion=1.8
```

## 2.制作Docker镜像

在运维主机shkf6-245.host.com上：

- 配置数据库连接串

```shell
[root@shkf6-245 ~]# cat /data/dockerfile/apollo-adminservice/config/application-github.properties
# DataSource
spring.datasource.url = jdbc:mysql://mysql.od.com:3306/ApolloConfigDB?characterEncoding=utf8
spring.datasource.username = apolloconfig
spring.datasource.password = 123456
```

- 修改启动脚本

```shell
[root@shkf6-245 ~]# cat /data/dockerfile/apollo-adminservice/scripts/startup.sh
#!/bin/bash
SERVICE_NAME=apollo-adminservice
## Adjust log dir if necessary
LOG_DIR=/opt/logs/apollo-admin-server
## Adjust server port if necessary
SERVER_PORT=8080
APOLLO_ADMIN_SERVICE_NAME=$(hostname -i)

# SERVER_URL="http://localhost:${SERVER_PORT}"
SERVER_URL="http://${APOLLO_ADMIN_SERVICE_NAME}:${SERVER_PORT}"

## Adjust memory settings if necessary
#export JAVA_OPTS="-Xms2560m -Xmx2560m -Xss256k -XX:MetaspaceSize=128m -XX:MaxMetaspaceSize=384m -XX:NewSize=1536m -XX:MaxNewSize=1536m -XX:SurvivorRatio=8"

## Only uncomment the following when you are using server jvm
#export JAVA_OPTS="$JAVA_OPTS -server -XX:-ReduceInitialCardMarks"

########### The following is the same for configservice, adminservice, portal ###########
export JAVA_OPTS="$JAVA_OPTS -XX:ParallelGCThreads=4 -XX:MaxTenuringThreshold=9 -XX:+DisableExplicitGC -XX:+ScavengeBeforeFullGC -XX:SoftRefLRUPolicyMSPerMB=0 -XX:+ExplicitGCInvokesConcurrent -XX:+HeapDumpOnOutOfMemoryError -XX:-OmitStackTraceInFastThrow -Duser.timezone=Asia/Shanghai -Dclient.encoding.override=UTF-8 -Dfile.encoding=UTF-8 -Djava.security.egd=file:/dev/./urandom"
export JAVA_OPTS="$JAVA_OPTS -Dserver.port=$SERVER_PORT -Dlogging.file=$LOG_DIR/$SERVICE_NAME.log -XX:HeapDumpPath=$LOG_DIR/HeapDumpOnOutOfMemoryError/"

# Find Java
if [[ -n "$JAVA_HOME" ]] && [[ -x "$JAVA_HOME/bin/java" ]]; then
    javaexe="$JAVA_HOME/bin/java"
elif type -p java > /dev/null 2>&1; then
    javaexe=$(type -p java)
elif [[ -x "/usr/bin/java" ]];  then
    javaexe="/usr/bin/java"
else
    echo "Unable to find Java"
    exit 1
fi

if [[ "$javaexe" ]]; then
    version=$("$javaexe" -version 2>&1 | awk -F '"' '/version/ {print $2}')
    version=$(echo "$version" | awk -F. '{printf("%03d%03d",$1,$2);}')
    # now version is of format 009003 (9.3.x)
    if [ $version -ge 011000 ]; then
        JAVA_OPTS="$JAVA_OPTS -Xlog:gc*:$LOG_DIR/gc.log:time,level,tags -Xlog:safepoint -Xlog:gc+heap=trace"
    elif [ $version -ge 010000 ]; then
        JAVA_OPTS="$JAVA_OPTS -Xlog:gc*:$LOG_DIR/gc.log:time,level,tags -Xlog:safepoint -Xlog:gc+heap=trace"
    elif [ $version -ge 009000 ]; then
        JAVA_OPTS="$JAVA_OPTS -Xlog:gc*:$LOG_DIR/gc.log:time,level,tags -Xlog:safepoint -Xlog:gc+heap=trace"
    else
        JAVA_OPTS="$JAVA_OPTS -XX:+UseParNewGC"
        JAVA_OPTS="$JAVA_OPTS -Xloggc:$LOG_DIR/gc.log -XX:+PrintGCDetails"
        JAVA_OPTS="$JAVA_OPTS -XX:+UseConcMarkSweepGC -XX:+UseCMSCompactAtFullCollection -XX:+UseCMSInitiatingOccupancyOnly -XX:CMSInitiatingOccupancyFraction=60 -XX:+CMSClassUnloadingEnabled -XX:+CMSParallelRemarkEnabled -XX:CMSFullGCsBeforeCompaction=9 -XX:+CMSClassUnloadingEnabled  -XX:+PrintGCDateStamps -XX:+PrintGCApplicationConcurrentTime -XX:+PrintHeapAtGC -XX:+UseGCLogFileRotation -XX:NumberOfGCLogFiles=5 -XX:GCLogFileSize=5M"
    fi
fi

printf "$(date) ==== Starting ==== \n"

cd `dirname $0`/..
chmod 755 $SERVICE_NAME".jar"
./$SERVICE_NAME".jar" start

rc=$?;

if [[ $rc != 0 ]];
then
    echo "$(date) Failed to start $SERVICE_NAME.jar, return code: $rc"
    exit $rc;
fi

tail -f /dev/null
```

在官网的基础上修改了这两个参数

```shell
SERVER_PORT=8080
APOLLO_ADMIN_SERVICE_NAME=$(hostname -i)
```

- 编写dockerfile

```shell
[root@shkf6-245 ~]# cd /data/dockerfile/apollo-adminservice/
[root@shkf6-245 apollo-adminservice]# vi Dockerfile

FROM sunrisenan/jre8:8u112

ENV VERSION 1.5.1

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime &&\
    echo "Asia/Shanghai" > /etc/timezone

ADD apollo-adminservice-${VERSION}.jar /apollo-adminservice/apollo-adminservice.jar
ADD config/ /apollo-adminservice/config
ADD scripts/ /apollo-adminservice/scripts

CMD ["/apollo-adminservice/scripts/startup.sh"]
```

- 制作镜像并推送

```shell
[root@shkf6-245 apollo-adminservice]# docker build . -t harbor.od.com/infra/apollo-adminservice:v1.5.1

[root@shkf6-245 apollo-adminservice]# docker push harbor.od.com/infra/apollo-adminservice:v1.5.1
```

## 3.准备配置文件

在运维主机shkf6-245.host.com

```shell
[root@shkf6-245 apollo-adminservice]# mkdir /data/k8s-yaml/apollo-adminservice && cd /data/k8s-yaml/apollo-adminservice
```

- ConfigMap

```shell
[root@shkf6-245 apollo-adminservice]# vi cm.yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: apollo-adminservice-cm
  namespace: infra
data:
  application-github.properties: |
    # DataSource
    spring.datasource.url = jdbc:mysql://mysql.od.com:3306/ApolloConfigDB?characterEncoding=utf8
    spring.datasource.username = apolloconfig
    spring.datasource.password = 123456
    eureka.service.url = http://config.od.com/eureka
  app.properties: |
    appId=100003172
```

- Deployment

```shell
[root@shkf6-245 apollo-adminservice]# vi dp.yaml
kind: Deployment
apiVersion: extensions/v1beta1
metadata:
  name: apollo-adminservice
  namespace: infra
  labels: 
    name: apollo-adminservice
spec:
  replicas: 1
  selector:
    matchLabels: 
      name: apollo-adminservice
  template:
    metadata:
      labels: 
        app: apollo-adminservice 
        name: apollo-adminservice
    spec:
      volumes:
      - name: configmap-volume
        configMap:
          name: apollo-adminservice-cm
      containers:
      - name: apollo-adminservice
        image: harbor.od.com/infra/apollo-adminservice:v1.5.1
        ports:
        - containerPort: 8080
          protocol: TCP
        volumeMounts:
        - name: configmap-volume
          mountPath: /apollo-adminservice/config
        terminationMessagePath: /dev/termination-log
        terminationMessagePolicy: File
        imagePullPolicy: IfNotPresent
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
    appId=100003172
```

## 4.应用配置清单

在任何一台运算节点上：

```shell
[root@shkf6-244 ~]# kubectl apply -f http://k8s-yaml.od.com/apollo-adminservice/cm.yaml
configmap/apollo-adminservice-cm created

[root@shkf6-244 ~]# kubectl apply -f http://k8s-yaml.od.com/apollo-adminservice/dp.yaml
deployment.extensions/apollo-adminservice created
```

## 5.浏览器访问

[http://config.od.com](http://config.od.com/)

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_9be9bc201d9670735918a3edc724097d_r.png)

```shell
[root@shkf6-244 ~]# curl http://172.6.244.4:8080/info
{"git":{"commit":{"time":{"seconds":1573275854,"nanos":0},"id":"c9eae54"},"branch":"1.5.1"}}
```

# 第六章：实战交付apollo配置中心组件–portal到kubernetes集群

## 1.准备软件包

在运维主机HDSS6-245.host.com上：

[下载官方release包](https://github.com/ctripcorp/apollo/releases/download/v1.5.1/apollo-portal-1.5.1-github.zip)

```shell
[root@shkf6-245 ~]# wget -O /opt/src/apollo-portal-1.5.1-github.zip https://github.com/ctripcorp/apollo/releases/download/v1.5.1/apollo-portal-1.5.1-github.zip

[root@shkf6-245 ~]# mkdir /data/dockerfile/apollo-portal && unzip -o /opt/src/apollo-portal-1.5.1-github.zip -d /data/dockerfile/apollo-portal
```

- 清理不用的文件

```shell
[root@shkf6-245 ~]# cd /data/dockerfile/apollo-portal/
[root@shkf6-245 apollo-portal]# rm -f apollo-portal-1.5.1-sources.jar 
[root@shkf6-245 apollo-portal]# rm -f apollo-portal.conf 
[root@shkf6-245 apollo-portal]# rm -f scripts/shutdown.sh 
```

## 2.执行数据库脚本

在数据库主机shkf6-241.host.com上：

[数据库脚本地址](https://raw.githubusercontent.com/ctripcorp/apollo/master/scripts/db/migration/portaldb/V1.0.0__initialization.sql)

[root[@shkf6](https://github.com/shkf6)-241 ~]# wget -O apolloportal.sql https://raw.githubusercontent.com/ctripcorp/apollo/master/scripts/db/migration/portaldb/V1.0.0__initialization.sql

```shell
[root@hdss7-11 ~]# mysql -uroot -p
MariaDB [ApolloConfigDB]> source ./apolloportal.sql
```

## 3.数据库用户授权

```shell
MariaDB [ApolloPortalDB]> grant INSERT,DELETE,UPDATE,SELECT on ApolloPortalDB.* to "apolloportal"@"192.168.6.%" identified by "123456";
MariaDB [ApolloPortalDB]> update ServerConfig set Value='[{"orgId":"od01","orgName":"Linux学院"},{"orgId":"od02","orgName":"云计算学院"},{"orgId":"od03","orgName":"Python学院"}]' where Id=2;
```

## 4.制作Docker镜像

在运维主机shkf6-245.host.com上：

- 配置数据库连接串（用cm的话这里可以不用修改）

```shell
[root@shkf6-245 apollo-portal]# cat config/application-github.properties
# DataSource
spring.datasource.url = jdbc:mysql://mysql.od.com:3306/ApolloPortalDB?characterEncoding=utf8
spring.datasource.username = apolloportal
spring.datasource.password = 123456
```

- 配置Portal的meta service（用cm的话这里可以不用修改）

```shell
[root@shkf6-245 apollo-portal]# cat config/apollo-env.properties
dev.meta=http://config.od.com
```

- 更新startup.sh

```shell
[root@shkf6-245 apollo-portal]# cat scripts/startup.sh
#!/bin/bash
SERVICE_NAME=apollo-portal
## Adjust log dir if necessary
LOG_DIR=/opt/logs/apollo-portal-server
## Adjust server port if necessary
SERVER_PORT=8080
APOLLO_PORTAL_SERVICE_NAME=$(hostname -i)

# SERVER_URL="http://localhost:$SERVER_PORT"
SERVER_URL="http://${APOLLO_PORTAL_SERVICE_NAME}:${SERVER_PORT}"

## Adjust memory settings if necessary
#export JAVA_OPTS="-Xms2560m -Xmx2560m -Xss256k -XX:MetaspaceSize=128m -XX:MaxMetaspaceSize=384m -XX:NewSize=1536m -XX:MaxNewSize=1536m -XX:SurvivorRatio=8"

## Only uncomment the following when you are using server jvm
#export JAVA_OPTS="$JAVA_OPTS -server -XX:-ReduceInitialCardMarks"

########### The following is the same for configservice, adminservice, portal ###########
export JAVA_OPTS="$JAVA_OPTS -XX:ParallelGCThreads=4 -XX:MaxTenuringThreshold=9 -XX:+DisableExplicitGC -XX:+ScavengeBeforeFullGC -XX:SoftRefLRUPolicyMSPerMB=0 -XX:+ExplicitGCInvokesConcurrent -XX:+HeapDumpOnOutOfMemoryError -XX:-OmitStackTraceInFastThrow -Duser.timezone=Asia/Shanghai -Dclient.encoding.override=UTF-8 -Dfile.encoding=UTF-8 -Djava.security.egd=file:/dev/./urandom"
export JAVA_OPTS="$JAVA_OPTS -Dserver.port=$SERVER_PORT -Dlogging.file=$LOG_DIR/$SERVICE_NAME.log -XX:HeapDumpPath=$LOG_DIR/HeapDumpOnOutOfMemoryError/"

# Find Java
if [[ -n "$JAVA_HOME" ]] && [[ -x "$JAVA_HOME/bin/java" ]]; then
    javaexe="$JAVA_HOME/bin/java"
elif type -p java > /dev/null 2>&1; then
    javaexe=$(type -p java)
elif [[ -x "/usr/bin/java" ]];  then
    javaexe="/usr/bin/java"
else
    echo "Unable to find Java"
    exit 1
fi

if [[ "$javaexe" ]]; then
    version=$("$javaexe" -version 2>&1 | awk -F '"' '/version/ {print $2}')
    version=$(echo "$version" | awk -F. '{printf("%03d%03d",$1,$2);}')
    # now version is of format 009003 (9.3.x)
    if [ $version -ge 011000 ]; then
        JAVA_OPTS="$JAVA_OPTS -Xlog:gc*:$LOG_DIR/gc.log:time,level,tags -Xlog:safepoint -Xlog:gc+heap=trace"
    elif [ $version -ge 010000 ]; then
        JAVA_OPTS="$JAVA_OPTS -Xlog:gc*:$LOG_DIR/gc.log:time,level,tags -Xlog:safepoint -Xlog:gc+heap=trace"
    elif [ $version -ge 009000 ]; then
        JAVA_OPTS="$JAVA_OPTS -Xlog:gc*:$LOG_DIR/gc.log:time,level,tags -Xlog:safepoint -Xlog:gc+heap=trace"
    else
        JAVA_OPTS="$JAVA_OPTS -XX:+UseParNewGC"
        JAVA_OPTS="$JAVA_OPTS -Xloggc:$LOG_DIR/gc.log -XX:+PrintGCDetails"
        JAVA_OPTS="$JAVA_OPTS -XX:+UseConcMarkSweepGC -XX:+UseCMSCompactAtFullCollection -XX:+UseCMSInitiatingOccupancyOnly -XX:CMSInitiatingOccupancyFraction=60 -XX:+CMSClassUnloadingEnabled -XX:+CMSParallelRemarkEnabled -XX:CMSFullGCsBeforeCompaction=9 -XX:+CMSClassUnloadingEnabled  -XX:+PrintGCDateStamps -XX:+PrintGCApplicationConcurrentTime -XX:+PrintHeapAtGC -XX:+UseGCLogFileRotation -XX:NumberOfGCLogFiles=5 -XX:GCLogFileSize=5M"
    fi
fi

printf "$(date) ==== Starting ==== \n"

cd `dirname $0`/..
chmod 755 $SERVICE_NAME".jar"
./$SERVICE_NAME".jar" start

rc=$?;

if [[ $rc != 0 ]];
then
    echo "$(date) Failed to start $SERVICE_NAME.jar, return code: $rc"
    exit $rc;
fi

tail -f /dev/null
```

- 写Dockerfile

```shell
[root@shkf6-245 apollo-portal]# cat Dockerfile
FROM sunrisenan/jre8:8u112

ENV VERSION 1.5.1

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime &&\
    echo "Asia/Shanghai" > /etc/timezone

ADD apollo-portal-${VERSION}.jar /apollo-portal/apollo-portal.jar
ADD config/ /apollo-portal/config
ADD scripts/ /apollo-portal/scripts

CMD ["/apollo-portal/scripts/startup.sh"]
```

- 制作镜像并推送

```shell
[root@shkf6-245 apollo-portal]# docker build . -t harbor.od.com/infra/apollo-portal:v1.5.1

[root@shkf6-245 apollo-portal]# docker push harbor.od.com/infra/apollo-portal:v1.5.1
```

## 5.解析域名

DNS主机shkf6-241.host.com上：

```shell
[root@shkf6-241 ~]# tail -1 /var/named/od.com.zone
portal             A    192.168.6.66
```

## 6.准备资源配置清单

在运维主机shkf6-245.host.com上

```shell
[root@shkf6-245 apollo-portal]# mkdir /data/k8s-yaml/apollo-portal && cd /data/k8s-yaml/apollo-portal
```

- ConfigMap

```shell
[root@shkf6-245 apollo-portal]# cat cm.yaml 
apiVersion: v1
kind: ConfigMap
metadata:
  name: apollo-portal-cm
  namespace: infra
data:
  application-github.properties: |
    # DataSource
    spring.datasource.url = jdbc:mysql://mysql.od.com:3306/ApolloPortalDB?characterEncoding=utf8
    spring.datasource.username = apolloportal
    spring.datasource.password = 123456
  app.properties: |
    appId=100003173
  apollo-env.properties: |
    dev.meta=http://config.od.com
```

- Deployment

```shell
[root@shkf6-245 apollo-portal]# cat dp.yaml 
kind: Deployment
apiVersion: extensions/v1beta1
metadata:
  name: apollo-portal
  namespace: infra
  labels: 
    name: apollo-portal
spec:
  replicas: 1
  selector:
    matchLabels: 
      name: apollo-portal
  template:
    metadata:
      labels: 
        app: apollo-portal 
        name: apollo-portal
    spec:
      volumes:
      - name: configmap-volume
        configMap:
          name: apollo-portal-cm
      containers:
      - name: apollo-portal
        image: harbor.od.com/infra/apollo-portal:v1.5.1
        ports:
        - containerPort: 8080
          protocol: TCP
        volumeMounts:
        - name: configmap-volume
          mountPath: /apollo-portal/config
        terminationMessagePath: /dev/termination-log
        terminationMessagePolicy: File
        imagePullPolicy: IfNotPresent
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

- service

```shell
[root@shkf6-245 apollo-portal]# cat svc.yaml 
kind: Service
apiVersion: v1
metadata: 
  name: apollo-portal
  namespace: infra
spec:
  ports:
  - protocol: TCP
    port: 8080
    targetPort: 8080
  selector: 
    app: apollo-portal
```

- ingress

```shell
[root@shkf6-245 apollo-portal]# cat ingress.yaml 
kind: Ingress
apiVersion: extensions/v1beta1
metadata: 
  name: apollo-portal
  namespace: infra
spec:
  rules:
  - host: portal.od.com
    http:
      paths:
      - path: /
        backend: 
          serviceName: apollo-portal
          servicePort: 8080
```

## 7.应用资源配置清单

在任意一台k8s运算节点执行：

```shell
[root@shkf6-243 ~]# kubectl apply -f http://k8s-yaml.od.com/apollo-portal/cm.yaml
configmap/apollo-portal-cm created
[root@shkf6-243 ~]# kubectl apply -f http://k8s-yaml.od.com/apollo-portal/dp.yaml
deployment.extensions/apollo-portal created
[root@shkf6-243 ~]# kubectl apply -f http://k8s-yaml.od.com/apollo-portal/svc.yaml
service/apollo-portal created
[root@shkf6-243 ~]# kubectl apply -f http://k8s-yaml.od.com/apollo-portal/ingress.yaml
ingress.extensions/apollo-portal created
```

## 8.浏览器访问

[http://portal.od.com](http://portal.od.com/)

- 用户名：apollo
- 密码： admin

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_c1a3e387ab6bae3d89d8ad61b5cd03c7_r.png)

# 第七章：实战配置dubbo微服务接入apollo配置中心管理

## 1.改造dubbo-demo-service项目

### 1.使用IDE拉取项目（这里使用git bash作为范例）

```shell
~]# git clone https://github.com/sunrisenan/dubbo-demo-service.git
```

### 2.切到apollo分支

```shell
~]# cd dubbo-demo-service/
dubbo-demo-service]# git checkout -b apollo
```

### 3.修改pom.xml

- 加入apollo客户端jar包的依赖

```java
dubbo-server/pom.xml

<dependency>
  <groupId>com.ctrip.framework.apollo</groupId>
  <artifactId>apollo-client</artifactId>
  <version>1.1.0</version>
</dependency>
```

- 修改resource段

```java
dubbo-server/pom.xml


<resource>
  <directory>src/main/resources</directory>
  <includes>
  <include>**/*</include>
  </includes>
  <filtering>false</filtering>
</resource>
```

### 4.增加resources目录

```java
/d/workspace/dubbo-demo-service/dubbo-server/src/main


$ mkdir -pv resources/META-INF
mkdir: created directory 'resources'
mkdir: created directory 'resources/META-INF'
```

### 5.修改config.properties文件

```java
/d/workspace/dubbo-demo-service/dubbo-server/src/main/resources/config.properties


dubbo.registry=${dubbo.registry}
dubbo.port=${dubbo.port}
```

### 6.修改srping-config.xml文件

- beans段新增属性

```java
/d/workspace/dubbo-demo-service/dubbo-server/src/main/resources/spring-config.xml

xmlns:apollo="http://www.ctrip.com/schema/apollo"
```

- xsi:schemaLocation段内新增属性

```java
/d/workspace/dubbo-demo-service/dubbo-server/src/main/resources/spring-config.xml

http://www.ctrip.com/schema/apollo http://www.ctrip.com/schema/apollo.xsd
```

- 新增配置项

```java
/d/workspace/dubbo-demo-service/dubbo-server/src/main/resources/spring-config.xml

<apollo:config/>
```

- 删除配置项（注释）

```java
/d/workspace/dubbo-demo-service/dubbo-server/src/main/resources/spring-config.xml

<!-- <context:property-placeholder location="classpath:config.properties"/> -->
```

- 增加app.properties文件

```java
/d/workspace/dubbo-demo-service/dubbo-server/src/main/resources/META-INF/app.properties

app.id=dubbo-demo-service
```

- 提交git中心仓库（github）

```java
$ git push origin apollo
```

## 2.配置apollo-portal

### 1.创建项目

- 部门

> 样例部门1（老男孩linux学院001）

- 应用id

> dubbo-demo-service

- 应用名称

> dubbo服务提供者

- 应用负责人

> apollo|apollo

- 项目管理员

> apollo|apollo

提交

### 2.进入配置界面

**新增配置项1**

- Key

> dubbo.registry

- Value

> zookeeper://zk1.od.com:2181

- 选择集群

> DEV

提交

**新增配置项2**

- Key

> dubbo.port

- Value

> 20880

- 选择集群

> DEV

提交

### 3.发布配置

点击发布，配置生效

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_289aeb7203f30ba86b82096de0a81b39_r.png)

## 3.使用jenkins进行CI

略（注意记录镜像的tag）

## 4.上线新构建的项目

### 1.准备资源配置清单

运维主机shkf6-245.host.com上：

```shell
[root@shkf6-245 ~]# cat /data/k8s-yaml/dubbo-demo-service-apollo/deployment.yaml
kind: Deployment
apiVersion: extensions/v1beta1
metadata:
  name: dubbo-demo-service
  namespace: app
  labels: 
    name: dubbo-demo-service
spec:
  replicas: 1
  selector:
    matchLabels: 
      name: dubbo-demo-service
  template:
    metadata:
      labels: 
        app: dubbo-demo-service
        name: dubbo-demo-service
    spec:
      containers:
      - name: dubbo-demo-service
        image: harbor.od.com/app/dubbo-demo-service:apollo_191211_1326
        ports:
        - containerPort: 20880
          protocol: TCP
        env:
        - name: JAR_BALL
          value: dubbo-server.jar
        - name: C_OPTS
          value: -Denv=dev -Dapollo.meta=http://config.od.com
        imagePullPolicy: IfNotPresent
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

注意：增加了env段配置
注意：docker镜像新版的tag

### 2.应用资源配置清单

在任意一台k8s运算节点上执行：

```shell
[root@shkf6-243 ~]# kubectl apply -f http://k8s-yaml.od.com/dubbo-demo-service-apollo/deployment.yaml
```

### 3.观察项目运行情况

[http://dubbo-monitor.od.com](http://dubbo-monitor.od.com/)

## 5.改造dubbo-demo-web

略

## 6.配置apollo-portal

### 1.创建项目

- 部门

> 样例部门1（linux学院od01）

- 应用id

> dubbo-demo-web

- 应用名称

> dubbo服务消费者

- 应用负责人

> apollo|apollo

- 项目管理员

> apollo|apollo

提交

### 2.进入配置页面

**新增配置项1**

- Key

> dubbo.registry

- Value

> zookeeper://zk1.od.com:2181

- 选择集群

> DEV

提交

### 3.发布配置

点击发布，配置生效

## 7.使用jenkins进行CI

略（注意记录镜像的tag）

## 8.上线新构建的项目

### 1.准备资源配置清单

运维主机shkf6-245.host.com上：

```shell
[root@shkf6-245 ~]# cat /data/k8s-yaml/dubbo-demo-consumer-apollo/dp.yaml 
kind: Deployment
apiVersion: extensions/v1beta1
metadata:
  name: dubbo-demo-consumer
  namespace: app
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
        image: harbor.od.com/app/dubbo-demo-consumer:apollo_191211_1417
        ports:
        - containerPort: 8080
          protocol: TCP
        - containerPort: 20880
          protocol: TCP
        env:
        - name: JAR_BALL
          value: dubbo-client.jar
        - name: C_OPTS
          value: -Denv=dev -Dapollo.meta=http://config.od.com
        imagePullPolicy: IfNotPresent
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

注意：增加了env段配置
注意：docker镜像新版的tag

### 2.应用资源配置清单

在任意一台k8s运算节点上执行：

```shell
[root@shkf6-243 ~]# kubectl apply -f http://k8s-yaml.od.com/dubbo-demo-consumer-apollo/dp.yaml
```

## 9.通过Apollo配置中心动态维护项目的配置

以dubbo-demo-service项目为例，不用修改代码

- 在[http://portal.od.com](http://portal.od.com/) 里修改dubbo.port配置项
- 重启dubbo-demo-service项目
- 配置生效

# 第八章：实战使用apollo配置中心管理测试环境和生产环境

## 1.实战维护多套dubbo微服务环境

### 1.生产实践

1. 迭代新需求/修复BUG（编码->提GIT）
2. 测试环境发版，测试（应用通过编译打包发布至TEST命名空间）
3. 测试通过，上线（应用镜像直接发布至PROD命名空间）

### 2.系统架构

- 物理架构

| 主机名             | 角色                  | ip            |
| :----------------- | :-------------------- | :------------ |
| shkf6-241.host.com | zk-test(测试环境Test) | 192.168.6.241 |
| shkf6-242.host.com | zk-prod(生产环境Prod) | 192.168.6.242 |
| shkf6-243.host.com | kubernetes运算节点    | 192.168.6.243 |
| shkf6-244.host.com | kubernetes运算节点    | 192.168.6.244 |
| shkf6-245.host.com | 运维主机，harbor仓库  | 192.168.6.245 |

- K8S内系统架构

| 环境             | 命名空间 | 应用                                  |
| :--------------- | :------- | :------------------------------------ |
| 测试环境（TEST） | test     | apollo-config，apollo-admin           |
| 测试环境（TEST） | test     | dubbo-demo-service，dubbo-demo-web    |
| 生产环境（PROD） | prod     | apollo-config，apollo-admin           |
| 生产环境（PROD） | prod     | dubbo-demo-service，dubbo-demo-web    |
| ops环境（infra） | infra    | jenkins，dubbo-monitor，apollo-portal |

### 3.修改/添加域名解析

DNS主机HDSS6-241.host.com上：

```shell
[root@shkf6-241 ~]# tail /var/named/od.com.zone
zk-test            A    192.168.6.241
zk-prod            A    192.168.6.242
config-test        A    192.168.6.66
config-prod        A    192.168.6.66
demo-test          A    192.168.6.66
demo-prod          A    192.168.6.66
```

### 4.Apollo的k8s应用配置

- 删除app命名空间内应用，创建test命名空间，创建prod命名空间

```shell
[root@shkf6-243 ~]# kubectl create ns test
[root@shkf6-243 ~]# kubectl create ns prod
```

- 配置连接docker仓库的认证

```shell
[root@shkf6-243 ~]# kubectl create secret docker-registry harbor --docker-server=harbor.od.com --docker-username=admin --docker-password=Harbor12345 -n test
[root@shkf6-243 ~]# kubectl create secret docker-registry harbor --docker-server=harbor.od.com --docker-username=admin --docker-password=Harbor12345 -n prod
```

- 删除infra命名空间内apollo-configservice，apollo-adminservice应用
- 数据库内删除ApolloConfigDB，创建ApolloConfigTestDB，创建ApolloConfigProdDB

[sql代码下载地址](http://down.sunrisenan.com/apollo/)

测试：

```shell
[root@shkf6-241 apollo]# mysql -uroot -p123456 < test/apolloconfig.sql 

> update ApolloConfigTestDB.ServerConfig set ServerConfig.Value="http://config-test.od.com/eureka" where ServerConfig.Key="eureka.service.url";

> grant INSERT,DELETE,UPDATE,SELECT on ApolloConfigTestDB.* to "apolloconfig"@"192.168.6.%" identified by "123456";
```

生产：

```shell
[root@shkf6-241 apollo]# mysql -uroot -p123456 < prod/apolloconfig.sql 

> update ApolloConfigProdDB.ServerConfig set ServerConfig.Value="http://config-prod.od.com/eureka" where ServerConfig.Key="eureka.service.url";

> grant INSERT,DELETE,UPDATE,SELECT on ApolloConfigProdDB.* to "apolloconfig"@"192.168.6.%" identified by "123456";
```

- 准备apollo-config，apollo-admin的资源配置清单（各2套）

注：apollo-config/apollo-admin的configmap配置要点

更改portal数据库-分环境

```shell
> update ApolloPortalDB.ServerConfig set Value='fat,pro' where Id=1;
```

- Test环境

```shell
application-github.properties: |
    # DataSource
    spring.datasource.url = jdbc:mysql://mysql.od.com:3306/ApolloConfigTestDB?characterEncoding=utf8
    spring.datasource.username = apolloconfig
    spring.datasource.password = 123456
    eureka.service.url = http://config-test.od.com/eureka
```

- Prod环境

```shell
application-github.properties: |
    # DataSource
    spring.datasource.url = jdbc:mysql://mysql.od.com:3306/ApolloConfigProdDB?characterEncoding=utf8
    spring.datasource.username = apolloconfig
    spring.datasource.password = 123456
    eureka.service.url = http://config-prod.od.com/eureka
```

- 依次应用，分别发布在test和prod命名空间
- 修改apollo-portal的configmap并重启portal

```shell
apollo-env.properties: |
    TEST.meta=http://config-test.od.com
    PROD.meta=http://config-prod.od.com
```

### 5.Apollo的portal配置

#### 1.管理员工具

删除应用、集群、AppNamespace，将已配置应用删除

#### 2.系统参数

- Key

> apollo.portal.envs

- Value

> fat,pro

查询

- Value

> fat,pro

保存

### 6.新建dubbo-demo-service和dubbo-demo-web项目

在TEST/PROD环境分别增加配置项并发布

### 7.发布dubbo微服务

- 准备dubbo-demo-service和dubbo-demo-web的资源配置清单（各2套）
- 依次应用，分别发布至test和prod命名空间
- 使用dubbo-monitor查验

## 2.互联网公司技术部的日常

- 产品经理整理需求，需求评审，出产品原型
- 开发同学夜以继日的开发，提测
- 测试同学使用Jenkins持续集成，并发布至测试环境
- 验证功能，通过->待上线or打回->修改代码
- 提交发版申请，运维同学将测试后的包发往生产环境
- 无尽的BUG修复（笑cry）

## 3.分环境

[配置资料](http://down.sunrisenan.com/apollo/apollo-fat-pro.tar.gz)

### 第一阶段 解析域名

```shell
[root@shkf6-241 ~]# tail -7 /var/named/od.com.zone
zk-test            A    192.168.6.241
zk-prod            A    192.168.6.242
mysql              A    192.168.6.241
config-test        A    192.168.6.66
config-prod        A    192.168.6.66
demo-test          A    192.168.6.66
demo-prod          A    192.168.6.66


[root@shkf6-241 ~]# dig -t A zk-test.od.com +short
192.168.6.241
```

### 第二阶段 Portal设置

1、查看并清除配置中心的历史记录

如果是新环境的话可以忽略这一步，直接操作第二步：

```shell
MariaDB [(none)]> use ApolloPortalDB

MariaDB [ApolloPortalDB]> select * from App;

MariaDB [ApolloPortalDB]> select * from AppNamespace;

MariaDB [ApolloPortalDB]> truncate table ApolloPortalDB.App;
Query OK, 0 rows affected (0.28 sec)

MariaDB [ApolloPortalDB]> truncate table ApolloPortalDB.AppNamespace;
Query OK, 0 rows affected (0.18 sec)
```

2、ApolloPortalDB分环境fat和pro：

```shell
> update ApolloPortalDB.ServerConfig set Value='fat,pro' where Id=1;
```

3、portal资源配置清单

```shell
[root@shkf6-245 k8s-yaml]# tree apollo-portal-fat-pro/
apollo-portal-fat-pro/
├── cm.yaml
├── dp.yaml
├── ingress.yaml
└── svc.yaml
```

### 第三阶段 测试环境

1、创建名称空间

```shell
[root@shkf6-243 ~]# kubectl create ns test
namespace/test created
[root@shkf6-243 ~]# kubectl create secret docker-registry harbor --docker-server=harbor.od.com --docker-username=admin --docker-password=Harbor12345 -n test
```

2、数据库设置

```shell
[root@shkf6-241 apollo]# mysql -uroot -p123456 < test/apolloconfig.sql 

> update ApolloConfigTestDB.ServerConfig set ServerConfig.Value="http://config-test.od.com/eureka" where ServerConfig.Key="eureka.service.url";

> grant INSERT,DELETE,UPDATE,SELECT on ApolloConfigTestDB.* to "apolloconfig"@"192.168.6.%" identified by "123456";
```

3、资源配置清单准备

```shell
[root@shkf6-245 k8s-yaml]# mkdir -pv test/{apollo-configservice,apollo-adminservice,dubbo-demo-consumer,dubbo-demo-service}

[root@shkf6-245 k8s-yaml]# tree test/
test/
├── apollo-adminservice
│   ├── cm.yaml
│   └── dp.yaml
├── apollo-configservice
│   ├── cm.yaml
│   ├── dp.yaml
│   ├── ingress.yaml
│   └── svc.yaml
├── dubbo-demo-consumer
│   ├── dp.yaml
│   ├── ingress.yaml
│   └── svc.yaml
└── dubbo-demo-service
    └── deployment.yaml
```

### 第四阶段 生产环境

1、创建名称空间

```shell
[root@shkf6-243 ~]# kubectl create ns prod
namespace/prod created
[root@shkf6-243 ~]# kubectl create secret docker-registry harbor --docker-server=harbor.od.com --docker-username=admin --docker-password=Harbor12345 -n prod
```

2、数据库设置

```shell
[root@shkf6-241 apollo]# mysql -uroot -p123456 < prod/apolloconfig.sql 

> update ApolloConfigProdDB.ServerConfig set ServerConfig.Value="http://config-prod.od.com/eureka" where ServerConfig.Key="eureka.service.url";

> grant INSERT,DELETE,UPDATE,SELECT on ApolloConfigProdDB.* to "apolloconfig"@"192.168.6.%" identified by "123456";
```

3、资源配置清单准备

```shell
[root@shkf6-245 k8s-yaml]# mkdir -pv prod/{apollo-configservice,apollo-adminservice,dubbo-demo-consumer,dubbo-demo-service}

[root@shkf6-245 k8s-yaml]# tree prod/
prod/
├── apollo-adminservice
│   ├── cm.yaml
│   └── dp.yaml
├── apollo-configservice
│   ├── cm.yaml
│   ├── dp.yaml
│   ├── ingress.yaml
│   └── svc.yaml
├── dubbo-demo-consumer
│   ├── dp.yaml
│   ├── ingress.yaml
│   └── svc.yaml
└── dubbo-demo-service
    └── deployment.yaml
```

### 知识点：

```shell
    env:
    - name: JAR_BALL
      value: dubbo-server.jar
    - name: C_OPTS
      value: -Denv=pro -Dapollo.meta=http://config-test.od.com        #这里用的是apollo-configservice的ingress地址


    env:
    - name: JAR_BALL
      value: dubbo-server.jar
    - name: C_OPTS
      value: -Denv=pro -Dapollo.meta=http://apollo-configservice:8080     #这里用的是apollo-configservice的svc地址

    [root@shkf6-243 ~]# kubectl get svc -n prod 
    NAME                   TYPE        CLUSTER-IP    EXTERNAL-IP   PORT(S)    AGE
    apollo-configservice   ClusterIP   10.96.2.8     <none>        8080/TCP   4d14h
```