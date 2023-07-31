
- StatefulSet 可以说是 Kubernetes 中作业编排的“集大成者”。 
- 因为，几乎每一种 Kubernetes 的编排功能，都可以在编写 StatefulSet 的 YAML 文件时被用到。 
- 案列：部署一个 MySQL 集群
- 如何使用 StatefulSet 将MySQL的集群搭建过程“容器化”。
- 首先，用自然语言来描述一下我们想要部署的“有状态应用”。
1、是一个“主从复制”（Maser-Slave Replication）的 MySQL 集群；
2、有 1 个主节点（Master）；
3、有多个从节点（Slave）；
4、从节点需要能水平扩展；
5、所有的写操作，只能在主节点上执行；
6、读操作可以在所有节点上执行。



- 在常规环境里，部署这样一个主从模式的 MySQL 集群的主要难点在于：如何让从节点能够拥有主节点的数据，即：如何配置主（Master）从（Slave）节点的复制与同步。
1、第一步工作，就是通过 XtraBackup 将 Master 节点的数据备份到指定目录。
2、配置 Slave 节点。Slave 节点在第一次启动前，需要先把 Master 节点的备份数据，连同备份信息文件，一起拷贝到自己的数据目录（/var/lib/mysql）下。
3、启动 Slave 节点。“CHANGE MASTER TO” 和 “START SLAVE”命令
4、在这个集群中添加更多的 Slave 节点。
- 通过上面的叙述，我们不难看到，将部署 MySQL 集群的流程迁移到 Kubernetes 项目上，需要能够“容器化”地解决下面的“三座大山”：
1、Master 节点和 Slave 节点需要有不同的配置文件（即：不同的 my.cnf）；
2、Master 节点和 Salve 节点需要能够传输备份信息文件；
3、在 Slave 节点第一次启动之前，需要执行一些初始化 SQL 操作；

- 可以看出：MySQL 本身同时拥有拓扑状态（主从节点的区别）和存储状态（MySQL 保存在本地的数据），自然要通过 StatefulSet 来解决这“三座大山”的问题。

- 针对1、“第一座大山：Master 节点和 Slave 节点需要有不同的配置文件”，很容易处理：我们只需要给主从节点分别准备两份不同的 MySQL 配置文件，然后根据 Pod 的序号（Index）挂载进去即可。
如 mysql-cm.yaml 和 mysql-svc.yaml 文件
- 针对2、先搭建框架，再完善细节。其中，Pod 部分如何定义，是完善细节时的重点。
如 mysql-sts文件
- 针对3、

