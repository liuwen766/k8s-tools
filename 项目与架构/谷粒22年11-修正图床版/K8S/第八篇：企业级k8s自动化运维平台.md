# 前情回顾

- 使用ELK Stack收集、分析K8S内应用的日志
  - 使用高配版ELK模型：
    - 安装部署ES
    - 安装部署kafka
    - 制作filebeat镜像
    - 使用“SideCar”模式构建业务pod
    - 启动logstah
    - 部署Kibana
  - kibana的使用要点：
    - 时间选择器
    - 环境选择器
    - 项目选择器
    - 关键字选择器

# 第一章：容器云概述

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_41aa1ee2d75ea210d135df771b315287_r.png)

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_a2e78b84116cc817fa0eae2b71915988_r.png)

- IaaS是云服务的最底层，主要提供一些基础资源。
- PaaS提供软件部署平台(runtime),抽象掉硬件和操作系统细节，可以无缝地扩展(scaling)。开发者只需要关注自己的业务逻辑，不需要关注底层。
- SaaS 是软件的开发、管理、部署都是交给第三方，不需要关心技术问题，可以拿来即用。

------

**kubernetes不是什么？**

kubernetes并不是传统的PaaS（平台即服务）系统。

- Kubernetes不限制支持应用的类型，不限制应用框架。不限制受支持的语言runtimes（例如，java，python，ruby），满足12-factor applications。不区分“apps”或者“services”。
  kubernetes支持不同负载应用，包括有状态、无状态、数据处理类型的应用。只要这个应用可以在容器运行，那么就能很好的运行在kubernetes上。
- kubernetes不提供中间件(如message buses)、数据处理框架（如spark）、数据库（如Mysql）或者集群存储系统（如ceph）作为内置服务。这些应用都可以运行在Kubernetes上面。
- kubernetes不部署源码不编译应用。持续集成的（CI）工作流方面，不同的用户有不同的需求和偏好的区域，因此，我们提供分层的CI工作流，但并不定义它应该如何工作。
- kubernetes允许用户选择自己的日志、监控和报警系统。
- kubernetes不提供或授权一个全面的应用程序配置 语言/系统（例如，jsonnet）。
- kubernetes不提供任何机器配置、维护、管理或者自我修复系统。

------

- 越来越多的云计算厂商，正在基于K8S构建PaaS平台
- 获得PaaS能力的几个必要条件：
  - 统一应用的运行时环境（docker）
  - laaS能力（K8S）
  - 有可靠的中间件集群、数据库集群（DBA的主要工作）
  - 有分布式存储集群（存储工程师的主要工作）
  - 有适配的监控、日志系统（Prometheus、ELK）
  - 有完善的CI、CD系统（jenkins、？）
- 常见的基于K8S的CD系统
  - 自研
  - Argo CD
  - OpenShift
  - Spinnaker

# 第二章：Spinnaker概述

