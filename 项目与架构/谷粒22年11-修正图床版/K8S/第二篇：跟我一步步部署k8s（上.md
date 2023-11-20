# 引子

**1.docker容器化封装应用的意义(好处)**

Docker引擎统一了基础设施环境 - docker环境

- 硬件的配置
- 操作系统的版本
- 运行时环境的异构

Docker引擎统一了程序打包(装箱)方式 - docker镜像

- java程序
- python程序
- nodejs程序
- …

Docker引擎统一了程序部署(运行)

- java -jar … -> docker run …
- python manage.py runserver … -> docker run …
- npm run dev -> docker run …

客户端/服务器 C/S

浏览器/服务器 B/S

移动端/服务器

小程序/服务器

**2.Docker容器化封装应用程序的缺点(坏处)**

- 单机使用，无法有效集群
- 随着容器数量的上升，管理成本攀升
- 没有有效的容灾\自愈机制
- 没有预设编排模板，无法实现快速、大规模容器调度
- 没有统一的配置管理中心工具
- 没有容器生命周期的管理工具
- 没有图形化运维管理工具 (个人开发的web界面)
  ….

因此我们需要一套容器编排工具！

**3.Docker容器引擎的开源容器编排工具目前市场上主要有：**

- docker compose、docker swarm
- Mesosphere + Marathon
- Kubernetes(K8S)

# 第一章：kubernetes概述

> 官网：[https://kubernetes.io](https://kubernetes.io/)
> GitHub：https://github.com/kubernetes/kubernetes
> 由来：谷歌的Borg系统，后经Go语言重写并捐给CNCF基金会开源
> 含义：词根源于希腊语：舵手/飞行员。K8S -> K12345678S
> 重要作用：开源的容器编排框架工具(生态极其丰富)
> 学习的意义：解决跑裸docker的若干痛点

kubernetes优势：

```shell
自动装箱，水平扩展，自我修复
服务发现和负载均衡
自动发布(默认滚动发布模式)和回滚
集中化配置管理和秘钥管理
存储编排
任务批量处理运行
....
```

# 第二章：kubernetes快速入门

## 1.四组基本概念

### 1.Pod/Pod控制器

Pod

```shell
Pod是K8S里能够被运行的做小的逻辑单元(原子单元)
1个Pod里面可以运行多个容器，他们共享UTS+NET+IPC名称空间
可以把Pod理解成豌豆荚，而同一个Pod内的每个容器是一颗颗豌豆
一个Pod里运行多个容器，又叫：边车(SideCar)模式
```

Pod控制器

```shell
Pod控制器是Pod启动的一种模板，用来保证在K8S里启动的Pod应始终按照人们的预期运行(副本数、生命周期、健康状态检查...)

K8S内提供了众多的Pod控制器，常用的有以下几种:
    Deployment
    DaemonSet
    ReplicaSet
    StatefulSet
    Job
    Cronjob
```

### 2.Name/Namespace

Name

```shell
由于K8S内部，使用“资源”来定义每一种逻辑概念(功能)故每种“资源”，都应该有自己的“名称”
“资源”有api版本(apiVersion)类别(Kind)、元数据(matadata)、定义清单(spec)、状态(status)等配置信息
“名称”通常定义在“资源”的“元数据”信息里
```

Namespace

```shell
随着项目增多、人员增加、集群规模的扩大，需要一种能够隔离K8S内各种“资源”的方法，这就是名称空间
名称空间可以理解为K8S内部的虚拟集群组
不同名称空间内的“资源”，名称可以相同，相同名称空间内的同种“资源”，“名称”不能相同
合理的使用K8S的名称空间，使得集群管理员能够更好的对交付到K8S里的服务进行分类管理和浏览
K8S里默认存在的名称空间有：default、kube-system、kube-public
查询K8S里特定“资源”要带上相应的名称空间
```

### 3.Label/Label选择器

Label

```shell
标签是k8s特色的管理方式，便于分类管理资源对象。
一个标签可以对应多个资源，一个资源也可以有多个标签，它们是多对多的关系。
一个资源拥有多个标签，可以实现不同维度的管理。
标签的组成：key=value
与标签类似的，还有一种“注解”(annotations)
```

Label选择器

```shell
给资源打上标签后，可以使用标签选择器过滤指定的标签
标签选择器目前有两个：基于等值关系(等于、不等于)和基于集合关系(属于、不属于，存在)
许多资源支持内嵌标签选择器字段
    matchLabels
    matchExpressions
```

### 4.Service/Ingress

Service

```shell
在K8S的世界里，虽然每个Pod都会被分配一个单独的IP地址，但这个IP地址会随着Pod的销毁而消失
Service(服务)就是用来解决这个问题的核心概念
一个Service可以看作一组提供相同服务的Pod的对外访问接口
Service作用于哪些Pod是通过标签选择器来定义的
```

Ingress

```shell
Ingress是K8S集群里工作在OSI网络参考模型下，第7层的应用，对外暴露的接口
Service只能进行L4流量调度，表现形式是ip+port
Ingress则可以调度不同业务域，不同URL访问路径的业务流量
```

## 2.K8S的组成

```shell
核心组件
    配置存储中心 -> etcd服务
    主控(Master)节点
        kube-apiserver服务
        kube-controller-manager服务
        kube-scheduler服务
    运算(node)节点
        kube-kubelet服务
        kube-proxy服务

CLI客户端
    kubectl

核心附件
    CNI网络插件 -> flannel/calico
    服务发现用插件 -> coredns
    服务暴露用插件 -> traefik
    GUI管理插件 -> Dashboard
```

apiserver：

```shell
    提供了集群管理的REST API接口(包括鉴权、数据校验及集群状态变更)
    负责其他模块之间的数据交互，承担通信枢纽功能
    是资源配额控制的入口
    提供完备的集群安全机制
```

controller-manager：

```shell
    由一系列控制器组成，通过apiserver监控整个集群的状态，并确保集群处于预期的工作状态
    Node Controller                    节点控制
    Deployment Controller           pod控制器
    Service Controller                服务控制器
    Volume Controller                   存储卷控制器
    Endpoint Controller             接入点控制器
    Garbage Controller              垃圾控制器
    Namespace Controller            名称空间控制器
    Job Controller                    任务控制器
    Resource quta Controller        资源配额控制器
    ...    
```

scheduler:

```shell
    主要功能是接收调度pod到适合的运算节点上
    预算策略(predict)
    优选策略(priorities)
```

kubelet:

```shell
    简单的说，kubelet的主要功能就是定时从某个地方获取节点上pod的期望状态(运行什么容器、运行副本数量、网络或者存储如何配置等等)，并调用对应的容器平台接口达到这个状态。
    定时汇报当前节点的状态给apiserver，以供调度的时候使用
    镜像和容器的清理工作，保证节点上的镜像不会占满磁盘空间，退出的容器不会占用太多资源
```

kube-proxy：

```shell
    是K8S在每个节点上运行网络代理，service资源的载体
    建立了pod网络和集群网络的关系(clusterip -> podip)
    常用的三种流量调度模式：
        Userspace(废弃)
        Iptables(濒临泛滥)
        Ipvs(推荐)

    负责建立和删除包括更新调度规则、通知apiserver自己的更新，或者从apiserver的调度规则变化来更新自己的。
```

## 3.K8S三条网络详解

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_9f8b90aeb7332c704fcad1ca417d1b6e_r.png)

# 第三章：实验部署集群架构详解

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_f5d8c5028b894e4a2c1ca111eb6b7d27_r.png)

常见的K8S安装部署方式：

```shell
1. Minikube单节点微型K8S(仅供学习、预览使用)
2. 二进制安装部署(生产首选，新手推荐)
3. 使用Kuberadmin进行部署，K8S的部署工具，跑在K8S里(相对简单，熟手推荐)
```

# 第四章：部署k8s集群前准备工作

### 1.准备虚拟机

- 5台vm，每台4C8G

