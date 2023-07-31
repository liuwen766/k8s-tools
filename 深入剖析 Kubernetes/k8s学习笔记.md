# 深入理解k8s

## 01-K8s开篇词

**因“容器”而起的技术变革中，Kubernetes 项目已然成为容器技术的事实标准，重新定义了基础设施领域对应用编排与管理的种种可能。** 



总有很多相似的问题被反复提及，比如：

1. 为什么容器里只能跑“一个进程”？
2. 为什么我原先一直在用的某个 JVM 参数，在容器里就不好使了？
3. 为什么 Kubernetes 就不能固定 IP 地址？容器网络连不通又该如何去 Debug？
4. Kubernetes 中 StatefulSet 和 Operator 到底什么区别？PV 和 PVC 这些概念又该怎么用？

这些问题乍一看与我们平常的认知非常矛盾，但它们的答案和原理却并不复杂。不过很遗憾，对于刚刚开始学习容器的技术人员来说，它们却很难用一两句话就能解释清楚。 

究其原因在于，**从过去以物理机和虚拟机为主体的开发运维环境，向以容器为核心的基础设施的转变过程，并不是一次温和的改革，而是涵盖了对网络、存储、调度、操作系统、分布式原理等各个方面的容器化理解和改造。** 



这些关于 Linux 内核、分布式系统、网络、存储等方方面面的积累，并不会在 Docker 或者 Kubernetes 的文档中交代清楚。可偏偏就是它们，才是真正掌握容器技术体系的精髓所在，是每一位技术从业者需要悉心修炼的“内功”。 

其实，容器技术体系看似纷乱繁杂，却存在着很多可以“牵一发而动全身”的主线。比如，Linux 的进程模型对于容器本身的重要意义；或者，“控制器”模式对整个 Kubernetes 项目提纲挈领的作用。 

借由这个专栏，给你讲清楚容器背后的这些技术本质与设计思想，并结合着对核心特性的剖析与实践，加深你对容器技术的理解。为此，我把专栏划分成了 4 大模块：

1. **“白话”容器技术基础：** 我希望用饶有趣味的解说，给你梳理容器技术生态的发展脉络，用最通俗易懂的语言描述容器底层技术的实现方式，让你知其然，也知其所以然。
2. **Kubernetes 集群的搭建与实践：** Kubernetes 集群号称“非常复杂”，但是如果明白了其中的架构和原理，选择了正确的工具和方法，它的搭建却也可以“一键安装”，它的应用部署也可以浅显易懂。
3. **容器编排与 Kubernetes 核心特性剖析：** 这是这个专栏最重要的内容。“编排”永远都是容器云项目的灵魂所在，也是 Kubernetes 社区持久生命力的源泉。在这一模块，我会从分布式系统设计的视角出发，抽象和归纳出这些特性中体现出来的普遍方法，然后带着这些指导思想去逐一阐述 Kubernetes 项目关于编排、调度和作业管理的各项核心特性。“不识庐山真面目，只缘身在此山中”，希望这样一个与众不同的角度，能够给你以全新的启发。
4. **Kubernetes 开源社区与生态：**“开源生态”永远都是容器技术和 Kubernetes 项目成功的关键。在这个模块，我会和你一起探讨，容器社区在开源软件工程指导下的演进之路；带你思考，如何同团队一起平衡内外部需求，让自己逐渐成为社区中不可或缺的一员。



## 02-容器技术预习篇 

Cloud Foundry → Docker → Kubernetes

### 1、初出茅庐

相比于的如日中天 AWS 和盛极一时的 OpenStack，以 Cloud Foundry 为代表的开源 PaaS 项目，却成为了当时云计算技术中的一股清流。 

当时还名叫 dotCloud 的 Docker 公司，眼看就要被如火如荼的 PaaS 风潮抛弃，dotCloud 公司却做出了这样一个决定：开源自己的容器项目 Docker。 

**PaaS 项目被大家接纳的一个主要原因，就是它提供了一种名叫“应用托管”的能力。** 在当时，虚拟机和云计算已经是比较普遍的技术和服务了，那时主流用户的普遍用法，就是租一批 AWS 或者 OpenStack 的虚拟机，然后像以前管理物理服务器那样，用脚本或者手工的方式在这些机器上部署应用。 

当然，这个部署过程难免会碰到云端虚拟机和本地环境不一致的问题，所以当时的云计算服务，比的就是谁能更好地模拟本地服务器环境，能带来更好的“上云”体验。而 PaaS 开源项目的出现，就是当时解决这个问题的一个最佳方案。 

事实上，**像 Cloud Foundry 这样的 PaaS 项目，最核心的组件就是一套应用的打包和分发机制。** Cloud Foundry 为每种主流编程语言都定义了一种打包格式，而“cf push”的作用，基本上等同于用户把应用的可执行文件和启动脚本打进一个压缩包内，上传到云上 Cloud Foundry 的存储中。接着，Cloud Foundry 会通过调度器选择一个可以运行这个应用的虚拟机，然后通知这个机器上的 Agent 把应用压缩包下载下来启动。 

这时候关键来了，由于需要在一个虚拟机上启动很多个来自不同用户的应用，Cloud Foundry 会调用操作系统的 Cgroups 和 Namespace 机制为每一个应用单独创建一个称作“沙盒”的隔离环境，然后在“沙盒”中启动这些应用进程。这样，就实现了把多个用户的应用互不干涉地在虚拟机里批量地、自动地运行起来的目的。**这，正是 PaaS 项目最核心的能力。** 而这些 Cloud Foundry 用来运行应用的隔离环境，或者说“沙盒”，就是所谓的“容器”。 

事实上，Docker 项目确实与 Cloud Foundry 的容器在大部分功能和实现原理上都是一样的，可偏偏就是这剩下的一小部分不一样的功能，成了 Docker 项目接下来“呼风唤雨”的不二法宝。**这个功能，就是 Docker 镜像。**

出现这个问题的根本原因是，一旦用上了 PaaS，用户就必须为每种语言、每种框架，甚至每个版本的应用维护一个打好的包。这个打包过程，没有任何章法可循，更麻烦的是，明明在本地运行得好好的应用，却需要做很多修改和配置工作才能在 PaaS 里运行起来。而这些修改和配置，并没有什么经验可以借鉴，基本上得靠不断试错，直到你摸清楚了本地应用和远端 PaaS 匹配的“脾气”才能够搞定。 

而**Docker 镜像解决的，恰恰就是打包这个根本性的问题。** 所谓 Docker 镜像，其实就是一个压缩包。但是这个压缩包里的内容，比 PaaS 的应用可执行文件 + 启停脚本的组合就要丰富多了。实际上，大多数 Docker 镜像是直接由一个完整操作系统的所有文件和目录构成的，所以这个压缩包里的内容跟你本地开发和测试环境用的操作系统是完全一样的。这，**正是 Docker 镜像的精髓。**本地环境和云端环境的高度一致！ 

所以，**Docker 项目给 PaaS 世界带来的“降维打击”，其实是提供了一种非常便利的打包机制。这种机制直接打包了应用运行所需要的整个操作系统，从而保证了本地环境和云端环境的高度一致，避免了用户通过“试错”来匹配两种不同运行环境之间差异的痛苦过程。** 



2013~2014 年，以 Cloud Foundry 为代表的 PaaS 项目，逐渐完成了教育用户和开拓市场的艰巨任务，也正是在这个将概念逐渐落地的过程中，应用“打包”困难这个问题，成了整个后端技术圈子的一块心病。

Docker 项目的出现，则为这个根本性的问题提供了一个近乎完美的解决方案。这正是 Docker 项目刚刚开源不久，就能够带领一家原本默默无闻的 PaaS 创业公司脱颖而出，然后迅速占领了所有云计算领域头条的技术原因。

而在成为了基础设施领域近十年难得一见的技术明星之后，dotCloud 公司则在 2013 年底大胆改名为 Docker 公司。不过，这个在当时就颇具争议的改名举动，也成为了日后容器技术圈风云变幻的一个关键伏笔。 

 

