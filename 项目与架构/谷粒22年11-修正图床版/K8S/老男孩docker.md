# 第一章：容器技术概述

### 1.容器化技术的由来

> 容器是一种基础工具。泛指任何可以用于容纳其它物品的工具，可以部分或完全封闭，被用于容纳、储存、运输物品。物体可以被放置在容器中，而容器则可以保护内容物。
>
> 容器的类型：
> 瓶、罐、箱、篮、桶、袋、瓮、碗、柜、盆、鞘 …

把系统里的三个服务拆分开，借助容器运行互补干扰
![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_eaa33ca0f7f960412b13a5ce47d2303b_r.png)

### 2.容器发展的历史

容器概念始于 1979 年提出的 UNIX chroot，它是一个 UNIX 操作系统的系统调用，将一个进程及其子进程的根目录改变到文件系统中的一个新位置，让这些进程只能访问到这个新的位置，从而达到了进程隔离的目的。

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_e451aa3e2ed85977f816c2ad8e02af43_r.png)

2000 年的时候 FreeBSD 开发了一个类似于 chroot 的容器技术 Jails，这是最早期，也是功能最多的容器技术。Jails 英译过来是监狱的意思，这个“监狱”（用沙盒更为准确）包含了文件系统、用户、网络、进程等的隔离。

2001 Linux 也发布自己的容器技术 Linux VServer，2004 Solaris 也发布了 Solaris Containers，两者都将资源进行划分，形成一个个 zones，又叫做虚拟服务器。

2005 年推出 OpenVZ，它通过对 Linux 内核进行补丁来提供虚拟化的支持，每个 OpenVZ 容器完整支持了文件系统、用户及用户组、进程、网络、设备和 IPC 对象的隔离。

2007 年 Google 实现了 Control Groups( cgroups )，并加入到 Linux 内核中，这是划时代的，为后期容器的资源配额提供了技术保障。

2008 年基于 cgroups 和 linux namespace 推出了第一个最为完善的 Linux 容器 LXC。

2013 年推出到现在为止最为流行和使用最广泛的容器 Docker，相比其他早期的容器技术，Docker 引入了一整套容器管理的生态系统，包括分层的镜像模型，容器注册库，友好的 Rest API。

2014 年 CoreOS 也推出了一个类似于 Docker 的容器 Rocket，CoreOS 一个更加轻量级的 Linux 操作系统，在安全性上比 Docker 更严格。

2016 年微软也在 Windows 上提供了容器的支持，Docker 可以以原生方式运行在 Windows 上，而不是需要使用 Linux 虚拟机。

基本上到这个时间节点，容器技术就已经很成熟了，再往后就是容器云的发展，由此也衍生出多种容器云的平台管理技术，其中以 kubernetes 最为出众，有了这样一些细粒度的容器集群管理技术，也为微服务的发展奠定了基石。因此，对于未来来说，应用的微服务化是一个较大的趋势。

### 3.Docker的诞生和幕后的公司

2010年，几个大胡子年轻人在旧金山成立了一家做PaaS平台的公司，起名为[dotCloud]，dotCloud主要是基于PaaS平台为开发者或开发商提供技术服务。

Docker于2013.03.27正式作为public项目发布。

DotCloud公司2013年10月改名为Docker Inc，转型专注于Docker引擎和Docker生态系统。

落腮胡，皮夹克和摩托车，是Docker的灵魂人物CTO Solomon Hykeys的标准记号，外形看起来像是推动社会运动的老大。

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_644f1e6b4ea79fd4b329d02c479998cf_r.png)

Docker 容器的思想就是采用集装箱思想，为应用提供了一个基于容器的标准化运输系统。Docker 可以将任何应用及其依赖打包成一个轻量级、可移植、自包含的容器。容器可以运行在几乎所有的操作系统上。这样容器就可以跑在任何环境中，因此才有了那句话：

> Build Once, Run Anywhere

### 4.什么是docker？

Docker是Docker.inc公司开源的一个基于LXC技术之上构建的Container容器引擎，源代码托管在GitHub上，基于Go语言并遵从Apache2.0协议开源（可以商业）。

Docker项目的目标是实现轻量级的操作系统虚拟化解决方案。

Docker是容器引擎，Docker是通过内核虚拟化技术（namespaces及cgroups等）来提供容器的资源隔离与安全保障等。由于Docker通过操作系统层的虚拟化实现隔离，所以Docker容器在运行时，不需要类似虚拟机VM额外的操作系统开销，提高资源利用率。

微软，红帽Linux，IBM，Oracle等主流IT厂商已经在自己的产品里增加了对Docker的支持。

相比其他早期的容器技术，Docker引入了一整套容器管理的生态系统，包括分层的镜像模型，容器注册库，友好Rest API。

下面图比较了Docker和传统虚拟化方式的不同之处，可见容器是在操作系统层面上实现虚拟化，直接复制本地主机的操作系统，而传统方式则是在硬件层面实现。

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_01695e3d340b880de53c321ca0458677_r.png)

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_5020cce19b474b4b7c3687b8015263f8_r.png)

为了用docker容器模型，应用A和应用B在哪些层面上[隔离](https://www.cnblogs.com/rdchenxi/p/10455831.html)呢？

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_9dac59255d7815e9f010c612b6bb148a_r.png)

### 5.使用docker的优势

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_8d9d11aed21823a5e0185845c2676475_r.png)

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_7127665b1220ebf62118e95f5d53c101_r.png)

### 6.docker容器引擎介绍

Docker引擎可以从Docker网站下载，也可以基于Github上的源码进行构建。无论是开源版本还是商业版本，都有Linux和windows版本。

Docker引擎主要有两个版本：企业版（EE）和社区版（CE）。

每个季度，企业版和社区版都会发布一个稳定版本。社区版本会提供4个月的支持，而企业版本会提供12个月的支持。

从2017年第一季度开始，Docker版本号遵循YY.MM-xx格式，类似于Ubuntu等项目。例如，2018年6月第一次发布的社区版为18.06.0-ce。

注：2017年第一季度以前，Docker版本号遵循大版本号.小版本号的格式。采用新格式前的最后一个版本是Docker1.13。

# 第二章：Docker容器引擎的安装部署

### 1.基础环境检查

```shell
[root@node4 ~]# uname -a
Linux node4.98yz.cn 3.10.0-693.21.1.el7.x86_64 #1 SMP Wed Mar 7 19:03:37 UTC 2018 x86_64 x86_64 x86_64 GNU/Linux
[root@node4 ~]# cat /etc/redhat-release 
CentOS Linux release 7.4.1708 (Core) 
[root@node4 ~]# getenforce 
Disabled
[root@node4 ~]# systemctl stop firewalld
[root@node4 ~]# free -m
              total        used        free      shared  buff/cache   available
Mem:           5806         135        5526           8         144        5459
Swap:          1023           0        1023
[root@node4 ~]# ping baidu.com
PING baidu.com (220.181.38.148) 56(84) bytes of data.
64 bytes from 220.181.38.148 (220.181.38.148): icmp_seq=1 ttl=51 time=92.4 ms
64 bytes from 220.181.38.148 (220.181.38.148): icmp_seq=2 ttl=51 time=42.5 ms
```