| 主机名             | 角色                     | ip            | 配置 | 部署服务                                                     |
| :----------------- | :----------------------- | :------------ | :--- | :----------------------------------------------------------- |
| shkf6-241.host.com | k8s1代理节点1            | 192.168.6.241 | 4C8G | bind9,nginx(四层代理),keeplived,supervisor                   |
| shkf6-242.host.com | k8s1代理节点2            | 192.168.6.242 | 4C8G | etcd,nginx(四层代理),keeplived,supervisor                    |
| shkf6-243.host.com | k8s运算节点1             | 192.168.6.243 | 4C8G | etcd,kube-apiserver,kube-controller-manager,kube-scheduler,kube-kubelet,kube-proxy,supervisor |
| shkf6-244.host.com | k8s运算节点2             | 192.168.6.244 | 4C8G | etcd,kube-apiserver,kube-controller-manager,kube-scheduler,kube-kubelet,kube-proxy,supervisor |
| shkf6-245.host.com | k8s运维节点(docker仓库） | 192.168.6.245 | 4C8G | 证书服务，docker私有仓库harbor，nginx代理harbor，pause       |

查看系统版本

```shell
~]# uname -a
Linux shkf6-241.host.com 3.10.0-693.21.1.el7.x86_64 #1 SMP Wed Mar 7 19:03:37 UTC 2018 x86_64 x86_64 x86_64 GNU/Linux
```

### 2.调整操作系统

所有机器上：

1.设置主机名

```shell
~]# hostnamectl set-hostname shkf6-241.host.com
~]# hostnamectl set-hostname shkf6-242.host.com
~]# hostnamectl set-hostname shkf6-243.host.com
~]# hostnamectl set-hostname shkf6-244.host.com
~]# hostnamectl set-hostname shkf6-245.host.com
```

2.关闭selinux和关闭防火墙

```shell
~]# sed -i 's#SELINUX=enforcing#SELINUX=disabled#g' /etc/selinux/config
~]# setenforce 0

~]# systemctl stop firewalld
```

3.安装epel-release

```shell
~]# yum install -y epel-release
```

4.安装必工具

```shell
~]# yum install wget net-tools telnet tree nmap sysstat lrzsz dos2unix bind-utils -y
```

### 3.DNS服务初始化

shkf6-241.host.com上：

#### 1.安装bind9软件

```shell
[root@shkf6-241 ~]# yum install bind -y
=====================================================================================================================================
 Package                                  Arch                     Version                              Repository              Size
=====================================================================================================================================
Installing:
 bind                                     x86_64                   32:9.11.4-9.P2.el7                   base                   2.3 M
```

#### 2.配置bind9

主配置文件

```shell
[root@shkf6-241 ~]# vim /etc/named.conf

listen-on port 53 { 192.168.6.241; };
allow-query     { any; };
forwarders      { 192.168.6.254; };      #向上查询(增加一条)
dnssec-enable no;
dnssec-validation no;

[root@shkf6-241 ~]# named-checkconf   #检查配置文件
```

区域配置文件

```shell
[root@shkf6-241 ~]# vim /etc/named.rfc1912.zones

[root@shkf6-241 ~]# tail -12 /etc/named.rfc1912.zones

zone "host.com" IN {
        type  master;
        file  "host.com.zone";
        allow-update { 192.168.6.241; };
};

zone "od.com" IN {
        type  master;
        file  "od.com.zone";
        allow-update { 192.168.6.241; };
};
```

配置区域数据文件

- 配置主机域数据文件

```shell
[root@shkf6-241 ~]# vim /var/named/host.com.zone
$ORIGIN host.com.
$TTL 600    ; 10 minutes
@       IN SOA    dns.host.com. dnsadmin.host.com. (
                2019111201 ; serial
                10800      ; refresh (3 hours)
                900        ; retry (15 minutes)
                604800     ; expire (1 week)
                86400      ; minimum (1 day)
                )
            NS   dns.host.com.
$TTL 60    ; 1 minute
dns                A    192.168.6.241
SHKF6-241          A    192.168.6.241
SHKF6-242          A    192.168.6.242
SHKF6-243          A    192.168.6.243
SHKF6-244          A    192.168.6.244
SHKF6-245          A    192.168.6.245
[root@shkf6-241 ~]# vim /var/named/od.com.zone
$ORIGIN od.com.
$TTL 600    ; 10 minutes
@           IN SOA    dns.od.com. dnsadmin.od.com. (
                2019111201 ; serial
                10800      ; refresh (3 hours)
                900        ; retry (15 minutes)
                604800     ; expire (1 week)
                86400      ; minimum (1 day)
                )
                NS   dns.od.com.
$TTL 60    ; 1 minute
dns                A    192.168.6.241
```

#### 3.启动bind9

```shell
[root@shkf6-241 ~]# named-checkconf
[root@shkf6-241 ~]# systemctl start named
[root@shkf6-241 ~]# systemctl enable named
```

#### 4.检查

```shell
[root@shkf6-245 ~]# dig -t A shkf6-244.host.com @192.168.6.241 +short
192.168.6.244
[root@shkf6-245 ~]# dig -t A shkf6-241.host.com @192.168.6.241 +short
192.168.6.241
[root@shkf6-245 ~]# dig -t A shkf6-242.host.com @192.168.6.241 +short
192.168.6.242
[root@shkf6-245 ~]# dig -t A shkf6-243.host.com @192.168.6.241 +short
192.168.6.243
[root@shkf6-245 ~]# dig -t A shkf6-244.host.com @192.168.6.241 +short
192.168.6.244
[root@shkf6-245 ~]# dig -t A shkf6-245.host.com @192.168.6.241 +short
192.168.6.245
```

#### 5.配置dns客户端

- Linux主机上

```shell
~]# cat /etc/resolv.conf 
# Generated by NetworkManager
search host.com
nameserver 192.168.6.241
```

- windows主机上

> 网络和共享中心 -> 网卡设置 -> 设置DNS服务器
> 如有必要，还应设置虚拟网卡的接口地跃点数为：10

#### 6.检查

```shell
[root@shkf6-245 ~]# ping shkf6-241
PING SHKF6-241.host.com (192.168.6.241) 56(84) bytes of data.
64 bytes from node1.98yz.cn (192.168.6.241): icmp_seq=1 ttl=64 time=0.213 ms
^C
--- SHKF6-241.host.com ping statistics ---
1 packets transmitted, 1 received, 0% packet loss, time 0ms
rtt min/avg/max/mdev = 0.213/0.213/0.213/0.000 ms
[root@shkf6-245 ~]# ping shkf6-241.host.com
PING SHKF6-241.host.com (192.168.6.241) 56(84) bytes of data.
64 bytes from node1.98yz.cn (192.168.6.241): icmp_seq=1 ttl=64 time=0.136 ms
^C
--- SHKF6-241.host.com ping statistics ---
1 packets transmitted, 1 received, 0% packet loss, time 0ms
rtt min/avg/max/mdev = 0.136/0.136/0.136/0.000 ms
[C:\Users\mcake]$ ping dns.od.com

正在 Ping dns.od.com [192.168.6.241] 具有 32 字节的数据:
来自 192.168.6.241 的回复: 字节=32 时间<1ms TTL=63
来自 192.168.6.241 的回复: 字节=32 时间<1ms TTL=63
```

### 4.准备自签证书

运维主机shkf6-245.host.com上：

#### 1.安装CFSSL

- 证书签发工具CFSSL:R1.2

[cfssl下载地址](https://pkg.cfssl.org/R1.2/cfssl_linux-amd64)

[cfssl-json下载地址](https://pkg.cfssl.org/R1.2/cfssljson_linux-amd64)

[cfssl-certinfo下载地址](https://pkg.cfssl.org/R1.2/cfssl-certinfo_linux-amd64)

```shell
[root@shkf6-245 ~]# wget https://pkg.cfssl.org/R1.2/cfssl_linux-amd64 -O /usr/bin/cfssl
[root@shkf6-245 ~]# wget https://pkg.cfssl.org/R1.2/cfssljson_linux-amd64 -O /usr/bin/cfssl-json
[root@shkf6-245 ~]# wget https://pkg.cfssl.org/R1.2/cfssl-certinfo_linux-amd64 -O /usr/bin/cfssl-certinfo
[root@shkf6-245 ~]# chmod +x /usr/bin/cfssl*
```

#### 2.创建生成CA证书签名请求(csr)的JSON配置文件

```shell
[root@shkf6-245 ~]# mkdir /opt/certs
[root@shkf6-245 ~]# vim /opt/certs/ca-csr.json
[root@shkf6-245 ~]# cat /opt/certs/ca-csr.json
{
    "CN": "OldboyEdu",
    "hosts": [
    ],
    "key": {
        "algo": "rsa",
        "size": 2048
    },
    "names": [
        {
            "C": "CN",
            "ST": "beijing",
            "L": "beijing",
            "O": "od",
            "OU": "ops"
        }
    ],
    "ca": {
        "expiry": "175200h"
    }
}
```

> CN:Common Name，浏览器使用该字段验证网站是否合法，一般写的是域名。非常重要。
> C：Country。国家
> ST：State，州，省
> L：Locality，城区，城市
> O：Organization Name，组织名称，公司名称
> OU：Organization Unit Name。组织单位名称，公司部门

#### 3.生成CA证书和私钥

```shell
/opt/certs

[root@shkf6-245 certs]# cfssl gencert -initca ca-csr.json | cfssl-json -bare ca
2019/11/12 16:31:15 [INFO] generating a new CA key and certificate from CSR
2019/11/12 16:31:15 [INFO] generate received request
2019/11/12 16:31:15 [INFO] received CSR
2019/11/12 16:31:15 [INFO] generating key: rsa-2048
2019/11/12 16:31:16 [INFO] encoded CSR
2019/11/12 16:31:16 [INFO] signed certificate with serial number 165156553242987548447967502951541624956409280173
[root@shkf6-245 certs]# ll
total 16
-rw-r--r-- 1 root root  993 Nov 12 16:31 ca.csr
-rw-r--r-- 1 root root  328 Nov 12 16:06 ca-csr.json
-rw------- 1 root root 1679 Nov 12 16:31 ca-key.pem
-rw-r--r-- 1 root root 1346 Nov 12 16:31 ca.pem
```

### 5.部署docker环境

在shkf6-243、shkf6-244、shkf6-245上：

#### 1.安装

```shell
[root@shkf6-243 ~]# curl -fsSL https://get.docker.com | bash -s docker --mirror Aliyun

[root@shkf6-244 ~]# curl -fsSL https://get.docker.com | bash -s docker --mirror Aliyun

[root@shkf6-245 ~]# curl -fsSL https://get.docker.com | bash -s docker --mirror Aliyun
```

#### 2.配置

```shell
[root@shkf6-243 ~]# mkdir /etc/docker
[root@shkf6-243 ~]# vim /etc/docker/daemon.json

{
  "graph": "/data/docker",
  "storage-driver": "overlay2",
  "insecure-registries": ["registry.access.redhat.com","quay.io","harbor.od.com"],
  "registry-mirrors": ["https://q2gr04ke.mirror.aliyuncs.com"],
  "bip": "172.6.243.1/24",
  "exec-opts": ["native.cgroupdriver=systemd"],
  "live-restore": true
}
[root@shkf6-244 ~]# mkdir /etc/docker
[root@shkf6-244 ~]# vim /etc/docker/daemon.json

{
  "graph": "/data/docker",
  "storage-driver": "overlay2",
  "insecure-registries": ["registry.access.redhat.com","quay.io","harbor.od.com"],
  "registry-mirrors": ["https://q2gr04ke.mirror.aliyuncs.com"],
  "bip": "172.6.244.1/24",
  "exec-opts": ["native.cgroupdriver=systemd"],
  "live-restore": true
}
[root@shkf6-245 ~]# mkdir /etc/docker
[root@shkf6-245 ~]# vim /etc/docker/daemon.json

{
  "graph": "/data/docker",
  "storage-driver": "overlay2",
  "insecure-registries": ["registry.access.redhat.com","quay.io","harbor.od.com"],
  "registry-mirrors": ["https://q2gr04ke.mirror.aliyuncs.com"],
  "bip": "172.6.245.1/24",
  "exec-opts": ["native.cgroupdriver=systemd"],
  "live-restore": true
}
```

#### 3.启动

在shkf6-243、shkf6-244、shkf6-245上操作

```shell
~]# systemctl start docker
~]# systemctl enable docker
```

### 6、部署docker镜像私有仓库harbor

shkf6-245上：

#### 1.下载软件二进制包并解压

[harbor官方github地址](https://github.com/goharbor/harbor)

[harbor下载地址](https://storage.googleapis.com/harbor-releases/release-1.8.0/harbor-offline-installer-v1.8.3.tgz)

```shell
[root@shkf6-245 ~]# mkdir -p /opt/src/harbor
[root@shkf6-245 ~]# cd /opt/src/harbor
[root@shkf6-245 harbor]# wget https://storage.googleapis.com/harbor-releases/release-1.8.0/harbor-offline-installer-v1.8.3.tgz
[root@shkf6-245 harbor]# tar xvf harbor-offline-installer-v1.8.3.tgz -C /opt
harbor/harbor.v1.8.3.tar.gz
harbor/prepare
harbor/LICENSE
harbor/install.sh
harbor/harbor.yml

[root@shkf6-245 harbor]# mv /opt/harbor /opt/harbor-v1.8.3
[root@shkf6-245 harbor]# ln -s /opt/harbor-v1.8.3 /opt/harbor
```

#### 2.配置

```shell
/opt/harbor/harbor.yml

hostname: harbor.od.com
http:
  port: 180
harbor_admin_password: Harbor12345
data_volume: /data/harbor
log:
  level: info
  rotate_count: 50
  rotate_size: 200M
  location: /data/harbor/logs

mkdir -p /data/harbor/logs
```

#### 3.安装docker-compose

运维主机shkf6-245.host.com上：

```shell
[root@shkf6-245 harbor]# yum install docker-compose -y
[root@shkf6-245 harbor]# rpm -qa docker-compose
docker-compose-1.18.0-4.el7.noarch
```

#### 4.安装harbor

```shell
[root@shkf6-245 harbor]# ll
total 569632
-rw-r--r-- 1 root root 583269670 Sep 16 11:53 harbor.v1.8.3.tar.gz
-rw-r--r-- 1 root root      4526 Nov 13 11:35 harbor.yml
-rwxr-xr-x 1 root root      5088 Sep 16 11:53 install.sh
-rw-r--r-- 1 root root     11347 Sep 16 11:53 LICENSE
-rwxr-xr-x 1 root root      1654 Sep 16 11:53 prepare
[root@shkf6-245 harbor]# sh install.sh
```

#### 5.检查harbor启动情况

```shell
[root@shkf6-245 harbor]# docker-compose ps
      Name                     Command               State             Ports          
--------------------------------------------------------------------------------------
harbor-core         /harbor/start.sh                 Up                               
harbor-db           /entrypoint.sh postgres          Up      5432/tcp                 
harbor-jobservice   /harbor/start.sh                 Up                               
harbor-log          /bin/sh -c /usr/local/bin/ ...   Up      127.0.0.1:1514->10514/tcp
harbor-portal       nginx -g daemon off;             Up      80/tcp                   
nginx               nginx -g daemon off;             Up      0.0.0.0:180->80/tcp      
redis               docker-entrypoint.sh redis ...   Up      6379/tcp                 
registry            /entrypoint.sh /etc/regist ...   Up      5000/tcp                 
registryctl         /harbor/start.sh                 Up
```

#### 6.配置harbor的dns内网解析

在shkf6-241上：

```shell
[root@shkf6-241 ~]# /var/named/od.com.zone
harbor             A    192.168.6.245

# 注意serial前滚一个序号
```

重启named

```shell
[root@shkf6-241 ~]# systemctl restart named
```

测试

```shell
[root@shkf6-241 ~]# dig -t A harbor.od.com +short
192.168.6.245
```

#### 7.安装nginx并配置

用nginx代理180端口：

```shell
[root@shkf6-245 harbor]# yum install nginx -y
[root@shkf6-245 harbor]# rpm -qa nginx
nginx-1.16.1-1.el7.x86_64

[root@shkf6-245 harbor]# vim /etc/nginx/conf.d/harbor.od.com.conf
server {
    listen       80;
    server_name  harbor.od.com;

    client_max_body_size 1000m;

    location / {
        proxy_pass http://127.0.0.1:180;
    }
}

[root@shkf6-245 harbor]# nginx -t
nginx: the configuration file /etc/nginx/nginx.conf syntax is ok
nginx: configuration file /etc/nginx/nginx.conf test is successful
[root@shkf6-245 harbor]# systemctl start nginx
[root@shkf6-245 harbor]# systemctl enable nginx
```

#### 8.浏览器打开[http://harbor.od.com](http://harbor.od.com/)

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_912067b201416e737952e930e79212e1_r.png)

> 账号为admin 密码是Harbor12345

#### 9.检查

1.登录harbor，创建public仓库
![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_07c23944a8431db55a23c99f478faa44_r.png)

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_c18cf349e136767c800c57a09c79ab59_r.png)

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_eafeec65828a52590657bdcea5eb659c_r.png)

2.从docker.io下载镜像nginx:1.7.9

```shell
[root@shkf6-245 ~]# docker pull nginx:1.7.9
1.7.9: Pulling from library/nginx
Image docker.io/library/nginx:1.7.9 uses outdated schema1 manifest format. Please upgrade to a schema2 image for better future compatibility. More information at https://docs.docker.com/registry/spec/deprecated-schema-v1/
a3ed95caeb02: Pull complete 
6f5424ebd796: Pull complete 
d15444df170a: Pull complete 
e83f073daa67: Pull complete 
a4d93e421023: Pull complete 
084adbca2647: Pull complete 
c9cec474c523: Pull complete 
Digest: sha256:e3456c851a152494c3e4ff5fcc26f240206abac0c9d794affb40e0714846c451
Status: Downloaded newer image for nginx:1.7.9
docker.io/library/nginx:1.7.9
```

3.打tag

```shell
[root@shkf6-245 ~]# docker images|grep 1.7.9
nginx                           1.7.9                      84581e99d807        4 years ago         91.7MB
[root@shkf6-245 ~]# docker tag 84581e99d807 harbor.od.com/public/nginx:v1.7.9
```

4.登录私有仓库，并推送镜像nginx

```shell
[root@shkf6-245 ~]# docker login harbor.od.com
Username: admin
Password: 
WARNING! Your password will be stored unencrypted in /root/.docker/config.json.
Configure a credential helper to remove this warning. See
https://docs.docker.com/engine/reference/commandline/login/#credentials-store

Login Succeeded
[root@shkf6-245 ~]# docker push harbor.od.com/public/nginx:v1.7.9
The push refers to repository [harbor.od.com/public/nginx]
5f70bf18a086: Pushed 
4b26ab29a475: Pushed 
ccb1d68e3fb7: Pushed 
e387107e2065: Pushed 
63bf84221cce: Pushed 
e02dce553481: Pushed 
dea2e4984e29: Pushed 
v1.7.9: digest: sha256:b1f5935eb2e9e2ae89c0b3e2e148c19068d91ca502e857052f14db230443e4c2 size: 3012
```

5.查看仓库

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_f43bd912bcc3eb48f776b460fe54736f_r.png)

# 第五章：部署主控节点服务

## 1.部署etcd集群

### 1.集群规划

| 主机名             | 角色        | ip            |
| :----------------- | :---------- | :------------ |
| shkf6-242.host.com | etcd lead   | 192.168.6.242 |
| shkf6-243.host.com | etcd foolow | 192.168.6.243 |
| shkf6-244.host.com | etcd foolow | 192.168.6.244 |

注意：这里部署文档以`shkf6-242.host.com`主机为例，另外两台主机安装部署方法类似

### 2.创建基于根证书的config配置文件

```shell
[root@shkf6-245 certs]# cd /opt/certs/
[root@shkf6-245 certs]# vim /opt/certs/ca-config.json

{
    "signing": {
        "default": {
            "expiry": "175200h"
        },
        "profiles": {
            "server": {
                "expiry": "175200h",
                "usages": [
                    "signing",
                    "key encipherment",
                    "server auth"
                ]
            },
            "client": {
                "expiry": "175200h",
                "usages": [
                    "signing",
                    "key encipherment",
                    "client auth"
                ]
            },
            "peer": {
                "expiry": "175200h",
                "usages": [
                    "signing",
                    "key encipherment",
                    "server auth",
                    "client auth"
                ]
            }
        }
    }
}
```

> 证书类型
> client certificate：客户端使用，用于服务端认证客户端，例如etcdctl、etcd proxy、fleetctl、docker客户端
> server certificate：服务器端使用，客户端已验证服务端身份，例如docker服务端、kube-apiserver
> peer certificate：双向证书，用于etcd集群成员间通信

### 3.创建生成自签证书签名请求(csr)的JSON配置文件

运维主机`shkf6-245.host.com`上：

```shell
[root@shkf6-245 certs]# vi etcd-peer-csr.json

{
    "CN": "k8s-etcd",
    "hosts": [
        "192.168.6.241",
        "192.168.6.242",
        "192.168.6.243",
        "192.168.6.244"
    ],
    "key": {
        "algo": "rsa",
        "size": 2048
    },
    "names": [
        {
            "C": "CN",
            "ST": "beijing",
            "L": "beijing",
            "O": "od",
            "OU": "ops"
        }
    ]
}
```

### 4.生成etcd证书和私钥

```shell
[root@shkf6-245 certs]# cfssl gencert -ca=ca.pem -ca-key=ca-key.pem -config=ca-config.json -profile=peer etcd-peer-csr.json|cfssl-json -bare etcd-peer
2019/11/13 17:00:03 [INFO] generate received request
2019/11/13 17:00:03 [INFO] received CSR
2019/11/13 17:00:03 [INFO] generating key: rsa-2048
2019/11/13 17:00:04 [INFO] encoded CSR
2019/11/13 17:00:04 [INFO] signed certificate with serial number 69997016866371968425072677347883174107938471757
2019/11/13 17:00:04 [WARNING] This certificate lacks a "hosts" field. This makes it unsuitable for
websites. For more information see the Baseline Requirements for the Issuance and Management
of Publicly-Trusted Certificates, v.1.1.6, from the CA/Browser Forum (https://cabforum.org);
specifically, section 10.2.3 ("Information Requirements").
```

### 5.检查生成的证书、私钥

```shell
[root@shkf6-245 certs]# ll
total 36
-rw-r--r-- 1 root root  836 Nov 13 16:32 ca-config.json
-rw-r--r-- 1 root root  993 Nov 12 16:31 ca.csr
-rw-r--r-- 1 root root  328 Nov 12 16:06 ca-csr.json
-rw------- 1 root root 1679 Nov 12 16:31 ca-key.pem
-rw-r--r-- 1 root root 1346 Nov 12 16:31 ca.pem
-rw-r--r-- 1 root root 1062 Nov 13 17:00 etcd-peer.csr
-rw-r--r-- 1 root root  379 Nov 13 16:34 etcd-peer-csr.json
-rw------- 1 root root 1679 Nov 13 17:00 etcd-peer-key.pem
-rw-r--r-- 1 root root 1428 Nov 13 17:00 etcd-peer.pem
```

### 6.创建etcd用户

在shkf6-242机器上：

```shell
[root@shkf6-242 ~]# useradd -s /sbin/nologin  -M etcd
```

### 7.下载软件，解压，做软连接

[etcd下载地址](https://github.com/etcd-io/etcd/tags)

[这里使用的是etcd-v3.1.20](https://github.com/etcd-io/etcd/releases/download/v3.1.20/etcd-v3.1.20-linux-amd64.tar.gz)

在shkf6-242机器上：

```shell
[root@shkf6-242 ~]# cd /opt/src/
[root@shkf6-242 src]# wget https://github.com/etcd-io/etcd/releases/download/v3.1.20/etcd-v3.1.20-linux-amd64.tar.gz
[root@shkf6-242 src]# tar xf etcd-v3.1.20-linux-amd64.tar.gz -C /opt
[root@shkf6-242 src]# cd /opt/
[root@shkf6-242 opt]# mv etcd-v3.1.20-linux-amd64/ etcd-v3.1.20
[root@shkf6-242 opt]# ln -s /opt/etcd-v3.1.20/ /opt/etcd
[root@shkf6-242 src]# ls -l /opt/
total 0
lrwxrwxrwx 1 root   root   18 Nov 13 17:30 etcd -> /opt/etcd-v3.1.20/
drwxr-xr-x 3 478493 89939 123 Oct 11  2018 etcd-v3.1.20
drwxr-xr-x 2 root   root   45 Nov 13 17:27 src
```

### 8.创建目录，拷贝证书，私钥

在shkf6-242机器上：

- 创建目录

```shell
[root@shkf6-242 src]# mkdir -p /opt/etcd/certs /data/etcd /data/logs/etcd-server
```

- 拷贝证书

```shell
[root@shkf6-242 src]# cd /opt/etcd/certs
[root@shkf6-242 certs]# scp -P52113 shkf6-245:/opt/certs/ca.pem /opt/etcd/certs/
root@shkf6-245's password: 
ca.pem                                                                                             100% 1346   133.4KB/s   00:00    
[root@shkf6-242 certs]# scp -P52113 shkf6-245:/opt/certs/etcd-peer.pem /opt/etcd/certs/
root@shkf6-245's password: 
etcd-peer.pem                                                                                      100% 1428   208.6KB/s   00:00    
[root@shkf6-242 certs]# scp -P52113 shkf6-245:/opt/certs/etcd-peer-key.pem /opt/etcd/certs/
root@shkf6-245's password: 
etcd-peer-key.pem
```

> 将运维主机上生成的ca.pem,etcd-peer-key.pem,etcd-peer.pem拷贝到/ope/etcd/certs目录中，注意私钥权限600

- 修改权限

```shell
/opt/etcd/certs

[root@shkf6-242 certs]# chown -R etcd.etcd /opt/etcd/certs /data/etcd /data/logs/etcd-server
[root@shkf6-242 certs]# ls -l
total 12
-rw-r--r-- 1 etcd etcd 1346 Nov 13 17:45 ca.pem
-rw------- 1 etcd etcd 1679 Nov 13 17:46 etcd-peer-key.pem
-rw-r--r-- 1 etcd etcd 1428 Nov 13 17:45 etcd-peer.pem
```

### 9.创建etcd服务启动脚本

在shkf6-242机器上：

```shell
[root@shkf6-242 certs]# vim /opt/etcd/etcd-server-startup.sh

#!/bin/sh
./etcd --name etcd-server-6-242 \
       --data-dir /data/etcd/etcd-server \
       --listen-peer-urls https://192.168.6.242:2380 \
       --listen-client-urls https://192.168.6.242:2379,http://127.0.0.1:2379 \
       --quota-backend-bytes 8000000000 \
       --initial-advertise-peer-urls https://192.168.6.242:2380 \
       --advertise-client-urls https://192.168.6.242:2379,http://127.0.0.1:2379 \
       --initial-cluster  etcd-server-6-242=https://192.168.6.242:2380,etcd-server-6-243=https://192.168.6.243:2380,etcd-server-6-244=https://192.168.6.244:2380 \
       --ca-file ./certs/ca.pem \
       --cert-file ./certs/etcd-peer.pem \
       --key-file ./certs/etcd-peer-key.pem \
       --client-cert-auth  \
       --trusted-ca-file ./certs/ca.pem \
       --peer-ca-file ./certs/ca.pem \
       --peer-cert-file ./certs/etcd-peer.pem \
       --peer-key-file ./certs/etcd-peer-key.pem \
       --peer-client-cert-auth \
       --peer-trusted-ca-file ./certs/ca.pem \
       --log-output stdout
```

注意：etcd集群各主机的启动脚本略有不同，部署其他节点是需要注意。

### 10.调整权限

在shkf6-242机器上：

```shell
[root@shkf6-242 certs]# cd ../
[root@shkf6-242 etcd]# chmod +x etcd-server-startup.sh 
[root@shkf6-242 etcd]# ll etcd-server-startup.sh 
-rwxr-xr-x 1 root root 1013 Nov 14 08:52 etcd-server-startup.sh
```

### 11.安装supervisor软件

在shkf6-242机器上：

```shell
[root@shkf6-242 etcd]# yum install supervisor -y
[root@shkf6-242 etcd]# systemctl start supervisord.service
[root@shkf6-242 etcd]# systemctl enable supervisord.service
```

### 12.创建etcd-server的启动配置

在shkf6-242机器上：

```shell
[root@shkf6-242 etcd]# vim /etc/supervisord.d/etcd-server.ini
[root@shkf6-242 etcd]# cat /etc/supervisord.d/etcd-server.ini
[program:etcd-server-6-242]
command=/opt/etcd/etcd-server-startup.sh                        ; the program (relative uses PATH, can take args)
numprocs=1                                                      ; number of processes copies to start (def 1)
directory=/opt/etcd                                             ; directory to cwd to before exec (def no cwd)
autostart=true                                                  ; start at supervisord start (default: true)
autorestart=true                                                ; retstart at unexpected quit (default: true)
startsecs=30                                                    ; number of secs prog must stay running (def. 1)
startretries=3                                                  ; max # of serial start failures (default 3)
exitcodes=0,2                                                   ; 'expected' exit codes for process (default 0,2)
stopsignal=QUIT                                                 ; signal used to kill process (default TERM)
stopwaitsecs=10                                                 ; max num secs to wait b4 SIGKILL (default 10)
user=etcd                                                       ; setuid to this UNIX account to run the program
redirect_stderr=true                                            ; redirect proc stderr to stdout (default false)
killasgroup=true                                                ; kill all process in a group
stopasgroup=true                                                ; stop all process in a group
stdout_logfile=/data/logs/etcd-server/etcd.stdout.log           ; stdout log path, NONE for none; default AUTO
stdout_logfile_maxbytes=64MB                                    ; max # logfile bytes b4 rotation (default 50MB)
stdout_logfile_backups=4                                        ; # of stdout logfile backups (default 10)
stdout_capture_maxbytes=1MB                                     ; number of bytes in 'capturemode' (default 0)
stdout_events_enabled=false                                     ; emit events on stdout writes (default false)
```

注意：etcd集群各主机启动配置略有不同，配置其他节点时注意修改。

### 13.启动etcd服务并检查

在shkf6-242机器上：

```shell
[root@shkf6-242 etcd]# supervisorctl update
etcd-server-6-242: added process group

[root@shkf6-242 etcd]# supervisorctl status
etcd-server-6-242                RUNNING   pid 10375, uptime 0:00:43
[root@shkf6-242 etcd]# netstat -lntup|grep etcd
tcp        0      0 192.168.6.242:2379      0.0.0.0:*               LISTEN      10376/./etcd
tcp        0      0 127.0.0.1:2379          0.0.0.0:*               LISTEN      10376/./etcd
tcp        0      0 192.168.6.242:2380      0.0.0.0:*               LISTEN      10376/./etcd 
```

### 14.安装部署启动检查所有集群规划的etcd服务

在shkf6-243机器上：

```shell
[root@shkf6-243 ~]# useradd -s /sbin/nologin -M etcd
[root@shkf6-243 ~]# mkdir /opt/src
[root@shkf6-243 ~]# cd /opt/src
[root@shkf6-243 src]# wget https://github.com/etcd-io/etcd/releases/download/v3.1.20/etcd-v3.1.20-linux-amd64.tar.gz
[root@shkf6-243 src]# tar xf etcd-v3.1.20-linux-amd64.tar.gz -C /opt
[root@shkf6-243 src]# cd /opt/
[root@shkf6-243 opt]# mv etcd-v3.1.20-linux-amd64/ etcd-v3.1.20
[root@shkf6-243 opt]# ln -s /opt/etcd-v3.1.20/ /opt/etcd
total 0
drwx--x--x  4 root   root   28 Nov 13 10:33 containerd
lrwxrwxrwx  1 root   root   18 Nov 14 09:28 etcd -> /opt/etcd-v3.1.20/
drwxr-xr-x  3 478493 89939 123 Oct 11  2018 etcd-v3.1.20
drwxr-xr-x. 2 root   root    6 Sep  7  2017 rh
drwxr-xr-x  2 root   root   45 Nov 14 09:26 src
[root@shkf6-243 opt]# rm -fr rh containerd/ containerd/
[root@shkf6-243 opt]# mkdir -p /opt/etcd/certs /data/etcd /data/logs/etcd-server
[root@shkf6-243 opt]# cd /opt/etcd/certs
[root@shkf6-243 certs]# scp -P52113 shkf6-245:/opt/certs/ca.pem /opt/etcd/certs/
[root@shkf6-243 certs]# scp -P52113 shkf6-245:/opt/certs/etcd-peer.pem /opt/etcd/certs/
[root@shkf6-243 certs]# scp -P52113 shkf6-245:/opt/certs/etcd-peer-key.pem /opt/etcd/certs/
[root@shkf6-243 certs]#  chown -R etcd.etcd /opt/etcd/certs /data/etcd /data/logs/etcd-server
[root@shkf6-243 certs]# vim /opt/etcd/etcd-server-startup.sh
[root@shkf6-243 certs]# cat /opt/etcd/etcd-server-startup.sh
#!/bin/sh
./etcd --name etcd-server-6-243 \
       --data-dir /data/etcd/etcd-server \
       --listen-peer-urls https://192.168.6.243:2380 \
       --listen-client-urls https://192.168.6.243:2379,http://127.0.0.1:2379 \
       --quota-backend-bytes 8000000000 \
       --initial-advertise-peer-urls https://192.168.6.243:2380 \
       --advertise-client-urls https://192.168.6.243:2379,http://127.0.0.1:2379 \
       --initial-cluster  etcd-server-6-242=https://192.168.6.242:2380,etcd-server-6-243=https://192.168.6.243:2380,etcd-server-6-244=https://192.168.6.244:2380 \
       --ca-file ./certs/ca.pem \
       --cert-file ./certs/etcd-peer.pem \
       --key-file ./certs/etcd-peer-key.pem \
       --client-cert-auth  \
       --trusted-ca-file ./certs/ca.pem \
       --peer-ca-file ./certs/ca.pem \
       --peer-cert-file ./certs/etcd-peer.pem \
       --peer-key-file ./certs/etcd-peer-key.pem \
       --peer-client-cert-auth \
       --peer-trusted-ca-file ./certs/ca.pem \
       --log-output stdout
[root@shkf6-243 certs]# chmod +x /opt/etcd/etcd-server-startup.sh
[root@shkf6-243 certs]# yum install supervisor -y
[root@shkf6-243 certs]# systemctl start supervisord.service
[root@shkf6-243 certs]# systemctl enable supervisord.service
[root@shkf6-243 certs]# vim /etc/supervisord.d/etcd-server.ini
[root@shkf6-243 certs]# cat /etc/supervisord.d/etcd-server.ini
[program:etcd-server-6-243]
command=/opt/etcd/etcd-server-startup.sh                        ; the program (relative uses PATH, can take args)
numprocs=1                                                      ; number of processes copies to start (def 1)
directory=/opt/etcd                                             ; directory to cwd to before exec (def no cwd)
autostart=true                                                  ; start at supervisord start (default: true)
autorestart=true                                                ; retstart at unexpected quit (default: true)
startsecs=30                                                    ; number of secs prog must stay running (def. 1)
startretries=3                                                  ; max # of serial start failures (default 3)
exitcodes=0,2                                                   ; 'expected' exit codes for process (default 0,2)
stopsignal=QUIT                                                 ; signal used to kill process (default TERM)
stopwaitsecs=10                                                 ; max num secs to wait b4 SIGKILL (default 10)
user=etcd                                                       ; setuid to this UNIX account to run the program
redirect_stderr=true                                            ; redirect proc stderr to stdout (default false)
killasgroup=true                                                ; kill all process in a group
stopasgroup=true                                                ; stop all process in a group
stdout_logfile=/data/logs/etcd-server/etcd.stdout.log           ; stdout log path, NONE for none; default AUTO
stdout_logfile_maxbytes=64MB                                    ; max # logfile bytes b4 rotation (default 50MB)
stdout_logfile_backups=4                                        ; # of stdout logfile backups (default 10)
stdout_capture_maxbytes=1MB                                     ; number of bytes in 'capturemode' (default 0)
stdout_events_enabled=false                                     ; emit events on stdout writes (default false)
[root@shkf6-243 certs]# supervisorctl update
etcd-server-6-243: added process group
[root@shkf6-243 certs]# supervisorctl status
etcd-server-6-243                STARTING  
[root@shkf6-243 certs]# netstat -lntup|grep etcd
tcp        0      0 192.168.6.243:2379      0.0.0.0:*               LISTEN      12113/./etcd        
tcp        0      0 127.0.0.1:2379          0.0.0.0:*               LISTEN      12113/./etcd        
tcp        0      0 192.168.6.243:2380      0.0.0.0:*               LISTEN      12113/./etcd
```

在shkf6-244机器上：

```shell
[root@shkf6-244 ~]# useradd -s /sbin/nologin  -M etcd
[root@shkf6-244 ~]# mkdir /opt/src
[root@shkf6-244 ~]# cd /opt/src
[root@shkf6-244 src]# wget https://github.com/etcd-io/etcd/releases/download/v3.1.20/etcd-v3.1.20-linux-amd64.tar.gz
[root@shkf6-244 src]# useradd -s /sbin/nologin  -M etcd
useradd: user 'etcd' already exists
[root@shkf6-244 src]# tar xf etcd-v3.1.20-linux-amd64.tar.gz -C /opt
[root@shkf6-244 src]# cd /opt/
[root@shkf6-244 opt]# mv etcd-v3.1.20-linux-amd64/ etcd-v3.1.20
[root@shkf6-244 opt]# ln -s /opt/etcd-v3.1.20/ /opt/etcd
[root@shkf6-244 opt]# mkdir -p /opt/etcd/certs /data/etcd /data/logs/etcd-server
[root@shkf6-244 opt]#  cd /opt/etcd/certs
[root@shkf6-244 certs]# scp -P52113 shkf6-245:/opt/certs/ca.pem /opt/etcd/certs/
[root@shkf6-244 certs]# scp -P52113 shkf6-245:/opt/certs/etcd-peer.pem /opt/etcd/certs/
[root@shkf6-244 certs]# scp -P52113 shkf6-245:/opt/certs/etcd-peer-key.pem /opt/etcd/certs/
[root@shkf6-244 certs]# chown -R etcd.etcd /opt/etcd/certs /data/etcd /data/logs/etcd-server
[root@shkf6-244 certs]# vim /opt/etcd/etcd-server-startup.sh
[root@shkf6-244 etcd]# cat /opt/etcd/etcd-server-startup.sh
#!/bin/sh
./etcd --name etcd-server-6-244 \
       --data-dir /data/etcd/etcd-server \
       --listen-peer-urls https://192.168.6.244:2380 \
       --listen-client-urls https://192.168.6.244:2379,http://127.0.0.1:2379 \
       --quota-backend-bytes 8000000000 \
       --initial-advertise-peer-urls https://192.168.6.244:2380 \
       --advertise-client-urls https://192.168.6.244:2379,http://127.0.0.1:2379 \
       --initial-cluster  etcd-server-6-242=https://192.168.6.242:2380,etcd-server-6-243=https://192.168.6.243:2380,etcd-server-6-244=https://192.168.6.244:2380 \
       --ca-file ./certs/ca.pem \
       --cert-file ./certs/etcd-peer.pem \
       --key-file ./certs/etcd-peer-key.pem \
       --client-cert-auth  \
       --trusted-ca-file ./certs/ca.pem \
       --peer-ca-file ./certs/ca.pem \
       --peer-cert-file ./certs/etcd-peer.pem \
       --peer-key-file ./certs/etcd-peer-key.pem \
       --peer-client-cert-auth \
       --peer-trusted-ca-file ./certs/ca.pem \
       --log-output stdout
[root@shkf6-244 certs]# cd ../
[root@shkf6-244 etcd]# chmod +x etcd-server-startup.sh 
[root@shkf6-244 etcd]# yum install supervisor -y
[root@shkf6-244 etcd]# systemctl start supervisord.service
[root@shkf6-244 etcd]# systemctl enable supervisord.service
[root@shkf6-244 etcd]#  vim /etc/supervisord.d/etcd-server.ini
[root@shkf6-244 etcd]# cat /etc/supervisord.d/etcd-server.ini
[program:etcd-server-6-244]
command=/opt/etcd/etcd-server-startup.sh                        ; the program (relative uses PATH, can take args)
numprocs=1                                                      ; number of processes copies to start (def 1)
directory=/opt/etcd                                             ; directory to cwd to before exec (def no cwd)
autostart=true                                                  ; start at supervisord start (default: true)
autorestart=true                                                ; retstart at unexpected quit (default: true)
startsecs=30                                                    ; number of secs prog must stay running (def. 1)
startretries=3                                                  ; max # of serial start failures (default 3)
exitcodes=0,2                                                   ; 'expected' exit codes for process (default 0,2)
stopsignal=QUIT                                                 ; signal used to kill process (default TERM)
stopwaitsecs=10                                                 ; max num secs to wait b4 SIGKILL (default 10)
user=etcd                                                       ; setuid to this UNIX account to run the program
redirect_stderr=true                                            ; redirect proc stderr to stdout (default false)
killasgroup=true                                                ; kill all process in a group
stopasgroup=true                                                ; stop all process in a group
stdout_logfile=/data/logs/etcd-server/etcd.stdout.log           ; stdout log path, NONE for none; default AUTO
stdout_logfile_maxbytes=64MB                                    ; max # logfile bytes b4 rotation (default 50MB)
stdout_logfile_backups=4                                        ; # of stdout logfile backups (default 10)
stdout_capture_maxbytes=1MB                                     ; number of bytes in 'capturemode' (default 0)
stdout_events_enabled=false                                     ; emit events on stdout writes (default false)
[root@shkf6-244 etcd]# supervisorctl update
[root@shkf6-244 etcd]# supervisorctl status
etcd-server-6-244                RUNNING   pid 11748, uptime 0:00:33
[root@shkf6-244 etcd]#  netstat -lntup|grep etcd
tcp        0      0 192.168.6.244:2379      0.0.0.0:*               LISTEN      11749/./etcd        
tcp        0      0 127.0.0.1:2379          0.0.0.0:*               LISTEN      11749/./etcd        
tcp        0      0 192.168.6.244:2380      0.0.0.0:*               LISTEN      11749/./etcd
[root@shkf6-244 etcd]# ./etcdctl cluster-health
member 4244d625c76d5482 is healthy: got healthy result from http://127.0.0.1:2379
member aa911af67b8285a2 is healthy: got healthy result from http://127.0.0.1:2379
member c751958d48e7e127 is healthy: got healthy result from http://127.0.0.1:2379
cluster is healthy
```

### 15.检查集群状态

**三个etcd节点都起来后**

在shkf6-242机器上：

```shell
[root@shkf6-242 etcd]# ./etcdctl cluster-health
member 4244d625c76d5482 is healthy: got healthy result from http://127.0.0.1:2379
member aa911af67b8285a2 is healthy: got healthy result from http://127.0.0.1:2379
member c751958d48e7e127 is healthy: got healthy result from http://127.0.0.1:2379
cluster is healthy
[root@shkf6-242 etcd]# ./etcdctl member list
4244d625c76d5482: name=etcd-server-6-242 peerURLs=https://192.168.6.242:2380 clientURLs=http://127.0.0.1:2379,https://192.168.6.242:2379 isLeader=true
aa911af67b8285a2: name=etcd-server-6-243 peerURLs=https://192.168.6.243:2380 clientURLs=http://127.0.0.1:2379,https://192.168.6.243:2379 isLeader=false
c751958d48e7e127: name=etcd-server-6-244 peerURLs=https://192.168.6.244:2380 clientURLs=http://127.0.0.1:2379,https://192.168.6.244:2379 isLeader=false
```

## 2.部署kube-apiserver集群

### 1.集群规划

| 主机名             | 角色           | ip            |
| :----------------- | :------------- | :------------ |
| shkf6-243.host.com | kube-apiserver | 192.168.6.243 |
| shkf6-244.host.com | kube-apiserver | 192.168.6.244 |
| shkf6-241.host.com | 4层负载均衡    | 192.168.6.241 |
| shkf6-242.host.com | 4层负载均衡    | 192.168.6.242 |

注意：这里`192.168.6.241`和`192.168.6.242`使用nginx做4层负载均衡器，用keepalived跑一个vip：`192.168.6.66`，代理两个kube-apiserver，实现高可用

这里部署文档以`shkf6-243.host.com`主机为例，另外一台运算节点安装部署方法类似

### 2.下载软件，解压，做软链

`shkf6-243.host.com`主机上：

[kubernetes官方Github地址](https://github.com/kubernetes/kubernetes)
[kubernetes下地址](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG-1.15.md#downloads-for-v1154)

```shell
[root@shkf6-243 src]# cd /opt/src/
[root@shkf6-243 src]# wget http://down.sunrisenan.com/k8s/kubernetes/kubernetes-server-linux-amd64-v1.15.2.tar.gz

[root@shkf6-243 src]# tar xf kubernetes-server-linux-amd64-v1.15.2.tar.gz -C /opt/
[root@shkf6-243 src]# cd /opt/

[root@shkf6-243 opt]# mv kubernetes/ kubernetes-v1.15.2
[root@shkf6-243 opt]# ln -s /opt/kubernetes-v1.15.2/ /opt/kubernetes
```

**删除源码**

```shell
[root@shkf6-243 opt]# cd kubernetes
[root@shkf6-243 kubernetes]# rm -f kubernetes-src.tar.gz 
[root@shkf6-243 kubernetes]# ll
total 1180
drwxr-xr-x 2 root root       6 Aug  5 18:01 addons
-rw-r--r-- 1 root root 1205293 Aug  5 18:01 LICENSES
drwxr-xr-x 3 root root      17 Aug  5 17:57 server
```

**删除docker镜像**

```shell
[root@shkf6-243 kubernetes]# cd server/bin
[root@shkf6-243 bin]# ll
total 1548800
-rwxr-xr-x 1 root root  43534816 Aug  5 18:01 apiextensions-apiserver
-rwxr-xr-x 1 root root 100548640 Aug  5 18:01 cloud-controller-manager
-rw-r--r-- 1 root root         8 Aug  5 17:57 cloud-controller-manager.docker_tag
-rw-r--r-- 1 root root 144437760 Aug  5 17:57 cloud-controller-manager.tar
-rwxr-xr-x 1 root root 200648416 Aug  5 18:01 hyperkube
-rwxr-xr-x 1 root root  40182208 Aug  5 18:01 kubeadm
-rwxr-xr-x 1 root root 164501920 Aug  5 18:01 kube-apiserver
-rw-r--r-- 1 root root         8 Aug  5 17:57 kube-apiserver.docker_tag
-rw-r--r-- 1 root root 208390656 Aug  5 17:57 kube-apiserver.tar
-rwxr-xr-x 1 root root 116397088 Aug  5 18:01 kube-controller-manager
-rw-r--r-- 1 root root         8 Aug  5 17:57 kube-controller-manager.docker_tag
-rw-r--r-- 1 root root 160286208 Aug  5 17:57 kube-controller-manager.tar
-rwxr-xr-x 1 root root  42985504 Aug  5 18:01 kubectl
-rwxr-xr-x 1 root root 119616640 Aug  5 18:01 kubelet
-rwxr-xr-x 1 root root  36987488 Aug  5 18:01 kube-proxy
-rw-r--r-- 1 root root         8 Aug  5 17:57 kube-proxy.docker_tag
-rw-r--r-- 1 root root  84282368 Aug  5 17:57 kube-proxy.tar
-rwxr-xr-x 1 root root  38786144 Aug  5 18:01 kube-scheduler
-rw-r--r-- 1 root root         8 Aug  5 17:57 kube-scheduler.docker_tag
-rw-r--r-- 1 root root  82675200 Aug  5 17:57 kube-scheduler.tar
-rwxr-xr-x 1 root root   1648224 Aug  5 18:01 mounter
[root@shkf6-243 bin]# rm -f *.tar
[root@shkf6-243 bin]# rm -f *_tag
[root@shkf6-243 bin]# ll
total 884636
-rwxr-xr-x 1 root root  43534816 Aug  5 18:01 apiextensions-apiserver
-rwxr-xr-x 1 root root 100548640 Aug  5 18:01 cloud-controller-manager
-rwxr-xr-x 1 root root 200648416 Aug  5 18:01 hyperkube
-rwxr-xr-x 1 root root  40182208 Aug  5 18:01 kubeadm
-rwxr-xr-x 1 root root 164501920 Aug  5 18:01 kube-apiserver
-rwxr-xr-x 1 root root 116397088 Aug  5 18:01 kube-controller-manager
-rwxr-xr-x 1 root root  42985504 Aug  5 18:01 kubectl
-rwxr-xr-x 1 root root 119616640 Aug  5 18:01 kubelet
-rwxr-xr-x 1 root root  36987488 Aug  5 18:01 kube-proxy
-rwxr-xr-x 1 root root  38786144 Aug  5 18:01 kube-scheduler
-rwxr-xr-x 1 root root   1648224 Aug  5 18:01 mounter
```

### 3.签发client证书

在运维机`shkf6.245.host.com`上：

#### 1.创建生成证书请求（csr）的JSON配置文件

```shell
vim /opt/certs/client-csr.json

{
    "CN": "k8s-node",
    "hosts": [
    ],
    "key": {
        "algo": "rsa",
        "size": 2048
    },
    "names": [
        {
            "C": "CN",
            "ST": "beijing",
            "L": "beijing",
            "O": "od",
            "OU": "ops"
        }
    ]
}
```

#### 2.生成client证书和私钥

```shell
[root@shkf6-245 certs]# cfssl gencert -ca=ca.pem -ca-key=ca-key.pem -config=ca-config.json -profile=client client-csr.json |cfssl-json -bare client
2019/11/14 13:59:24 [INFO] generate received request
2019/11/14 13:59:24 [INFO] received CSR
2019/11/14 13:59:24 [INFO] generating key: rsa-2048
2019/11/14 13:59:24 [INFO] encoded CSR
2019/11/14 13:59:24 [INFO] signed certificate with serial number 71787071397684874048844497862502145400133190813
2019/11/14 13:59:24 [WARNING] This certificate lacks a "hosts" field. This makes it unsuitable for
websites. For more information see the Baseline Requirements for the Issuance and Management
of Publicly-Trusted Certificates, v.1.1.6, from the CA/Browser Forum (https://cabforum.org);
specifically, section 10.2.3 ("Information Requirements").
```

#### 3.检查生成的证书和私钥

```shell
[root@shkf6-245 certs]# ll client*
-rw-r--r-- 1 root root  993 Nov 14 13:59 client.csr
-rw-r--r-- 1 root root  280 Nov 14 13:59 client-csr.json
-rw------- 1 root root 1679 Nov 14 13:59 client-key.pem
-rw-r--r-- 1 root root 1363 Nov 14 13:59 client.pem
```

### 4.签发kube-apiserver证书

在运维机`shkf6.245.host.com`上：

#### 1.创建生成证书签名请求（csr）的josn配置文件

```shell
[root@shkf6-245 certs]# vim /opt/certs/apiserver-csr.json
[root@shkf6-245 certs]# cat /opt/certs/apiserver-csr.json
{
    "CN": "k8s-apiserver",
    "hosts": [
        "127.0.0.1",
        "10.96.0.1",
        "kubernetes.default",
        "kubernetes.default.svc",
        "kubernetes.default.svc.cluster",
        "kubernetes.default.svc.cluster.local",
        "192.168.6.66",
        "192.168.6.243",
        "192.168.6.244",
        "192.168.6.245"
    ],
    "key": {
        "algo": "rsa",
        "size": 2048
    },
    "names": [
        {
            "C": "CN",
            "ST": "beijing",
            "L": "beijing",
            "O": "od",
            "OU": "ops"
        }
    ]
} 
```

#### 2.生成kube-apiserver证书和私钥

```shell
[root@shkf6-245 certs]# cfssl gencert -ca=ca.pem -ca-key=ca-key.pem -config=ca-config.json -profile=server apiserver-csr.json |cfssl-json -bare apiserver
2019/11/14 14:10:01 [INFO] generate received request
2019/11/14 14:10:01 [INFO] received CSR
2019/11/14 14:10:01 [INFO] generating key: rsa-2048
2019/11/14 14:10:02 [INFO] encoded CSR
2019/11/14 14:10:02 [INFO] signed certificate with serial number 531358145467350237994138515547646071524442824033
2019/11/14 14:10:02 [WARNING] This certificate lacks a "hosts" field. This makes it unsuitable for
websites. For more information see the Baseline Requirements for the Issuance and Management
of Publicly-Trusted Certificates, v.1.1.6, from the CA/Browser Forum (https://cabforum.org);
specifically, section 10.2.3 ("Information Requirements").
```

#### 3.检查生成的证书和私钥

```shell
[root@shkf6-245 certs]# ll apiserver*
-rw-r--r-- 1 root root 1249 Nov 14 14:10 apiserver.csr
-rw-r--r-- 1 root root  581 Nov 14 14:09 apiserver-csr.json
-rw------- 1 root root 1679 Nov 14 14:10 apiserver-key.pem
-rw-r--r-- 1 root root 1598 Nov 14 14:10 apiserver.pem
```

### 5.拷贝证书至各个运算节点，并创建配置

在运维机`shkf6.243.host.com`上：

拷贝证书

```shell
[root@shkf6-243 bin]# pwd
/opt/kubernetes/server/bin
[root@shkf6-243 bin]# mkdir cert
[root@shkf6-243 bin]# scp -P52113 shkf6-245:/opt/certs/apiserver-key.pem /opt/kubernetes/server/bin/cert/
[root@shkf6-243 bin]# scp -P52113 shkf6-245:/opt/certs/apiserver.pem /opt/kubernetes/server/bin/cert/
[root@shkf6-243 bin]# scp -P52113 shkf6-245:/opt/certs/ca-key.pem /opt/kubernetes/server/bin/cert/
[root@shkf6-243 bin]# scp -P52113 shkf6-245:/opt/certs/ca.pem /opt/kubernetes/server/bin/cert/  
[root@shkf6-243 bin]# scp -P52113 shkf6-245:/opt/certs/client-key.pem /opt/kubernetes/server/bin/cert/
[root@shkf6-243 bin]# scp -P52113 shkf6-245:/opt/certs/client.pem /opt/kubernetes/server/bin/cert/
```

创建配置文件

```shell
[root@shkf6-243 bin]# mkdir conf
[root@shkf6-243 bin]# vi conf/audit.yaml
[root@shkf6-243 bin]# cat conf/audit.yaml 
apiVersion: audit.k8s.io/v1beta1 # This is required.
kind: Policy
# Don't generate audit events for all requests in RequestReceived stage.
omitStages:
  - "RequestReceived"
rules:
  # Log pod changes at RequestResponse level
  - level: RequestResponse
    resources:
    - group: ""
      # Resource "pods" doesn't match requests to any subresource of pods,
      # which is consistent with the RBAC policy.
      resources: ["pods"]
  # Log "pods/log", "pods/status" at Metadata level
  - level: Metadata
    resources:
    - group: ""
      resources: ["pods/log", "pods/status"]

  # Don't log requests to a configmap called "controller-leader"
  - level: None
    resources:
    - group: ""
      resources: ["configmaps"]
      resourceNames: ["controller-leader"]

  # Don't log watch requests by the "system:kube-proxy" on endpoints or services
  - level: None
    users: ["system:kube-proxy"]
    verbs: ["watch"]
    resources:
    - group: "" # core API group
      resources: ["endpoints", "services"]

  # Don't log authenticated requests to certain non-resource URL paths.
  - level: None
    userGroups: ["system:authenticated"]
    nonResourceURLs:
    - "/api*" # Wildcard matching.
    - "/version"

  # Log the request body of configmap changes in kube-system.
  - level: Request
    resources:
    - group: "" # core API group
      resources: ["configmaps"]
    # This rule only applies to resources in the "kube-system" namespace.
    # The empty string "" can be used to select non-namespaced resources.
    namespaces: ["kube-system"]

  # Log configmap and secret changes in all other namespaces at the Metadata level.
  - level: Metadata
    resources:
    - group: "" # core API group
      resources: ["secrets", "configmaps"]

  # Log all other resources in core and extensions at the Request level.
  - level: Request
    resources:
    - group: "" # core API group
    - group: "extensions" # Version of group should NOT be included.

  # A catch-all rule to log all other requests at the Metadata level.
  - level: Metadata
    # Long-running requests like watches that fall under this rule will not
    # generate an audit event in RequestReceived.
    omitStages:
      - "RequestReceived"
```

### 6.创建启动脚本

在运维机`shkf6.243.host.com`上：

```shell
[root@shkf6-243 bin]# vim /opt/kubernetes/server/bin/kube-apiserver.sh
[root@shkf6-243 bin]# cat /opt/kubernetes/server/bin/kube-apiserver.sh
#!/bin/bash
./kube-apiserver \
  --apiserver-count 2 \
  --audit-log-path /data/logs/kubernetes/kube-apiserver/audit-log \
  --audit-policy-file ./conf/audit.yaml \
  --authorization-mode RBAC \
  --client-ca-file ./cert/ca.pem \
  --requestheader-client-ca-file ./cert/ca.pem \
  --enable-admission-plugins NamespaceLifecycle,LimitRanger,ServiceAccount,DefaultStorageClass,DefaultTolerationSeconds,MutatingAdmissionWebhook,ValidatingAdmissionWebhook,ResourceQuota \
  --etcd-cafile ./cert/ca.pem \
  --etcd-certfile ./cert/client.pem \
  --etcd-keyfile ./cert/client-key.pem \
  --etcd-servers https://192.168.6.242:2379,https://192.168.6.243:2379,https://192.168.6.244:2379 \
  --service-account-key-file ./cert/ca-key.pem \
  --service-cluster-ip-range 10.96.0.0/22 \
  --service-node-port-range 3000-29999 \
  --target-ram-mb=1024 \
  --kubelet-client-certificate ./cert/client.pem \
  --kubelet-client-key ./cert/client-key.pem \
  --log-dir  /data/logs/kubernetes/kube-apiserver \
  --tls-cert-file ./cert/apiserver.pem \
  --tls-private-key-file ./cert/apiserver-key.pem \
  --v 2
```

### 7.调整权限和目录

在运维机`shkf6.243.host.com`上：

```shell
[root@shkf6-243 bin]# chmod +x /opt/kubernetes/server/bin/kube-apiserver.sh
[root@shkf6-243 bin]# mkdir -p /data/logs/kubernetes/kube-apiserver
```

### 8.创建supervisor配置

在运维机`shkf6.243.host.com`上：

```shell
[root@shkf6-243 bin]# vi /etc/supervisord.d/kube-apiserver.ini
[root@shkf6-243 bin]# cat /etc/supervisord.d/kube-apiserver.ini
[program:kube-apiserver-6-243]
command=/opt/kubernetes/server/bin/kube-apiserver.sh            ; the program (relative uses PATH, can take args)
numprocs=1                                                      ; number of processes copies to start (def 1)
directory=/opt/kubernetes/server/bin                            ; directory to cwd to before exec (def no cwd)
autostart=true                                                  ; start at supervisord start (default: true)
autorestart=true                                                ; retstart at unexpected quit (default: true)
startsecs=30                                                    ; number of secs prog must stay running (def. 1)
startretries=3                                                  ; max # of serial start failures (default 3)
exitcodes=0,2                                                   ; 'expected' exit codes for process (default 0,2)
stopsignal=QUIT                                                 ; signal used to kill process (default TERM)
stopwaitsecs=10                                                 ; max num secs to wait b4 SIGKILL (default 10)
user=root                                                       ; setuid to this UNIX account to run the program
redirect_stderr=true                                            ; redirect proc stderr to stdout (default false)
killasgroup=true                                                ; kill all process in a group
stopasgroup=true                                                ; stop all process in a group
stdout_logfile=/data/logs/kubernetes/kube-apiserver/apiserver.stdout.log        ; stderr log path, NONE for none; default AUTO
stdout_logfile_maxbytes=64MB                                    ; max # logfile bytes b4 rotation (default 50MB)
stdout_logfile_backups=4                                        ; # of stdout logfile backups (default 10)
stdout_capture_maxbytes=1MB                                     ; number of bytes in 'capturemode' (default 0)
stdout_events_enabled=false                                     ; emit events on stdout writes (default false)
```

### 9.启动服务并检查

```shell
[root@shkf6-243 bin]# supervisorctl update
kube-apiserver-6-243: added process group
[root@shkf6-243 bin]# supervisorctl status
etcd-server-6-243                RUNNING   pid 12112, uptime 5:06:23
kube-apiserver-6-243             RUNNING   pid 12824, uptime 0:00:46

[root@shkf6-243 bin]# netstat -lntup|grep kube-apiser
tcp        0      0 127.0.0.1:8080          0.0.0.0:*               LISTEN      12825/./kube-apiser 
tcp6       0      0 :::6443                 :::*                    LISTEN      12825/./kube-apiser 
```

### 10.安装部署启动检查所有集群规划机器

```shell
[root@shkf6-244 src]# tar xf kubernetes-server-linux-amd64-v1.15.2.tar.gz -C /opt/
[root@shkf6-244 src]# cd /opt/
[root@shkf6-244 opt]# mv kubernetes kubernetes-v1.15.2
[root@shkf6-244 opt]# ln -s /opt/kubernetes-v1.15.2/ /opt/kubernetes
[root@shkf6-244 opt]# cd kubernetes
[root@shkf6-244 kubernetes]# rm -f kubernetes-src.tar.gz 
[root@shkf6-244 kubernetes]# cd server/bin
[root@shkf6-244 bin]# rm -f *.tar
[root@shkf6-244 bin]# rm -f *_tag
[root@shkf6-244 bin]# mkdir cert
[root@shkf6-244 bin]# scp -P52113 shkf6-245:/opt/certs/apiserver-key.pem /opt/kubernetes/server/bin/cert/
[root@shkf6-244 bin]# scp -P52113 shkf6-245:/opt/certs/apiserver.pem /opt/kubernetes/server/bin/cert/
[root@shkf6-244 bin]# scp -P52113 shkf6-245:/opt/certs/ca-key.pem /opt/kubernetes/server/bin/cert/
[root@shkf6-244 bin]# scp -P52113 shkf6-245:/opt/certs/ca.pem /opt/kubernetes/server/bin/cert/  
[root@shkf6-244 bin]# scp -P52113 shkf6-245:/opt/certs/client-key.pem /opt/kubernetes/server/bin/cert/
[root@shkf6-244 bin]# scp -P52113 shkf6-245:/opt/certs/client.pem /opt/kubernetes/server/bin/cert/
[root@shkf6-244 bin]# mkdir conf
[root@shkf6-244 bin]# vi conf/audit.yaml
[root@shkf6-244 bin]# cat conf/audit.yaml 
apiVersion: audit.k8s.io/v1beta1 # This is required.
kind: Policy
# Don't generate audit events for all requests in RequestReceived stage.
omitStages:
  - "RequestReceived"
rules:
  # Log pod changes at RequestResponse level
  - level: RequestResponse
    resources:
    - group: ""
      # Resource "pods" doesn't match requests to any subresource of pods,
      # which is consistent with the RBAC policy.
      resources: ["pods"]
  # Log "pods/log", "pods/status" at Metadata level
  - level: Metadata
    resources:
    - group: ""
      resources: ["pods/log", "pods/status"]

  # Don't log requests to a configmap called "controller-leader"
  - level: None
    resources:
    - group: ""
      resources: ["configmaps"]
      resourceNames: ["controller-leader"]

  # Don't log watch requests by the "system:kube-proxy" on endpoints or services
  - level: None
    users: ["system:kube-proxy"]
    verbs: ["watch"]
    resources:
    - group: "" # core API group
      resources: ["endpoints", "services"]

  # Don't log authenticated requests to certain non-resource URL paths.
  - level: None
    userGroups: ["system:authenticated"]
    nonResourceURLs:
    - "/api*" # Wildcard matching.
    - "/version"

  # Log the request body of configmap changes in kube-system.
  - level: Request
    resources:
    - group: "" # core API group
      resources: ["configmaps"]
    # This rule only applies to resources in the "kube-system" namespace.
    # The empty string "" can be used to select non-namespaced resources.
    namespaces: ["kube-system"]

  # Log configmap and secret changes in all other namespaces at the Metadata level.
  - level: Metadata
    resources:
    - group: "" # core API group
      resources: ["secrets", "configmaps"]

  # Log all other resources in core and extensions at the Request level.
  - level: Request
    resources:
    - group: "" # core API group
    - group: "extensions" # Version of group should NOT be included.

  # A catch-all rule to log all other requests at the Metadata level.
  - level: Metadata
    # Long-running requests like watches that fall under this rule will not
    # generate an audit event in RequestReceived.
    omitStages:
      - "RequestReceived"
[root@shkf6-244 bin]# vi /opt/kubernetes/server/bin/kube-apiserver.sh
[root@shkf6-244 bin]# cat /opt/kubernetes/server/bin/kube-apiserver.sh
#!/bin/bash
./kube-apiserver \
  --apiserver-count 2 \
  --audit-log-path /data/logs/kubernetes/kube-apiserver/audit-log \
  --audit-policy-file ./conf/audit.yaml \
  --authorization-mode RBAC \
  --client-ca-file ./cert/ca.pem \
  --requestheader-client-ca-file ./cert/ca.pem \
  --enable-admission-plugins NamespaceLifecycle,LimitRanger,ServiceAccount,DefaultStorageClass,DefaultTolerationSeconds,MutatingAdmissionWebhook,ValidatingAdmissionWebhook,ResourceQuota \
  --etcd-cafile ./cert/ca.pem \
  --etcd-certfile ./cert/client.pem \
  --etcd-keyfile ./cert/client-key.pem \
  --etcd-servers https://192.168.6.242:2379,https://192.168.6.243:2379,https://192.168.6.244:2379 \
  --service-account-key-file ./cert/ca-key.pem \
  --service-cluster-ip-range 10.96.0.0/22 \
  --service-node-port-range 3000-29999 \
  --target-ram-mb=1024 \
  --kubelet-client-certificate ./cert/client.pem \
  --kubelet-client-key ./cert/client-key.pem \
  --log-dir  /data/logs/kubernetes/kube-apiserver \
  --tls-cert-file ./cert/apiserver.pem \
  --tls-private-key-file ./cert/apiserver-key.pem \
  --v 2
[root@shkf6-244 bin]# chmod +x /opt/kubernetes/server/bin/kube-apiserver.sh
[root@shkf6-244 bin]# mkdir -p /data/logs/kubernetes/kube-apiserver
[root@shkf6-244 bin]# vi /etc/supervisord.d/kube-apiserver.ini
[root@shkf6-244 bin]# cat /etc/supervisord.d/kube-apiserver.ini
[program:kube-apiserver-6-244]
command=/opt/kubernetes/server/bin/kube-apiserver.sh            ; the program (relative uses PATH, can take args)
numprocs=1                                                      ; number of processes copies to start (def 1)
directory=/opt/kubernetes/server/bin                            ; directory to cwd to before exec (def no cwd)
autostart=true                                                  ; start at supervisord start (default: true)
autorestart=true                                                ; retstart at unexpected quit (default: true)
startsecs=30                                                    ; number of secs prog must stay running (def. 1)
startretries=3                                                  ; max # of serial start failures (default 3)
exitcodes=0,2                                                   ; 'expected' exit codes for process (default 0,2)
stopsignal=QUIT                                                 ; signal used to kill process (default TERM)
stopwaitsecs=10                                                 ; max num secs to wait b4 SIGKILL (default 10)
user=root                                                       ; setuid to this UNIX account to run the program
redirect_stderr=true                                            ; redirect proc stderr to stdout (default false)
killasgroup=true                                                ; kill all process in a group
stopasgroup=true                                                ; stop all process in a group
stdout_logfile=/data/logs/kubernetes/kube-apiserver/apiserver.stdout.log        ; stderr log path, NONE for none; default AUTO
stdout_logfile_maxbytes=64MB                                    ; max # logfile bytes b4 rotation (default 50MB)
stdout_logfile_backups=4                                        ; # of stdout logfile backups (default 10)
stdout_capture_maxbytes=1MB                                     ; number of bytes in 'capturemode' (default 0)
stdout_events_enabled=false                                     ; emit events on stdout writes (default false)
[root@shkf6-244 bin]# supervisorctl update
kube-apiserver-6-244: added process group
[root@shkf6-244 bin]# supervisorctl status
etcd-server-6-244                RUNNING   pid 11748, uptime 5:10:52
kube-apiserver-6-244             RUNNING   pid 12408, uptime 0:00:43
```

### 11.配四层反向代理

#### 1.部署nginx

在`shkf6-241`和`shkf6-242`上：

```shell
 ~]# yum install nginx -y
```

#### 2.配置4层代理

在`shkf6-241`和`shkf6-242`上：

```shell
 ~]# vim /etc/nginx/nginx.conf

stream {
    upstream kube-apiserver {
        server 192.168.6.243:6443     max_fails=3 fail_timeout=30s;
        server 192.168.6.244:6443     max_fails=3 fail_timeout=30s;
    }
    server {
        listen 7443;
        proxy_connect_timeout 2s;
        proxy_timeout 900s;
        proxy_pass kube-apiserver;
    }
}
```

#### 3.启动nginx

在`shkf6-241`和`shkf6-242`上：

```shell
 ~]# nginx -t
nginx: the configuration file /etc/nginx/nginx.conf syntax is ok
nginx: configuration file /etc/nginx/nginx.conf test is successful
 ~]# systemctl start nginx
 ~]# systemctl enable nginx
```

#### 4.部署keepalived服务

在`shkf6-241`和`shkf6-242`上：

```shell
[root@shkf6-241 ~]# yum install keepalived -y
```

#### 5.配置keepalived服务

在`shkf6-241`和`shkf6-242`上：

```shell
[root@shkf6-241 ~]# vi /etc/keepalived/check_port.sh
[root@shkf6-241 ~]# chmod +x /etc/keepalived/check_port.sh
[root@shkf6-241 ~]# cat /etc/keepalived/check_port.sh
#!/bin/bash
#keepalived 监控端口脚本
#使用方法：
#在keepalived的配置文件中
#vrrp_script check_port {#创建一个vrrp_script脚本,检查配置
#    script "/etc/keepalived/check_port.sh 6379" #配置监听的端口
#    interval 2 #检查脚本的频率,单位（秒）
#}
CHK_PORT=$1
if [ -n "$CHK_PORT" ];then
        PORT_PROCESS=`ss -lnt|grep $CHK_PORT|wc -l`
        if [ $PORT_PROCESS -eq 0 ];then
                echo "Port $CHK_PORT Is Not Used,End."
                exit 1
        fi
else
        echo "Check Port Cant Be Empty!"
fi
```

在`shkf6-241`上：

配置keepalived主：

```shell
[root@shkf6-241 ~]# vim /etc/keepalived/keepalived.conf 
[root@shkf6-241 ~]# cat /etc/keepalived/keepalived.conf
! Configuration File for keepalived

global_defs {
   router_id 192.168.6.241

}

vrrp_script chk_nginx {
    script "/etc/keepalived/check_port.sh 7443"
    interval 2
    weight -20
}

vrrp_instance VI_1 {
    state MASTER
    interface eth0
    virtual_router_id 251
    priority 100
    advert_int 1
    mcast_src_ip 192.168.6.241
    nopreempt

    authentication {
        auth_type PASS
        auth_pass 11111111
    }
    track_script {
         chk_nginx
    }
    virtual_ipaddress {
        192.168.6.66
    }
}
```

在`shkf6-242`上：

配置keepalived备：

```shell
[root@shkf6-242 ~]# vim /etc/keepalived/keepalived.conf 
[root@shkf6-242 ~]# cat /etc/keepalived/keepalived.conf
! Configuration File for keepalived
global_defs {
    router_id 192.168.6.242
}
vrrp_script chk_nginx {
    script "/etc/keepalived/check_port.sh 7443"
    interval 2
    weight -20
}
vrrp_instance VI_1 {
    state BACKUP
    interface eth0
    virtual_router_id 251
    mcast_src_ip 192.168.6.242
    priority 90
    advert_int 1
    authentication {
        auth_type PASS
        auth_pass 11111111
    }
    track_script {
        chk_nginx
    }
    virtual_ipaddress {
        192.168.6.66
    }
}
```

#### 6.启动keepalived

在`shkf6-241`和`shkf6-242`上：

```shell
 ~]# systemctl start keepalived.service 
 ~]# systemctl enable keepalived.service
```

#### 7.检查VIP

```shell
[root@shkf6-241 ~]# ip a |grep eth0
2: eth0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc mq state UP qlen 1000
    inet 192.168.6.241/24 brd 192.168.6.255 scope global eth0
    inet 192.168.6.66/32 scope global eth0
```

### 12.启动代理并检查

```shell
[root@shkf6-243 bin]# netstat -lntup|grep kube-apiser
tcp        0      0 127.0.0.1:8080          0.0.0.0:*               LISTEN      12825/./kube-apiser 
tcp6       0      0 :::6443                 :::*                    LISTEN      12825/./kube-apiser 

[root@shkf6-244 bin]# netstat -lntup|grep kube-apiser
tcp        0      0 127.0.0.1:8080          0.0.0.0:*               LISTEN      12409/./kube-apiser 
tcp6       0      0 :::6443                 :::*                    LISTEN      12409/./kube-apiser 

[root@shkf6-241 ~]# netstat -lntup|grep 7443
tcp        0      0 0.0.0.0:7443            0.0.0.0:*               LISTEN      12936/nginx: master

[root@shkf6-242 ~]# netstat -lntup|grep 7443
tcp        0      0 0.0.0.0:7443            0.0.0.0:*               LISTEN      11254/nginx: master
```

## 3.部署controller-manager

### 1.集群规划

| 主机名             | 角色               | ip            |
| :----------------- | :----------------- | :------------ |
| shkf6-243.host.com | controller-manager | 192.168.6.243 |
| shkf6-244.host.com | controller-manager | 192.168.6.244 |

注意：这里部署文档以`shkf6-243.host.com`主机为例，另外一台运算节点安装部署方法类似

### 2.创建启动脚本

shkf6-243上和shkf6-244：

```shell
[root@shkf6-243 bin]# cat /opt/kubernetes/server/bin/kube-controller-manager.sh
#!/bin/sh
./kube-controller-manager \
  --cluster-cidr 172.6.0.0/16 \
  --leader-elect true \
  --log-dir /data/logs/kubernetes/kube-controller-manager \
  --master http://127.0.0.1:8080 \
  --service-account-private-key-file ./cert/ca-key.pem \
  --service-cluster-ip-range 10.96.0.0/22 \
  --root-ca-file ./cert/ca.pem \
  --v 2
```

### 3.调整文件权限，创建目录

shkf6-243上和shkf6-244：

```shell
[root@shkf6-243 bin]# chmod +x /opt/kubernetes/server/bin/kube-controller-manager.sh
[root@shkf6-243 bin]# mkdir -p /data/logs/kubernetes/kube-controller-manager
```

### 4.创建supervisor配置

shkf6-243上和shkf6-244：

```shell
[root@shkf6-243 bin]# vi /etc/supervisord.d/kube-conntroller-manager.ini
[root@shkf6-243 bin]# cat /etc/supervisord.d/kube-conntroller-manager.ini
[program:kube-controller-manager-6.243]
command=/opt/kubernetes/server/bin/kube-controller-manager.sh                     ; the program (relative uses PATH, can take args)
numprocs=1                                                                        ; number of processes copies to start (def 1)
directory=/opt/kubernetes/server/bin                                              ; directory to cwd to before exec (def no cwd)
autostart=true                                                                    ; start at supervisord start (default: true)
autorestart=true                                                                  ; retstart at unexpected quit (default: true)
startsecs=30                                                                      ; number of secs prog must stay running (def. 1)
startretries=3                                                                    ; max # of serial start failures (default 3)
exitcodes=0,2                                                                     ; 'expected' exit codes for process (default 0,2)
stopsignal=QUIT                                                                   ; signal used to kill process (default TERM)
stopwaitsecs=10                                                                   ; max num secs to wait b4 SIGKILL (default 10)
user=root                                                                         ; setuid to this UNIX account to run the program
redirect_stderr=true                                                              ; redirect proc stderr to stdout (default false)
killasgroup=true                                                                  ; kill all process in a group
stopasgroup=true                                                                  ; stop all process in a group
stdout_logfile=/data/logs/kubernetes/kube-controller-manager/controller.stdout.log  ; stderr log path, NONE for none; default AUTO
stdout_logfile_maxbytes=64MB                                                      ; max # logfile bytes b4 rotation (default 50MB)
stdout_logfile_backups=4                                                          ; # of stdout logfile backups (default 10)
stdout_capture_maxbytes=1MB                                                       ; number of bytes in 'capturemode' (default 0)
stdout_events_enabled=false                                                       ; emit events on stdout writes (default false)
```

### 5.启动服务并检查

shkf6-243上和shkf6-244：

```shell
[root@shkf6-243 bin]# supervisorctl update
kube-controller-manager-6.243: added process group
[root@shkf6-243 bin]# supervisorctl status
etcd-server-6-243                RUNNING   pid 12112, uptime 23:36:23
kube-apiserver-6-243             RUNNING   pid 12824, uptime 18:30:46
kube-controller-manager-6.243    RUNNING   pid 14952, uptime 0:01:00
```

### 6.安装部署启动检查所有集群规划主机的kube-controller-manager服务

略

## 4.部署kube-scheduler

### 1.集群规划

| 主机名             | 角色                    | ip            |
| :----------------- | :---------------------- | :------------ |
| shkf6-243.host.com | kube-controller-manager | 192.168.6.243 |
| shkf6-244.host.com | kube-controller-manager | 192.168.6.244 |

注意：这里部署文档以`shkf6-243.host.com`主机为例，另外一台运算节点安装部署方法类似

### 2.创建启动脚本

在shkf6-243和shkf6-244上：

```shell
[root@shkf6-243 bin]# vi /opt/kubernetes/server/bin/kube-scheduler.sh
[root@shkf6-243 bin]# cat /opt/kubernetes/server/bin/kube-scheduler.sh
#!/bin/sh
./kube-scheduler \
  --leader-elect  \
  --log-dir /data/logs/kubernetes/kube-scheduler \
  --master http://127.0.0.1:8080 \
  --v 2
```

### 3.调整文件权限，创建目录

在shkf6-243和shkf6-244上：

```shell
[root@shkf6-243 bin]# chmod +x /opt/kubernetes/server/bin/kube-scheduler.sh
[root@shkf6-243 bin]# mkdir -p /data/logs/kubernetes/kube-scheduler
```

### 4.创建supervisor配置

在shkf6-243和shkf6-244上：

```shell
[root@shkf6-243 bin]# vi /etc/supervisord.d/kube-scheduler.ini
[root@shkf6-243 bin]# cat /etc/supervisord.d/kube-scheduler.ini
[program:kube-scheduler-6-243]
command=/opt/kubernetes/server/bin/kube-scheduler.sh                     ; the program (relative uses PATH, can take args)
numprocs=1                                                               ; number of processes copies to start (def 1)
directory=/opt/kubernetes/server/bin                                     ; directory to cwd to before exec (def no cwd)
autostart=true                                                           ; start at supervisord start (default: true)
autorestart=true                                                         ; retstart at unexpected quit (default: true)
startsecs=30                                                             ; number of secs prog must stay running (def. 1)
startretries=3                                                           ; max # of serial start failures (default 3)
exitcodes=0,2                                                            ; 'expected' exit codes for process (default 0,2)
stopsignal=QUIT                                                          ; signal used to kill process (default TERM)
stopwaitsecs=10                                                          ; max num secs to wait b4 SIGKILL (default 10)
user=root                                                                ; setuid to this UNIX account to run the program
redirect_stderr=true                                                     ; redirect proc stderr to stdout (default false)
killasgroup=true                                                         ; kill all process in a group
stopasgroup=true                                                         ; stop all process in a group
stdout_logfile=/data/logs/kubernetes/kube-scheduler/scheduler.stdout.log ; stderr log path, NONE for none; default AUTO
stdout_logfile_maxbytes=64MB                                             ; max # logfile bytes b4 rotation (default 50MB)
stdout_logfile_backups=4                                                 ; # of stdout logfile backups (default 10)
stdout_capture_maxbytes=1MB                                              ; number of bytes in 'capturemode' (default 0)
stdout_events_enabled=false                                              ; emit events on stdout writes (default false)
```

### 5.启动服务并检查

在shkf6-243和shkf6-244上：

```shell
[root@shkf6-243 bin]# supervisorctl update
kube-scheduler-6-243: added process group
[root@shkf6-243 bin]# supervisorctl status
etcd-server-6-243                RUNNING   pid 12112, uptime 23:52:01
kube-apiserver-6-243             RUNNING   pid 12824, uptime 18:46:24
kube-controller-manager-6.243    RUNNING   pid 14952, uptime 0:16:38
kube-scheduler-6-243             RUNNING   pid 15001, uptime 0:01:39
[root@shkf6-243 bin]# ln -s /opt/kubernetes/server/bin/kubectl /usr/bin/kubectl
```

### 6.kubect1命令自动补全

各运算节点上:

```shell
~]# yum install bash-completion -y
~]# kubectl completion bash > /etc/bash_completion.d/kubectl
```

重新登录终端即可

### 7.安装部署启动检查所有集群规划主机的kube-controller-manager服务

略

## 5.检查主控节点

在shkf6-243和shkf6-244上：

```shell
[root@shkf6-243 bin]# which kubectl 
/usr/bin/kubectl
[root@shkf6-243 bin]# kubectl get cs
NAME                 STATUS    MESSAGE              ERROR
controller-manager   Healthy   ok                   
scheduler            Healthy   ok                   
etcd-1               Healthy   {"health": "true"}   
etcd-0               Healthy   {"health": "true"}   
etcd-2               Healthy   {"health": "true"} 
```

# 第六章：部署运算节点服务

## 1.部署kubelet

### 1.集群规划

| 主机名             | 角色    | ip            |
| :----------------- | :------ | :------------ |
| shkf6-243.host.com | kubelet | 192.168.6.243 |
| shkf6-244.host.com | kubelet | 192.168.6.244 |

注意：这里部署文档以`shkf6-243.host.com`主机为例，另外一台运算节点安装部署方法类似

### 2.签发kubelet证书

运维主机shkf6-245.host.com上：

#### 1.创建生成正事签名请求（csr）的JSON配置文件

```shell
[root@shkf6-245 certs]# vi /opt/certs/kubelet-csr.json
[root@shkf6-245 certs]# cat /opt/certs/kubelet-csr.json

{
    "CN": "k8s-kubelet",
    "hosts": [
    "127.0.0.1",
    "192.168.6.66",
    "192.168.6.243",
    "192.168.6.244",
    "192.168.6.245",
    "192.168.6.246",
    "192.168.6.247",
    "192.168.6.248"
    ],
    "key": {
        "algo": "rsa",
        "size": 2048
    },
    "names": [
        {
            "C": "CN",
            "ST": "beijing",
            "L": "beijing",
            "O": "od",
            "OU": "ops"
        }
    ]
}
```

注意：把所有有可能用到的kubulet主机全加进去

#### 2.生成kubelet证书和私钥

```shell
[root@shkf6-245 certs]# cfssl gencert -ca=ca.pem -ca-key=ca-key.pem -config=ca-config.json -profile=server kubelet-csr.json | cfssl-json -bare kubelet
2019/11/15 09:49:58 [INFO] generate received request
2019/11/15 09:49:58 [INFO] received CSR
2019/11/15 09:49:58 [INFO] generating key: rsa-2048
2019/11/15 09:49:59 [INFO] encoded CSR
2019/11/15 09:49:59 [INFO] signed certificate with serial number 609294877015122932833154151112494803106290808681
2019/11/15 09:49:59 [WARNING] This certificate lacks a "hosts" field. This makes it unsuitable for
websites. For more information see the Baseline Requirements for the Issuance and Management
of Publicly-Trusted Certificates, v.1.1.6, from the CA/Browser Forum (https://cabforum.org);
specifically, section 10.2.3 ("Information Requirements").
```

#### 3.检查生成证书的证书、私钥

```shell
[root@shkf6-245 certs]# ll kubelet*
-rw-r--r-- 1 root root 1098 Nov 15 09:49 kubelet.csr
-rw-r--r-- 1 root root  445 Nov 15 09:47 kubelet-csr.json
-rw------- 1 root root 1675 Nov 15 09:49 kubelet-key.pem
-rw-r--r-- 1 root root 1452 Nov 15 09:49 kubelet.pem
```

### 3.拷贝证书至各运算节点，并创建配置

shkf6-243上：

#### 1.拷贝证书，私钥，注意私钥文件属性600

```shell
[root@shkf6-243 bin]#  scp -P52113 shkf6-245:/opt/certs/kubelet-key.pem /opt/kubernetes/server/bin/cert/
[root@shkf6-243 bin]#  scp -P52113 shkf6-245:/opt/certs/kubelet.pem /opt/kubernetes/server/bin/cert/
[root@shkf6-243 bin]# ll cert/
total 32
-rw------- 1 root root 1679 Nov 14 14:18 apiserver-key.pem
-rw-r--r-- 1 root root 1598 Nov 14 14:18 apiserver.pem
-rw------- 1 root root 1679 Nov 14 14:18 ca-key.pem
-rw-r--r-- 1 root root 1346 Nov 14 14:19 ca.pem
-rw------- 1 root root 1679 Nov 14 14:19 client-key.pem
-rw-r--r-- 1 root root 1363 Nov 14 14:19 client.pem
-rw------- 1 root root 1675 Nov 15 10:01 kubelet-key.pem
-rw-r--r-- 1 root root 1452 Nov 15 10:02 kubelet.pem
```

#### 2.创建配置

##### 1.set-cluster

注意：在conf目录下

```shell
[root@shkf6-243 conf]# kubectl config set-cluster myk8s \
  --certificate-authority=/opt/kubernetes/server/bin/cert/ca.pem \
  --embed-certs=true \
  --server=https://192.168.6.66:7443 \
  --kubeconfig=kubelet.kubeconfig

Cluster "myk8s" set.
```

##### 2.set-credentials

注意：在conf目录下

```shell
[root@shkf6-243 conf]# kubectl config set-credentials k8s-node \
  --client-certificate=/opt/kubernetes/server/bin/cert/client.pem \
  --client-key=/opt/kubernetes/server/bin/cert/client-key.pem \
  --embed-certs=true \
  --kubeconfig=kubelet.kubeconfig 

User "k8s-node" set.
```

##### 3.set-context

注意：在conf目录下

```shell
[root@shkf6-243 conf]# kubectl config set-context myk8s-context \
  --cluster=myk8s \
  --user=k8s-node \
  --kubeconfig=kubelet.kubeconfig

Context "myk8s-context" created.
```

##### 4.use-context

注意：在conf目录下

```shell
[root@shkf6-243 conf]# kubectl config use-context myk8s-context --kubeconfig=kubelet.kubeconfig

Switched to context "myk8s-context".
```

##### 5.k8s-node.yaml

- 创建资源配置文件

```shell
[root@shkf6-243 conf]# vim /opt/kubernetes/server/bin/conf/k8s-node.yaml
[root@shkf6-243 conf]# cat /opt/kubernetes/server/bin/conf/k8s-node.yaml 
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: k8s-node
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: system:node
subjects:
- apiGroup: rbac.authorization.k8s.io
  kind: User
  name: k8s-node
```

- 是集群角色用户生效

```shell
[root@shkf6-243 conf]# kubectl create -f k8s-node.yaml
```

- 查看集群角色

```shell
[root@shkf6-243 conf]# kubectl get clusterrolebinding k8s-node
NAME       AGE
k8s-node   22m
[root@shkf6-243 conf]# kubectl get clusterrolebinding k8s-node -o yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  creationTimestamp: "2019-11-15T02:14:34Z"
  name: k8s-node
  resourceVersion: "17884"
  selfLink: /apis/rbac.authorization.k8s.io/v1/clusterrolebindings/k8s-node
  uid: e09ed617-936f-4936-8adc-d7cc9b3bce63
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: system:node
subjects:
- apiGroup: rbac.authorization.k8s.io
  kind: User
  name: k8s-node
```

在shkf6-244上：

```shell
[root@shkf6-244 bin]# scp -P52113 shkf6-243:/opt/kubernetes/server/bin/conf/kubelet.kubeconfig /opt/kubernetes/server/bin/conf/
```

可略

```shell
[root@shkf6-244 bin]# scp -P52113 shkf6-243:/opt/kubernetes/server/bin/conf/k8s-node.yaml /opt/kubernetes/server/bin/conf/
```

### 4.准备pause基础镜像

运维主机shkf6-244.host.com上：

#### 1.下载

```shell
[root@shkf6-245 certs]# docker pull kubernetes/pause
Using default tag: latest
latest: Pulling from kubernetes/pause
4f4fb700ef54: Pull complete 
b9c8ec465f6b: Pull complete 
Digest: sha256:b31bfb4d0213f254d361e0079deaaebefa4f82ba7aa76ef82e90b4935ad5b105
Status: Downloaded newer image for kubernetes/pause:latest
docker.io/kubernetes/pause:latest
```

#### 2.打标签

```shell
[root@shkf6-245 certs]# docker images|grep pause
kubernetes/pause                latest                     f9d5de079539        5 years ago         240kB
[root@shkf6-245 certs]# docker tag f9d5de079539 harbor.od.com/public/pause:latest
```

#### 3.推送私有仓库(harbor)中

```shell
[root@shkf6-245 certs]# docker push harbor.od.com/public/pause:latest
The push refers to repository [harbor.od.com/public/pause]
5f70bf18a086: Mounted from public/nginx 
e16a89738269: Pushed 
latest: digest: sha256:b31bfb4d0213f254d361e0079deaaebefa4f82ba7aa76ef82e90b4935ad5b105 size: 938
```

### 5.创建kubelet启动脚本

在shkf6-243：

```shell
[root@shkf6-243 conf]# vi /opt/kubernetes/server/bin/kubelet.sh
[root@shkf6-243 conf]# cat /opt/kubernetes/server/bin/kubelet.sh
#!/bin/sh
./kubelet \
  --anonymous-auth=false \
  --cgroup-driver systemd \
  --cluster-dns 10.96.0.2 \
  --cluster-domain cluster.local \
  --runtime-cgroups=/systemd/system.slice \
  --kubelet-cgroups=/systemd/system.slice \
  --fail-swap-on="false" \
  --client-ca-file ./cert/ca.pem \
  --tls-cert-file ./cert/kubelet.pem \
  --tls-private-key-file ./cert/kubelet-key.pem \
  --hostname-override shkf6-243.host.com \
  --image-gc-high-threshold 20 \
  --image-gc-low-threshold 10 \
  --kubeconfig ./conf/kubelet.kubeconfig \
  --log-dir /data/logs/kubernetes/kube-kubelet \
  --pod-infra-container-image harbor.od.com/public/pause:latest \
  --root-dir /data/kubelet
```

注意：kubelet集群各主机的启动脚本略有不同，部署其节点时注意修改

```
hostname-override
```

### 6.检查配置，权限，创建日志目录

在shkf6-243：

```shell
[root@shkf6-243 conf]# ls -l|grep kubelet.kubeconfig 
-rw------- 1 root root 6202 Nov 15 10:11 kubelet.kubeconfig

[root@shkf6-243 conf]# chmod +x /opt/kubernetes/server/bin/kubelet.sh
[root@shkf6-243 conf]# mkdir -p /data/logs/kubernetes/kube-kubelet /data/kubelet
```

### 7.创建supervisor配置

在shkf6-243：

```shell
[root@shkf6-243 conf]# vi /etc/supervisord.d/kube-kubelet.ini
[root@shkf6-243 conf]# cat /etc/supervisord.d/kube-kubelet.ini

[program:kube-kubelet-6-243]
command=/opt/kubernetes/server/bin/kubelet.sh     ; the program (relative uses PATH, can take args)
numprocs=1                                        ; number of processes copies to start (def 1)
directory=/opt/kubernetes/server/bin              ; directory to cwd to before exec (def no cwd)
autostart=true                                    ; start at supervisord start (default: true)
autorestart=true                                ; retstart at unexpected quit (default: true)
startsecs=30                                      ; number of secs prog must stay running (def. 1)
startretries=3                                    ; max # of serial start failures (default 3)
exitcodes=0,2                                     ; 'expected' exit codes for process (default 0,2)
stopsignal=QUIT                                   ; signal used to kill process (default TERM)
stopwaitsecs=10                                   ; max num secs to wait b4 SIGKILL (default 10)
user=root                                         ; setuid to this UNIX account to run the program
redirect_stderr=true                              ; redirect proc stderr to stdout (default false)
killasgroup=true                                  ; kill all process in a group
stopasgroup=true                                  ; stop all process in a group
stdout_logfile=/data/logs/kubernetes/kube-kubelet/kubelet.stdout.log   ; stderr log path, NONE for none; default AUTO
stdout_logfile_maxbytes=64MB                      ; max # logfile bytes b4 rotation (default 50MB)
stdout_logfile_backups=4                          ; # of stdout logfile backups (default 10)
stdout_capture_maxbytes=1MB                       ; number of bytes in 'capturemode' (default 0)
stdout_events_enabled=false                       ; emit events on stdout writes (default false)
```

### 8.启动服务并检查

- 启动服务并检查

```shell
[root@shkf6-243 conf]# supervisorctl update
kube-kubelet-6-243: added process group
[root@shkf6-243 conf]# supervisorctl status
etcd-server-6-243                RUNNING   pid 12112, uptime 1 day, 1:43:37
kube-apiserver-6-243             RUNNING   pid 12824, uptime 20:38:00
kube-controller-manager-6.243    RUNNING   pid 14952, uptime 2:08:14
kube-kubelet-6-243               RUNNING   pid 15398, uptime 0:01:25
kube-scheduler-6-243             RUNNING   pid 15001, uptime 1:53:15

[root@shkf6-243 conf]# tail -fn 200 /data/logs/kubernetes/kube-kubelet/kubelet.stdout.log
```

### 9.检查运算节点

```shell
[root@shkf6-243 conf]# kubectl get nodes
NAME                 STATUS   ROLES    AGE     VERSION
shkf6-243.host.com   Ready    <none>   16m     v1.15.2
shkf6-244.host.com   Ready    <none>   2m12s   v1.15.2
[root@shkf6-243 conf]# kubectl get nodes -o wide
NAME                 STATUS   ROLES    AGE     VERSION   INTERNAL-IP     EXTERNAL-IP   OS-IMAGE                KERNEL-VERSION               CONTAINER-RUNTIME
shkf6-243.host.com   Ready    <none>   17m     v1.15.2   192.168.6.243   <none>        CentOS Linux 7 (Core)   3.10.0-693.21.1.el7.x86_64   docker://19.3.4
shkf6-244.host.com   Ready    <none>   2m45s   v1.15.2   192.168.6.244   <none>        CentOS Linux 7 (Core)   3.10.0-693.21.1.el7.x86_64   docker://19.3.4
```

- 给node节点打标签

  ```shell
  [root@shkf6-243 conf]# kubectl label node shkf6-243.host.com node-role.kubernetes.io/node=
  node/shkf6-243.host.com labeled
  [root@shkf6-243 conf]# kubectl label node shkf6-243.host.com node-role.kubernetes.io/master=
  node/shkf6-243.host.com labeled
  [root@shkf6-243 conf]# kubectl label node shkf6-244.host.com node-role.kubernetes.io/node=
  node/shkf6-244.host.com labeled
  [root@shkf6-243 conf]# kubectl label node shkf6-244.host.com node-role.kubernetes.io/master=
  node/shkf6-244.host.com labeled
  ```

### 10.安装部署启动检查所有集群规划主机的kube-kubelet服务

略

### 11.检查所有运算节点

```shell
[root@shkf6-243 conf]# kubectl get nodes -o wide
NAME                 STATUS   ROLES         AGE     VERSION   INTERNAL-IP     EXTERNAL-IP   OS-IMAGE                KERNEL-VERSION               CONTAINER-RUNTIME
shkf6-243.host.com   Ready    master,node   20m     v1.15.2   192.168.6.243   <none>        CentOS Linux 7 (Core)   3.10.0-693.21.1.el7.x86_64   docker://19.3.4
shkf6-244.host.com   Ready    master,node   6m34s   v1.15.2   192.168.6.244   <none>        CentOS Linux 7 (Core)   3.10.0-693.21.1.el7.x86_64   docker://19.3.4
[root@shkf6-243 conf]# kubectl get nodes
NAME                 STATUS   ROLES         AGE     VERSION
shkf6-243.host.com   Ready    master,node   21m     v1.15.2
shkf6-244.host.com   Ready    master,node   6m42s   v1.15.2
```

## 2.部署kube-proxy

### 1.集群规划

| 主机名             | 角色       | ip            |
| :----------------- | :--------- | :------------ |
| shkf6-243.host.com | kube-proxy | 192.168.6.243 |
| shkf6-244.host.com | kube-proxy | 192.168.6.244 |

注意：这里部署文档以shkf6-243.host.com主机为例，另外一台运算节点安装部署方法类似

### 2.签发kube-proxy证书

运维主机shkf6-245.host.com上：

#### 1.创建生成证书签名请求（csr）的JSON文件

```shell
[root@shkf6-245 certs]# vi /opt/certs/kube-proxy-csr.json
[root@shkf6-245 certs]# cat kube-proxy-csr.json 
{
    "CN": "system:kube-proxy",
    "key": {
        "algo": "rsa",
        "size": 2048
    },
    "names": [
        {
            "C": "CN",
            "ST": "beijing",
            "L": "beijing",
            "O": "od",
            "OU": "ops"
        }
    ]
}
```

#### 2.生成kubelet证书和私钥

```shell
[root@shkf6-245 certs]# cfssl gencert -ca=ca.pem -ca-key=ca-key.pem -config=ca-config.json -profile=client kube-proxy-csr.json |cfssl-json -bare kube-proxy-client
2019/11/15 12:28:23 [INFO] generate received request
2019/11/15 12:28:23 [INFO] received CSR
2019/11/15 12:28:23 [INFO] generating key: rsa-2048
2019/11/15 12:28:24 [INFO] encoded CSR
2019/11/15 12:28:24 [INFO] signed certificate with serial number 499210659443234759487015805632579178164834077987
2019/11/15 12:28:24 [WARNING] This certificate lacks a "hosts" field. This makes it unsuitable for
websites. For more information see the Baseline Requirements for the Issuance and Management
of Publicly-Trusted Certificates, v.1.1.6, from the CA/Browser Forum (https://cabforum.org);
specifically, section 10.2.3 ("Information Requirements").
```

注意：这里的clent不能与其他的通用，上面CN变了，`"CN": "system:kube-proxy",`

#### 3.检查生成证书的证书、私钥

```shell
[root@shkf6-245 certs]# ll kube-proxy*
-rw-r--r-- 1 root root 1005 Nov 15 12:28 kube-proxy-client.csr
-rw------- 1 root root 1675 Nov 15 12:28 kube-proxy-client-key.pem
-rw-r--r-- 1 root root 1375 Nov 15 12:28 kube-proxy-client.pem
-rw-r--r-- 1 root root  267 Nov 15 12:28 kube-proxy-csr.json
```

### 3.拷贝证书至各个运算节点，并创建配置

#### 1.拷贝`kube-proxy-client-key.pem`和`kube-proxy-client.pem`至运算节点

```shell
[root@shkf6-243 conf]# scp -P52113 shkf6-245:/opt/certs/kube-proxy-client-key.pem /opt/kubernetes/server/bin/cert/
[root@shkf6-243 conf]# scp -P52113 shkf6-245:/opt/certs/kube-proxy-client.pem /opt/kubernetes/server/bin/cert/
```

#### 2.创建配置

##### 1.set-cluster

注意：在conf目录下

```shell
[root@shkf6-243 conf]# kubectl config set-cluster myk8s \
  --certificate-authority=/opt/kubernetes/server/bin/cert/ca.pem \
  --embed-certs=true \
  --server=https://192.168.6.66:7443 \
  --kubeconfig=kube-proxy.kubeconfig
```

##### 2.set-credentials

注意：在conf目录下

```shell
[root@shkf6-243 conf]# kubectl config set-credentials kube-proxy \
  --client-certificate=/opt/kubernetes/server/bin/cert/kube-proxy-client.pem \
  --client-key=/opt/kubernetes/server/bin/cert/kube-proxy-client-key.pem \
  --embed-certs=true \
  --kubeconfig=kube-proxy.kubeconfig
```

##### 3.set-context

注意：在conf目录下

```shell
[root@shkf6-243 conf]# kubectl config set-context myk8s-context \
  --cluster=myk8s \
  --user=kube-proxy \
  --kubeconfig=kube-proxy.kubeconfig
```

##### 4.use-context

注意：在conf目录下

```shell
[root@shkf6-243 conf]# kubectl config use-context myk8s-context --kubeconfig=kube-proxy.kubeconfig
```

### 4.创建kube-proxy启动脚本

在shkf6-243上：

- 加载ipvs模块

```shell
[root@shkf6-243 conf]# vi /root/ipvs.sh
[root@shkf6-243 conf]# cat /root/ipvs.sh
#!/bin/bash
ipvs_mods_dir="/usr/lib/modules/$(uname -r)/kernel/net/netfilter/ipvs"
for i in $(ls $ipvs_mods_dir|grep -o "^[^.]*")
do
  /sbin/modinfo -F filename $i &>/dev/null
  if [ $? -eq 0 ];then
    /sbin/modprobe $i
  fi
done

[root@shkf6-243 bin]# sh /root/ipvs.sh 

[root@shkf6-243 bin]# lsmod |grep ip_vs
[root@shkf6-243 bin]# lsmod |grep ip_vs
ip_vs_wlc              12519  0 
ip_vs_sed              12519  0 
ip_vs_pe_sip           12697  0 
nf_conntrack_sip       33860  1 ip_vs_pe_sip
ip_vs_nq               12516  0 
ip_vs_lc               12516  0 
ip_vs_lblcr            12922  0 
ip_vs_lblc             12819  0 
ip_vs_ftp              13079  0 
ip_vs_dh               12688  0 
ip_vs_sh               12688  0 
ip_vs_wrr              12697  0 
ip_vs_rr               12600  0 
ip_vs                 141092  24 ip_vs_dh,ip_vs_lc,ip_vs_nq,ip_vs_rr,ip_vs_sh,ip_vs_ftp,ip_vs_sed,ip_vs_wlc,ip_vs_wrr,ip_vs_pe_sip,ip_vs_lblcr,ip_vs_lblc
nf_nat                 26787  3 ip_vs_ftp,nf_nat_ipv4,nf_nat_masquerade_ipv4
nf_conntrack          133387  8 ip_vs,nf_nat,nf_nat_ipv4,xt_conntrack,nf_nat_masquerade_ipv4,nf_conntrack_netlink,nf_conntrack_sip,nf_conntrack_ipv4
libcrc32c              12644  4 xfs,ip_vs,nf_nat,nf_conntrack
```

- 创建启动脚本

```shell
[root@shkf6-243 conf]# vi /opt/kubernetes/server/bin/kube-proxy.sh
[root@shkf6-243 conf]# cat /opt/kubernetes/server/bin/kube-proxy.sh
#!/bin/sh
./kube-proxy \
  --cluster-cidr 172.6.0.0/16 \
  --hostname-override shkf6-243.host.com \
  --proxy-mode=ipvs \
  --ipvs-scheduler=nq \
  --kubeconfig ./conf/kube-proxy.kubeconfig
```

注意：kube-proxy集群各主机的启动脚本略有不同，部署其他节点时注意修改。

### 5.检查配置，权限，创建日志目录

在shkf6-243上：

```shell
[root@shkf6-243 conf]# ls -l|grep kube-proxy.kubeconfig 
-rw------- 1 root root 6207 Nov 15 12:32 kube-proxy.kubeconfig

[root@shkf6-243 conf]# chmod +x /opt/kubernetes/server/bin/kube-proxy.sh
[root@shkf6-243 conf]# mkdir -p /data/logs/kubernetes/kube-proxy
```

### 6.创建supervisor配置

```shell
[root@shkf6-243 conf]# vi /etc/supervisord.d/kube-proxy.ini
[root@shkf6-243 conf]# cat /etc/supervisord.d/kube-proxy.ini
[program:kube-proxy-6-243]
command=/opt/kubernetes/server/bin/kube-proxy.sh                     ; the program (relative uses PATH, can take args)
numprocs=1                                                           ; number of processes copies to start (def 1)
directory=/opt/kubernetes/server/bin                                 ; directory to cwd to before exec (def no cwd)
autostart=true                                                       ; start at supervisord start (default: true)
autorestart=true                                                     ; retstart at unexpected quit (default: true)
startsecs=30                                                         ; number of secs prog must stay running (def. 1)
startretries=3                                                       ; max # of serial start failures (default 3)
exitcodes=0,2                                                        ; 'expected' exit codes for process (default 0,2)
stopsignal=QUIT                                                      ; signal used to kill process (default TERM)
stopwaitsecs=10                                                      ; max num secs to wait b4 SIGKILL (default 10)
user=root                                                            ; setuid to this UNIX account to run the program
redirect_stderr=true                                                 ; redirect proc stderr to stdout (default false)
killasgroup=true                                                     ; kill all process in a group
stopasgroup=true                                                     ; stop all process in a group
stdout_logfile=/data/logs/kubernetes/kube-proxy/proxy.stdout.log     ; stderr log path, NONE for none; default AUTO
stdout_logfile_maxbytes=64MB                                         ; max # logfile bytes b4 rotation (default 50MB)
stdout_logfile_backups=4                                             ; # of stdout logfile backups (default 10)
stdout_capture_maxbytes=1MB                                          ; number of bytes in 'capturemode' (default 0)
stdout_events_enabled=false                                          ; emit events on stdout writes (default false)
```

### 7.启动服务并检查

```shell
[root@shkf6-243 conf]# supervisorctl update
kube-proxy-6-243: added process group

[root@shkf6-243 conf]# supervisorctl status
etcd-server-6-243                RUNNING   pid 12112, uptime 1 day, 8:13:06
kube-apiserver-6-243             RUNNING   pid 12824, uptime 1 day, 3:07:29
kube-controller-manager-6.243    RUNNING   pid 14952, uptime 8:37:43
kube-kubelet-6-243               RUNNING   pid 15398, uptime 6:30:54
kube-proxy-6-243                 RUNNING   pid 8055, uptime 0:01:19
kube-scheduler-6-243             RUNNING   pid 15001, uptime 8:22:44
[root@shkf6-243 conf]# yum install ipvsadm -y
[root@shkf6-243 ~]# ipvsadm -Ln
IP Virtual Server version 1.2.1 (size=4096)
Prot LocalAddress:Port Scheduler Flags
  -> RemoteAddress:Port           Forward Weight ActiveConn InActConn
TCP  10.96.0.1:443 nq
  -> 192.168.6.243:6443           Masq    1      0          0         
  -> 192.168.6.244:6443           Masq    1      0          0 

[root@shkf6-243 ~]# kubectl get svc
NAME         TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
kubernetes   ClusterIP   10.96.0.1    <none>        443/TCP   24h
```

# 第七章：完成部署并验证集群

- 创建配置清单

```shell
[root@shkf6-243 conf]# cat /root/nginx-ds.yaml 
apiVersion: extensions/v1beta1
kind: DaemonSet
metadata:
  name: nginx-ds
spec:
  template:
    metadata:
      labels:
        app: nginx-ds
    spec:
      containers:
      - name: my-nginx
        image: harbor.od.com/public/nginx:v1.7.9
        ports:
        - containerPort: 80
```

- 集群运算节点登录harbor

```shell
[root@shkf6-243 conf]# docker login harbor.od.com
Username: admin  
Password: 

[root@shkf6-244 conf]# docker login harbor.od.com
Username: admin
Password:
```

- 创建pod

```shell
[root@shkf6-243 conf]# kubectl create -f nginx-ds.yaml
```

- 创建pod

```shell
[root@shkf6-243 conf]# kubectl get pods 
NAME             READY   STATUS    RESTARTS   AGE
nginx-ds-ftxpz   1/1     Running   0          2m50s
nginx-ds-wb6wt   1/1     Running   0          2m51s

[root@shkf6-243 conf]# kubectl get cs
NAME                 STATUS    MESSAGE              ERROR
scheduler            Healthy   ok                   
controller-manager   Healthy   ok                   
etcd-0               Healthy   {"health": "true"}   
etcd-1               Healthy   {"health": "true"}   
etcd-2               Healthy   {"health": "true"}
```

# 第八章：资源需求说明