

- 1、安装go 
- 2、安装 operator-sdk 到 /uer/local/bin



- 1、operator-sdk init --domain testdomain --repo github.com/test/test 
- 2、operator-sdk create api --group group --version v1alpha1 --kind Test --resource --controller
- 3、修改controller【vim controllers/test_controller.go】
  - 安装 kustomize 到 当前根目录bin下
- 4、make generate
- 5、make manifests
- 6、make install
- 7、修改Dockerfile文件
  - 添加如下两行
  - RUN go env -w GO111MODULE=on 
  - RUN go env -w GOPROXY=https://goproxy.cn,direct
- 8、make docker-build IMG=abc:v1
- 9、修改config/default/manager_auth_proxy_patch.yaml文件 
  - gcr.io/kubebuilder/kube-rbac-proxy:v0.13.0镜像下载不到，换个镜像
- 10、make deploy IMG=abc:v1
- 11、生成CR
   - 验证crd资源 kubectl get crd|grep test
```yaml
apiVersion: tests.group.testdomain/v1alpha1
Kind: Test
metadata:
   name: test
   namespace: my-operator-system
spec:
   foo: hello-world
```