### 2、崭露头角

**而 Docker 项目之所以能取得如此高的关注，一方面正如前面我所说的那样，它解决了应用打包和发布这一困扰运维人员多年的技术难题；而另一方面，就是因为它第一次把一个纯后端的技术概念，通过非常友好的设计和封装，交到了最广大的开发者群体手里。** 

**解决了应用打包这个根本性的问题，同开发者与生俱来的的亲密关系，再加上 PaaS 概念已经深入人心的完美契机，成为 Docker 这个技术上看似平淡无奇的项目一举走红的重要原因。** 

总结：Docker 项目在短时间内迅速崛起的三个重要原因：

1. Docker 镜像通过技术手段解决了 PaaS 的根本性问题；
2. Docker 容器同开发者之间有着与生俱来的密切关系；
3. PaaS 概念已经深入人心的完美契机。

Docker 项目从发布之初就全面发力，从技术、社区、商业、市场全方位争取到的开发者群体，实际上是为此后吸引整个生态到自家“PaaS”上的一个铺垫。**只不过这时，“PaaS”的定义已经全然不是 Cloud Foundry 描述的那个样子，而是变成了一套以 Docker 容器为技术核心，以 Docker 镜像为打包标准的、全新的“容器化”思路。** 

### 3、群雄并起

虽然 Docker 项目备受追捧，但用户们最终要部署的，还是他们的网站、服务、数据库，甚至是云计算业务。

这就意味着，只有那些能够为用户提供平台层能力的工具，才会真正成为开发者们关心和愿意付费的产品。而 Docker 项目这样一个只能用来创建和启停容器的小工具，最终只能充当这些平台项目的“幕后英雄”。



Docker 项目发布后，CoreOS 公司很快就认识到可以把“容器”的概念无缝集成到自己的这套方案中，从而为用户提供更高层次的 PaaS 能力。所以，CoreOS 很早就成了 Docker 项目的贡献者，并在短时间内成为了 Docker 项目中第二重要的力量。

相较于 CoreOS 是依托于一系列开源项目（比如 Container Linux 操作系统、Fleet 作业调度工具、systemd 进程管理和 rkt 容器），一层层搭建起来的平台产品，Swarm 项目则是以一个完整的整体来对外提供集群管理功能。而 Swarm 的最大亮点，则是它完全使用 Docker 项目原本的容器管理 API 来完成集群管理。 

所以在部署了 Swarm 的多机环境下，用户只需要使用原先的 Docker 指令创建一个容器，这个请求就会被 Swarm 拦截下来处理，然后通过具体的调度算法找到一个合适的 Docker Daemon 运行起来。 

**Docker 公司，开始及时地借助这波浪潮通过并购来完善自己的平台层能力**。其中一个最成功的案例，莫过于对 Fig 项目的收购。**Fig 项目之所以受欢迎，在于它在开发者面前第一次提出了“容器编排”（Container Orchestration）的概念。** 

编排，它主要是指用户如何通过某些工具或者配置来完成一组虚拟机以及关联资源的定义、配置、创建、删除等工作，然后由云计算平台按照这些指定的逻辑来完成的过程。 

Fig 就会把这些容器的定义和配置交给 Docker API 按照访问逻辑依次创建，你的一系列容器就都启动了；而容器 A 与 B 之间的关联关系，也会交给 Docker 的 Link 功能通过写入 hosts 文件的方式进行配置。更重要的是，你还可以在 Fig 的配置文件里定义各种容器的副本个数等编排参数，再加上 Swarm 的集群管理能力，一个活脱脱的 PaaS 呼之欲出。 

Fig 项目被收购后改名为 Compose，它成了 Docker 公司到目前为止第二大受欢迎的项目，一直到今天也依然被很多人使用。 

当时的这个容器生态里，还有很多令人眼前一亮的开源项目或公司。比如，专门负责处理容器网络的 SocketPlane 项目（后来被 Docker 公司收购），专门负责处理容器存储的 Flocker 项目（后来被 EMC 公司收购），专门给 Docker 集群做图形化管理界面和对外提供云服务的 Tutum 项目（后来被 Docker 公司收购）等等。

一时之间，整个后端和云计算领域的聪明才俊都汇集在了这个“小鲸鱼”的周围，为 Docker 生态的蓬勃发展献上了自己的智慧。



而除了这个异常繁荣的、围绕着 Docker 项目和公司的生态之外，还有一个势力在当时也是风头无两，这就是老牌集群管理项目 Mesos 和它背后的创业公司 Mesosphere。 它发布了一个名为 Marathon 的项目，而这个项目很快就成为了 Docker Swarm 的一个有力竞争对手。**虽然不能提供像 Swarm 那样的原生 Docker API，Mesos 社区却拥有一个独特的竞争力：超大规模集群的管理经验。**它旨在使用户能够像管理一台机器那样管理一个万级别的物理机集群，并且使用 Docker 容器在这个集群里自由地部署应用。而这，对很多大型企业来说具有着非同寻常的吸引力。 

这时，如果你再去审视当时的容器技术生态，就不难发现 CoreOS 公司竟然显得有些尴尬了。它的 rkt 容器完全打不开局面，Fleet 集群管理项目更是少有人问津，CoreOS 完全被 Docker 公司压制了。

而处境同样不容乐观的似乎还有 RedHat，作为 Docker 项目早期的重要贡献者，RedHat 也是因为对 Docker 公司平台化战略不满而愤愤退出。但此时，它竟只剩下 OpenShift 这个跟 Cloud Foundry 同时代的经典 PaaS 一张牌可以打，跟 Docker Swarm 和转型后的 Mesos 完全不在同一个“竞技水平”之上。

 2014 年注定是一个神奇的年份。就在这一年的 6 月，基础设施领域的翘楚 Google 公司突然发力，正式宣告了一个名叫 Kubernetes 项目的诞生。而这个项目，不仅挽救了当时的 CoreOS 和 RedHat，还如同当年 Docker 项目的横空出世一样，再一次改变了整个容器市场的格局。 



### 4、尘埃落定

容器领域的其他几位玩家开始商议“切割”Docker 项目的话语权。而“切割”的手段也非常经典，那就是成立一个中立的基金会。

于是，2015 年 6 月 22 日，由 Docker 公司牵头，CoreOS、Google、RedHat 等公司共同宣布，Docker 公司将 Libcontainer 捐出，并改名为 RunC 项目，交由一个完全中立的基金会管理，然后以 RunC 为依据，大家共同制定一套容器和镜像的标准和规范。

这套标准和规范，就是 OCI（ Open Container Initiative ）。**OCI 的提出，意在将容器运行时和镜像的实现从 Docker 项目中完全剥离出来**。



所以这次，Google、RedHat 等开源基础设施领域玩家们，共同牵头发起了一个名为 CNCF（Cloud Native Computing Foundation）的基金会。这个基金会的目的其实很容易理解：它希望，以 Kubernetes 项目为基础，建立一个由开源基础设施领域厂商主导的、按照独立基金会方式运营的平台级社区，来对抗以 Docker 公司为核心的容器商业生态。 

 CNCF 社区就需要至少确保两件事情： 

1. Kubernetes 项目必须能够在容器编排领域取得足够大的竞争优势；
2. CNCF 社区必须以 Kubernetes 项目为核心，覆盖足够多的场景。

Kubernetes 项目的基础特性，并不是几个工程师突然“拍脑袋”想出来的东西，而是 Google 公司在容器化基础设施领域多年来实践经验的沉淀与升华。这，正是 Kubernetes 项目能够从一开始就避免同 Swarm 和 Mesos 社区同质化的重要手段。 

Kubernetes 项目并没有跟 Swarm 项目展开同质化的竞争，所以“Docker Native”的说辞并没有太大的杀伤力。相反地，Kubernetes 项目让人耳目一新的设计理念和号召力，很快就构建出了一个与众不同的容器编排与管理的生态。 

 在 2016 年，Docker 公司宣布了一个震惊所有人的计划：放弃现有的 Swarm 项目，将容器编排和集群管理功能全部内置到 Docker 项目当中。而**Kubernetes 的应对策略则是反其道而行之，开始在整个社区推进“民主化”架构**，即：从 API 到容器运行时的每一层，Kubernetes 项目都为开发者暴露出了可以扩展的插件机制，鼓励用户通过代码的方式介入到 Kubernetes 项目的每一个阶段。 