### 2.安装epel源

```shell
[root@node4 ~]# yum install epel-release -y
```

这里可以看到epel源里也有docker1.13.1

```shell
[root@node4 ~]# yum list docker --show-duplicates
Loaded plugins: fastestmirror, langpacks
Loading mirror speeds from cached hostfile
 * base: mirrors.aliyun.com
 * epel: mirrors.aliyun.com
 * extras: mirrors.aliyun.com
 * updates: mirrors.aliyun.com
Available Packages
docker.x86_64                           2:1.13.1-102.git7f2769b.el7.centos                           extras
docker.x86_64                           2:1.13.1-103.git7f2769b.el7.centos                           extras
```

### 3.安装官方docker源

```shell
[root@node4 ~]# yum install yum-utils -y
[root@node4 ~]# yum-config-manager --add-repo http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo
```

这里可以看到docker-ce源里的docker镜像了。

```shell
[root@node4 ~]# yum list docker-ce --show-duplicates
Loaded plugins: fastestmirror, langpacks
docker-ce-stable                                                                                              | 3.5 kB  00:00:00     
(1/2): docker-ce-stable/x86_64/updateinfo                                                                     |   55 B  00:00:05     
(2/2): docker-ce-stable/x86_64/primary_db                                                                     |  37 kB  00:00:05     
Loading mirror speeds from cached hostfile
 * base: mirrors.aliyun.com
 * epel: mirrors.aliyun.com
 * extras: mirrors.aliyun.com
 * updates: mirrors.aliyun.com
Available Packages
docker-ce.x86_64               17.03.0.ce-1.el7.centos                       docker-ce-stable
docker-ce.x86_64               17.03.1.ce-1.el7.centos                       docker-ce-stable
docker-ce.x86_64               17.03.2.ce-1.el7.centos                       docker-ce-stable
docker-ce.x86_64               17.03.3.ce-1.el7                              docker-ce-stable
docker-ce.x86_64               17.06.0.ce-1.el7.centos                       docker-ce-stable
docker-ce.x86_64               17.06.1.ce-1.el7.centos                       docker-ce-stable
docker-ce.x86_64               17.06.2.ce-1.el7.centos                       docker-ce-stable
docker-ce.x86_64               17.09.0.ce-1.el7.centos                       docker-ce-stable
docker-ce.x86_64               17.09.1.ce-1.el7.centos                       docker-ce-stable
docker-ce.x86_64               17.12.0.ce-1.el7.centos                       docker-ce-stable
docker-ce.x86_64               17.12.1.ce-1.el7.centos                       docker-ce-stable
docker-ce.x86_64               18.03.0.ce-1.el7.centos                       docker-ce-stable
docker-ce.x86_64               18.03.1.ce-1.el7.centos                       docker-ce-stable
docker-ce.x86_64               18.06.0.ce-3.el7                              docker-ce-stable
docker-ce.x86_64               18.06.1.ce-3.el7                              docker-ce-stable
docker-ce.x86_64               18.06.2.ce-3.el7                              docker-ce-stable
docker-ce.x86_64               18.06.3.ce-3.el7                              docker-ce-stable
docker-ce.x86_64               3:18.09.0-3.el7                               docker-ce-stable
docker-ce.x86_64               3:18.09.1-3.el7                               docker-ce-stable
docker-ce.x86_64               3:18.09.2-3.el7                               docker-ce-stable
docker-ce.x86_64               3:18.09.3-3.el7                               docker-ce-stable
docker-ce.x86_64               3:18.09.4-3.el7                               docker-ce-stable
docker-ce.x86_64               3:18.09.5-3.el7                               docker-ce-stable
docker-ce.x86_64               3:18.09.6-3.el7                               docker-ce-stable
docker-ce.x86_64               3:18.09.7-3.el7                               docker-ce-stable
docker-ce.x86_64               3:18.09.8-3.el7                               docker-ce-stable
docker-ce.x86_64               3:18.09.9-3.el7                               docker-ce-stable
docker-ce.x86_64               3:19.03.0-3.el7                               docker-ce-stable
docker-ce.x86_64               3:19.03.1-3.el7                               docker-ce-stable
docker-ce.x86_64               3:19.03.2-3.el7                               docker-ce-stable
docker-ce.x86_64               3:19.03.3-3.el7                               docker-ce-stable
docker-ce.x86_64               3:19.03.4-3.el7                               docker-ce-stable
```

### 4.安装最新版docker并配置

```shell
[root@node4 ~]# yum install docker-ce -y

[root@node4 ~]# mkdir /etc/docker
[root@node4 ~]# cd /etc/docker/
[root@node4 docker]# vi /etc/docker/daemon.json
[root@node4 docker]# cat /etc/docker/daemon.json 
{
  "graph": "/data/docker",     # docker工作目录
  "storage-driver": "overlay2",   # 存储驱动
  "insecure-registries": ["registry.access.redhat.com","quay.io"],  # 不安全的仓库
  "registry-mirrors": ["https://q2gr04ke.mirror.aliyuncs.com"],  # 加速镜像
  "bip": "172.6.244.1/24",   # docker的网络,尽量要与宿主机有个对照关系
  "exec-opts": ["native.cgroupdriver=systemd"],  # cgroup的类型
  "live-restore": true  # 让docker容器不依懒docker引擎的死与活
}
```

### 5.启动docker

```shell
[root@node4 docker]# systemctl enable docker.service 
[root@node4 docker]# systemctl start docker.service 

[root@node4 ~]# docker version
Client: Docker Engine - Community
 Version:           19.03.4
 API version:       1.40
 Go version:        go1.12.10
 Git commit:        9013bf583a
 Built:             Fri Oct 18 15:52:22 2019
 OS/Arch:           linux/amd64
 Experimental:      false

Server: Docker Engine - Community
 Engine:
  Version:          19.03.4
  API version:      1.40 (minimum version 1.12)
  Go version:       go1.12.10
  Git commit:       9013bf583a
  Built:            Fri Oct 18 15:50:54 2019
  OS/Arch:          linux/amd64
  Experimental:     false
 containerd:
  Version:          1.2.10
  GitCommit:        b34a5c8af56e510852c35414db4c1f4fa6172339
 runc:
  Version:          1.0.0-rc8+dev
  GitCommit:        3e425f80a8c931f88e6d94a8c831b9d5aa481657
 docker-init:
  Version:          0.18.0
  GitCommit:        fec3683
```

### 6.核查docker配置及启动情况

