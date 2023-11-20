## 基于Docker和Kubernetes的企业级DevOps实践训练营

### 课程准备

1. CentOS7.4版本以上  虚拟机3台（4C+8G+50G），内网互通，可连外网

2. 离线镜像包

   链接：https://pan.baidu.com/s/1N1AYGCYftYGn6L0QPMWIMw  提取码：ev2h

3. 课件文档

   - 《训练营课件》
   - 《安装手册》

4. git仓库

   https://gitee.com/agagin/python-demo.git  python demo项目

   https://gitee.com/agagin/demo-resources.git  demo项目演示需要的资源文件

### 课程安排

2020.4.11 Docker + kubernetes

2020.4.18 DevOps平台实践

2天的时间，节奏会相对快一些



本着实践为核心的思想，本课程使用企业常见的基于Django + uwsgi + Nginx架构的Python Demo项目，分别讲述三个事情：

- 项目的容器化

  教大家如何把公司的项目做成容器，并且运行在docker环境中

- 使用Kubernetes集群来管理容器化的项目

  带大家一步一步部署k8s集群，并把容器化后的demo项目使用k8s来管理起来

- 使用Jenkins和Kubernetes集成，实现demo项目的持续集成/持续交付(CI/CD)

  会使用k8s管理应用生命周期后，还差最后的环节，就是如何把开发、测试、部署的流程使用自动化工具整合起来，最后一部分呢，课程会教会大家如何优雅的使用gitlab+Jenkins+k8s构建企业级的DevOps平台



### 流程示意