Kubernetes 项目的这个变革的效果立竿见影，很快在整个容器社区中催生出了大量的、基于 Kubernetes API 和扩展接口的二次创新工作，比如：

- 目前热度极高的微服务治理项目 Istio；
- 被广泛采用的有状态应用部署框架 Operator；
- 还有像 Rook 这样的开源创业项目，它通过 Kubernetes 的可扩展接口，把 Ceph 这样的重量级产品封装成了简单易用的容器存储插件。

从 2017 年开始，Docker 公司先是将 Docker 项目的容器运行时部分 Containerd 捐赠给 CNCF 社区，标志着 Docker 项目已经全面升级成为一个 PaaS 平台；紧接着，Docker 公司宣布将 Docker 项目改名为 Moby，然后交给社区自行维护，而 Docker 公司的商业产品将占有 Docker 这个注册商标。 



2017 年 10 月，Docker 公司出人意料地宣布，将在自己的主打产品 Docker 企业版中内置 Kubernetes 项目，这标志着持续了近两年之久的“编排之争”至此落下帷幕。

2018 年 1 月 30 日，RedHat 宣布斥资 2.5 亿美元收购 CoreOS。

2018 年 3 月 28 日，这一切纷争的始作俑者，Docker 公司的 CTO Solomon Hykes 宣布辞职，曾经纷纷扰扰的容器技术圈子，到此尘埃落定。



容器本身没有价值，有价值的是“容器编排”。 

开源才是王道：以开发者为核心，构建一个相对民主和开放的容器生态。 



## 03-容器技术概念入门篇 

### 5、白话容器基础（一）：从进程说开去

**一个程序运起来后的计算机执行环境的总和，就是我们今天的主角：进程**。

对于进程来说，它的静态表现就是程序，平常都安安静静地待在磁盘上；而一旦运行起来，它就变成了计算机里的数据和状态的总和，这就是它的动态表现。 



而**容器技术的核心功能，就是通过约束和修改进程的动态表现，从而为其创造出一个“边界”。** 

对于 Docker 等大多数 Linux 容器来说，**Cgroups 技术**是用来制造约束的主要手段，而**Namespace 技术**则是用来修改进程视图的主要方法。 



**除了 PID Namespace，Linux 操作系统还提供了 Mount、UTS、IPC、Network 和 User 这些 Namespace，用来对各种不同的进程上下文进行“障眼法”操作。** 



所以，Docker 容器这个听起来玄而又玄的概念，实际上是在创建容器进程时，指定了这个进程所需要启用的一组 Namespace 参数。这样，容器就只能“看”到当前 Namespace 所限定的资源、文件、设备、状态，或者配置。而对于宿主机以及其他不相关的程序，它就完全看不到了。

**所以说，容器，其实是一种特殊的进程而已。**

### 6、白话容器基础（二）：隔离与限制

**“敏捷”和“高性能”是容器相较于虚拟机最大的优势，也是它能够在 PaaS 这种更细粒度的资源管理平台上大行其道的重要原因。** 

根据实验，一个运行着 CentOS 的 KVM 虚拟机启动后，在不做优化的情况下，虚拟机自己就需要占用 100~200 MB 内存。此外，用户应用运行在虚拟机里面，它对宿主机操作系统的调用就不可避免地要经过虚拟化软件的拦截和处理，这本身又是一层性能损耗，尤其对计算资源、网络和磁盘 I/O 的损耗非常大。

而相比之下，容器化后的用户应用，却依然还是一个宿主机上的普通进程，这就意味着这些因为虚拟化而带来的性能损耗都是不存在的；而另一方面，使用 Namespace 作为隔离手段的容器并不需要单独的 Guest OS，这就使得容器额外的资源占用几乎可以忽略不计。



Docker容器相对于虚拟机的缺点：

首先，既然容器只是运行在宿主机上的一种特殊的进程，那么多个容器之间使用的就还是同一个宿主机的操作系统内核。 

其次，在 Linux 内核中，有很多资源和对象是不能被 Namespace 化的，最典型的例子就是：时间。 

 后续会讲到的基于虚拟化或者独立内核技术的容器实现，则可以比较好地在隔离与性能之间做出平衡。 



**Linux Cgroups 就是 Linux 内核中用来为进程设置资源限制的一个重要功能。Linux Cgroups 的全称是 Linux Control Group。它最主要的作用，就是限制一个进程组能够使用的资源上限，包括 CPU、内存、磁盘、网络带宽等等。** 



**Linux Cgroups 的设计还是比较易用的，简单粗暴地理解呢，它就是一个子系统目录加上一组资源限制文件的组合**。



一个正在运行的 Docker 容器，其实就是一个启用了多个 Linux Namespace 的应用进程，而这个进程能够使用的资源量，则受 Cgroups 配置的限制。

这也是容器技术中一个非常重要的概念，即：**容器是一个“单进程”模型。** 



Linux 下的 /proc 目录存储的是记录当前内核运行状态的一系列特殊文件，用户可以通过访问这些文件，查看系统以及当前正在运行的进程的信息，比如 CPU 使用情况、内存占用率等，这些文件也是 top 指令查看系统信息的主要数据来源。 

但是，你如果在容器里执行 top 指令，就会发现，它显示的信息居然是宿主机的 CPU 和内存数据，而不是当前容器的数据。 

造成这个问题的原因就是，/proc 文件系统并不知道用户通过 Cgroups 给这个容器做了什么样的资源限制，即：/proc 文件系统不了解 Cgroups 限制的存在。 



### 7、白话容器基础（三）：深入理解容器镜像

在 Linux 操作系统里，有一个名为 chroot 的命令可以帮助你在 shell 中方便地完成这个工作。顾名思义，它的作用就是帮你“change root file system”，即改变进程的根目录到你指定的位置。它的用法也非常简单。 

使用chroot实现当前进程与宿主机的目录隔离。chroot 是把某个目录修改为根目录，从而无法访问外部的内容。

**实际上，Mount Namespace 正是基于对 chroot 的不断改良才被发明出来的，它也是 Linux 操作系统里的第一个 Namespace。** 



**而这个挂载在容器根目录上、用来为容器进程提供隔离后执行环境的文件系统，就是所谓的“容器镜像”。它还有一个更为专业的名字，叫作：rootfs（根文件系统）。** 



现在，应该可以理解，对 Docker 项目来说，它最核心的原理实际上就是为待创建的用户进程：

1. 启用 Linux Namespace 配置；
2. 设置指定的 Cgroups 参数；
3. 切换进程的根目录（Change Root）。

另外，**需要明确的是，rootfs 只是一个操作系统所包含的文件、配置和目录，并不包括操作系统内核。在 Linux 操作系统中，这两部分是分开存放的，操作系统只有在开机启动时才会加载指定版本的内核镜像。** 

所以，rootfs 只包括了操作系统的“躯壳”，并没有包括操作系统的“灵魂” 。实际上，同一台机器上的所有容器，都共享宿主机操作系统的内核。 

**正是由于 rootfs 的存在，容器才有了一个被反复宣传至今的重要特性：一致性。** 

**这种深入到操作系统级别的运行环境一致性，打通了应用在本地开发和远端执行环境之间难以逾越的鸿沟。** 



但有了容器之后，更准确地说，有了容器镜像（即 rootfs）之后，这个问题【由于云端与本地服务器环境不同，应用的打包过程，一直是使用 PaaS 时最“痛苦”的一个步骤 】被非常优雅地解决了。

**由于 rootfs 里打包的不只是应用，而是整个操作系统的文件和目录，也就意味着，应用以及它运行所需要的所有依赖，都被封装在了一起。**