```shell
[root@node4 ~]# docker info
Client:
 Debug Mode: false

Server:
 Containers: 0
  Running: 0
  Paused: 0
  Stopped: 0
 Images: 0
 Server Version: 19.03.4
 Storage Driver: overlay2
  Backing Filesystem: xfs
  Supports d_type: true
  Native Overlay Diff: true
 Logging Driver: json-file
 Cgroup Driver: systemd
 Plugins:
  Volume: local
  Network: bridge host ipvlan macvlan null overlay
  Log: awslogs fluentd gcplogs gelf journald json-file local logentries splunk syslog
 Swarm: inactive
 Runtimes: runc
 Default Runtime: runc
 Init Binary: docker-init
 containerd version: b34a5c8af56e510852c35414db4c1f4fa6172339
 runc version: 3e425f80a8c931f88e6d94a8c831b9d5aa481657
 init version: fec3683
 Security Options:
  seccomp
   Profile: default
 Kernel Version: 3.10.0-693.21.1.el7.x86_64
 Operating System: CentOS Linux 7 (Core)
 OSType: linux
 Architecture: x86_64
 CPUs: 4
 Total Memory: 5.671GiB
 Name: node4.98yz.cn
 ID: BSY4:7UBM:AA4T:OYVU:CRVZ:H5V3:L5X2:KJD4:2PHI:H74W:33N4:23KB
 Docker Root Dir: /data/docker
 Debug Mode: false
 Registry: https://index.docker.io/v1/
 Labels:
 Experimental: false
 Insecure Registries:
  quay.io
  registry.access.redhat.com
  127.0.0.0/8
 Registry Mirrors:
  https://q2gr04ke.mirror.aliyuncs.com/
 Live Restore Enabled: true
```

### 7.启动第一个Docker容器

```shell
[root@node4 ~]# docker run hello-world
Unable to find image 'hello-world:latest' locally
latest: Pulling from library/hello-world
1b930d010525: Pull complete 
Digest: sha256:c3b4ada4687bbaa170745b3e4dd8ac3f194ca95b2d0518b417fb47e5879d9b5f
Status: Downloaded newer image for hello-world:latest

Hello from Docker!
This message shows that your installation appears to be working correctly.

To generate this message, Docker took the following steps:
 1. The Docker client contacted the Docker daemon.
 2. The Docker daemon pulled the "hello-world" image from the Docker Hub.
    (amd64)
 3. The Docker daemon created a new container from that image which runs the
    executable that produces the output you are currently reading.
 4. The Docker daemon streamed that output to the Docker client, which sent it
    to your terminal.

To try something more ambitious, you can run an Ubuntu container with:
 $ docker run -it ubuntu bash

Share images, automate workflows, and more with a free Docker ID:
 https://hub.docker.com/

For more examples and ideas, visit:
 https://docs.docker.com/get-started/
```

# 第三章：Docker的镜像管理

### 1.hub.docker.com

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_e35df76da442b5988c59e003b0f79ee3_r.png)

### 2.注册一个dockerhub账号

登录hub.docker.com

```shell
[root@node4 ~]# docker login docker.io
Login with your Docker ID to push and pull images from Docker Hub. If you don't have a Docker ID, head over to https://hub.docker.com to create one.
Username: sunrisenan
Password: 
WARNING! Your password will be stored unencrypted in /root/.docker/config.json.
Configure a credential helper to remove this warning. See
https://docs.docker.com/engine/reference/commandline/login/#credentials-store

Login Succeeded
```

密码存在位置，用base64可以解码：

```shell
[root@node4 ~]# cat /root/.docker/config.json 
{
    "auths": {
        "https://index.docker.io/v1/": {
            "auth": "c3VucmlzZW5hbjpseXo1MjAx="
        }
    },
    "HttpHeaders": {
        "User-Agent": "Docker-Client/19.03.4 (linux)"
    }
}[root@node4 ~]# echo "c3VucmlzZW5hbjpseXo1MjAx="|base64 -d
sunrisenan:123
```

### 3.搜索一个镜像

- 命令行搜索

```shell
[root@node4 ~]# docker search alpine
NAME                                   DESCRIPTION                                     STARS               OFFICIAL            AUTOMATED
alpine                                 A minimal Docker image based on Alpine Linux…   5795                [OK]                
mhart/alpine-node                      Minimal Node.js built on Alpine Linux           444                                     
anapsix/alpine-java                    Oracle Java 8 (and 7) with GLIBC 2.28 over A…   428                                     [OK]
frolvlad/alpine-glibc                  Alpine Docker image with glibc (~12MB)          218                                     [OK]
gliderlabs/alpine                      Image based on Alpine Linux will help you wi…   180                                     
mvertes/alpine-mongo                   light MongoDB container                         106                                     [OK]
alpine/git                             A  simple git container running in alpine li…   100                                     [OK]
yobasystems/alpine-mariadb             MariaDB running on Alpine Linux [docker] [am…   52                                      [OK]
kiasaki/alpine-postgres                PostgreSQL docker image based on Alpine Linux   45                                      [OK]
alpine/socat                           Run socat command in alpine container           39                                      [OK]
davidcaste/alpine-tomcat               Apache Tomcat 7/8 using Oracle Java 7/8 with…   37                                      [OK]
zzrot/alpine-caddy                     Caddy Server Docker Container running on Alp…   35                                      [OK]
easypi/alpine-arm                      AlpineLinux for RaspberryPi                     32                                      
jfloff/alpine-python                   A small, more complete, Python Docker image …   28                                      [OK]
byrnedo/alpine-curl                    Alpine linux with curl installed and set as …   27                                      [OK]
hermsi/alpine-sshd                     Dockerize your OpenSSH-server with rsync and…   26                                      [OK]
etopian/alpine-php-wordpress           Alpine WordPress Nginx PHP-FPM WP-CLI           22                                      [OK]
hermsi/alpine-fpm-php                  Dockerize your FPM PHP 7.4 upon a lightweigh…   21                                      [OK]
bashell/alpine-bash                    Alpine Linux with /bin/bash as a default she…   14                                      [OK]
zenika/alpine-chrome                   Chrome running in headless mode in a tiny Al…   14                                      [OK]
davidcaste/alpine-java-unlimited-jce   Oracle Java 8 (and 7) with GLIBC 2.21 over A…   13                                      [OK]
tenstartups/alpine                     Alpine linux base docker image with useful p…   9                                       [OK]
spotify/alpine                         Alpine image with `bash` and `curl`.            9                                       [OK]
roribio16/alpine-sqs                   Dockerized ElasticMQ server + web UI over Al…   7                                       [OK]
rawmind/alpine-traefik                 This image is the traefik base. It comes fro…   5                                       [OK]
```

- 在网站上搜索

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_57e5346ff7517d24af888fdc6d1fabd6_r.png)

### 4.下载一个镜像

- 直接下载

```shell
[root@node4 ~]# docker pull alpine
Using default tag: latest
latest: Pulling from library/alpine
89d9c30c1d48: Pull complete 
Digest: sha256:c19173c5ada610a5989151111163d28a67368362762534d8a8121ce95cf2bd5a
Status: Downloaded newer image for alpine:latest
docker.io/library/alpine:latest
```

- 下载指定tag

```shell
[root@node4 ~]# docker pull alpine:3.10.1
3.10.1: Pulling from library/alpine
050382585609: Pull complete 
Digest: sha256:6a92cd1fcdc8d8cdec60f33dda4db2cb1fcdcacf3410a8e05b3741f44a9b5998
Status: Downloaded newer image for alpine:3.10.1
docker.io/library/alpine:3.10.1
```

> 镜像的结构：registry_name/repository_name/image_name:tag_name
> 例如：docker.io/library/alpine:3.10.1

### 5.查看本地镜像

