

#### docker的架构组成

![在这里插入图片描述](https://img-blog.csdnimg.cn/20210601165408954.png)

```
docker daemon：安装使用docker得先运行docker daemon进程，用于管理docker，如：镜像，容器，网络数据集。
rest接口：提供daemon交互的API接口
docker client：客户端使用restapi 和docker daemon进行访问
images：镜像是一个只读模板，用于创建容器，也可以通过dockerfile文件描述镜像的内容
Registry：docker镜像需要进行管理，docker提供了registry（注册表）仓库，其实他是一个容器，可以用于基于该容器运行私有仓库。也可以使用docker hub 联网公有仓库。
container：容器是一个镜像的运行实例。
```



#### 创建容器的过程

```
获取镜像，docker pull centos 从镜像仓库拉取
使用镜像创建容器
分配文件系统，挂着读写从，在读写从加载镜像
分配网络/网桥接口，创建一个网络接口，让容器和宿主机通信
容器获取IP
执行容器命令
反馈容器启动结果
```



#### 小结

```
Docker是一种CS架构的软件产品，可以把代码及依赖打包成镜像，作为交付介质，并且把镜像启动成为容器，提供容器生命周期的管理

docker组成：docker daemon | restapi接口 | docker client | iamges | registry | container
```

需要掌握的内容

```
了解docker理念，熟悉docker架构组成
```



#### docker的安装与部署

安装

环境

```
配置好防火墙及selinux，同步时间，修改hostname,可通外网
```

#### 配置宿主机网卡转发

```bash
# 若没有配置，需执行如下
cat <<EOF >  /etc/sysctl.d/docker.conf
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
net.ipv4.ip_forward=1
EOF

# sysctl修改内核运行时参数，-p从配置文件/etc/sysctl.d/加载内核参数设置
sysctl -p /etc/sysctl.d/docker.conf

```

##### yum安装docker

```bash
2.# 下载阿里源的repo文件
先备份
mkdir /etc/yum.repos.d/bak
mv /etc/yum.repos.d/*.repo  /etc/yum.repos.d/bak

curl -o /etc/yum.repos.d/Centos-7.repo http://mirrors.aliyun.com/repo/Centos-7.repo
curl -o /etc/yum.repos.d/docker-ce.repo http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo
yum clean all && yum makecache

2.# yum 安装
yum install docker-ce -y

# 查看源中可用版本
yum list docker-ce --showduplicates | sort -r

#也可以指定版本安装
yum install -y docker-ce-18.09.9

3.#配置源加速
mkdir -p /etc/docker
vim /etc/docker/daemom.json
{
  "registry-mirrors": [
    "https://8xpk5wnt.mirror.aliyuncs.com"
  ]
}

4.# 设置开机启动
systemctl enable docker
systemctl daemon-reload
systemctl start docker

5.# docker详细信息
docker info

6.#docker是一个cs架构
# c端：docker-client
~]# which docker
/usr/bin/docker

# s端：docker daemon
 ~]# ps aux | grep docker 
root       6796  0.0  0.8 749652 73012 ?        Ssl  09:47   0:00 /usr/bin/dockerd -H fd:// --containerd=/run/containerd/containerd.sock

# 容器container
 ~]# ps aux | grep containerd
root       6784  0.6  0.5 667652 45708 ?        Ssl  09:47   0:04 /usr/bin/containerd
root       6796  0.0  0.9 749652 75396 ?        Ssl  09:47   0:00 /usr/bin/dockerd -H fd:// --containerd=/run/container/containerd.sock

7.# docker 命令帮助信息
docker --help
docker run --help

```



小结

```
修改尽量做到备份
配置宿主机网卡转发
同步时间，注意防火墙等
```



#### 基本操作

1.查看镜像列表

```bash
docker images
```

2.获取镜像

```bash
#从远处仓库拉取
docker pull nginx:alpine

#本地构建
docker build . -t my-nginx:ubuntu  -f Dockerfile
#参数解释：
docker build 用于指定dockerfile创建镜像
-t, --tag 镜像的名字及标签，通常 name:tag 或者 name 格式；可以在一次构建中为一个镜像设置多个标签
-f 指定要使用的dockerfile路径

```

3.通过镜像启动容器

```bash
docker run --name my-nginx-alpine -d nginx:alpine

# 参数解释:
docker run 创建一个新的容器并运行一个命令
--name 为容器指定一个名称
-d 后台运行容器，并返回容器id

```

4.如何知道容器内部运行了什么程序

```bash
# 进入容器内部
$ docker exec -it my-nginx-alpine /bin/sh

# 容器内部执行
/ # ps aux
PID   USER     TIME  COMMAND
    1 root      0:00 nginx: master process nginx -g daemon off;
   33 nginx     0:00 nginx: worker process
   34 nginx     0:00 nginx: worker process
   35 nginx     0:00 nginx: worker process
   36 nginx     0:00 nginx: worker process
   43 root      0:00 sh
   49 root      0:00 ps aux

# 参数解释：
docker exec 在运行的容器中执行命令
-t 分配一个伪终端
-i 即使没有附加也保存STDIN(标准输入)打开
-d 分离模式，在后台运行

```

5.docker如何指定容器启动后该执行什么命令

通过docker build 来模拟构建一个Nginx镜像

```bash
1.创建dockerfile
mkdir demo
cd demo
vim Dockerfile

# 告诉docker使用哪个基础镜像作为模板,后续命令都以这个镜像为基础
FROM ubuntu

# RUN命令会在上面指定的镜像里面执行命令
RUN apt-get update && apt install -y nginx

# 告诉docker，启动容器时执行如下命令
CMD ["/usr/sbin/nginx", "-g", "daemon off;"]

# 参数解释：
# -g 从配置文件中设置全局指令
# daemon off 后台运行
# docker容器中pid为1的进程结束, 容器也就停止运行，所以要加daemon off

2.构建本地镜像
docker build . -t my-nginx:ubuntu -f Dockerfile

3.使用新镜像启动容器
docker run --name my-nginx-ubuntu -d my-nginx:ubuntu

4.进入容器查看进程
docker exec -ti my-nginx-ubuntu /bin/sh
# ps aux
# apt install -y  curl
# curl localhost:80

```

6.如何访问容器内的服务

```bash
docker exec -ti my-nginx-alpine  curl localhost:80

或进入容器内部访问
docker exec -ti my-nginx-alpine /bin/sh
# ps aux | grep nginx
# curl localhost:80

```

7.宿主机中如何访问容器服务

```bash
## 删除旧服务，重新做做端口映射启动
docker rm -f my-nginx-alpine
docker run --name my-nginx-alpine -d -p 8080:80 nginx:alpine
curl 192.168.178.79:8080

#参数解释：
-p 指定端口映射，格式为：主机端口:容器端口
-P 随机端口映射，容器内部端口随机映射到主机的端口，格式同上
rm 删除容器，-f 删除运行中的容器

```

#### 操作演示

![在这里插入图片描述](https://img-blog.csdnimg.cn/20210601165452750.jpg)

1.查看所有镜像

```bash
docker images

```

2.拉取镜像

```bash
docker pull name:tang

例：docker pull nginx:alpine

```

3.如何唯一确定镜像

- image_id
- repository:tag

```bash
例：
 ~]# docker images
REPOSITORY   TAG       IMAGE ID       CREATED       SIZE
my-nginx     ubuntu    a626628023ce   2 days ago    160MB
nginx        alpine    a6eb2a334a9f   5 days ago    22.6MB

```

4.导出镜像到文件中

```bash
docker save -o nginx-alpine.tar  nginx:alpine
docker save 将指定镜像保存成tar归档文件
-o 输出到文件
nginx:alpine 表示要指定的镜像

```

5.从文件中加载镜像

```bash
docker load -i nginx-alpine.tar
docker load 导入使用docker save命令导出的镜像
-i 指定导入的文件，代替STDIN

```

6.部署镜像仓库

https://docs.docker.com/registry/

```bash
# 使用docker镜像启动镜像仓库服务
docker run -d -p 5000:5000 --restart always --name registry registry:2
--restart always docker在重启的时候，只要启动docker，就自动启动这个容器

## 默认仓库不带认证，若需要认证参考：
https://docs.docker.com/registry/deploying/#restricting-access

```

7.推送本地镜像到镜像仓库中

```bash
docker tag nginx:alpine localhost:5000/nginx:alpine
docker push localhost:5000/nginx:alpine 

## 镜像仓库给外部访问，不能通过localhost，可以使用内网地址
 ~]# docker tag nginx:alpine 192.168.178.79:5000/nginx:alpine
 ~]# docker push 192.168.178.79:5000/nginx:alpine
The push refers to repository [192.168.178.79:5000/nginx]
Get https://192.168.178.79:5000/v2/: http: server gave HTTP response to HTTPS client

#docker默认不允许向http的仓库地址推送，如何做成https的，可以参考如下：
https://docs.docker.com/registry/deploying/#run-an-externally-accessible-registry

# 我们没有可信证书机构颁发的证书或域名，自签名证书需要再每个节点中拷贝证书文件，比较麻烦，因此我们通过配置daemon的方式，来跳过证书的验证:
 ~]# vim /etc/docker/daemon.json 
 ~]# cat /etc/docker/daemon.json
{
  "registry-mirrors" : [
    "https://8xpk5wnt.mirror.aliyuncs.com"
  ],
  "insecure-registries": [
     "192.168.178.79:5000"
  ]
}
 ~]# systemctl restart docker
 ~]# docker push 192.168.178.79:5000/nginx:alpine
 ~]# docker images
REPOSITORY                  TAG       IMAGE ID       CREATED       SIZE
192.168.178.79:5000/nginx   alpine    a6eb2a334a9f   5 days ago    22.6MB
nginx                       alpine    a6eb2a334a9f   5 days ago    22.6MB
ubuntu                      latest    7e0aa2d69a15   5 weeks ago   72.7MB
registry                    2         1fd8e1b0bb7e   6 weeks ago   26.2MB

```

##### 8.删除镜像

```bash
docker rmi [REPOSITORY]  
docker rmi  [IMAGE ID] 

```

##### 9.查看容器列表

```bash
## 查看运行状态的容器列表
docker ps 

## 查看全部状态的容器列表
docker ps -a

-a 显示所有的容器，包括未运行的
-f 根据条件过滤显示的内容，docker ps -f name='关键字搜索'
-l 显示最近创建的容器
-n 列出最近创建的n个容器
-q 静默模式，只显示容器编号
-s 显示总的文件大小

```

##### 10.启动容器

```bash
## 启动
docker run --name nginx -d nginx:alpine

## 映射端口，把容器的端口映射到宿主机中，-p 格式：宿主机端口:容器端口
docker run --name my-nginx -d -p 8080:80 nginx:alpine

## 资源限制，最大可用内存500M
 docker run --name lb-nginx -d --memory=500m  nginx:alpine

## 更多参考 docker run --help

```

#### 11.容器数据持久化

挂载

```bash
# 挂载主机目录
docker run --name nginx -d -v /opt:/opt nginx:alpine
docker run --name mysql -e MYSQL_ROOT_PASSWORD=123456 -d -v /opt/mysql/:/var/lib/mysql mysql:5.7

#参数解释：
-v 绑定一个卷，给容器挂载存储卷，挂载到容器的某个目录
-e 设置环境变量
# 宿主机目录是/opt/mysql，容器目录是/var/lib/mysql ，默认没有会自动创建

验证数据：
 ~]# ll /opt/mysql/
total 188484
-rw-r----- 1 polkitd input       56 May 31 16:03 auto.cnf
-rw------- 1 polkitd input     1680 May 31 16:03 ca-key.pem
...
```

**使用volumes卷**

```bash
# 查看volumes
docker volume  ls

# 创建
docker volume create my-vol

# 删除
docker volume rm my-vol

#创建volume并运行容器
docker run --name test-1 -d -v my-vol:/opt/my-vol nginx:alpine

# 查看容器是否运行
~]# docker ps -a -f name=test
CONTAINER ID   IMAGE          COMMAND                  CREATED              STATUS              PORTS     NAMES
1d178f091f89   nginx:alpine   "/docker-entrypoint.…"   About a minute ago   Up About a minute   80/tcp    test-1

# 不进入容器，但在容器内创建数据
docker exec -ti 1d178f091f89 touch /opt/my-vol/a.py


# 验证数据共享
# 再创建一个容器，查看数据是否存在
 ~]# docker run --name test-2 -d -v my-vol:/opt/hh nginx:alpine
c729a96f6f582b8006b34147a50667826705c0f0c637d432683515d9f2c03b38
~]# docker exec -ti test-2 ls /opt/hh
a.py

```

##### 12.进入容器或者执行容器内的命令

```bash
docker exec -ti <container_id_or_name> /bin/sh
docker exec <container_id_name> hostname

```

#### 13.主机与容器直接拷贝数据

```bash
## 主机拷贝到容器
 ~]# echo "love python" >> /tmp/test.py
 ~]# docker cp /tmp/test.py test-1:/tmp
 ~]# docker exec -ti test-1 cat /tmp/test.py
love python

## 容器拷贝到主机
 ~]# docker cp test-1:/tmp ./
 ~]# ls /tmp/
test.py  vmware-root

```

#### 14.查看容器日志

```bash
## 查看全部日志
docker logs container_name

## 实时查看最新日志
docker logs -f nginx

## 从最新的100条开始查看
docker logs --tail=100 -f nginx

```



##### 15.停止或删除容器

```bash
## 停止运行中的容器
docker stop nginx

## 启动退出容器
dockerstart nginx

## 删除非运行状态的容器
docker rm nginx

## 删除运行中的容器
docker rm -f nginx

## 过滤出指定容器，然后删除
docker rm -f `docker ps -f name=nginx -q`
```

#### 16.查看容器或者镜像的明细

```bash
# 查看容器详细信息，包括容器IP地址等
docker inspect nginx | more

#查看镜像的明细信息
docker inspect nginx:alpine

# 显示 Docker 系统信息，包括镜像和容器数
docker info
```

##### 17.挂载已有的数据，重新创建镜像仓库容器

```bash
1.导入本地进行文件
 ~]# tar -xzf  registry.tar.gz -C /opt/
 ~]# ls /opt/registry/
docker

2.删除当前镜像仓库文件
 ~]# docker rm -f registry 
registry

3.使用docker镜像启动镜像仓库服务
docker run -d -p 5000:5000 --restart always -v /opt/registry:/var/lib/registry --name registry registry:2

4.检查启动情况
 ~]# docker ps -a | grep registry
3e8ff4e819a0   registry:2     "/entrypoint.sh /etc…"   26 seconds ago      Up 25 seconds      0.0.0.0:5000->5000/tcp, :::5000->5000/tcp  registry

5.拉取镜像成功，则表示挂载成功
 ~]# docker pull 192.168.178.79:5000/centos:centos7.5.1804
 ~]# docker images | grep centos
192.168.178.79:5000/centos   centos7.5.1804   cf49811e3cdb   2 years ago   200MB

```

假设启动镜像仓库服务的主机地址为192.168.178.79，该目录中已存在的镜像列表：

| 现镜像仓库地址                                               | 原镜像仓库地址                                               |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| 192.168.178.79:5000/coreos/flannel:v0.11.0-amd64             | quay.io/coreos/flannel:v0.11.0-amd64                         |
| 192.168.178.79:5000/mysql:5.7                                | mysql:5.7                                                    |
| 192.168.178.79:5000/nginx:alpine                             | nginx:alpine                                                 |
| 192.168.178.79:5000/centos:centos7.5.1804                    | centos:centos7.5.1804                                        |
| 192.168.178.79:5000/elasticsearch/elasticsearch:7.4.2        | docker.elastic.co/elasticsearch/elasticsearch:7.4.2          |
| 192.168.178.79:5000/fluentd-es-root:v1.6.2-1.0               | quay.io/fluentd_elasticsearch/fluentd:v2.5.2                 |
| 192.168.178.79:5000/kibana/kibana:7.4.2                      | docker.elastic.co/kibana/kibana:7.4.2                        |
| 192.168.178.79:5000/kubernetesui/dashboard:v2.0.0-beta5      | kubernetesui/dashboard:v2.0.0-beta5                          |
| 192.168.178.79:5000/kubernetesui/metrics-scraper:v1.0.1      | kubernetesui/metrics-scraper:v1.0.1                          |
| 192.168.178.79:5000/kubernetes-ingress-controller/nginx-ingress-controller:0.30.0 | quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.30.0 |
| 192.168.178.79:5000/jenkinsci/blueocean:latest               | jenkinsci/blueocean:latest                                   |
| 192.168.178.79:5000/sonarqube:7.9-community                  | sonarqube:7.9-community                                      |
| 192.168.178.79:5000/postgres:11.4                            | postgres:11.4                                                |



#### 通过1号进程理解容器的本质

```bash
 ~]# docker run --name m-nginx-alpine -d nginx:alpine
81360b4251cf7384c709449ca1a6739f670098bc9cde0b78c4a92570a07689c1
 ~]# docker exec -ti m-nginx-alpine sh
/ # ps aux
PID   USER     TIME  COMMAND
    1 root      0:00 nginx: master process nginx -g daemon off;
   32 nginx     0:00 nginx: worker process
   33 nginx     0:00 nginx: worker process
   34 nginx     0:00 nginx: worker process
   35 nginx     0:00 nginx: worker process
   36 root      0:00 sh
   42 root      0:00 ps aux

```

容器启动的时候可以通过命令去覆盖默认的CMD

```bash
$ docker run -d --name xxx nginx:alpine <自定义命令
# <自定义命令> 会覆盖镜像中指定的CMD指令，作为容器的1号进程启动

# 命令执行完毕，1号进程退出即容器退出
 ~]# docker ps -a | grep test-4
a27336cd5b62   nginx:alpine   "/docker-entrypoint.…"   18 seconds ago      Exited (0) 17 seconds ago                                               test-4

# 可以看出只有1号进程存在，容器即存在
 ~]# docker run -d --name test-5  nginx:alpine ping www.baidu.com
376ce88c92a7a2ae2dd85149775874b811fc5fdd4e5f427b1383a3c9cb3929ce
 ~]# docker ps -a | grep test-5
376ce88c92a7   nginx:alpine   "/docker-entrypoint.…"   5 seconds ago        Up 5 seconds                    80/tcp                                      test-5

```

本质上讲容器是利用namespace和cgroup等技术在宿主机中创建独立的虚拟空间，这个空间内的网络，进程，挂载等资源都是隔离的。

```bash
 ~]# docker exec -ti m-nginx-alpine sh
/ # ifconfig | awk 'NR==2{print $2}'
addr:172.17.0.8
/ # ls /
bin                   etc                   mnt                   run                   tmp
dev                   home                  opt                   sbin                  usr
docker-entrypoint.d   lib                   proc                  srv                   var
docker-entrypoint.sh  media                 root                  sys

# 在容器内安装软件，创建修改文件等，对宿主机和其他容器没有任何影响，和虚拟机不同的是，容器间共享一个内核，所以容器内没法升级内核，且容器会随着1号进程的消亡而退出容器。

```

#### 容器的七种状态

```bash
create 已创建
restarting 重启中
running 运行中
removing 迁移中
paused 暂停
exited 停止
dead 死亡
```

#### 其他常用知识点

可以使用关键字进行搜索官方仓库中的镜像，并且可以使用docker pull 命令来将它下载到本地

```bash
docker search ubuntu
```

熟悉常用命令即可，其他命令可以使用–help 获取帮助

```bash
docker --help
docker run --help
```

标记本地镜像，将其归入某一仓库

```bash
docker tag [OPTIONS] IMAGE[:TAG] [REGISTRYHOST/][USERNAME/]NAME[:TAG]

例如：
docker tag nginx:alpine 192.168.178.79:5000/nginx:alpine
```

杀掉一个运行中的容器

```bash
docker kill -s KILL nginx
```

#### 小结

```
1.学习了容器的安装
2.容器的操作注意围绕三大核心要素，即镜像，容器，仓库
3.学会容器的基本操作，使用日志排查故障
4.通过1号进程理解容器的本质

小技巧：
可以使用 docker ps -aq 获取容器编号，批量删除容器，还可以通过-f过滤
```

##### 需要掌握的内容

```
熟练掌握常用命令，不知道的可以通过--help查询帮助，
docker run ： -v -e -d --name 等参数
docker create 
docker ps -a -q -f 等参数
docker logs --tail -f 参数
docker images
docker pull | push
docker cp 
docker rm -f | rmi -f

数据的持久化

学会使用官方网站查看：https://docs.docker.com/desktop/
```

容器实现原理
容器实现原理
虚拟化核心需要解决的问题：资源隔离与资源限制

虚拟机硬件虚拟化技术，通过一个hypervisor【虚拟】层实现对资源的彻底隔离
容器则是操作系统级别的虚拟化，利用的是内核的Cgroup和Namespace特性，此功能完全通过软件实现。
Namespace资源隔离
命名空间是全局资源的一种抽象，将资源放到不同的命名空间中，各个命名空间中的资源是相互隔离的。

| **分类**           | **系统调用参数** | **相关内核版本**                                             |
| ------------------ | ---------------- | ------------------------------------------------------------ |
| Mount namespaces   | CLONE_NEWNS      | [Linux 2.4.19](http://lwn.net/2001/0301/a/namespaces.php3)   |
| UTS namespaces     | CLONE_NEWUTS     | [Linux 2.6.19](http://lwn.net/Articles/179345/)              |
| IPC namespaces     | CLONE_NEWIPC     | [Linux 2.6.19](http://lwn.net/Articles/187274/)              |
| PID namespaces     | CLONE_NEWPID     | [Linux 2.6.24](http://lwn.net/Articles/259217/)              |
| Network namespaces | CLONE_NEWNET     | [始于Linux 2.6.24 完成于 Linux 2.6.29](http://lwn.net/Articles/219794/) |
| User namespaces    | CLONE_NEWUSER    | [始于 Linux 2.6.23 完成于 Linux 3.8](http://lwn.net/Articles/528078/) |



docker在启动一个容器的时候，会调用Linux Kemel Namespace的接口，来创建一块虚拟空间，创建的时候，可以支持设置下面这几种(可以随意选择)。docker默认都设置。

pid：用于进程隔离(PID：进程ID)
net：管理网络接口(NET：网络)
ipc：管理对IPC资源的访问(IPC：进程间通信，(信号量，消息队列和共享内存))
mnt：管理文件系统挂载点(MNT：挂载)
uts：隔离主机名和域名
user：隔离用户和用户组
CGroup资源限制
通过namespace可以保证容器之间的隔离，但是无法控制每个容器可以占有多少资源，如果其中的某一个容器正在执行CPU密集型的任务，那么就会影响其他容器中任务的性能和执行效率，导致多个容器相互影响并抢占资源。|

![在这里插入图片描述](https://img-blog.csdnimg.cn/20210601165618651.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L20wXzQ5NjU0MjI4,size_16,color_FFFFFF,t_70)



Control Groups（简称CGroups）就是能够隔离宿主机上的物理资源，例如CPU，内存磁盘I/O和网络带宽。每一个CGroup都是一组被相同的标准和参数限制的进程。而我们需要做的就是把容器这个进程加入到指定的CGroup中，来实现对多个容器资源使用进行限制。

UnionFS联合文件系统
Linux namespace和cgroup分别解决了容器的资源隔离与资源限制，那么容器是很轻量的，通常每台机器中可以运行几十上百个容器，这些容器是共用一个image，还是各自将这个image复制了一份，然后各自独立运行呢？如果每个容器之间都是全量的文件系统拷贝，那么会导致至少如下问题：

运行容器的速度会变慢
容器和镜像对宿主机的磁盘空间的压力
怎么解决这个问题—Docker的存储驱动

镜像分层存储
UnionFS
Docker 镜像是由一系列的层组成的，每层代表Docker中的一条指令，比如下面的Dockerfile文件：

```bash
FROM ubuntu:15.04
COPY ./app
RUN make /app
CMD python /app/app.py

```

这里的 Dockerfile 包含4条命令，其中每一行就创建了一层，下面显示了上述Dockerfile构建出来的镜像运行的容器层的结构：

![在这里插入图片描述](https://img-blog.csdnimg.cn/20210601165645261.jpg?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L20wXzQ5NjU0MjI4,size_16,color_FFFFFF,t_70)

镜像就是由这些一层一层堆叠起来的，镜像中的这些层都是只读的，当我们运行容器的时候，就可以在这些基础层上添加新的可写层，也就是我们通常说的容器层，对于运行中的容器所做的所有更改(比如写入新文件，修改现有文件，删除文件)都将写入到这个容器层。

对容器层的操作，主要利用了写时复制（CoW）技术。CoW就是copy-on-write，表示只在需要写时才去复制，这个是针对已有文件的修改场景。 CoW技术可以让所有的容器共享image的文件系统，所有数据都从image中读取，只有当要对文件进行写操作时，才从image里把要写的文件复制到自己的文件系统进行修改。所以无论有多少个容器共享同一个image，所做的写操作都是对从image中复制到自己的文件系统中的复本上进行，并不会修改image的源文件，且多个容器操作同一个文件，会在每个容器的文件系统里生成一个复本，每个容器修改的都是自己的复本，相互隔离，相互不影响。使用CoW可以有效的提高磁盘的利用率。



![在这里插入图片描述](https://img-blog.csdnimg.cn/20210601165715375.jpg?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L20wXzQ5NjU0MjI4,size_16,color_FFFFFF,t_70)

镜像中每一层的文件都是分散在不同的目录中的，如何把这些不同目录的文件整合到一起呢？

UnionFS 其实是一种为 Linux 操作系统设计的用于把多个文件系统联合到同一个挂载点的文件系统服务。 它能够将不同文件夹中的层联合（Union）到了同一个文件夹中，整个联合的过程被称为联合挂载（Union Mount）



![在这里插入图片描述](https://img-blog.csdnimg.cn/20210601165732532.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L20wXzQ5NjU0MjI4,size_16,color_FFFFFF,t_70)



上图是AUFS的实现，AUFS是作为Docker存储驱动的一种实现，Docker 还支持了不同的存储驱动，包括 aufs、devicemapper、overlay2、zfs 和 Btrfs 等等，在最新的 Docker 中，overlay2 取代了 aufs 成为了推荐的存储驱动，但是在没有 overlay2 驱动的机器上仍然会使用 aufs 作为 Docker 的默认驱动。

小结

```bash
容器的实现依赖于内核模块提供的namespace和control-group的功能，通过namespace创建一块虚拟空间，空间内实现了各类资源(进程，网络，文件系统)的隔离，提供control-group实现看对隔离空间的资源的使用的限制。

对容器层的操作，主要利用了写时复制（CoW）技术。

UnionFS 是一种为 Linux 操作系统设计的用于把多个文件系统联合到同一个挂载点的文件系统服务。  它能够将不同文件夹中的层联合（Union）到了同一个文件夹中，整个联合的过程被称为联合挂载（Union Mount）。

```



# docker网络原理

Docker使用Linux桥接，在宿主机虚拟一个Docker容器网桥(docker0)，Docker启动一个容器时会根据Docker网桥的网段分配给容器一个IP地址，称为`Container-IP`，同时**Docker网桥是每个容器的默认网关**。因为在同一宿主机内的容器都接入同一个网桥，这样容器之间就能够通过容器的Container-IP直接通信。

Docker网桥是宿主机虚拟出来的，并不是真实存在的网络设备，外部网络是无法寻址到的，这也意味着**外部网络无法通过直接Container-IP访问到容器**。如果容器希望外部访问能够访问到，可以通过**映射容器端口到宿主主机（端口映射）**，即docker run创建容器时候通过 -p 或 -P 参数来启用，访问容器的时候就通过`[宿主机IP]:[容器端口]`访问容器。



#### docker网络模式

我们在使用docker run创建docker容器时，可以用-net选型指定容器的网络模式，docker有以下四种网络模式：

| Docker网络模式     | 配置                       | 说明                                                         |
| ------------------ | -------------------------- | ------------------------------------------------------------ |
| host模式           | --net=host                 | 容器和宿主机共享Network namespace。                          |
| container模式      | --net=container:NAME_or_ID | 容器和另外一个容器共享Network namespace。 kubernetes中的pod就是多个容器共享一个Network namespace。 |
| none模式           | --net=none                 | 容器有独立的Network namespace，但并没有对其进行任何网络设置，如分配veth pair 和网桥连接，配置IP等。 |
| bridge模式（默认） | --net=bridge               |                                                              |

### 1）bridge模式(NAT)

当Docker进程启动时，会在主机上创建一个名为docker0的虚拟网桥，此主机上启动的Docker容器会连接到这个虚拟网桥上。虚拟网桥的工作方式和物理交换机类似，这样主机上的所有容器就通过交换机连在了一个二层网络中。

从`docker0子网`中分配一个IP给容器使用，并设置**docker0的IP地址**为容器的**默认网关**。在主机上创建一对虚拟网卡veth pair设备（可以理解为网线），Docker将veth pair设备的一端放在新创建的容器中，并命名为`eth0`（容器的网卡）；另一端放在主机中，以`vethxxx`这样类似的名字命名，并将这个网络设备加入到docker0网桥中。可以通过`brctl show`命令查看。

bridge模式是docker的默认网络模式，不写--net参数，就是bridge模式。使用docker run -p时，docker实际是在iptables做了DNAT规则，实现端口转发功能。可以使用`iptables -t nat -vnL`查看。

如果不指定的话，默认就会使用bridge模式，bridge本意是桥的意思，其实就是网桥模式。我们怎么理解网桥，如果需要做类比的话，我们可以把网桥看出一个二层的交换机设备。

交换机通信简图

![在这里插入图片描述](https://img-blog.csdnimg.cn/20210601165751738.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L20wXzQ5NjU0MjI4,size_16,color_FFFFFF,t_70)

> 也叫NAT
>
> ```bash
> [root@node4 ~]# docker run -ti --rm sunrisenan/alpine:3.10.3 /bin/sh
> / # ip a
> 1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN qlen 1
>     link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
>     inet 127.0.0.1/8 scope host lo
>        valid_lft forever preferred_lft forever
> 100: eth0@if101: <BROADCAST,MULTICAST,UP,LOWER_UP,M-DOWN> mtu 1500 qdisc noqueue state UP # docker里为eth0
>     link/ether 02:42:ac:06:f4:02 brd ff:ff:ff:ff:ff:ff
>     inet 172.6.244.2/24 brd 172.6.244.255 scope global eth0 # IP为172.
>        valid_lft forever preferred_lft forever
> ```



#### **网桥模式示意图**

![在这里插入图片描述](https://img-blog.csdnimg.cn/2021060116580583.png)![](https://i0.hdslb.com/bfs/album/75b9ebe95f54babd161d41dda322333b3332de06.png)

Linux能够起到**虚拟交换机作用**的网络设备，是网桥。他是一个工作在**数据链路层**(data link)的设备，主要的功能是**根据MAC地址将数据包转发到网桥的不同端口上**。

查看网桥

```bash
#下载安装
yum install -y  bridge-utils

[root@localhost ~]# brctl show
bridge name     bridge id               STP enabled     interfaces
docker0         8000.0242cf195186       no              veth10904eb
                                                        veth10ff0a9
                                                        veth1787878
                                                        veth8244792
                                                        veth9760247
                                                        vethbca21e7
```

**关于brctl命令参数说明和示例**

| 参数            | 说明                   | 示例                  |
| --------------- | ---------------------- | --------------------- |
| `addbr `        | 创建网桥               | brctl addbr br10      |
| `delbr `        | 删除网桥               | brctl delbr br10      |
| `addif  `       | 将网卡接口接入网桥     | brctl addif br10 eth0 |
| `delif  `       | 删除网桥接入的网卡接口 | brctl delif br10 eth0 |
| `show `         | 查询网桥信息           | brctl show br10       |
| `stp  {on|off}` | 启用禁用 STP           | brctl stp br10 off/on |
| `showstp `      | 查看网桥 STP 信息      | brctl showstp br10    |
| `setfd  `       | 设置网桥延迟           | brctl setfd br10 10   |
| `showmacs `     | 查看 mac 信息          | brctl showmacs br10   |

##### docker网络

> docker在启动一个容器时时如何实现容器间的互联互通的？

Docker创建一个容器的时候，会执行如下操作：

- 创建一对虚拟接口/网卡，也就是一对虚拟接口（veth pair)
- 本地主机一端桥接到默认的docker0或指定网桥上，并具有一个唯一的名字，如`veth9953b75`
- 容器一端放到新启动的容器内部，并修改名字为`eth0`，这个网卡/接口只在容器的命名空间可见；
- 从网桥可用地址段中(也就是与该bridge对应的network)获取一个空闲地址分配给容器的eth0
- 配置默认路由到网桥

整个过程其实都是docker自动帮我们完成的，清理掉是所有的容器来验证：

> 基础知识
>
> 主机们在不在一个广播域，完全取决于**主机连接的交换机端口们在不在同一个VLAN**：
>
> **1. 如果在同一个VLAN，即使主机们的网段不相同，也是工作在一个广播域。**
>    1.1 主机们的网段相同，可以ARP发现彼此的MAC，直接通信，不需要任何三层设备（网关）的介入。
>
>    1.2 主机们的网段不相同，即使在一个广播域，也不能直接通信，需要三层设备（网关）的介入。
>
> **2. 如果不在一个VLAN，主机们不在一个广播域**
>    2.1 一个VLAN对应一个网段，那么主机之间的通信需要三层设备（网关）的介入。
>
>    2.2 如果很不巧，两个VLAN里的主机使用相同的网段，主机并不知道有VLAN 的存在，所以依然认为其它主机和自己在一个广播域，但是这个广播域被交换机VLAN逻辑隔离，成为两个广播域，这时无法使用三层设备使得它们通信，唯一的方法，使用一个网桥将两个VLAN二层桥接起来，它们就可以通信了。
>
> **所谓网关，就是上文提到的三层设备，可以是路由器、或三层交换机、防火墙。**
>
> 哈哈，这个2.2的情况估计没有几个看得懂，凡是没点赞的都是看不懂的，整个知乎用户能看懂2.2情况的不会超过1000人…

> 网关是负责不同网络通信使用的。不同网络指的是 ip 地址中的网络号不同，比如 192.168.2.3/24，这个 ip 表示网络号为192.168.2(前24位)
>
> 比如 a 节点 ip 为`192.168.2.1/24`， b 节点 ip 为 `192.168.2.3/24` ，a 给 b 发送消息的时候，会先看是否在同一个网络，根据 b 的 ip 很容易判断出都是在 192. 168. 2这个网络内，所以 a 会直接给 b 发送消息(通过 b 的 mac 地址，这个 mac 地址是通过 `arp `协议获取的） 。
>
> c节点 ip 地址为 `192.168.3.2/24`，a 发送消息给 c， a 很容易知道 c 的网络地址是 192.168.3，与自己的网络地址不一样，这时候就会把这个消息发送给网关(也是通过 mac 地址)，网关负责把消息发送给 c。
>
> 说白了，就是通信协议规定了，在同网段内可以直接通过 mac 通信，不同网段需要通过网关通信。

```bash
# 清理所有容器
docker rm -f `docker ps -aq`
docker ps -a

# 查看网桥中的接口，目前是没有的
[root@docker ~]# brctl show
bridge name	bridge id		STP enabled	interfaces
docker0		8000.0242cfb7aaca	no	    空

# 创建测试容器test1
docker run -d --name test1 nginx:alpine

# 查看网桥中的接口，已经把test1 的veth端接入到网桥中
[root@docker ~]# brctl show
bridge name	bridge id		STP enabled	interfaces
docker0		8000.0242cfb7aaca	no		veth95f0f29

#已经在宿主机中可以查看到
[root@docker ~]# ip a | grep veth
131: veth95f0f29@if130: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue master docker0 state UP group default 

#进入容器查看容器的eth0网卡及分配的容器IP  # docker exec test1 ip a
[root@docker ~]# docker exec -ti test1 sh
/ # ifconfig | awk 'NR==2{print $2}'
addr:172.17.0.2
/ # route -n
Kernel IP routing table
Destination     Gateway         Genmask         Flags Metric Ref    Use Iface
0.0.0.0         172.17.0.1      0.0.0.0         UG    0      0        0 eth0
172.17.0.0      0.0.0.0         255.255.0.0     U     0      0        0 eth0



# 再启动一个测试容器，测试容器间的通信
docker run -d --name test2 nginx:alpine
docker exec -it test2 sh
/# curl 172.17.0.2:80

## 为啥可以通信
/ # route -n
Kernel IP routing table
Destination     Gateway         Genmask         Flags Metric Ref    Use Iface
0.0.0.0         172.17.0.1      0.0.0.0         UG    0      0        0 eth0
172.17.0.0      0.0.0.0         255.255.0.0     U     0      0        0 eth0
# eth0 网卡是这个容器里的默认路由设备，所有对172.17.0.0/16网段的请求，也会被交给eth0来处理(第二条 172.17.0.0 路由规则)，这条路由规则的网关(Gateway)是0.0.0.0，这就意味着这是一条直连规则，即：凡是匹配到这条规则的IP包，应该经过本机的eth0网卡，通过二层网络(数据链路层)直接发往目的主机。

# 网桥会维护一份Mac映射表，我们可以通过命令来查看一下
[root@docker ~]# brctl showmacs docker0
port no	mac addr		is local?	ageing timer
  1	32:76:d0:a3:0d:5c	yes		   0.00
  1	32:76:d0:a3:0d:5c	yes		   0.00
  2	86:4e:1c:5d:07:83	yes		   0.00
  2	86:4e:1c:5d:07:83	yes		   0.00
  
 # 这些Mac地址是主机端的veth网卡对于的Mac，可以查看运行
ip a |grep -n3 eth
```

![](https://i0.hdslb.com/bfs/album/f0471c86c6d749bd74e9fce0877affb2b9391bf3.png)

<img src="https://img-blog.csdnimg.cn/20210601165835644.png" alt="在这里插入图片描述" style="zoom:50%;" />



**我们如何指定网桥上的这些虚拟网卡与容器端是如何对应？**

通过ifindex，网卡索引号

```bash
# 分别查看test1，test2 容器的网卡索引
[root@docker ~]# docker exec -ti test1 cat /sys/class/net/eth0/ifindex
130
[root@docker ~]# docker exec -ti test2 cat /sys/class/net/eth0/ifindex
134

#再通过在主机中找到虚拟网卡后面这个@ifxx的值，如果是同一个值，说明这个虚拟网卡和这个容器的eth0是配对的
[root@docker ~]#  ip a | grep @if
131: veth95f0f29@if130: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue master docker0 state UP group default 
135: veth25234c9@if134: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue master docker0 state UP group default 

```

#### 容器与宿主机的通信

添加端口映射：

```bash
# 启动容器的时候通过-p 参数添加宿主机端口与容器内部服务端口的映射
docker run --name test -d -p 8080:80 nginx:alpine
curl localhost:8080
```

![在这里插入图片描述](https://img-blog.csdnimg.cn/20210601170222380.jpeg?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L20wXzQ5NjU0MjI4,size_16,color_FFFFFF,t_70)

**端口映射通过iptables如何实现**

```bash

```

![在这里插入图片描述](https://img-blog.csdnimg.cn/2021060117023866.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L20wXzQ5NjU0MjI4,size_16,color_FFFFFF,t_70)



访问本机的8088端口，数据包会从流入方向进入本机，因此涉及到PREROUTING【路由前】和INPUT链，我们是通过做宿主机与容器之间加的端口映射，所以肯定会涉及到端口转换，那哪个表是负责存储端口转换信息的呢，那就是nat表，负责维护网络地址转换信息的。

```bash
# 查看一下PREROUTING链的nat表
[root@docker ~]# iptables -t nat -nvL PREROUTING
Chain PREROUTING (policy ACCEPT 0 packets, 0 bytes)
 pkts bytes target     prot opt in     out     source               destination         
   26  1488 DOCKER     all  --  *      *       0.0.0.0/0            0.0.0.0/0            ADDRTYPE match dst-type LOCAL

参数解释：
-t 对指定的表进行操作
-n 以数字的方式显示ip，它会将ip直接显示出来，如果不加-n，则会将ip反向解析成主机名。
-v 详细模式；-vvv :越多越详细
-L 列出链chain上面的所有规则，如果没有指定链，列出表上所有规则
```

规则利用了iptables的addrtype【地址类型】扩展，匹配网络类型为本地的包

```bash
# 如何确定哪些是匹配本地的
[root@docker ~]# ip route show table local type local
local 127.0.0.0/8 dev lo proto kernel scope host src 127.0.0.1 
local 127.0.0.1 dev lo proto kernel scope host src 127.0.0.1 
local 172.17.0.1 dev docker0 proto kernel scope host src 172.17.0.1 
local 192.168.178.79 dev ens33 proto kernel scope host src 192.168.178.79 

```

也就是说目标地址类型匹配到这些的，会转发到我们的TARGET中，TARGET是动作，意味着对符合要求的数据包执行什么样的操作，最常见的为ACCEPT(接受)和DROP(终止)，此处的target(目标)为docker，很明显docker不是标准的动作，那docker是什么呢？我们通常会定义自定义的链，这样把某类对应的规则放在自定义链中，然后把自定义的链绑定到标准的链路中，因此此处docker是自定义的链。那我们现在就来看一下docker这个自定义链上的规则。

```bash
[root@docker ~]# iptables -t nat -nvL DOCKER
Chain DOCKER (2 references)
 pkts bytes target     prot opt in     out     source               destination         
   17  1020 RETURN     all  --  docker0 *       0.0.0.0/0            0.0.0.0/0           
    0     0 DNAT       tcp  --  !docker0 *       0.0.0.0/0            0.0.0.0/0            tcp dpt:8080 to:172.17.0.2:80
```

此条规则就是对主机收到的目的端口为8080的tcp流量进行DNAT转换，将流量发往172.17.0.2:80，172.17.0.2地址是不是就是我们上面创建的docker容器的ip地址，流量走到网桥上了，后面就走网桥的转发就ok了。所以，外加只需要访问192.168.178.79:8080就可以访问到容器中的服务了。



```bash
数据包在出来方向走postrouting链，查看一下规则
[root@docker ~]#  iptables -t nat -nvL POSTROUTING
Chain POSTROUTING (policy ACCEPT 0 packets, 0 bytes)
 pkts bytes target     prot opt in     out     source               destination         
   96  5925 MASQUERADE  all  --  *      !docker0  172.17.0.0/16        0.0.0.0/0           
    0     0 MASQUERADE  tcp  --  *      *       172.17.0.2           172.17.0.2           tcp dpt:80
    
MASQUERADE这个动作其实是一种更灵活的SNAT，把原地址转换成主机的出口ip地址，解释一下这条规则的意思：
这条规则会将原地址为172.17.0.0/16的包(也就是从docker容器产生的包)，并且不是从docker0网卡发出的，进行原地址转换，转发成主机网卡的地址。大概的过程就是ACK的包在容器里面发出来，会路由到网桥docker0，网桥根据宿主机的路由规则会转给宿主机网卡eth0，这时候包就从docker0网卡转到eth0网卡了，并从eth0网卡发出去，这时候这条规则就会生效了，把源地址换成了eth0的ip地址。

```

> 注意一下，刚才这个过程涉及到了网卡间包的传递，那一定要打开主机的ip_forward转发服务，要不然包转不了，服务肯定访问不到。

### 2）host模式

如果启动容器的时候使用host模式，那么这个容器将不会获得一个独立的Network Namespace，而是**和宿主机共用一个Network Namespace**。容器将不会虚拟出自己的网卡，配置自己的IP等，而是使用宿主机的IP和端口。但是，容器的其他方面，如文件系统、进程列表等还是和宿主机隔离的。

使用host模式的容器可以直接使用宿主机的IP地址与外界通信，容器内部的服务端口也可以使用宿主机的端口，不需要进行NAT，host最大的优势就是网络性能比较好，但是docker host上已经使用的端口就不能再用了，网络的隔离性不好。

![](https://i0.hdslb.com/bfs/album/1fd613ed1b4624a56be999a8d5a91bf6c37056d5.png)

容器内部不会创建网络空间，共享宿主机的网络空间，比如直接通过host模式创建mysql容器：

```bash
$ docker run --net host -d --name mysql -e MYSQL_ROOT_PASSWORD=123456 mysql:5.7
$ curl localhost:3306
5.7.3Uid4ÿÿ󿿕1[M/Z5NGRy}mysql_native_password!ÿ#08S01Got packets out of order

```

容器启动后，会默认监听3306端口，由于网络是host，因为可以直接通过宿主机的3306端口进行访问服务，效果等同于在宿主机直接启动mysqld的进程。

```bash
[root@node4 ~]# docker run -ti --rm --net=host sunrisenan/alpine:3.10.3 /bin/sh
/ # ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN qlen 1
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host 
       valid_lft forever preferred_lft forever
2: ens192: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc mq state UP qlen 1000
    link/ether 00:50:56:b3:4c:8f brd ff:ff:ff:ff:ff:ff
    inet 192.168.6.244/24 brd 192.168.6.255 scope global ens192
       valid_lft forever preferred_lft forever
    inet6 fe80::250:56ff:feb3:4c8f/64 scope link 
       valid_lft forever preferred_lft forever
3: docker0: <NO-CARRIER,BROADCAST,MULTICAST,UP> mtu 1500 qdisc noqueue state DOWN 
    link/ether 02:42:65:95:28:d8 brd ff:ff:ff:ff:ff:ff
    inet 172.6.244.1/24 brd 172.6.244.255 scope global docker0
       valid_lft forever preferred_lft forever
    inet6 fe80::42:65ff:fe95:28d8/64 scope link 
       valid_lft forever preferred_lft forever
```



### 3）container模式

这个模式指定新创建的容器**和已经存在的一个容器共享一个network namespace**，而不是和宿主机共享。新创建的容器不会创建自己的网卡，配置自己的IP，而是和一个指定的容器共享IP，端口范围等。同样，两个容器除了网络方面，其他的如文件系统，进程列表等还是隔离的，两个容器的进程可以通过IO网卡设备通信。

![在这里插入图片描述](https://img-blog.csdnimg.cn/20210601170303227.jpeg)![](https://i0.hdslb.com/bfs/album/2b39e5726282da9da21af58302eb17213a9147ea.png)

```bash
# 启动容器测试，共享mysql的网络空间
$ docker run --net host -d --name mysql -e MYSQL_ROOT_PASSWORD=123456 mysql:5.7
# 冒号后面是mysql，意思是新建一个容器，同时容器的网络指向mysql的网络空间
$ docker run -ti --rm --net=container:mysql busybox sh
/ # ps aux
PID   USER     TIME  COMMAND
    1 root      0:00 sh
   12 root      0:00 ps aux
/ # ifconfig | awk 'NR==2{print $2}'
addr:172.17.0.1
/ # netstat -tlp | grep 3306
tcp        0      0 :::3306                 :::*                    LISTEN      -
/ # telnet localhost 3306
Connected to localhost
J
5.7.34Z..bHeOÿ
             Nu/qXS
                   vl0mysql_native_password

!#08S01Got packets out of orderConnection closed by foreign host

## --rm 退出容器后，容器会自动删除

```

在一下特殊的场景中非常有用，例如kubernetes的pod，kubernetes为pod创建一个基础设施容器，同一pod下的其他容器都以container模式共享这个基础设施的网络命名空间，相互之间以localhost访问，构成一个统一的整体。

> 联合网络
>
> ```bash
> [root@node4 ~]# docker run -d sunrisenan/nginx:v1.12.2
> 76f844e6057b6493a4a8933819ade7c29325ef17bc1fddc88bdf4c7ec60c620b
> [root@node4 ~]# 
> [root@node4 ~]# docker ps -qa
> 76f844e6057b
> [root@node4 ~]# docker run -ti --rm --net=container:76f844e6057b sunrisenan/nginx:curl bash
> root@76f844e6057b:/# curl localhost
> <!DOCTYPE html>
> <html>
> <head>
> <title>Welcome to nginx!</title>
> <style>
>     body {
>         width: 35em;
>         margin: 0 auto;
>         font-family: Tahoma, Verdana, Arial, sans-serif;
>     }
> </style>
> </head>
> <body>
> <h1>Welcome to nginx!</h1>
> <p>If you see this page, the nginx web server is successfully installed and
> working. Further configuration is required.</p>
> 
> <p>For online documentation and support please refer to
> <a href="http://nginx.org/">nginx.org</a>.<br/>
> Commercial support is available at
> <a href="http://nginx.com/">nginx.com</a>.</p>
> 
> <p><em>Thank you for using nginx.</em></p>
> </body>
> </html>
> root@76f844e6057b:/# 
> ```
>
> 

### 4）none模式

使用none模式，**Docker容器拥有自己的Network Namespace，但是，并不为Docker容器进行任何网络配置**。也就是说，这个Docker容器没有网卡、IP、路由等信息。需要我们自己为Docker容器添加网卡、配置IP等。

这种网络模式下容器只有lo回环网络，没有其他网卡。none模式可以在容器创建时通过--network=none来指定。这种类型的网络没有办法联网，封闭的网络能很好的保证容器的安全性。

None模式示意图:

![](https://i0.hdslb.com/bfs/album/5f081aa8c9d10d87a518474031fd7e74e10c69a9.png)

```bash
[root@node4 ~]# docker run -ti --rm --net=none  sunrisenan/alpine:3.10.3 /bin/sh
/ # ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN qlen 1
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever

```



#### 小结

```bash
了解容器网络原理
docker网络的四种模式：bridge，host，container，none
掌握容器与容器间，容器与宿主机之间的通信原理

```

##### 使用技巧

1.清理主机上所有退出的容器

```bash
docker rm `docker ps -aq`
docker rm $(docker ps -aq)
```

2.调试或者排查容器启动错误

```bash
# 若有时遇到容器启动失败的情况，可以使用相同的镜像启动一个临时容器，先进入容器
docker run --rm -ti <image_id> sh

# 进入容器后，手动执行该容器对应的ENTRYPOINT或者CMD命令，这样即使出错，容器也不会退出，因为bash进程作为1号进程，我们只要不退出容器，该容器就不会自动退出。

```







