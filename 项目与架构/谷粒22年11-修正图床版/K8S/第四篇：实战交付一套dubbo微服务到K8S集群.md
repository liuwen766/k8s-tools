# 前情回顾

- K8S核心资源管理方法（CRID）
  - 陈述式管理 –> 基于众多kuberctl命令
  - 声明式管理 –> 基于K8S资源配置清单
  - GUI式管理 –> 基于K8S仪表盘（dashboard）
- K8S的CNI网络插件
  - 种类众多，以flannel为例
  - 三种常用工作模式
  - 优化SNAT规则
- K8S服务发现
  - 集群网络 –> Cluster IP
  - Service资源 –> Service Name
  - Coredns软件 –> 实现了Service Name和Cluster IP的自动关联
- K8S的服务暴露
  - Ingress资源 –> 专用于暴露7层应用到K8S集群外的一种核心资源（http/https）
  - Ingress控制器 –> 一个简化版的nginx（调度流量） + go脚本（动态识别yaml）
  - Traefik软件 –> 实现了Ingress控制器的一个软件
- Dashboard（仪表盘）
  - 基于RBAC认证的一个GUI资源管理软件
  - 连个常用版本：V1.8.3和v1.10.1
  - K8S如何基于RBAC进行鉴权
  - 手撕ssl证书签发

# 第一章：Dubbo微服务概述

## 1.dubbo什么？

- dubbo是阿里巴巴SOA服务化治理方案的核心框架，每天为2000+个服务提供3000000000+次访问量支持，并被广泛应用于阿里巴巴集团的各成员站点
- dubbo是一个分布式服务框架，致力于提供高可用性能和透明化的RPC远程服务调用方案，以及SOA服务治理方案。
- 简单的说，dubbo就是一个服务框架，如果没有分布式的需求，其实是不需要用的，只是在分布式的时候，才有dubbo这样的分布式服务框架的需求，并且本质上是个服务调用的东西，说白了就是个远程服务调用的分布式框架。

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_3cc763ccfe6967147255ec9ca0c48b53_r.png)

## 2.dubbo能做什么？

- 透明化的远程方法调用，就像调用本地方法一样调用远程方法，只需要配置，没有任何API侵入。
- 软负载均衡及容错机制，可在内网替代F5等硬件负载均衡器，降低成本，减少单点。
- 服务自动注册与发现，不再需要写死服务提供方地址，注册中心基于接口名查询服务提供者的IP地址，并且能够平滑添加或删除服务提供者。

# 第二章：实验架构详解

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_151103d7cb4d3faf4e1c2c5cb996febf_r.png)

# 第三章：部署zookeeper集群

- Zookeeper是Dubbo微服务集群的注册中心
- 它的高可用机制和K8S的etcd集群一致
- 由Java编写，所以需要jdk环境

## 1.集群规划

| 主机名    | 角色                      | IP            |
| :-------- | :------------------------ | :------------ |
| shkf6-241 | k8s代理节点1,zk1,jdk      | 192.168.6.241 |
| shkf6-242 | k8s代理节点2,zk2,jdk      | 192.168.6.242 |
| shkf6-243 | k8s运算节点1,zk3,jdk      | 192.168.6.243 |
| shkf6-244 | k8s运算节点2,jenkins      | 192.168.6.244 |
| shkf6-245 | k8s运维节点（docker仓库） | 192.168.6.245 |

## 2.安装jdk1.8（3台zk角色主机）