```shell
[root@node4 ~]# docker images
REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
alpine              latest              965ea09ff2eb        2 weeks ago         5.55MB
alpine              3.10.1              b7b28af77ffe        3 months ago        5.58MB
hello-world         latest              fce289e99eb9        10 months ago       1.84kB

[root@node4 ~]# docker image ls
REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
alpine              latest              965ea09ff2eb        2 weeks ago         5.55MB
alpine              3.10.1              b7b28af77ffe        3 months ago        5.58MB
hello-world         latest              fce289e99eb9        10 months ago       1.84kB
```

### 6.给镜像打标签

```shell
[root@node4 ~]# docker tag b7b28af77ffe sunrisenan/alpine:3.10.1
[root@node4 ~]# docker images
REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
alpine              latest              965ea09ff2eb        2 weeks ago         5.55MB
alpine              3.10.1              b7b28af77ffe        3 months ago        5.58MB
sunrisenan/alpine   3.10.1              b7b28af77ffe        3 months ago        5.58MB
hello-world         latest              fce289e99eb9        10 months ago       1.84kB
```

### 7.推送镜像

```shell
[root@node4 ~]# docker push sunrisenan/alpine:3.10.1
The push refers to repository [docker.io/sunrisenan/alpine]
1bfeebd65323: Mounted from library/alpine 
3.10.1: digest: sha256:57334c50959f26ce1ee025d08f136c2292c128f84e7b229d1b0da5dac89e9866 size: 528
```

### 8.给镜像再打一个标签

```shell
[root@node4 ~]# docker tag 965ea09ff2eb docker.io/sunrisenan/alpine:3.10.3
[root@node4 ~]# docker images
REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
alpine              latest              965ea09ff2eb        2 weeks ago         5.55MB
sunrisenan/alpine   3.10.3              965ea09ff2eb        2 weeks ago         5.55MB
alpine              3.10.1              b7b28af77ffe        3 months ago        5.58MB
sunrisenan/alpine   3.10.1              b7b28af77ffe        3 months ago        5.58MB
hello-world         latest              fce289e99eb9        10 months ago       1.84kB
```

### 9.再次推送镜像

```shell
[root@node4 ~]# docker push docker.io/sunrisenan/alpine:3.10.3
The push refers to repository [docker.io/sunrisenan/alpine]
77cae8ab23bf: Mounted from library/alpine 
3.10.3: digest: sha256:e4355b66995c96b4b468159fc5c7e3540fcef961189ca13fee877798649f531a size: 528
```

### 10.删除镜像

- 删除镜像标签

```shell
[root@node4 ~]# docker rmi sunrisenan/alpine:3.10.1
Untagged: sunrisenan/alpine:3.10.1
Untagged: sunrisenan/alpine@sha256:57334c50959f26ce1ee025d08f136c2292c128f84e7b229d1b0da5dac89e9866
[root@node4 ~]# docker images | grep alpine
alpine              latest              965ea09ff2eb        2 weeks ago         5.55MB
sunrisenan/alpine   3.10.3              965ea09ff2eb        2 weeks ago         5.55MB
alpine              3.10.1              b7b28af77ffe        3 months ago        5.58MB
```

- 删除镜像

```shell
[root@node4 ~]# docker rmi b7b28af77ffe
Untagged: alpine:3.10.1
Untagged: alpine@sha256:6a92cd1fcdc8d8cdec60f33dda4db2cb1fcdcacf3410a8e05b3741f44a9b5998
Deleted: sha256:b7b28af77ffec6054d13378df4fdf02725830086c7444d9c278af25312aa39b9
Deleted: sha256:1bfeebd65323b8ddf5bd6a51cc7097b72788bc982e9ab3280d53d3c613adffa7
```

- 删除镜像标签及镜像

```shell
[root@node4 ~]# docker images | grep alpine
alpine              latest              965ea09ff2eb        2 weeks ago         5.55MB
sunrisenan/alpine   3.10.3              965ea09ff2eb        2 weeks ago         5.55MB
[root@node4 ~]# docker rmi -f 965ea09ff2eb   # -f 强制删除（包括标签及镜像）
Untagged: alpine:latest
Untagged: alpine@sha256:c19173c5ada610a5989151111163d28a67368362762534d8a8121ce95cf2bd5a
Untagged: sunrisenan/alpine:3.10.3
Untagged: sunrisenan/alpine@sha256:e4355b66995c96b4b468159fc5c7e3540fcef961189ca13fee877798649f531a
Deleted: sha256:965ea09ff2ebd2b9eeec88cd822ce156f6674c7e99be082c7efac3c62f3ff652
Deleted: sha256:77cae8ab23bf486355d1b3191259705374f4a11d483b24964d2f729dd8c076a0
[root@node4 ~]# docker images | grep alpine
[root@node4 ~]# 
```

### 11.Docker 镜像新特性

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_2bb8359fdcc9e0812b295d36a3150bb9_r.png)

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_f24dd512522de9ee061bd32e34070db2_r.png)

# 第四章：Docker的基本操作

### 1.查看本地的容器进程

```shell
[root@node4 ~]# docker ps -a
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS                   PORTS               NAMES
e63ab3e10bd6        hello-world         "/hello"            2 hours ago         Exited (0) 2 hours ago                       epic_bohr
```

### 2.启动容器（运行镜像）

```shell
docker run是日常用的最频繁的命令之一，同样也是较为复杂的命令之一
命令格式：docker run [OPTIONS] IMAGE [COMMAND] [ARG..]
OPTIONS:选项
    -i：表示启动一个可交互的容器，并持续打开标准输入
    -t:表示使用终端关联到容器的标准输入输出上
    -d：表示将容器放置后台运行
    --rm：退出后即删除容器
    --name：表示定义容器的唯一名称
    IMAGE：表示要运行的镜像
    COMMAND：表示启动容器时要运行的命令
```

**交互式启动一个容器**

```shell
[root@node4 ~]# docker run -ti sunrisenan/alpine:3.10.3 /bin/sh
Unable to find image 'sunrisenan/alpine:3.10.3' locally
3.10.3: Pulling from sunrisenan/alpine
89d9c30c1d48: Pull complete 
Digest: sha256:e4355b66995c96b4b468159fc5c7e3540fcef961189ca13fee877798649f531a
Status: Downloaded newer image for sunrisenan/alpine:3.10.3
/ # cat /etc/issue 
Welcome to Alpine Linux 3.10
Kernel \r on an \m (\l)
```

**非交互是启动一个容器**

```shell
[root@node4 ~]# docker run --rm sunrisenan/alpine:3.10.3 /bin/echo hello
hello
```

**非交互式启动一个后台容器**

```shell
[root@node4 ~]# docker run -d --name myalpine sunrisenan/alpine:3.10.3 /bin/sleep 300
3d015bfe776af7c14d8acbfecb6a09d2573a8e332d73a67030e4d3f651c03d2a
```

### 3.查看容器

**查看运行中的容器**

```shell
[root@node4 ~]# docker ps
CONTAINER ID        IMAGE                      COMMAND             CREATED             STATUS              PORTS               NAMES
3d015bfe776a        sunrisenan/alpine:3.10.3   "/bin/sleep 300"    4 minutes ago       Up 4 minutes                            myalpine
```