Docker 公司在实现 Docker 镜像时并没有沿用以前制作 rootfs 的标准流程，而是做了一个小小的创新：Docker 在镜像的设计中，引入了层（layer）的概念。也就是说，用户制作镜像的每一步操作，都会生成一个层，也就是一个增量 rootfs。 这个设计 用到了一种叫作联合文件系统（Union File System）的能力。 

UnionFS，最主要的功能是将多个不同位置的目录联合挂载（union mount）到同一个目录下 。



通过结合使用 Mount Namespace 和 rootfs，容器就能够为进程构建出一个完善的文件系统隔离环境。当然，这个功能的实现还必须感谢 chroot 和 pivot_root 这两个系统调用切换进程根目录的能力。 



 更重要的是，一旦容器镜像被发布，那么你在全世界的任何一个地方下载这个镜像，得到的内容都完全一致，可以完全复现这个镜像制作者当初的完整环境。这，就是容器技术“强一致性”的重要体现。 

因此，容器镜像必将会成为未来软件的主流发布方式。 



总结：docker的核心原理

一、使用Namespaces做主机名、网络、pid等资源的隔离；
二、使用Control Groups对进程、进程组做资源的限制；
三、使用Union FileSystem用来做镜像构建和容器运行环境等。

### 8、白话容器基础（四）：重新认识Docker容器

相较于之前介绍的制作 rootfs 的过程，Docker 提供了一种更便捷的方式，叫作 Dockerfile。 **Dockerfile 的设计思想，是使用一些标准的原语（即大写高亮的词语），描述我们所要构建的 Docker 镜像。并且这些原语，都是按顺序处理的。** 

```dockerfile
# 使用官方提供的 Python 开发镜像作为基础镜像
FROM python:2.7-slim
# 将工作目录切换为 /app
WORKDIR /app
# 将当前目录下的所有内容复制到 /app 下
ADD . /app
# 使用 pip 命令安装这个应用所需要的依赖
RUN pip install --trusted-host pypi.python.org -r requirements.txt
# 允许外界访问容器的 80 端口
EXPOSE 80
# 设置环境变量
ENV NAME World
# 设置容器进程为：python app.py，即：这个 Python 应用的启动命令
CMD ["python", "app.py"]
```

**我们统一称 Docker 容器的启动进程为 ENTRYPOINT，而不是 CMD【 即ENTRYPOINT CMD 】。** 



**一个进程，可以选择加入到某个进程已有的 Namespace 当中，从而达到“进入”这个进程所在容器的目的，这正是 docker exec 的实现原理。**而这个操作所依赖的，乃是一个名叫 setns() 的 Linux 系统调用。 



Docker Volume



Docker Volume 要解决的问题：**Volume 机制，允许你将宿主机上指定的目录或者文件，挂载到容器里面进行读取和修改操作。** 



挂载原理：只需要在 rootfs 准备好之后，在执行 chroot 之前，把 Volume 指定的宿主机目录（比如 /home 目录），挂载到指定的容器目录（比如 /test 目录）在宿主机上对应的目录（即 /var/lib/docker/aufs/mnt/[可读写层 ID]/test）上，这个 Volume 的挂载工作就完成了。 

更重要的是，由于执行这个挂载操作时，“容器进程”已经创建了，也就意味着此时 Mount Namespace 已经开启了。所以，这个挂载事件只在这个容器里可见。你在宿主机上，是看不见容器内部的这个挂载点的。这就**保证了容器的隔离性不会被 Volume 打破**。这里提到的 " 容器进程 "，是 Docker 创建的一个容器初始化进程 (dockerinit)，而不是应用进程 (ENTRYPOINT + CMD)。dockerinit 会负责完成根目录的准备、挂载设备和目录、配置 hostname 等一系列需要在容器内进行的初始化操作。最后，它通过 execv() 系统调用，让应用进程取代自己，成为容器里的 PID=1 的进程。 

容器 Volume 里的信息，并不会被 docker commit 提交掉；但这个挂载点目录 /test 本身，则会出现在新的镜像当中。 

### 9、从容器到容器云：谈谈Kubernetes的本质

一个“容器”，实际上是一个由 Linux Namespace、Linux Cgroups 和 rootfs 三种技术构建出来的进程的隔离环境。

从这个结构中我们不难看出，一个正在运行的 Linux 容器，其实可以被“一分为二”地看待：

1. 一组联合挂载在 /var/lib/docker/aufs/mnt 上的 rootfs，这一部分我们称为“容器镜像”（Container Image），是容器的静态视图；
2. 一个由 Namespace+Cgroups 构成的隔离环境，这一部分我们称为“容器运行时”（Container Runtime），是容器的动态视图。

从一个开发者和单一的容器镜像，到无数开发者和庞大的容器集群，容器技术实现了从“容器”到“容器云”的飞跃，标志着它真正得到了市场和生态的认可。

这样，**容器就从一个开发者手里的小工具，一跃成为了云计算领域的绝对主角；而能够定义容器组织和管理规范的“容器编排”技术，则当仁不让地坐上了容器技术领域的“头把交椅”。**

Borg 系统，一直以来都被誉为 Google 公司内部最强大的“秘密武器” ，Borg 项目当仁不让地位居整个基础设施技术栈的最底层。 正是由于这样的定位，Borg 可以说是 Google 最不可能开源的一个项目。而幸运地是，得益于 Docker 项目和容器技术的风靡，它却终于得以以另一种方式与开源社区见面，这个方式就是 Kubernetes 项目。 

 Kubernetes 项目在 Borg 体系的指导下，体现出了一种独有的“先进性”与“完备性”，而这些特质才是一个基础设施领域开源项目赖以生存的核心价值。 

![1690691147879](C:\Users\HP\AppData\Roaming\Typora\typora-user-images\1690691147879.png)

**Kubernetes 项目的架构**，跟它的原型项目 Borg 非常类似，都由 Master 和 Node 两种节点组成，而这两种角色分别对应着控制节点和计算节点。其中，控制节点，即 Master 节点，由三个紧密协作的独立组件组合而成，它们分别是负责 API 服务的 kube-apiserver、负责调度的 kube-scheduler，以及负责容器编排的 kube-controller-manager。整个集群的持久化数据，则由 kube-apiserver 处理后保存在 Etcd 中。

而计算节点上最核心的部分，则是一个叫作 **kubelet** 的组件。
**在 Kubernetes 项目中，kubelet 主要负责同容器运行时（比如 Docker 项目）打交道**。而这个交互所依赖的，是一个称作 CRI（Container Runtime Interface）的远程调用接口，这个接口定义了容器运行时的各项核心操作，比如：启动一个容器需要的所有参数。 

而**kubelet 的另一个重要功能，则是调用网络插件和存储插件为容器配置网络和持久化存储**。分别是 CNI（Container Networking Interface）和 CSI（Container Storage Interface）。 



**从一开始，Kubernetes 项目就没有像同时期的各种“容器云”项目那样，把 Docker 作为整个架构的核心，而仅仅把它作为最底层的一个容器运行时实现。** 



**Kubernetes 项目最主要的设计思想是，从更宏观的角度，以统一的方式来定义任务之间的各种关系，并且为将来支持更多种类的关系留有余地。** 



K8s里**Service 服务的主要作用，就是作为 Pod 的代理入口（Portal），从而代替 Pod 对外暴露一个固定的网络地址**。

这样，对于 Web 应用的 Pod 来说，它需要关心的就是数据库 Pod 的 Service 信息。不难想象，Service 后端真正代理的 Pod 的 IP 地址、端口等信息的自动更新、维护，则是 Kubernetes 项目的职责。



 Kubernetes 项目核心功能的“全景图” 

![1690691111836](C:\Users\HP\AppData\Roaming\Typora\typora-user-images\1690691111836.png)

在 Kubernetes 项目中，所推崇的使用方法是：

- 首先，通过一个“编排对象”，比如 Pod、Job、CronJob 等，来描述你试图管理的应用；
- 然后，再为它定义一些“服务对象”，比如 Service、Secret、Horizontal Pod Autoscaler（自动水平扩展器）等。这些对象，会负责具体的平台级功能。