[JDK_ALL下载地址](https://www.oracle.com/technetwork/java/archive-139210.html)

[jdk1.8下载](http://down.sunrisenan.com/oracle/jdk-8u221-linux-x64.tar.gz)

在shkf6-241机器上：

```shell
[root@shkf6-241 ~]# mkdir /opt/src
[root@shkf6-241 ~]# wget -O /opt/src/jdk-8u221-linux-x64.tar.gz http://down.sunrisenan.com/oracle/jdk-8u221-linux-x64.tar.gz
[root@shkf6-241 ~]# ls -l /opt/src/  | grep jdk
-rw-r--r-- 1 root root 195094741 Nov 28 10:44 jdk-8u221-linux-x64.tar.gz
[root@shkf6-241 ~]# mkdir /usr/java
[root@shkf6-241 ~]# tar xf /opt/src/jdk-8u221-linux-x64.tar.gz -C /usr/java
[root@shkf6-241 ~]# ln -s /usr/java/jdk1.8.0_221 /usr/java/jdk
[root@shkf6-241 ~]# vi /etc/profile
[root@shkf6-241 ~]# tail -4 /etc/profile

export JAVA_HOME=/usr/java/jdk
export PATH=$JAVA_HOME/bin:$JAVA_HOME/bin:$PATH
export CLASSPATH=$CLASSPATH:$JAVA_HOME/lib:$JAVA_HOME/lib/tools.jar
[root@shkf6-241 ~]# source /etc/profile

[root@shkf6-241 ~]# java -version
java version "1.8.0_221"
Java(TM) SE Runtime Environment (build 1.8.0_221-b11)
Java HotSpot(TM) 64-Bit Server VM (build 25.221-b11, mixed mode)
```

注意：这里以shkf6-241为例，分别在shkf6-242，shkf6-243上部署

## 3.安装zookeeper（3台zk角色主机）

[zk下载](https://mirrors.tuna.tsinghua.edu.cn/apache/zookeeper/)

[zookeeper](http://archive.apache.org/dist/zookeeper/)

### 1.解压配置

```shell
[root@shkf6-241 ~]# wget -O /opt/src/zookeeper-3.4.14.tar.gz https://mirrors.tuna.tsinghua.edu.cn/apache/zookeeper/zookeeper-3.4.14/zookeeper-3.4.14.tar.gz
[root@shkf6-241 ~]# tar xf /opt/src/zookeeper-3.4.14.tar.gz -C /opt/
[root@shkf6-241 ~]# ln -s /opt/zookeeper-3.4.14 /opt/zookeeper
[root@shkf6-241 ~]# mkdir -pv /data/zookeeper/data /data/zookeeper/logs
[root@shkf6-241 ~]# vi /opt/zookeeper/conf/zoo.cfg
[root@shkf6-241 ~]# cat /opt/zookeeper/conf/zoo.cfg
tickTime=2000
initLimit=10
syncLimit=5
dataDir=/data/zookeeper/data
dataLogDir=/data/zookeeper/logs
clientPort=2181
server.1=zk1.od.com:2888:3888
server.2=zk2.od.com:2888:3888
server.3=zk3.od.com:2888:3888
```

注意：各节点zk配置相同

### 2.myid

hdsh6-241.host.com上：

```shell
[root@shkf6-241 ~]# echo "1" > /data/zookeeper/data/myid
```

hdsh6-242.host.com上：

```shell
[root@shkf6-242 ~]# echo "2" > /data/zookeeper/data/myid
```

hdsh6-243.host.com上：

```shell
[root@shkf6-243 ~]# echo "3" > /data/zookeeper/data/myid
```

### 3.做dns解析

hdsh6-241.host.com上：

```shell
[root@shkf6-241 ~]# vi /var/named/od.com.zone 
[root@shkf6-241 ~]# cat /var/named/od.com.zone
$ORIGIN od.com.
$TTL 600    ; 10 minutes
@           IN SOA    dns.od.com. dnsadmin.od.com. (
                2019111209 ; serial
                10800      ; refresh (3 hours)
                900        ; retry (15 minutes)
                604800     ; expire (1 week)
                86400      ; minimum (1 day)
                )
                NS   dns.od.com.
$TTL 60    ; 1 minute
dns                A    192.168.6.241
harbor             A    192.168.6.245
k8s-yaml           A    192.168.6.245
traefik            A    192.168.6.66
dashboard          A    192.168.6.66
zk1                A    192.168.6.241
zk2                A    192.168.6.242
zk3                A    192.168.6.243

[root@shkf6-241 ~]# systemctl restart named.service

[root@shkf6-241 ~]# dig -t A zk1.od.com @192.168.6.241 +short
192.168.6.241
```

### 4.依次启动

```shell
[root@shkf6-241 ~]# /opt/zookeeper/bin/zkServer.sh start

[root@shkf6-242 ~]# /opt/zookeeper/bin/zkServer.sh start

[root@shkf6-243 ~]# /opt/zookeeper/bin/zkServer.sh start
```

### 5.常用命令

- 查看当前角色

```shell
[root@shkf6-241 ~]# /opt/zookeeper/bin/zkServer.sh status
ZooKeeper JMX enabled by default
Using config: /opt/zookeeper/bin/../conf/zoo.cfg
Mode: follower

[root@shkf6-242 ~]# /opt/zookeeper/bin/zkServer.sh status
ZooKeeper JMX enabled by default
Using config: /opt/zookeeper/bin/../conf/zoo.cfg
Mode: leader

[root@shkf6-243 ~]# /opt/zookeeper/bin/zkServer.sh status
ZooKeeper JMX enabled by default
Using config: /opt/zookeeper/bin/../conf/zoo.cfg
Mode: follower
```

# 第四章：部署jenkins

## 1.准备镜像

[jenkins官网](https://jenkins.io/download/)

[jenkins镜像](https://hub.docker.com/_/jenkins)

在运维主机下载官网上的稳定版（这里下载2.190.3）

```shell
[root@shkf6-245 ~]# docker pull jenkins/jenkins:2.190.3
[root@shkf6-245 ~]# docker images | grep jenkins
jenkins/jenkins                                   2.190.3                    22b8b9a84dbe        7 days ago          568MB
[root@shkf6-245 ~]# docker tag 22b8b9a84dbe harbor.od.com/public/jenkins:v2.190.3
[root@shkf6-245 ~]# docker pull !$
docker push harbor.od.com/public/jenkins:v2.190.3
```

## 2.自定义Dockerfile

在运维主机shkf6-245.host.com上：

```shell
[root@shkf6-245 ~]# mkdir -p  /data/dockerfile/jenkins/
[root@shkf6-245 ~]# vi /data/dockerfile/jenkins/Dockerfile
[root@shkf6-245 ~]# cat /data/dockerfile/jenkins/Dockerfile
FROM harbor.od.com/public/jenkins:v2.190.3
USER root
RUN /bin/cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime &&\ 
    echo 'Asia/Shanghai' >/etc/timezone
ADD id_rsa /root/.ssh/id_rsa
ADD config.json /root/.docker/config.json
ADD get-docker.sh /get-docker.sh
RUN echo "    StrictHostKeyChecking no" >> /etc/ssh/ssh_config &&\
    /get-docker.sh
```

- get-docker加速版

```shell
FROM harbor.od.com/public/jenkins:v2.190.3
USER root
RUN /bin/cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime &&\ 
   echo 'Asia/Shanghai' >/etc/timezone
ADD id_rsa /root/.ssh/id_rsa
ADD config.json /root/.docker/config.json
ADD get-docker.sh /get-docker.sh
RUN echo "    StrictHostKeyChecking no" >> /etc/ssh/ssh_config &&\
   /get-docker.sh --mirror Aliyun   # 阿里云加速
```

这个Dockerfile里我们主要做了以下几件事

- 设置容器用户为root
- 设置容器内的时区
- 将ssh私钥加入（使用git拉取代码时要用到，配置的公钥应配置在gitlab中）
- 加入了登录自建harbor仓库的config文件
- 修改了ssh客户端的配置
- 安装一个docker的客户端

√ 1.生成ssh秘钥：

```shell
[root@shkf6-245 ~]# ssh-keygen -t rsa -b 2048 -C "yanzhao.li@qq.com" -N "" -f /root/.ssh/id_rsa

[root@shkf6-245 ~]# cat /root/.ssh/id_rsa.pub   #可以看到自己设置的邮箱
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDzRHGRCF3F/IaI5EMwbZJ5V0AFJDogQVUeWEGLiqskyOVhoVAM/mPRXzNXPz/CkMKOkclCt/gPUWYgowVqFFnBobacVCmTATSdp0CDYhEjB54LAeTuOrbXb4uB957LlLRdiM3gsLtmjYxbs5dNRCGHZ4dXJ729nwAUofMkH+duVuN4OZ2GqNBz4ZCStgTOsM/vcyUex/N/mfET+ZLJO6+gLN0WzhjjmrynKueDXRsFSC+qHVIEi1WWHpGkr6sXX5FXIoviBQk8wJiFLvfEtjILDRMKxIMi3/uZeDrHKP4/9wGfu6OgLFKXWYsQByKnzIp3LsRZoI3EjGy6nx/VgnGZ yanzhao.li@qq.com
```

√ 2.拷贝文件

```shell
[root@shkf6-245 ~]# cp /root/.ssh/id_rsa /data/dockerfile/jenkins/

[root@shkf6-245 ~]# cp /root/.docker/config.json /data/dockerfile/jenkins/

[root@shkf6-245 ~]# cd /data/dockerfile/jenkins/ && curl -fsSL get.docker.com -o get-docker.sh && chmod +x get-docker.sh
```

√ 3.查看docker harbor config

```shell
[root@shkf6-245 jenkins]#cat /root/.docker/config.json
{
    "auths": {
        "harbor.od.com": {
            "auth": "YWRtaW46SGFyYm9yMTIzNDU="
        },
        "https://index.docker.io/v1/": {
            "auth": "c3VucmlzZW5hbjpseXo1MjA="
        }
    },
    "HttpHeaders": {
        "User-Agent": "Docker-Client/19.03.5 (linux)"
    }
}
```

## 3.制作自定义镜像

/data/dockerfile/jenkins

```shell
[root@shkf6-245 jenkins]# ls -l
total 28
-rw------- 1 root root   229 Nov 28 13:50 config.json
-rw-r--r-- 1 root root   394 Nov 28 13:15 Dockerfile
-rwxr-xr-x 1 root root 13216 Nov 28 13:53 get-docker.sh
-rw------- 1 root root  1679 Nov 28 13:40 id_rsa


[root@shkf6-245 jenkins]# docker build . -t harbor.od.com/infra/jenkins:v2.190.3
Sending build context to Docker daemon  19.46kB
Step 1/7 : FROM harbor.od.com/public/jenkins:v2.190.3
 ---> 22b8b9a84dbe
Step 2/7 : USER root
 ---> Running in 6347ef23acfd
Removing intermediate container 6347ef23acfd
 ---> ff18352d230e
Step 3/7 : RUN /bin/cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime &&    echo 'Asia/Shanghai' >/etc/timezone
 ---> Running in 970da85d013e
Removing intermediate container 970da85d013e
 ---> ca63098fe359
Step 4/7 : ADD id_rsa /root/.ssh/id_rsa
 ---> 0274b5facac2
Step 5/7 : ADD config.json /root/.docker/config.json
 ---> 75d0e57592c3
Step 6/7 : ADD get-docker.sh /get-docker.sh
 ---> a0ec7cf884a4
Step 7/7 : RUN echo "    StrictHostKeyChecking no" >> /etc/ssh/ssh_config &&    /get-docker.sh
 ---> Running in cd18e5417de5
# Executing docker install script, commit: f45d7c11389849ff46a6b4d94e0dd1ffebca32c1
+ sh -c apt-get update -qq >/dev/null
+ sh -c DEBIAN_FRONTEND=noninteractive apt-get install -y -qq apt-transport-https ca-certificates curl >/dev/null
debconf: delaying package configuration, since apt-utils is not installed
+ sh -c curl -fsSL "https://download.docker.com/linux/debian/gpg" | apt-key add -qq - >/dev/null
Warning: apt-key output should not be parsed (stdout is not a terminal)
+ sh -c echo "deb [arch=amd64] https://download.docker.com/linux/debian stretch stable" > /etc/apt/sources.list.d/docker.list
+ sh -c apt-get update -qq >/dev/null
+ [ -n  ]
+ sh -c apt-get install -y -qq --no-install-recommends docker-ce >/dev/null
debconf: delaying package configuration, since apt-utils is not installed
If you would like to use Docker as a non-root user, you should now consider
adding your user to the "docker" group with something like:

  sudo usermod -aG docker your-user

Remember that you will have to log out and back in for this to take effect!

WARNING: Adding a user to the "docker" group will grant the ability to run
         containers which can be used to obtain root privileges on the
         docker host.
         Refer to https://docs.docker.com/engine/security/security/#docker-daemon-attack-surface
         for more information.
Removing intermediate container cd18e5417de5
 ---> 7170e12fccfe
Successfully built 7170e12fccfe
Successfully tagged harbor.od.com/infra/jenkins:v2.190.3
```

## 4.创建infra仓库

在Harbor页面，创建infra仓库，注意：私有仓库

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_fa310a840919e2af2f8bfbf58e9203ab_r.png)

## 5.推送镜像

```shell
[root@shkf6-245 jenkins]# docker push harbor.od.com/infra/jenkins:v2.190.3
```

√ gitee.com 添加私钥，测试jenkins镜像：

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_8c8788f678c45b04375937a4c13290d5_r.png)

```shell
[root@shkf6-245 jenkins]# docker run --rm harbor.od.com/infra/jenkins:v2.190.3 ssh -i /root/.ssh/id_rsa  -T git@gitee.com
Warning: Permanently added 'gitee.com,212.64.62.174' (ECDSA) to the list of known hosts.
Hi Sunrise! You've successfully authenticated, but GITEE.COM does not provide shell access.
```

## 6.创建kubernetes命名空间，私有仓库鉴权

在任意运算节点上：

```shell
[root@shkf6-243 ~]# kubectl create ns infra
namespace/infra created

[root@shkf6-243 ~]# kubectl create secret docker-registry harbor --docker-server=harbor.od.com --docker-username=admin --docker-password=Harbor12345 -n infra
secret/harbor created
```

## 7.准备共享存储

运维主机，以及所有运算节点上：

```shell
[root@shkf6-243 ~]# yum install nfs-utils -y

[root@shkf6-244 ~]# yum install nfs-utils -y

[root@shkf6-245 ~]# yum install nfs-utils -y
```

- 配置NFS服务

运维主机shkf6-245上：

```shell
[root@shkf6-245 ~]# cat /etc/exports
/data/nfs-volume 192.168.6.0/24(rw,no_root_squash)
```

- 启动NFS服务

运维主机shkf6-245上：

```shell
[root@shkf6-245 ~]# mkdir -p  /data/nfs-volume/jenkins_home
[root@shkf6-245 ~]# systemctl start nfs
[root@shkf6-245 ~]# systemctl enable nfs
```

## 8.准备资源配置清单

运维主机shkf6-245上：

```shell
[root@shkf6-245 ~]# mkdir /data/k8s-yaml/jenkins
```

- Deployment

```shell
[root@shkf6-245 ~]# cat /data/k8s-yaml/jenkins/dp.yaml 
kind: Deployment
apiVersion: extensions/v1beta1
metadata:
  name: jenkins
  namespace: infra
  labels: 
    name: jenkins
spec:
  replicas: 1
  selector:
    matchLabels: 
      name: jenkins
  template:
    metadata:
      labels: 
        app: jenkins 
        name: jenkins
    spec:
      volumes:
      - name: data
        nfs: 
          server: shkf6-245
          path: /data/nfs-volume/jenkins_home
      - name: docker
        hostPath: 
          path: /run/docker.sock
          type: ''
      containers:
      - name: jenkins
        image: harbor.od.com/infra/jenkins:v2.190.3
        imagePullPolicy: IfNotPresent
        ports:
        - containerPort: 8080
          protocol: TCP
        env:
        - name: JAVA_OPTS
          value: -Xmx512m -Xms512m
        volumeMounts:
        - name: data
          mountPath: /var/jenkins_home
        - name: docker
          mountPath: /run/docker.sock
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

- service

```shell
[root@shkf6-245 ~]# cat /data/k8s-yaml/jenkins/svc.yaml 
kind: Service
apiVersion: v1
metadata: 
  name: jenkins
  namespace: infra
spec:
  ports:
  - protocol: TCP
    port: 80
    targetPort: 8080
  selector:
    app: jenkins
```

- ingress

```shell
[root@shkf6-245 ~]# cat /data/k8s-yaml/jenkins/ingress.yaml 
kind: Ingress
apiVersion: extensions/v1beta1
metadata: 
  name: jenkins
  namespace: infra
spec:
  rules:
  - host: jenkins.od.com
    http:
      paths:
      - path: /
        backend: 
          serviceName: jenkins
          servicePort: 80
```

## 9.应用资源配置清单

在任意运算节点上：

```shell
[root@shkf6-243 ~]# kubectl apply -f http://k8s-yaml.od.com/jenkins/dp.yaml

[root@shkf6-243 ~]# kubectl apply -f http://k8s-yaml.od.com/jenkins/svc.yaml

[root@shkf6-243 ~]# kubectl apply -f http://k8s-yaml.od.com/jenkins/ingress.yaml
```

- 检查

```shell
[root@shkf6-243 ~]# kubectl apply -f http://k8s-yaml.od.com/jenkins/dp.yaml
deployment.extensions/jenkins created
[root@shkf6-243 ~]# kubectl get all -n infra
NAME                           READY   STATUS    RESTARTS   AGE
pod/jenkins-74f7d66687-gjgth   1/1     Running   0          56m

NAME              TYPE        CLUSTER-IP    EXTERNAL-IP   PORT(S)   AGE
service/jenkins   ClusterIP   10.96.2.239   <none>        80/TCP    63m

NAME                      READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/jenkins   1/1     1            1           56m

NAME                                 DESIRED   CURRENT   READY   AGE
replicaset.apps/jenkins-74f7d66687   1         1         1       56m
```

## 10.解析域名

在shkf6-241上：

- 增加配置

```shell
[root@shkf6-241 ~]# cat /var/named/od.com.zone 
$ORIGIN od.com.
$TTL 600    ; 10 minutes
@           IN SOA    dns.od.com. dnsadmin.od.com. (
                2019111210 ; serial    # 滚动加一
                10800      ; refresh (3 hours)
                900        ; retry (15 minutes)
                604800     ; expire (1 week)
                86400      ; minimum (1 day)
                )
                NS   dns.od.com.
$TTL 60    ; 1 minute
dns                A    192.168.6.241
harbor             A    192.168.6.245
k8s-yaml           A    192.168.6.245
traefik            A    192.168.6.66
dashboard          A    192.168.6.66
zk1                A    192.168.6.241
zk2                A    192.168.6.242
zk3                A    192.168.6.243
jenkins            A    192.168.6.66       # 添加解析
```

- 重启，检查

```shell
[root@shkf6-241 ~]# systemctl restart named
[root@shkf6-241 ~]# dig -t A jenkins.od.com @192.168.6.241 +short
192.168.6.66
```

## 11.配置jenkins加速

- jenkins插件清华大学镜像地址

  ```shell
  [root@shkf6-245 ~]# wget -O /data/nfs-volume/jenkins_home/updates/default.json https://mirrors.tuna.tsinghua.edu.cn/jenkins/updates/update-center.json
  ```

- 其他方法

操作步骤

以上的配置Json其实在Jenkins的工作目录中

```shell
$ cd {你的Jenkins工作目录}/updates  #进入更新配置位置
```

第一种方式：使用vim

```shell
$ vim default.json   #这个Json文件与上边的配置文件是相同的

这里wiki和github的文档不用改，我们就可以成功修改这个配置

使用vim的命令，如下，替换所有插件下载的url

:1,$s/http:\/\/updates.jenkins-ci.org\/download/https:\/\/mirrors.tuna.tsinghua.edu.cn\/jenkins/g

替换连接测试url

:1,$s/http:\/\/www.google.com/https:\/\/www.baidu.com/g

    进入vim先输入：然后再粘贴上边的：后边的命令，注意不要写两个冒号！

修改完成保存退出:wq
```

第二种方式：使用sed

```shell
$ sed -i 's/http:\/\/updates.jenkins-ci.org\/download/https:\/\/mirrors.tuna.tsinghua.edu.cn\/jenkins/g' default.json && sed -i 's/http:\/\/www.google.com/https:\/\/www.baidu.com/g' default.json

    这是直接修改的配置文件，如果前边Jenkins用sudo启动的话，那么这里的两个sed前均需要加上sudo
```

[重启Jenkins，安装插件试试，简直超速](https://www.cnblogs.com/hellxz/p/jenkins_install_plugins_faster.html)！！

## 12.浏览器访问

浏览器访问 http://jenkins.od.com/

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_9b05da934bea1cc9b12790e4f06ba227_r.png)

## 13.页面配置jenkins

### 1.初始化密码

```shell
[root@shkf6-243 ~]# kubectl get pods -n infra
NAME                       READY   STATUS    RESTARTS   AGE
jenkins-74f7d66687-gjgth   1/1     Running   0          68m
[root@shkf6-243 ~]# kubectl exec jenkins-74f7d66687-gjgth /bin/cat /var/jenkins_home/secrets/initialAdminPassword -n infra 
59be7fd64b2b4c18a3cd927e0123f609


[root@shkf6-245 ~]# cat /data/nfs-volume/jenkins_home/secrets/initialAdminPassword 
59be7fd64b2b4c18a3cd927e0123f609
```

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_4079ad580f9908779fad463a361da95e_r.png)

### 2.跳过安装插件

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_142d1541420384969ad4a2bf66cc5c8a_r.png)

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_a33ab46e65c3d187b7443cd43a4cb8d6_r.png)

### 3.更改admin密码

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_ad3ee426b7710d427a4c2bfc652e4212_r.png)

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_fb066a63442d974943d2dc9d30dc0e6e_r.png)

### 4.使用admin登录

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_a373a959e064dce76be35505b4c88af1_r.png)

### 5.调整安全选项

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_b254e056cf8d6ccd78e8b862a36fcd97_r.png)

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_8cbb70f4591522b61501ad019b544c53_r.png)

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_dfbaa5b108f88df82246abf817183507_r.png)

### 6.安装Blue Ocean插件

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_0c43496501b1ffbc81e0de1225829eb0_r.png)

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_50c51fccc496c23857fc0d00efbe49b4_r.png)

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_111097df8cd6513890005c24dee704b6_r.png)

> 我们勾上这个允许匿名登录主要也是配合最后spinnaker

如果不允许匿名访问可进行如下操作：

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_6dc2fe58f288ecb137322e5da9f854c6_r.png)

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_8c4e6c8b6d1c0155b7859960e5dfc69c_r.png)

## 14.配置New job

- create new jobs
- Enter anitem name

> dubbo-demo

- Pipeline –> ok
- Discard old builds

> Days to keep builds：3
> Max # of builds to keep:30

- This project is parameterized

1.Add Parameter –> String Parameter

> Name：app_name
> Default Value:
> Description：project name，e.g：dubbo-demo-service

2.Add Parameter -> String Parameter

> Name : image_name
> Default Value :
> Description : project docker image name. e.g: app/dubbo-demo-service

3.Add Parameter -> String Parameter

> Name : git_repo
> Default Value :
> Description : project git repository. e.g: https://gitee.com/stanleywang/dubbo-demo-service.git

4.Add Parameter -> String Parameter

> Name : git_ver
> Default Value :
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
> Default Value : ./target
> Description : the relative path of target file such as .jar or .war package. e.g: ./dubbo-server/target

8.Add Parameter -> String Parameter

> Name : mvn_cmd
> Default Value : mvn clean package -Dmaven.test.skip=true
> Description : maven command. e.g: mvn clean package -e -q -Dmaven.test.skip=true

9.Add Parameter -> Choice Parameter

> Name : base_image
> Default Value :
>
> - base/jre7:7u80
> - base/jre8:8u112
>   Description : project base image list in harbor.od.com.

10.Add Parameter -> Choice Parameter

> Name : maven
> Default Value :
>
> - 3.6.0-8u181
> - 3.2.5-6u025
> - 2.2.1-6u025
>   Description : different maven edition.

## 15.Pipeline Script

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
      stage('package') { //move jar file into project_dir
        steps {
          sh "cd ${params.app_name}/${env.BUILD_NUMBER} && cd ${params.target_dir} && mkdir project_dir && mv *.jar ./project_dir"
        }
      }
      stage('image') { //build image and push to registry
        steps {
          writeFile file: "${params.app_name}/${env.BUILD_NUMBER}/Dockerfile", text: """FROM harbor.od.com/${params.base_image}
ADD ${params.target_dir}/project_dir /opt/project_dir"""
          sh "cd  ${params.app_name}/${env.BUILD_NUMBER} && docker build -t harbor.od.com/${params.image_name}:${params.git_ver}_${params.add_tag} . && docker push harbor.od.com/${params.image_name}:${params.git_ver}_${params.add_tag}"
        }
      }
    }
}
```

# 4.最后的准备工作

## 1.检查jenkins容器里的docker客户端

进入jenkins的docker容器里，检查docker客户端是否可用。

```shell
[root@shkf6-243 ~]# kubectl get pods -n infra 
NAME                       READY   STATUS    RESTARTS   AGE
jenkins-74f7d66687-6hdr7   1/1     Running   0          4d22h
[root@shkf6-243 ~]# kubectl exec -it jenkins-74f7d66687-6hdr7 /bin/sh -n infra 
# exit
[root@shkf6-243 ~]# kubectl exec -it jenkins-74f7d66687-6hdr7 bash -n infra 
root@jenkins-74f7d66687-6hdr7:/# docker ps -a
CONTAINER ID        IMAGE                               COMMAND                  CREATED             STATUS              PORTS                NAMES
96cc0389be29        7170e12fccfe                        "/sbin/tini -- /usr/…"   4 days ago          Up 4 days                                k8s_jenkins_jenkins-74f7d66687-6hdr7_infra_09e864de-341a-4a7d-a773-3803e19f428e_0
3bb2b7530c2c        harbor.od.com/public/pause:latest   "/pause"                 4 days ago          Up 4 days                                k8s_POD_jenkins-74f7d66687-6hdr7_infra_09e864de-341a-4a7d-a773-3803e19f428e_0
95c0c0485530        0c60bcf89900                        "/dashboard --insecu…"   5 days ago          Up 5 days                                k8s_kubernetes-dashboard_kubernetes-dashboard-5dbdd9bdd7-dtm98_kube-system_9a0475f5-2f02-4fac-bab1-ae295d4808c2_0
1d659b7beb93        harbor.od.com/public/pause:latest   "/pause"                 5 days ago          Up 5 days                                k8s_POD_kubernetes-dashboard-5dbdd9bdd7-dtm98_kube-system_9a0475f5-2f02-4fac-bab1-ae295d4808c2_0
598726a6347f        add5fac61ae5                        "/entrypoint.sh --ap…"   5 days ago          Up 5 days                                k8s_traefik-ingress_traefik-ingress-whtw9_kube-system_6ac78a23-81e9-48d0-a424-df2012e0ae2e_0
4d04878ff060        harbor.od.com/public/pause:latest   "/pause"                 5 days ago          Up 5 days           0.0.0.0:81->80/tcp   k8s_POD_traefik-ingress-whtw9_kube-system_6ac78a23-81e9-48d0-a424-df2012e0ae2e_0
root@jenkins-74f7d66687-6hdr7:/# 
```

## 2.检查jenkins容器里的SSH key

```shell
root@jenkins-74f7d66687-6hdr7:/# ssh -i /root/.ssh/id_rsa -T git@gitee.com
Warning: Permanently added 'gitee.com,212.64.62.174' (ECDSA) to the list of known hosts.
Hi StanleyWang (DeployKey)! You've successfully authenticated, but GITEE.COM does not provide shell access.
Note: Perhaps the current use is DeployKey.
Note: DeployKey only supports pull/fetch operations
```

## 3.部署maven软件

maven官方下载地址：

[maven3](https://archive.apache.org/dist/maven/maven-3/)
[maven2](https://archive.apache.org/dist/maven/maven-2/)
[maven1](https://archive.apache.org/dist/maven/maven-1/)

在运维主机shkf6-245.host.com上二进制部署，这里部署maven-3.6.1版

/opt/src

```shell
[root@shkf6-245 src]# wget https://archive.apache.org/dist/maven/maven-3/3.6.1/binaries/apache-maven-3.6.1-bin.tar.gz
[root@shkf6-245 src]# ls -l
total 8924
-rw-r--r-- 1 root root 9136463 Sep  4 00:54 apache-maven-3.6.1-bin.tar.gz
[root@shkf6-245 src]# mkdir /data/nfs-volume/jenkins_home/maven-3.6.1-8u232     # 8u232是jenkins中java的版本
[root@shkf6-245 src]# tar xf apache-maven-3.6.1-bin.tar.gz -C /data/nfs-volume/jenkins_home/maven-3.6.1-8u232
[root@shkf6-245 src]# cd /data/nfs-volume/jenkins_home/maven-3.6.1-8u232

[root@shkf6-245 maven-3.6.1-8u232]# mv apache-maven-3.6.1 ../
[root@shkf6-245 maven-3.6.1-8u232]# mv ../apache-maven-3.6.1/* .
```

- 设置国内镜像源

```shell
[root@shkf6-245 ~]# vi /data/nfs-volume/jenkins_home/maven-3.6.1-8u232/conf/settings.xml
<mirror>
  <id>alimaven</id>
  <name>aliyun maven</name>
  <url>http://maven.aliyun.com/nexus/content/groups/public/</url>
  <mirrorOf>central</mirrorOf>        
</mirror>
```

实例：

```shell
    146   <mirrors>
    147     <!-- mirror
    148      | Specifies a repository mirror site to use instead of a given repository. The repository that
    149      | this mirror serves has an ID that matches the mirrorOf element of this mirror. IDs are used
    150      | for inheritance and direct lookup purposes, and must be unique across the set of mirrors.
    151      |
    152     <mirror>
    153       <id>alimaven</id>
    154       <name>aliyun maven</name>
    155       <url>http://maven.aliyun.com/nexus/content/groups/public/</url>
    156       <mirrorOf>central</mirrorOf>
    157     </mirror>
    158     <mirror>
    159       <id>mirrorId</id>
    160       <mirrorOf>repositoryId</mirrorOf>
    161       <name>Human Readable Name for this Mirror.</name>
    162       <url>http://my.repository.com/repo/path</url>
    163     </mirror>
    164      -->
    165   </mirrors>
```

其他版本略

## 2.制作dubbo微服务的底包镜像

在运维主机shkf6-245.host.com上

1. 下载底包

```shell
[root@shkf6-245 jre8]# docker pull sunrisenan/jre8:8u112

[root@shkf6-245 jre8]# docker images|grep jre
sunrisenan/jre8                                   8u112                      fa3a085d6ef1        2 years ago         363MB

[root@shkf6-245 jre8]# docker tag fa3a085d6ef1 harbor.od.com/public/jre:8u112
[root@shkf6-245 jre8]# docker push harbor.od.com/public/jre:8u112
```

1. 自定义Dockerfile

- Dockerfile

```shell
[root@shkf6-245 jre8]# pwd
/data/dockerfile/jre8

[root@shkf6-245 jre8]# cat Dockerfile 
FROM harbor.od.com/public/jre:8u112
RUN /bin/cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime &&\
    echo 'Asia/Shanghai' >/etc/timezone
ADD config.yml /opt/prom/config.yml
ADD jmx_javaagent-0.3.1.jar /opt/prom/
WORKDIR /opt/project_dir
ADD entrypoint.sh /entrypoint.sh
CMD ["/entrypoint.sh"]
```

- config.yml

```shell
[root@shkf6-245 jre8]# cat config.yml 
---
rules:
  - pattern: '.*'

[root@shkf6-245 jre8]# wget https://repo1.maven.org/maven2/io/prometheus/jmx/jmx_prometheus_javaagent/0.3.1/jmx_prometheus_javaagent-0.3.1.jar -O jmx_javaagent-0.3.1.jar
```

- jmx_javaagent-0.3.1.jar

```shell
[root@shkf6-245 jre8]# wget https://repo1.maven.org/maven2/io/prometheus/jmx/jmx_prometheus_javaagent/0.3.1/jmx_prometheus_javaagent-0.3.1.jar -O jmx_javaagent-0.3.1.jar
```

- vi entrypoint.sh (不要忘了给执行权限)

```shell
[root@shkf6-245 jre8]# vi entrypoint.sh
[root@shkf6-245 jre8]# chmod +x entrypoint.sh
[root@shkf6-245 jre8]# cat entrypoint.sh 
#!/bin/sh
M_OPTS="-Duser.timezone=Asia/Shanghai -javaagent:/opt/prom/jmx_javaagent-0.3.1.jar=$(hostname -i):${M_PORT:-"12346"}:/opt/prom/config.yml"
C_OPTS=${C_OPTS}
JAR_BALL=${JAR_BALL}
exec java -jar ${M_OPTS} ${C_OPTS} ${JAR_BALL}
```

1. 制作dubbo服务docker底包

```shell
[root@shkf6-245 jre8]# pwd
/data/dockerfile/jre8
[root@shkf6-245 jre8]# ls -l
total 372
-rw-r--r-- 1 root root     29 Dec  4 09:50 config.yml
-rw-r--r-- 1 root root    297 Dec  4 09:49 Dockerfile
-rwxr-xr-x 1 root root    234 Dec  4 09:54 entrypoint.sh
-rw-r--r-- 1 root root 367417 May 10  2018 jmx_javaagent-0.3.1.jar

[root@shkf6-245 jre8]# docker build . -t harbor.od.com/base/jre8:8u112

[root@shkf6-245 jre8]# docker push harbor.od.com/base/jre8:8u112
```

注意：jre7底包制作类似，这里略

# 5.交付dubbo微服务至kubernetes集群

## 1.dubbo服务提供者（dubbo-demo-service）

### 1.通过jenkins进行一次CI

打开jenkins页面，使用admin登录，准备构建`dubbo-demo`项目

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_9454d95ccceca6db81095028a1a3e1fc_r.png)

点`Build with Parameters`

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_e5e8d9ac4a03cf7c428272a6f879dcd8_r.png)

依次填入/选择：

app_name

> dubbo-demo-service

image_name

> app/dubbo-demo-service

git_repo

> https://gitee.com/stanleywang/dubbo-demo-service.git

git_ver

> master

add_tag

> 191204_1942

mvn_dir

> ./

target_dir

> ./dubbo-server/target

mvn_cmd

> mvn clean package -Dmaven.test.skip=true

base_image

> base/jre8:8u112

maven

> 3.6.1-8u232

点击`Build`进行构建，等待构建完成。

test $? -eq 0 && 成功，进行下一步 || 失败，排错直到成功

### 2.检查harbor仓库内镜像

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_38df9d1af4bf5370d6112ecfad084aff_r.png)

### 3.准备k8s资源配置清单

运维主机shkf6-245.host.com上，准备资源配置清单：

```shell
[root@shkf6-245 ~]# mkdir /data/k8s-yaml/dubbo-demo-service/
[root@shkf6-245 ~]# cd /data/k8s-yaml/dubbo-demo-service/
[root@shkf6-245 dubbo-demo-service]# vi /data/k8s-yaml/dubbo-demo-service/deployment.yaml
[root@shkf6-245 dubbo-demo-service]# cat /data/k8s-yaml/dubbo-demo-service/deployment.yaml
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
        image: harbor.od.com/app/dubbo-demo-service:master_191204_1942
        ports:
        - containerPort: 20880
          protocol: TCP
        env:
        - name: JAR_BALL
          value: dubbo-server.jar
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

### 4.应用资源配置清单

在任意一台k8s运算节点执行：

- 创建kubernetes命名空间，私有仓库鉴权

  ```shell
  [root@shkf6-243 ~]# kubectl create ns app
  namespace/app created
  [root@shkf6-243 ~]# kubectl create secret docker-registry harbor --docker-server=harbor.od.com --docker-username=admin --docker-password=Harbor12345 -n app
  secret/harbor created
  ```

- 应用资源配置清单

```shell
[root@shkf6-243 ~]# kubectl apply -f http://k8s-yaml.od.com/dubbo-demo-service/deployment.yaml
deployment.extensions/dubbo-demo-service created
```

### 5.检查docker运行情况及zk里的信息

```
/opt/zookeeper/bin/zkCli.sh
[root@shkf6-243 ~]# /opt/zookeeper/bin/zkCli.sh -server localhost:2181

[zk: localhost:2181(CONNECTED) 0] ls /
[dubbo, zookeeper]

[zk: localhost:2181(CONNECTED) 1] ls /dubbo
[com.od.dubbotest.api.HelloService]
```

## 2.dubbo-monitor工具

[dubbo-monitor源码](https://github.com/Jeromefromcn/dubbo-monitor)

### 1.准备docker镜像

#### 1.下载源码并解压

下载到运维主机shkf6-245.host.com上

```shell
[root@shkf6-245 ~]# cd /opt/src/
[root@shkf6-245 src]# wget -O /opt/src/dubbo-monitor-master.zip http://down.sunrisenan.com/dubbo-monitor/dubbo-monitor-master.zip

[root@shkf6-245 src]# yum install unzip -y
[root@shkf6-245 src]# unzip dubbo-monitor-master.zip

[root@shkf6-245 src]# mv dubbo-monitor-master /data/dockerfile/dubbo-monitor
[root@shkf6-245 src]# cd /data/dockerfile/dubbo-monitor
```

#### 2.修改配置

- 修改dubbo-monitor主配置文件

```shell
[root@shkf6-245 dubbo-monitor]# vi dubbo-monitor-simple/conf/dubbo_origin.properties
[root@shkf6-245 dubbo-monitor]# cat dubbo-monitor-simple/conf/dubbo_origin.properties
dubbo.container=log4j,spring,registry,jetty
dubbo.application.name=dubbo-monitor
dubbo.application.owner=OldboyEdu
dubbo.registry.address=zookeeper://zk1.od.com:2181?backup=zk2.od.com:2181,zk3.od.com:2181
dubbo.protocol.port=20880
dubbo.jetty.port=8080
dubbo.jetty.directory=/dubbo-monitor-simple/monitor
dubbo.charts.directory=/dubbo-monitor-simple/charts
dubbo.statistics.directory=/dubbo-monitor-simple/statistics
dubbo.log4j.file=logs/dubbo-monitor-simple.log
dubbo.log4j.level=WARN
```

- 修改duboo-monitor启动脚本

```shell
[root@shkf6-245 dubbo-monitor]# sed -r -i -e '/^nohup/{p;:a;N;$!ba;d}'  ./dubbo-monitor-simple/bin/start.sh && sed  -r -i -e "s%^nohup(.*)%exec \1%"  ./dubbo-monitor-simple/bin/start.sh
    JAVA_MEM_OPTS=" -server -Xmx128g -Xms128g -Xmn32m -XX:PermSize=16m -Xss256k -XX:+DisableExplicitGC -XX:+UseConcMarkSweepGC -XX:+CMSParallelRemarkEnabled -XX:+UseCMSCompactAtFullCollection -XX:LargePageSizeInBytes=128m -XX:+UseFastAccessorMethods -XX:+UseCMSInitiatingOccupancyOnly -XX:CMSInitiatingOccupancyFraction=70 "
else
    JAVA_MEM_OPTS=" -server -Xms128g -Xmx128g -XX:PermSize=16m -XX:SurvivorRatio=2 -XX:+UseParallelGC "
fi
```

> 提示：记得最后的`&`符删除掉

示例：启动脚本完整配置

```shell
[root@shkf6-245 dubbo-monitor]# cat dubbo-monitor-simple/bin/start.sh
#!/bin/bash
sed -e "s/{ZOOKEEPER_ADDRESS}/$ZOOKEEPER_ADDRESS/g" /dubbo-monitor-simple/conf/dubbo_origin.properties > /dubbo-monitor-simple/conf/dubbo.properties
cd `dirname $0`
BIN_DIR=`pwd`
cd ..
DEPLOY_DIR=`pwd`
CONF_DIR=$DEPLOY_DIR/conf

SERVER_NAME=`sed '/dubbo.application.name/!d;s/.*=//' conf/dubbo.properties | tr -d '\r'`
SERVER_PROTOCOL=`sed '/dubbo.protocol.name/!d;s/.*=//' conf/dubbo.properties | tr -d '\r'`
SERVER_PORT=`sed '/dubbo.protocol.port/!d;s/.*=//' conf/dubbo.properties | tr -d '\r'`
LOGS_FILE=`sed '/dubbo.log4j.file/!d;s/.*=//' conf/dubbo.properties | tr -d '\r'`

if [ -z "$SERVER_NAME" ]; then
    SERVER_NAME=`hostname`
fi

PIDS=`ps -f | grep java | grep "$CONF_DIR" |awk '{print $2}'`
if [ -n "$PIDS" ]; then
    echo "ERROR: The $SERVER_NAME already started!"
    echo "PID: $PIDS"
    exit 1
fi

if [ -n "$SERVER_PORT" ]; then
    SERVER_PORT_COUNT=`netstat -tln | grep $SERVER_PORT | wc -l`
    if [ $SERVER_PORT_COUNT -gt 0 ]; then
        echo "ERROR: The $SERVER_NAME port $SERVER_PORT already used!"
        exit 1
    fi
fi

LOGS_DIR=""
if [ -n "$LOGS_FILE" ]; then
    LOGS_DIR=`dirname $LOGS_FILE`
else
    LOGS_DIR=$DEPLOY_DIR/logs
fi
if [ ! -d $LOGS_DIR ]; then
    mkdir $LOGS_DIR
fi
STDOUT_FILE=$LOGS_DIR/stdout.log

LIB_DIR=$DEPLOY_DIR/lib
LIB_JARS=`ls $LIB_DIR|grep .jar|awk '{print "'$LIB_DIR'/"$0}'|tr "\n" ":"`

JAVA_OPTS=" -Djava.awt.headless=true -Djava.net.preferIPv4Stack=true "
JAVA_DEBUG_OPTS=""
if [ "$1" = "debug" ]; then
    JAVA_DEBUG_OPTS=" -Xdebug -Xnoagent -Djava.compiler=NONE -Xrunjdwp:transport=dt_socket,address=8000,server=y,suspend=n "
fi
JAVA_JMX_OPTS=""
if [ "$1" = "jmx" ]; then
    JAVA_JMX_OPTS=" -Dcom.sun.management.jmxremote.port=1099 -Dcom.sun.management.jmxremote.ssl=false -Dcom.sun.management.jmxremote.authenticate=false "
fi
JAVA_MEM_OPTS=""
BITS=`java -version 2>&1 | grep -i 64-bit`
if [ -n "$BITS" ]; then
    JAVA_MEM_OPTS=" -server -Xmx128g -Xms128g -Xmn32m -XX:PermSize=16m -Xss256k -XX:+DisableExplicitGC -XX:+UseConcMarkSweepGC -XX:+CMSParallelRemarkEnabled -XX:+UseCMSCompactAtFullCollection -XX:LargePageSizeInBytes=128m -XX:+UseFastAccessorMethods -XX:+UseCMSInitiatingOccupancyOnly -XX:CMSInitiatingOccupancyFraction=70 "
else
    JAVA_MEM_OPTS=" -server -Xms128g -Xmx128g -XX:PermSize=16m -XX:SurvivorRatio=2 -XX:+UseParallelGC "
fi

echo -e "Starting the $SERVER_NAME ...\c"
exec  java $JAVA_OPTS $JAVA_MEM_OPTS $JAVA_DEBUG_OPTS $JAVA_JMX_OPTS -classpath $CONF_DIR:$LIB_JARS com.alibaba.dubbo.container.Main > $STDOUT_FILE 2>&1
```

#### 3.制作镜像

- 准备Dockerfile

```shell
[root@shkf6-245 dubbo-monitor]# cat Dockerfile 
FROM jeromefromcn/docker-alpine-java-bash
MAINTAINER Jerome Jiang
COPY dubbo-monitor-simple/ /dubbo-monitor-simple/
CMD /dubbo-monitor-simple/bin/start.sh
```

- build镜像

```shell
[root@shkf6-245 dubbo-monitor]# docker build . -t harbor.od.com/infra/dubbo-monitor:latest

[root@shkf6-245 dubbo-monitor]# docker push harbor.od.com/infra/dubbo-monitor:latest
```

### 2.解析域名

在DNS主机shkf6-241.hosts.com上：

```shell
[root@shkf6-241 ~]# tail -1 /var/named/od.com.zone 
dubbo-monitor      A    192.168.6.66
```

### 3.准备k8s资源配置清单

运维主机shkf6-245.host.com上:

- 创建目录

```shell
[root@shkf6-245 ~]# mkdir /data/k8s-yaml/dubbo-monitor
```

- Deployment

```shell
[root@shkf6-245 ~]# vi /data/k8s-yaml/dubbo-monitor/dp.yaml
[root@shkf6-245 ~]# cat /data/k8s-yaml/dubbo-monitor/dp.yaml
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

- server

```shell
[root@shkf6-245 ~]# vi /data/k8s-yaml/dubbo-monitor/svc.yaml
[root@shkf6-245 ~]# cat /data/k8s-yaml/dubbo-monitor/svc.yaml
kind: Service
apiVersion: v1
metadata: 
  name: dubbo-monitor
  namespace: infra
spec:
  ports:
  - protocol: TCP
    port: 8080
    targetPort: 8080
  selector: 
    app: dubbo-monitor
```

- ingress

```shell
[root@shkf6-245 ~]# vi /data/k8s-yaml/dubbo-monitor/ingress.yaml
[root@shkf6-245 ~]# cat /data/k8s-yaml/dubbo-monitor/ingress.yaml
kind: Ingress
apiVersion: extensions/v1beta1
metadata: 
  name: dubbo-monitor
  namespace: infra
spec:
  rules:
  - host: dubbo-monitor.od.com
    http:
      paths:
      - path: /
        backend: 
          serviceName: dubbo-monitor
          servicePort: 8080
```

### 4.应用资源配置清单

在任意一台k8s运算节点执行：

```shell
[root@shkf6-243 ~]# kubectl apply -f http://k8s-yaml.od.com/dubbo-monitor/dp.yaml
deployment.extensions/dubbo-monitor created
[root@shkf6-243 ~]# kubectl apply -f http://k8s-yaml.od.com/dubbo-monitor/svc.yaml
service/dubbo-monitor created
[root@shkf6-243 ~]# kubectl apply -f http://k8s-yaml.od.com/dubbo-monitor/ingress.yaml
ingress.extensions/dubbo-monitor created
```

### 5.浏览器访问

[http://dubbo-monitor.od.com](http://dubbo-monitor.od.com/)

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_b54fbf8a32850059c621325f5fab81ca_r.png)

## 2.dubbo服务消费者（dubbo-demo-consumer）

### 1.通过jenkins进行一次CI

打开jenkins页面，使用admin登录，准备构建dubbo-demo项目

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_593391e6e785e8b874f83a1fca2d588c_r.png)

点`Build with Parameters`

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_b8249fd0b079641297b6fa36ff189cea_r.png)

依次填入/选择：

app_name

> dubbo-demo-consumer

image_name

> app/dubbo-demo-consumer

git_repo

> [git@gitee.com](mailto:git@gitee.com):stanleywang/dubbo-demo-web.git

git_ver

> master

add_tag

> 191205_1908

mvn_dir

> ./

target_dir

> ./dubbo-client/target

mvn_cmd

> mvn clean package -e -q -Dmaven.test.skip=true

base_image

> base/jre8:8u112

maven

> 3.6.1-8u232

点击`Build`进行构建，等待构建完成。

test $? -eq 0 && 成功，进行下一步 || 失败，排错直到成功

### 2.检查harbor仓库内镜像

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_3eda79f023990bcca458c33cfbae815b_r.png)

### 3.解析域名

在DNS主机shkf6-241.host.com上：

```shell
[root@shkf6-241 ~]# tail -1 /var/named/od.com.zone 
demo               A    192.168.6.66
```

### 4.准备k8s资源配置清单

运维主机shkf6-245.host.com上，准备资源配置清单

- 创建目录

```shell
[root@shkf6-245 ~]# mkdir /data/k8s-yaml/dubbo-demo-consumer
```

- deployment

```shell
[root@shkf6-245 ~]# vi /data/k8s-yaml/dubbo-demo-consumer/dp.yaml
[root@shkf6-245 ~]# cat /data/k8s-yaml/dubbo-demo-consumer/dp.yaml
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
        image: harbor.od.com/app/dubbo-demo-consumer:master_191205_1908
        ports:
        - containerPort: 8080
          protocol: TCP
        - containerPort: 20880
          protocol: TCP
        env:
        - name: JAR_BALL
          value: dubbo-client.jar
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
[root@shkf6-245 ~]# vi /data/k8s-yaml/dubbo-demo-consumer/svc.yaml
[root@shkf6-245 ~]# cat /data/k8s-yaml/dubbo-demo-consumer/svc.yaml
kind: Service
apiVersion: v1
metadata: 
  name: dubbo-demo-consumer
  namespace: app
spec:
  ports:
  - protocol: TCP
    port: 8080
    targetPort: 8080
  selector: 
    app: dubbo-demo-consumer
```

- ingress

```shell
[root@shkf6-245 ~]# vi /data/k8s-yaml/dubbo-demo-consumer/ingress.yaml
[root@shkf6-245 ~]# cat /data/k8s-yaml/dubbo-demo-consumer/ingress.yaml
kind: Ingress
apiVersion: extensions/v1beta1
metadata: 
  name: dubbo-demo-consumer
  namespace: app
spec:
  rules:
  - host: demo.od.com
    http:
      paths:
      - path: /
        backend: 
          serviceName: dubbo-demo-consumer
          servicePort: 8080
```

### 5.应用资源配置清单

在任意一台k8s运算节点执行：

```shell
[root@shkf6-243 ~]# kubectl apply -f http://k8s-yaml.od.com/dubbo-demo-consumer/dp.yaml
deployment.extensions/dubbo-demo-consumer created
[root@shkf6-243 ~]# kubectl apply -f http://k8s-yaml.od.com/dubbo-demo-consumer/svc.yaml
service/dubbo-demo-consumer created
[root@shkf6-243 ~]# kubectl apply -f http://k8s-yaml.od.com/dubbo-demo-consumer/ingress.yaml
ingress.extensions/dubbo-demo-consumer created
```

### 6.检查docker运行情况及dubbo-monitor

[http://dubbo-monitor.od.com](http://dubbo-monitor.od.com/)

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_e60e27ccda4d04ac6e91b006bf63331a_r.png)

### 7.浏览器访问

http://demo.od.com/hello?name=sunrise

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_a515a2aec28885078f92a1e9192d9c95_r.png)