**查看所有容器**

```shell
[root@node4 ~]# docker ps -a
CONTAINER ID        IMAGE                      COMMAND             CREATED              STATUS                     PORTS               NAMES
7ecf398c77e4        hello-world                "/hello"            6 seconds ago        Exited (0) 5 seconds ago                       funny_wright
c44608564744        sunrisenan/alpine:3.10.3   "/bin/sleep 300"    About a minute ago   Up About a minute                              myalpine
```

**查看宿主机进程**

```shell
[root@node4 ~]# ps -aux|grep sleep|grep -v grep
root      6226  0.0  0.0   1540   248 ?        Ss   16:39   0:00 /bin/sleep 300
```

### 4.进入容器

```shell
[root@node4 ~]# docker exec -it myalpine /bin/sh
/ # ps aux
PID   USER     TIME  COMMAND
    1 root      0:00 /bin/sleep 300
    7 root      0:00 /bin/sh
   12 root      0:00 ps aux

/ # exit   # 退出容器
```

> docker attach 不推荐使用

### 5.停止/启动/重启容器

**停止容器**

```shell
[root@node4 ~]# docker ps
CONTAINER ID        IMAGE                      COMMAND             CREATED             STATUS              PORTS               NAMES
2d60b6eac204        sunrisenan/alpine:3.10.3   "/bin/sleep 300"    3 minutes ago       Up 3 minutes    
myalpine
[root@node4 ~]# docker stop myalpine
myalpine

[root@node4 ~]# docker ps -a
CONTAINER ID        IMAGE                      COMMAND             CREATED             STATUS                        PORTS               NAMES
2d60b6eac204        sunrisenan/alpine:3.10.3   "/bin/sleep 300"    4 minutes ago       Exited (137) 11 seconds ago                       myalpine
```

**启动容器**

```shell
[root@node4 ~]# docker start myalpine
myalpine
[root@node4 ~]# docker ps -a
CONTAINER ID        IMAGE                      COMMAND             CREATED             STATUS              PORTS               NAMES
2d60b6eac204        sunrisenan/alpine:3.10.3   "/bin/sleep 300"    7 minutes ago       Up 8 seconds                            myalpine
```

**重启容器**

```shell
[root@node4 ~]# docker restart myalpine
myalpine
[root@node4 ~]# docker ps -a|grep myalpine
2d60b6eac204        sunrisenan/alpine:3.10.3   "/bin/sleep 300"    9 minutes ago       Up 28 seconds                           myalpine
```

### 6.删除容器

```shell
[root@node4 ~]# docker rm myalpine
Error response from daemon: You cannot remove a running container db3335af09774651f55715b1da66ba930060b488ebdc7c4dc4ce67ce933dfc8e. Stop the container before attempting removal or force remove
[root@node4 ~]# docker rm -f myalpine
myalpine
```

> 如果没有停止的容器，需要用docker rm -f 进行强制删除

**删除已停止的容器**

```shell
[root@node4 ~]# for i in `docker ps -a|grep -i exit|awk '{print $1}'`;do docker rm -f $i;done
```

**删除所有的容器**

```shell
[root@node4 ~]# for i in `docker ps -qa`;do docker rm -f $i;done
```

### 7.修改/提交容器

**修改容器**

```shell
[root@node4 ~]# docker run -d --name myalpine sunrisenan/alpine:3.10.3 /bin/sleep 300
279bf6153576afcc8ca902d8c93282ec138b08704beb1f11f81fa106e0e2ad68

[root@node4 ~]# docker ps 
CONTAINER ID        IMAGE                      COMMAND             CREATED             STATUS              PORTS               NAMES
279bf6153576        sunrisenan/alpine:3.10.3   "/bin/sleep 300"    13 seconds ago      Up 11 seconds                           myalpine

[root@node4 ~]# docker exec -it 279bf6153576 /bin/sh
/ # pwd
/
/ # echo Hello > 1.txt
/ # ls
1.txt  bin    dev    etc    home   lib    media  mnt    opt    proc   root   run    sbin   srv    sys    tmp    usr    var
/ # exit
```

**提交容器**

```shell
[root@node4 ~]# docker commit -p 279bf6153576 sunrisenan/alpine:3.10.3_with_1.txt
sha256:5a766219fd567129bbc12f3e1293b782c35d7213ed5b54d8522e69a9f11c28b9
```

> -p 不让其他的内容再写进来了

**检验**

```shell
[root@node4 ~]# docker run --rm sunrisenan/alpine:3.10.3_with_1.txt /bin/cat /1.txt
Hello
```

> 也可以使用docker export 279bf6153576 >myalpine_with_1.txt.tar

### 8.导出/导入容器

**导出镜像**

```shell
[root@node4 ~]# docker images
REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
sunrisenan/alpine   3.10.3_with_1.txt   5a766219fd56        7 minutes ago       5.55MB
sunrisenan/alpine   3.10.3              965ea09ff2eb        2 weeks ago         5.55MB
sunrisenan/alpine   3.10.1              b7b28af77ffe        3 months ago        5.58MB
hello-world         latest              fce289e99eb9        10 months ago       1.84kB
[root@node4 ~]# docker save 5a766219fd56 > alpine_with_1.txt.tar
[root@node4 ~]# ls
alpine_with_1.txt.tar  anaconda-ks.cfg
```

> 导出的时候尽量加上版本号

**导入镜像**

```shell
[root@node4 ~]# docker rmi -f 5a766219fd56
Untagged: sunrisenan/alpine:3.10.3_with_1.txt
Deleted: sha256:5a766219fd567129bbc12f3e1293b782c35d7213ed5b54d8522e69a9f11c28b9
Deleted: sha256:34e6c5bebdd97e107ec3de8fb873b6f8bdd68fd593b511a737d41e62f80f58a8

[root@node4 ~]# docker load < ./alpine_with_1.txt.tar 
a87bd13a8b89: Loading layer [==================================================>]  3.584kB/3.584kB
Loaded image ID: sha256:5a766219fd567129bbc12f3e1293b782c35d7213ed5b54d8522e69a9f11c28b9
[root@node4 ~]# docker images
REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
<none>              <none>              5a766219fd56        10 minutes ago      5.55MB
sunrisenan/alpine   3.10.3              965ea09ff2eb        2 weeks ago         5.55MB
sunrisenan/alpine   3.10.1              b7b28af77ffe        3 months ago        5.58MB
hello-world         latest              fce289e99eb9        10 months ago       1.84kB

[root@node4 ~]# docker tag 5a766219fd56 sunrisenan/alpine:v3.10.3_with_1.txt
[root@node4 ~]# docker images
REPOSITORY          TAG                  IMAGE ID            CREATED             SIZE
sunrisenan/alpine   v3.10.3_with_1.txt   5a766219fd56        12 minutes ago      5.55MB
sunrisenan/alpine   3.10.3               965ea09ff2eb        2 weeks ago         5.55MB
sunrisenan/alpine   3.10.1               b7b28af77ffe        3 months ago        5.58MB
hello-world         latest               fce289e99eb9        10 months ago       1.84kB
```