**这种使用方法，就是所谓的“声明式 API”。这种 API 对应的“编排对象”和“服务对象”，都是 Kubernetes 项目中的 API 对象（API Object）。**



Kubernetes 项目如何启动一个容器化任务呢？ 

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
  labels:
    app: nginx
spec:
  replicas: 2
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.7.9
        ports:
        - containerPort: 80
```

实际上，过去很多的集群管理项目（比如 Yarn、Mesos，以及 Swarm）所擅长的，都是把一个容器，按照某种规则，放置在某个最佳节点上运行起来。这种功能，我们称为“调度”。

而 Kubernetes 项目所擅长的，是按照用户的意愿和整个系统的规则，完全自动化地处理好容器之间的各种关系。**这种功能，就是我们经常听到的一个概念：编排。**

所以说，Kubernetes 项目的本质，是为用户提供一个具有普遍意义的容器编排工具。

不过，更重要的是，Kubernetes 项目为用户提供的不仅限于一个工具。它真正的价值，乃在于**提供了一套基于容器构建分布式系统的基础依赖。**



## 04-Kubernetes集群搭建与实践

### 10、Kubernetes一键部署利器：kubeadm

Kubernetes 的架构和它的组件。在部署时，它的每一个组件都是一个需要被执行的、单独的二进制文件。 



kubelet 是 Kubernetes 项目用来操作 Docker 等容器运行时的核心组件。可是，除了跟容器运行时打交道外，kubelet 在配置容器网络、管理容器数据卷时，都需要直接操作宿主机。 所以需要把 kubelet 直接运行在宿主机上，然后使用容器部署其他的 Kubernetes 组件。 
所以，使用 kubeadm 的第一步，是在机器上手动安装 kubeadm、kubelet 和 kubectl 这三个二进制文件。 

kubeadm 的工作流程：【 kubeadm 几乎完全是一位高中生的作品。他叫 Lucas Käldström，芬兰人，今年只有 18 岁。kubeadm，是他 17 岁时用业余时间完成的一个社区项目。】

执行 kubeadm init 指令 。

1、**kubeadm 首先要做的，是一系列的检查工作，以确定这台机器可以用来部署 Kubernetes**。这一步检查，我们称为“Preflight Checks”。

2、**在通过了 Preflight Checks 之后，kubeadm 要为你做的，是生成 Kubernetes 对外提供服务所需的各种证书和对应的目录。**  /etc/kubernetes/pki/ca.{crt,key} 

3、**证书生成后，kubeadm 接下来会为其他组件生成访问 kube-apiserver 所需的配置文件**。这些文件的路径是：/etc/kubernetes/xxx.conf： 

4、**接下来，kubeadm 会为 Master 组件生成 Pod 配置文件**。位于/etc/kubernetes/manifests/【master节点的kube-apiserver、kube-controller-manager、kube-scheduler以及etcd，而它们都会被使用 Pod 的方式部署起来 】【**在 Kubernetes 中，有一种特殊的容器启动方法叫做“Static Pod”【它的启动不依赖与k8s】。它允许你把要部署的 Pod 的 YAML 文件放在一个指定的目录里。这样，当这台机器上的 kubelet 启动时，它会自动检查这个目录，加载所有的 Pod YAML 文件，然后在这台机器上启动它们。从这一点也可以看出，kubelet 在 Kubernetes 项目中的地位非常高，在设计上它就是一个完全独立的组件，而其他 Master 组件，则更像是辅助性的系统容器。**】

5、**kubeadm 就会为集群生成一个 bootstrap token**。 用于其它节点进行 kubeadm join。【kubeadm 至少需要发起一次“不安全模式”的访问到 kube-apiserver，从而拿到保存在 ConfigMap 中的 cluster-info（它保存了 APIServer 的授权信息）。而 bootstrap token，扮演的就是这个过程中的安全验证的角色。 】

6、 **kubeadm init 的最后一步，就是安装默认插件**。Kubernetes 默认 kube-proxy 和 DNS 这两个插件是必须安装的。 

### 11、从0到1：搭建一个完整的Kubernetes集群

从0开始安装K8s1.25：https://blog.csdn.net/qq_41822345/article/details/126679925

容器网络插件：Flannel
可视化插件DashBoard
容器持久化插件Rook【它巧妙地依赖了 Kubernetes 提供的编排能力，合理的使用了很多诸如 Operator、CRD 等重要的扩展特性 】

存储插件会在容器里挂载一个基于网络或者其他机制的远程数据卷，使得在容器里创建的文件，实际上是保存在远程存储服务器上，或者以分布式的方式保存在多个节点上，而与当前宿主机没有任何绑定关系。这样，无论你在其他哪个宿主机上启动新的容器，都可以请求挂载指定的持久化存储卷，从而访问到数据卷里保存的内容。**这就是“持久化”的含义。** 



另外，k8s集群的部署过程并不像传说中那么繁琐，这主要得益于：

1. kubeadm 项目大大简化了部署 Kubernetes 的准备工作，尤其是配置文件、证书、二进制文件的准备和制作，以及集群版本管理等操作，都被 kubeadm 接管了。
2. Kubernetes 本身“一切皆容器”的设计思想，加上良好的可扩展机制，使得插件的部署非常简便。

上述思想，也是开发和使用 Kubernetes 的重要指导思想，即：基于 Kubernetes 开展工作时，你一定要优先考虑这两个问题：

1. 我的工作是不是可以容器化？
2. 我的工作是不是可以借助 Kubernetes API 和可扩展机制来完成？

而一旦这项工作能够基于 Kubernetes 实现容器化，就很有可能像上面的部署过程一样，大幅简化原本复杂的运维工作。对于时间宝贵的技术人员来说，这个变化的重要性是不言而喻的。

- 思考题

1. 你是否使用其他工具部署过 Kubernetes 项目？经历如何？
2. 你是否知道 Kubernetes 项目当前（v1.11）能够有效管理的集群规模是多少个节点？你在生产环境中希望部署或者正在部署的集群规模又是多少个节点呢？



### 12、牛刀小试：我的第一个容器化应用

一个 YAML 文件，对应到 Kubernetes 中，就是一个 API Object（API 对象）。当你为这个对象的各个字段填好值并提交给 Kubernetes 之后，Kubernetes 就会负责创建出这些对象所定义的容器或者其他类型的 API 资源。 

Kubernetes 里“最小”的 API 对象是 Pod。Pod 可以等价为一个应用，所以，Pod 可以由多个紧密协作的容器组成。由于“最小”，所以它往往都是被其他对象控制的。这种组合方式，正是 Kubernetes 进行容器编排的重要模式。 

```yaml
apiVersion: apps/v1
kind: Deployment    #API 对象的类型（Type），是一个 Deployment。
metadata:
  name: nginx-deployment
spec:
  selector:
    matchLabels:  #Label Selector。
      app: nginx 
  replicas: 2
  template:
    metadata:
      labels:
        app: nginx #Deployment控制器对象，通过Labels字段从Kubernetes中过滤出它关心的被控制对象。
    spec:
      containers:   #Pod 就是 Kubernetes 世界里的“应用”；而一个应用，可以由多个容器组成。
      - name: nginx
        image: nginx:1.7.9
        ports:
        - containerPort: 80
        volumeMounts:
        - mountPath: "/usr/share/nginx/html"
          name: nginx-vol
        volumes:
      	- name: nginx-vol2
          hostPath: 
           path: /var/data  
      volumes:       
      - name: nginx-vol  
        emptyDir: {}      #不显式声明宿主机目录的 Volume。[创建临时目录]
      - name: nginx-vol2
        hostPath:         #显式声明宿主机目录的 Volume。
          path: /apps/logs/qfusion/cloud-monitor/
          type: DirectoryOrCreate