## 6.实战维护dubbo微服务集群

- 更新（rolling update）

  - 修改代码提git（发版）

  - 使用jenkins进行CI

  - 修改并应用k8s资源配置清单

    > 或者在k8s的dashboard上直接操作

- 扩容（scaling）

  - k8s的dashboard上直接操作

## 7.k8s灾难性毁灭测试

================

运行中的集群在某天挂了一台

```shell
[root@shkf6-243 ~]# halt
```

这时访问业务会有短暂的 Bad Gateway

================

1、K8S中移除坏的节点（这时会触发自愈机制）：

```shell
[root@shkf6-244 ~]# kubectl delete node shkf6-243.host.com
node "shkf6-243.host.com" deleted
```

2、这时需要判定负载均衡是否要移除节点
略。

3、机器修复完，自动加入集群，打标签

```shell
[root@shkf6-244 ~]# kubectl label node shkf6-243.host.com node-role.kubernetes.io/master=
node/shkf6-243.host.com labeled
[root@shkf6-244 ~]# kubectl label node shkf6-243.host.com node-role.kubernetes.io/node=
node/shkf6-243.host.com labeled
```

4、根据测试结果是要重启docker引擎的

```shell
[root@shkf6-243 ~]# systemctl restart docker 
```

5、跟据情况平衡POD负载

