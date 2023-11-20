- gitee个人代码：[https://gitee.com/HanFerm/gulimall](https://gitee.com/HanFerm/gulimall)
- 笔记-基础篇-1(P1-P28)：[https://blog.csdn.net/hancoder/article/details/106922139](https://blog.csdn.net/hancoder/article/details/106922139)
- 笔记-基础篇-2(P28-P100)：[https://blog.csdn.net/hancoder/article/details/107612619](https://blog.csdn.net/hancoder/article/details/107612619)
- 笔记-高级篇(P340)：[https://blog.csdn.net/hancoder/article/details/107612746](https://blog.csdn.net/hancoder/article/details/107612746)
- 笔记-vue：[https://blog.csdn.net/hancoder/article/details/107007605](https://blog.csdn.net/hancoder/article/details/107007605)
- 笔记-elastic search、上架、检索：[https://blog.csdn.net/hancoder/article/details/113922398](https://blog.csdn.net/hancoder/article/details/113922398)
- 笔记-认证服务：[https://blog.csdn.net/hancoder/article/details/114242184](https://blog.csdn.net/hancoder/article/details/114242184)
- 笔记-分布式锁与缓存：[https://blog.csdn.net/hancoder/article/details/114004280](https://blog.csdn.net/hancoder/article/details/114004280)
- 笔记-集群篇：[https://blog.csdn.net/hancoder/article/details/107612802](https://blog.csdn.net/hancoder/article/details/107612802)
- 笔记-k8s、devOps专栏：[https://blog.csdn.net/hancoder/category_11140481.html](https://blog.csdn.net/hancoder/category_11140481.html)
- springcloud笔记：[https://blog.csdn.net/hancoder/article/details/109063671](https://blog.csdn.net/hancoder/article/details/109063671)
- 笔记版本说明：2020年提供过笔记文档，但只有P1-P50的内容，2021年整理了P340的内容。请点击标题下面分栏查看系列笔记
- 声明：

  - 可以白嫖，但请勿转载发布，笔记手打不易
  - 本系列笔记不断迭代优化，csdn：hancoder上是最新版内容，10W字都是在csdn免费开放观看的。
  - 离线md笔记文件获取方式见文末。2021-3版本的md笔记打完压缩包共500k（云图床），包括本项目笔记，还有cloud、docker、mybatis-plus、rabbitMQ等个人相关笔记
- 本项目其他笔记见专栏：[https://blog.csdn.net/hancoder/category_10822407.html](https://blog.csdn.net/hancoder/category_10822407.html)



# 0、谷粒最后一篇

- Jenkins：[https://blog.csdn.net/hancoder/article/details/118233786](https://blog.csdn.net/hancoder/article/details/118233786)
- kubeSphere：[https://blog.csdn.net/hancoder/article/details/118053239](https://blog.csdn.net/hancoder/article/details/118053239)

# 一、DevOps

基础知识详见：[https://blog.csdn.net/hancoder/article/details/118233786](https://blog.csdn.net/hancoder/article/details/118233786)

- 持续集成CI：拉取代码、自动化测试等
- 持续部署CD：代码通过评审后部署到生成环境中
- 流水线：[https://www.jenkins.io/zh/doc/book/pipeline/](https://www.jenkins.io/zh/doc/book/pipeline/)
  - https://kubesphere.com.cn/docs/devops-user-guide/how-to-use/create-a-pipeline-using-jenkinsfile/
  - 因为kubeSphere有可视化界面，所以无需写JenkinsFile了。
  - 从代码库中检出代码

#### 流水线概述

本示例流水线包括以下八个阶段。

![流水线概览](https://kubesphere.com.cn/images/docs/zh-cn/devops-user-guide/use-devops/create-a-pipeline-using-a-jenkinsfile/pipeline-overview.png)

备注

- **阶段 1：Checkout SCM**：从 GitHub 仓库检出源代码。
- **阶段 2：单元测试**：待该测试通过后才会进行下一阶段。
- **阶段 3：SonarQube 分析**：SonarQube 代码质量分析。
- **阶段 4：构建并推送快照镜像**：根据**行为策略**中选定的分支来构建镜像，并将 `SNAPSHOT-$BRANCH_NAME-$BUILD_NUMBER` 标签推送至 Docker Hub，其中 `$BUILD_NUMBER` 为流水线活动列表中的运行序号。
- **阶段 5：推送最新镜像**：将 SonarQube 分支标记为 `latest`，并推送至 Docker Hub。
- **阶段 6：部署至开发环境**：将 SonarQube 分支部署到开发环境，此阶段需要审核。
- **阶段 7：带标签推送**：生成标签并发布到 GitHub，该标签会推送到 Docker Hub。
- **阶段 8：部署至生产环境**：将已发布的标签部署到生产环境。

上述内容详见：[https://kubesphere.com.cn/docs/devops-user-guide/how-to-use/create-a-pipeline-using-jenkinsfile/](https://kubesphere.com.cn/docs/devops-user-guide/how-to-use/create-a-pipeline-using-jenkinsfile/)

去kubeSphere中点击之前创建的devOps项目，添加dockerhub凭证（账号密码）、gitee凭证、kubeconfig凭证，创建图示https://kubesphere.com.cn/docs/devops-user-guide/how-to-use/credential-management/

#### 安装SonarQube 

kubeSphere v3中已无SonarQube ，自己安装：https://kubesphere.com.cn/docs/devops-user-guide/how-to-integrate/sonarqube/

```sh
# 查看k8s集群中是否有SonarQube 
kubectl get svc -n kubesphere-devops-system | grep sonarqube-sonarqube
```

```sh
helm version
# 
helm upgrade --install sonarqube sonarqube --repo https://charts.kubesphere.io/main -n kubesphere-devops-system  --create-namespace --set service.type=NodePort
# 上面命令执行出错，查阅发现是需要helm版本是3
# 安装helm3，打不开就手动创建文件
curl https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3 | bash

```

```sh
[root@k8s-node1 vagrant]# helm upgrade --install sonarqube sonarqube --repo https://charts.kubesphere.io/main -n kubesphere-devops-system  --create-namespace --set service.type=NodePort
Release "sonarqube" does not exist. Installing it now.
NAME: sonarqube
LAST DEPLOYED: Sun Oct 10 08:07:24 2021
NAMESPACE: kubesphere-devops-system
STATUS: deployed
REVISION: 1
NOTES:
1. Get the application URL by running these commands:
  export NODE_PORT=$(kubectl get --namespace kubesphere-devops-system -o jsonpath="{.spec.ports[0].nodePort}" services sonarqube-sonarqube)
  export NODE_IP=$(kubectl get nodes --namespace kubesphere-devops-system -o jsonpath="{.items[0].status.addresses[0].address}")
  echo http://$NODE_IP:$NODE_PORT
  
[root@k8s-node1 vagrant]#   export NODE_PORT=$(kubectl get --namespace kubesphere-devops-system -o jsonpath="{.spec.ports[0].nodePort}" services sonarqube-sonarqube)
[root@k8s-node1 vagrant]#   export NODE_IP=$(kubectl get nodes --namespace kubesphere-devops-system -o jsonpath="{.items[0].status.addresses[0].address}")
[root@k8s-node1 vagrant]#   echo http://$NODE_IP:$NODE_PORT
http://192.168.56.100:30276


kubectl get pod -n kubesphere-devops-system | grep sonarqube-sonarqube
```

稍微等会就好了，admin  admin登录

创建token：  gulimall-analyze 生成，然后取kubeSphere添加凭证，类型为秘密文本，复制过来 `d6ecb75473607ba5e7c60fd4444ebbeeb7091a84`

![image-20211010165505365](https://i0.hdslb.com/bfs/album/27640cb50a501a9df65a87b5c14e36e368b20e33.png)

fork项目https://gitee.com/runzexia/devops-java-sample到gitee

编辑Jenkinsfile-online，修改如下目的内容（注意我们使用的是gitee）

```
environment {
        DOCKER_CREDENTIAL_ID = 'dockerhub-id'
        GITHUB_CREDENTIAL_ID = 'gitee-id'
        KUBECONFIG_CREDENTIAL_ID = 'demo-kubeconfig'
        REGISTRY = 'docker.io'
        DOCKERHUB_NAMESPACE = 'hanferm'
        GITHUB_ACCOUNT = 'hanferm'
        APP_NAME = 'devops-java-sample'
        SONAR_CREDENTIAL_ID= 'sonar-qube'
}


省略很多。。。

sh 'git push http://$GIT_USERNAME:$GIT_PASSWORD@gitee.com/$GITHUB_ACCOUNT/devops-java-sample.git --tags --ipv4'


去掉文件中所有的-o（mvn后面的-o）
```

上面Jenkinsfile就是定义了流水线

> 注意qube里面token要选择java

登录project-admin这个账号。创建项目kubesphere-sample-dev、kubesphere-sample-prod，并邀请成员

登录project账号，点击devops-project项目，创建流水线，

![](https://gitee.com/HanFerm/image-bed/raw/master/img/20211010171215.png)



#### 动手实验

官网都有，不贴了

### 有/无状态服务

一、定义：

- 无状态服务：就是没有特殊状态的服务,各个请求对于服务器来说统一无差别处理,请求自身携带了所有服务端所需要的所有参数(服务端自身不存储跟请求相关的任何数据,不包括数据库存储信息)
- 有状态服务：与之相反,有状态服务在服务端保留之前请求的信息,用以处理当前请求,比如session等

二、如何选择：

有状态服务常常用于实现事务（并不是唯一办法，下文有另外的方案）。举一个常见的例子，在商城里购买一件商品。需要经过放入购物车、确认订单、付款等多个步骤。由于HTTP协议本身是无状态的，所以为了实现有状态服务，就需要通过一些额外的方案。比如最常见的session，将用户挑选的商品（购物车），保存到session中，当付款的时候，再从购物车里取出商品信息 。

有状态服务可以很容易地实现事务，所以也是有价值的。但是经常听到一种说法，即server要设计为无状态的，这主要是从可伸缩性来考虑的。如果server是无状态的，那么对于客户端来说，就可以将请求发送到任意一台server上，然后就可以通过负载均衡等手段，实现水平扩展。如果server是有状态的，那么就无法很容易地实现了，因为客户端需要始终把请求发到同一台server才行，所谓“session迁移”等方案，也就是为了解决这个问题。


状态服务和无状态服务各有优劣，它们在一些情况下是可以转换的，或者有时候可以共用，并非一定要全部否定。

在一定需要处理请求上下文的情况下又想使用无状态服务,可以将相关的请求信息存储到共享内存中或者数据库中,参考分布式session的实现方式：

- 基于数据库的Session共享
- 基于NFS共享文件系统
- 基于memcached 的session
- 基于resin/tomcat web容器本身的session复制机制
- 基于TT/Redis 或 jbosscache 进行 session 共享

- 基于cookie 进行session共享



# 二、k8s部署中间件

在gulimall项目中

### 1、部署mysql

gulimall-mysql-master，16G内存不够用了，就不集群了。

1）配置中心里创建配置

![](https://gitee.com/HanFerm/image-bed/raw/master/img/20211011141450.png)

内存不够不用集群，所以不用填写主从复制配置内容了

2）创建mysql-pvc

3）创建有状态服务

创建有状态服务，选择mysql:3.7镜像，环境变量

![](https://gitee.com/HanFerm/image-bed/raw/master/img/20211011142159.png)

![](https://gitee.com/HanFerm/image-bed/raw/master/img/20211011143345.png)

挂载配置文件

![](https://gitee.com/HanFerm/image-bed/raw/master/img/20211011143257.png)

挂载存储卷

![](https://gitee.com/HanFerm/image-bed/raw/master/img/20211011143627.png)

![](https://gitee.com/HanFerm/image-bed/raw/master/img/20211011143648.png)

因为是有状态服务，所以默认会帮我们用无头service DNS创建域名

![](https://gitee.com/HanFerm/image-bed/raw/master/img/20211011143745.png)

继续创建从库的PVC、配置文件等内容，然后创建有状态服务。。。

然后进入到master容器组里，授权可以访问的登录用户名和IP、授权可以同步的IP和用户。

但是这样岂不是重启后master容器就变了？不会的，因为这些配置是持久化到了PVC中。

在从库容器中 change master to.....然后输入start slave

在主容器新增内容测试主从

#### 思路总结

- 每个mysql进程都被一个由状态服务包裹
- 每一个Mysql进程用配置文件CM和PVC存储内容，防止重新失效
- 以后的IP的都是使用的域名，而不是IP

### 2、部署Redis

再像前面一样创建CM(挂载/etc/redis  选定特定建和路径  redis-conf  redis.conf)和PVC（挂载/data）、

![](https://gitee.com/HanFerm/image-bed/raw/master/img/20211011150508.png)

有状态服务，在第二个配置里 容器镜像-启动命令：redis-server 参数：/etc/redis/redis.conf

### 3、部署ES、Kibana

添加配置 http.host值为0.0.0.0 ；discovery.type值为single-node； 还有 ES_JAVA_OPS "-Xms64m -Xmx512m"等等，然后创建服务时引入环境变量

挂载pvc：挂载到  /usr/share/elasticsearch/data



Kibana：无状态服务、环境变量主动ES主机（域名），选择外网访问Nodeport



#### 4、RabbitMQ

创建PVC，使用MQ-management（带管理界面），选择存储全卷挂载/var/lib/rabbitmq

### 5、Nacos

创建PVC、挂载/home/nacos

环境变量MODE standalone

服务状态切换：删除时不选择有状态副本集，容器没被删（相当于部署删了pod没删）。新建服务时 指定工作负载，选定原来的有状态副本集

### 6、Zipkin

无状态，它把数据交给了其他服务保存 

- STORAGE_TYPE   elasticsearch
- ES_HOST  elasticsearch.gulimall:9200

外网访问

### 7、Sentinel

无状态，官方没有镜像，有个人镜像，端口8333

# 三、部署微服务

### 1、流程

- 打包镜像上传仓库Dockerfile，Docker按该文件制作成镜像
- 编写Deploy文件部署到k8s集群  k8s的yaml  `kind:Deplyment`
- 编写service文件暴露到k8s集群 k8s的yaml  `kind:Service`
- 集群内访问测试，外部访问测试

Jenkins串联流程 Jenkinsfile

### 2、IP->域名

在微服务中创建application-prod.yaml，把原来application.yaml中的IP都改为k8s的service或ingress域名

因为像Nacos服务原来是暴露的外部端口，该端口会改变，所以我们不使用ingress条件下可以使用无头service。创建服务，关联工作负载，有状态副本集中选择nacos部署（不是pod），暴露到外部端口8848，端口名称如http-nacos-8848

### 3、Dockerfile

```dockerfile
FROM java:8
EXPOSE 8080
# 同时把微服务里的server.port改成8080
VOLUME /tmp
# 本地IDEA target/app.jar
ADD target/app.jar  /app.jar
# 主要是为了修改创建时间，可以去掉
RUN bash -c 'touch /app.jar'
# 容器启动默认运行命令
ENTRYPOINT ["java","-jar","/app.jar","--spring.profiles.active=prod"]

```

```properties
# 最后的.是jar包路径
docker build -f Dockerfile -t docker.io/hanferm/test:v1 .

docker login -u 账号 -p 密码
docker push 镜像
```

### 4、k8s.yaml

在每个微服务下创建deploy/devops.yaml 

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: 微服务名
  namespace: gulimall
  labels:
    app: 微服务名

spec:
  replicas: 1
  selector:
    matchLabels:
      app: 微服务名
  template:
    metadata:
      labels:
        app: 微服务名
    spec:
      containers:
        - name: ks
          image: $REGISTRY/$DOCKERHUB_NAMESPACE/$APP_NAME:$TAG_NAME
          ports:
            - containerPort: 8080
              protocol: TCP
          resources:
            limits:
              cpu: 200m
              memory: 500Mi
            requests:
              cpu: 100m
              memory: 100Mi
          imagePullPolicy: IfNotPresent
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
    strategy:
      type: RollingUpdate
      rollingUpdate:
        maxUnavailable: 25%
        maxSurge: 25%
---
kind: Service
apiVersion: v1
metadata:
  name: 微服务
  namespace: gulimall
  lables:
    app: 微服务
spec:
  ports:
    - name: http-8080
      protocol: TCP
      port: 8080service端口
      targetPort: 8080容器端口
      nodePort: 外部端口

  selector:
    app: 微服务
  clusterIP:
  type: NodePort
```

### 5、流水线

在ks中创建流水线gulimall-cicd

类型Node、label：maven

- 拉取代码
- 参数化构建parameters。运行流水线时就需要填入该参数，这个参数可以让我们指定具体哪个微服务参与此次部署，微服务选项制作成下拉栏的形式
- 环境变量environment，容器使用的环境变量
- Sonar代码质量检查
- 打包微服务、构建镜像、push到镜像镜像仓库
- k8s部署，此时k8s会拉镜像然后部署：`kubernetesDeploy(configs: '$微服务/deploy/**.yaml' , enableConfigSubstitution: true, kubeconfigId: "$KUBECONFIG_CREDENTIAL_ID")`

```properties
pipeline {
  agent {
    node {
      label 'maven'
    }
  }

    parameters {
        string(name:'TAG_NAME',defaultValue: '',description:'')
    }

    environment {
        DOCKER_CREDENTIAL_ID = 'dockerhub-id'
        GITHUB_CREDENTIAL_ID = 'gitee-id'
        KUBECONFIG_CREDENTIAL_ID = 'demo-kubeconfig'
        REGISTRY = 'docker.io'
        DOCKERHUB_NAMESPACE = 'hanferm'
        GITHUB_ACCOUNT = 'hanferm'
        APP_NAME = 'devops-java-sample'
        SONAR_CREDENTIAL_ID= 'sonar-qube'
    }

    stages {
        stage ('checkout scm') {
            steps {
                checkout(scm)
            }
        }

        stage ('unit test') {
            steps {
                container ('maven') {
                    sh 'mvn clean  -gs `pwd`/configuration/settings.xml test'
                }
            }
        }

        stage('sonarqube analysis') {
          steps {
            container ('maven') {
              withCredentials([string(credentialsId: "$SONAR_CREDENTIAL_ID", variable: 'SONAR_TOKEN')]) {
                withSonarQubeEnv('sonar') {
                 sh "mvn sonar:sonar  -gs `pwd`/configuration/settings.xml -Dsonar.branch=$BRANCH_NAME -Dsonar.login=$SONAR_TOKEN"
                }
              }
              timeout(time: 1, unit: 'HOURS') {
                waitForQualityGate abortPipeline: true
              }
            }
          }
        }

        stage ('构建build & push') {
            steps {
                container ('maven') {
                    sh 'mvn  -Dmaven.test.skip=true -gs `pwd`/configuration/settings.xml clean package'
                    sh 'docker build -f Dockerfile-online -t $REGISTRY/$DOCKERHUB_NAMESPACE/$APP_NAME:SNAPSHOT-$BRANCH_NAME-$BUILD_NUMBER .'
                    withCredentials([usernamePassword(passwordVariable : 'DOCKER_PASSWORD' ,usernameVariable : 'DOCKER_USERNAME' ,credentialsId : "$DOCKER_CREDENTIAL_ID" ,)]) {
                        sh 'echo "$DOCKER_PASSWORD" | docker login $REGISTRY -u "$DOCKER_USERNAME" --password-stdin'
                        sh 'docker push  $REGISTRY/$DOCKERHUB_NAMESPACE/$APP_NAME:SNAPSHOT-$BRANCH_NAME-$BUILD_NUMBER'
                    }
                }
            }
        }

        stage('push latest'){
           when{
             branch 'master'
           }
           steps{
                container ('maven') {
                  sh 'docker tag  $REGISTRY/$DOCKERHUB_NAMESPACE/$APP_NAME:SNAPSHOT-$BRANCH_NAME-$BUILD_NUMBER $REGISTRY/$DOCKERHUB_NAMESPACE/$APP_NAME:latest '
                  sh 'docker push  $REGISTRY/$DOCKERHUB_NAMESPACE/$APP_NAME:latest '
                }
           }
        }

        stage('deploy to dev') {
          when{
            branch 'master'
          }
          steps {
            input(id: 'deploy-to-dev', message: 'deploy to dev?')
            # k8s部署
            kubernetesDeploy(configs: 'deploy/dev-ol/**', enableConfigSubstitution: true, kubeconfigId: "$KUBECONFIG_CREDENTIAL_ID")
          }
        }
        stage('发布版push with tag'){
          when{
            expression{
              return params.TAG_NAME =~ /v.*/
            }
          }
          steps {
              container ('maven') {
                input(id: 'release-image-with-tag', message: 'release image with tag?')
                  withCredentials([usernamePassword(credentialsId: "$GITHUB_CREDENTIAL_ID", passwordVariable: 'GIT_PASSWORD', usernameVariable: 'GIT_USERNAME')]) {
                    sh 'git config --global user.email "kubesphere@yunify.com" '
                    sh 'git config --global user.name "kubesphere" '
                    sh 'git tag -a $TAG_NAME -m "$TAG_NAME" '
                    sh 'git push http://$GIT_USERNAME:$GIT_PASSWORD@gitee.com/$GITHUB_ACCOUNT/devops-java-sample.git --tags --ipv4'
                  }
                sh 'docker tag  $REGISTRY/$DOCKERHUB_NAMESPACE/$APP_NAME:SNAPSHOT-$BRANCH_NAME-$BUILD_NUMBER $REGISTRY/$DOCKERHUB_NAMESPACE/$APP_NAME:$TAG_NAME '
                sh 'docker push  $REGISTRY/$DOCKERHUB_NAMESPACE/$APP_NAME:$TAG_NAME '
          }
          }
        }
        stage('deploy to production') {
          when{
            expression{
              return params.TAG_NAME =~ /v.*/
            }
          }
          steps {
            input(id: 'deploy-to-production', message: 'deploy to production?')
            kubernetesDeploy(configs: 'deploy/prod-ol/**', enableConfigSubstitution: true, kubeconfigId: "$KUBECONFIG_CREDENTIAL_ID")
          }
        }
    }
}

```



## 笔记不易：

离线笔记均为markdown格式，图片也是云图，10多篇笔记20W字，压缩包仅500k，推荐使用typora阅读。也可以自己导入有道云笔记等软件中

阿里云图床现在**每周得几十元充值**，都要自己往里搭了，麻烦不要散播与转发

![](https://i0.hdslb.com/bfs/album/ff3fb7e24f05c6a850ede4b1f3acc54312c3b0c6.png)

打赏后请主动发支付信息到邮箱  553736044@qq.com  ，上班期间很容易忽略收账信息，邮箱回邮基本秒回（请备注付费的内容）

禁止转载发布，禁止散播，若发现大量散播，将对本系统文章图床进行重置处理。