> [spinnaker](https://www.spinnaker.io/)是Netflix在2015年，开源的一款持续交付平台，它继承了Netflix上一代集群和部署管理工具 Asgard：web-based Cloud Management and Deployment的优点，同时根据公司业务以及技术的发展抛弃了一些过时的设计：提高了持续交付系统的可复用行，提供了稳定可靠的API，提供了对基础设施和程序全局性的视图，配置、管理、运维都更简单，而且还完全兼容Asgard，总之对于Netflix来说Spinnaker是更牛逼的持续交付平台。

[官方github](https://github.com/spinnaker)

## 1.主要功能

集群管理

> 集群管理主要用于管理云资源，Spinnaker所说的“云”可以理解成AWS，即主要是laas的资源，比如OpenStack，Google云，微软云等，后来还支持了容器和Kubernetes，但是管理方式还是按照管理基础设施的模式来设计的。

部署管理

> 管理部署流程是Spinaker的核心功能，他负责将jenkins流水线创建的镜像，部署到kubernetes集群中去，让服务真正运行起来。

# 第三章：自动化运维平台架构详解

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_d46fa9af85576cfdf4b95a0380f08253_r.png)

[来源及解释](https://www.spinnaker.io/reference/architecture/)

# 第四章：安装部署Spinnaker微服务集群

## 1.部署对象式存储组件—minio

运维主机shkf6-245.host.com上：

### 1.下载镜像

[镜像下载地址](https://hub.docker.com/r/minio/minio)

```shell
[root@shkf6-245 ~]# docker pull minio/minio:latest

[root@shkf6-245 ~]# docker images|grep minio
minio/minio                                       latest                     902f6a03bf69        5 days ago          53.6MB

[root@shkf6-245 ~]# docker tag 902f6a03bf69 harbor.od.com/armory/minio:latest
[root@shkf6-245 ~]# docker push harbor.od.com/armory/minio:latest
[root@shkf6-245 ~]# mkdir /data/nfs-volume/minio
```

### 2.准备资源配置清单

- Deployment

```shell
[root@shkf6-245 ~]# mkdir /data/k8s-yaml/armory/minio

[root@shkf6-245 ~]# vi /data/k8s-yaml/armory/minio/dp.yaml
[root@shkf6-245 ~]# cat /data/k8s-yaml/armory/minio/dp.yaml
kind: Deployment
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  labels:
    name: minio
  name: minio
  namespace: armory
spec:
  progressDeadlineSeconds: 600
  replicas: 1
  revisionHistoryLimit: 7
  selector:
    matchLabels:
      name: minio
  template:
    metadata:
      labels:
        app: minio
        name: minio
    spec:
      containers:
      - name: minio
        image: harbor.od.com/armory/minio:latest
        imagePullPolicy: IfNotPresent
        ports:
        - containerPort: 9000
          protocol: TCP
        args:
        - server
        - /data
        env:
        - name: MINIO_ACCESS_KEY
          value: admin
        - name: MINIO_SECRET_KEY
          value: admin123
        readinessProbe:
          failureThreshold: 3
          httpGet:
            path: /minio/health/ready
            port: 9000
            scheme: HTTP
          initialDelaySeconds: 10
          periodSeconds: 10
          successThreshold: 1
          timeoutSeconds: 5
        volumeMounts:
        - mountPath: /data
          name: data
      imagePullSecrets:
      - name: harbor
      volumes:
      - nfs:
          server: shkf6-245
          path: /data/nfs-volume/minio
        name: data
```

- Service

```shell
[root@shkf6-245 ~]# vi /data/k8s-yaml/armory/minio/svc.yaml
[root@shkf6-245 ~]# cat /data/k8s-yaml/armory/minio/svc.yaml
apiVersion: v1
kind: Service
metadata:
  name: minio
  namespace: armory
spec:
  ports:
  - port: 80
    protocol: TCP
    targetPort: 9000
  selector:
    app: minio
```

- Ingress

```shell
[root@shkf6-245 ~]# vi /data/k8s-yaml/armory/minio/ingress.yaml

[root@shkf6-245 ~]# cat /data/k8s-yaml/armory/minio/ingress.yaml
kind: Ingress
apiVersion: extensions/v1beta1
metadata: 
  name: minio
  namespace: armory
spec:
  rules:
  - host: minio.od.com
    http:
      paths:
      - path: /
        backend: 
          serviceName: minio
          servicePort: 80
```

### 3.解析域名

```shell
[root@shkf6-241 ~]# tail -1 /var/named/od.com.zone
minio               A    192.168.6.66
```

### 4.应用资源配置清单

在任意一台运算节点

```shell
[root@shkf6-243 ~]#  kubectl create secret docker-registry harbor --docker-server=harbor.od.com --docker-username=admin --docker-password=Harbor12345 -n armory

[root@shkf6-243 ~]# kubectl create ns armory
namespace/armory created
[root@shkf6-243 ~]# kubectl apply -f http://k8s-yaml.od.com/armory/minio/dp.yaml 
deployment.extensions/minio created
[root@shkf6-243 ~]# kubectl apply -f http://k8s-yaml.od.com/armory/minio/svc.yaml 
service/minio created
[root@shkf6-243 ~]# kubectl apply -f http://k8s-yaml.od.com/armory/minio/ingress.yaml 
ingress.extensions/minio created
```

### 5.浏览器访问

```shell
http://minio.od.com
```

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_3cfa0d7379eba292be188a4c0deaec21_r.png)

## 2.部署缓存组件—Redis

### 1.准备docker镜像

运维主机shkf6-245.host.com上：

[镜像下载地址](https://hub.docker.com/_/redis)

```shell
[root@shkf6-245 ~]# docker pull redis:4.0.14

[root@shkf6-245 ~]# docker images|grep redis
redis                                             4.0.14                     b93767ee535f        6 days ago          89.2MB

[root@shkf6-245 ~]# docker tag b93767ee535f harbor.od.com/armory/redis:v4.0.14
[root@shkf6-245 ~]# docker push harbor.od.com/armory/redis:v4.0.14
```

### 2.准备资源配置清单

- Deployment

```shell
[root@shkf6-245 ~]# mkdir /data/k8s-yaml/armory/redis
[root@shkf6-245 ~]# vi /data/k8s-yaml/armory/redis/dp.yaml
[root@shkf6-245 ~]# cat /data/k8s-yaml/armory/redis/dp.yaml
kind: Deployment
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  labels:
    name: redis
  name: redis
  namespace: armory
spec:
  replicas: 1
  revisionHistoryLimit: 7
  selector:
    matchLabels:
      name: redis
  template:
    metadata:
      labels:
        app: redis
        name: redis
    spec:
      containers:
      - name: redis
        image: harbor.od.com/armory/redis:v4.0.14
        imagePullPolicy: IfNotPresent
        ports:
        - containerPort: 6379
          protocol: TCP
      imagePullSecrets:
      - name: harbor
```

- Service

```shell
[root@shkf6-245 ~]# vi /data/k8s-yaml/armory/redis/svc.yaml
[root@shkf6-245 ~]# cat /data/k8s-yaml/armory/redis/svc.yaml
apiVersion: v1
kind: Service
metadata:
  name: redis
  namespace: armory
spec:
  ports:
  - port: 6379
    protocol: TCP
    targetPort: 6379
  selector:
    app: redis
```

### 3.应用资源配置清单

任意运算节点上：

```shell
[root@shkf6-243 ~]# kubectl apply -f http://k8s-yaml.od.com/redis/dp.yaml
deployment.extensions/redis created
[root@shkf6-243 ~]# kubectl apply -f http://k8s-yaml.od.com/redis/svc.yaml
service/redis created
```

## 3.部署k8s云驱动组件—CloudDriver

运维主机shkf6-245.host.com上：

### 1.准备docker镜像

[镜像下载地址](https://quay.io/repository/container-image/spinnaker-clouddriver)

```shell
[root@shkf6-245 ~]# docker pull docker.io/armory/spinnaker-clouddriver-slim:release-1.8.x-14c9664

[root@shkf6-245 ~]# docker images|grep spinnaker-clouddriver-slim
armory/spinnaker-clouddriver-slim                 release-1.8.x-14c9664      edb2507fdb62        18 months ago       662MB

[root@shkf6-245 ~]# docker tag edb2507fdb62 harbor.od.com/armory/clouddriver:v1.8.x
[root@shkf6-245 ~]# docker push harbor.od.com/armory/clouddriver:v1.8.x
```

### 2.准备minio的secret

- 准备配置文件

运维主机shkf6-245.host.com上：

```shell
[root@shkf6-245 ~]# mkdir -p /data/k8s-yaml/armory/clouddriver
[root@shkf6-245 ~]# vi /data/k8s-yaml/armory/clouddriver/credentials
[root@shkf6-245 ~]# cat /data/k8s-yaml/armory/clouddriver/credentials
[default]
aws_access_key_id=admin
aws_secret_access_key=admin123
```

- 创建secret

任意运算节点上：

```shell
[root@shkf6-243 ~]# wget http://k8s-yaml.od.com/armory/clouddriver/credentials
[root@shkf6-243 ~]# kubectl create secret generic credentials --from-file=./credentials -n armory
```

### 3.准备k8s的用户配置

#### 1.签发证书和私钥

在运维主机shkf6-245.host.com上：

- 准备证书签发文件admin-csr.json

```shell
[root@shkf6-245 ~]# cp /opt/certs/client-csr.json /opt/certs/admin-csr.json
[root@shkf6-245 ~]# cd /opt/certs/
[root@shkf6-245 certs]#cat admin-csr.json
{
    "CN": "cluster-admin",
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

**注：CN要设置为：cluster-admin**

- 签发生成admin.pem、admin-key.pem

```shell
[root@shkf6-245 certs]# cfssl gencert -ca=ca.pem -ca-key=ca-key.pem -config=ca-config.json -profile=client admin-csr.json | cfssl-json -bare admin
2020/01/09 17:36:20 [INFO] generate received request
2020/01/09 17:36:20 [INFO] received CSR
2020/01/09 17:36:20 [INFO] generating key: rsa-2048
2020/01/09 17:36:20 [INFO] encoded CSR
2020/01/09 17:36:20 [INFO] signed certificate with serial number 13522671111456276473606163567487943761642592599
2020/01/09 17:36:20 [WARNING] This certificate lacks a "hosts" field. This makes it unsuitable for
websites. For more information see the Baseline Requirements for the Issuance and Management
of Publicly-Trusted Certificates, v.1.1.6, from the CA/Browser Forum (https://cabforum.org);
specifically, section 10.2.3 ("Information Requirements").
[root@shkf6-245 certs]# ll admin*
-rw-r--r-- 1 root root 1001 Jan  9 17:36 admin.csr
-rw-r--r-- 1 root root  285 Jan  9 17:32 admin-csr.json
-rw------- 1 root root 1675 Jan  9 17:36 admin-key.pem
-rw-r--r-- 1 root root 1371 Jan  9 17:36 admin.pem
```

#### 2.做kubeconfig配置

任意运算节点上：

```shell
[root@shkf6-243 ~]# scp shkf6-245:/opt/certs/ca.pem .
ca.pem                                                                                             100% 1346    44.4KB/s   00:00    
[root@shkf6-243 ~]# scp shkf6-245:/opt/certs/admin.pem .
admin.pem                                                                                          100% 1371    86.5KB/s   00:00    
[root@shkf6-243 ~]# scp shkf6-245:/opt/certs/admin-key.pem .
admin-key.pem                                                                                      100% 1675   115.6KB/s   00:00    
[root@shkf6-243 ~]# kubectl config set-cluster myk8s --certificate-authority=./ca.pem --embed-certs=true --server=https://192.168.6.66:7443 --kubeconfig=config
Cluster "myk8s" set.
[root@shkf6-243 ~]# kubectl config set-credentials cluster-admin --client-certificate=./admin.pem --client-key=./admin-key.pem --embed-certs=true --kubeconfig=config
User "cluster-admin" set.
[root@shkf6-243 ~]# kubectl config set-context myk8s-context --cluster=myk8s --user=cluster-admin --kubeconfig=config
Context "myk8s-context" created.
[root@shkf6-243 ~]# kubectl config use-context myk8s-context --kubeconfig=config
Switched to context "myk8s-context".
[root@shkf6-243 ~]# kubectl create clusterrolebinding myk8s-admin --clusterrole=cluster-admin --user=cluster-admin
clusterrolebinding.rbac.authorization.k8s.io/myk8s-admin created
```

#### 3.验证cluster-admin用户

> 将config文件拷贝至任意运算节点/root/.kube下，使用kubectl验证

#### 4.创建configmap配置

```shell
[root@shkf6-243 ~]# cp config default-kubeconfig
[root@shkf6-243 ~]# kubectl create cm default-kubeconfig --from-file=default-kubeconfig -n armory
configmap/default-kubeconfig created
```

#### 5.配置运维主机管理k8s集群

```shell
[root@shkf6-245 ~]# mkdir /root/.kube
[root@shkf6-245 ~]# cd /root/.kube
[root@shkf6-245 .kube]# scp shkf6-243:/root/config .

[root@shkf6-245 .kube]# scp shkf6-243:/opt/kubernetes/server/bin/kubectl /sbin/kubectl

[root@shkf6-245 .kube]# kubectl top node
NAME                 CPU(cores)   CPU%   MEMORY(bytes)   MEMORY%   
shkf6-243.host.com   2062m        34%    7511Mi          63%       
shkf6-244.host.com   1571m        26%    6525Mi          55% 

[root@shkf6-245 ~]# yum install bash-completion -y
[root@shkf6-245 ~]# kubectl completion bash > /etc/bash_completion.d/kubectl
```

#### 6.控制节点config和远程config 对比：（分析）

```shell
[root@shkf6-243 ~]# kubectl config view
apiVersion: v1
clusters: []
contexts: []
current-context: ""
kind: Config
preferences: {}
users: []

[root@shkf6-245 ~]# kubectl config view
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: DATA+OMITTED
    server: https://192.168.6.66:7443
  name: myk8s
contexts:
- context:
    cluster: myk8s
    user: cluster-admin
  name: myk8s-context
current-context: myk8s-context
kind: Config
preferences: {}
users:
- name: cluster-admin
  user:
    client-certificate-data: REDACTED
    client-key-data: REDACTED
```

### 4.准备资源配置清单

- ConfigMap1

```shell
[root@shkf6-245 k8s-yaml]# vi /data/k8s-yaml/armory/clouddriver/init-env.yaml
[root@shkf6-245 k8s-yaml]# cat /data/k8s-yaml/armory/clouddriver/init-env.yaml 
kind: ConfigMap
apiVersion: v1
metadata:
  name: init-env
  namespace: armory
data:
  API_HOST: http://spinnaker.od.com/api
  ARMORY_ID: c02f0781-92f5-4e80-86db-0ba8fe7b8544
  ARMORYSPINNAKER_CONF_STORE_BUCKET: armory-platform
  ARMORYSPINNAKER_CONF_STORE_PREFIX: front50
  ARMORYSPINNAKER_GCS_ENABLED: "false"
  ARMORYSPINNAKER_S3_ENABLED: "true"
  AUTH_ENABLED: "false"
  AWS_REGION: us-east-1
  BASE_IP: 127.0.0.1
  CLOUDDRIVER_OPTS: -Dspring.profiles.active=armory,configurator,local
  CONFIGURATOR_ENABLED: "false"
  DECK_HOST: http://spinnaker.od.com
  ECHO_OPTS: -Dspring.profiles.active=armory,configurator,local
  GATE_OPTS: -Dspring.profiles.active=armory,configurator,local
  IGOR_OPTS: -Dspring.profiles.active=armory,configurator,local
  PLATFORM_ARCHITECTURE: k8s
  REDIS_HOST: redis://redis:6379
  SERVER_ADDRESS: 0.0.0.0
  SPINNAKER_AWS_DEFAULT_REGION: us-east-1
  SPINNAKER_AWS_ENABLED: "false"
  SPINNAKER_CONFIG_DIR: /home/spinnaker/config
  SPINNAKER_GOOGLE_PROJECT_CREDENTIALS_PATH: ""
  SPINNAKER_HOME: /home/spinnaker
  SPRING_PROFILES_ACTIVE: armory,configurator,local
```

- ConfigMap2

```shell
wget -O /data/k8s-yaml/armory/clouddriver/default-config.yaml http://down.sunrisenan.com/k8s/default-config.yaml
```

- ConfigMap3

```shell
[root@shkf6-245 k8s-yaml]# vi /data/k8s-yaml/armory/clouddriver/custom-config.yaml 
[root@shkf6-245 k8s-yaml]# cat /data/k8s-yaml/armory/clouddriver/custom-config.yaml 
kind: ConfigMap
apiVersion: v1
metadata:
  name: custom-config
  namespace: armory
data:
  clouddriver-local.yml: |
    kubernetes:
      enabled: true
      accounts:
        - name: cluster-admin
          serviceAccount: false
          dockerRegistries:
            - accountName: harbor
              namespace: []
          namespaces:
            - test
            - prod
          kubeconfigFile: /opt/spinnaker/credentials/custom/default-kubeconfig
      primaryAccount: cluster-admin
    dockerRegistry:
      enabled: true
      accounts:
        - name: harbor
          requiredGroupMembership: []
          providerVersion: V1
          insecureRegistry: true
          address: http://harbor.od.com
          username: admin
          password: Harbor12345
      primaryAccount: harbor
    artifacts:
      s3:
        enabled: true
        accounts:
        - name: armory-config-s3-account
          apiEndpoint: http://minio
          apiRegion: us-east-1
      gcs:
        enabled: false
        accounts:
        - name: armory-config-gcs-account
  custom-config.json: ""
  echo-configurator.yml: |
    diagnostics:
      enabled: true
  front50-local.yml: |
    spinnaker:
      s3:
        endpoint: http://minio
  igor-local.yml: |
    jenkins:
      enabled: true
      masters:
        - name: jenkins-admin
          address: http://jenkins.od.com
          username: admin
          password: admin123
      primaryAccount: jenkins-admin
  nginx.conf: |
    gzip on;
    gzip_types text/plain text/css application/json application/x-javascript text/xml application/xml application/xml+rss text/javascript application/vnd.ms-fontobject application/x-font-ttf font/opentype image/svg+xml image/x-icon;

    server {
           listen 80;

           location / {
                proxy_pass http://armory-deck/;
           }

           location /api/ {
                proxy_pass http://armory-gate:8084/;
           }

           rewrite ^/login(.*)$ /api/login$1 last;
           rewrite ^/auth(.*)$ /api/auth$1 last;
    }
  spinnaker-local.yml: |
    services:
      igor:
        enabled: true
```

- Deployment

```shell
[root@shkf6-245 k8s-yaml]# cat /data/k8s-yaml/armory/clouddriver/dp.yaml 
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  labels:
    app: armory-clouddriver
  name: armory-clouddriver
  namespace: armory
spec:
  replicas: 1
  revisionHistoryLimit: 7
  selector:
    matchLabels:
      app: armory-clouddriver
  template:
    metadata:
      annotations:
        artifact.spinnaker.io/location: '"armory"'
        artifact.spinnaker.io/name: '"armory-clouddriver"'
        artifact.spinnaker.io/type: '"kubernetes/deployment"'
        moniker.spinnaker.io/application: '"armory"'
        moniker.spinnaker.io/cluster: '"clouddriver"'
      labels:
        app: armory-clouddriver
    spec:
      containers:
      - name: armory-clouddriver
        image: harbor.od.com/armory/clouddriver:v1.8.x
        imagePullPolicy: IfNotPresent
        command:
        - bash
        - -c
        args:
        - bash /opt/spinnaker/config/default/fetch.sh && cd /home/spinnaker/config
          && /opt/clouddriver/bin/clouddriver
        ports:
        - containerPort: 7002
          protocol: TCP
        env:
        - name: JAVA_OPTS
          value: -Xmx2000M
        envFrom:
        - configMapRef:
            name: init-env
        livenessProbe:
          failureThreshold: 5
          httpGet:
            path: /health
            port: 7002
            scheme: HTTP
          initialDelaySeconds: 600
          periodSeconds: 3
          successThreshold: 1
          timeoutSeconds: 1
        readinessProbe:
          failureThreshold: 5
          httpGet:
            path: /health
            port: 7002
            scheme: HTTP
          initialDelaySeconds: 180
          periodSeconds: 3
          successThreshold: 5
          timeoutSeconds: 1
        securityContext: 
          runAsUser: 0
        volumeMounts:
        - mountPath: /etc/podinfo
          name: podinfo
        - mountPath: /home/spinnaker/.aws
          name: credentials
        - mountPath: /opt/spinnaker/credentials/custom
          name: default-kubeconfig
        - mountPath: /opt/spinnaker/config/default
          name: default-config
        - mountPath: /opt/spinnaker/config/custom
          name: custom-config
      imagePullSecrets:
      - name: harbor
      volumes:
      - configMap:
          defaultMode: 420
          name: default-kubeconfig
        name: default-kubeconfig
      - configMap:
          defaultMode: 420
          name: custom-config
        name: custom-config
      - configMap:
          defaultMode: 420
          name: default-config
        name: default-config
      - name: credentials
        secret:
          defaultMode: 420
          secretName: credentials
      - downwardAPI:
          defaultMode: 420
          items:
          - fieldRef:
              apiVersion: v1
              fieldPath: metadata.labels
            path: labels
          - fieldRef:
              apiVersion: v1
              fieldPath: metadata.annotations
            path: annotations
        name: podinfo
```

- Service

```shell
[root@shkf6-245 k8s-yaml]# vi /data/k8s-yaml/armory/clouddriver/svc.yaml
[root@shkf6-245 k8s-yaml]# cat /data/k8s-yaml/armory/clouddriver/svc.yaml 
apiVersion: v1
kind: Service
metadata:
  name: armory-clouddriver
  namespace: armory
spec:
  ports:
  - port: 7002
    protocol: TCP
    targetPort: 7002
  selector:
    app: armory-clouddriver
```

### 5.应用资源配置清单

在任意一台运算节点

```shell
[root@shkf6-245 ~]# cd /data/k8s-yaml/armory/clouddriver/
[root@shkf6-245 clouddriver]# kubectl apply -f ./
configmap/custom-config created
configmap/default-config created
deployment.extensions/armory-clouddriver created
configmap/init-env created
service/armory-clouddriver created
```

### 6.检查

准备配置文件

```shell
[root@shkf6-243 ~]# cat nginx-armory.yaml 
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: nginx-armory
  namespace: armory
spec:
  template:
    metadata:
      labels:
        app: nginx-armory
    spec:
      containers:
      - name: my-nginx
        image: harbor.od.com/public/nginx:curl
        command: ["nginx","-g","daemon off;"]
        ports:
        - containerPort: 80
```

应用配置文件

```shell
[root@shkf6-243 ~]# kubectl apply -f nginx-armory.yaml 
deployment.extensions/nginx-armory created
```

验证clouddriver

```shell
[root@shkf6-243 ~]# kubectl exec -it armory-clouddriver-5f6cff8bb8-hmhrk bash -n armory
bash-4.4# exit
[root@shkf6-243 ~]# kubectl exec -it nginx-armory-cb6cfdfd-5lcsr bash -n armory
root@nginx-armory-cb6cfdfd-5lcsr:/# curl armory-clouddriver:7002/health
{"status":"UP","kubernetes":{"status":"UP"},"redisHealth":{"status":"UP","maxIdle":100,"minIdle":25,"numActive":0,"numIdle":4,"numWaiters":0},"dockerRegistry":{"status":"UP"},"diskSpace":{"status":"UP","total":44550057984,"free":35043196928,"threshold":10485760}}
```

## 4.部署数据持久化组件—Front50

在运维主机shkf6-245.host.com上

### 1.准备镜像

[镜像下载地址](https://quay.io/repository/container-image/spinnaker-front50)

```shell
[root@shkf6-245 clouddriver]# docker pull docker.io/armory/spinnaker-front50-slim:release-1.8.x-93febf2

[root@shkf6-245 clouddriver]# docker images | grep front50
armory/spinnaker-front50-slim                     release-1.8.x-93febf2      0d353788f4f2        15 months ago       273MB

[root@shkf6-245 clouddriver]# docker tag 0d353788f4f2 harbor.od.com/armory/front50:v1.8.x
[root@shkf6-245 clouddriver]# docker push harbor.od.com/armory/front50:v1.8.x
```

### 2.准备资源配置清单

- Deployment

```shell
[root@shkf6-245 armory]# mkdir /data/k8s-yaml/armory/front50
[root@shkf6-245 armory]# vi /data/k8s-yaml/armory/front50/dp.yaml
[root@shkf6-245 armory]# cat /data/k8s-yaml/armory/front50/dp.yaml
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  labels:
    app: armory-front50
  name: armory-front50
  namespace: armory
spec:
  replicas: 1
  revisionHistoryLimit: 7
  selector:
    matchLabels:
      app: armory-front50
  template:
    metadata:
      annotations:
        artifact.spinnaker.io/location: '"armory"'
        artifact.spinnaker.io/name: '"armory-front50"'
        artifact.spinnaker.io/type: '"kubernetes/deployment"'
        moniker.spinnaker.io/application: '"armory"'
        moniker.spinnaker.io/cluster: '"front50"'
      labels:
        app: armory-front50
    spec:
      containers:
      - name: armory-front50
        image: harbor.od.com/armory/front50:v1.8.x
        imagePullPolicy: IfNotPresent
        command:
        - bash
        - -c
        args:
        - bash /opt/spinnaker/config/default/fetch.sh && cd /home/spinnaker/config
          && /opt/front50/bin/front50
        ports:
        - containerPort: 8080
          protocol: TCP
        env:
        - name: JAVA_OPTS
          value: -javaagent:/opt/front50/lib/jamm-0.2.5.jar -Xmx1000M
        envFrom:
        - configMapRef:
            name: init-env
        livenessProbe:
          failureThreshold: 3
          httpGet:
            path: /health
            port: 8080
            scheme: HTTP
          initialDelaySeconds: 600
          periodSeconds: 3
          successThreshold: 1
          timeoutSeconds: 1
        readinessProbe:
          failureThreshold: 3
          httpGet:
            path: /health
            port: 8080
            scheme: HTTP
          initialDelaySeconds: 180
          periodSeconds: 5
          successThreshold: 8
          timeoutSeconds: 1
        volumeMounts:
        - mountPath: /etc/podinfo
          name: podinfo
        - mountPath: /home/spinnaker/.aws
          name: credentials
        - mountPath: /opt/spinnaker/config/default
          name: default-config
        - mountPath: /opt/spinnaker/config/custom
          name: custom-config
      imagePullSecrets:
      - name: harbor
      volumes:
      - configMap:
          defaultMode: 420
          name: custom-config
        name: custom-config
      - configMap:
          defaultMode: 420
          name: default-config
        name: default-config
      - name: credentials
        secret:
          defaultMode: 420
          secretName: credentials
      - downwardAPI:
          defaultMode: 420
          items:
          - fieldRef:
              apiVersion: v1
              fieldPath: metadata.labels
            path: labels
          - fieldRef:
              apiVersion: v1
              fieldPath: metadata.annotations
            path: annotations
        name: podinfo
```

- Service

```shell
[root@shkf6-245 armory]# vi /data/k8s-yaml/armory/front50/svc.yaml
[root@shkf6-245 armory]# cat /data/k8s-yaml/armory/front50/svc.yaml
apiVersion: v1
kind: Service
metadata:
  name: armory-front50
  namespace: armory
spec:
  ports:
  - port: 8080
    protocol: TCP
    targetPort: 8080
  selector:
    app: armory-front50
```

### 3.应用资源清单

```shell
[root@shkf6-245 armory]# kubectl apply -f ./front50/
deployment.extensions/armory-front50 created
service/armory-front50 created
```

### 4.檢查：

```shell
[root@shkf6-243 ~]# kubectl exec -it nginx-armory-cb6cfdfd-5lcsr bash -n armory 
root@nginx-armory-cb6cfdfd-5lcsr:/# curl armory-front50:8080/health
{"status":"UP"}
```

### 5.浏览器访问

[http://minio.od.com](http://minio.od.com/) 登录并观察存储是否创建（已創建）
![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_fc6319e4150831ecf32ea9770e6c2dfe_r.png)

```shell
[root@shkf6-245 armory]# ll /data/nfs-volume/minio/
total 0
drwxr-xr-x 2 root root 6 Jan 10 15:26 armory-platform
```

## 5.部署任务编排组件–Orca

运维主机shkf6-245.host.com

### 1.准备docker镜像

[镜像下载地址](https://quay.io/repository/container-image/spinnaker-orca)

```shell
[root@shkf6-245 armory]# docker pull docker.io/armory/spinnaker-orca-slim:release-1.8.x-de4ab55

[root@shkf6-245 armory]# docker images | grep orca
armory/spinnaker-orca-slim                        release-1.8.x-de4ab55      5103b1f73e04        15 months ago       141MB

[root@shkf6-245 armory]# docker tag 5103b1f73e04 harbor.od.com/armory/orca:v1.8.x
[root@shkf6-245 armory]# docker push harbor.od.com/armory/orca:v1.8.x
```

### 2.准备资源配置清单

- Deployment

```shell
[root@shkf6-245 armory]# mkdir /data/k8s-yaml/armory/orca
[root@shkf6-245 armory]# vi /data/k8s-yaml/armory/orca/dp.yaml
[root@shkf6-245 armory]# cat /data/k8s-yaml/armory/orca/dp.yaml
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  labels:
    app: armory-orca
  name: armory-orca
  namespace: armory
spec:
  replicas: 1
  revisionHistoryLimit: 7
  selector:
    matchLabels:
      app: armory-orca
  template:
    metadata:
      annotations:
        artifact.spinnaker.io/location: '"armory"'
        artifact.spinnaker.io/name: '"armory-orca"'
        artifact.spinnaker.io/type: '"kubernetes/deployment"'
        moniker.spinnaker.io/application: '"armory"'
        moniker.spinnaker.io/cluster: '"orca"'
      labels:
        app: armory-orca
    spec:
      containers:
      - name: armory-orca
        image: harbor.od.com/armory/orca:v1.8.x
        imagePullPolicy: IfNotPresent
        command:
        - bash
        - -c
        args:
        - bash /opt/spinnaker/config/default/fetch.sh && cd /home/spinnaker/config
          && /opt/orca/bin/orca
        ports:
        - containerPort: 8083
          protocol: TCP
        env:
        - name: JAVA_OPTS
          value: -Xmx1000M
        envFrom:
        - configMapRef:
            name: init-env
        livenessProbe:
          failureThreshold: 5
          httpGet:
            path: /health
            port: 8083
            scheme: HTTP
          initialDelaySeconds: 600
          periodSeconds: 5
          successThreshold: 1
          timeoutSeconds: 1
        readinessProbe:
          failureThreshold: 3
          httpGet:
            path: /health
            port: 8083
            scheme: HTTP
          initialDelaySeconds: 180
          periodSeconds: 3
          successThreshold: 5
          timeoutSeconds: 1
        volumeMounts:
        - mountPath: /etc/podinfo
          name: podinfo
        - mountPath: /opt/spinnaker/config/default
          name: default-config
        - mountPath: /opt/spinnaker/config/custom
          name: custom-config
      imagePullSecrets:
      - name: harbor
      volumes:
      - configMap:
          defaultMode: 420
          name: custom-config
        name: custom-config
      - configMap:
          defaultMode: 420
          name: default-config
        name: default-config
      - downwardAPI:
          defaultMode: 420
          items:
          - fieldRef:
              apiVersion: v1
              fieldPath: metadata.labels
            path: labels
          - fieldRef:
              apiVersion: v1
              fieldPath: metadata.annotations
            path: annotations
        name: podinfo
```

- Service

```shell
[root@shkf6-245 armory]# vi /data/k8s-yaml/armory/orca/svc.yaml
[root@shkf6-245 armory]# cat /data/k8s-yaml/armory/orca/svc.yaml
apiVersion: v1
kind: Service
metadata:
  name: armory-orca
  namespace: armory
spec:
  ports:
  - port: 8083
    protocol: TCP
    targetPort: 8083
  selector:
    app: armory-orca
```

### 3.应用资源配置清单

```shell
[root@shkf6-245 armory]# kubectl apply -f ./orca/
deployment.extensions/armory-orca created
service/armory-orca created
```

### 4.檢查：

```shell
[root@shkf6-243 ~]# kubectl exec -it nginx-armory-cb6cfdfd-5lcsr bash -n armory 
root@nginx-armory-cb6cfdfd-5lcsr:/# curl armory-orca:8083/health
{"status":"UP"}
```

## 6.部署消息总线组件–Echo

### 1.准备镜像

```shell
[root@shkf6-245 armory]# docker pull docker.io/armory/echo-armory:c36d576-release-1.8.x-617c567

[root@shkf6-245 armory]# docker images|grep echo
armory/echo-armory                                c36d576-release-1.8.x-617c567   415efd46f474        18 months ago       287MB

[root@shkf6-245 armory]# docker tag 415efd46f474 harbor.od.com/armory/echo:v1.8.x
[root@shkf6-245 armory]# docker push harbor.od.com/armory/echo:v1.8.x
```

### 2.准备资源配置清单

- Deployment

```shell
[root@shkf6-245 armory]# mkdir /data/k8s-yaml/armory/echo
mkdir: cannot create directory ‘/data/k8s-yaml/armory/echo’: File exists
[root@shkf6-245 armory]# vi /data/k8s-yaml/armory/echo/dp.yaml
[root@shkf6-245 armory]# cat /data/k8s-yaml/armory/echo/dp.yaml
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  labels:
    app: armory-echo
  name: armory-echo
  namespace: armory
spec:
  replicas: 1
  revisionHistoryLimit: 7
  selector:
    matchLabels:
      app: armory-echo
  template:
    metadata:
      annotations:
        artifact.spinnaker.io/location: '"armory"'
        artifact.spinnaker.io/name: '"armory-echo"'
        artifact.spinnaker.io/type: '"kubernetes/deployment"'
        moniker.spinnaker.io/application: '"armory"'
        moniker.spinnaker.io/cluster: '"echo"'
      labels:
        app: armory-echo
    spec:
      containers:
      - name: armory-echo
        image: harbor.od.com/armory/echo:v1.8.x
        imagePullPolicy: IfNotPresent
        command:
        - bash
        - -c
        args:
        - bash /opt/spinnaker/config/default/fetch.sh && cd /home/spinnaker/config
          && /opt/echo/bin/echo
        ports:
        - containerPort: 8089
          protocol: TCP
        env:
        - name: JAVA_OPTS
          value: -javaagent:/opt/echo/lib/jamm-0.2.5.jar -Xmx1000M
        envFrom:
        - configMapRef:
            name: init-env
        livenessProbe:
          failureThreshold: 3
          httpGet:
            path: /health
            port: 8089
            scheme: HTTP
          initialDelaySeconds: 600
          periodSeconds: 3
          successThreshold: 1
          timeoutSeconds: 1
        readinessProbe:
          failureThreshold: 3
          httpGet:
            path: /health
            port: 8089
            scheme: HTTP
          initialDelaySeconds: 180
          periodSeconds: 3
          successThreshold: 5
          timeoutSeconds: 1
        volumeMounts:
        - mountPath: /etc/podinfo
          name: podinfo
        - mountPath: /opt/spinnaker/config/default
          name: default-config
        - mountPath: /opt/spinnaker/config/custom
          name: custom-config
      imagePullSecrets:
      - name: harbor
      volumes:
      - configMap:
          defaultMode: 420
          name: custom-config
        name: custom-config
      - configMap:
          defaultMode: 420
          name: default-config
        name: default-config
      - downwardAPI:
          defaultMode: 420
          items:
          - fieldRef:
              apiVersion: v1
              fieldPath: metadata.labels
            path: labels
          - fieldRef:
              apiVersion: v1
              fieldPath: metadata.annotations
            path: annotations
        name: podinfo
```

- Service

```shell
[root@shkf6-245 armory]# vi /data/k8s-yaml/armory/echo/svc.yaml
[root@shkf6-245 armory]# cat /data/k8s-yaml/armory/echo/svc.yaml
apiVersion: v1
kind: Service
metadata:
  name: armory-echo
  namespace: armory
spec:
  ports:
  - port: 8089
    protocol: TCP
    targetPort: 8089
  selector:
    app: armory-echo
```

### 3.应用资源配置清单

```shell
[root@shkf6-245 armory]# kubectl apply -f ./echo/
deployment.extensions/armory-echo created
service/armory-echo created
```

## 7.部署流水线交互组件–lgor

### 1.准备镜像

```shell
[root@shkf6-245 armory]# docker pull docker.io/armory/spinnaker-igor-slim:release-1.8-x-new-install-healthy-ae2b329

[root@shkf6-245 armory]# docker images |grep igor
armory/spinnaker-igor-slim                        release-1.8-x-new-install-healthy-ae2b329   23984f5b43f6        18 months ago       135MB

[root@shkf6-245 armory]# docker tag 23984f5b43f6 harbor.od.com/armory/igor:v1.8.x
[root@shkf6-245 armory]# docker push harbor.od.com/armory/igor:v1.8.x
```

### 2.准备资源配置清单

- Deployment

```shell
[root@shkf6-245 armory]# mkdir /data/k8s-yaml/armory/igor
[root@shkf6-245 armory]# vi /data/k8s-yaml/armory/igor/dp.yaml
[root@shkf6-245 armory]# cat /data/k8s-yaml/armory/igor/dp.yaml
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  labels:
    app: armory-igor
  name: armory-igor
  namespace: armory
spec:
  replicas: 1
  revisionHistoryLimit: 7
  selector:
    matchLabels:
      app: armory-igor
  template:
    metadata:
      annotations:
        artifact.spinnaker.io/location: '"armory"'
        artifact.spinnaker.io/name: '"armory-igor"'
        artifact.spinnaker.io/type: '"kubernetes/deployment"'
        moniker.spinnaker.io/application: '"armory"'
        moniker.spinnaker.io/cluster: '"igor"'
      labels:
        app: armory-igor
    spec:
      containers:
      - name: armory-igor
        image: harbor.od.com/armory/igor:v1.8.x
        imagePullPolicy: IfNotPresent
        command:
        - bash
        - -c
        args:
        - bash /opt/spinnaker/config/default/fetch.sh && cd /home/spinnaker/config
          && /opt/igor/bin/igor
        ports:
        - containerPort: 8088
          protocol: TCP
        env:
        - name: IGOR_PORT_MAPPING
          value: -8088:8088
        - name: JAVA_OPTS
          value: -Xmx1000M
        envFrom:
        - configMapRef:
            name: init-env
        livenessProbe:
          failureThreshold: 3
          httpGet:
            path: /health
            port: 8088
            scheme: HTTP
          initialDelaySeconds: 600
          periodSeconds: 3
          successThreshold: 1
          timeoutSeconds: 1
        readinessProbe:
          failureThreshold: 3
          httpGet:
            path: /health
            port: 8088
            scheme: HTTP
          initialDelaySeconds: 180
          periodSeconds: 5
          successThreshold: 5
          timeoutSeconds: 1
        volumeMounts:
        - mountPath: /etc/podinfo
          name: podinfo
        - mountPath: /opt/spinnaker/config/default
          name: default-config
        - mountPath: /opt/spinnaker/config/custom
          name: custom-config
      imagePullSecrets:
      - name: harbor
      securityContext:
        runAsUser: 0
      volumes:
      - configMap:
          defaultMode: 420
          name: custom-config
        name: custom-config
      - configMap:
          defaultMode: 420
          name: default-config
        name: default-config
      - downwardAPI:
          defaultMode: 420
          items:
          - fieldRef:
              apiVersion: v1
              fieldPath: metadata.labels
            path: labels
          - fieldRef:
              apiVersion: v1
              fieldPath: metadata.annotations
            path: annotations
        name: podinfo
```

- Service

```shell
[root@shkf6-245 armory]# vi /data/k8s-yaml/armory/igor/svc.yaml
[root@shkf6-245 armory]# cat /data/k8s-yaml/armory/igor/svc.yaml
apiVersion: v1
kind: Service
metadata:
  name: armory-igor
  namespace: armory
spec:
  ports:
  - port: 8088
    protocol: TCP
    targetPort: 8088
  selector:
    app: armory-igor
```

### 3.应用资源配置清单

```shell
[root@shkf6-245 armory]# kubectl apply -f ./igor/
deployment.extensions/armory-igor created
service/armory-igor created
```

## 8.部署api提供组件–Gate

### 1.准备镜像

```shell
[root@shkf6-245 armory]# docker pull docker.io/armory/gate-armory:dfafe73-release-1.8.x-5d505ca

[root@shkf6-245 armory]# docker images | grep gate
armory/gate-armory                                dfafe73-release-1.8.x-5d505ca               b092d4665301        18 months ago       179MB

[root@shkf6-245 armory]# docker tag b092d4665301 harbor.od.com/armory/gate:v1.8.x
[root@shkf6-245 armory]# docker push harbor.od.com/armory/gate:v1.8.x
```

### 2.准备资源配置清单

- Deployment

```shell
[root@shkf6-245 armory]# mkdir /data/k8s-yaml/armory/gate
[root@shkf6-245 armory]# vi /data/k8s-yaml/armory/gate/dp.yaml
[root@shkf6-245 armory]# cat /data/k8s-yaml/armory/gate/dp.yaml
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  labels:
    app: armory-gate
  name: armory-gate
  namespace: armory
spec:
  replicas: 1
  revisionHistoryLimit: 7
  selector:
    matchLabels:
      app: armory-gate
  template:
    metadata:
      annotations:
        artifact.spinnaker.io/location: '"armory"'
        artifact.spinnaker.io/name: '"armory-gate"'
        artifact.spinnaker.io/type: '"kubernetes/deployment"'
        moniker.spinnaker.io/application: '"armory"'
        moniker.spinnaker.io/cluster: '"gate"'
      labels:
        app: armory-gate
    spec:
      containers:
      - name: armory-gate
        image: harbor.od.com/armory/gate:v1.8.x
        imagePullPolicy: IfNotPresent
        command:
        - bash
        - -c
        args:
        - bash /opt/spinnaker/config/default/fetch.sh gate && cd /home/spinnaker/config
          && /opt/gate/bin/gate
        ports:
        - containerPort: 8084
          name: gate-port
          protocol: TCP
        - containerPort: 8085
          name: gate-api-port
          protocol: TCP
        env:
        - name: GATE_PORT_MAPPING
          value: -8084:8084
        - name: GATE_API_PORT_MAPPING
          value: -8085:8085
        - name: JAVA_OPTS
          value: -Xmx1000M
        envFrom:
        - configMapRef:
            name: init-env
        livenessProbe:
          exec:
            command:
            - /bin/bash
            - -c
            - wget -O - http://localhost:8084/health || wget -O - https://localhost:8084/health
          failureThreshold: 5
          initialDelaySeconds: 600
          periodSeconds: 5
          successThreshold: 1
          timeoutSeconds: 1
        readinessProbe:
          exec:
            command:
            - /bin/bash
            - -c
            - wget -O - http://localhost:8084/health?checkDownstreamServices=true&downstreamServices=true
              || wget -O - https://localhost:8084/health?checkDownstreamServices=true&downstreamServices=true
          failureThreshold: 3
          initialDelaySeconds: 180
          periodSeconds: 5
          successThreshold: 10
          timeoutSeconds: 1
        volumeMounts:
        - mountPath: /etc/podinfo
          name: podinfo
        - mountPath: /opt/spinnaker/config/default
          name: default-config
        - mountPath: /opt/spinnaker/config/custom
          name: custom-config
      imagePullSecrets:
      - name: harbor
      securityContext:
        runAsUser: 0
      volumes:
      - configMap:
          defaultMode: 420
          name: custom-config
        name: custom-config
      - configMap:
          defaultMode: 420
          name: default-config
        name: default-config
      - downwardAPI:
          defaultMode: 420
          items:
          - fieldRef:
              apiVersion: v1
              fieldPath: metadata.labels
            path: labels
          - fieldRef:
              apiVersion: v1
              fieldPath: metadata.annotations
            path: annotations
        name: podinfo
```

- Service

```shell
[root@shkf6-245 armory]# vi /data/k8s-yaml/armory/gate/svc.yaml
[root@shkf6-245 armory]# cat /data/k8s-yaml/armory/gate/svc.yaml
apiVersion: v1
kind: Service
metadata:
  name: armory-gate
  namespace: armory
spec:
  ports:
  - name: gate-port
    port: 8084
    protocol: TCP
    targetPort: 8084
  - name: gate-api-port
    port: 8085
    protocol: TCP
    targetPort: 8085
  selector:
    app: armory-gate
```

### 3.应用资源配置清单

```shell
[root@shkf6-245 armory]# kubectl apply -f ./gate/
deployment.extensions/armory-gate created
service/armory-gate created
```

### 4.检查

```shell
root@nginx-armory-cb6cfdfd-5lcsr:/# curl armory-gate:8084/health
{"status":"UP"}
```

## 9.部署前端网站项目–Deck

### 1.准备docker镜像

```shell
[root@shkf6-245 armory]# docker pull docker.io/armory/deck-armory:d4bf0cf-release-1.8.x-0a33f94

[root@shkf6-245 armory]# docker images |grep deck
armory/deck-armory                                d4bf0cf-release-1.8.x-0a33f94               9a87ba3b319f        18 months ago       518MB

[root@shkf6-245 armory]# docker tag 9a87ba3b319f harbor.od.com/armory/deck:v1.8.x
[root@shkf6-245 armory]# docker push harbor.od.com/armory/deck:v1.8.x
```

### 2.准备资源配置清单

- Deployment

```shell
[root@shkf6-245 armory]# mkdir /data/k8s-yaml/armory/deck
[root@shkf6-245 armory]# vi /data/k8s-yaml/armory/deck/dp.yaml
[root@shkf6-245 armory]# cat /data/k8s-yaml/armory/deck/dp.yaml
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  labels:
    app: armory-deck
  name: armory-deck
  namespace: armory
spec:
  replicas: 1
  revisionHistoryLimit: 7
  selector:
    matchLabels:
      app: armory-deck
  template:
    metadata:
      annotations:
        artifact.spinnaker.io/location: '"armory"'
        artifact.spinnaker.io/name: '"armory-deck"'
        artifact.spinnaker.io/type: '"kubernetes/deployment"'
        moniker.spinnaker.io/application: '"armory"'
        moniker.spinnaker.io/cluster: '"deck"'
      labels:
        app: armory-deck
    spec:
      containers:
      - name: armory-deck
        image: harbor.od.com/armory/deck:v1.8.x
        imagePullPolicy: IfNotPresent
        command:
        - bash
        - -c
        args:
        - bash /opt/spinnaker/config/default/fetch.sh && /entrypoint.sh
        ports:
        - containerPort: 9000
          protocol: TCP
        envFrom:
        - configMapRef:
            name: init-env
        livenessProbe:
          failureThreshold: 3
          httpGet:
            path: /
            port: 9000
            scheme: HTTP
          initialDelaySeconds: 180
          periodSeconds: 3
          successThreshold: 1
          timeoutSeconds: 1
        readinessProbe:
          failureThreshold: 5
          httpGet:
            path: /
            port: 9000
            scheme: HTTP
          initialDelaySeconds: 30
          periodSeconds: 3
          successThreshold: 5
          timeoutSeconds: 1
        volumeMounts:
        - mountPath: /etc/podinfo
          name: podinfo
        - mountPath: /opt/spinnaker/config/default
          name: default-config
        - mountPath: /opt/spinnaker/config/custom
          name: custom-config
      imagePullSecrets:
      - name: harbor
      volumes:
      - configMap:
          defaultMode: 420
          name: custom-config
        name: custom-config
      - configMap:
          defaultMode: 420
          name: default-config
        name: default-config
      - downwardAPI:
          defaultMode: 420
          items:
          - fieldRef:
              apiVersion: v1
              fieldPath: metadata.labels
            path: labels
          - fieldRef:
              apiVersion: v1
              fieldPath: metadata.annotations
            path: annotations
        name: podinfo
```

- Service

```shell
[root@shkf6-245 armory]# vi /data/k8s-yaml/armory/deck/svc.yaml
[root@shkf6-245 armory]# cat /data/k8s-yaml/armory/deck/svc.yaml
apiVersion: v1
kind: Service
metadata:
  name: armory-deck
  namespace: armory
spec:
  ports:
  - port: 80
    protocol: TCP
    targetPort: 9000
  selector:
    app: armory-deck
```

### 3.应用资源配置清单

```shell
[root@shkf6-245 armory]# kubectl apply -f ./deck/
deployment.extensions/armory-deck created
service/armory-deck created
```

## 10.部署前端代理–Nginx

### 1.准备docker镜像

```shell
[root@shkf6-245 armory]# docker pull nginx:1.12.2

[root@shkf6-245 armory]# docker images | grep nginx
goharbor/nginx-photon                             v1.8.3                                      3a016e0dc7de        3 months ago        37MB
nginx                                             1.12.2                                      4037a5562b03        20 months ago       108MB
sunrisenan/nginx                                  v1.12.2                                     4037a5562b03        20 months ago       108MB
nginx                                             1.7.9                                       84581e99d807        4 years ago         91.7MB
harbor.od.com/public/nginx                        v1.7.9                                      84581e99d807        4 years ago         91.7MB
[root@shkf6-245 armory]# docker tag 4037a5562b03 harbor.od.com/armory/nginx:v1.12.2
[root@shkf6-245 armory]# docker push harbor.od.com/armory/nginx:v1.12.2
```

### 2.准备资源配置清单

- Deployment

```shell
[root@shkf6-245 armory]# mkdir /data/k8s-yaml/armory/nginx
[root@shkf6-245 armory]# vi /data/k8s-yaml/armory/nginx/dp.yaml
[root@shkf6-245 armory]# cat /data/k8s-yaml/armory/nginx/dp.yaml
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  labels:
    app: armory-nginx
  name: armory-nginx
  namespace: armory
spec:
  replicas: 1
  revisionHistoryLimit: 7
  selector:
    matchLabels:
      app: armory-nginx
  template:
    metadata:
      annotations:
        artifact.spinnaker.io/location: '"armory"'
        artifact.spinnaker.io/name: '"armory-nginx"'
        artifact.spinnaker.io/type: '"kubernetes/deployment"'
        moniker.spinnaker.io/application: '"armory"'
        moniker.spinnaker.io/cluster: '"nginx"'
      labels:
        app: armory-nginx
    spec:
      containers:
      - name: armory-nginx
        image: harbor.od.com/armory/nginx:v1.12.2
        imagePullPolicy: Always
        command:
        - bash
        - -c
        args:
        - bash /opt/spinnaker/config/default/fetch.sh nginx && nginx -g 'daemon off;'
        ports:
        - containerPort: 80
          name: http
          protocol: TCP
        - containerPort: 443
          name: https
          protocol: TCP
        - containerPort: 8085
          name: api
          protocol: TCP
        livenessProbe:
          failureThreshold: 3
          httpGet:
            path: /
            port: 80
            scheme: HTTP
          initialDelaySeconds: 180
          periodSeconds: 3
          successThreshold: 1
          timeoutSeconds: 1
        readinessProbe:
          failureThreshold: 3
          httpGet:
            path: /
            port: 80
            scheme: HTTP
          initialDelaySeconds: 30
          periodSeconds: 3
          successThreshold: 5
          timeoutSeconds: 1
        volumeMounts:
        - mountPath: /opt/spinnaker/config/default
          name: default-config
        - mountPath: /etc/nginx/conf.d
          name: custom-config
      imagePullSecrets:
      - name: harbor
      volumes:
      - configMap:
          defaultMode: 420
          name: custom-config
        name: custom-config
      - configMap:
          defaultMode: 420
          name: default-config
        name: default-config
```

- Service

```shell
[root@shkf6-245 armory]# vi /data/k8s-yaml/armory/nginx/svc.yaml
[root@shkf6-245 armory]# cat /data/k8s-yaml/armory/nginx/svc.yaml
apiVersion: v1
kind: Service
metadata:
  name: armory-nginx
  namespace: armory
spec:
  ports:
  - name: http
    port: 80
    protocol: TCP
    targetPort: 80
  - name: https
    port: 443
    protocol: TCP
    targetPort: 443
  - name: api
    port: 8085
    protocol: TCP
    targetPort: 8085
  selector:
    app: armory-nginx
```

- Ingress

```shell
[root@shkf6-245 armory]# vi /data/k8s-yaml/armory/nginx/ingress.yaml
[root@shkf6-245 armory]# cat /data/k8s-yaml/armory/nginx/ingress.yaml
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  labels:
    app: spinnaker
    web: spinnaker.od.com
  name: armory-nginx
  namespace: armory
spec:
  rules:
  - host: spinnaker.od.com
    http:
      paths:
      - backend:
          serviceName: armory-nginx
          servicePort: 80
```

### 3.应用资源配置清单

```shell
[root@shkf6-245 armory]# kubectl apply -f ./nginx/
deployment.extensions/armory-nginx created
ingress.extensions/armory-nginx created
service/armory-nginx created
```

## 11.域名解析

```shell
[root@shkf6-241 ~]# tail -1 /var/named/od.com.zone 
spinnaker          A    192.168.6.66
```

## 12.打开浏览器

http://spinnaker.od.com/

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_9df403d8699c4e7ead8a8f4b2b1dfe64_r.png)

# 第五章：使用spinnaker结合jenkins构建镜像

## 1.使用spinnaker前期准备

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_d603a00304cf1ce1d321e0649a045e30_r.png)

可以到看到aopplo、dubbo，这是前面在CloudDriver中custom-config.yaml 里配置了管理prod和test名称空间。

把dubbo项目从k8s集群中移除

```shell
[root@shkf6-245 k8s-yaml]# kubectl delete -f test/dubbo-demo-consumer/dp.yaml 
deployment.extensions "dubbo-demo-consumer" deleted

[root@shkf6-245 k8s-yaml]# kubectl delete -f test/dubbo-demo-service/deployment.yaml 
deployment.extensions "dubbo-demo-service" deleted


[root@shkf6-245 k8s-yaml]# kubectl delete -f prod/dubbo-demo-consumer/dp.yaml 
deployment.extensions "dubbo-demo-consumer" deleted

[root@shkf6-245 k8s-yaml]# kubectl delete -f prod/dubbo-demo-service/deployment.yaml 
deployment.extensions "dubbo-demo-service" deleted
```

刷新页面就没有了

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_4c9591eaa08903d5e31af2b6a379fbac_r.png)

## 2.使用spinnaker创建dubbo

1.创建应用集

- Application -> Actions -> CreateApplication

  - Name

    > test0dubbo

  - Owner Emial

    > [1210353303@qq.com](mailto:1210353303@qq.com)

- Create

效果图

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_c194c272a426d72791426f9383434707_r.png)

2.创建pipelines

- PIPELINES -> Configure a new pipeline

  - Type

    > Pipeline

  - Pipeline Name

    > dubbo-demo-service

- create

3.配置加4个参数

解释：Triggers触发器 Parameters选项 Artifacts手工 Notifications通知 Description描述

- Parameters第一个参数

  - Name

    > git_ver

  - Required

    > 打勾

- Parameters第二个参数

  - Name

    > add_tag

  - Required

    > 打勾

- Parameters第三个参数

  - Name

    > app_name

  - Required

    > 打勾

  - Default Value

    > dubbo-demo-service

- Parameters第四个参数

  - Name

    > image_name

  - Required

    > 打勾

  - Default Value

    > app/dubbo-demo-service

- Save Changes

4.增加一个流水线的阶段

- Add stage

  - Type

    > Jenkins

  - Master

    > jenkins-admin

  - Job

    > dubbo-demo

  - add_tag

    > ${ parameters.add_tag }

  - app_name

    > ${ parameters.app_name }

  - base_image

    > base/jre8:8u112_with_logs

  - git_repo

    > https://github.com/sunrisenan/dubbo-demo-service.git

  - git_ver

    > ${ parameters.git_ver }

  - image_name

    > ${ parameters.image_name }

  - target_dir

    > ./dubbo-server/target

- Save Changes

- PIPELIENS

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_82f4f2bd0744cdb820dd6dd4c7e702ad_r.png)

5.运行流水线

- Start Manual Execution

  - git_ver

    > apollo

  - add_tag

    > 200113_1650

- RUN

效果演示：

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_fcdc8f7b827ea8dbdffbc512429b74d3_r.png)

# 第六章：实战配置、使用Spinnaker配置dubbo服务提供者项目

- Application -> PIPELIENS -> Configure -> Add stage

- Basic Settings

  - Type

    > Deploy

  - Stage Name

    > Deploy

- Add server group

  - Account

    > cluster-admin

  - Namespace

    > test

  - Detail

    > [项目名]dubbo-dubbo-service

  - Containers

    > harbor.od.com/app/dubbo-demo-service:apollo_200113_1843
    > harbor.od.com/infra/filebeat:v7.4.0

  - Strategy

    > None

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_51361b669f47d84374b5ebf03d0bb1db_r.png)

- Deployment

  - Deployment

    > 打勾

  - Strategy

    > RollingUpdate

  - History Limit

    > 7

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_b5557ae68dae1ab2d2dc4b13e40cf3d3_r.png)

- Replices

  - Capacity

    > 1 (默认起一份)

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_c80935d88ddfa665eedb9acfc7d40bd9_r.png)

- Volume Sources -> Add Volume Source

  - Volume Source

    > EMPTYPE

  - Name

    > logm

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_9a7024d746ac0861280ff04def849691_r.png)

- Advanced Settings

  - DNS Policy

    > ClusterFirst

  - Termination Grace Period

    > 30

  - Pod Annotations

    - Key

      > blackbox_scheme

    - Value

      > tcp

    - Key

      > blackbox_port

    - Value

      > 20880

    - Key

      > prometheus_io_scrape

    - Value

      > true

    - Key

      > prometheus_io_path

    - Value

      > /

    - Key

      > prometheus_io_port

    - Value

      > 12346

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_b62fd52fa894e42d6ba18915658e9a79_r.png)

------

配置第一个容器

- Container -> Basic Settings -> Environment Variables -> Add Environment Variables

  - Name

    > JAR_BALL

  - Source

    > Explicit

  - Value

    > dubbo-server.jar

  - Name

    > C_OPTS

  - Source

    > Explicit

  - Value

    > -Denv=fat -Dapollo.meta=[http://config-test.od.com](http://config-test.od.com/)

- Volume Mounts

  - Source Name

    > logm

  - Mount Path

    > /opt/logs

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_1d996a12716cd6ee7b82fe3937b84610_r.png)

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_4fbd0116dcd9e4ba107653df0fd6a5d1_r.png)

------

配置第二个容器

- Container -> Basic Settings -> Environment Variables -> Add Environment Variables

  - Name

    > ENV

  - Source

    > Explicit

  - Value

    > test

  - Name

    > PROJ_NAME

  - Source

    > Explicit

  - Value

    > dubbo-demo-service

- VolumeMounts

  - Source Name

    > logm

  - Mount Path

    > /logm

- Add

- Save Changes

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_c43fb348d9b65659f58f896b69b0e34e_r.png)

调整JSON，改为通用型

```shell
"imageId": "harbor.od.com/${parameters.image_name}:${parameters.git_ver}_${parameters.add_tag}",
"registry": "harbor.od.com",
"repository": "${parameters.image_name}",
"tag": "${parameters.git_ver}_${parameters.add_tag}"
```

- Save Changes

# 第七章：实战配置、使用Spinnaker配置dubbo服务消费者项目

1.用test0dubbo应用集

2.创建pipelines

- PIPELINES -> Create

  - Type

    > Pipeline

  - Pipeline Name 注：名字要和gitlab项目名字一致

    > dubbo-demo-web

  - Copy From

    > None

- create

3.配置加4个参数

- Parameters第一个参数

  - Name

    > git_ver

  - Required

    > 打勾

- Parameters第二个参数

  - Name

    > add_tag

  - Required

    > 打勾

- Parameters第三个参数

  - Name

    > app_name

  - Required

    > 打勾

  - Default Value

    > dubbo-demo-web

- Parameters第四个参数

  - Name

    > image_name

  - Required

    > 打勾

  - Default Value

    > app/dubbo-demo-web

- Save Changes

4.增加一个流水线的阶段

- Add stage

  - Type

    > Jenkins

  - Master

    > jenkins-admin

  - Job

    > tomcat-demo

  - add_tag

    > ${ parameters.add_tag }

  - app_name

    > ${ parameters.app_name }

  - base_image

    > base/jre8:8u112_with_logs

  - git_repo

    > https://github.com/sunrisenan/dubbo-demo-web.git

  - git_ver

    > ${ parameters.git_ver }

  - image_name

    > ${ parameters.image_name }

  - target_dir

    > ./dubbo-server/target

- Save Changes

5.server

- Application -> INFRASTRUCTURE -> LOAD BALANCERS -> Create Load Balancer

  - Basic Settings

    - Account

      > cluster-admin

    - Namespace

      > test

    - Detail

      > demo-web

  - Ports

    - Name

      > http

    - Port

      > 80

    - Target Port

      > 8080

- create

6.ingress

- Application -> INFRASTRUCTURE -> FIREWALLS -> Create Firewall

- Basic Settings

  - Account

    > cluster-admin

  - Namespace

    > test

  - Detail

    > dubbo-web

- Rules -> Add New Rule

  - Host

    > demo-test.od.com

- Add New Path

  - Load Balancer

    > test0dubbo-web

  - Path

    > /

  - Port

    > 80

- create

7.deploy

- Application -> PIPELIENS -> Configure -> Add stage

- Basic Settings

  - Type

    > Deploy

  - Stage Name

    > Deploy

- Add server group

  - Account

    > cluster-admin

  - Namespace

    > test

  - Detail

    > [项目名]dubbo-dubbo-web

  - Containers

    > harbor.od.com/app/dubbo-demo-web:tomcat_200114_1613
    > harbor.od.com/infra/filebeat:v7.4.0

  - Strategy

    > None

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_194193f51d8c0c1a10500f46dee63124_r.png)

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_0294ea1019dc9f37f4b6721954066cb9_r.png)

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_956bb5bf57e27ac081b1a783c818bdac_r.png)

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_57ecb02e5317a6b5923b68e6aed0d44e_r.png)

------

配置第一个容器

- Container -> Basic Settings -> Environment Variables -> Add Environment Variables

  - Name

    > C_OPTS

  - Source

    > Explicit

  - Value

    > -Denv=fat -Dapollo.meta=[http://config-test.od.com](http://config-test.od.com/)

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_7743440833efd8ab6f62f13a7cefb4a6_r.png)

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_3817ad67b09b38beeaf2e67009040022_r.png)

------

配置第二个容器

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_cb44e54ad656e5c9d6e3dc6f9b99d94a_r.png)

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_de5e2ffe5e897b520db216d030b5f6eb_r.png)

- Add

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_c5c3e7ca526ad2f9143cc0537578aca9_r.png)

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_6e839264b45ec18f3285a937f4ad2614_r.png)

更改为：

```shell
"imageId": "harbor.od.com/${parameters.image_name}:${parameters.git_ver}_${parameters.add_tag}",
"registry": "harbor.od.com",
"repository": "${parameters.image_name}",
"tag": "${parameters.git_ver}_${parameters.add_tag}"
```

- Save Changes

第七章：实战使用Spinnaker进行灰度发布、金丝雀发布