```shell
[root@shkf6-244 ~]# kubectl get pods -n app
NAME                                  READY   STATUS    RESTARTS   AGE
dubbo-demo-consumer-5668798c5-86g7w   1/1     Running   0          26m
dubbo-demo-consumer-5668798c5-p2n4f   1/1     Running   0          21h
dubbo-demo-service-b4fd94448-j5lfx    1/1     Running   0          26m
dubbo-demo-service-b4fd94448-jdtmd    1/1     Running   0          43h

[root@shkf6-244 ~]# kubectl get pods -n app -o wide
NAME                                  READY   STATUS    RESTARTS   AGE   IP             NODE                 NOMINATED NODE   READINESS GATES
dubbo-demo-consumer-5668798c5-86g7w   1/1     Running   0          26m   172.6.244.9    shkf6-244.host.com   <none>           <none>
dubbo-demo-consumer-5668798c5-p2n4f   1/1     Running   0          21h   172.6.244.7    shkf6-244.host.com   <none>           <none>
dubbo-demo-service-b4fd94448-j5lfx    1/1     Running   0          26m   172.6.244.10   shkf6-244.host.com   <none>           <none>
dubbo-demo-service-b4fd94448-jdtmd    1/1     Running   0          43h   172.6.244.5    shkf6-244.host.com   <none>           <none>

[root@shkf6-244 ~]# kubectl delete pod dubbo-demo-consumer-5668798c5-86g7w -n app
pod "dubbo-demo-consumer-5668798c5-86g7w" deleted
[root@shkf6-244 ~]# kubectl delete pod dubbo-demo-service-b4fd94448-j5lfx -n app
pod "dubbo-demo-service-b4fd94448-j5lfx" deleted
```

6、总结：

1、删除k8s坏的node节点，这时故障自愈
2、注释掉坏负载均衡器上坏节点ip
3、修复好机器加入集群
4、打标签，平衡节点pods