```

像这样使用一种 API 对象（Deployment）管理另一种 API 对象（Pod）的方法，在 Kubernetes 中，叫作“控制器”模式（controller pattern）。 

```shell
kubectl create -f nginx-deployment.yaml 
kubectl replace -f nginx-deployment.yaml 
kubectl apply -f nginx-deployment.yaml 
```

**当应用本身发生变化时，开发人员和运维人员可以依靠容器镜像来进行同步；当应用部署参数发生变化时，这些 YAML 文件就是他们相互沟通和信任的媒介。** 



## 05-容器编排与Kubernetes作业管理

### 13、为什么我们需要Pod？

容器，就是未来云计算系统中的进程；容器镜像就是这个系统里的“.exe”安装包。

Pod，是 Kubernetes 项目中最小的 API 对象。 Pod，是 Kubernetes 项目的原子调度单位。 

Kubernetes 就是操作系统！ 

- 为什么需要Pod？——**容器设计模式**

1、用于解决具有”超亲密关系容器“【进程[容器]与进程组[Pod]】的调度问题。Pod 是 Kubernetes 里的原子调度单位。这就意味着，Kubernetes 项目的调度器，是统一按照 Pod 而非容器的资源需求进行计算的。 

2、**容器设计模式**。 Pod只是一个逻辑概念。Pod，其实是一组共享了某些资源的容器。**Pod 里的所有容器，共享的是同一个 Network Namespace，并且可以声明共享同一个 Volume。** 

**为了保证一个Pod中容器之间的对等关系**，在 Kubernetes 项目里，Pod 的实现需要使用一个中间容器，这个容器叫作 Infra 容器。在这个 Pod 中，Infra 容器永远都是第一个被创建的容器，而其他用户定义的容器，则通过 Join Network Namespace 的方式，与 Infra 容器关联在一起。 



而对于同一个 Pod 里面的所有用户容器来说，它们的进出流量，也可以认为都是通过 Infra 容器完成的。这一点很重要，因为**将来如果你要为 Kubernetes 开发一个网络插件时，应该重点考虑的是如何配置这个 Pod 的 Network Namespace，而不是每一个用户容器如何使用你的网络配置，这是没有意义的。** 

对于Volume，Kubernetes 项目只要把所有 Volume 的定义都设计在 Pod 层级即可。 

Pod 这种“超亲密关系”容器的设计思想，实际上就是希望，当用户想在一个容器里跑多个功能并不相关的应用时，应该优先考虑它们是不是更应该被描述成一个 Pod 里的多个容器。 

**第一个最典型的例子是：WAR 包与 Web 服务器。**我们知道对于一个java的 WAR 包应用，它需要被放在 Tomcat 的 webapps 目录下运行起来。那么如何更新应用——可以把 WAR 包和 Tomcat 分别做成镜像，然后把它们作为一个 Pod 里的两个容器“组合”在一起。通过两个容器挂载同一个目录，解决 WAR 包与 Tomcat 容器之间耦合关系的问题。 如下：

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: javaweb-2
spec:
  initContainers:  # Init Container 定义的容器都会最先启动
  - image: geektime/sample:v2
    name: war
    command: ["cp", "/sample.war", "/app"]
    volumeMounts:
    - mountPath: /app
      name: app-volume
  containers:    # Tomcat 容器启动时，它的 webapps 目录下就一定会存在 sample.war 文件
  - image: geektime/tomcat:7.0
    name: tomcat
    command: ["sh","-c","/root/apache-tomcat-7.0.42-v2/bin/start.sh"]
    volumeMounts:
    - mountPath: /root/apache-tomcat-7.0.42-v2/webapps
      name: app-volume
    ports:
    - containerPort: 8080
      hostPort: 8001 
  volumes:
  - name: app-volume
    emptyDir: {}
```

这个所谓的“组合”操作，正是容器设计模式里最常用的一种模式，它的名字叫：sidecar。**sidecar 指的就是我们可以在一个 Pod 中，启动一个辅助容器，来完成一些独立于主进程（主容器）之外的工作。** 

**第二个例子，则是容器的日志收集。**比如现在有一个应用，需要不断地把日志文件输出到容器的 /var/log 目录中。通过使用共享的 Volume 来完成对文件的操作。 

Pod 的另一个重要特性是，它的所有容器都共享同一个 Network Namespace。这就使得很多与 Pod 网络相关的配置和管理，也都可以交给 sidecar 完成，而完全无须干涉用户容器。这里最典型的例子莫过于 Istio 这个微服务治理项目。

可以这么理解 Pod 的本质：

> Pod，实际上是在扮演传统基础设施里“虚拟机”的角色；而容器，则是这个虚拟机里运行的用户程序。

Pod 这个概念，提供的是一种编排思想，而不是具体的技术方案。 

容器化，这个“上云”工作的完成，最终还是要靠深入理解容器的本质，即：进程。 

当你需要把一个运行在虚拟机里的应用迁移到 Docker 容器中时，一定要仔细分析到底有哪些进程（组件）运行在这个虚拟机里。

然后，你就可以把整个虚拟机想象成为一个 Pod，把这些进程分别做成容器镜像，把有顺序关系的容器，定义为 Init Container。这才是更加合理的、松耦合的容器编排诀窍，也是从传统应用架构，到“微服务架构”最自然的过渡方式。

### 14、深入解析Pod对象（一）：基本概念

一定要理解Pod 扮演的是传统部署环境里“虚拟机”的角色。这样的设计，是为了使用户从传统环境（虚拟机环境）向 Kubernetes（容器环境）的迁移，更加平滑。 

**凡是调度、网络、存储，以及安全相关的属性，基本上是 Pod 级别的。**这些属性的共同特征是，它们描述的是“机器”这个整体，而不是里面运行的“程序”。比如，配置这个“机器”的网卡（即：Pod 的网络定义），配置这个“机器”的磁盘（即：Pod 的存储定义），配置这个“机器”的防火墙（即：Pod 的安全定义）。更不用说，这台“机器”运行在哪个服务器之上（即：Pod 的调度）。 

**凡是跟容器的 Linux Namespace 相关的属性，也一定是 Pod 级别的**。这个原因也很容易理解：Pod 的设计，就是要让它里面的容器尽可能多地共享 Linux Namespace，仅保留必要的隔离和限制能力。 

**凡是 Pod 中的容器要共享宿主机的 Namespace，也一定是 Pod 级别的定义** 

```shell
apiVersion: v1
kind: Pod
metadata:
  name: nginx
spec:
  hostNetwork: true
  hostIPC: true
  hostPID: true
  containers:
  - name: nginx
    image: nginx
  - name: shell
    image: busybox
    stdin: true
    tty: true
  - name: lifecycle-demo-container
    image: nginx
    lifecycle:      # Lifecycle 字段
      postStart:    # 在容器启动后，立刻执行一个指定的操作，但它并不严格保证顺序。即在postStart启动时，ENTRYPOINT有可能还没有结束。
        exec:
          command: ["/bin/sh", "-c", "echo Hello from the postStart handler > /usr/share/message"]
      preStop:     # 同步，它会阻塞当前的容器杀死流程，直到这个 Hook 定义操作完成之后，才允许容器被杀死，这跟 postStart 不一样。容器被删除之前，先调用了 nginx 的退出指令（即 preStop 定义的操作），从而实现了容器的“优雅退出”。
        exec:
          command: ["/usr/sbin/nginx","-s","quit"]
```

在这个 Pod 中，我定义了共享宿主机的 Network、IPC 和 PID Namespace。这就意味着，这个 Pod 里的所有容器，会直接使用宿主机的网络、直接与宿主机进行 IPC 通信、看到宿主机里正在运行的所有进程。 

**Pod 对象在 Kubernetes 中的生命周期**。 

Pending 、 Running 、 Succeeded 、 Failed 、 Unknown



### 15、深入解析Pod对象（二）：使用进阶

k8s中有一些特殊的Volume， 这种特殊的 Volume，叫作 Projected Volume，可以把它翻译为“投射数据卷”。这些特殊 Volume 的作用，是为容器提供预先定义好的数据。所以，从容器的角度来看，这些 Volume 里的信息就是仿佛是**被 Kubernetes“投射”（Project）进入容器当中的**。 

到目前为止，Kubernetes 支持的 Projected Volume 一共有四种：