**检验镜像**

```shell
[root@node4 ~]# docker run --rm -it --name myalpine_with_1.txt sunrisenan/alpine:v3.10.3_with_1.txt /bin/cat /1.txt
Hello
```

### 9.查看容器日志

```shell
[root@node4 ~]# docker run hello-world 2>&1 >>/dev/null 
[root@node4 ~]# docker ps -a
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS                      PORTS               NAMES
2c71417a3b27        hello-world         "/hello"            16 seconds ago      Exited (0) 15 seconds ago                       goofy_goldstine

[root@node4 ~]# docker logs 2c71417a3b27

Hello from Docker!
This message shows that your installation appears to be working correctly.

To generate this message, Docker took the following steps:
 1. The Docker client contacted the Docker daemon.
 2. The Docker daemon pulled the "hello-world" image from the Docker Hub.
    (amd64)
 3. The Docker daemon created a new container from that image which runs the
    executable that produces the output you are currently reading.
 4. The Docker daemon streamed that output to the Docker client, which sent it
    to your terminal.

To try something more ambitious, you can run an Ubuntu container with:
 $ docker run -it ubuntu bash

Share images, automate workflows, and more with a free Docker ID:
 https://hub.docker.com/

For more examples and ideas, visit:
 https://docs.docker.com/get-started/
```

> docker logs -f 2c71417a3b27 (-f 可以看实时日志)

# 第五章：Docker的高级操作

### 1.下载nginx镜像

```shell
[root@node4 ~]# docker pull nginx:1.12.2
1.12.2: Pulling from library/nginx
f2aa67a397c4: Pull complete 
e3eaf3d87fe0: Pull complete 
38cb13c1e4c9: Pull complete 
Digest: sha256:72daaf46f11cc753c4eab981cbf869919bd1fee3d2170a2adeac12400f494728
Status: Downloaded newer image for nginx:1.12.2
docker.io/library/nginx:1.12.2
[root@node4 ~]# docker images | grep nginx
nginx               1.12.2               4037a5562b03        18 months ago       108MB
[root@node4 ~]# docker tag 4037a5562b03 sunrisenan/nginx:v1.12.2
[root@node4 ~]# docker images | grep nginx
nginx               1.12.2               4037a5562b03        18 months ago       108MB
sunrisenan/nginx    v1.12.2              4037a5562b03        18 months ago       108MB
```

### 2.映射端口

- docker run -p 容器外的端口：容器内的端口

```shell
[root@node4 ~]# docker run -d -p81:80 sunrisenan/nginx:v1.12.2
19a8f420ac6fdc8c1a8d6a293dc49313134efaa63592a5d5d2b6272caf5ce1e8
[root@node4 ~]# netstat -lntup|grep 81
tcp6       0      0 :::81                   :::*                    LISTEN      9313/docker-proxy

[root@node4 ~]# curl 127.0.0.1:81
<!DOCTYPE html>
<html>
<head>
<title>Welcome to nginx!</title>
<style>
    body {
        width: 35em;
        margin: 0 auto;
        font-family: Tahoma, Verdana, Arial, sans-serif;
    }
</style>
</head>
<body>
<h1>Welcome to nginx!</h1>
<p>If you see this page, the nginx web server is successfully installed and
working. Further configuration is required.</p>

<p>For online documentation and support please refer to
<a href="http://nginx.org/">nginx.org</a>.<br/>
Commercial support is available at
<a href="http://nginx.com/">nginx.com</a>.</p>

<p><em>Thank you for using nginx.</em></p>
</body>
</html>
```

### 3.挂载数据卷

- docker run -v 容器外的目录：容器内的目录

```shell
[root@node4 ~]# mkdir html
[root@node4 ~]# cd html/
[root@node4 html]# wget www.baidu.com -O index.html
--2019-11-06 18:12:06--  http://www.baidu.com/
Resolving www.baidu.com (www.baidu.com)... 180.101.49.12, 180.101.49.11
Connecting to www.baidu.com (www.baidu.com)|180.101.49.12|:80... connected.
HTTP request sent, awaiting response... 200 OK
Length: 2381 (2.3K) [text/html]
Saving to: ‘index.html’

100%[===========================================================================================>] 2,381       --.-K/s   in 0s      

2019-11-06 18:12:11 (186 MB/s) - ‘index.html’ saved [2381/2381]

[root@node4 html]# docker run -d -p82:80 -v /root/html:/usr/share/nginx/html sunrisenan/nginx:v1.12.2
d44ce0d1bb4e42773dda95825cbc1a1989a38b704d0a133487f0adc054fbc88f
[root@node4 html]# docker ps -a | grep nginx
d44ce0d1bb4e        sunrisenan/nginx:v1.12.2   "nginx -g 'daemon of…"   24 seconds ago      Up 23 seconds       0.0.0.0:82->80/tcp   elated_williams
a5393a15159b        sunrisenan/nginx:v1.12.2   "nginx -g 'daemon of…"   4 minutes ago       Up 4 minutes        0.0.0.0:81->80/tcp   optimistic_rubin
```

浏览器打开[http://192.168.6.244:82](http://192.168.6.244:82/)

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_0242e19875fcab5b236bd7b25875c43a_r.png)

**查看挂载：**

```shell
[root@node4 ~]# docker inspect d44ce0d1bb4e
...
        "Mounts": [
            {
                "Type": "bind",
                "Source": "/root/html",
                "Destination": "/usr/share/nginx/html",
                "Mode": "",
                "RW": true,
                "Propagation": "rprivate"
            }
...
```

### 4.传递环境变量

- docker run -e 环境变量key：环境变量value

**传递一个环境标量**

```shell
[root@node4 ~]# docker run --rm -e E_OPTS=abcdefg sunrisenan/nginx:v1.12.2 printenv
PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
HOSTNAME=3011a2978c7c
E_OPTS=abcdefg
NGINX_VERSION=1.12.2-1~stretch
NJS_VERSION=1.12.2.0.1.14-1~stretch
HOME=/root
```

**传递两个环境标量**

```shell
[root@node4 ~]# docker run --rm -e E_OPTS=abcdefg -e G_OPES=123@PASS  sunrisenan/nginx:v1.12.2 printenv
PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
HOSTNAME=1bbb50ae0017
G_OPES=123@PASS
E_OPTS=abcdefg
NGINX_VERSION=1.12.2-1~stretch
NJS_VERSION=1.12.2.0.1.14-1~stretch
HOME=/root
```

### 5.容器内安装软件（工具）

- yum/apt-get/apt等

```shell
[root@node4 ~]# docker run --rm -it  sunrisenan/nginx:v1.12.2 bash
root@626e6ba6a358:/# curl
bash: curl: command not found
root@626e6ba6a358:/# tee /etc/apt/sources.list << EOF
> deb http://mirrors.163.com/debian/ jessie main non-free contrib
> deb http://mirrors.163.com/debian/ jessie-updates main non-free contrib
> EOF
deb http://mirrors.163.com/debian/ jessie main non-free contrib
deb http://mirrors.163.com/debian/ jessie-updates main non-free contrib

root@626e6ba6a358:/# apt-get update && apt-get install curl -y
```

