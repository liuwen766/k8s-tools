- 针对LoadBalancer 类型的 Service，每个 Service 都要有一个负载均衡服务，所以这个做法实际上既浪费成本又高。
- 其实，Kubernetes 中可以内置一个全局的负载均衡器。通过访问的 URL，把请求转发给不同的后端 Service。这种全局的、为了代理不同后端
  Service 而设置的负载均衡服务，就是 Kubernetes 里的 Ingress 服务。

- Ingress 的功能其实很容易理解：所谓 Ingress，就是 Service 的“Service”。
- 如何能使用 Kubernetes 的 Ingress 来创建一个统一的负载均衡器，从而实现当用户访问不同的域名时，能够访问到不同的 Deployment
  呢？【example-ingress.yaml】


- 所谓 Ingress 对象，其实就是 Kubernetes 项目对“反向代理”的一种抽象。一个Ingress对象的主要内容，实际上就是一个“反向代理”服务
  （比如：Nginx）的配置文件的描述。而这个代理服务对应的转发规则，就是IngressRule。
- 一个 Nginx Ingress Controller 为你提供的服务，其实是一个可以根据 Ingress 对象和被代理后端 Service 的变化，来自动进行更新的
  Nginx 负载均衡器。

- 在实际的使用中，只需要从社区里选择一个具体的 Ingress Controller，把它部署在 Kubernetes 集群里即可。然后，这个 Ingress
  Controller 会根据定义的 Ingress 对象，提供对应的代理能力。目前——业界常用的各种反向代理项目，比如
  Nginx、HAProxy、Envoy、Traefik等，都已经为 Kubernetes 专门维护了对应的 Ingress Controller。
- eg：部署 Nginx Ingress Controller
- 原理：当一个新的 Ingress 对象由用户创建后，nginx-ingress-controller 就会根据 Ingress 对象里定义的内容，生成一份对应的
  Nginx 配置文件（/etc/nginx/nginx.conf），并使用这个配置文件启动一个 Nginx 服务。
- 一个 Nginx Ingress Controller 提供的服务，其实是一个可以根据 Ingress 对象和被代理后端 Service 的变化，来自动进行更新的
  Nginx 负载均衡器。

```shell
# 在下面的 YAML 文件中，定义了一个使用 nginx-ingress-controller 镜像的 Pod。
# 这个 Pod 本身，就是一个监听 Ingress 对象以及它所代理的后端 Service 变化的控制器。
$ kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/mandatory.yaml
# 为了让用户能够用到这个 Nginx，就需要创建一个 Service 来把 Nginx Ingress Controller 管理的 Nginx 服务暴露出去。
# 这个 Service 的唯一工作，就是将所有携带 ingress-nginx 标签的 Pod 的 80 和 433 端口暴露出去。
$ kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/provider/baremetal/service-nodeport.yaml

# 记录下这个 Service 的访问入口，即：宿主机的地址和 NodePort 的端口，如下所示：
$ kubectl get svc -n ingress-nginx
NAME            TYPE       CLUSTER-IP     EXTERNAL-IP   PORT(S)                      AGE
ingress-nginx   NodePort   10.105.72.96   <none>        80:30044/TCP,443:31453/TCP   3h

$ IC_IP=10.168.0.2 # 任意一台宿主机的地址
$ IC_HTTPS_PORT=31453 # NodePort 端口
```

- 在 Ingress Controller 和它所需要的 Service 部署完成后，我们就可以使用它。

```shell
$ kubectl create -f cafe.yaml
$ kubectl create -f cafe-secret.yaml
$ kubectl create -f cafe-ingress.yaml
$ kubectl get ingress
NAME           HOSTS              ADDRESS   PORTS     AGE
cafe-ingress   cafe.example.com             80, 443   2h
 
$ kubectl describe ingress cafe-ingress
Name:             cafe-ingress
Namespace:        default
Address:          
Default backend:  default-http-backend:80 (<none>)
TLS:
  cafe-secret terminates cafe.example.com
Rules:
  Host              Path  Backends
  ----              ----  --------
  cafe.example.com  
                    /tea      tea-svc:80 (<none>)
                    /coffee   coffee-svc:80 (<none>)
Annotations:
Events:
  Type    Reason  Age   From                      Message
  ----    ------  ----  ----                      -------
  Normal  CREATE  4m    nginx-ingress-controller  Ingress default/cafe-ingress
```

- 如果请求没有匹配到任何一条 IngressRule，那么会默认返回一个 Nginx 的 404 页面。不过Ingress Controller 也允许通过 Pod
  启动命令里的–default-backend-service 参数，设置一条默认规则，比如：–default-backend-service=nginx-default-backend。


- 目前，Ingress 只能工作在七层，而 Service 只能工作在四层。所以当你想要在 Kubernetes 里为应用进行 TLS 配置等 HTTP
  相关的操作时，都必须通过 Ingress 来进行。

- Kubernetes 提出 Ingress 概念的原因其实也非常容易理解，有了 Ingress 这个抽象，用户就可以根据自己的需求来自由选择 Ingress
  Controller。比如，如果应用对代理服务的中断非常敏感，那么就应该考虑选择类似于 Traefik 这样支持“热加载”的 Ingress
  Controller 实现。