Secret；ConfigMap；Downward API【 它的作用是：让 Pod 里的容器能够直接获取到这个 Pod API 对象本身的信息。 】；ServiceAccountToken。

比如：

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: test-projected-volume 
spec:
  containers:
  - name: test-secret-volume
    image: busybox
    args:
    - sleep
    - "86400"
    volumeMounts:
    - name: mysql-cred
      mountPath: "/projected-volume"
      readOnly: true
  volumes:
  - name: mysql-cred
    projected:      # 这里声明的Volume，并不是emptyDir或者hostPath类型，而是projected类型。
      sources:
      - secret:
          name: user   #kubectl create secret generic user --from-file=./username.txt
      - secret:
          name: pass   #kubectl create secret generic pass --from-file=./password.txt
---
apiVersion: v1
kind: Secret
metadata:
  name: mysecret
type: Opaque
data:
  user: YWRtaW4=   # Secret 对象要求这些数据必须是经过 Base64 转码的
  pass: MWYyZDFlMmU2N2Rm # 在真正的生产环境中，需要在 Kubernetes 中开启 Secret 的加密插件，增强数据的安全性
```



ServiceAccountToken为了方便使用，Kubernetes 已经为你提供了一个的默认“服务账户”（default Service Account）。并且，任何一个运行在 Kubernetes 里的 Pod，都可以直接使用这个默认的 Service Account，而无需显示地声明挂载它 

**这种把 Kubernetes 客户端以容器的方式运行在集群里，然后使用 default Service Account 自动授权的方式，被称作“InClusterConfig”，也是我最推荐的进行 Kubernetes API 编程的授权方式。** 



 Pod 的另一个重要的配置：容器健康检查和恢复机制。 

```yaml
apiVersion: v1
kind: Pod
metadata:
  labels:
    test: liveness
  name: test-liveness-exec
spec:
  containers:
  - name: liveness
    image: busybox
    args:
    - /bin/sh
    - -c
    - touch /tmp/healthy; sleep 30; rm -rf /tmp/healthy; sleep 600
    livenessProbe:    # 健康检查
      exec:
        command:
        - cat
        - /tmp/healthy
      initialDelaySeconds: 5
      periodSeconds: 5
    readinessProbe:
      xxxx
      # readinessProbe 检查结果的成功与否，决定的这个 Pod 是不是能被通过 Service 的方式访问到，而并不影响 Pod 的生命周期。
```

Pod 的 Spec 部分的一个标准字段（pod.spec.restartPolicy），默认值是 Always，即：任何时候这个容器发生了异常，它一定会被重新创建。 

Pod 的恢复策略。除了 Always、OnFailure 、Never。

 Pod 的恢复过程，永远都是发生在当前节点上，而不会跑到别的节点上去。事实上，一旦一个 Pod 与一个节点（Node）绑定，除非这个绑定发生了变化（pod.spec.node 字段被修改），否则它永远都不会离开这个节点。这也就意味着，如果这个宿主机宕机了，这个 Pod 也不会主动迁移到其他节点上去。 



 Kubernetes 能不能自动给 Pod 填充某些字段呢？ 

所以，这个时候，运维人员就可以定义一个 PodPreset 对象。在这个对象中，凡是想在开发人员编写的 Pod 里追加的字段，都可以预先定义好。 

```yaml
apiVersion: settings.k8s.io/v1alpha1
kind: PodPreset     # PodPreset 对象
metadata:
  name: allow-database
spec:
  selector:
    matchLabels:
      role: frontend
  env:
    - name: DB_PORT
      value: "6379"
  volumeMounts:
    - mountPath: /cache
      name: cache-volume
  volumes:
    - name: cache-volume
      emptyDir: {}
```

**PodPreset 里定义的内容，只会在 Pod API 对象被创建之前追加在这个对象本身上，而不会影响任何 Pod 的控制器的定义。**比如，我们现在提交的是一个 nginx-deployment，那么这个 Deployment 对象本身是永远不会被 PodPreset 改变的，被修改的只是这个 Deployment 创建出来的所有 Pod。 

PodPreset 这样专门用来对 Pod 进行批量化、自动化修改的工具对象。 

最后认真体会一下 Kubernetes“一切皆对象”的设计思想：比如应用是 Pod 对象，应用的配置是 ConfigMap 对象，应用要访问的密码则是 Secret 对象。 

### 16、编排其实很简单：谈谈“控制器”模型

**Pod 这个看似复杂的 API 对象，实际上就是对容器的进一步抽象和封装而已。** 

实际上 kube-controller-manager 这个组件，就是一系列控制器的集合。 

这些控制器之所以被统一放在 pkg/controller 目录下，就是因为它们都遵循 Kubernetes 项目中的一个通用编排模式，即：控制循环（control loop）。伪代码表示：

```yaml
for {
  实际状态 := 获取集群中对象 X 的实际状态（Actual State）
  期望状态 := 获取集群中对象 X 的期望状态（Desired State）
  if 实际状态 == 期望状态{
    什么都不做
  } else {
    执行编排动作，将实际状态调整为期望状态
  }
}

其中：
实际状态：在具体实现中，实际状态往往来自于 Kubernetes 集群本身。
比如，kubelet 通过心跳汇报的容器状态和节点状态，或者监控系统中保存的应用监控数据，或者控制器主动收集的它自己感兴趣的信息，这些都是常见的实际状态的来源。
期望状态：一般来自于用户提交的 YAML 文件。

很明显，这些信息往往都保存在 Etcd 中。
```

控制循环的最后结果，往往都是对被控制对象的某种写操作。比如，增加 Pod，删除已有的 Pod，或者更新 Pod 的某个字段。**这也是 Kubernetes 项目“面向 API 对象编程”的一个直观体现。**

**类似 Deployment 这样的一个控制器，实际上都是由上半部分的控制器定义（包括期望状态），加上下半部分的被控制对象的模板组成的。**



```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  selector:
    matchLabels:
      app: nginx
  replicas: 2    # 控制器定义【上半部分】
  ------------------------------------------------------------------------
  template:      # 控制即PodTemplate 类似还有VolumeTemplate。被控制对象【下半部分】
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.7.9
        ports:
        - containerPort: 80
```



### 17、经典PaaS的记忆：作业副本与水平扩展

```shell
# 对当前Deployment资源的操作都会被记录下来
kubectl create -f nginx-deployment.yaml --record

kubectl scale deployment nginx-deployment --replicas=4
# 查看更新过程
kubectl rollout status deployment/nginx-deployment

kubectl set image deployment/nginx-deployment nginx=nginx:1.91
kubectl rollout undo deployment/nginx-deployment
kubectl rollout history deployment/nginx-deployment --revision=2

# 让Deployment进入“暂停”状态，任何修改不会触发“滚动更新”，也不会创建新的 ReplicaSet。
kubectl rollout pause deployment/nginx-deployment
# 
kubectl rollout resume deploy/nginx-deployment
```

Deployment 遵循一种叫作“滚动更新”（rolling update）的方式，来升级现有的容器。而这个能力的实现，依赖的是 Kubernetes 项目中的一个非常重要的概念（API 对象）：ReplicaSet。

**Deployment 控制器实际操纵的，正是这样的 ReplicaSet 对象，而不是 Pod 对象。**  

```yaml
apiVersion: apps/v1
kind: ReplicaSet
metadata:
  name: nginx-set
  labels:
    app: nginx