> 提交容器并推送至dockerhub，这个景象还有用！

**镜像固化并推送到镜像仓库**

```shell
[root@node4 ~]# docker commit -p 626e6ba6a358 sunrisenan/nginx:curl
sha256:6ea518f9d8f6da45c14b9aa83043a62b7d0f639ceb821d718db545eea30f8693

[root@node4 ~]# docker images | grep curl
sunrisenan/nginx    curl                 6ea518f9d8f6        About a minute ago   136MB

[root@node4 ~]# docker push sunrisenan/nginx:curl
The push refers to repository [docker.io/sunrisenan/nginx]
b7f991e1e7fd: Pushed 
4258832b2570: Mounted from library/nginx 
683a28d1d7fd: Mounted from library/nginx 
d626a8ad97a1: Mounted from library/nginx 
curl: digest: sha256:03d6586ea5dcb82233f028340d85b96e7534aa3878180847729f1bd44c8384e6 size: 1160
```

### 6.容器的生命周期

- 检查本地是否存在镜像，如果不存在即从远端仓库检索
- 利用镜像启动容器
- 分配一个文件系统，并在只读的镜像层外挂载一层可读写层
- 从宿主机配置的网桥接口中桥接一个虚拟接口到容器
- 从地址池配置一个ip地址给容器
- 执行用户指定的指令
- 执行完毕后容器终止

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_81a9f584cb7dd9a685d3524018654531_r.png)

# 第六章：Docker file详解

### 1.dockerfile介绍

构建镜像的两种方式：

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_d97f57d111f60cca54b2337d4c9c5e75_r.png)

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_ff6fef7a652299892c47896bfa844f4e_r.png)

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_a17fd06bc4d871030347291db2d72174_r.png)

- **Dockerfile**
  - Dockerfile是一个文本文件，文件名：Dockerfile
  - 所有Dockerfile都是由FROM指令开始的
  - 构建命令：docker build . -t [registry]/[repositort]:[tag]

### 2.USER/WORKDIR指令

- Dockerfile

```shell
vi /data/dockerfile/Dockerfile

FROM sunrisenan/nginx:v1.12.2
USER nginx
WORKDIR /usr/share/nginx/html
```

- 构建镜像

```shell
[root@node4 dockerfile]# docker build . -t sunrisenan/nginx:v1.12.2_with_user_workdir
Sending build context to Docker daemon  2.048kB
Step 1/3 : FROM sunrisenan/nginx:v1.12.2
 ---> 4037a5562b03
Step 2/3 : USER nginx
 ---> Running in ed1dad767416
Removing intermediate container ed1dad767416
 ---> 2e96a40b4b85
Step 3/3 : WORKDIR /usr/share/nginx/html
 ---> Running in 35613b2afbe0
Removing intermediate container 35613b2afbe0
 ---> 20a3068e906c
Successfully built 20a3068e906c
Successfully tagged sunrisenan/nginx:v1.12.2_with_user_workdir
```

- 验证

```shell
[root@node4 dockerfile]# docker run --rm  -it --name nginx123  20a3068e906c /bin/bash
nginx@114eb12feb40:/usr/share/nginx/html$ whoami
nginx
nginx@114eb12feb40:/usr/share/nginx/html$ pwd
/usr/share/nginx/html
```

### 3.ADD/EXPOSE指令

- Dockerfile

```shell
wget -O /data/dockerfile/index.html www.baidu.com
vi /data/dockerfile/Dockerfile

FROM sunrisenan/nginx:v1.12.2
ADD index.html /usr/share/nginx/html/index.html
EXPOSE 80
```

- 构建镜像

```shell
[root@node4 dockerfile]# docker build . -t sunrisenan/nginx:v1.12.2_with_index_expose
Sending build context to Docker daemon   5.12kB
Step 1/3 : FROM sunrisenan/nginx:v1.12.2
 ---> 4037a5562b03
Step 2/3 : ADD index.html /usr/share/nginx/html/index.html
 ---> 6e35f2fd8a8c
Step 3/3 : EXPOSE 80
 ---> Running in d5f714442b32
Removing intermediate container d5f714442b32
 ---> 643ede2cc9e0
Successfully built 643ede2cc9e0
Successfully tagged sunrisenan/nginx:v1.12.2_with_index_expose
```

- 验证1

```shell
[root@node4 dockerfile]# docker run -P  --rm  -it --name nginx123  sunrisenan/nginx:v1.12.2_with_index_expose /bin/bash
root@f5b5689c3442:/# nginx -g "daemon off;"

[root@node4 dockerfile]# netstat -lntup|grep docker-proxy
tcp6       0      0 :::81                   :::*                    LISTEN      10609/docker-proxy  
tcp6       0      0 :::82                   :::*                    LISTEN      10806/docker-proxy  
tcp6       0      0 :::4000                 :::*                    LISTEN      22183/docker-proxy 
```

> -P随机暴露一个端口

- 验证2

```shell
[root@node4 dockerfile]# docker run -P -d --rm --name nginx123  sunrisenan/nginx:v1.12.2_with_index_expose
54007722a39f28c7c7ede6125fbb3f51c76bcdb7004d7635d1f37385573cb70a

[root@node4 dockerfile]# netstat -lntup|grep docker
tcp6       0      0 :::81                   :::*                    LISTEN      10609/docker-proxy  
tcp6       0      0 :::82                   :::*                    LISTEN      10806/docker-proxy  
tcp6       0      0 :::4001                 :::*                    LISTEN      22324/docker-proxy  
```

> EXPOSE 80 写不写都行,都能用-p 都能映射出来

### 4.RUN/ENV指令

- Dockerfile

```shell
vi /data/dockerfile/Dockerfile

FROM centos:7
ENV VER 9.11.4
RUN yum install bind-$VER -y
```

> ENV VER 空格 什么，或VER= 什么,ENV可以在构建镜像里面用，同时也可以在系统里面用
> RUN 在构建镜像的时候执行一些可执行命令

- 构建镜像

```shell
[root@node4 dockerfile]# docker build . -t sunrisenan/bind:v9.11.4_with_env_run
Sending build context to Docker daemon   5.12kB
Step 1/3 : FROM centos:7
7: Pulling from library/centos
d8d02d457314: Already exists 
Digest: sha256:307835c385f656ec2e2fec602cf093224173c51119bbebd602c53c3653a3d6eb
Status: Downloaded newer image for centos:7
 ---> 67fa590cfc1c
Step 2/3 : ENV VER 9.11.4
 ---> Running in 5a554643a9d0
Removing intermediate container 5a554643a9d0
 ---> 41bb82f39287
Step 3/3 : RUN yum install bind-$VER -y
 ---> Running in c80fa75b61af
Loaded plugins: fastestmirror, ovl
Determining fastest mirrors
 * base: ftp.sjtu.edu.cn
 * extras: mirrors.aliyun.com
 * updates: mirrors.aliyun.com
Resolving Dependencies
--> Running transaction check
---> Package bind.x86_64 32:9.11.4-9.P2.el7 will be installed
.....
 ---> c6f46e66d43d
Successfully built c6f46e66d43d
Successfully tagged sunrisenan/bind:v9.11.4_with_env_run
```