![](https://i0.hdslb.com/bfs/album/c34fd839fb5d50ef423cd53a80546fb2ff6c3d35.png)

### 你将学到哪些

- Docker相关
  - 如何使用Dockerfile快速构建镜像
  - Docker镜像、容器、仓库的常用操作
  - Docker容器的网络（Bridge下的SNAT、DNAT）

- Kubernetes相关
  - 集群的快速搭建
  - kubernetes的架构及工作流程
  - 使用Pod控制器管理业务应用的生命周期
  - 使用CoreDNS、Service和Ingress实现服务发现、负载均衡及四层、七层网络的访问
  - Kubernetes的认证授权体系
- 使用EFK构建集群业务应用的日志收集系统
  
- 基于Gitlab+Jenkins+k8s构建DevOps平台
  - Jenkins介绍及流水线的使用
  - Jenkinsfile及多分支流水线的实际应用
  - Jenkins集成sonarQube、Docker、Kubernetes
  - 使用groovy编写sharedLibrary，实现CI/CD流程的优化



### 第一章  走进Docker的世界

介绍docker的前世今生，了解docker的实现原理，以Django项目为例，带大家如何编写最佳的Dockerfile构建镜像。通过本章的学习，大家会知道docker的概念及基本操作，并学会构建自己的业务镜像，并通过抓包的方式掌握Docker最常用的bridge网络模式的通信。

#### 认识docker

###### 怎么出现的

- 轻量、高效的虚拟化

  Docker 公司位于旧金山,原名dotCloud，底层利用了Linux容器技术（在操作系统中实现资源隔离与限制）。为了方便创建和管理这些容器，dotCloud 开发了一套内部工具，之后被命名为“Docker”。Docker就是这样诞生的。

  （思考为啥要用Linux容器技术？）

  ![](https://i0.hdslb.com/bfs/album/646c8df117928a81f2f3d8cd697a288bf94ec1d2.jpg)

Hypervisor： 一种运行在基础物理服务器和操作系统之间的中间软件层，可允许多个操作系统和应用共享硬件 。常见的VMware的 Workstation 、ESXi、微软的Hyper-V或者思杰的XenServer。

Container Runtime：通过Linux内核虚拟化能力管理多个容器，多个容器共享一套操作系统内核。因此摘掉了内核占用的空间及运行所需要的耗时，使得容器极其轻量与快速。

- 软件交付过程中的环境依赖

  ![](https://i0.hdslb.com/bfs/album/8dc7b2302f9cd1e2c3be7683ab778d13cfffc00b.png)

###### 几个知识点

- 可以把应用程序代码及运行依赖环境打包成镜像，作为交付介质，在各环境部署

- 可以将镜像（image）启动成为容器(container)，并且提供多容器的生命周期进行管理（启、停、删）

- container容器之间相互隔离，且每个容器可以设置资源限额

- 提供轻量级虚拟化功能，容器就是在宿主机中的一个个的虚拟的空间，彼此相互隔离，完全独立

- CS架构的软件产品

  ![](https://i0.hdslb.com/bfs/album/7a74c72cbbe6334471bf7da1874cd0390ec25904.png)

###### 版本管理

- Docker 引擎主要有两个版本：企业版（EE）和社区版（CE）
- 每个季度(1-3,4-6,7-9,10-12)，企业版和社区版都会发布一个稳定版本(Stable)。社区版本会提供 4 个月的支持，而企业版本会提供 12 个月的支持
- 每个月社区版还会通过 Edge 方式发布月度版
- 从 2017 年第一季度开始，Docker 版本号遵循 YY.MM-xx 格式，类似于 Ubuntu 等项目。例如，2018 年 6 月第一次发布的社区版本为 18.06.0-ce

![](https://i0.hdslb.com/bfs/album/ea30a2b73b4d98382e100f29405b801d1ea7df8a.png)

###### 发展史

13年成立，15年开始，迎来了飞速发展。

Docker 1.8之前，使用[LXC](https://linuxcontainers.org/fr/lxc/introduction/)，Docker在上层做了封装， 把LXC复杂的容器创建与使用方式简化为自己的一套命令体系。 

之后，为了实现跨平台等复杂的场景，Docker抽出了libcontainer项目，把对namespace、cgroup的操作封装在libcontainer项目里，支持不同的平台类型。

2015年6月，Docker牵头成立了 OCI（Open Container Initiative开放容器计划）组织，这个组织的目的是建立起一个围绕容器的通用标准 。 容器格式标准是一种不受上层结构绑定的协议，即不限于某种特定操作系统、硬件、CPU架构、公有云等 ， 允许任何人在遵循该标准的情况下开发应用容器技术，这使得容器技术有了一个更广阔的发展空间。

OCI成立后，libcontainer 交给OCI组织来维护，但是libcontainer中只包含了与kernel交互的库，因此基于libcontainer项目，后面又加入了一个CLI工具，并且项目改名为runC (https://github.com/opencontainers/runc )， 目前runC已经成为一个功能强大的runtime工具。

Docker也做了架构调整。将容器运行时相关的程序从docker daemon剥离出来，形成了**containerd**。containerd向上为Docker Daemon提供了`gRPC接口`，使得Docker Daemon屏蔽下面的结构变化，确保原有接口向下兼容。向下通过`containerd-shim`结合`runC`，使得引擎可以独立升级，避免之前Docker Daemon升级会导致所有容器不可用的问题。



![](https://i0.hdslb.com/bfs/album/111871c43ac50ea0813071f2e11b9d5a7399de6f.png)



也就是说

- runC（libcontainer）是符合OCI标准的一个实现，与底层系统交互
- containerd是实现了OCI之上的容器的高级功能，比如镜像管理、容器执行的调用等
- Dockerd目前是最上层与CLI交互的进程，接收cli的请求并与containerd协作

###### 小结

1. 为了解决软件交付过程中的环境依赖，同时提供一种更加轻量的虚拟化技术，Docker出现了
2. Docker是一种CS架构的软件产品，可以把代码及依赖打包成镜像，作为交付介质，并且把镜像启动成为容器，提供容器生命周期的管理
3. docker-ce，每季度发布stable版本。18.06，18.09，19.03
4. 发展至今，docker已经通过制定OCI标准对最初的项目做了拆分，其中runC和containerd是docker的核心项目，理解docker整个请求的流程，对我们深入理解docker有很大的帮助



#### 安装

###### 配置宿主机网卡转发

```powershell
## 配置网卡转发,看值是否为1
$ sysctl -a |grep -w net.ipv4.ip_forward
net.ipv4.ip_forward = 1

## 若未配置，需要执行如下
$ cat <<EOF >  /etc/sysctl.d/docker.conf
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
net.ipv4.ip_forward=1
EOF
$ sysctl -p /etc/sysctl.d/docker.conf
```

###### Yum安装配置docker

```powershell
## 下载阿里源repo文件
$ curl -o /etc/yum.repos.d/Centos-7.repo http://mirrors.aliyun.com/repo/Centos-7.repo
$ curl -o /etc/yum.repos.d/docker-ce.repo http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo

$ yum clean all && yum makecache
## yum安装
$ yum install -y docker-ce
## 查看源中可用版本
$ yum list docker-ce --showduplicates | sort -r
## 安装指定版本
##yum install -y docker-ce-18.09.9

## 配置源加速
## https://cr.console.aliyun.com/cn-hangzhou/instances/mirrors
mkdir -p /etc/docker
vi /etc/docker/daemon.json
{
  "registry-mirrors" : [
    "https://8xpk5wnt.mirror.aliyuncs.com",
    "https://dockerhub.azk8s.cn",
    "https://registry.docker-cn.com",
    "https://ot2k4d59.mirror.aliyuncs.com/"
  ]
}

## 设置开机自启
systemctl enable docker  
systemctl daemon-reload

## 启动docker
systemctl start docker 

## 查看docker信息
docker info

## docker-client
which docker
## docker daemon
ps aux |grep docker
```



#### 核心要素及常用操作详解

![](https://i0.hdslb.com/bfs/album/22f4fcabcd939c68c94229830ce64394c2afd861.png)

三大核心要素：镜像(Image)、容器(Container)、仓库(Registry)

（先整体看下流程，再逐个演示）

###### 镜像（Image）

打包了业务代码及运行环境的包，是静态的文件，不能直接对外提供服务。

###### 容器（Container）

镜像的运行时，可以对外提供服务。本质上讲是利用namespace和cgroup等技术在宿主机中创建的独立的虚拟空间。

###### 仓库（Registry）

- 公有仓库，Docker Hub，阿里，网易...
- 私有仓库，企业内部搭建
  - Docker Registry，Docker官方提供的镜像仓库存储服务
  - Harbor, 是Docker Registry的更高级封装，它除了提供友好的Web UI界面，角色和用户权限管理，用户操作审计等功能 
- 镜像访问地址形式 registry.devops.com/demo/hello:latest,若没有前面的url地址，则默认寻找Docker Hub中的镜像，若没有tag标签，则使用latest作为标签
- 公有的仓库中，一般存在这么几类镜像
  - 操作系统基础镜像（centos，ubuntu，suse，alpine）
  - 中间件（nginx，redis，mysql，tomcat）
  - 语言编译环境（python，java，golang）
  - 业务镜像（django-demo...）

###### 操作演示

![](https://i0.hdslb.com/bfs/album/3a8c04d98b41e8a860dfd16ba858694c5ff0ebd6.jpg)

1. 解压离线包

   为了保证镜像下载的速度，因此提前在一台节点下载了离线镜像包，做解压：

   ```powershell
   $ tar zxf registry.tar.gz -C /opt
   $ ll /opt/registry-data
   total 25732
   drwxr-xr-x 3 root root     4096 Apr  9 20:11 registry
   -rw------- 1 root root 26344448 Apr  9 22:15 registry-v2.tar
   ```

2. 查看所有镜像：

```powershell
$ docker images
```

2. 拉取镜像:

```powershell
$ docker pull nginx:alpine
```

3. 如何唯一确定镜像:

- image_id
- repository:tag

```powershell
$ docker images
REPOSITORY    TAG                 IMAGE ID            CREATED             SIZE
nginx         alpine              377c0837328f        2 weeks ago         19.7MB
```

4. 导出镜像到文件中

   ```powershell
$ docker save -o nginx-alpine.tar nginx:alpine
   ```
   
5. 从文件中加载镜像

   ```powershell
   $ docker load -i nginx-alpine.tar
   ```

6. 部署镜像仓库

    https://docs.docker.com/registry/ 

   ```powershell
   ## 使用docker镜像启动镜像仓库服务
   $ docker run -d -p 5000:5000 --restart always -v /opt/registry-data/registry:/var/lib/registry --name registry registry:2
   
   ## 默认仓库不带认证，若需要认证，参考https://docs.docker.com/registry/deploying/#restricting-access
   ```

   假设启动镜像仓库服务的主机地址为172.21.32.6，该目录中已存在的镜像列表：

   | 现镜像仓库地址                                               | 原镜像仓库地址                                               |
   | ------------------------------------------------------------ | ------------------------------------------------------------ |
   | 172.21.32.6:5000/coreos/flannel:v0.11.0-amd64                | quay.io/coreos/flannel:v0.11.0-amd64                         |
   | 172.21.32.6:5000/mysql:5.7                                   | mysql:5.7                                                    |
   | 172.21.32.6:5000/nginx:alpine                                | nginx:alpine                                                 |
   | 172.21.32.6:5000/centos:centos7.5.1804                       | centos:centos7.5.1804                                        |
   | 172.21.32.6:5000/elasticsearch/elasticsearch:7.4.2           | docker.elastic.co/elasticsearch/elasticsearch:7.4.2          |
   | 172.21.32.6:5000/fluentd-es-root:v1.6.2-1.0                  | gcr.io/google_containers/fluentd-elasticsearch:v2.4.0        |
   | 172.21.32.6:5000/kibana/kibana:7.4.2                         | docker.elastic.co/kibana/kibana:7.4.2                        |
   | 172.21.32.6:5000/kubernetesui/dashboard:v2.0.0-beta5         | kubernetesui/dashboard:v2.0.0-beta5                          |
   | 172.21.32.6:5000/kubernetesui/metrics-scraper:v1.0.1         | kubernetesui/metrics-scraper:v1.0.1                          |
   | 172.21.32.6:5000/kubernetes-ingress-controller/nginx-ingress-controller:0.30.0 | quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.30.0 |

   

7. 推送本地镜像到镜像仓库中

   ```powershell
   $ docker tag nginx:alpine localhost:5000/nginx:alpine
   $ docker push localhost:5000/nginx:alpine
   ## 我的镜像仓库给外部访问，不能通过localhost，尝试使用内网地址172.21.16.3:5000/nginx:alpine
   $ docker tag nginx:alpine 172.21.16.3:5000/nginx:alpine
   $ docker push 172.21.16.3:5000/nginx:alpine
   The push refers to repository [172.21.16.3:5000/nginx]
   Get https://172.21.16.3:5000/v2/: http: server gave HTTP response to HTTPS client
   ## docker默认不允许向http的仓库地址推送，如何做成https的，参考：https://docs.docker.com/registry/deploying/#run-an-externally-accessible-registry
   ## 我们没有可信证书机构颁发的证书和域名，自签名证书需要在每个节点中拷贝证书文件，比较麻烦，因此我们通过配置daemon的方式，来跳过证书的验证：
   $ cat /etc/docker/daemon.json
   {
     "registry-mirrors": [
       "https://8xpk5wnt.mirror.aliyuncs.com"
     ],
     "insecure-registries": [
        "172.21.16.3:5000"
     ]
   }
   $ systemctl restart docker
   $ docker push 172.21.16.3:5000/nginx:alpine
   $ docker images	# IMAGE ID相同，等于起别名或者加快捷方式
   REPOSITORY               TAG                 IMAGE ID            CREATED             SIZE
   172.21.16.3:5000/nginx   alpine              377c0837328f        4 weeks ago         
   nginx                    alpine              377c0837328f        4 weeks ago         
   localhost:5000/nginx     alpine              377c0837328f        4 weeks ago         
   registry                 2                   708bc6af7e5e        2 months ago       
   ```

8. 删除镜像

   ```powershell
   docker rmi nginx:alpine
   ```

9. 查看容器列表

   ```powershell
   ## 查看运行状态的容器列表
   $ docker ps
   
   ## 查看全部状态的容器列表
   $ docker ps -a
   ```

10. 启动容器

    ```powershell
    ## 后台启动
    $ docker run --name nginx -d nginx:alpine
    ##查看run流程#
    ##查看容器进程
    ## 等同于在虚拟机中开辟了一块隔离的独立的虚拟空间
    ## 启动容器的同时进入容器，-ti与/bin/sh或者/bin/bash配套使用，意思未分配一个tty终端
    $ docker run --name nginx -ti nginx:alpine /bin/sh
    （注意：退出容器后，该容器会变成退出状态，因为容器内部的1号进程退出）
    
    ## 实际上，在运行容器的时候，镜像地址后面跟的命令等于是覆盖了原有的容器的CMD命令，因此，执行的这些命令在容器内部就是1号进程，若该进程不存在了，那么容器就会处于退出的状态，比如，宿主机中执行
    1. echo 1,执行完后，该命令立马就结束了
    2. ping www.baidu.com,执行完后，命令的进程会持续运行
    $ docker run -d --name test_echo nginx:alpine echo 1,容器会立马退出
    $ docker run -d --name test_ping nginx:alpine ping www.baidu.com,容器不会退出，但是因为没有加-d参数，因此一直在前台运行，若ctrl+C终止，则容器退出，因为1号进程被终止了
    
    
    ## 映射端口,把容器的端口映射到宿主机中,-p <host_port>:<container_port>
    $ docker run --name nginx -d -p 8080:80 nginx:alpine
    
    ## 资源限制,-cpuset-cpus用于设置容器可以使用的 vCPU 核。-c,--cpu-shares用于设置多个容器竞争 CPU 时，各个容器相对能分配到的 CPU 时间比例。假设有三个正在运行的容器，这三个容器中的任务都是 CPU 密集型的。第一个容器的 cpu 共享权值是 1024，其它两个容器的 cpu 共享权值是 512。第一个容器将得到 50% 的 CPU 时间，而其它两个容器就只能各得到 25% 的 CPU 时间了。如果再添加第四个 cpu 共享值为 1024 的容器，每个容器得到的 CPU 时间将重新计算。第一个容器的CPU 时间变为 33%，其它容器分得的 CPU 时间分别为 16.5%、16.5%、33%。必须注意的是，这个比例只有在 CPU 密集型的任务执行时才有用。在四核的系统上，假设有四个单进程的容器，它们都能各自使用一个核的 100% CPU 时间，不管它们的 cpu 共享权值是多少。
    $ docker run --cpuset-cpus="0-3" --cpu-shares=512 --memory=500m nginx:alpine
    ```

    ![](https://i0.hdslb.com/bfs/album/e9045be78fd67d0d476db74cfb3965bcfc9603d1.jpg)

11. 容器数据持久化

    ```powershell
    ## 挂载主机目录
    $ docker run --name nginx -d  -v /opt:/opt -v /var/log:/var/log nginx:alpine
    $ docker run --name mysql -d -v /opt/mysql/:/var/lib/mysql mysql:5.7
    
    ## 使用volumes卷
    $ docker volume ls
    $ docker volume create my-vol
    $ docker run --name nginx -d -v my-vol:/opt/my-vol nginx:alpine
    $ docker exec -ti nginx touch /opt/my-vol/a.txt
    
    ## 验证数据共享
    $ docker run --name nginx2 -d -v my-vol:/opt/hh nginx:alpine
    $ docker exec -ti nginx2 ls /opt/hh/
    a.txt
    ```

12. 进入容器或者执行容器内的命令

    ```powershell
    $ docker exec -ti <container_id_or_name> /bin/sh
    $ docker exec -ti <container_id_or_name> hostname
    ```

13. 主机与容器之间拷贝数据

    ```powershell
    ## 主机拷贝到容器
    $ echo '123'>/tmp/test.txt
    $ docker cp /tmp/test.txt nginx:/tmp
    $ docker exec -ti nginx cat /tmp/test.txt
    123
    
    ## 容器拷贝到主机
    $ docker cp nginx:/tmp/test.txt ./
    ```

14. 查看容器日志

    ```powershell
    ## 查看全部日志
    $ docker logs nginx
    
    ## 实时查看最新日志
    $ docker logs -f nginx
    
    ## 从最新的100条开始查看
    $ docker logs --tail=100 -f nginx
    ```

15. 停止或者删除容器

    ```powershell
    ## 停止运行中的容器
    $ docker stop nginx
    
    ## 启动退出容器
    $ docker start nginx
    
    ## 删除退出容器
    $ docker rm nginx
    
    ## 删除运行中的容器
    $ docker rm -f nginx
    ```

16. 查看容器或者镜像的明细

    ```powershell
    ## 查看容器详细信息，包括容器IP地址等
    $ docker inspect nginx
    
    ## 查看镜像的明细信息
    $ docker inspect nginx:alpine
    ```

    

#### Django应用容器化实践

###### django项目介绍

- 项目地址：https://gitee.com/agagin/python-demo.git
- python3 + uwsgi + nginx + mysql

- 内部服务端口8002

###### 构建命令

```powershell
$ docker build . -t ImageName:ImageTag -f Dockerfile
```

如何理解构建镜像的过程？

Dockerfile是一堆指令，在docker build的时候，按照该指令进行操作，最终生成我们期望的镜像

- FROM 指定基础镜像，必须为第一个命令

  ```
  格式：
  	FROM <image>
  	FROM <image>:<tag>
  示例：
  	FROM mysql:5.7
  注意：
  	tag是可选的，如果不使用tag时，会使用latest版本的基础镜像
  ```

- MAINTAINER 镜像维护者的信息

  ```
  格式：
  	MAINTAINER <name>
  示例：
  	MAINTAINER Yongxin Li
      MAINTAINER inspur_lyx@hotmail.com
      MAINTAINER Yongxin Li <inspur_lyx@hotmail.com>
  ```

- COPY|ADD 添加本地文件到镜像中

  ```
  格式：
  	COPY <src>... <dest>
  示例：
      ADD hom* /mydir/          # 添加所有以"hom"开头的文件
      ADD test relativeDir/     # 添加 "test" 到 `WORKDIR`/relativeDir/
      ADD test /absoluteDir/    # 添加 "test" 到 /absoluteDir/
  ```

- WORKDIR 工作目录

  ```
  格式：
  	WORKDIR /path/to/workdir
  示例：
      WORKDIR /a  (这时工作目录为/a)
  注意：
  	通过WORKDIR设置工作目录后，Dockerfile中其后的命令RUN、CMD、ENTRYPOINT、ADD、COPY等命令都会在该目录下执行
  ```

- RUN 构建镜像过程中执行命令

  ```
  格式：
  	RUN <command>
  示例：
      RUN yum install nginx
      RUN pip install django
      RUN mkdir test && rm -rf /var/lib/unusedfiles
  注意：
  	RUN指令创建的中间镜像会被缓存，并会在下次构建中使用。如果不想使用这些缓存镜像，可以在构建时指定--no-cache参数，如：docker build --no-cache
  ```

- CMD 构建容器后调用，也就是在容器启动时才进行调用

  ```
  格式：
      CMD ["executable","param1","param2"] (执行可执行文件，优先)
      CMD ["param1","param2"] (设置了ENTRYPOINT，则直接调用ENTRYPOINT添加参数)
      CMD command param1 param2 (执行shell内部命令)
  示例：
      CMD ["/usr/bin/wc","--help"]
      CMD ping www.baidu.com
  注意：
  	CMD不同于RUN，CMD用于指定在容器启动时所要执行的命令，而RUN用于指定镜像构建时所要执行的命令。
  ```

- ENTRYPOINT 设置容器初始化命令，使其可执行化

  ```
  格式：
      ENTRYPOINT ["executable", "param1", "param2"] (可执行文件, 优先)
      ENTRYPOINT command param1 param2 (shell内部命令)
  示例：
      ENTRYPOINT ["/usr/bin/wc","--help"]
  注意：
  	ENTRYPOINT与CMD非常类似，不同的是通过docker run执行的命令不会覆盖ENTRYPOINT，而docker run命令中指定的任何参数，都会被当做参数再次传递给ENTRYPOINT。Dockerfile中只允许有一个ENTRYPOINT命令，多指定时会覆盖前面的设置，而只执行最后的ENTRYPOINT指令
  ```

- ENV

  ```
  格式：
      ENV <key> <value>
      ENV <key>=<value>
  示例：
      ENV myName John
      ENV myCat=fluffy
  ```

- EXPOSE

  ```
  格式：
      EXPOSE <port> [<port>...]
  示例：
      EXPOSE 80 443
      EXPOSE 8080
      EXPOSE 11211/tcp 11211/udp
  注意：
      EXPOSE并不会让容器的端口访问到主机。要使其可访问，需要在docker run运行容器时通过-p来发布这些端口，或通过-P参数来发布EXPOSE导出的所有端口
  
  ```

  ![](https://i0.hdslb.com/bfs/album/3d070d8b05e2a2b905c35ee46a6fe0056ab95b09.png)

###### Dockerfile

*dockerfiles/Dockerfile*

```dockerfile
# This my first django Dockerfile
# Version 1.0

# Base images 基础镜像
FROM centos:centos7.5.1804

#MAINTAINER 维护者信息
LABEL maintainer="inspur_lyx@hotmail.com"

#ENV 设置环境变量
ENV LANG en_US.UTF-8
ENV LC_ALL en_US.UTF-8

#RUN 执行以下命令
RUN curl -so /etc/yum.repos.d/Centos-7.repo http://mirrors.aliyun.com/repo/Centos-7.repo
RUN yum install -y  python36 python3-devel gcc pcre-devel zlib-devel make net-tools

#工作目录
WORKDIR /opt/myblog

#拷贝文件至工作目录
COPY . .

#安装nginx
RUN tar -zxf nginx-1.13.7.tar.gz -C /opt  && cd /opt/nginx-1.13.7 && ./configure --prefix=/usr/local/nginx \
&& make && make install && ln -s /usr/local/nginx/sbin/nginx /usr/bin/nginx

RUN cp myblog.conf /usr/local/nginx/conf/myblog.conf

#安装依赖的插件
RUN pip3 install -i http://mirrors.aliyun.com/pypi/simple/ --trusted-host mirrors.aliyun.com -r requirements.txt

RUN chmod +x run.sh && rm -rf ~/.cache/pip

#EXPOSE 映射端口
EXPOSE 8002

#容器启动时执行命令
CMD ["./run.sh"]
```

执行构建：

```powershell
$ docker build . -t myblog:v1 -f Dockerfile
```



###### 定制化基础镜像

`dockerfiles/myblog/Dockerfile-base`

```dockerfile
# Base images 基础镜像
FROM centos:centos7.5.1804

#MAINTAINER 维护者信息
LABEL maintainer="inspur_lyx@hotmail.com"

#ENV 设置环境变量
ENV LANG en_US.UTF-8
ENV LC_ALL en_US.UTF-8

#RUN 执行以下命令
RUN curl -so /etc/yum.repos.d/Centos-7.repo http://mirrors.aliyun.com/repo/Centos-7.repo
RUN yum install -y  python36 python3-devel gcc pcre-devel zlib-devel make net-tools

COPY nginx-1.13.7.tar.gz  /opt

#安装nginx
RUN tar -zxf /opt/nginx-1.13.7.tar.gz -C /opt  && cd /opt/nginx-1.13.7 && ./configure --prefix=/usr/local/nginx && make && make install && ln -s /usr/local/nginx/sbin/nginx /usr/bin/nginx
```

```powershell
## 构建基础镜像
$ docker build . -t centos-python3-nginx:v1 -f Dockerfile-base
$ docker tag centos-python3-nginx:v1 172.21.32.6:5000/base/centos-python3-nginx:v1
$ docker push 172.21.32.6:5000/base/centos-python3-nginx:v1
```

###### 简化Dockerfile

`dockerfiles/myblog/Dockerfile-optimized`

```dockerfile
# This my first django Dockerfile
# Version 1.0

# Base images 基础镜像
FROM centos-python3-nginx:v1

#MAINTAINER 维护者信息
LABEL maintainer="inspur_lyx@hotmail.com"

#工作目录
WORKDIR /opt/myblog

#拷贝文件至工作目录
COPY . .

RUN cp myblog.conf /usr/local/nginx/conf/myblog.conf

#安装依赖的插件
RUN pip3 install -i http://mirrors.aliyun.com/pypi/simple/ --trusted-host mirrors.aliyun.com -r requirements.txt

RUN chmod +x run.sh && rm -rf ~/.cache/pip

#EXPOSE 映射端口
EXPOSE 8002

#容器启动时执行命令
CMD ["./run.sh"]
```

```powershell
$ docker build . -t myblog -f Dockerfile-optimized
```



###### 运行mysql

```powershell
$ docker run -d -p 3306:3306 --name mysql  -v /opt/mysql/mysql-data/:/var/lib/mysql -e MYSQL_DATABASE=myblog -e MYSQL_ROOT_PASSWORD=123456 mysql:5.7

## 查看数据库
$ docker exec -ti mysql bash
#/ mysql -uroot -p123456
#/ show databases;

## navicator连接
```



###### 启动Django应用

```powershell
## 启动容器
$ docker run -d -p 8002:8002 --name myblog -e MYSQL_HOST=172.21.32.6 -e MYSQL_USER=root -e MYSQL_PASSWD=123456  myblog

## migrate
$ docker exec -ti myblog bash
#/ python3 manage.py makemigrations
#/ python3 manage.py migrate
#/ python3 manage.py createsuperuser

## 创建超级用户
$ docker exec -ti myblog python3 manage.py createsuperuser

## 收集静态文件
## $ docker exec -ti myblog python3 manage.py collectstatic
```

访问62.234.214.206:8002/admin

构建镜像，替换默认编码：

`dockerfiles/mysql/my.cnf`

```powershell
$ cat my.cnf
[mysqld]
user=root
character-set-server=utf8
lower_case_table_names=1

[client]
default-character-set=utf8
[mysql]
default-character-set=utf8

!includedir /etc/mysql/conf.d/
!includedir /etc/mysql/mysql.conf.d/
```

`dockerfiles/mysql/Dockerfile`

```dockerfile
FROM mysql:5.7
COPY my.cnf /etc/mysql/my.cnf
## CMD或者ENTRYPOINT默认继承
```

```powershell
$ docker build . -t mysql:5.7-utf8
$ docker tag mysql:5.7-utf8 172.21.16.3:5000/mysql:5.7-utf8
$ docker push 172.21.16.3:5000/mysql:5.7-utf8

## 删除旧的mysql容器，使用新镜像启动,不用再次初始化
$ docker rm -f mysql
$ rm -rf /opt/mysql/mysql-data/*
$ docker run -d -p 3306:3306 --name mysql -v /opt/mysql/mysql-data/:/var/lib/mysql -e MYSQL_DATABASE=myblog -e MYSQL_ROOT_PASSWORD=123456 172.21.32.6:5000/mysql:5.7-utf8

## 重新migrate
$ docker exec -ti myblog bash
#/ python3 manage.py makemigrations
#/ python3 manage.py migrate
#/ python3 manage.py createsuperuser
```



#### 实现原理

虚拟化核心需要解决的问题：资源隔离与资源限制

- 虚拟机硬件虚拟化技术， 通过一个 hypervisor 层实现对资源的彻底隔离。
- 容器则是操作系统级别的虚拟化，利用的是内核的 Cgroup 和 Namespace 特性，此功能完全通过软件实现。

###### Namespace 资源隔离

命名空间是全局资源的一种抽象，将资源放到不同的命名空间中，各个命名空间中的资源是相互隔离的。 通俗来讲，就是docker在启动一个容器的时候，会调用Linux Kernel Namespace的接口，来创建一块虚拟空间，创建的时候，可以支持设置下面这几种（可以随意选择）,docker默认都设置。

- pid：用于进程隔离（PID：进程ID）
- net：管理网络接口（NET：网络）
- ipc：管理对 IPC 资源的访问（IPC：进程间通信（信号量、消息队列和共享内存））
- mnt：管理文件系统挂载点（MNT：挂载）
- uts：隔离主机名和域名
- user：隔离用户和用户组（3.8以后的内核才支持）

```go
func setNamespaces(daemon *Daemon, s *specs.Spec, c *container.Container) error {
    // user
    // network
    // ipc
    // uts
    // pid
    if c.HostConfig.PidMode.IsContainer() {
        ns := specs.LinuxNamespace{Type: "pid"}
        pc, err := daemon.getPidContainer(c)
        if err != nil {
            return err
        }
        ns.Path = fmt.Sprintf("/proc/%d/ns/pid", pc.State.GetPID())
        setNamespace(s, ns)
    } else if c.HostConfig.PidMode.IsHost() {
        oci.RemoveNamespace(s, specs.LinuxNamespaceType("pid"))
    } else {
        ns := specs.LinuxNamespace{Type: "pid"}
        setNamespace(s, ns)
    }
    return nil
}
```

###### CGroup 资源限制

通过namespace可以保证容器之间的隔离，但是无法控制每个容器可以占用多少资源， 如果其中的某一个容器正在执行 CPU 密集型的任务，那么就会影响其他容器中任务的性能与执行效率，导致多个容器相互影响并且抢占资源。如何对多个容器的资源使用进行限制就成了解决进程虚拟资源隔离之后的主要问题。

![](https://i0.hdslb.com/bfs/album/f88da30754255738e08c5949fa536156324fa69b.png)

Control Groups（简称 CGroups）就是能够隔离宿主机器上的物理资源，例如 CPU、内存、磁盘 I/O 和网络带宽。每一个 CGroup 都是一组被相同的标准和参数限制的进程。而我们需要做的，其实就是把容器这个进程加入到指定的Cgroup中。深入理解CGroup，请[点此](![image-20200323195718300](C:\Users\liyongxin\AppData\Roaming\Typora\typora-user-images\image-20200323195718300.png))。

###### UnionFS 联合文件系统

Linux namespace和cgroup分别解决了容器的资源隔离与资源限制，那么容器是很轻量的，通常每台机器中可以运行几十上百个容器， 这些个容器是共用一个image，还是各自将这个image复制了一份，然后各自独立运行呢？ 如果每个容器之间都是全量的文件系统拷贝，那么会导致至少如下问题：

- 运行容器的速度会变慢
- 容器和镜像对宿主机的磁盘空间的压力

怎么解决这个问题------Docker的存储驱动

- 镜像分层存储
-  UnionFS

Docker 镜像是由一系列的层组成的，每层代表 Dockerfile 中的一条指令，比如下面的 Dockerfile 文件：

```dockerfile
FROM ubuntu:15.04
COPY . /app
RUN make /app
CMD python /app/app.py
```

这里的 Dockerfile 包含4条命令，其中每一行就创建了一层，下面显示了上述Dockerfile构建出来的镜像运行的容器层的结构：

![](https://i0.hdslb.com/bfs/album/c6354991c7892a38207dddf9f07b4328d3dbeed3.jpg)

镜像就是由这些层一层一层堆叠起来的，镜像中的这些层都是只读的，当我们运行容器的时候，就可以在这些基础层至上添加新的可写层，也就是我们通常说的`容器层`，对于运行中的容器所做的所有更改（比如写入新文件、修改现有文件、删除文件）都将写入这个容器层。

对容器层的操作，主要利用了写时复制（CoW）技术。CoW就是copy-on-write，表示只在需要写时才去复制，这个是针对已有文件的修改场景。 CoW技术可以让所有的容器共享image的文件系统，所有数据都从image中读取，只有当要对文件进行写操作时，才从image里把要写的文件复制到自己的文件系统进行修改。所以无论有多少个容器共享同一个image，所做的写操作都是对从image中复制到自己的文件系统中的复本上进行，并不会修改image的源文件，且多个容器操作同一个文件，会在每个容器的文件系统里生成一个复本，每个容器修改的都是自己的复本，相互隔离，相互不影响。使用CoW可以有效的提高磁盘的利用率。 

![](https://i0.hdslb.com/bfs/album/82b3e992947958d7935aeb52753c9e6a2f07e171.jpg)

镜像中每一层的文件都是分散在不同的目录中的，如何把这些不同目录的文件整合到一起呢？

UnionFS 其实是一种为 Linux 操作系统设计的用于把多个文件系统联合到同一个挂载点的文件系统服务。  它能够将不同文件夹中的层联合（Union）到了同一个文件夹中，整个联合的过程被称为联合挂载（Union Mount）。

![](https://i0.hdslb.com/bfs/album/dc779a853a48ff36651c03a86b4d2a66d0dd7924.png)

上图是AUFS的实现，AUFS是作为Docker存储驱动的一种实现，Docker 还支持了不同的存储驱动，包括 aufs、devicemapper、overlay2、zfs 和  Btrfs 等等，在最新的 Docker 中，overlay2 取代了 aufs 成为了推荐的存储驱动，但是在没有 overlay2 驱动的机器上仍然会使用 aufs 作为 Docker 的默认驱动。 

#### Docker网络

docker容器是一块具有隔离性的虚拟系统，容器内可以有自己独立的网络空间，

- 多个容器之间是如何实现通信的呢？
- 容器和宿主机之间又是如何实现的通信呢？
- 使用-p参数是怎么实现的端口映射?

带着我们就这些问题，我们来学习一下docker的网络模型，最后我会通过抓包的方式，给大家演示一下数据包在容器和宿主机之间的转换过程。

##### 网络模式

我们在使用docker run创建Docker容器时，可以用--net选项指定容器的网络模式，Docker有以下4种网络模式：

- bridge模式，使用--net=bridge指定，默认设置

- host模式，使用--net=host指定，容器内部网络空间共享宿主机的空间，效果类似直接在宿主机上启动一个进程，端口信息和宿主机共用。

- container模式，使用--net=container:NAME_or_ID指定

  指定容器与特定容器共享网络命名空间

- none模式，使用--net=none指定

  网络模式为空，即仅保留网络命名空间，但是不做任何网络相关的配置(网卡、IP、路由等)

##### bridge模式

那我们之前在演示创建docker容器的时候其实是没有指定的网络模式的，如果不指定的话默认就会使用bridge模式，bridge本意是桥的意思，其实就是网桥模式，那我们怎么理解网桥，如果需要做类比的话，我们可以把网桥看成一个二层的交换机设备，我们来看下这张图：

交换机通信简图

![](https://i0.hdslb.com/bfs/album/b45b14439190162440b44b7342e40ab3acd49399.png)



网桥模式示意图

![](https://i0.hdslb.com/bfs/album/b1b4245b8251894fc8e8677132d8f62a501ac6b7.png)



网桥在哪，查看网桥

```powershell
$ yum install -y bridge-utils
$ brctl show
bridge name     bridge id               STP enabled     interfaces
docker0         8000.0242b5fbe57b       no              veth3a496ed
```

有了网桥之后，那我们看下docker在启动一个容器的时候做了哪些事情才能实现容器间的互联互通

Docker 创建一个容器的时候，会执行如下操作：

- 创建一对虚拟接口/网卡，也就是veth pair；
- 本地主机一端桥接 到默认的 docker0 或指定网桥上，并具有一个唯一的名字，如 veth9953b75；
- 容器一端放到新启动的容器内部，并修改名字作为 eth0，这个网卡/接口只在容器的命名空间可见；
- 从网桥可用地址段中（也就是与该bridge对应的network）获取一个空闲地址分配给容器的 eth0
- 配置默认路由到网桥

那整个过程其实是docker自动帮我们完成的，清理掉所有容器，来验证。

```powershell
## 清掉所有容器
$ docker rm -f `docker ps -aq`
$ docker ps
$ brctl show # 查看网桥中的接口，目前没有

## 创建测试容器test1
$ docker run -d --name test1 nginx:alpine
$ brctl show # 查看网桥中的接口，已经把test1的veth端接入到网桥中
$ ip a |grep veth # 已在宿主机中可以查看到
$ docker exec -ti test1 sh 
/ # ifconfig  # 查看容器的eth0网卡及分配的容器ip
/ # route -n  # 观察默认网关都指向了网桥的地址，即所有流量都转向网桥，等于是在veth pair接通了网线
Kernel IP routing table
Destination     Gateway         Genmask         Flags Metric Ref    Use Iface
0.0.0.0         172.17.0.1      0.0.0.0         UG    0      0        0 eth0
172.17.0.0      0.0.0.0         255.255.0.0     U     0      0        0 eth0

# 再来启动一个测试容器，测试容器间的通信
$ docker run -d --name test2 nginx:alpine
$ docker exec -ti test sh
/ # sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories
/ # apk add curl
/ # curl 172.17.0.8:80

## 为啥可以通信，因为两个容器是接在同一个网桥中的，通信其实是通过mac地址和端口的的记录来做转发的。test1访问test2，通过test1的eth0发送ARP广播，网桥会维护一份mac映射表，我们可以大概通过命令来看一下，
$ brctl showmacs docker0
## 这些mac地址是主机端的veth网卡对应的mac，可以查看一下
$ ip a 

```

我们如何知道网桥上的这些虚拟网卡与容器端是如何对应？

通过ifindex，网卡索引号

```powershell
## 查看test1容器的网卡索引
$ docker exec -ti test1 cat /sys/class/net/eth0/ifindex

## 主机中找到虚拟网卡后面这个@ifxx的值，如果是同一个值，说明这个虚拟网卡和这个容器的eth0网卡是配对的。
$ ip a |grep @if
```

整理脚本，快速查看对应：

```powershell
for container in $(docker ps -q); do
    iflink=`docker exec -it $container sh -c 'cat /sys/class/net/eth0/iflink'`
    iflink=`echo $iflink|tr -d '\r'`
    veth=`grep -l $iflink /sys/class/net/veth*/ifindex`
    veth=`echo $veth|sed -e 's;^.*net/\(.*\)/ifindex$;\1;'`
    echo $container:$veth
done
```



上面我们讲解了容器之间的通信，那么容器与宿主机的通信是如何做的？

![](https://i0.hdslb.com/bfs/album/d9e4a028cf9557a6d8ccec57ab8ac5689d7bca76.png)



添加端口映射：

```powershell
## 启动容器的时候通过-p参数添加宿主机端口与容器内部服务端口的映射
$ docker run --name test -d -p 8088:80 nginx:alpine
$ curl localhost:8088
```

端口映射如何实现的？先来回顾iptables链表图

![](https://i0.hdslb.com/bfs/album/2af4d0946e9ba8ec9d7df36255d62d20e23b5279.png)



访问本机的8088端口，数据包会从流入方向进入本机，因此涉及到PREROUTING和INPUT链，我们是通过做宿主机与容器之间加的端口映射，所以肯定会涉及到端口转换，那哪个表是负责存储端口转换信息的呢，就是nat表，负责维护网络地址转换信息的。因此我们来查看一下PREROUTING链的nat表：

```powershell
$ iptables -t nat -nvL PREROUTING
Chain PREROUTING (policy ACCEPT 159 packets, 20790 bytes)
 pkts bytes target     prot opt in     out     source               destination
    3   156 DOCKER     all  --  *      *       0.0.0.0/0            0.0.0.0/0            ADDRTYPE match dst-type LOCAL
```

规则利用了iptables的addrtype拓展，匹配网络类型为本地的包，如何确定哪些是匹配本地，

```powershell
$ ip route show table local type local
local 127.0.0.0/8 dev lo proto kernel scope host src 127.0.0.1
local 127.0.0.1 dev lo proto kernel scope host src 127.0.0.1
local 172.17.0.1 dev docker0 proto kernel scope host src 172.17.0.1
local 192.168.136.133 dev ens33 proto kernel scope host src 192.168.136.133
```

也就是说目标地址类型匹配到这些的，会转发到我们的TARGET中，TARGET是动作，意味着对符合要求的数据包执行什么样的操作，最常见的为ACCEPT或者DROP，此处的TARGET为DOCKER，很明显DOCKER不是标准的动作，那DOCKER是什么呢？我们通常会定义自定义的链，这样把某类对应的规则放在自定义链中，然后把自定义的链绑定到标准的链路中，因此此处DOCKER 是自定义的链。那我们现在就来看一下DOCKER这个自定义链上的规则。

```powershell
$ iptables -t nat -nvL DOCKER
Chain DOCKER (2 references)                                                                                                
 pkts bytes target     prot opt in     out     source               destination                                            
    0     0 RETURN     all  --  docker0 *       0.0.0.0/0            0.0.0.0/0                                             
    0     0 DNAT       tcp  --  !docker0 *       0.0.0.0/0            0.0.0.0/0            tcp dpt:8088 to:172.17.0.2:80 

```

此条规则就是对主机收到的目的端口为8088的tcp流量进行DNAT转换，将流量发往172.17.0.2:80，172.17.0.2地址是不是就是我们上面创建的Docker容器的ip地址，流量走到网桥上了，后面就走网桥的转发就ok了。
所以，外界只需访问192.168.136.133:8088就可以访问到容器中的服务了。

 数据包在出口方向走POSTROUTING链，我们查看一下规则：

```powershell
$ iptables -t nat -nvL POSTROUTING
Chain POSTROUTING (policy ACCEPT 1099 packets, 67268 bytes)
 pkts bytes target     prot opt in     out     source               destination
   86  5438 MASQUERADE  all  --  *      !docker0  172.17.0.0/16        0.0.0.0/0
    0     0 MASQUERADE  tcp  --  *      *       172.17.0.4           172.17.0.4           tcp dpt:80
```

大家注意MASQUERADE这个动作是什么意思，其实是一种更灵活的SNAT，把源地址转换成主机的出口ip地址，那解释一下这条规则的意思:

这条规则会将源地址为172.17.0.0/16的包（也就是从Docker容器产生的包），并且不是从docker0网卡发出的，进行源地址转换，转换成主机网卡的地址。大概的过程就是ACK的包在容器里面发出来，会路由到网桥docker0，网桥根据宿主机的路由规则会转给宿主机网卡eth0，这时候包就从docker0网卡转到eth0网卡了，并从eth0网卡发出去，这时候这条规则就会生效了，把源地址换成了eth0的ip地址。

> 注意一下，刚才这个过程涉及到了网卡间包的传递，那一定要打开主机的ip_forward转发服务，要不然包转不了，服务肯定访问不到。
> 	

###### 抓包演示

我们先想一下，我们要抓哪个网卡的包

- 首先访问宿主机的8088端口，我们抓一下宿主机的eth0

  ```powershell
  $ tcpdump -i eth0 port 8088 -w host.cap
  ```

- 然后最终包会流入容器内，那我们抓一下容器内的eth0网卡

  ```powershell
  # 容器内安装一下tcpdump
  $ sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories
  $ apk add tcpdump
  $ tcpdump -i eth0 port 80 -w container.cap
  ```

到另一台机器访问一下，

```powershell
$ curl 172.21.32.6:8088/
```

停止抓包，拷贝容器内的包到宿主机

```powershell
$ docker cp test:/root/container.cap /root/
```

把抓到的内容拷贝到本地，使用wireshark进行分析。

```powershell
$ scp root@172.21.32.6:/root/*.cap /d/packages
```

（wireshark合并包进行分析）

![](https://i0.hdslb.com/bfs/album/e0a1a22c02674617892dd370a49bb6be02bbe5b7.png)

![](https://i0.hdslb.com/bfs/album/138caa0f0578d8068f7f8d076efb77333a6f6749.png)



进到容器内的包做DNAT，出去的包做SNAT，这样对外面来讲，根本就不知道机器内部是谁提供服务，其实这就和一个内网多个机器公用一个外网IP地址上网的效果是一样的，对吧，那这也属于NAT功能的一个常见的应用场景。

##### Host模式

容器内部不会创建网络空间，共享宿主机的网络空间

```powershell
$ docker run --net host -d --name mysql:5.7
```

##### Conatiner模式

这个模式指定新创建的容器和已经存在的一个容器共享一个 Network Namespace，而不是和宿主机共享。新创建的容器不会创建自己的网卡，配置自己的 IP，而是和一个指定的容器共享 IP、端口范围等。同样，两个容器除了网络方面，其他的如文件系统、进程列表等还是隔离的。两个容器的进程可以通过 lo 网卡设备通信。 

![](https://i0.hdslb.com/bfs/album/4f25b4bad3ec8e5d5f113bba1aed72cdfbcd193e.jpg)

```powershell
## 启动测试容器，共享mysql的网络空间
$ docker run -ti --rm --net=container:mysql busybox sh
/ # ip a
/ # netstat -tlp|grep 3306
/ # telnet localhost 3306
```



#### 实用技巧

1. 清理主机上所有退出的容器

   ```powershell
   $ docker rm $(docker ps -aq)
   ```

2. 调试或者排查容器启动错误

   ```powershell
   ## 若有时遇到容器启动失败的情况，可以先使用相同的镜像启动一个临时容器，先进入容器
   $ docker exec -ti --rm <image_id> bash
   ## 进入容器后，手动执行该容器对应的ENTRYPOINT或者CMD命令，这样即使出错，容器也不会退出，因为bash作为1号进程，我们只要不退出容器，该容器就不会自动退出
   ```

   

#### 本章小结

1. 为了解决软件交付过程中的环境依赖，同时提供一种更加轻量的虚拟化技术，Docker出现了

2. 2013年诞生，15年开始迅速发展，从17.03月开始，使用时间日期管理版本，稳定版以每季度为准

3. Docker是一种CS架构的软件产品，可以把代码及依赖打包成镜像，作为交付介质，并且把镜像启动成为容器，提供容器生命周期的管理

4. 使用yum部署docker，启动后通过操作docker这个命令行，自动调用docker daemon完成容器相关操作

5. 常用操作

   - systemctl  start|stop|restart docker
   - docker build | pull  -> docker tag -> docker push
   - docker run --name my-demo  -d  -p 8080:80 -v  /opt/data:/data  demo:v20200327
   - docker cp  /path/a.txt  mycontainer:/opt
   - docker exec -ti  mycontainer  /bin/sh
   - docker logs -f mycontainer

6. 通过dockerfile构建业务镜像，先使用基础镜像，然后通过一系列的指令把我们的业务应用所需要的运行环境和依赖都打包到镜像中，然后通过CMD或者ENTRYPOINT指令把镜像启动时的入口制定好，完成封装即可。有点类似于，先找来一个空的集装箱(基础镜像)，然后把项目依赖的服务都扔到集装箱中，然后设置好服务的启动入口，关闭箱门，即完成了业务镜像的制作。

7. 容器的实现依赖于内核模块提供的namespace和control-group的功能，通过namespace创建一块虚拟空间，空间内实现了各类资源(进程、网络、文件系统)的隔离，提供control-group实现了对隔离的空间的资源使用的限制。

8. docker镜像使用分层的方式进行存储，根据主机的存储驱动的不同，实现方式会不同，kernel在3.10.0-514以上自动支持overlay2 存储驱动，也是目前Docker推荐的方式。

9. 得益于分层存储的模式，多个容器可以通过copy-on-write的策略，在镜像的最上层加一个可写层，实现一个镜像快速启动多个容器的场景

10. docker的网络模式分为4种，最常用的为bridge和host模式。bridge模式通过docker0网桥，启动容器的时候通过创建一对虚拟网卡，将容器连接在桥上，同时维护了虚拟网卡与网桥端口的关系，实现容器间的通信。容器与宿主机之间的通信通过iptables端口映射的方式，docker利用iptables的PREROUTING和POSTROUTING的nat功能，实现了SNAT与DNAT，使得容器内部的服务被完美的保护起来。

    

### 第二章  Kubernetes实践之旅

本章学习kubernetes的架构及工作流程，重点介绍如何使用Deployment管理Pod生命周期，实现服务不中断的滚动更新，通过服务发现来实现集群内部的服务间访问，并通过ingress-nginx实现外部使用域名访问集群内部的服务。同时介绍基于EFK如何搭建Kubernetes集群的日志收集系统。

学完本章，我们的Django demo项目已经可以运行在k8s集群中，同时我们可以使用域名进行服务的访问。

- 架构及核心组件介绍
- 使用kubeadm快速搭建集群
- 运行第一个Pod应用
- Pod进阶
- Pod控制器的使用
- 实现服务与Node绑定的几种方式
- 负载均衡与服务发现
- 使用Ingress实现集群服务的7层代理
- Django项目k8s落地实践
- 基于EFK实现kubernetes集群的日志平台（扩展）
- 集群认证与授权

#### 纯容器模式的问题

1. 业务容器数量庞大，哪些容器部署在哪些节点，使用了哪些端口，如何记录、管理，需要登录到每台机器去管理？
2. 跨主机通信，多个机器中的容器之间相互调用如何做，iptables规则手动维护？
3. 跨主机容器间互相调用，配置如何写？写死固定IP+端口？
4. 如何实现业务高可用？多个容器对外提供服务如何实现负载均衡？
5. 容器的业务中断了，如何可以感知到，感知到以后，如何自动启动新的容器?
6. 如何实现滚动升级保证业务的连续性？
7. ......

#### 容器调度管理平台

Docker Swarm       Mesos        Google Kubernetes

2017年开始Kubernetes凭借强大的容器集群管理功能, 逐步占据市场,目前在容器编排领域一枝独秀

 https://kubernetes.io/ 

#### 架构图

区分组件与资源

![](https://i0.hdslb.com/bfs/album/e9fb85ddbffdbbe8e8f1fd972d52b3a6d3cd4c70.png)



#### 核心组件

- ETCD：分布式高性能键值数据库,存储整个集群的所有元数据

- ApiServer:  API服务器,集群资源访问控制入口,提供restAPI及安全访问控制

- Scheduler：调度器,负责把业务容器调度到最合适的Node节点

- Controller Manager：控制器管理,确保集群资源按照期望的方式运行
  - Replication Controller
  - Node controller
  - ResourceQuota Controller
  - Namespace Controller
  - ServiceAccount Controller
  - Tocken Controller
  - Service Controller
  - Endpoints Controller

- kubelet：运行在每运行在每个节点上的主要的“节点代理”个节点上的主要的“节点代理”
  - pod 管理：kubelet 定期从所监听的数据源获取节点上 pod/container 的期望状态（运行什么容器、运行的副本数量、网络或者存储如何配置等等），并调用对应的容器平台接口达到这个状态。
  - 容器健康检查：kubelet 创建了容器之后还要查看容器是否正常运行，如果容器运行出错，就要根据 pod 设置的重启策略进行处理.
  - 容器监控：kubelet 会监控所在节点的资源使用情况，并定时向 master 报告，资源使用数据都是通过 cAdvisor 获取的。知道整个集群所有节点的资源情况，对于 pod 的调度和正常运行至关重要
- kubectl: 命令行接口，用于对 Kubernetes 集群运行命令  https://kubernetes.io/zh/docs/reference/kubectl/ 
- CNI实现: 通用网络接口, 我们使用flannel来作为k8s集群的网络插件, 实现跨节点通信

#### 工作流程

![](https://i0.hdslb.com/bfs/album/5fb4494b23a8279807231a1fa7a32a1926df0bf8.png)

1. 用户准备一个资源文件（记录了业务应用的名称、镜像地址等信息），通过调用APIServer执行创建Pod
2. APIServer收到用户的Pod创建请求，将Pod信息写入到etcd中
3. 调度器通过list-watch的方式，发现有新的pod数据，但是这个pod还没有绑定到某一个节点中
4. 调度器通过调度算法，计算出最适合该pod运行的节点，并调用APIServer，把信息更新到etcd中
5. kubelet同样通过list-watch方式，发现有新的pod调度到本机的节点了，因此调用容器运行时，去根据pod的描述信息，拉取镜像，启动容器，同时生成事件信息
6. 同时，把容器的信息、事件及状态也通过APIServer写入到etcd中



#### 实践--集群安装

kubeadm  https://kubernetes.io/zh/docs/reference/setup-tools/kubeadm/kubeadm/ 

###### 核心组件

静态Pod的方式：

```powershell
## etcd、apiserver、controller-manager、kube-scheduler
$ kubectl -n kube-system get po
```

systemd服务方式：

```powershell
$ systemctl status kubelet
```

kubectl：二进制命令行工具

###### 理解集群资源

组件是为了支撑k8s平台的运行，安装好的软件。

资源是如何去使用k8s的能力的定义。比如，k8s可以使用Pod来管理业务应用，那么Pod就是k8s集群中的一类资源，集群中的所有资源可以提供如下方式查看：

```powershell
$ kubectl api-resources
```

如何理解namespace：

命名空间，集群内一个虚拟的概念，类似于资源池的概念，一个池子里可以有各种资源类型，绝大多数的资源都必须属于某一个namespace。集群初始化安装好之后，会默认有如下几个namespace：

```powershell
$ kubectl get namespaces
NAME                   STATUS   AGE
default                Active   84m
kube-node-lease        Active   84m
kube-public            Active   84m
kube-system            Active   84m
kubernetes-dashboard   Active   71m
```

- 所有NAMESPACED的资源，在创建的时候都需要指定namespace，**若不指定，默认会在default命名空间下**
- **相同namespace下的同类资源不可以重名**，不同类型的资源可以重名
- 不同namespace下的同类资源可以重名
- 通常在项目使用的时候，我们会创建带有业务含义的namespace来做逻辑上的整合

###### kubectl的使用

类似于docker，kubectl是命令行工具，用于与APIServer交互，内置了丰富的子命令，功能极其强大。 https://kubernetes.io/docs/reference/kubectl/overview/ 

```powershell
$ kubectl -h
$ kubectl get -h
$ kubectl create -h
$ kubectl create namespace -h
```

kubectl如何管理集群资源

```powershell
$ kubectl get po -v=7
```



#### 实践--使用k8s管理业务应用

##### 最小调度单元 Pod

docker调度的是容器，在k8s集群中，最小的调度单元是Pod（豆荚）

![](https://i0.hdslb.com/bfs/album/578d7bed4715bf5725b4b96dfbf9d2635fbb851f.png)

###### 为什么引入Pod

- 与容器引擎解耦

  Docker、Rkt。平台设计与引擎的具体的实现解耦

- 多容器共享网络|存储|进程 空间, 支持的业务场景更加灵活

###### 使用yaml格式定义Pod

*myblog/one-pod/pod.yaml*

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: myblog
  namespace: demo
  labels:
    component: myblog
spec:
  containers:
  - name: myblog
    image: 172.21.32.6:5000/myblog
    env:
    - name: MYSQL_HOST   #  指定root用户的用户名
      value: "127.0.0.1"
    - name: MYSQL_PASSWD
      value: "123456"
    ports:
    - containerPort: 8002
  - name: mysql
    image: 172.21.32.6:5000/mysql:5.7-utf8
    ports:
    - containerPort: 3306
    env:
    - name: MYSQL_ROOT_PASSWORD
      value: "123456"
    - name: MYSQL_DATABASE
      value: "myblog"
```

```json
{
	"apiVersion": "v1",		
	"kind": "Pod",
	"metadata": {
		"name": "myblog",
        "namespace": "demo",
        "labels": {
            "component": "myblog"
        }
	},
	"spec": {
		"containers": [
			{
				"name": "myblog",
				"image": "172.21.32.6:5000/myblog",
                "env": [
                    {
                        "name": "MYSQL_HOST",
                        "value": "127.0.0.1"
                    },
                    {
                        "name": "MYSQL_PASSWD",
                        "value": "123456"
                    }
                ],
				"ports": [
					{
						"containerPort": 8002
					}
				]
			},
    		{
    			"name": "mysql",
                ...
			}
		]
	}
}
```

##### apiVersion

| apiVersion | 含义                                                         |
| :--------- | :----------------------------------------------------------- |
| alpha      | 进入K8s功能的早期候选版本，可能包含Bug，最终不一定进入K8s    |
| beta       | 已经过测试的版本，最终会进入K8s，但功能、对象定义可能会发生变更。 |
| stable     | 可安全使用的稳定版本                                         |
| v1         | stable 版本之后的首个版本，包含了更多的核心对象              |
| apps/v1    | 使用最广泛的版本，像Deployment、ReplicaSets都已进入该版本    |

##### kind

资源类型与apiVersion对照表

| Kind                  | apiVersion                              |
| :-------------------- | :-------------------------------------- |
| ClusterRoleBinding    | rbac.authorization.k8s.io/v1            |
| ClusterRole           | rbac.authorization.k8s.io/v1            |
| ConfigMap             | v1                                      |
| CronJob               | batch/v1beta1                           |
| DaemonSet             | extensions/v1beta1                      |
| Node                  | v1                                      |
| Namespace             | v1                                      |
| Secret                | v1                                      |
| PersistentVolume      | v1                                      |
| PersistentVolumeClaim | v1                                      |
| Pod                   | v1                                      |
| Deployment            | v1、apps/v1、apps/v1beta1、apps/v1beta2 |
| Service               | v1                                      |
| Ingress               | extensions/v1beta1                      |
| ReplicaSet            | apps/v1、apps/v1beta2                   |
| Job                   | batch/v1                                |
| StatefulSet           | apps/v1、apps/v1beta1、apps/v1beta2     |

快速获得资源和版本

```powershell
$ kubectl explain pod
$ kubectl explain Pod.apiVersion
```

###### 创建和访问Pod

```powershell
## 创建namespace, namespace是逻辑上的资源池
$ kubectl create namespace demo

## 使用指定文件创建Pod
$ kubectl create -f demo-pod.yaml

## 查看pod，可以简写po
## 所有的操作都需要指定namespace，如果是在default命名空间下，则可以省略
$ kubectl -n demo get pods -o wide
NAME     READY   STATUS    RESTARTS   AGE    IP             NODE
myblog   2/2     Running   0          3m     10.244.1.146   k8s-slave1

## 使用Pod Ip访问服务,3306和8002
$ curl 10.244.1.146:8002/blog/index/

## 进入容器,执行初始化, 不必到对应的主机执行docker exec
$ kubectl -n demo exec -ti myblog -c myblog bash
/ # env
/ # python3 manage.py migrate
$ kubectl -n demo exec -ti myblog -c mysql bash
/ # mysql -p123456

## 再次访问服务,3306和8002
$ curl 10.244.1.146:8002/blog/index/

```

###### Infra容器

登录`k8s-slave1`节点

```powershell
$ docker ps -a |grep myblog  ## 发现有三个容器
## 其中包含mysql和myblog程序以及Infra容器
## 为了实现Pod内部的容器可以通过localhost通信，每个Pod都会启动Infra容器，然后Pod内部的其他容器的网络空间会共享该Infra容器的网络空间(Docker网络的container模式)，Infra容器只需要hang住网络空间，不需要额外的功能，因此资源消耗极低。

## 登录master节点，查看pod内部的容器ip均相同，为pod ip
$ kubectl -n demo exec -ti myblog -c myblog bash
/ # ifconfig
$ kubectl -n demo exec -ti myblog -c mysql bash
/ # ifconfig
```

pod容器命名: ```k8s_<container_name>_<pod_name>_<namespace>_<random_string>```



###### 查看pod详细信息

```powershell
## 查看pod调度节点及pod_ip
$ kubectl -n demo get pods -o wide
## 查看完整的yaml
$ kubectl -n demo get po myblog -o yaml
## 查看pod的明细信息及事件
$ kubectl -n demo describe pod myblog
```

###### Troubleshooting and Debugging

```powershell
#进入Pod内的容器
$ kubectl -n <namespace> exec <pod_name> -c <container_name> -ti /bin/sh

#查看Pod内容器日志,显示标准或者错误输出日志
$ kubectl -n <namespace> logs -f <pod_name> -c <container_name>
```



###### 更新服务版本

```powershell
$ kubectl apply -f demo-pod.yaml
```

###### 删除Pod服务

```powershell
#根据文件删除
$ kubectl delete -f demo-pod.yaml

#根据pod_name删除
$ kubectl -n <namespace> delete pod <pod_name>
```

###### Pod数据持久化

若删除了Pod，由于mysql的数据都在容器内部，会造成数据丢失，因此需要数据进行持久化。

- 定点使用hostpath挂载，nodeSelector定点

  `myblog/one-pod/pod-with-volume.yaml`

  ```yaml
  apiVersion: v1
  kind: Pod
  metadata:
    name: myblog
    namespace: demo
    labels:
      component: myblog
  spec:
    volumes: 
    - name: mysql-data
      hostPath: 
        path: /opt/mysql/data
    nodeSelector:   # 使用节点选择器将Pod调度到指定label的节点
      component: mysql
    containers:
    - name: myblog
      image: 172.21.32.6:5000/myblog
      env:
      - name: MYSQL_HOST   #  指定root用户的用户名
        value: "127.0.0.1"
      - name: MYSQL_PASSWD
        value: "123456"
      ports:
      - containerPort: 8002
    - name: mysql
      image: 172.21.32.6:5000/mysql:5.7-utf8
      ports:
      - containerPort: 3306
      env:
      - name: MYSQL_ROOT_PASSWORD
        value: "123456"
      - name: MYSQL_DATABASE
        value: "myblog"
      volumeMounts:
      - name: mysql-data
        mountPath: /var/lib/mysql
  ```

  保存文件为`pod-with-volume.yaml`，执行创建

  ```powershell
  ## 若存在旧的同名服务，先删除掉，后创建
  $ kubectl -n demo delete pod myblog
  ## 创建
  $ kubectl create -f pod-with-volume.yaml
  
  ## 此时pod状态Pending
  $ kubectl -n demo get po
  NAME     READY   STATUS    RESTARTS   AGE
  myblog   0/2     Pending   0          32s
  
  ## 查看原因，提示调度失败，因为节点不满足node selector
  $ kubectl -n demo describe po myblog
  Events:
    Type     Reason            Age                From               Message
    ----     ------            ----               ----               -------
    Warning  FailedScheduling  12s (x2 over 12s)  default-scheduler  0/3 nodes are available: 3 node(s) didn't match node selector.
    
  ## 为节点打标签
  $ kubectl label node k8s-slave1 component=mysql
  
  ## 再次查看，已经运行成功
  $ kubectl -n demo get po
  NAME     READY   STATUS    RESTARTS   AGE     IP             NODE
  myblog   2/2     Running   0          3m54s   10.244.1.150   k8s-slave1
  
  ## 到k8s-slave1节点，查看/opt/mysql/data
  $ ll /opt/mysql/data/
  total 188484
  -rw-r----- 1 polkitd input       56 Mar 29 09:20 auto.cnf
  -rw------- 1 polkitd input     1676 Mar 29 09:20 ca-key.pem
  -rw-r--r-- 1 polkitd input     1112 Mar 29 09:20 ca.pem
  drwxr-x--- 2 polkitd input     8192 Mar 29 09:20 sys
  ...
  
  ## 执行migrate，创建数据库表，然后删掉pod，再次创建后验证数据是否存在
  $ kubectl -n demo exec -ti myblog python3 manage.py migrate
  
  ## 访问服务，正常
  $ curl 10.244.1.150:8002/blog/index/ 
  
  ## 删除pod
  $ kubectl delete -f pod-with-volume.yaml
  
  ## 再次创建Pod
  $ kubectl create -f pod-with-volume.yaml
  
  ## 查看pod ip并访问服务
  $ kubectl -n demo get po -o wide
  NAME     READY   STATUS    RESTARTS   AGE   IP             NODE  
  myblog   2/2     Running   0          7s    10.244.1.151   k8s-slave1
  
  ## 未做migrate，服务正常
  $ curl 10.244.1.151:8002/blog/index/
  ```

- 使用PV+PVC连接分布式存储解决方案
  - ceph
  - glusterfs
  - nfs

###### 服务健康检查

检测容器服务是否健康的手段，若不健康，会根据设置的重启策略（restartPolicy）进行操作，两种检测机制可以分别单独设置，若不设置，默认认为Pod是健康的。

两种机制：

- LivenessProbe探针
  用于判断容器是否存活，即Pod是否为running状态，如果LivenessProbe探针探测到容器不健康，则kubelet将kill掉容器，并根据容器的重启策略是否重启，如果一个容器不包含LivenessProbe探针，则Kubelet认为容器的LivenessProbe探针的返回值永远成功。 
- ReadinessProbe探针
  用于判断容器是否正常提供服务，即容器的Ready是否为True，是否可以接收请求，如果ReadinessProbe探测失败，则容器的Ready将为False，控制器将此Pod的Endpoint从对应的service的Endpoint列表中移除，从此不再将任何请求调度此Pod上，直到下次探测成功。（剔除此pod不参与接收请求不会将流量转发给此Pod）。

三种类型：

- exec：通过执行命令来检查服务是否正常，回值为0则表示容器健康
- httpGet方式：通过发送http请求检查服务是否正常，返回200-399状态码则表明容器健康
- tcpSocket：通过容器的IP和Port执行TCP检查，如果能够建立TCP连接，则表明容器健康

示例：

完整文件路径 `myblog/one-pod/pod-with-healthcheck.yaml`

```yaml
  containers:
  - name: myblog
    image: 172.21.32.6:5000/myblog
    env:
    - name: MYSQL_HOST   #  指定root用户的用户名
      value: "127.0.0.1"
    - name: MYSQL_PASSWD
      value: "123456"
    ports:
    - containerPort: 8002
    livenessProbe:
      httpGet:
        path: /blog/index/
        port: 8002
        scheme: HTTP
      initialDelaySeconds: 10  # 容器启动后第一次执行探测是需要等待多少秒
      periodSeconds: 15 	# 执行探测的频率
      timeoutSeconds: 2		# 探测超时时间
    readinessProbe: 
      httpGet: 
        path: /blog/index/
        port: 8002
        scheme: HTTP
      initialDelaySeconds: 10 
      timeoutSeconds: 2
      periodSeconds: 15
```

- initialDelaySeconds：容器启动后第一次执行探测是需要等待多少秒。
- periodSeconds：执行探测的频率。默认是10秒，最小1秒。
- timeoutSeconds：探测超时时间。默认1秒，最小1秒。
- successThreshold：探测失败后，最少连续探测成功多少次才被认定为成功。默认是1。对于liveness必须是1，最小值是1。
- failureThreshold：探测成功后，最少连续探测失败多少次
  才被认定为失败。默认是3，最小值是1。

重启策略：

Pod的重启策略（RestartPolicy）应用于Pod内的所有容器，并且仅在Pod所处的Node上由kubelet进行判断和重启操作。当某个容器异常退出或者健康检查失败时，kubelet将根据RestartPolicy的设置来进行相应的操作。
 Pod的重启策略包括Always、OnFailure和Never，默认值为Always。

- Always：当容器失败时，由kubelet自动重启该容器；
- OnFailure：当容器终止运行且退出码不为0时，有kubelet自动重启该容器；
- Never：不论容器运行状态如何，kubelet都不会重启该容器。

###### 镜像拉取策略

```yaml
spec:
  containers:
  - name: myblog
    image: 172.21.32.6:5000/demo/myblog
    imagePullPolicy: IfNotPresent
```

设置镜像的拉取策略，默认为IfNotPresent

- Always，总是拉取镜像，即使本地有镜像也从仓库拉取
- IfNotPresent ，本地有则使用本地镜像，本地没有则去仓库拉取
- Never，只使用本地镜像，本地没有则报错

###### Pod资源限制

为了保证充分利用集群资源，且确保重要容器在运行周期内能够分配到足够的资源稳定运行，因此平台需要具备

Pod的资源限制的能力。 对于一个pod来说，资源最基础的2个的指标就是：CPU和内存。

Kubernetes提供了个采用requests和limits 两种类型参数对资源进行预分配和使用限制。

完整文件路径：`myblog/one-pod/pod-with-resourcelimits.yaml`

```yaml
...
  containers:
  - name: myblog
    image: 172.21.32.6:5000/myblog
    env:
    - name: MYSQL_HOST   #  指定root用户的用户名
      value: "127.0.0.1"
    - name: MYSQL_PASSWD
      value: "123456"
    ports:
    - containerPort: 8002
    resources:
      requests:
        memory: 100Mi
        cpu: 50m
      limits:
        memory: 500Mi
        cpu: 100m
...
```

requests：

- 容器使用的最小资源需求,作用于schedule阶段，作为容器调度时资源分配的判断依赖
- 只有当前节点上可分配的资源量 >= request 时才允许将容器调度到该节点
- request参数不限制容器的最大可使用资源
- requests.cpu被转成docker的--cpu-shares参数，与cgroup cpu.shares功能相同 (无论宿主机有多少个cpu或者内核，--cpu-shares选项都会按照比例分配cpu资源）
- requests.memory没有对应的docker参数，仅作为k8s调度依据

limits：

- 容器能使用资源的最大值
- 设置为0表示对使用的资源不做限制, 可无限的使用
- 当pod 内存超过limit时，会被oom
- 当cpu超过limit时，不会被kill，但是会限制不超过limit值
- limits.cpu会被转换成docker的–cpu-quota参数。与cgroup cpu.cfs_quota_us功能相同
- limits.memory会被转换成docker的–memory参数。用来限制容器使用的最大内存

 对于 CPU，我们知道计算机里 CPU 的资源是按`“时间片”`的方式来进行分配的，系统里的每一个操作都需要 CPU 的处理，所以，哪个任务要是申请的 CPU 时间片越多，那么它得到的 CPU 资源就越多。

然后还需要了解下 CGroup 里面对于 CPU 资源的单位换算：

```shell
1 CPU =  1000 millicpu（1 Core = 1000m）
```

 这里的 `m` 就是毫、毫核的意思，Kubernetes 集群中的每一个节点可以通过操作系统的命令来确认本节点的 CPU 内核数量，然后将这个数量乘以1000，得到的就是节点总 CPU 总毫数。比如一个节点有四核，那么该节点的 CPU 总毫量为 4000m。 

`docker run`命令和 CPU 限制相关的所有选项如下：

| 选项                  | 描述                                                    |
| --------------------- | ------------------------------------------------------- |
| `--cpuset-cpus=""`    | 允许使用的 CPU 集，值可以为 0-3,0,1                     |
| `-c`,`--cpu-shares=0` | CPU 共享权值（相对权重）                                |
| `cpu-period=0`        | 限制 CPU CFS 的周期，范围从 100ms~1s，即[1000, 1000000] |
| `--cpu-quota=0`       | 限制 CPU CFS 配额，必须不小于1ms，即 >= 1000，绝对限制  |

```shell
docker run -it --cpu-period=50000 --cpu-quota=25000 ubuntu:16.04 /bin/bash
```

将 CFS 调度的周期设为 50000，将容器在每个周期内的 CPU 配额设置为 25000，表示该容器每 50ms 可以得到 50% 的 CPU 运行时间。

> 注意：若内存使用超出限制，会引发系统的OOM机制，因CPU是可压缩资源，不会引发Pod退出或重建

###### yaml优化

目前完善后的yaml，`myblog/one-pod/pod-completed.yaml`

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: myblog
  namespace: demo
  labels:
    component: myblog
spec:
  volumes: 
  - name: mysql-data
    hostPath: 
      path: /opt/mysql/data
  nodeSelector:   # 使用节点选择器将Pod调度到指定label的节点
    component: mysql
  containers:
  - name: myblog
    image: 172.21.32.6:5000/myblog
    env:
    - name: MYSQL_HOST   #  指定root用户的用户名
      value: "127.0.0.1"
    - name: MYSQL_PASSWD
      value: "123456"
    ports:
    - containerPort: 8002
    resources:
      requests:
        memory: 100Mi
        cpu: 50m
      limits:
        memory: 500Mi
        cpu: 100m
    livenessProbe:
      httpGet:
        path: /blog/index/
        port: 8002
        scheme: HTTP
      initialDelaySeconds: 10  # 容器启动后第一次执行探测是需要等待多少秒
      periodSeconds: 15 	# 执行探测的频率
      timeoutSeconds: 2		# 探测超时时间
    readinessProbe: 
      httpGet: 
        path: /blog/index/
        port: 8002
        scheme: HTTP
      initialDelaySeconds: 10 
      timeoutSeconds: 2
      periodSeconds: 15
  - name: mysql
    image: 172.21.32.6:5000/mysql:5.7-utf8
    ports:
    - containerPort: 3306
    env:
    - name: MYSQL_ROOT_PASSWORD
      value: "123456"
    - name: MYSQL_DATABASE
      value: "myblog"
    resources:
      requests:
        memory: 100Mi
        cpu: 50m
      limits:
        memory: 500Mi
        cpu: 100m
    readinessProbe:
      tcpSocket:
        port: 3306
      initialDelaySeconds: 5
      periodSeconds: 10
    livenessProbe:
      tcpSocket:
        port: 3306
      initialDelaySeconds: 15
      periodSeconds: 20
    volumeMounts:
    - name: mysql-data
      mountPath: /var/lib/mysql
```

为什么要优化

- 考虑真实的使用场景，像数据库这类中间件，是作为公共资源，为多个项目提供服务，不适合和业务容器绑定在同一个Pod中，因为业务容器是经常变更的，而数据库不需要频繁迭代
- yaml的环境变量中存在敏感信息（账号、密码），存在安全隐患

解决问题一，需要拆分yaml

`myblog/two-pod/mysql.yaml`

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: mysql
  namespace: demo
  labels:
    component: mysql
spec:
  hostNetwork: true	# 声明pod的网络模式为host模式，效果通docker run --net=host
  volumes: 
  - name: mysql-data
    hostPath: 
      path: /opt/mysql/data
  nodeSelector:   # 使用节点选择器将Pod调度到指定label的节点
    component: mysql
  containers:
  - name: mysql
    image: 172.21.32.6:5000/mysql:5.7-utf8
    ports:
    - containerPort: 3306
    env:
    - name: MYSQL_ROOT_PASSWORD
      value: "123456"
    - name: MYSQL_DATABASE
      value: "myblog"
    resources:
      requests:
        memory: 100Mi
        cpu: 50m
      limits:
        memory: 500Mi
        cpu: 100m
    readinessProbe:
      tcpSocket:
        port: 3306
      initialDelaySeconds: 5
      periodSeconds: 10
    livenessProbe:
      tcpSocket:
        port: 3306
      initialDelaySeconds: 15
      periodSeconds: 20
    volumeMounts:
    - name: mysql-data
      mountPath: /var/lib/mysql
```

myblog.yaml

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: myblog
  namespace: demo
  labels:
    component: myblog
spec:
  containers:
  - name: myblog
    image: 172.21.32.6:5000/myblog
    imagePullPolicy: IfNotPresent
    env:
    - name: MYSQL_HOST   #  指定root用户的用户名
      value: "172.21.32.6"
    - name: MYSQL_PASSWD
      value: "123456"
    ports:
    - containerPort: 8002
    resources:
      requests:
        memory: 100Mi
        cpu: 50m
      limits:
        memory: 500Mi
        cpu: 100m
    livenessProbe:
      httpGet:
        path: /blog/index/
        port: 8002
        scheme: HTTP
      initialDelaySeconds: 10  # 容器启动后第一次执行探测是需要等待多少秒
      periodSeconds: 15 	# 执行探测的频率
      timeoutSeconds: 2		# 探测超时时间
    readinessProbe: 
      httpGet: 
        path: /blog/index/
        port: 8002
        scheme: HTTP
      initialDelaySeconds: 10 
      timeoutSeconds: 2
      periodSeconds: 15
```

创建测试

```powershell
## 先删除旧pod
$ kubectl -n demo delete po myblog

## 分别创建mysql和myblog
$ kubectl create -f mysql.yaml
$ kubectl create -f myblog.yaml

## 查看pod，注意mysqlIP为宿主机IP，因为网络模式为host
$ kubectl -n demo get po -o wide 
NAME     READY   STATUS    RESTARTS   AGE   IP                NODE
myblog   1/1     Running   0          41s   10.244.1.152      k8s-slave1
mysql    1/1     Running   0          52s   192.168.136.131   k8s-slave1

## 访问myblog服务正常
$ curl 10.244.1.152:8002/blog/index/
```



解决问题二，环境变量中敏感信息带来的安全隐患

为什么要统一管理环境变量

- 环境变量中有很多敏感的信息，比如账号密码，直接暴漏在yaml文件中存在安全性问题
- 团队内部一般存在多个项目，这些项目直接存在配置相同环境变量的情况，因此可以统一维护管理
- 对于开发、测试、生产环境，由于配置均不同，每套环境部署的时候都要修改yaml，带来额外的开销

k8s提供两类资源，configMap和Secret，可以用来实现业务配置的统一管理， 允许将配置文件与镜像文件分离，以使容器化的应用程序具有可移植性 。

![](https://i0.hdslb.com/bfs/album/1e64a761a3ef97056c795904715d28d18d6a7366.png)

##### configMap

- configMap，通常用来管理应用的配置文件或者环境变量，`myblog/two-pod/configmap.yaml`

  ```yaml
  apiVersion: v1
  kind: ConfigMap
  metadata:
    name: myblog
    namespace: demo
  data:
    MYSQL_HOST: "172.21.32.6"
    MYSQL_PORT: "3306"
  ```

- Secret，管理敏感类的信息，默认会base64编码存储，有三种类型

  - Service Account ：用来访问Kubernetes API，由Kubernetes自动创建，并且会自动挂载到Pod的/run/secrets/kubernetes.io/serviceaccount目录中；创建ServiceAccount后，Pod中指定serviceAccount后，自动创建该ServiceAccount对应的secret；
  - Opaque ： base64编码格式的Secret，用来存储密码、密钥等；
  - kubernetes.io/dockerconfigjson ：用来存储私有docker registry的认证信息。

  `myblog/two-pod/secret.yaml`
  
  ```yaml
  apiVersion: v1
  kind: Secret
  metadata:
    name: myblog
    namespace: demo
  type: Opaque
  data:
    MYSQL_USER: cm9vdA==		#注意加-n参数， echo -n root|base64
  MYSQL_PASSWD: MTIzNDU2
  ```

  创建并查看：
  
  ```powershell
  $ kubectl create -f secret.yaml
  $ kubectl -n demo get secret
  ```
  
  如果不习惯这种方式，可以通过如下方式：
  
  ```powershell
  $ cat secret.txt
  MYSQL_USER=root
  MYSQL_PASSWD=123456
  $ kubectl -n demo create secret generic myblog --from-env-file=secret.txt 
  ```

修改后的mysql的yaml，资源路径：`myblog/two-pod/mysql-with-config.yaml`

```yaml
...
spec:
  containers:
  - name: mysql
    image: 172.21.32.6:5000/mysql:5.7-utf8
    env:
    - name: MYSQL_USER
      valueFrom:
        secretKeyRef:
          name: myblog
          key: MYSQL_USER
    - name: MYSQL_PASSWD
      valueFrom:
        secretKeyRef:
          name: myblog
          key: MYSQL_PASSWD
    - name: MYSQL_DATABASE
      value: "myblog"
...
```

修改后的myblog的yaml，资源路径：`myblog/two-pod/myblog-with-config.yaml`

```yaml
spec:
  containers:
  - name: myblog
    image: 172.21.32.6:5000/myblog
    imagePullPolicy: IfNotPresent
    env:
    - name: MYSQL_HOST
      valueFrom:
        configMapKeyRef:
          name: myblog
          key: MYSQL_HOST
    - name: MYSQL_PORT
      valueFrom:
        configMapKeyRef:
          name: myblog
          key: MYSQL_PORT
    - name: MYSQL_USER
      valueFrom:
        secretKeyRef:
          name: myblog
          key: MYSQL_USER
    - name: MYSQL_PASSWD
      valueFrom:
        secretKeyRef:
          name: myblog
          key: MYSQL_PASSWD
```

在部署不同的环境时，pod的yaml无须再变化，只需要在每套环境中维护一套ConfigMap和Secret即可。但是注意configmap和secret不能跨namespace使用，且更新后，pod内的env不会自动更新，重建后方可更新。



###### 如何编写资源yaml

1. 拿来主义，从机器中已有的资源中拿

   ```powershell
   $ kubectl -n kube-system get po,deployment,ds
   ```

2. 学会在官网查找， https://kubernetes.io/docs/home/ 

3. 从kubernetes-api文档中查找， https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.16/#pod-v1-core 

4. kubectl explain 查看具体字段含义



###### pod状态与生命周期

Pod的状态如下表所示：

| 状态值            | 描述                                                         |
| ----------------- | ------------------------------------------------------------ |
| Pending           | API Server已经创建该Pod，等待调度器调度或者                  |
| ContainerCreating | 镜像正在创建                                                 |
| Running           | Pod内容器均已创建，且至少有一个容器处于运行状态、正在启动状态或正在重启状态 |
| Succeeded         | Pod内所有容器均已成功执行退出，且不再重启                    |
| Failed            | Pod内所有容器均已退出，但至少有一个容器退出为失败状态        |
| CrashLoopBackOff  | Pod内有容器启动失败，比如配置文件丢失导致主进程启动失败      |
| Unknown           | 由于某种原因无法获取该Pod的状态，可能由于网络通信不畅导致    |

生命周期示意图：

![](https://i0.hdslb.com/bfs/album/645d9ce6b36cc1f102d286da6a78ee6c3167bfaa.jpg)

启动和关闭示意：

![](https://i0.hdslb.com/bfs/album/82299a117becf7529123657548874026562159a6.jpg)

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: demo-start-stop
  namespace: demo
  labels:
    component: demo-start-stop
spec:
  initContainers:
  - name: init
    image: busybox
    command: ['sh', '-c', 'echo $(date +%s): INIT >> /loap/timing']
    volumeMounts:
    - mountPath: /loap
      name: timing
  containers:
  - name: main
    image: busybox
    command: ['sh', '-c', 'echo $(date +%s): START >> /loap/timing;
sleep 10; echo $(date +%s): END >> /loap/timing;']
    volumeMounts:
    - mountPath: /loap 
      name: timing
    livenessProbe:
      exec:
        command: ['sh', '-c', 'echo $(date +%s): LIVENESS >> /loap/timing']
    readinessProbe:
      exec:
        command: ['sh', '-c', 'echo $(date +%s): READINESS >> /loap/timing']
    lifecycle:
      postStart:
        exec:
          command: ['sh', '-c', 'echo $(date +%s): POST-START >> /loap/timing']
      preStop:
        exec:
          command: ['sh', '-c', 'echo $(date +%s): PRE-STOP >> /loap/timing']
  volumes:
  - name: timing
    hostPath:
      path: /tmp/loap
```

创建pod测试：

```powershell
$ kubectl create -f demo-pod-start.yaml

## 查看demo状态
$ kubectl -n demo get po -o wide -w

## 查看调度节点的/tmp/loap/timing
$ cat /tmp/loap/timing
1585424708: INIT
1585424746: START
1585424746: POST-START
1585424754: READINESS
1585424756: LIVENESS
1585424756: END
```

>  须主动杀掉 Pod 才会触发 `pre-stop hook`，如果是 Pod 自己 Down 掉，则不会执行 `pre-stop hook` 

###### 小结

1. 实现k8s平台与特定的容器运行时解耦，提供更加灵活的业务部署方式，引入了Pod概念
2. k8s使用yaml格式定义资源文件，yaml中Map与List的语法，与json做类比
3. 通过kubectl create | get | exec | logs | delete 等操作k8s资源，必须指定namespace
4. 每启动一个Pod，为了实现网络空间共享，会先创建Infra容器，并把其他容器网络加入该容器
5. 通过livenessProbe和readinessProbe实现Pod的存活性和就绪健康检查
6. 通过requests和limit分别限定容器初始资源申请与最高上限资源申请
7. 通过Pod IP访问具体的Pod服务，实现是



##### Pod控制器

只使用Pod, 将会面临如下需求:

1. 业务应用启动多个副本
2. Pod重建后IP会变化，外部如何访问Pod服务
3. 运行业务Pod的某个节点挂了，可以自动帮我把Pod转移到集群中的可用节点启动起来
4. 我的业务应用功能是收集节点监控数据,需要把Pod运行在k8集群的各个节点上

###### Workload (工作负载)

控制器又称工作负载是用于实现管理pod的中间层，确保pod资源符合预期的状态，pod的资源出现故障时，会尝试 进行重启，当根据重启策略无效，则会重新新建pod的资源。 

![](https://i0.hdslb.com/bfs/album/2e995c695e38f0d5ac696ada66c81ded2d35a83e.png)

- ReplicaSet: 代用户创建指定数量的pod副本数量，确保pod副本数量符合预期状态，并且支持滚动式自动扩容和缩容功能
- Deployment：工作在ReplicaSet之上，用于管理无状态应用，目前来说最好的控制器。支持滚动更新和回滚功能，还提供声明式配置
- DaemonSet：用于确保集群中的每一个节点只运行特定的pod副本，通常用于实现系统级后台任务。比如ELK服务
- Job：只要完成就立即退出，不需要重启或重建
- Cronjob：周期性任务控制，不需要持续后台运行
- StatefulSet：管理有状态应用



###### Deployment

`myblog/deployment/deploy-mysql.yaml`

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: mysql
  namespace: demo
spec:
  replicas: 1	#指定Pod副本数
  selector:		#指定Pod的选择器
    matchLabels:
      app: mysql # 必须和下面的对应
  template:
    metadata:
      labels:	#给Pod打label
        app: mysql
    spec:
      hostNetwork: true	# 声明pod的网络模式为host模式，效果通docker run --net=host
      volumes: 
      - name: mysql-data
        hostPath: 
          path: /opt/mysql/data
      nodeSelector:   # 使用节点选择器将Pod调度到指定label的节点
        component: mysql
      containers:
      - name: mysql
        image: 172.21.32.6:5000/mysql:5.7
        args:
        - "--character-set-server=utf8"     #
        - "--collation-server=utf8_general_ci"    #  指定字符编码
        ports:
        - containerPort: 3306
        env:
        - name: MYSQL_USER
          valueFrom:
            secretKeyRef:
              name: myblog
              key: MYSQL_USER
        - name: MYSQL_PASSWD
          valueFrom:
            secretKeyRef:
              name: myblog
              key: MYSQL_PASSWD
        - name: MYSQL_DATABASE
          value: "myblog"
        resources:
          requests:
            memory: 100Mi
            cpu: 50m
          limits:
            memory: 500Mi
            cpu: 100m
        readinessProbe:
          tcpSocket:
            port: 3306
          initialDelaySeconds: 5
          periodSeconds: 10
        livenessProbe:
          tcpSocket:
            port: 3306
          initialDelaySeconds: 15
          periodSeconds: 20
        volumeMounts:
        - name: mysql-data
          mountPath: /var/lib/mysql
```

deploy-myblog.yaml:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: myblog
  namespace: demo
spec:
  replicas: 1	#指定Pod副本数
  selector:		#指定Pod的选择器
    matchLabels:
      app: myblog
  template:
    metadata:
      labels:	#给Pod打label
        app: myblog
    spec:
      containers:
      - name: myblog
        image: 172.21.32.6:5000/myblog
        imagePullPolicy: IfNotPresent
        env:
        - name: MYSQL_HOST
          valueFrom:
            configMapKeyRef:
              name: myblog
              key: MYSQL_HOST
        - name: MYSQL_PORT
          valueFrom:
            configMapKeyRef:
              name: myblog
              key: MYSQL_PORT
        - name: MYSQL_USER
          valueFrom:
            secretKeyRef:
              name: myblog
              key: MYSQL_USER
        - name: MYSQL_PASSWD
          valueFrom:
            secretKeyRef:
              name: myblog
              key: MYSQL_PASSWD
        ports:
        - containerPort: 8002
        resources:
          requests:
            memory: 100Mi
            cpu: 50m
          limits:
            memory: 500Mi
            cpu: 100m
        livenessProbe:
          httpGet:
            path: /blog/index/
            port: 8002
            scheme: HTTP
          initialDelaySeconds: 10  # 容器启动后第一次执行探测是需要等待多少秒
          periodSeconds: 15 	# 执行探测的频率
          timeoutSeconds: 2		# 探测超时时间
        readinessProbe: 
          httpGet: 
            path: /blog/index/
            port: 8002
            scheme: HTTP
          initialDelaySeconds: 10 
          timeoutSeconds: 2
          periodSeconds: 15
```



###### 创建Deployment

```powershell
$ kubectl create -f deploy.yaml
```

###### 查看Deployment

```powershell
# kubectl api-resources
$ kubectl -n demo get deploy
NAME     READY   UP-TO-DATE   AVAILABLE   AGE
myblog   1/1     1            1           2m22s
mysql    1/1     1            1           2d11h

  * `NAME` 列出了集群中 Deployments 的名称。
  * `READY`显示当前正在运行的副本数/期望的副本数。
  * `UP-TO-DATE`显示已更新以实现期望状态的副本数。
  * `AVAILABLE`显示应用程序可供用户使用的副本数。
  * `AGE` 显示应用程序运行的时间量。

# 查看pod
$ kubectl -n demo get po
NAME                      READY   STATUS    RESTARTS   AGE
myblog-7c96c9f76b-qbbg7   1/1     Running   0          109s
mysql-85f4f65f99-w6jkj    1/1     Running   0          2m28s
```



###### 副本保障机制

controller实时检测pod状态，并保障副本数一直处于期望的值。

```powershell
## 删除pod，观察pod状态变化
$ kubectl -n demo delete pod myblog-7c96c9f76b-qbbg7

# 观察pod
$ kubectl get pods -o wide

## 设置两个副本, 或者通过kubectl -n demo edit deploy myblog的方式，最好通过修改文件，然后apply的方式，这样yaml文件可以保持同步
$ kubectl -n demo scale deploy myblog --replicas=2
deployment.extensions/myblog scaled

# 观察pod
$ kubectl get pods -o wide
NAME                      READY   STATUS    RESTARTS   AGE
myblog-7c96c9f76b-qbbg7   1/1     Running   0          11m
myblog-7c96c9f76b-s6brm   1/1     Running   0          55s
mysql-85f4f65f99-w6jkj    1/1     Running   0          11m
```

###### Pod驱逐策略

 K8S 有个特色功能叫 pod eviction，它在某些场景下如节点 NotReady，或者资源不足时，把 pod 驱逐至其它节点，这也是出于业务保护的角度去考虑的。

1. Kube-controller-manager: 周期性检查所有节点状态，当节点处于 NotReady 状态超过一段时间后，驱逐该节点上所有 pod。停掉kubelet
   - `pod-eviction-timeout`：NotReady 状态节点超过该时间后，执行驱逐，默认 5 min

2. Kubelet: 周期性检查本节点资源，当资源不足时，按照优先级驱逐部分 pod
   - `memory.available`：节点可用内存
   - `nodefs.available`：节点根盘可用存储空间
   - `nodefs.inodesFree`：节点inodes可用数量
   - `imagefs.available`：镜像存储盘的可用空间
   - `imagefs.inodesFree`：镜像存储盘的inodes可用数量

###### 服务更新

修改dockerfile，重新打tag模拟服务更新。

更新方式：

- 修改yaml文件，使用`kubectl -n demo apply -f deploy-myblog.yaml`来应用更新
- `kubectl -n demo edit deploy myblog`在线更新
- `kubectl set image deploy myblog myblog=172.21.32.6:5000/myblog:v2 --record` 

修改文件测试：

```powershell
$ vi mybolg/blog/template/index.html

$ docker build . -t 172.21.32.6:5000/myblog:v2 -f Dockerfile_optimized
$ docker push 172.21.32.6:5000/myblog:v2
```



###### 更新策略

```yaml
...
spec:
  replicas: 2	#指定Pod副本数
  selector:		#指定Pod的选择器
    matchLabels:
      app: myblog
  strategy:
    rollingUpdate:
      maxSurge: 25%
      maxUnavailable: 25%
    type: RollingUpdate		#指定更新方式为滚动更新，默认策略，通过get deploy yaml查看
    ...
```

![](https://i0.hdslb.com/bfs/album/187ba47adc0aa0ffa99b9d7367365bf2d98076d5.png)

策略控制：

- maxSurge：最大激增数, 指更新过程中, 最多可以比replicas预先设定值多出的pod数量, 可以为固定值或百分比,默认为desired Pods数的25%。计算时向上取整(比如3.4，取4)，更新过程中最多会有replicas + maxSurge个pod
- maxUnavailable： 指更新过程中, 最多有几个pod处于无法服务状态 , 可以为固定值或百分比，默认为desired Pods数的25%。计算时向下取整(比如3.6，取3)

*在Deployment rollout时，需要保证Available(Ready) Pods数不低于 desired pods number - maxUnavailable; 保证所有的非异常状态Pods数不多于 desired pods number + maxSurge*。

以myblog为例，使用默认的策略，更新过程:

1. maxSurge 25%，2个实例，向上取整，则maxSurge为1，意味着最多可以有2+1=3个Pod，那么此时会新创建1个ReplicaSet，RS-new，把副本数置为1，此时呢，副本控制器就去创建这个新的Pod
2. 同时，maxUnavailable是25%，副本数2*25%，向下取整，则为0，意味着，滚动更新的过程中，不能有少于2个可用的Pod，因此，旧的Replica（RS-old）会先保持不动，等RS-new管理的Pod状态Ready后，此时已经有3个Ready状态的Pod了，那么由于只要保证有2个可用的Pod即可，因此，RS-old的副本数会有2个变成1个，此时，会删掉一个旧的Pod
3. 删掉旧的Pod的时候，由于总的Pod数量又变成2个了，因此，距离最大的3个还有1个Pod可以创建，所以，RS-new把管理的副本数由1改成2，此时又会创建1个新的Pod，等RS-new管理了2个Pod都ready后，那么就可以把RS-old的副本数由1置为0了，这样就完成了滚动更新

```powershell
#查看滚动更新事件
$ kubectl -n demo describe deploy myblog
...
Events:
  Type    Reason             Age   From                   Message
  ----    ------             ----  ----                   -------
  Normal  ScalingReplicaSet  11s   deployment-controller  Scaled up replica set myblog-6cf56fc848 to 1
  Normal  ScalingReplicaSet  11s   deployment-controller  Scaled down replica set myblog-6fdcf98f9 to 1
  Normal  ScalingReplicaSet  11s   deployment-controller  Scaled up replica set myblog-6cf56fc848 to 2
  Normal  ScalingReplicaSet  6s    deployment-controller  Scaled down replica set myblog-6fdcf98f9 to 0
$ kubectl get rs
NAME                     DESIRED   CURRENT   READY   AGE
myblog-6cf56fc848   2         2         2       16h
myblog-6fdcf98f9    0         0         0       16h
```

###### 服务回滚

通过滚动升级的策略可以平滑的升级Deployment，若升级出现问题，需要最快且最好的方式回退到上一次能够提供正常工作的版本。为此K8S提供了回滚机制。

**revision**：更新应用时，K8S都会记录当前的版本号，即为revision，当升级出现问题时，可通过回滚到某个特定的revision，默认配置下，K8S只会保留最近的几个revision，可以通过Deployment配置文件中的spec.revisionHistoryLimit属性增加revision数量，默认是10。

查看当前：

```powershell
$ kubectl -n demo rollout history deploy myblog ##CHANGE-CAUSE为空
$ kubectl delete -f deploy-myblog.yaml    ## 方便演示到具体效果，删掉已有deployment
```

记录回滚：

```powershell
$ kubectl create -f deploy-myblog.yaml --record

$ kubectl -n demo set image deploy myblog myblog=172.21.32.6:5000/myblog:v2 --record=true
```

查看deployment更新历史：

```powershell
$ kubectl -n demo rollout history deploy myblog
deployment.extensions/myblog
REVISION  CHANGE-CAUSE
1         kubectl create --filename=deploy-myblog.yaml --record=true
2         kubectl set image deploy myblog myblog=172.21.32.6:5000/demo/myblog:v1 --record=true
```

回滚到具体的REVISION:

```powershell
$ kubectl -n demo rollout undo deploy myblog --to-revision=1
deployment.extensions/myblog rolled back

# 访问应用测试
```



##### Kubernetes调度

######  为何要控制Pod应该如何调度 

- 集群中有些机器的配置高（SSD，更好的内存等），我们希望核心的服务（比如说数据库）运行在上面
- 某两个服务的网络传输很频繁，我们希望它们最好在同一台机器上 
- ......

###### NodeSelector

 `label`是`kubernetes`中一个非常重要的概念，用户可以非常灵活的利用 label 来管理集群中的资源，POD 的调度可以根据节点的 label 进行特定的部署。 

查看节点的label：

```powershell
$ kubectl get nodes --show-labels
```

为节点打label：

```powershell
$ kubectl label node k8s-master disktype=ssd
```

 当 node 被打上了相关标签后，在调度的时候就可以使用这些标签了，只需要在spec 字段中添加`nodeSelector`字段，里面是我们需要被调度的节点的 label。 

```yaml
...
spec:
  hostNetwork: true	# 声明pod的网络模式为host模式，效果通docker run --net=host
  volumes: 
  - name: mysql-data
    hostPath: 
      path: /opt/mysql/data
  nodeSelector:   # 使用节点选择器将Pod调度到指定label的节点
    component: mysql
  containers:
  - name: mysql
  	image: 172.21.32.6:5000/demo/mysql:5.7
...
```

###### nodeAffinity

节点亲和性 ， 比上面的`nodeSelector`更加灵活，它可以进行一些简单的逻辑组合，不只是简单的相等匹配 。分为两种，软策略和硬策略。

 preferredDuringSchedulingIgnoredDuringExecution：软策略，如果你没有满足调度要求的节点的话，Pod就会忽略这条规则，继续完成调度过程，说白了就是满足条件最好了，没有满足就忽略掉的策略。

 requiredDuringSchedulingIgnoredDuringExecution ： 硬策略，如果没有满足条件的节点的话，就不断重试直到满足条件为止，简单说就是你必须满足我的要求，不然我就不会调度Pod。

```yaml
#要求 Pod 不能运行在128和132两个节点上，如果有个节点满足disktype=ssd的话就优先调度到这个节点上
...
spec:
      containers:
      - name: demo
        image: 172.21.32.6:5000/demo/myblog
        ports:
        - containerPort: 8002
      affinity:
          nodeAffinity:
            requiredDuringSchedulingIgnoredDuringExecution:
                nodeSelectorTerms:
                - matchExpressions:
                    - key: kubernetes.io/hostname
                      operator: NotIn
                      values:
                        - 192.168.136.128
                        - 192.168.136.132
            preferredDuringSchedulingIgnoredDuringExecution:
                - weight: 1
                  preference:
                    matchExpressions:
                    - key: disktype
                      operator: In
                      values:
                        - ssd
                        - sas
...
```

这里的匹配逻辑是 label 的值在某个列表中，现在`Kubernetes`提供的操作符有下面的几种：

- In：label 的值在某个列表中
- NotIn：label 的值不在某个列表中
- Gt：label 的值大于某个值
- Lt：label 的值小于某个值
- Exists：某个 label 存在
- DoesNotExist：某个 label 不存在

*如果nodeSelectorTerms下面有多个选项的话，满足任何一个条件就可以了；如果matchExpressions有多个选项的话，则必须同时满足这些条件才能正常调度 Pod*

###### 污点（Taints）与容忍（tolerations）

对于`nodeAffinity`无论是硬策略还是软策略方式，都是调度 Pod 到预期节点上，而`Taints`恰好与之相反，如果一个节点标记为 Taints ，除非 Pod 也被标识为可以容忍污点节点，否则该 Taints 节点不会被调度Pod。

Taints(污点)是Node的一个属性，设置了Taints(污点)后，因为有了污点，所以Kubernetes是不会将Pod调度到这个Node上的。于是Kubernetes就给Pod设置了个属性Tolerations(容忍)，只要Pod能够容忍Node上的污点，那么Kubernetes就会忽略Node上的污点，就能够(不是必须)把Pod调度过去。

比如用户希望把 Master 节点保留给 Kubernetes 系统组件使用，或者把一组具有特殊资源预留给某些 Pod，则污点就很有用了，Pod 不会再被调度到 taint 标记过的节点。taint 标记节点举例如下：

设置污点：

```powershell
$ kubectl taint node [node_name] key=value:[effect]   
      其中[effect] 可取值： [ NoSchedule | PreferNoSchedule | NoExecute ]
       NoSchedule：一定不能被调度。
       PreferNoSchedule：尽量不要调度。
       NoExecute：不仅不会调度，还会驱逐Node上已有的Pod。
  示例：kubectl taint node k8s-master smoke=true:NoSchedule
```

去除污点：

```powershell
去除指定key及其effect：
     kubectl taint nodes [node_name] key:[effect]-    #这里的key不用指定value
                
 去除指定key所有的effect: 
     kubectl taint nodes node_name key-
 
 示例：
     kubectl taint node k8s-master smoke=true:NoSchedule
     kubectl taint node k8s-master smoke:NoExecute-
     kubectl taint node k8s-master smoke-
```

污点演示：

```powershell
## 给k8s-slave1打上污点，smoke=true:NoSchedule
$ kubectl taint node k8s-slave1 smoke=true:NoSchedule
$ kubectl taint node k8s-slave2 drunk=true:NoSchedule


## 扩容myblog的Pod，观察新Pod的调度情况
$ kuebctl -n demo scale deploy myblog --replicas=3
$ kubectl -n demo get po -w    ## pending
```



Pod容忍污点示例：`myblog/deployment/deploy-myblog-taint.yaml`

```powershell
...
spec:
      containers:
      - name: demo
        image: 172.21.32.6:5000/demo/myblog
      tolerations: #设置容忍性
      - key: "smoke" 
        operator: "Equal"  #如果操作符为Exists，那么value属性可省略,不指定operator，默认为Equal
        value: "true"
        effect: "NoSchedule"
      - key: "drunk" 
        operator: "Equal"  #如果操作符为Exists，那么value属性可省略,不指定operator，默认为Equal
        value: "true"
        effect: "NoSchedule"
	  #意思是这个Pod要容忍的有污点的Node的key是smoke Equal true,效果是NoSchedule，
      #tolerations属性下各值必须使用引号，容忍的值都是设置Node的taints时给的值。
```

```powershell
$ kubectl apply -f deploy-myblog-taint.yaml
```

```powershell
spec:
      containers:
      - name: demo
        image: 172.21.32.6:5000/demo/myblog
      tolerations:
        - operator: "Exists"
```



##### Kubernetes服务访问之`Service`

通过以前的学习，我们已经能够通过Deployment来创建一组Pod来提供具有高可用性的服务。虽然每个Pod都会分配一个单独的Pod IP，然而却存在如下两个问题：

- Pod IP仅仅是集群内可见的虚拟IP，外部无法访问。
- Pod IP会随着Pod的销毁而消失，当ReplicaSet对Pod进行动态伸缩时，Pod IP可能随时随地都会变化，这样对于我们访问这个服务带来了难度。

> docker使用的是docker0

###### Service 负载均衡之Cluster IP

service是一组pod的服务抽象，相当于一组pod的LB，负责将请求分发给对应的pod。service会为这个LB提供一个IP，一般称为cluster IP 。使用Service对象，通过selector进行标签选择，找到对应的Pod:

`myblog/deployment/svc-myblog.yaml`

```yaml
apiVersion: v1
kind: Service
metadata:
  name: myblog
  namespace: demo
spec:
  ports:
  - port: 80 # 监听的端口
    protocol: TCP
    targetPort: 8002 # pod的端口
  selector:
    app: myblog
  type: ClusterIP
```

操作演示：

```powershell
## 别名
$ alias kd='kubectl -n demo'

## 创建服务
$ kd create -f svc-myblog.yaml
$ kd get po --show-labels
NAME                      READY   STATUS    RESTARTS   AGE    LABELS
myblog-5c97d79cdb-jn7km   1/1     Running   0          6m5s   app=myblog
mysql-85f4f65f99-w6jkj    1/1     Running   0          176m   app=mysql

$ kd get svc
NAME     TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)   AGE
myblog   ClusterIP   10.99.174.93   <none>        80/TCP    7m50s

$ kd describe svc myblog
Name:              myblog
Namespace:         demo
Labels:            <none>
Annotations:       <none>
Selector:          app=myblog
Type:              ClusterIP
IP:                10.99.174.93
Port:              <unset>  80/TCP
TargetPort:        8002/TCP
Endpoints:         10.244.0.68:8002
Session Affinity:  None
Events:            <none>

## 扩容myblog服务
$ kd scale deploy myblog --replicas=2
deployment.extensions/myblog scaled

## 再次查看
$ kd describe svc myblog
Name:              myblog
Namespace:         demo
Labels:            <none>
Annotations:       <none>
Selector:          app=myblog
Type:              ClusterIP
IP:                10.99.174.93
Port:              <unset>  80/TCP
TargetPort:        8002/TCP
Endpoints:         10.244.0.68:8002,10.244.1.158:8002
Session Affinity:  None
Events:            <none>
```

Service与Pod如何关联:

service对象创建的同时，会创建同名的endpoints对象，若服务设置了readinessProbe, 当readinessProbe检测失败时，endpoints列表中会剔除掉对应的pod_ip，这样流量就不会分发到健康检测失败的Pod中

```powershell
$ kd get endpoints myblog
NAME     ENDPOINTS                            AGE
myblog   10.244.0.68:8002,10.244.1.158:8002   7m
```

Service Cluster-IP如何访问:

```powershell
$ kd get svc myblog
NAME   TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)   AGE
myblog   ClusterIP   10.99.174.93   <none>        80/TCP    13m
$ curl 10.99.174.93/blog/index/
```

为mysql服务创建service：

```yaml
apiVersion: v1
kind: Service
metadata:
  name: mysql
  namespace: demo
spec:
  ports:
  - port: 3306
    protocol: TCP
    targetPort: 3306
  selector:
    app: mysql
  type: ClusterIP
```

访问mysql：

```powershell
$ kd get svc mysql
mysql    ClusterIP   10.108.214.84   <none>        3306/TCP   3s
$ curl 10.108.214.84:3306
```

目前使用hostNetwork部署，通过宿主机ip+port访问，弊端：

- 服务使用hostNetwork，使得宿主机的端口大量暴漏，存在安全隐患
- 容易引发端口冲突

服务均属于k8s集群，尽可能使用k8s的网络访问，因此可以对目前myblog访问mysql的方式做改造：

- 为mysql创建一个固定clusterIp的Service，把clusterIp配置在myblog的环境变量中
- 利用集群服务发现的能力，组件之间通过service name来访问

###### 服务发现

在k8s集群中，组件之间可以通过定义的Service名称实现通信。

演示服务发现：

```powershell
## 演示思路：在myblog的容器中直接通过service名称访问服务，观察是否可以访问通

# 先查看服务
$ kd get svc
NAME     TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)    AGE
myblog   ClusterIP   10.99.174.93    <none>        80/TCP     59m
mysql    ClusterIP   10.108.214.84   <none>        3306/TCP   35m

# 进入myblog容器
$ kd exec -ti myblog-5c97d79cdb-j485f bash
[root@myblog-5c97d79cdb-j485f myblog]# curl mysql:3306
5.7.29 )→  (mysql_native_password ot packets out of order
[root@myblog-5c97d79cdb-j485f myblog]# curl myblog/blog/index/
我的博客列表

```

虽然podip和clusterip都不固定，但是service name是固定的，而且具有完全的跨集群可移植性，因此组件之间调用的同时，完全可以通过service name去通信，这样避免了大量的ip维护成本，使得服务的yaml模板更加简单。因此可以对mysql和myblog的部署进行优化改造：

1. mysql可以去掉hostNetwork部署，使得服务只暴漏在k8s集群内部网络
2. configMap中数据库地址可以换成Service名称，这样跨环境的时候，配置内容基本上可以保持不用变化

修改deploy-mysql.yaml

```yaml
    spec:
      hostNetwork: true	# 去掉此行
      volumes: 
      - name: mysql-data
        hostPath: 
          path: /opt/mysql/data

```

修改configmap.yaml

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: myblog
  namespace: demo
data:
  MYSQL_HOST: "mysql"	# 此处替换为mysql
  MYSQL_PORT: "3306"
```

应用修改：

```powershell
$ kubectl apply -f configmap.yaml
$ kubectl apply -f deploy-mysql.yaml

## 重建pod
$ kubectl -n demo delete po mysql-7f747644b8-6npzn

#去掉taint
$ kubectl taint node k8s-slave1 smoke-
$ kubectl taint node k8s-slave2 drunk-

## myblog不用动，会自动因健康检测不过而重启
```

服务发现实现：

 `CoreDNS`是一个`Go`语言实现的链式插件`DNS服务端`，是CNCF成员，是一个高性能、易扩展的`DNS服务端`。 

```powershell
$ kubectl -n kube-system get po -o wide|grep dns
coredns-d4475785-2w4hk             1/1     Running   0          4d22h   10.244.0.64       
coredns-d4475785-s49hq             1/1     Running   0          4d22h   10.244.0.65

# 查看myblog的pod解析配置
$ kubectl -n demo exec -ti myblog-5c97d79cdb-j485f bash
[root@myblog-5c97d79cdb-j485f myblog]# cat /etc/resolv.conf
nameserver 10.96.0.10
search demo.svc.cluster.local svc.cluster.local cluster.local
options ndots:5

## 10.96.0.10 从哪来
$ kubectl -n kube-system get svc
NAME       TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)         AGE
kube-dns   ClusterIP   10.96.0.10   <none>        53/UDP,53/TCP   51d

## 启动pod的时候，会把kube-dns服务的cluster-ip地址注入到pod的resolve解析配置中，同时添加对应的namespace的search域。 因此跨namespace通过service name访问的话，需要添加对应的namespace名称，
service_name.namespace_name
$ kubectl get svc
NAME         TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
kubernetes   ClusterIP   10.96.0.1    <none>        443/TCP   26h
```



###### Service负载均衡之NodePort

cluster-ip为虚拟地址，只能在k8s集群内部进行访问，集群外部如果访问内部服务，实现方式之一为使用NodePort方式。NodePort会默认在 30000-32767 ，不指定的会随机使用其中一个。

`myblog/deployment/svc-myblog-nodeport.yaml`

```powershell
apiVersion: v1
kind: Service
metadata:
  name: myblog-np
  namespace: demo
spec:
  ports:
  - port: 80
    protocol: TCP
    targetPort: 8002
  selector:
    app: myblog
  type: NodePort
```

查看并访问服务：

```powershell
$ kd create -f svc-myblog-nodeport.yaml
service/myblog-np created
$ kd get svc
NAME        TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)        AGE
myblog      ClusterIP   10.99.174.93     <none>        80/TCP         102m
myblog-np   NodePort    10.105.228.101   <none>        80:30647/TCP   4s
mysql       ClusterIP   10.108.214.84    <none>        3306/TCP       77m

#集群内每个节点的NodePort端口都会进行监听
$ curl 192.168.136.128:30647/blog/index/
我的博客列表
$ curl 192.168.136.131:30647/blog/index/
我的博客列表
## 浏览器访问
```

思考：

1. NodePort的端口监听如何转发到对应的Pod服务？

2. CLUSTER-IP为虚拟IP，集群内如何通过虚拟IP访问到具体的Pod服务？

###### kube-proxy

 运行在每个节点上，监听 API Server 中服务对象的变化，再通过创建流量路由规则来实现网络的转发。[参照]( https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies )

有三种模式：

- User space, 让 Kube-Proxy 在用户空间监听一个端口，所有的 Service 都转发到这个端口，然后 Kube-Proxy 在内部应用层对其进行转发 ， 所有报文都走一遍用户态，性能不高，k8s v1.2版本后废弃。
- Iptables， 当前默认模式，完全由 IPtables 来实现， 通过各个node节点上的iptables规则来实现service的负载均衡，但是随着service数量的增大，iptables模式由于线性查找匹配、全量更新等特点，其性能会显著下降。 
- IPVS， 与iptables同样基于Netfilter，但是采用的hash表，因此当service数量达到一定规模时，hash查表的速度优势就会显现出来，从而提高service的服务性能。 k8s 1.8版本开始引入，1.11版本开始稳定，需要开启宿主机的ipvs模块。

 IPtables模式示意图：

![](https://i0.hdslb.com/bfs/album/1a223ec0abe90ee5e174559cf42ebb907cae72e9.jpg)

```powershell
$ iptables-save |grep -v myblog-np|grep  "demo/myblog"
-A KUBE-SERVICES ! -s 10.244.0.0/16 -d 10.99.174.93/32 -p tcp -m comment --comment "demo/myblog: cluster IP" -m tcp --dport 80 -j KUBE-MARK-MASQ
-A KUBE-SERVICES -d 10.99.174.93/32 -p tcp -m comment --comment "demo/myblog: cluster IP" -m tcp --dport 80 -j KUBE-SVC-WQNGJ7YFZKCTKPZK

$ iptables-save |grep KUBE-SVC-WQNGJ7YFZKCTKPZK
-A KUBE-SVC-WQNGJ7YFZKCTKPZK -m statistic --mode random --probability 0.50000000000 -j KUBE-SEP-GB5GNOM5CZH7ICXZ
-A KUBE-SVC-WQNGJ7YFZKCTKPZK -j KUBE-SEP-7GWC3FN2JI5KLE47

$  iptables-save |grep KUBE-SEP-GB5GNOM5CZH7ICXZ
-A KUBE-SEP-GB5GNOM5CZH7ICXZ -p tcp -m tcp -j DNAT --to-destination 10.244.1.158:8002

$ iptables-save |grep KUBE-SEP-7GWC3FN2JI5KLE47
-A KUBE-SEP-7GWC3FN2JI5KLE47 -p tcp -m tcp -j DNAT --to-destination 10.244.1.159:8002
```

##### Kubernetes服务访问之Ingress

对于Kubernetes的Service，无论是Cluster-Ip和NodePort均是四层的负载，集群内的服务如何实现七层的负载均衡，这就需要借助于Ingress，Ingress控制器的实现方式有很多，比如nginx, Contour, Haproxy, trafik, Istio，我们以nginx的实现为例做演示。

Ingress-nginx是7层的负载均衡器 ，负责统一管理外部对k8s cluster中service的请求。主要包含：

- ingress-nginx-controller：根据用户编写的ingress规则（创建的ingress的yaml文件），动态的去更改nginx服务的配置文件，并且reload重载使其生效（是自动化的，通过lua脚本来实现）；
- ingress资源对象：将Nginx的配置抽象成一个Ingress对象，每添加一个新的Service资源对象只需写一个新的Ingress规则的yaml文件即可（或修改已存在的ingress规则的yaml文件）

###### 示意图：

![](https://i0.hdslb.com/bfs/album/56bf40ff1c6e7520244c33fabae2983f077b17ff.webp)

###### 实现逻辑

1）ingress controller通过和kubernetes api交互，动态的去感知集群中ingress规则变化
2）然后读取ingress规则(规则就是写明了哪个域名对应哪个service)，按照自定义的规则，生成一段nginx配置
3）再写到nginx-ingress-controller的pod里，这个Ingress controller的pod里运行着一个Nginx服务，控制器把生成的nginx配置写入/etc/nginx.conf文件中
4）然后reload一下使配置生效。以此达到域名分别配置和动态更新的问题。

###### 安装

  [官方文档](https://github.com/kubernetes/ingress-nginx/blob/master/docs/deploy/index.md)

```powershell
$ wget https://raw.githubusercontent.com/kubernetes/ingress-nginx/nginx-0.30.0/deploy/static/mandatory.yaml
## 或者使用myblog/deployment/ingress/mandatory.yaml
## 修改部署节点
$ grep -n5 nodeSelector mandatory.yaml
212-    spec:
213-      hostNetwork: true #添加为host模式
214-      # wait up to five minutes for the drain of connections
215-      terminationGracePeriodSeconds: 300
216-      serviceAccountName: nginx-ingress-serviceaccount
217:      nodeSelector:
218-        ingress: "true"		#替换此处，来决定将ingress部署在哪些机器
219-      containers:
220-        - name: nginx-ingress-controller
221-          image: quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.30.0
222-          args:
```

使用示例：`myblog/deployment/ingress.yaml`

```yaml
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: myblog
  namespace: demo
spec:
  rules:
  - host: myblog.devops.cn
    http:
      paths:
      - path: /
        backend:
          serviceName: myblog
          servicePort: 80
```

ingress-nginx动态生成upstream配置：

```yaml
...
                server_name myblog.devops.cn ;

                listen 80  ;
                listen [::]:80  ;
                listen 443  ssl http2 ;
                listen [::]:443  ssl http2 ;

                set $proxy_upstream_name "-";

                ssl_certificate_by_lua_block {
                        certificate.call()
                }

                location / {

                        set $namespace      "demo";
                        set $ingress_name   "myblog";
 ...
```

###### 访问

域名解析服务，将 `myblog.devops.cn`解析到ingress的地址上。ingress是支持多副本的，高可用的情况下，生产的配置是使用lb服务（内网F5设备，公网elb、slb、clb，解析到各ingress的机器，如何域名指向lb地址）

本机，添加如下hosts记录来演示效果。

```json
192.168.136.128 myblog.devops.cn
```

然后，访问 http://myblog.devops.cn/blog/index/

HTTPS访问：

```powershell
#自签名证书
$ openssl req -x509 -nodes -days 2920 -newkey rsa:2048 -keyout tls.key -out tls.crt -subj "/CN=*.devops.cn/O=ingress-nginx"

# 证书信息保存到secret对象中，ingress-nginx会读取secret对象解析出证书加载到nginx配置中
$ kubectl -n demo create secret tls https-secret --key tls.key --cert tls.crt
secret/https-secret created
```

修改yaml

```yaml
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: myblog-tls
  namespace: demo
spec:
  rules:
  - host: myblog.devops.cn
    http:
      paths:
      - path: /
        backend:
          serviceName: myblog
          servicePort: 80
  tls:
  - hosts:
    - myblog.devops.cn
    secretName: https-secret
```

然后，访问 https://myblog.devops.cn/blog/index/



##### Kubernetes认证与授权

###### APIService安全控制

![](https://i0.hdslb.com/bfs/album/0fe544921de3c3c235319b62d48c352c4569dca0.jpg)



- Authentication：身份认证

  1. 这个环节它面对的输入是整个`http request`，负责对来自client的请求进行身份校验，支持的方法包括:

  - client证书验证（https双向验证）

  - `basic auth`
  - 普通token
  - `jwt token`(用于serviceaccount)

  2. APIServer启动时，可以指定一种Authentication方法，也可以指定多种方法。如果指定了多种方法，那么APIServer将会逐个使用这些方法对客户端请求进行验证， 只要请求数据通过其中一种方法的验证，APIServer就会认为Authentication成功；

  3. 使用kubeadm引导启动的k8s集群的apiserver初始配置中，默认支持`client证书`验证和`serviceaccount`两种身份验证方式。 证书认证通过设置`--client-ca-file`根证书以及`--tls-cert-file`和`--tls-private-key-file`来开启。

  4. 在这个环节，apiserver会通过client证书或 `http header`中的字段(比如serviceaccount的`jwt token`)来识别出请求的`用户身份`，包括”user”、”group”等，这些信息将在后面的`authorization`环节用到。

- Authorization：鉴权，你可以访问哪些资源

  1. 这个环节面对的输入是`http request context`中的各种属性，包括：`user`、`group`、`request path`（比如：`/api/v1`、`/healthz`、`/version`等）、 `request verb`(比如：`get`、`list`、`create`等)。

  2. APIServer会将这些属性值与事先配置好的访问策略(`access policy`）相比较。APIServer支持多种`authorization mode`，包括`Node、RBAC、Webhook`等。

  3. APIServer启动时，可以指定一种`authorization mode`，也可以指定多种`authorization mode`，如果是后者，只要Request通过了其中一种mode的授权， 那么该环节的最终结果就是授权成功。在较新版本kubeadm引导启动的k8s集群的apiserver初始配置中，`authorization-mode`的默认配置是`”Node,RBAC”`。

- Admission Control：[准入控制](http://docs.kubernetes.org.cn/144.html)，一个控制链(层层关卡)，偏集群安全控制、管理方面。为什么说是安全相关的机制？
  - 以NamespaceLifecycle为例， 该插件确保处于Termination状态的Namespace不再接收新的对象创建请求，并拒绝请求不存在的Namespace。该插件还可以防止删除系统保留的Namespace:default，kube-system，kube-public。 
  - NodeRestriction， 此插件限制kubelet修改Node和Pod对象，这样的kubelets只允许修改绑定到Node的Pod API对象，以后版本可能会增加额外的限制 。

为什么我们执行命令kubectl命令，可以直接管理k8s集群资源？

###### kubectl的认证授权

kubectl的日志调试级别：

| 信息 | 描述                                                         |
| :--- | :----------------------------------------------------------- |
| v=0  | 通常，这对操作者来说总是可见的。                             |
| v=1  | 当您不想要很详细的输出时，这个是一个合理的默认日志级别。     |
| v=2  | 有关服务和重要日志消息的有用稳定状态信息，这些信息可能与系统中的重大更改相关。这是大多数系统推荐的默认日志级别。 |
| v=3  | 关于更改的扩展信息。                                         |
| v=4  | 调试级别信息。                                               |
| v=6  | 显示请求资源。                                               |
| v=7  | 显示 HTTP 请求头。                                           |
| v=8  | 显示 HTTP 请求内容。                                         |
| v=9  | 显示 HTTP 请求内容，并且不截断内容。                         |

```powershell
$ kubectl get nodes -v=7
I0329 20:20:08.633065    3979 loader.go:359] Config loaded from file /root/.kube/config
I0329 20:20:08.633797    3979 round_trippers.go:416] GET https://192.168.136.128:6443/api/v1/nodes?limit=500
```

`kubeadm init`启动完master节点后，会默认输出类似下面的提示内容：

```bash
... ...
Your Kubernetes master has initialized successfully!

To start using your cluster, you need to run the following as a regular user:
  mkdir -p $HOME/.kube
  sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
  sudo chown $(id -u):$(id -g) $HOME/.kube/config
... ...
```

这些信息是在告知我们如何配置`kubeconfig`文件。按照上述命令配置后，master节点上的`kubectl`就可以直接使用`$HOME/.kube/config`的信息访问`k8s cluster`了。 并且，通过这种配置方式，`kubectl`也拥有了整个集群的管理员(root)权限。

很多K8s初学者在这里都会有疑问：

- 当`kubectl`使用这种`kubeconfig`方式访问集群时，`Kubernetes`的`kube-apiserver`是如何对来自`kubectl`的访问进行身份验证(`authentication`)和授权(`authorization`)的呢？
- 为什么来自`kubectl`的请求拥有最高的管理员权限呢？

查看`/root/.kube/config`文件：

前面提到过apiserver的authentication支持通过`tls client certificate、basic auth、token`等方式对客户端发起的请求进行身份校验， 从kubeconfig信息来看，kubectl显然在请求中使用了`tls client certificate`的方式，即客户端的证书。 

证书base64解码：

```powershell
$ echo xxxxxxxxxxxxxx |base64 -d > kubectl.crt
```

 说明在认证阶段，`apiserver`会首先使用`--client-ca-file`配置的CA证书去验证kubectl提供的证书的有效性,基本的方式 ：

```powershell
$  openssl verify -CAfile /etc/kubernetes/pki/ca.crt kubectl.crt
kubectl.crt: OK
```

除了认证身份，还会取出必要的信息供授权阶段使用，文本形式查看证书内容：

```powershell
$ openssl x509 -in kubectl.crt -text
Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number: 4736260165981664452 (0x41ba9386f52b74c4)
    Signature Algorithm: sha256WithRSAEncryption
        Issuer: CN=kubernetes
        Validity
            Not Before: Feb 10 07:33:39 2020 GMT
            Not After : Feb  9 07:33:40 2021 GMT
        Subject: O=system:masters, CN=kubernetes-admin
        ...
```

认证通过后，提取出签发证书时指定的CN(Common Name),`kubernetes-admin`，作为请求的用户名 (User Name), 从证书中提取O(Organization)字段作为请求用户所属的组 (Group)，`group = system:masters`，然后传递给后面的授权模块。

 kubeadm在init初始引导集群启动过程中，创建了许多`default`的`role、clusterrole、rolebinding`和`clusterrolebinding`， 在k8s有关RBAC的官方文档中，我们看到下面一些`default clusterrole`列表: 

![](C:\Users\liyongxin\Desktop\wework\老男孩\训练营\images\kubeadm-default-clusterrole-list.png)

其中第一个cluster-admin这个cluster role binding绑定了system:masters group，这和authentication环节传递过来的身份信息不谋而合。 沿着system:masters group对应的cluster-admin clusterrolebinding“追查”下去，真相就会浮出水面。

我们查看一下这一binding：

```yaml
$ kubectl describe clusterrolebinding cluster-admin
Name:         cluster-admin
Labels:       kubernetes.io/bootstrapping=rbac-defaults
Annotations:  rbac.authorization.kubernetes.io/autoupdate: true
Role:
  Kind:  ClusterRole
  Name:  cluster-admin
Subjects:
  Kind   Name            Namespace
  ----   ----            ---------
  Group  system:masters
```

我们看到在kube-system名字空间中，一个名为cluster-admin的clusterrolebinding将cluster-admin cluster role与system:masters Group绑定到了一起， 赋予了所有归属于system:masters Group中用户cluster-admin角色所拥有的权限。

我们再来查看一下cluster-admin这个role的具体权限信息：

```powershell
$ kubectl describe clusterrole cluster-admin
Name:         cluster-admin
Labels:       kubernetes.io/bootstrapping=rbac-defaults
Annotations:  rbac.authorization.kubernetes.io/autoupdate: true
PolicyRule:
  Resources  Non-Resource URLs  Resource Names  Verbs
  ---------  -----------------  --------------  -----
  *.*        []                 []              [*]
             [*]                []              [*]
```

非资源类，如查看集群健康状态。

![](https://i0.hdslb.com/bfs/album/fec70d55e35319fec9f34187322d5ad58974296f.png)

###### RBAC

 Role-Based Access Control，基于角色的访问控制， apiserver启动参数添加--authorization-mode=RBAC 来启用RBAC认证模式，kubeadm安装的集群默认已开启。[官方介绍](https://kubernetes.io/docs/reference/access-authn-authz/rbac/)

查看开启：

```powershell
# master节点查看apiserver进程
$ ps aux |grep apiserver
```

RBAC模式引入了4个资源：

- Role，角色

   一个Role只能授权访问单个namespace 

  ```yaml
  ## 示例定义一个名为pod-reader的角色，该角色具有读取default这个命名空间下的pods的权限
  kind: Role
  apiVersion: rbac.authorization.k8s.io/v1
  metadata:
    namespace: default
    name: pod-reader
  rules:
  - apiGroups: [""] # "" indicates the core API group
    resources: ["pods"]
    verbs: ["get", "watch", "list"]
    
  ## apiGroups: "","apps", "autoscaling", "batch", kubectl api-versions
  ## resources: "services", "pods","deployments"... kubectl api-resources
  ## verbs: "get", "list", "watch", "create", "update", "patch", "delete", "exec"
  ```

- ClusterRole

   一个ClusterRole能够授予和Role一样的权限，但是它是集群范围内的。 

  ```yaml
  ## 定义一个集群角色，名为secret-reader，该角色可以读取所有的namespace中的secret资源
  kind: ClusterRole
  apiVersion: rbac.authorization.k8s.io/v1
  metadata:
    # "namespace" omitted since ClusterRoles are not namespaced
    name: secret-reader
  rules:
  - apiGroups: [""]
    resources: ["secrets"]
    verbs: ["get", "watch", "list"]
  ```

- Rolebinding

  将role中定义的权限分配给用户和用户组。RoleBinding包含主题（users,groups,或service accounts）和授予角色的引用。对于namespace内的授权使用RoleBinding，集群范围内使用ClusterRoleBinding。

  ```yaml
  ## 定义一个角色绑定，将pod-reader这个role的权限授予给jane这个User，使得jane可以在读取default这个命名空间下的所有的pod数据
  kind: RoleBinding
  apiVersion: rbac.authorization.k8s.io/v1
  metadata:
    name: read-pods
    namespace: default
  subjects:
  - kind: User   #这里可以是User,Group,ServiceAccount
    name: jane 
    apiGroup: rbac.authorization.k8s.io
  roleRef:
    kind: Role #这里可以是Role或者ClusterRole,若是ClusterRole，则权限也仅限于rolebinding的内部
    name: pod-reader # match the name of the Role or ClusterRole you wish to bind to
    apiGroup: rbac.authorization.k8s.io
  ```

  *注意：rolebinding既可以绑定role，也可以绑定clusterrole，当绑定clusterrole的时候，subject的权限也会被限定于rolebinding定义的namespace内部，若想跨namespace，需要使用clusterrolebinding*

  ```yaml
  ## 定义一个角色绑定，将dave这个用户和secret-reader这个集群角色绑定，虽然secret-reader是集群角色，但是因为是使用rolebinding绑定的，因此dave的权限也会被限制在development这个命名空间内
  apiVersion: rbac.authorization.k8s.io/v1
  # This role binding allows "dave" to read secrets in the "development" namespace.
  # You need to already have a ClusterRole named "secret-reader".
  kind: RoleBinding
  metadata:
    name: read-secrets
    #
    # The namespace of the RoleBinding determines where the permissions are granted.
    # This only grants permissions within the "development" namespace.
    namespace: development
  subjects:
  - kind: User
    name: dave # Name is case sensitive
    apiGroup: rbac.authorization.k8s.io
  - kind: ServiceAccount
    name: dave # Name is case sensitive
    namespace: demo
  roleRef:
    kind: ClusterRole
    name: secret-reader
    apiGroup: rbac.authorization.k8s.io
  ```

  考虑一个场景：  如果集群中有多个namespace分配给不同的管理员，每个namespace的权限是一样的，就可以只定义一个clusterrole，然后通过rolebinding将不同的namespace绑定到管理员身上，否则就需要每个namespace定义一个Role，然后做一次rolebinding。

- ClusterRolebingding

  允许跨namespace进行授权

  ```yaml
  apiVersion: rbac.authorization.k8s.io/v1
  # This cluster role binding allows anyone in the "manager" group to read secrets in any namespace.
  kind: ClusterRoleBinding
  metadata:
    name: read-secrets-global
  subjects:
  - kind: Group
    name: manager # Name is case sensitive
    apiGroup: rbac.authorization.k8s.io
  roleRef:
    kind: ClusterRole
    name: secret-reader
    apiGroup: rbac.authorization.k8s.io
  ```

  

![](https://i0.hdslb.com/bfs/album/63c6e88ad760067041f4edceb81dd2ad087a0f5d.png)

###### kubelet的认证授权

查看kubelet进程

```powershell
$ systemctl status kubelet
● kubelet.service - kubelet: The Kubernetes Node Agent
   Loaded: loaded (/etc/systemd/system/kubelet.service; enabled; vendor preset: disabled)
  Drop-In: /etc/systemd/system/kubelet.service.d
           └─10-kubeadm.conf
   Active: active (running) since Wed 2020-04-01 02:34:13 CST; 1 day 14h ago
     Docs: https://kubernetes.io/docs/
 Main PID: 851 (kubelet)
    Tasks: 21
   Memory: 127.1M
   CGroup: /system.slice/kubelet.service
           └─851 /usr/bin/kubelet --bootstrap-kubeconfig=/etc/kubernetes/bootstrap-kubelet.conf --kubeconfig=/etc/kubernetes/kubelet.conf
```

查看`/etc/kubernetes/kubelet.conf`，解析证书：

```powershell
$ echo xxxxx |base64 -d >kubelet.crt
$ openssl x509 -in kubelet.crt -text
Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number: 9059794385454520113 (0x7dbadafe23185731)
    Signature Algorithm: sha256WithRSAEncryption
        Issuer: CN=kubernetes
        Validity
            Not Before: Feb 10 07:33:39 2020 GMT
            Not After : Feb  9 07:33:40 2021 GMT
        Subject: O=system:nodes, CN=system:node:master-1
```

得到我们期望的内容：

```bash
Subject: O=system:nodes, CN=system:node:k8s-master
```

我们知道，k8s会把O作为Group来进行请求，因此如果有权限绑定给这个组，肯定在clusterrolebinding的定义中可以找得到。因此尝试去找一下绑定了system:nodes组的clusterrolebinding

```powershell
$ kubectl get clusterrolebinding|awk 'NR>1{print $1}'|xargs kubectl get clusterrolebinding -oyaml|grep -n10 system:nodes
98-  roleRef:
99-    apiGroup: rbac.authorization.k8s.io
100-    kind: ClusterRole
101-    name: system:certificates.k8s.io:certificatesigningrequests:selfnodeclient
102-  subjects:
103-  - apiGroup: rbac.authorization.k8s.io
104-    kind: Group
105:    name: system:nodes
106-- apiVersion: rbac.authorization.k8s.io/v1
107-  kind: ClusterRoleBinding
108-  metadata:
109-    creationTimestamp: "2020-02-10T07:34:02Z"
110-    name: kubeadm:node-proxier
111-    resourceVersion: "213"
112-    selfLink: /apis/rbac.authorization.k8s.io/v1/clusterrolebindings/kubeadm%3Anode-proxier

$ kubectl describe clusterrole system:certificates.k8s.io:certificatesigningrequests:selfnodeclient
Name:         system:certificates.k8s.io:certificatesigningrequests:selfnodeclient
Labels:       kubernetes.io/bootstrapping=rbac-defaults
Annotations:  rbac.authorization.kubernetes.io/autoupdate: true
PolicyRule:
  Resources                                                      Non-Resource URLs  Resource Names  Verbs
  ---------                                                      -----------------  --------------  -----
  certificatesigningrequests.certificates.k8s.io/selfnodeclient  []                 []              [create]
```

 结局有点意外，除了`system:certificates.k8s.io:certificatesigningrequests:selfnodeclient`外，没有找到system相关的rolebindings，显然和我们的理解不一样。 尝试去找[资料](https://kubernetes.io/docs/reference/access-authn-authz/rbac/)，发现了这么一段 :

| Default ClusterRole            | Default ClusterRoleBinding          | Description                                                  |
| :----------------------------- | :---------------------------------- | :----------------------------------------------------------- |
| system:kube-scheduler          | system:kube-scheduler user          | Allows access to the resources required by the [scheduler](https://kubernetes.io/docs/reference/generated/kube-scheduler/)component. |
| system:volume-scheduler        | system:kube-scheduler user          | Allows access to the volume resources required by the kube-scheduler component. |
| system:kube-controller-manager | system:kube-controller-manager user | Allows access to the resources required by the [controller manager](https://kubernetes.io/docs/reference/command-line-tools-reference/kube-controller-manager/) component. The permissions required by individual controllers are detailed in the [controller roles](https://kubernetes.io/docs/reference/access-authn-authz/rbac/#controller-roles). |
| system:node                    | None                                | Allows access to resources required by the kubelet, **including read access to all secrets, and write access to all pod status objects**. You should use the [Node authorizer](https://kubernetes.io/docs/reference/access-authn-authz/node/) and [NodeRestriction admission plugin](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/#noderestriction) instead of the `system:node` role, and allow granting API access to kubelets based on the Pods scheduled to run on them. The `system:node` role only exists for compatibility with Kubernetes clusters upgraded from versions prior to v1.8. |
| system:node-proxier            | system:kube-proxy user              | Allows access to the resources required by the [kube-proxy](https://kubernetes.io/docs/reference/command-line-tools-reference/kube-proxy/)component. |

大致意思是说：之前会定义system:node这个角色，目的是为了kubelet可以访问到必要的资源，包括所有secret的读权限及更新pod状态的写权限。如果1.8版本后，是建议使用 [Node authorizer](https://kubernetes.io/docs/reference/access-authn-authz/node/) and [NodeRestriction admission plugin](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/#noderestriction) 来代替这个角色的。

我们目前使用1.16，查看一下授权策略：

```powershell
$ ps axu|grep apiserver
kube-apiserver --authorization-mode=Node,RBAC  --enable-admission-plugins=NodeRestriction
```

查看一下官网对Node authorizer的介绍：

*Node authorization is a special-purpose authorization mode that specifically authorizes API requests made by kubelets.*

*In future releases, the node authorizer may add or remove permissions to ensure kubelets have the minimal set of permissions required to operate correctly.*

*In order to be authorized by the Node authorizer, kubelets must use a credential that identifies them as being in the `system:nodes` group, with a username of `system:node:<nodeName>`*



###### Service Account

前面说，认证可以通过证书，也可以通过使用ServiceAccount（服务账户）的方式来做认证。大多数时候，我们在基于k8s做二次开发时都是选择通过serviceaccount的方式。我们之前访问dashboard的时候，是如何做的？

```yaml
## 新建一个名为admin的serviceaccount，并且把名为cluster-admin的这个集群角色的权限授予新建的serviceaccount
apiVersion: v1
kind: ServiceAccount
metadata:
  name: admin
  namespace: kubernetes-dashboard
---
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1beta1
metadata:
  name: admin
  annotations:
    rbac.authorization.kubernetes.io/autoupdate: "true"
roleRef:
  kind: ClusterRole
  name: cluster-admin
  apiGroup: rbac.authorization.k8s.io
subjects:
- kind: ServiceAccount
  name: admin
  namespace: kubernetes-dashboard
```

我们查看一下：

```powershell
$ kubectl -n kubernetes-dashboard get sa admin -o yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  creationTimestamp: "2020-04-01T11:59:21Z"
  name: admin
  namespace: kubernetes-dashboard
  resourceVersion: "1988878"
  selfLink: /api/v1/namespaces/kubernetes-dashboard/serviceaccounts/admin
  uid: 639ecc3e-74d9-11ea-a59b-000c29dfd73f
secrets:
- name: admin-token-lfsrf
```

注意到serviceaccount上默认绑定了一个名为admin-token-lfsrf的secret，我们查看一下secret

```powershell
$ kubectl -n kubernetes-dashboard describe secret admin-token-lfsrf
Name:         admin-token-lfsrf
Namespace:    kubernetes-dashboard
Labels:       <none>
Annotations:  kubernetes.io/service-account.name: admin
              kubernetes.io/service-account.uid: 639ecc3e-74d9-11ea-a59b-000c29dfd73f

Type:  kubernetes.io/service-account-token
Data
====
ca.crt:     1025 bytes
namespace:  4 bytes
token:      eyJhbGciOiJSUzI1NiIsImtpZCI6IiJ9.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJkZW1vIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZWNyZXQubmFtZSI6ImFkbWluLXRva2VuLWxmc3JmIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQubmFtZSI6ImFkbWluIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQudWlkIjoiNjM5ZWNjM2UtNzRkOS0xMWVhLWE1OWItMDAwYzI5ZGZkNzNmIiwic3ViIjoic3lzdGVtOnNlcnZpY2VhY2NvdW50OmRlbW86YWRtaW4ifQ.ffGCU4L5LxTsMx3NcNixpjT6nLBi-pmstb4I-W61nLOzNaMmYSEIwAaugKMzNR-2VwM14WbuG04dOeO67niJeP6n8-ALkl-vineoYCsUjrzJ09qpM3TNUPatHFqyjcqJ87h4VKZEqk2qCCmLxB6AGbEHpVFkoge40vHs56cIymFGZLe53JZkhu3pwYuS4jpXytV30Ad-HwmQDUu_Xqcifni6tDYPCfKz2CZlcOfwqHeGIHJjDGVBKqhEeo8PhStoofBU6Y4OjObP7HGuTY-Foo4QindNnpp0QU6vSb7kiOiQ4twpayybH8PTf73dtdFt46UF6mGjskWgevgolvmO8A
```

开发的时候如何去调用k8s的api:

1. curl演示

```powershell
$ curl -k  -H "Authorization: Bearer eyJhbGciOiJSUzI1NiIsImtpZCI6IiJ9.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJkZW1vIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZWNyZXQubmFtZSI6ImFkbWluLXRva2VuLWxmc3JmIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQubmFtZSI6ImFkbWluIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQudWlkIjoiNjM5ZWNjM2UtNzRkOS0xMWVhLWE1OWItMDAwYzI5ZGZkNzNmIiwic3ViIjoic3lzdGVtOnNlcnZpY2VhY2NvdW50OmRlbW86YWRtaW4ifQ.ffGCU4L5LxTsMx3NcNixpjT6nLBi-pmstb4I-W61nLOzNaMmYSEIwAaugKMzNR-2VwM14WbuG04dOeO67niJeP6n8-ALkl-vineoYCsUjrzJ09qpM3TNUPatHFqyjcqJ87h4VKZEqk2qCCmLxB6AGbEHpVFkoge40vHs56cIymFGZLe53JZkhu3pwYuS4jpXytV30Ad-HwmQDUu_Xqcifni6tDYPCfKz2CZlcOfwqHeGIHJjDGVBKqhEeo8PhStoofBU6Y4OjObP7HGuTY-Foo4QindNnpp0QU6vSb7kiOiQ4twpayybH8PTf73dtdFt46UF6mGjskWgevgolvmO8A" https://62.234.214.206:6443/api/v1/namespaces/demo/pods?limit=500
```

2. postman

#### 查看etcd数据

拷贝etcdctl命令行工具：

```powershell
$ docker exec -ti  etcd_container which etcdctl
$ docker cp etcd_container:/usr/local/bin/etcdctl /usr/bin/etcdctl
```

查看所有key值：

```powershell
$  ETCDCTL_API=3 etcdctl --endpoints=https://[127.0.0.1]:2379 --cacert=/etc/kubernetes/pki/etcd/ca.crt --cert=/etc/kubernetes/pki/etcd/healthcheck-client.crt --key=/etc/kubernetes/pki/etcd/healthcheck-client.key get / --prefix --keys-only
```

查看具体的key对应的数据：

```powershell
$ ETCDCTL_API=3 etcdctl --endpoints=https://[127.0.0.1]:2379 --cacert=/etc/kubernetes/pki/etcd/ca.crt --cert=/etc/kubernetes/pki/etcd/healthcheck-client.crt --key=/etc/kubernetes/pki/etcd/healthcheck-client.key get /registry/pods/jenkins/sonar-postgres-7fc5d748b6-gtmsb
```

#### 基于EFK实现kubernetes集群的日志平台（扩展）

##### EFK介绍

EFK工作示意

![](https://i0.hdslb.com/bfs/album/640d8dc0351f2feec93764157cebacfb14557b57.png)



- Elasticsearch

  一个开源的分布式、Restful 风格的搜索和数据分析引擎，它的底层是开源库Apache Lucene。它可以被下面这样准确地形容：

  - 一个分布式的实时文档存储，每个字段可以被索引与搜索；
  - 一个分布式实时分析搜索引擎；
  - 能胜任上百个服务节点的扩展，并支持 PB 级别的结构化或者非结构化数据。

- Fluentd

  一个针对日志的收集、处理、转发系统。通过丰富的插件系统，可以收集来自于各种系统或应用的日志，转化为用户指定的格式后，转发到用户所指定的日志存储系统之中。 

  ![](https://i0.hdslb.com/bfs/album/dae3b430bb8d05002c609341c3a06b2a9f8adfe6.jpg)

  Fluentd 通过一组给定的数据源抓取日志数据，处理后（转换成结构化的数据格式）将它们转发给其他服务，比如 Elasticsearch、对象存储、kafka等等。Fluentd 支持超过300个日志存储和分析服务，所以在这方面是非常灵活的。主要运行步骤如下

  1. 首先 Fluentd 从多个日志源获取数据
  2. 结构化并且标记这些数据
  3. 然后根据匹配的标签将数据发送到多个目标服务

- Kibana

  Kibana是一个开源的分析和可视化平台，设计用于和Elasticsearch一起工作。可以通过Kibana来搜索，查看，并和存储在Elasticsearch索引中的数据进行交互。也可以轻松地执行高级数据分析，并且以各种图标、表格和地图的形式可视化数据。

##### 部署es服务

###### 部署分析

1. es生产环境是部署es集群，通常会使用statefulset进行部署，此例由于演示环境资源问题，部署为单点
2. 数据存储挂载主机路径
3. es默认使用elasticsearch用户启动进程，es的数据目录是通过宿主机的路径挂载，因此目录权限被主机的目录权限覆盖，因此可以利用init container容器在es进程启动之前把目录的权限修改掉，注意init container要用特权模式启动。

###### 部署并验证

`efk/elasticsearch.yaml`

```powershell
apiVersion: apps/v1
kind: StatefulSet
metadata:
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
    k8s-app: elasticsearch-logging
    version: v7.4.2
  name: elasticsearch-logging
  namespace: logging
spec:
  replicas: 1
  revisionHistoryLimit: 10
  selector:
    matchLabels:
      k8s-app: elasticsearch-logging
      version: v7.4.2
  serviceName: elasticsearch-logging
  template:
    metadata:
      labels:
        k8s-app: elasticsearch-logging
        version: v7.4.2
    spec:
      nodeSelector:
        log: "true"	## 指定部署在哪个节点。需根据环境来修改
      containers:
      - env:
        - name: NAMESPACE
          valueFrom:
            fieldRef:
              apiVersion: v1
              fieldPath: metadata.namespace
        - name: cluster.initial_master_nodes
          value: elasticsearch-logging-0
        - name: ES_JAVA_OPTS
          value: "-Xms512m -Xmx512m"
        image: 172.21.32.6:5000/elasticsearch/elasticsearch:7.4.2
        name: elasticsearch-logging
        ports:
        - containerPort: 9200
          name: db
          protocol: TCP
        - containerPort: 9300
          name: transport
          protocol: TCP
        volumeMounts:
        - mountPath: /usr/share/elasticsearch/data
          name: elasticsearch-logging
      dnsConfig:
        options:
        - name: single-request-reopen
      initContainers:
      - command:
        - /sbin/sysctl
        - -w
        - vm.max_map_count=262144
        image: alpine:3.6
        imagePullPolicy: IfNotPresent
        name: elasticsearch-logging-init
        resources: {}
        securityContext:
          privileged: true
      - name: fix-permissions
        image: alpine:3.6
        command: ["sh", "-c", "chown -R 1000:1000 /usr/share/elasticsearch/data"]
        securityContext:
          privileged: true
        volumeMounts:
        - name: elasticsearch-logging
          mountPath: /usr/share/elasticsearch/data
      volumes:
      - name: elasticsearch-logging
        hostPath:
          path: /esdata
---
apiVersion: v1
kind: Service
metadata:
  labels:
    k8s-app: elasticsearch-logging
  name: elasticsearch
  namespace: logging
spec:
  ports:
  - port: 9200
    protocol: TCP
    targetPort: db
  selector:
    k8s-app: elasticsearch-logging
  type: ClusterIP
```

```powershell
$ kubectl create namespace logging
## 给slave1节点打上label，将es服务调度到slave1节点
$ kubectl label node k8s-slave1 log=true
## 部署服务，可以先去部署es的节点把镜像下载到本地
$ kubectl create -f elasticsearch.yaml
statefulset.apps/elasticsearch-logging created  
service/elasticsearch created

## 等待片刻，查看一下es的pod部署到了k8s-slave1节点，状态变为running
$ kubectl -n logging get po -o wide  
NAME                      READY   STATUS    RESTARTS   AGE   IP             NODE       
elasticsearch-logging-0   1/1     Running   0          69m   10.244.1.104   k8s-slave1   
# 然后通过curl命令访问一下服务，验证es是否部署成功
$ kubectl -n logging get svc  
NAME            TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)    AGE  
elasticsearch   ClusterIP   10.109.174.58   <none>        9200/TCP   71m  
$ curl 10.109.174.58:9200  
{  
  "name" : "elasticsearch-logging-0",  
  "cluster_name" : "docker-cluster",  
  "cluster_uuid" : "uic8xOyNSlGwvoY9DIBT1g",  
  "version" : {  
	"number" : "7.4.2",  
	"build_flavor" : "default",  
	"build_type" : "docker",  
	"build_hash" : "2f90bbf7b93631e52bafb59b3b049cb44ec25e96",  
	"build_date" : "2019-10-28T20:40:44.881551Z",  
	"build_snapshot" : false,  
	"lucene_version" : "8.2.0",  
	"minimum_wire_compatibility_version" : "6.8.0",  
	"minimum_index_compatibility_version" : "6.0.0-beta1"  
  },  
  "tagline" : "You Know, for Search"  
}  
```

##### 部署kibana

###### 部署分析

2. kibana需要暴漏web页面给前端使用，因此使用ingress配置域名来实现对kibana的访问
3. kibana为无状态应用，直接使用Deployment来启动
4. kibana需要访问es，直接利用k8s服务发现访问此地址即可，http://elasticsearch:9200

###### 部署并验证

资源文件 `efk/kibana.yaml`

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: kibana
  namespace: logging
  labels:
    app: kibana
spec:
  selector:
    matchLabels:
      app: kibana
  template:
    metadata:
      labels:
        app: kibana
    spec:
      containers:
      - name: kibana
        image: 172.21.32.6:5000/kibana/kibana:7.4.2
        resources:
          limits:
            cpu: 1000m
          requests:
            cpu: 100m
        env:
          - name: ELASTICSEARCH_URL
            value: http://elasticsearch:9200
        ports:
        - containerPort: 5601
---
apiVersion: v1
kind: Service
metadata:
  name: kibana
  namespace: logging
  labels:
    app: kibana
spec:
  ports:
  - port: 5601
    protocol: TCP
    targetPort: 5601
  type: ClusterIP
  selector:
    app: kibana
---
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: kibana
  namespace: logging
spec:
  rules:
  - host: kibana.devops.cn
    http:
      paths:
      - path: /
        backend:
          serviceName: kibana
          servicePort: 5601
```

```powershell
$ kubectl create -f kibana.yaml  
deployment.apps/kibana created
service/kibana created  
ingress/kibana created

# 然后查看pod，等待状态变成running
$ kubectl -n logging get po  
NAME                      READY   STATUS    RESTARTS   AGE  
elasticsearch-logging-0   1/1     Running   0          88m  
kibana-944c57766-ftlcw    1/1     Running   0          15m

## 配置域名解析 kibana.devops.cn，并访问服务进行验证，若可以访问，说明连接es成功
```

##### 部署fluentd

###### 部署分析

1. fluentd为日志采集服务，kubernetes集群的每个业务节点都有日志产生，因此需要使用daemonset的模式进行部署
2. 为进一步控制资源，会为daemonset指定一个选择表情，fluentd=true来做进一步过滤，只有带有此标签的节点才会部署fluentd
3. 日志采集，需要采集哪些目录下的日志，采集后发送到es端，因此需要配置的内容比较多，我们选择使用configmap的方式把配置文件整个挂载出来

###### 部署服务

配置文件，`efk/fluentd-es-main.yaml`

```yaml
apiVersion: v1
data:
  fluent.conf: |-
    # This is the root config file, which only includes components of the actual configuration
    #
    #  Do not collect fluentd's own logs to avoid infinite loops.
    <match fluent.**>
    @type null
    </match>

    @include /fluentd/etc/config.d/*.conf
kind: ConfigMap
metadata:
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
  name: fluentd-es-config-main
  namespace: logging
```

配置文件，fluentd-config.yaml，注意点：

1. 数据源source的配置，k8s会默认把容器的标准和错误输出日志重定向到宿主机中
2. 默认集成了 [kubernetes_metadata_filter](https://github.com/fabric8io/fluent-plugin-kubernetes_metadata_filter) 插件，来解析日志格式，得到k8s相关的元数据，raw.kubernetes
3. match输出到es端的flush配置

```yaml
kind: ConfigMap
apiVersion: v1
metadata:
  name: fluentd-config
  namespace: logging
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
data:
  system.conf: |-
    <system>
      root_dir /tmp/fluentd-buffers/
    </system>
  containers.input.conf: |-
    <source>
      @id fluentd-containers.log
      @type tail
      path /var/log/containers/*.log
      pos_file /var/log/es-containers.log.pos
      time_format %Y-%m-%dT%H:%M:%S.%NZ
      localtime
      tag raw.kubernetes.*
      format json
      read_from_head true
    </source>
    # Detect exceptions in the log output and forward them as one log entry.
    <match raw.kubernetes.**>
      @id raw.kubernetes
      @type detect_exceptions
      remove_tag_prefix raw
      message log
      stream stream
      multiline_flush_interval 5
      max_bytes 500000
      max_lines 1000
    </match>
  forward.input.conf: |-
    # Takes the messages sent over TCP
    <source>
      @type forward
    </source>
  output.conf: |-
    # Enriches records with Kubernetes metadata
    <filter kubernetes.**>
      @type kubernetes_metadata
    </filter>
    <match **>
      @id elasticsearch
      @type elasticsearch
      @log_level info
      include_tag_key true
      host elasticsearch
      port 9200
      logstash_format true
      request_timeout    30s
      <buffer>
        @type file
        path /var/log/fluentd-buffers/kubernetes.system.buffer
        flush_mode interval
        retry_type exponential_backoff
        flush_thread_count 2
        flush_interval 5s
        retry_forever
        retry_max_interval 30
        chunk_limit_size 2M
        queue_limit_length 8
        overflow_action block
      </buffer>
    </match>
```

daemonset定义文件，fluentd.yaml，注意点：

1. 需要配置rbac规则，因为需要访问k8s api去根据日志查询元数据
2. 需要将/var/log/containers/目录挂载到容器中
3. 需要将fluentd的configmap中的配置文件挂载到容器内
4. 想要部署fluentd的节点，需要添加fluentd=true的标签

`efk/fluentd.yaml`

```yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: fluentd-es
  namespace: logging
  labels:
    k8s-app: fluentd-es
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
---
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: fluentd-es
  labels:
    k8s-app: fluentd-es
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
rules:
- apiGroups:
  - ""
  resources:
  - "namespaces"
  - "pods"
  verbs:
  - "get"
  - "watch"
  - "list"
---
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: fluentd-es
  labels:
    k8s-app: fluentd-es
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
subjects:
- kind: ServiceAccount
  name: fluentd-es
  namespace: logging
  apiGroup: ""
roleRef:
  kind: ClusterRole
  name: fluentd-es
  apiGroup: ""
---
apiVersion: apps/v1
kind: DaemonSet
metadata:
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
    k8s-app: fluentd-es
  name: fluentd-es
  namespace: logging
spec:
  selector:
    matchLabels:
      k8s-app: fluentd-es
  template:
    metadata:
      labels:
        k8s-app: fluentd-es
    spec:
      containers:
      - env:
        - name: FLUENTD_ARGS
          value: --no-supervisor -q
        image: 172.21.32.6:5000/fluentd-es-root:v1.6.2-1.0
        imagePullPolicy: IfNotPresent
        name: fluentd-es
        resources:
          limits:
            memory: 500Mi
          requests:
            cpu: 100m
            memory: 200Mi
        volumeMounts:
        - mountPath: /var/log
          name: varlog
        - mountPath: /var/lib/docker/containers
          name: varlibdockercontainers
          readOnly: true
        - mountPath: /home/docker/containers
          name: varlibdockercontainershome
          readOnly: true
        - mountPath: /fluentd/etc/config.d
          name: config-volume
        - mountPath: /fluentd/etc/fluent.conf
          name: config-volume-main
          subPath: fluent.conf
      nodeSelector:
        fluentd: "true"
      securityContext: {}
      serviceAccount: fluentd-es
      serviceAccountName: fluentd-es
      volumes:
      - hostPath:
          path: /var/log
          type: ""
        name: varlog
      - hostPath:
          path: /var/lib/docker/containers
          type: ""
        name: varlibdockercontainers
      - hostPath:
          path: /home/docker/containers
          type: ""
        name: varlibdockercontainershome
      - configMap:
          defaultMode: 420
          name: fluentd-config
        name: config-volume
      - configMap:
          defaultMode: 420
          items:
          - key: fluent.conf
            path: fluent.conf
          name: fluentd-es-config-main
        name: config-volume-main
```

```powershell
## 给slave1和slave2打上标签，进行部署fluentd日志采集服务
$ kubectl label node k8s-slave1 fluentd=true  
node/k8s-slave1 labeled  
$ kubectl label node k8s-slave2 fluentd=true  
node/k8s-slave2 labeled  

# 创建服务
$ kubectl create -f fluentd-es-config-main.yaml  
configmap/fluentd-es-config-main created  
$ kubectl create -f fluentd-configmap.yaml  
configmap/fluentd-config created  
$ kubectl create -f fluentd.yaml  
serviceaccount/fluentd-es created  
clusterrole.rbac.authorization.k8s.io/fluentd-es created  
clusterrolebinding.rbac.authorization.k8s.io/fluentd-es created  
daemonset.extensions/fluentd-es created 

## 然后查看一下pod是否已经在k8s-slave1和k8s-slave2节点启动成功
$ kubectl -n logging get po -o wide
NAME                      READY   STATUS    RESTARTS   AGE  
elasticsearch-logging-0   1/1     Running   0          123m  
fluentd-es-246pl   		  1/1     Running   0          2m2s  
fluentd-es-4e21w   		  1/1     Running   0          2m10s 
kibana-944c57766-ftlcw    1/1     Running   0          50m
```

##### EFK功能验证

###### 验证思路

k8s-slave1和slave2中启动服务，同时往标准输出中打印测试日志，到kibana中查看是否可以收集

###### 创建测试容器

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: counter
spec:
  nodeSelector:
    fluentd: "true"
  containers:
  - name: count
    image: alpine:3.6
    args: [/bin/sh, -c,
            'i=0; while true; do echo "$i: $(date)"; i=$((i+1)); sleep 1; done']
```

```powershell
$ kubectl get po  
NAME                          READY   STATUS    RESTARTS   AGE  
counter                       1/1     Running   0          6s
```

###### 配置kibana

登录kibana界面，按照截图的顺序操作：

![](https://i0.hdslb.com/bfs/album/edfd40c0eadc2e32370e5b5805e0029156c2789f.png)

![](https://i0.hdslb.com/bfs/album/f3e59b9e3b25ac5a0fd9613cdf7903c9ccc5e811.png)

![](https://i0.hdslb.com/bfs/album/c2cafd9a6f7b8182edb0f40a3806c22c69708ddb.png)

![](https://i0.hdslb.com/bfs/album/6f034a6f3621e3c7a26f46d3af0da72ee51cec66.png)

也可以通过其他元数据来过滤日志数据，比如可以单击任何日志条目以查看其他元数据，如容器名称，Kubernetes 节点，命名空间等，比如kubernetes.pod_name : counter

到这里，我们就在 Kubernetes 集群上成功部署了 EFK ，要了解如何使用 Kibana 进行日志数据分析，可以参考 Kibana 用户指南文档：https://www.elastic.co/guide/en/kibana/current/index.html