spec:
  replicas: 3
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.7.9
```

修改 Deployment 有很多方法。比如，可以直接使用 **kubectl edit 指令编辑 Etcd 里的 API 对象**。 kubectl set image更换镜像等。

kubectl edit 并不神秘，它不过是把 API 对象的内容下载到了本地文件，让你修改完成后再提交上去。 



**将一个集群中正在运行的多个 Pod 版本，交替地逐一升级的过程，就是“滚动更新”** 。很明显，它能使我们的服务即使在出现问题时，也不会受到太大的影响。

当然，这也就要求一定要使用 Pod 的 Health Check 机制检查应用的运行状态，而不是简单地依赖于容器的 Running 状态。要不然的话，虽然容器已经变成 Running 了，但服务很有可能尚未启动，“滚动更新”的效果也就达不到。

-  RollingUpdateStrategy 

这是为了进一步保证服务的连续性，Deployment Controller 会确保，在任何时间窗口内，只有指定比例的 Pod 处于离线状态。同时，它也会确保，在任何时间窗口内，只有指定比例的新 Pod 被创建出来。这两个比例的值都是可以配置的，默认都是 DESIRED 值的 25%。 

Deployment 对象有一个字段，叫作 spec.revisionHistoryLimit，就是 Kubernetes 为 Deployment 保留的“历史版本”个数。

Deployment 实际上是一个**两层控制器**。首先，它通过**ReplicaSet 的个数**来描述应用的版本；然后，它再通过**ReplicaSet 的属性**（比如 replicas 的值），来保证 Pod 的副本数量。 

Deployment 控制 ReplicaSet（版本），ReplicaSet 控制 Pod（副本数）。

### 18、深入理解StatefulSet（一）：拓扑状态

Deployment 对应用做了一个简单化假设。它认为，一个应用的所有 Pod，是完全一样的。所以，它们互相之间没有顺序，也无所谓运行在哪台宿主机上。需要的时候，Deployment 就可以通过 Pod 模板创建新的 Pod；不需要的时候，Deployment 就可以“杀掉”任意一个 Pod。

这种实例之间有不对等关系，以及实例对外部数据有依赖关系的应用，就被称为“有状态应用”（Stateful Application）。 

StatefulSet 的设计其实非常容易理解。它把真实世界里的应用状态，抽象为了两种情况： 
**拓扑状态**：这种情况意味着，应用的多个实例之间不是完全对等的关系。这些应用实例，必须按照某些顺序启动，比如应用的主节点 A 要先于从节点 B 启动。 
**存储状态**：最典型的例子，就是一个数据库应用的多个存储实例。 

所以，**StatefulSet 的核心功能，就是通过某种方式记录这些状态，然后在 Pod 被重新创建时，能够为新 Pod 恢复这些状态。** 

Service服务之 Headless Service。 

Service服务有两种访问方式：1、**是以 Service 的 VIP（Virtual IP，即：虚拟 IP）方式** ；2、**以 Service 的 DNS 方式**。其中第2种访问方式又有两种处理方式：a、Normal Service【即dns解析到的就是vip】；b、Headless Service【即dns解析到的就是某个pod的ip，不需要再分配vip】

```yaml
apiVersion: v1
kind: Service
metadata:
  name: nginx
  labels:
    app: nginx
spec:
  ports:
  - port: 80
    name: web
  clusterIP: None      # Headless Service
  selector:
    app: nginx
```

**StatefulSet 保证了 Pod 网络标识的稳定性**。

**通过sts，Kubernetes 成功地将 Pod 的拓扑状态（比如：哪个节点先启动，哪个节点后启动），按照 Pod 的“名字 + 编号”的方式固定了下来**。 

比如，如果 web-0 是一个需要先启动的主节点，web-1 是一个后启动的从节点，那么只要这个 StatefulSet 不被删除，你访问 web-0.nginx 时始终都会落在主节点上，访问 web-1.nginx 时，则始终都会落在从节点上，这个关系绝对不会发生任何变化。

StatefulSet 这个控制器的主要作用之一，就是使用 Pod 模板创建 Pod 的时候，对它们进行编号，并且按照编号顺序逐一完成创建工作。而当 StatefulSet 的“控制循环”发现 Pod 的“实际状态”与“期望状态”不一致，需要新建或者删除 Pod 进行“调谐”的时候，它会严格按照这些 Pod 编号的顺序，逐一完成这些操作。 



### 19、深入理解StatefulSet（二）：存储状态

为什么需要PVC？如下，这种yaml定义有什么问题？

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: rbd
spec:
  containers:
    - image: kubernetes/pause
      name: rbd-rw
      volumeMounts:
      - name: rbdpd
        mountPath: /mnt/rbd
  volumes:
    - name: rbdpd
      rbd:
        monitors:
        - '10.16.154.78:6789'
        - '10.16.154.82:6789'
        - '10.16.154.83:6789'
        pool: kube
        image: foo
        fsType: ext4
        readOnly: true
        user: admin
        keyring: /etc/ceph/keyring
        imageformat: "2"
        imagefeatures: "layering"
```

其一，如果不懂得 Ceph RBD 的使用方法，那么这个 Pod 里 Volumes 字段，你十有八九也完全看不懂。其二，这个 Ceph RBD 对应的存储服务器的地址、用户名、授权文件的位置，也都被轻易地暴露给了全公司的所有开发人员，这是一个典型的信息被“过度暴露”的例子。

这也是为什么，在后来的演化中，**Kubernetes 项目引入了一组叫作 Persistent Volume Claim（PVC）和 Persistent Volume（PV）的 API 对象，大大降低了用户声明和使用持久化 Volume 的门槛。**

Kubernetes 中 PVC 和 PV 的设计，**实际上类似于“接口”和“实现”的思想**。开发者只要知道并会使用“接口”，即：PVC；而运维人员则负责给“接口”绑定具体的实现，即：PV。这种解耦，就避免了因为向开发者暴露过多的存储系统细节而带来的隐患。此外，这种职责的分离，往往也意味着出现事故时可以更容易定位问题和明确责任，从而避免“扯皮”现象的出现。 

这种PVC、PV 的设计，也使得 StatefulSet 对存储状态的管理成为了可能。 

可以简单理解为：**PVC 其实就是一种特殊的 Volume**。只不过一个 PVC 具体是什么类型的 Volume，要在跟某个 PV 绑定之后才知道。 

StatefulSet 控制器恢复一个删除 Pod 的过程举例：
1、首先，当你把一个 Pod，比如 web-0，删除之后，这个 Pod 对应的 PVC 和 PV，并不会被删除，而这个 Volume 里已经写入的数据，也依然会保存在远程存储服务里（比如，我们在这个例子里用到的 Ceph 服务器）。 
2、此时，StatefulSet 控制器发现，一个名叫 web-0 的 Pod 消失了。所以，控制器就会重新创建一个新的、名字还是叫作 web-0 的 Pod 来，“纠正”这个不一致的情况。而在这个新的 Pod 对象的定义里，它声明使用的 PVC 的名字，还是叫作：www-web-0。这个 PVC 的定义，还是来自于 PVC 模板（volumeClaimTemplates），这是 StatefulSet 创建 Pod 的标准流程。 
3、所以，在这个新的 web-0 Pod 被创建出来之后，Kubernetes 为它查找名叫 www-web-0 的 PVC 时，就会直接找到旧 Pod 遗留下来的同名的 PVC，进而找到跟这个 PVC 绑定在一起的 PV。这样，新的 Pod 就可以挂载到旧 Pod 对应的那个 Volume，并且获取到保存在 Volume 里的数据。 

**Kubernetes 的 StatefulSet 就实现了对应用存储状态的管理。** 
sts控制器工作原理：
**首先，StatefulSet 的控制器直接管理的是 Pod**。
**其次，Kubernetes 通过 Headless Service，为这些有编号的 Pod，在 DNS 服务器中生成带有同样编号的 DNS 记录**。只要 StatefulSet 能够保证这些 Pod 名字里的编号不变，那么 Service 里类似于 web-0.nginx.default.svc.cluster.local 这样的 DNS 记录也就不会变，而这条记录解析出来的 Pod 的 IP 地址，则会随着后端 Pod 的删除和再创建而自动更新。这当然是 Service 机制本身的能力，不需要 StatefulSet 操心。 
**最后，StatefulSet 还为每一个 Pod 分配并创建一个同样编号的 PVC**。 

在这种情况下，即使 Pod 被删除，它所对应的 PVC 和 PV 依然会保留下来。所以当这个 Pod 被重新创建出来之后，Kubernetes 会为它找到同样编号的 PVC，挂载这个 PVC 对应的 Volume，从而获取到以前保存在 Volume 里的数据。 



























