- 运行容器

```shell
[root@node4 dockerfile]# docker run --rm sunrisenan/bind:v9.11.4_with_env_run rpm -qa bind
bind-9.11.4-9.P2.el7.x86_64

[root@node4 dockerfile]# docker run --rm sunrisenan/bind:v9.11.4_with_env_run printenv
PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
HOSTNAME=d71390b7a05c
VER=9.11.4
HOME=/root
```

### 5.CMD/ENTRYPOINT指令

> CMD指令和ENTRYPOINT指令作用相同，使用方法略有不同

#### 5.1 CMD指令

- Dockerfile

```shell
vi /data/dockerfile/Dockerfile

FROM centos:7
RUN yum install httpd -y
CMD ["httpd","-D","FOREGROUND"]    # 让其前台运行
```

- 构建镜像

```shell
[root@node4 dockerfile]# docker build . -t sunrisenan/httpd:test
```

- 运行容器

```shell
[root@node4 dockerfile]# docker run -d --rm  --name myhttpd -p83:80 sunrisenan/httpd:test
9d8153320901e7ebf609eecacea5c4ce371ecabd31a0ec287ab27791cb25288d
[root@node4 dockerfile]# docker ps -a
CONTAINER ID        IMAGE                                        COMMAND                  CREATED             STATUS                      PORTS                  NAMES
9d8153320901        sunrisenan/httpd:test                        "httpd -D FOREGROUND"    33 seconds ago      Up 32 seconds               0.0.0.0:83->80/tcp     myhttpd
```

- 浏览器访问[http://192.168.6.244:83](http://192.168.6.244:83/)

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_62579dbf724990c3a287eea1176ae4ae_r.png)

#### 5.2 ENTRYPOINT指令

- Dockerfile

```shell
vi /data/dockerfile/Dockerfile

FROM centos:7
ADD entrypoint.sh /entrypoint.sh
RUN yum install epel-release -q -y && yum install nginx -y
ENTRYPOINT /entrypoint.sh
```

- entrypoint.sh

```shell
vi /data/dockerfile/entrypoint.sh

#!/bin/bash
/sbin/nginx -g "daemon off;"
```

**注意：**entrypoint.sh不要忘了给执行权限！

- 构建镜像

```shell
[root@node4 dockerfile]# docker build . -t sunrisenan/nginx:mynginx
```

- 运行镜像

```shell
[root@node4 dockerfile]# docker run -d --rm -p84:80 sunrisenan/nginx:mynginx 
c9b6f81b177022fe45d1eb14b6e20c02cf1a0307b648b35cc964f4c094512564

[root@node4 dockerfile]# docker ps | grep mynginx
c9b6f81b1770        sunrisenan/nginx:mynginx                     "/bin/sh -c /entrypo…"   36 seconds ago      Up 34 seconds       0.0.0.0:84->80/tcp     sweet_williams
```

- 浏览器访问[http://192.168.6.244:84](http://192.168.6.244:84/)

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_d71e5344f6971fda8e4fb830628e9ed3_r.png)

# 第七章：综合实验

> 运行一个docker容器，在浏览器打开demo.od.com能访问到百度首页

### 1.准备Docker镜像

- Dockerfile

```shell
vi /data/dockerfile/Dockerfile

FROM sunrisenan/nginx:v1.12.2
USER root
ENV WWW /usr/share/nginx/html
ENV CONF /etc/nginx/conf.d
RUN /bin/cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime &&\ 
    echo 'Asia/Shanghai' >/etc/timezone
WORKDIR $WWW
ADD index.html $WWW/index.html
ADD demo.od.com.conf $CONF/demo.od.com.conf
EXPOSE 80
CMD ["nginx","-g","daemon off;"]
```

- index.html

```shell
/data/dockerfile/nginx/index.html

wget -O /data/dockerfile/index.html www.baidu.com
```

- demo.od.com.conf

```shell
/data/dockerfile/nginx/demo.od.com.conf

server {
   listen 80;
   server_name demo.od.com;

   root /usr/share/nginx/html;
}
```

### 2.构建镜像

```shell
[root@node4 nginx]# docker build . -t sunrisenan/nginx:baidu
```

### 3.运行容器

```shell
[root@node4 nginx]# docker run -d --rm -p80:80 --name baidu sunrisenan/nginx:baidu
e3c5ea1d36c35e4f35d36456781e9a021fba4a8c0cc91e2df33a4d13dc9fe311

[root@node4 nginx]# docker ps | grep baidu
e3c5ea1d36c3        sunrisenan/nginx:baidu                       "nginx -g 'daemon of…"   2 minutes ago       Up 2 minutes        0.0.0.0:80->80/tcp     baidu
```

### 4.验证

- 本地绑定host

C:\Windows\System32\drivers\etc\hosts

```shell
192.168.6.244 demo.od.com
```

- 浏览器访问

打开浏览器访问http://demo.od.com/

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_b0cc03d0f7b6e64e8c331a6dc7f5da86_r.png)

# 第八章：Docker的网络模型

### 1.NAT(默认)

```shell
[root@node4 ~]# docker run -ti --rm sunrisenan/alpine:3.10.3 /bin/sh
/ # ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN qlen 1
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
100: eth0@if101: <BROADCAST,MULTICAST,UP,LOWER_UP,M-DOWN> mtu 1500 qdisc noqueue state UP 
    link/ether 02:42:ac:06:f4:02 brd ff:ff:ff:ff:ff:ff
    inet 172.6.244.2/24 brd 172.6.244.255 scope global eth0
       valid_lft forever preferred_lft forever
```

### 2.None

```shell
[root@node4 ~]# docker run -ti --rm --net=none  sunrisenan/alpine:3.10.3 /bin/sh
/ # ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN qlen 1
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
/ # 
```

### 3.Host

```shell
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
/ # 
```

### 4.联合网络*

```shell
[root@node4 ~]# docker run -d sunrisenan/nginx:v1.12.2
76f844e6057b6493a4a8933819ade7c29325ef17bc1fddc88bdf4c7ec60c620b
[root@node4 ~]# 
[root@node4 ~]# docker ps -qa
76f844e6057b
[root@node4 ~]# docker run -ti --rm --net=container:76f844e6057b sunrisenan/nginx:curl bash
root@76f844e6057b:/# curl localhost
<!DOCTYPE html>
<html>
<head>
<title>Welcome to nginx!</title>
<style>
    body {
        width: 35em;
        margin: 0 auto;
        font-family: Tahoma, Verdana, Arial, sans-serif;
    }
</style>
</head>
<body>
<h1>Welcome to nginx!</h1>
<p>If you see this page, the nginx web server is successfully installed and
working. Further configuration is required.</p>

<p>For online documentation and support please refer to
<a href="http://nginx.org/">nginx.org</a>.<br/>
Commercial support is available at
<a href="http://nginx.com/">nginx.com</a>.</p>

<p><em>Thank you for using nginx.</em></p>
</body>
</html>
root@76f844e6057b:/# 
```

# 第九篇：快速更改密码

```shell
[root@shkf6-240 ~]# echo "1" |passwd root --stdin
```