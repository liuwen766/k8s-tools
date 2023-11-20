# 前情回顾

1.使用二进制安装部署K8S的要点：

- 基础设施环境准备好
  - centos7.6系统（内核在3.8x以上）
  - 关闭SELinux，关闭firewalld服务
  - 时间同步（chronyd）
  - 调整Base源，epel源
  - 内核优化（文件描述符大小，内核转发，等等。。。）
- 安装部署bind9内网DNS系统
- 安装部署docker的私有仓库–harbor
- 准备证书签发环境–cfssl
- 安装部署主控节点服务（4个）
  - Etd
  - Apiserver
  - Controller-manager
  - Scheduler
- 安装部署运算节点服务（2个）
  - Kubelet
  - Kube-proxy

2.关于cfssl工具：

- cfssl：证书签发的主要工具
- cfssl-json：讲cfssl生成的证书（json格式）变为文件承载式证书
- cfssl-certinfo：验证证书的信息

把证书`apiserver.pem`的信息列出来：

```shell
[root@shkf6-245 certs]# cfssl-certinfo -cert apiserver.pem 
{
  "subject": {
    "common_name": "k8s-apiserver",
    "country": "CN",
    "organization": "od",
    "organizational_unit": "ops",
    "locality": "beijing",
    "province": "beijing",
    "names": [
      "CN",
      "beijing",
      "beijing",
      "od",
      "ops",
      "k8s-apiserver"
    ]
  },
  "issuer": {
    "common_name": "OldboyEdu",
    "country": "CN",
    "organization": "od",
    "organizational_unit": "ops",
    "locality": "beijing",
    "province": "beijing",
    "names": [
      "CN",
      "beijing",
      "beijing",
      "od",
      "ops",
      "OldboyEdu"
    ]
  },
  "serial_number": "654380197932285157915575843459046152400709392018",
  "sans": [
    "kubernetes.default",
    "kubernetes.default.svc",
    "kubernetes.default.svc.cluster",
    "kubernetes.default.svc.cluster.local",
    "127.0.0.1",
    "10.96.0.1",
    "192.168.0.1",
    "192.168.6.66",
    "192.168.6.243",
    "192.168.6.244",
    "192.168.6.245"
  ],
  "not_before": "2019-11-18T05:54:00Z",
  "not_after": "2039-11-13T05:54:00Z",
  "sigalg": "SHA256WithRSA",
  "authority_key_id": "72:13:FC:72:0:8:4:A0:5A:90:B2:E2:10:FA:6A:7E:7A:2D:2C:22",
  "subject_key_id": "FC:36:3:9F:2C:D0:9F:D4:1E:F1:A4:56:59:41:A8:29:81:35:38:F7",
  "pem": "-----BEGIN CERTIFICATE-----\nMIIEdTCCA12gAwIBAgIUcp9sROX/bzXK8oiG5ImDB3n8lpIwDQYJKoZIhvcNAQEL\nBQAwYDELMAkGA1UEBhMCQ04xEDAOBgNVBAgTB2JlaWppbmcxEDAOBgNVBAcTB2Jl\naWppbmcxCzAJBgNVBAoTAm9kMQwwCgYDVQQLEwNvcHMxEjAQBgNVBAMTCU9sZGJv\neUVkdTAeFw0xOTExMTgwNTU0MDBaFw0zOTExMTMwNTU0MDBaMGQxCzAJBgNVBAYT\nAkNOMRAwDgYDVQQIEwdiZWlqaW5nMRAwDgYDVQQHEwdiZWlqaW5nMQswCQYDVQQK\nEwJvZDEMMAoGA1UECxMDb3BzMRYwFAYDVQQDEw1rOHMtYXBpc2VydmVyMIIBIjAN\nBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAzkm60A1MQiWsPgWHrp4JjJKVFAEP\ndzm+f4ZB3GsZnuiuMo3ypygClwsjDcPZL0+zKaAmwOMxa3cSmbgY8rYkbqGM8Tdd\nBN4ns0pDuwl5EbcMAfyL1ZsgpxctMwtaCO1wR2N6fAhk1BZt6VH7a9ruw83UKfDK\n3W77JxLdcgAEDdaizddQPhOE3W2BMJztLI03PIIQ2yZ3vikfZQjYwZLAWBBDpoOJ\neMm9J3RS0FL8YXTefjG02An4Z1BSTlTDOPogUxxAMAJ1jaaOXSpUSj/qACYpHx7N\n9zgKmkMKby2N+VfJs60MD60jn0CJXALJ1U/lyAieX8KRW3IEW1L/xhVhxwIDAQAB\no4IBITCCAR0wDgYDVR0PAQH/BAQDAgWgMBMGA1UdJQQMMAoGCCsGAQUFBwMBMAwG\nA1UdEwEB/wQCMAAwHQYDVR0OBBYEFPw2A58s0J/UHvGkVllBqCmBNTj3MB8GA1Ud\nIwQYMBaAFHIT/HIACASgWpCy4hD6an56LSwiMIGnBgNVHREEgZ8wgZyCEmt1YmVy\nbmV0ZXMuZGVmYXVsdIIWa3ViZXJuZXRlcy5kZWZhdWx0LnN2Y4Iea3ViZXJuZXRl\ncy5kZWZhdWx0LnN2Yy5jbHVzdGVygiRrdWJlcm5ldGVzLmRlZmF1bHQuc3ZjLmNs\ndXN0ZXIubG9jYWyHBH8AAAGHBApgAAGHBMCoAAGHBMCoBkKHBMCoBvOHBMCoBvSH\nBMCoBvUwDQYJKoZIhvcNAQELBQADggEBAAIKX7PJXI5Uj/nn1nfp4M2HTk1m+8OH\nzTwRVy1igM+2ull+S7rXO4MN9Sw16Coov6MyHviWXLWJgGTypOenz+tTQkdjzHLF\nrkDxM6Fu4b5pm4W1H5QQaoFfY7XaCGnoD0Mf3rMsG8Pi5n3e7pb2tw73ebxwHQx7\nnl43fOoRAyDVMmju5BKG8QOA5dfW3qi2BpCCM7KCgSVq/U8URaI3zBm6Mfm7eUnI\nhwnpufar08JCLVtoduKelbdaaBSEmDR+7bl0aQ5YHwAHuQZRxQB4qG+QbiTabuoV\npXATIAmGLqretFGp9rlsvh6kxIKw+NJ8k2DBpzOCzJALCYIbHhv40oA=\n-----END CERTIFICATE-----\n"
}
```

检查`h5.mcake.com`域名证书的信息

```shell
[root@shkf6-245 certs]# cfssl-certinfo -domain h5.mcake.com
{
  "subject": {
    "common_name": "h5.mcake.com",
    "names": [
      "h5.mcake.com"
    ]
  },
  "issuer": {
    "common_name": "Encryption Everywhere DV TLS CA - G1",
    "country": "US",
    "organization": "DigiCert Inc",
    "organizational_unit": "www.digicert.com",
    "names": [
      "US",
      "DigiCert Inc",
      "www.digicert.com",
      "Encryption Everywhere DV TLS CA - G1"
    ]
  },
  "serial_number": "20945969268560300749908123751452433550",
  "sans": [
    "h5.mcake.com"
  ],
  "not_before": "2019-07-30T00:00:00Z",
  "not_after": "2020-07-29T12:00:00Z",
  "sigalg": "SHA256WithRSA",
  "authority_key_id": "55:74:4F:B2:72:4F:F5:60:BA:50:D1:D7:E6:51:5C:9A:1:87:1A:D7",
  "subject_key_id": "3F:86:11:8A:11:D:24:58:D2:CF:33:58:20:52:4A:A4:AE:94:90:33",
  "pem": "-----BEGIN CERTIFICATE-----\nMIIFgTCCBGmgAwIBAgIQD8IMBHJkmkUGUQ2ohH7cjjANBgkqhkiG9w0BAQsFADBu\nMQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3\nd3cuZGlnaWNlcnQuY29tMS0wKwYDVQQDEyRFbmNyeXB0aW9uIEV2ZXJ5d2hlcmUg\nRFYgVExTIENBIC0gRzEwHhcNMTkwNzMwMDAwMDAwWhcNMjAwNzI5MTIwMDAwWjAX\nMRUwEwYDVQQDEwxoNS5tY2FrZS5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAw\nggEKAoIBAQCar3JOaRYZEehyIH4SACHo+IANki25XaazjJFYT4MUzU8u+fWRe21r\npnlosxT5z4sGyuk6DxhNdPi0e1zLXeEcNGwqDsO5NPf7BWx9u+ljxdlcWXYO/aY8\nCHyuFPwkrKIPqLsLgy47U7Wbm8WamqoF4ywKoCTQnZGfRIWCMzkLgFRP7a0IxUdP\nRtggxSXNjCcSD5KkUJLfSPMPKLR8pZ9pVRczaPj4Y6vLpryxU2HqVecW/+CYZq/a\nkZ3o2LAJo9Z+aNMCMWM0IMXv1teb81/M06qD4OYXLGOWzy/pWXYKA4m+NOj/fJLC\nSLjpLVlkB/CJAQPP0P4+Idqsc5gg+22bAgMBAAGjggJwMIICbDAfBgNVHSMEGDAW\ngBRVdE+yck/1YLpQ0dfmUVyaAYca1zAdBgNVHQ4EFgQUP4YRihENJFjSzzNYIFJK\npK6UkDMwFwYDVR0RBBAwDoIMaDUubWNha2UuY29tMA4GA1UdDwEB/wQEAwIFoDAd\nBgNVHSUEFjAUBggrBgEFBQcDAQYIKwYBBQUHAwIwTAYDVR0gBEUwQzA3BglghkgB\nhv1sAQIwKjAoBggrBgEFBQcCARYcaHR0cHM6Ly93d3cuZGlnaWNlcnQuY29tL0NQ\nUzAIBgZngQwBAgEwgYAGCCsGAQUFBwEBBHQwcjAkBggrBgEFBQcwAYYYaHR0cDov\nL29jc3AuZGlnaWNlcnQuY29tMEoGCCsGAQUFBzAChj5odHRwOi8vY2FjZXJ0cy5k\naWdpY2VydC5jb20vRW5jcnlwdGlvbkV2ZXJ5d2hlcmVEVlRMU0NBLUcxLmNydDAJ\nBgNVHRMEAjAAMIIBBAYKKwYBBAHWeQIEAgSB9QSB8gDwAHYApLkJkLQYWBSHuxOi\nzGdwCjw1mAT5G9+443fNDsgN3BAAAAFsQhPl9AAABAMARzBFAiEAgtcA1rcRVeTf\n3KLO54eXR3sLLDTW3XPLqj+VI+a28IICIBMBVWYIaNERBmejleyMnOoYbHULFEgi\ndi9eHiPT2sXpAHYAXqdz+d9WwOe1Nkh90EngMnqRmgyEoRIShBh1loFxRVgAAAFs\nQhPlLAAABAMARzBFAiAqnY9AE1/zVkYG6R16tO+i3ojXnM3CmBs3iXm6ywI5wQIh\nAN9jpUnfPguDaf9/LBwG8wgx8+6ybeoVv4SUUhUlRbzmMA0GCSqGSIb3DQEBCwUA\nA4IBAQBg/n6IyoEUBiIwikm+vKciqJpbOrL0urqkdTbQIwCbdxHe8s5ewULM8nJ9\n+77voDywIHaj0tLZvKxBmJUQliYKh1hthGY4IxQIsaxRNC/D3/5s2unz1yyTQaPr\n+CU0/xKpUyzh63wzS0w6/IRkXNwCaSwZWvFFR2HHJNQ9Y9NwkmyhCg7Sm2MLBTxj\nVmmjgTt5E47TiuCqYkEdH7KCoPSKh0Z6Jv46Bj0Ls5oFOZKa93QHipuHfuqmF6G/\nAsB9tfS4ATvxBb5hOxpfX6Tyv5cvFRKcAwGJxQ7fq9cuEAvkga7FkFfKD+JxdLuK\nbw3xQzHwX6kEB54Z88C/VRQ1oILw\n-----END CERTIFICATE-----\n"
}
```

3.关于kubeconfig文件：

- 这是一个K8S用户的配置文件
- 它里面含有证书信息
- 证书过期或更换，需要同步替换改文件

根据`MD5sum`计算`kubelet.kubeconfig`是否一样

```shell
[root@shkf6-243 ~]# cd /opt/kubernetes/server/bin/conf/
[root@shkf6-243 conf]# md5sum kubelet.kubeconfig 
8b2f777f58e4e36d09594db5d1d83ef2  kubelet.kubeconfig

[root@shkf6-244 ~]# cd /opt/kubernetes/server/bin/conf/
[root@shkf6-244 conf]# md5sum kubelet.kubeconfig 
8b2f777f58e4e36d09594db5d1d83ef2  kubelet.kubeconfig
```

由`kubelet.kubeconfig`找到证书信息

```shell
[root@shkf6-244 conf]# cat kubelet.kubeconfig 
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUR0RENDQXB5Z0F3SUJBZ0lVTlp2SjFqakpRL1ZXdk1HQ1RZUmRJNFpDRU40d0RRWUpLb1pJaHZjTkFRRUwKQlFBd1lERUxNQWtHQTFVRUJoTUNRMDR4RURBT0JnTlZCQWdUQjJKbGFXcHBibWN4RURBT0JnTlZCQWNUQjJKbAphV3BwYm1jeEN6QUpCZ05WQkFvVEFtOWtNUXd3Q2dZRFZRUUxFd052Y0hNeEVqQVFCZ05WQkFNVENVOXNaR0p2CmVVVmtkVEFlRncweE9URXhNVGd3TXpFd01EQmFGdzB6T1RFeE1UTXdNekV3TURCYU1HQXhDekFKQmdOVkJBWVQKQWtOT01SQXdEZ1lEVlFRSUV3ZGlaV2xxYVc1bk1SQXdEZ1lEVlFRSEV3ZGlaV2xxYVc1bk1Rc3dDUVlEVlFRSwpFd0p2WkRFTU1Bb0dBMVVFQ3hNRGIzQnpNUkl3RUFZRFZRUURFd2xQYkdSaWIzbEZaSFV3Z2dFaU1BMEdDU3FHClNJYjNEUUVCQVFVQUE0SUJEd0F3Z2dFS0FvSUJBUUNqWUFLaXlJaGIxcDkwUzc4SEY0Y1d5K3BSRWNRZUpVNjEKdFplelFkOVdocjgyY2pMUTlQRmVjMnFqL0Uxb2c3ZmNRZVdpT1pKN2oxczE2RGVHQUZqampVYUVHc1VjQ2VnUAovUmQ5TjRUb0pKT3dJYlJWTlcvWkYvQ21jSGdHdEpjWG8xMDdmVGFYQVdYNXo3SUVlTzNmSUVHZDM5WHFMUFJsClhNdVJHQzBxVklKdmxpNUp3eWhGTS9lNnR0VjdPMFIyZ2lKZUpxZWo0cUFRWXVKaUhVSmtHNW9xM2dvNUVoT04KMW4xS1ZrTDcrT2h5RitCYTVFNzlNWTNzT1E2MWlZM2IwY3BPaE5qWFdLYllNOC9XWUtNUkdDNjhyVEE0WVByWgppL1FZc2RVT0VmSG1HTmh5ai9KakJVMFArZlJzMVIraDBaR0Vac3pJQlRxb0c4dTFNUFJ2QWdNQkFBR2paakJrCk1BNEdBMVVkRHdFQi93UUVBd0lCQmpBU0JnTlZIUk1CQWY4RUNEQUdBUUgvQWdFQ01CMEdBMVVkRGdRV0JCUnkKRS94eUFBZ0VvRnFRc3VJUSttcCtlaTBzSWpBZkJnTlZIU01FR0RBV2dCUnlFL3h5QUFnRW9GcVFzdUlRK21wKwplaTBzSWpBTkJna3Foa2lHOXcwQkFRc0ZBQU9DQVFFQVc1S3JZVUlWRkx1SGc2MTVuYmVMU2FHaWl0SjZ0VHF4CkRHRmdHb21CbElSUzJiNGNRd1htbWtpend5TVkzakdTcjl3QlgzOEFTV1dWSXBKYk9NZXpLVnJyQkRCMlBUODgKVW5XcUlqVmlNOUJRa2k2WVlRdld1eHo0N2h6cnRzbFRiaHBFREZ2aVlueWcvenMra2l6emY4RmNSOEd4MHdCRAoyc2FvZE1od21WWTloNnhzZSthQzRLbW9ieFB1MWdINUtKRGh5MjZKTitkSkxHTVB2MlVLRmRYb1JzaVlsanBaCmh5bGlTOVJ2dm1jODZ4Vk9UVWlOZnFvTzFza1hiZW1HdHg1QU0zcHNoUzN4NmhLdXQzUUpXSkRUM1dYUXpyQjgKQXdBMy9NWW12aE1FWlUzTExtclo5eERGRmZTeFYzN0JtUmV2MGhwN2dSWGRiblRJVW8yait3PT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=
    server: https://192.168.6.66:7443
  name: myk8s
contexts:
- context:
    cluster: myk8s
    user: k8s-node
  name: myk8s-context
current-context: myk8s-context
kind: Config
preferences: {}
users:
- name: k8s-node
  user:
    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUR3akNDQXFxZ0F3SUJBZ0lVSU4wSnJubVlWR1UwNkQyZEE1dFZpUzNYenhvd0RRWUpLb1pJaHZjTkFRRUwKQlFBd1lERUxNQWtHQTFVRUJoTUNRMDR4RURBT0JnTlZCQWdUQjJKbGFXcHBibWN4RURBT0JnTlZCQWNUQjJKbAphV3BwYm1jeEN6QUpCZ05WQkFvVEFtOWtNUXd3Q2dZRFZRUUxFd052Y0hNeEVqQVFCZ05WQkFNVENVOXNaR0p2CmVVVmtkVEFlRncweE9URXhNVGd3TlRVeU1EQmFGdzB6T1RFeE1UTXdOVFV5TURCYU1GOHhDekFKQmdOVkJBWVQKQWtOT01SQXdEZ1lEVlFRSUV3ZGlaV2xxYVc1bk1SQXdEZ1lEVlFRSEV3ZGlaV2xxYVc1bk1Rc3dDUVlEVlFRSwpFd0p2WkRFTU1Bb0dBMVVFQ3hNRGIzQnpNUkV3RHdZRFZRUURFd2hyT0hNdGJtOWtaVENDQVNJd0RRWUpLb1pJCmh2Y05BUUVCQlFBRGdnRVBBRENDQVFvQ2dnRUJBTW5TeHovdWpVUjhJalBJRzJLL1BNeTIxY244ZVo0eVl0eG0KTERGcDdkYjJXcVNVRjFKSEV4VWRyVGZLcGF4TXFESCt3N01mR0F1YXFHNnA2Tm1RS1dUK1QrS1lQeW1KTkRNSAo5L25NTnFMOGFoT1lCNFFFVWZLbnRLRHNRSENDb0RhUG5nbkwySDJLZS8rZGJrTThPUXBRSU9MdkJnRmdjUENnCmR2S1hBOGtUOGY1RXZVUUhMZEdrZXJxTTZ2TFhkdlEweCtSS3RQMFhwajhxRCs3azIxNVNBZXJzQmVOZExXS2MKaUFkUWhBUmg2VUNYb2lzL1k1ZXBOeDYrMWhmQlg5QTYycnNuZCtzcTByQ1puSFh1cVY1eEVMbnBaMmIwemVxQwpnTzVObksyaGJ0Z1BQdkc3d2NwV24wU1ZHb3IweUExMXBBeUpaMzNIY2NKQWQ0Wi9uUnNDQXdFQUFhTjFNSE13CkRnWURWUjBQQVFIL0JBUURBZ1dnTUJNR0ExVWRKUVFNTUFvR0NDc0dBUVVGQndNQ01Bd0dBMVVkRXdFQi93UUMKTUFBd0hRWURWUjBPQkJZRUZPQkQySlVMYWNNekJzUmxWRXBJRUZjYUhoMlpNQjhHQTFVZEl3UVlNQmFBRkhJVAovSElBQ0FTZ1dwQ3k0aEQ2YW41NkxTd2lNQTBHQ1NxR1NJYjNEUUVCQ3dVQUE0SUJBUUNVd0p0Zk4xdzlXT05VCmZlbnFnVW80Ukk1blM4Zzdla1lhUytyQjBnVTFBL0NKZ2lRaWZyMC9zS0xRd29YY1FjZFpTNFhtUHcvY05tNVcKRlUwdnRTODhPL2k0N1RFYzV3NWtZZkNYemR3NHhZeVFIVHhyQkk5RUVRMGxkeUdsY293cUk0RGVFeUZ4d3o3bApsUnNZMmNPS3hZQmpjSENjb29oMUJkaEhHZVI1SXB2Nks1SEJmNWtweURKUGs1NXZwMTRIdzRkTDlMNFE4R2JZCjI1bDhKWE95ampGdGVDZmdUTkFmZnhTYmpCR0hLK2UreGRGU1V1eUc5WG9FSEJNc2l1L2o0QUsvb0tiRXNteUgKMFpYdit0c2FlWkd4Ukpyb3BVWldFOXNPU0NxQ0ZEWWl3dkZmSmhnRENzbFZHRmdFc3FoY1JTVHdWUTNDeVh5bApWS25XTTFEOAotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==
    client-key-data: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcEFJQkFBS0NBUUVBeWRMSFArNk5SSHdpTThnYllyODh6TGJWeWZ4NW5qSmkzR1lzTVdudDF2WmFwSlFYClVrY1RGUjJ0TjhxbHJFeW9NZjdEc3g4WUM1cW9icW5vMlpBcFpQNVA0cGcvS1lrME13ZjMrY3cyb3Z4cUU1Z0gKaEFSUjhxZTBvT3hBY0lLZ05vK2VDY3ZZZllwNy81MXVRenc1Q2xBZzR1OEdBV0J3OEtCMjhwY0R5UlB4L2tTOQpSQWN0MGFSNnVvenE4dGQyOURUSDVFcTAvUmVtUHlvUDd1VGJYbElCNnV3RjQxMHRZcHlJQjFDRUJHSHBRSmVpCkt6OWpsNmszSHI3V0Y4RmYwRHJhdXlkMzZ5clNzSm1jZGU2cFhuRVF1ZWxuWnZUTjZvS0E3azJjcmFGdTJBOCsKOGJ2QnlsYWZSSlVhaXZUSURYV2tESWxuZmNkeHdrQjNobitkR3dJREFRQUJBb0lCQUYyN1U2aGdmU0Z5V1Z3ZApNb0xRK0ViSEgxRTR2YTc0RGF2NGs4dTdPNmViTUl2QTczZlo1SVhwQzNxZTFnVEljVkVPMWdySmhSeFdqcVVlCnFqTG8zaUMyYjVsNFJkVmZrR3VtNXNjUHpjd3lXSDJUSE9KMk15ejBNRktRaG5qNlliZ1ZTVHVaZllrSW1RQWwKT0lGblpjSmhablNldC9aSnVRbzRMQ1lNZHNpYWM4Vis2dk1CdWtiL2pwL2JXT1F4aFM4MmtPREdaa3BaTDVFVgorR3NyeGFTaDB6aHpkaGlnSk5TRWR0S1lsR0xOUVdwU3ZscXNoMDhtNlRQWld4UkdzaDB4TG51ckkzeWJnZkJxCittWXRPUEh5dUZqeWlCZno4OHJEWUtYKytweTJwUzB5VGVrbHdtSW9NVk9SaGdDVG9sNkt3RENZeGQxVXJ4UE4KSWUyL3Joa0NnWUVBejRKU2EweU5sZGtudEhQSFVTQkNXOWIxYlF0NFRlV1NIWEppUU1ZWXVUUnlUZWZkeTEyTwp0RTE3c1ljWlU0Mkh4UitzTTBFenFka3kydWl5ZzdCME1ibmY3UXJYU0Q4YlhLNVZBbXVKd1Jxc1pJMXhldG9PCnJhcGttc09GWXVVT1BMU2d2UEVOVGJTckQ2d1U3eDFFUWU1L2xrWlgyeWg1UmpuTmJmdENtK1VDZ1lFQStQeFMKemlsVUh2M2xubGFpaWMzSTZkSVpWQ3lmK3JRdGJvb0FwVEF1TklNM3pzWFlKcDg0MVJNUXJZZ0FLRkhEY0VJbwphaFU2V011U3JTclFvNndDNXB6M3dsL1JCWlAvUWRRclF1L1lLd2RCVi9XQk9kM3Z6eGpObjIvN0plbE1ueUt2Ci9sZ29IcTZTN2hPUGdaUmg0aHZLRFBqazdoLzU1ZzQ2NDVIdnhQOENnWUVBcHd6Yit1TkM3QXBJYTMzMVREcnoKRU9vbzQ2TWpNMXFIMlVyWERCd3RsUk5DbmJMMm01dnlvUFhyaVF3Z2VHSHNsZVdjaEJxT1U4SzFyUU05aXNSSApsaXh6dDJsTnpDeDVnNUFZZ1gwL0JZVEttWnhBYWMwWG1ma2RTblh5Y0ozRGExMWlOUmk5US93WTVlSDdiRStjClBwT1loTXFXT2FrSWtGOUNJTEx3ZVgwQ2dZRUFoeG9iTUdTNmtZcUJVdFo5b2JxNDN5OHlzVHI1bjhhdXRFRkwKc2xhZmE3MGJ4aVlTY0hxTEV3c2lUSmIwUnV4KzJPWDlHZnJreXhQRFJoVnFXclZXYVo0WXppN0JzMzRuenFkNgp4ZnB3MklBNlU2a1NjcnpiaUF0VVg4UWFpZXE2dWNyUHBucGRZckNsWjJ2VHZhTXZMY3FZYTB1T3BTdFNwU05wCmp0dzhOeThDZ1lBdk9VKzJpYnphYXNNL3BjUHp3TStCWERub0NrcTdrdk4wRjh5dDJFdmJkUFlIMWc4UGVwa0cKWDYxTXIxVVQ2bVQ1ZC9HeTcrOXhOd2FGZzZENFk5VW5DUTBlU3lWL1plUWpGSGZtQS8rUUUxUy82K0pib1J4MwpQMUVsZ2psN0RXU3RodkJsYmhWYjdYc2MyTGhSMUR2RFJmUURqWE1MRWdvNC9LUXJULzRqd3c9PQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo=

[root@shkf6-245 certs]# echo "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUR3akNDQXFxZ0F3SUJBZ0lVSU4wSnJubVlWR1UwNkQyZEE1dFZpUzNYenhvd0RRWUpLb1pJaHZjTkFRRUwKQlFBd1lERUxNQWtHQTFVRUJoTUNRMDR4RURBT0JnTlZCQWdUQjJKbGFXcHBibWN4RURBT0JnTlZCQWNUQjJKbAphV3BwYm1jeEN6QUpCZ05WQkFvVEFtOWtNUXd3Q2dZRFZRUUxFd052Y0hNeEVqQVFCZ05WQkFNVENVOXNaR0p2CmVVVmtkVEFlRncweE9URXhNVGd3TlRVeU1EQmFGdzB6T1RFeE1UTXdOVFV5TURCYU1GOHhDekFKQmdOVkJBWVQKQWtOT01SQXdEZ1lEVlFRSUV3ZGlaV2xxYVc1bk1SQXdEZ1lEVlFRSEV3ZGlaV2xxYVc1bk1Rc3dDUVlEVlFRSwpFd0p2WkRFTU1Bb0dBMVVFQ3hNRGIzQnpNUkV3RHdZRFZRUURFd2hyT0hNdGJtOWtaVENDQVNJd0RRWUpLb1pJCmh2Y05BUUVCQlFBRGdnRVBBRENDQVFvQ2dnRUJBTW5TeHovdWpVUjhJalBJRzJLL1BNeTIxY244ZVo0eVl0eG0KTERGcDdkYjJXcVNVRjFKSEV4VWRyVGZLcGF4TXFESCt3N01mR0F1YXFHNnA2Tm1RS1dUK1QrS1lQeW1KTkRNSAo5L25NTnFMOGFoT1lCNFFFVWZLbnRLRHNRSENDb0RhUG5nbkwySDJLZS8rZGJrTThPUXBRSU9MdkJnRmdjUENnCmR2S1hBOGtUOGY1RXZVUUhMZEdrZXJxTTZ2TFhkdlEweCtSS3RQMFhwajhxRCs3azIxNVNBZXJzQmVOZExXS2MKaUFkUWhBUmg2VUNYb2lzL1k1ZXBOeDYrMWhmQlg5QTYycnNuZCtzcTByQ1puSFh1cVY1eEVMbnBaMmIwemVxQwpnTzVObksyaGJ0Z1BQdkc3d2NwV24wU1ZHb3IweUExMXBBeUpaMzNIY2NKQWQ0Wi9uUnNDQXdFQUFhTjFNSE13CkRnWURWUjBQQVFIL0JBUURBZ1dnTUJNR0ExVWRKUVFNTUFvR0NDc0dBUVVGQndNQ01Bd0dBMVVkRXdFQi93UUMKTUFBd0hRWURWUjBPQkJZRUZPQkQySlVMYWNNekJzUmxWRXBJRUZjYUhoMlpNQjhHQTFVZEl3UVlNQmFBRkhJVAovSElBQ0FTZ1dwQ3k0aEQ2YW41NkxTd2lNQTBHQ1NxR1NJYjNEUUVCQ3dVQUE0SUJBUUNVd0p0Zk4xdzlXT05VCmZlbnFnVW80Ukk1blM4Zzdla1lhUytyQjBnVTFBL0NKZ2lRaWZyMC9zS0xRd29YY1FjZFpTNFhtUHcvY05tNVcKRlUwdnRTODhPL2k0N1RFYzV3NWtZZkNYemR3NHhZeVFIVHhyQkk5RUVRMGxkeUdsY293cUk0RGVFeUZ4d3o3bApsUnNZMmNPS3hZQmpjSENjb29oMUJkaEhHZVI1SXB2Nks1SEJmNWtweURKUGs1NXZwMTRIdzRkTDlMNFE4R2JZCjI1bDhKWE95ampGdGVDZmdUTkFmZnhTYmpCR0hLK2UreGRGU1V1eUc5WG9FSEJNc2l1L2o0QUsvb0tiRXNteUgKMFpYdit0c2FlWkd4Ukpyb3BVWldFOXNPU0NxQ0ZEWWl3dkZmSmhnRENzbFZHRmdFc3FoY1JTVHdWUTNDeVh5bApWS25XTTFEOAotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg=="|base64 -d  > 123.pem
[root@shkf6-245 certs]# cfssl-certinfo -cert 123.pem 
{
  "subject": {
    "common_name": "k8s-node",
    "country": "CN",
    "organization": "od",
    "organizational_unit": "ops",
    "locality": "beijing",
    "province": "beijing",
    "names": [
      "CN",
      "beijing",
      "beijing",
      "od",
      "ops",
      "k8s-node"
    ]
  },
  "issuer": {
    "common_name": "OldboyEdu",
    "country": "CN",
    "organization": "od",
    "organizational_unit": "ops",
    "locality": "beijing",
    "province": "beijing",
    "names": [
      "CN",
      "beijing",
      "beijing",
      "od",
      "ops",
      "OldboyEdu"
    ]
  },
  "serial_number": "187617012736570890928549677433593587138601406234",
  "not_before": "2019-11-18T05:52:00Z",
  "not_after": "2039-11-13T05:52:00Z",
  "sigalg": "SHA256WithRSA",
  "authority_key_id": "72:13:FC:72:0:8:4:A0:5A:90:B2:E2:10:FA:6A:7E:7A:2D:2C:22",
  "subject_key_id": "E0:43:D8:95:B:69:C3:33:6:C4:65:54:4A:48:10:57:1A:1E:1D:99",
  "pem": "-----BEGIN CERTIFICATE-----\nMIIDwjCCAqqgAwIBAgIUIN0JrnmYVGU06D2dA5tViS3XzxowDQYJKoZIhvcNAQEL\nBQAwYDELMAkGA1UEBhMCQ04xEDAOBgNVBAgTB2JlaWppbmcxEDAOBgNVBAcTB2Jl\naWppbmcxCzAJBgNVBAoTAm9kMQwwCgYDVQQLEwNvcHMxEjAQBgNVBAMTCU9sZGJv\neUVkdTAeFw0xOTExMTgwNTUyMDBaFw0zOTExMTMwNTUyMDBaMF8xCzAJBgNVBAYT\nAkNOMRAwDgYDVQQIEwdiZWlqaW5nMRAwDgYDVQQHEwdiZWlqaW5nMQswCQYDVQQK\nEwJvZDEMMAoGA1UECxMDb3BzMREwDwYDVQQDEwhrOHMtbm9kZTCCASIwDQYJKoZI\nhvcNAQEBBQADggEPADCCAQoCggEBAMnSxz/ujUR8IjPIG2K/PMy21cn8eZ4yYtxm\nLDFp7db2WqSUF1JHExUdrTfKpaxMqDH+w7MfGAuaqG6p6NmQKWT+T+KYPymJNDMH\n9/nMNqL8ahOYB4QEUfKntKDsQHCCoDaPngnL2H2Ke/+dbkM8OQpQIOLvBgFgcPCg\ndvKXA8kT8f5EvUQHLdGkerqM6vLXdvQ0x+RKtP0Xpj8qD+7k215SAersBeNdLWKc\niAdQhARh6UCXois/Y5epNx6+1hfBX9A62rsnd+sq0rCZnHXuqV5xELnpZ2b0zeqC\ngO5NnK2hbtgPPvG7wcpWn0SVGor0yA11pAyJZ33HccJAd4Z/nRsCAwEAAaN1MHMw\nDgYDVR0PAQH/BAQDAgWgMBMGA1UdJQQMMAoGCCsGAQUFBwMCMAwGA1UdEwEB/wQC\nMAAwHQYDVR0OBBYEFOBD2JULacMzBsRlVEpIEFcaHh2ZMB8GA1UdIwQYMBaAFHIT\n/HIACASgWpCy4hD6an56LSwiMA0GCSqGSIb3DQEBCwUAA4IBAQCUwJtfN1w9WONU\nfenqgUo4RI5nS8g7ekYaS+rB0gU1A/CJgiQifr0/sKLQwoXcQcdZS4XmPw/cNm5W\nFU0vtS88O/i47TEc5w5kYfCXzdw4xYyQHTxrBI9EEQ0ldyGlcowqI4DeEyFxwz7l\nlRsY2cOKxYBjcHCcooh1BdhHGeR5Ipv6K5HBf5kpyDJPk55vp14Hw4dL9L4Q8GbY\n25l8JXOyjjFteCfgTNAffxSbjBGHK+e+xdFSUuyG9XoEHBMsiu/j4AK/oKbEsmyH\n0ZXv+tsaeZGxRJropUZWE9sOSCqCFDYiwvFfJhgDCslVGFgEsqhcRSTwVQ3CyXyl\nVKnWM1D8\n-----END CERTIFICATE-----\n"
}
```

# 第一章：kubectl命令工具使用详解

管理K8S核心资源的三种基本方法：

- 陈述式管理方法–主要依赖命令行CLI工具进行管理
- 声明式管理方法–主要依赖统一资源配置清单（manifest）进行管理
- GUI式管理方法–主要依赖图形化操作界面（web页面）进行管理

## 1.陈述式资源管理方法

### 1.管理名称空间

- 查看名称空间

```shell
[root@shkf6-243 ~]# kubectl get namespaces
NAME              STATUS   AGE
default           Active   25h
kube-node-lease   Active   25h
kube-public       Active   25h
kube-system       Active   25h

[root@shkf6-243 ~]# kubectl get ns
NAME              STATUS   AGE
default           Active   25h
kube-node-lease   Active   25h
kube-public       Active   25h
kube-system       Active   25h
```

- 查看名称空间内的资源

```shell
[root@shkf6-243 ~]# kubectl get all -n default
NAME                 READY   STATUS    RESTARTS   AGE
pod/nginx-ds-2692m   1/1     Running   0          104m
pod/nginx-ds-gf6hs   1/1     Running   0          104m


NAME                 TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
service/kubernetes   ClusterIP   10.96.0.1    <none>        443/TCP   25h

NAME                      DESIRED   CURRENT   READY   UP-TO-DATE   AVAILABLE   NODE SELECTOR   AGE
daemonset.apps/nginx-ds   2         2         2       2            2           <none>          104m
```

- 创建名称空间

```shell
[root@shkf6-243 ~]# kubectl create namespace app
namespace/app created
```

- 删除名称空间

```shell
[root@shkf6-243 ~]# kubectl delete namespace app
namespace "app" deleted
```

### 2.管理deployment资源

- 创建deployment

```shell
[root@shkf6-243 ~]# kubectl create deployment nginx-dp --image=harbor.od.com/public/nginx:v1.7.9 -n kube-public
deployment.apps/nginx-dp created
```

- 查看deployment

  - 简单查看

    ```shell
    [root@shkf6-243 ~]# kubectl get deploy -n kube-public
    NAME       READY   UP-TO-DATE   AVAILABLE   AGE
    nginx-dp   1/1     1            1           5m10s
    ```

  - 扩展查看

    ```shell
    [root@shkf6-243 ~]# kubectl get deployment -o wide -n kube-public
    NAME       READY   UP-TO-DATE   AVAILABLE   AGE     CONTAINERS   IMAGES                              SELECTOR
    nginx-dp   1/1     1            1           5m21s   nginx        harbor.od.com/public/nginx:v1.7.9   app=nginx-dp
    ```

  - 详细查看

    ```shell
    [root@shkf6-243 ~]# kubectl describe deployment nginx-dp -n kube-public
    Name:                   nginx-dp
    Namespace:              kube-public
    CreationTimestamp:      Tue, 19 Nov 2019 15:58:10 +0800
    Labels:                 app=nginx-dp
    ...
    ```

- 查看pod资源

```shell
[root@shkf6-243 ~]# kubectl get pods -n kube-public
NAME                        READY   STATUS    RESTARTS   AGE
nginx-dp-5dfc689474-lz5lz   1/1     Running   0          28m
```

- 进入pod资源

```shell
[root@shkf6-243 ~]# kubectl exec -it nginx-dp-5dfc689474-lz5lz bash -n kube-public
root@nginx-dp-5dfc689474-lz5lz:/# ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN qlen 1
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
14: eth0@if15: <BROADCAST,MULTICAST,UP,LOWER_UP,M-DOWN> mtu 1500 qdisc noqueue state UP 
    link/ether 02:42:ac:06:f3:03 brd ff:ff:ff:ff:ff:ff
    inet 172.6.243.3/24 brd 172.6.243.255 scope global eth0
       valid_lft forever preferred_lft forever
```

> 当然你也可以docker exec进入容器

- 删除pod资源（重启）

```shell
[root@shkf6-243 ~]# kubectl delete pod nginx-dp-5dfc689474-lz5lz -n kube-public
pod "nginx-dp-5dfc689474-lz5lz" deleted
[root@shkf6-243 ~]# kubectl get pod  -n kube-public
NAME                        READY   STATUS    RESTARTS   AGE
nginx-dp-5dfc689474-vtrwj   1/1     Running   0          21s
```

> 使用watch观察pod重建状态变化
> 强制删除参数： `--force --grace-period=0`

```
[root@shkf6-243 ~]# watch -n 1 'kubectl describe deployment nginx-dp -n kube-public|grep -C 5 Event'
```

- 删除deployment

```shell
[root@shkf6-243 ~]# kubectl delete deployment nginx-dp -n kube-public
deployment.extensions "nginx-dp" deleted
[root@shkf6-243 ~]# kubectl get deployment -n kube-public
No resources found.
[root@shkf6-243 ~]# kubectl get pods -n kube-public
No resources found.
```

### 3.管理service资源

```shell
[root@shkf6-243 ~]# kubectl create deployment nginx-dp --image=harbor.od.com/public/nginx:v1.7.9 -n kube-public
deployment.apps/nginx-dp created

[root@shkf6-243 ~]# kubectl get all -n kube-public
NAME                            READY   STATUS    RESTARTS   AGE
pod/nginx-dp-5dfc689474-z5rfh   1/1     Running   0          44m


NAME                       READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/nginx-dp   1/1     1            1           44m

NAME                                  DESIRED   CURRENT   READY   AGE
replicaset.apps/nginx-dp-5dfc689474   1         1         1       44m
```

- 创建service

```shell
[root@shkf6-243 ~]# kubectl expose deployment nginx-dp --port=80 -n kube-public
service/nginx-dp exposed
```

查看一下代码查看规律，增加一个副本，两个副本，三个副本，ipvs变化：

```shell
[root@shkf6-243 ~]# kubectl get service -n kube-public
NAME       TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)   AGE
nginx-dp   ClusterIP   10.102.187.18   <none>        80/TCP    54s
[root@shkf6-243 ~]# kubectl get svc -n kube-public
NAME       TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)   AGE
nginx-dp   ClusterIP   10.102.187.18   <none>        80/TCP    58s
[root@shkf6-243 ~]# ipvsadm -Ln
IP Virtual Server version 1.2.1 (size=4096)
Prot LocalAddress:Port Scheduler Flags
  -> RemoteAddress:Port           Forward Weight ActiveConn InActConn
TCP  10.96.0.1:443 nq
  -> 192.168.6.243:6443           Masq    1      0          0         
  -> 192.168.6.244:6443           Masq    1      0          0         
TCP  10.102.187.18:80 nq
  -> 172.6.243.3:80               Masq    1      0          0 
[root@shkf6-243 ~]# kubectl scale deployment nginx-dp --replicas=2 -n kube-public
deployment.extensions/nginx-dp scaled
[root@shkf6-243 ~]# ipvsadm -Ln
IP Virtual Server version 1.2.1 (size=4096)
Prot LocalAddress:Port Scheduler Flags
  -> RemoteAddress:Port           Forward Weight ActiveConn InActConn
TCP  10.96.0.1:443 nq
  -> 192.168.6.243:6443           Masq    1      0          0         
  -> 192.168.6.244:6443           Masq    1      0          0         
TCP  10.102.187.18:80 nq
  -> 172.6.243.3:80               Masq    1      0          0         
  -> 172.6.244.3:80               Masq    1      0          0  
[root@shkf6-243 ~]# kubectl scale deployment nginx-dp --replicas=3 -n kube-public
deployment.extensions/nginx-dp scaled

[root@shkf6-243 ~]# ipvsadm -Ln
IP Virtual Server version 1.2.1 (size=4096)
Prot LocalAddress:Port Scheduler Flags
  -> RemoteAddress:Port           Forward Weight ActiveConn InActConn
TCP  10.96.0.1:443 nq
  -> 192.168.6.243:6443           Masq    1      0          0         
  -> 192.168.6.244:6443           Masq    1      0          0         
TCP  10.102.187.18:80 nq
  -> 172.6.243.3:80               Masq    1      0          0         
  -> 172.6.243.4:80               Masq    1      0          0         
  -> 172.6.244.3:80               Masq    1      0          0  
```

说明：这里删除了svc这里的ipvs策略还在，需要用`ipvsadm -D -t`清理

```shell
[root@shkf6-243 ~]# ipvsadm -Ln
IP Virtual Server version 1.2.1 (size=4096)
Prot LocalAddress:Port Scheduler Flags
  -> RemoteAddress:Port           Forward Weight ActiveConn InActConn
TCP  10.96.0.1:443 nq
  -> 192.168.6.243:6443           Masq    1      0          0         
  -> 192.168.6.244:6443           Masq    1      0          0                 
TCP  10.96.1.83:80 nq
  -> 172.6.243.3:80               Masq    1      0          0         
  -> 172.6.243.4:80               Masq    1      0          0         
  -> 172.6.244.3:80               Masq    1      0          0         
TCP  10.102.187.18:80 nq
  -> 172.6.243.3:80               Masq    1      0          0         
  -> 172.6.243.4:80               Masq    1      0          0         
  -> 172.6.244.3:80               Masq    1      0          0         

[root@shkf6-243 ~]# ipvsadm -D -t  10.102.187.18:80
[root@shkf6-243 ~]# ipvsadm -Ln
IP Virtual Server version 1.2.1 (size=4096)
Prot LocalAddress:Port Scheduler Flags
  -> RemoteAddress:Port           Forward Weight ActiveConn InActConn
TCP  10.96.0.1:443 nq
  -> 192.168.6.243:6443           Masq    1      0          0         
  -> 192.168.6.244:6443           Masq    1      0          0         
TCP  10.96.1.83:80 nq
  -> 172.6.243.3:80               Masq    1      0          0         
  -> 172.6.243.4:80               Masq    1      0          0         
  -> 172.6.244.3:80               Masq    1      0          0 
```

陈述式有局限性，不能给daemonset创建service，有局限性

```shell
[root@shkf6-243 ~]# kubectl expose daemonset nginx-ds --port=880
error: cannot expose a DaemonSet.extensions
```

- 查看service

```shell
[root@shkf6-243 ~]# kubectl describe svc nginx-dp -n kube-public
Name:              nginx-dp
Namespace:         kube-public
Labels:            app=nginx-dp
Annotations:       <none>
Selector:          app=nginx-dp
Type:              ClusterIP
IP:                10.102.187.18
Port:              <unset>  80/TCP
TargetPort:        80/TCP
Endpoints:         172.6.243.3:80,172.6.243.4:80,172.6.244.3:80
Session Affinity:  None
Events:            <none>
```

### 4.kubectl用法总结

陈述式资源管理方法小结：

- kubernetes集群管理集群资源的唯一入口是通过相应的方法调用apiserver的接口
- kubulet是官方的CLI命令行工具，用于与apiserver进行通信，将用户在命令行输入的命令，组织并转化为apiserver能识别的信息，进而实现管理k8s各种资源的一种有效途径
- kubectl的命令大全
  - kubectl –help
  - http://docs.kubernetes.org.cn/
- 陈述式资源管理方法可以满足90%以上的资源需求，但它的缺点也很明显
  - 命令冗长、复杂、难以记忆
  - 特定场景下，无法实现管理需求
  - 对资源的增、删、查操作比较容易，改就很痛苦了

声明式资源管理小结：

- 声明式资源管理方法，依赖于统一资源配置清单文件对资源进行管理
- 对资源的管理，是通过事先定义在统一资源配置清单内，在通过陈述式命令应用到K8S集群里
- 语法格式：kubectl create/apply/delete -f /path/to/yaml
- 资源配置清单的学习方法：
  - tip1：多看别人（官方）写的，能读懂
  - tip2：能照着现成的文件改着用
  - tip3：遇到不懂的，善用kubectl explain..查
  - tip4：初学切记上来就无中生有，自己憋着写

## 2.声明式资源管理方法

声明式资源管理方法

- 声明式资源管理方法依赖于—资源配置清单（yaml/json）
- 查看资源配置清单的方法
  - ~]# kubectl get svc nginx-dp -o yaml -n kube-public
- 解释资源配置清单
  - ~]# kubectl explain service
- 创建资源配置清单
  - ~]# vi /root/nginx-ds-svc.yaml
- 应用资源配置清单
  - ~]# kubectl apply -f nginx-ds-svc.yaml
- 修改资源配置清单并应用
  - 在线修改
  - 离线修改
- 删除资源配置清单
  - 陈述式
  - 声明式

### 1.陈述式资源管理方法的局限性

- 命令冗长、复杂、难以记忆
- 特定场景下，无法实现管理需求
- 对资源的增、删、查操作比较容易，改就很痛苦了

### 2.查看资源配置清单

```shell
[root@shkf6-243 ~]# kubectl get svc nginx-dp -o yaml -n kube-public
apiVersion: v1
kind: Service
metadata:
  creationTimestamp: "2019-11-20T00:38:32Z"
  labels:
    app: nginx-dp
  name: nginx-dp
  namespace: kube-public
  resourceVersion: "218290"
  selfLink: /api/v1/namespaces/kube-public/services/nginx-dp
  uid: d48e3085-cea5-41ad-afe2-259c80524d9d
spec:
  clusterIP: 10.96.1.70
  ports:
  - port: 80
    protocol: TCP
    targetPort: 80
  selector:
    app: nginx-dp
  sessionAffinity: None
  type: ClusterIP
status:
  loadBalancer: {}
```

### 3.解释资源配置清单

```shell
[root@shkf6-243 ~]# kubectl explain service.metadata
KIND:     Service
VERSION:  v1

RESOURCE: metadata <Object>

DESCRIPTION:
     Standard object's metadata. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata

     ObjectMeta is metadata that all persisted resources must have, which
     includes all objects users must create.

FIELDS:
   annotations    <map[string]string>
     Annotations is an unstructured key value map stored with a resource that
     may be set by external tools to store and retrieve arbitrary metadata. They
     are not queryable and should be preserved when modifying objects. More
     info: http://kubernetes.io/docs/user-guide/annotations

   clusterName    <string>
     The name of the cluster which the object belongs to. This is used to
     distinguish resources with same name and namespace in different clusters.
     This field is not set anywhere right now and apiserver is going to ignore
     it if set in create or update request.

   creationTimestamp    <string>
     CreationTimestamp is a timestamp representing the server time when this
     object was created. It is not guaranteed to be set in happens-before order
     across separate operations. Clients may not set this value. It is
     represented in RFC3339 form and is in UTC. Populated by the system.
     Read-only. Null for lists. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata

   deletionGracePeriodSeconds    <integer>
     Number of seconds allowed for this object to gracefully terminate before it
     will be removed from the system. Only set when deletionTimestamp is also
     set. May only be shortened. Read-only.

   deletionTimestamp    <string>
     DeletionTimestamp is RFC 3339 date and time at which this resource will be
     deleted. This field is set by the server when a graceful deletion is
     requested by the user, and is not directly settable by a client. The
     resource is expected to be deleted (no longer visible from resource lists,
     and not reachable by name) after the time in this field, once the
     finalizers list is empty. As long as the finalizers list contains items,
     deletion is blocked. Once the deletionTimestamp is set, this value may not
     be unset or be set further into the future, although it may be shortened or
     the resource may be deleted prior to this time. For example, a user may
     request that a pod is deleted in 30 seconds. The Kubelet will react by
     sending a graceful termination signal to the containers in the pod. After
     that 30 seconds, the Kubelet will send a hard termination signal (SIGKILL)
     to the container and after cleanup, remove the pod from the API. In the
     presence of network partitions, this object may still exist after this
     timestamp, until an administrator or automated process can determine the
     resource is fully terminated. If not set, graceful deletion of the object
     has not been requested. Populated by the system when a graceful deletion is
     requested. Read-only. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata

   finalizers    <[]string>
     Must be empty before the object is deleted from the registry. Each entry is
     an identifier for the responsible component that will remove the entry from
     the list. If the deletionTimestamp of the object is non-nil, entries in
     this list can only be removed.

   generateName    <string>
     GenerateName is an optional prefix, used by the server, to generate a
     unique name ONLY IF the Name field has not been provided. If this field is
     used, the name returned to the client will be different than the name
     passed. This value will also be combined with a unique suffix. The provided
     value has the same validation rules as the Name field, and may be truncated
     by the length of the suffix required to make the value unique on the
     server. If this field is specified and the generated name exists, the
     server will NOT return a 409 - instead, it will either return 201 Created
     or 500 with Reason ServerTimeout indicating a unique name could not be
     found in the time allotted, and the client should retry (optionally after
     the time indicated in the Retry-After header). Applied only if Name is not
     specified. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#idempotency

   generation    <integer>
     A sequence number representing a specific generation of the desired state.
     Populated by the system. Read-only.

   initializers    <Object>
     An initializer is a controller which enforces some system invariant at
     object creation time. This field is a list of initializers that have not
     yet acted on this object. If nil or empty, this object has been completely
     initialized. Otherwise, the object is considered uninitialized and is
     hidden (in list/watch and get calls) from clients that haven't explicitly
     asked to observe uninitialized objects. When an object is created, the
     system will populate this list with the current set of initializers. Only
     privileged users may set or modify this list. Once it is empty, it may not
     be modified further by any user. DEPRECATED - initializers are an alpha
     field and will be removed in v1.15.

   labels    <map[string]string>
     Map of string keys and values that can be used to organize and categorize
     (scope and select) objects. May match selectors of replication controllers
     and services. More info: http://kubernetes.io/docs/user-guide/labels

   managedFields    <[]Object>
     ManagedFields maps workflow-id and version to the set of fields that are
     managed by that workflow. This is mostly for internal housekeeping, and
     users typically shouldn't need to set or understand this field. A workflow
     can be the user's name, a controller's name, or the name of a specific
     apply path like "ci-cd". The set of fields is always in the version that
     the workflow used when modifying the object. This field is alpha and can be
     changed or removed without notice.

   name    <string>
     Name must be unique within a namespace. Is required when creating
     resources, although some resources may allow a client to request the
     generation of an appropriate name automatically. Name is primarily intended
     for creation idempotence and configuration definition. Cannot be updated.
     More info: http://kubernetes.io/docs/user-guide/identifiers#names

   namespace    <string>
     Namespace defines the space within each name must be unique. An empty
     namespace is equivalent to the "default" namespace, but "default" is the
     canonical representation. Not all objects are required to be scoped to a
     namespace - the value of this field for those objects will be empty. Must
     be a DNS_LABEL. Cannot be updated. More info:
     http://kubernetes.io/docs/user-guide/namespaces

   ownerReferences    <[]Object>
     List of objects depended by this object. If ALL objects in the list have
     been deleted, this object will be garbage collected. If this object is
     managed by a controller, then an entry in this list will point to this
     controller, with the controller field set to true. There cannot be more
     than one managing controller.

   resourceVersion    <string>
     An opaque value that represents the internal version of this object that
     can be used by clients to determine when objects have changed. May be used
     for optimistic concurrency, change detection, and the watch operation on a
     resource or set of resources. Clients must treat these values as opaque and
     passed unmodified back to the server. They may only be valid for a
     particular resource or set of resources. Populated by the system.
     Read-only. Value must be treated as opaque by clients and . More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency

   selfLink    <string>
     SelfLink is a URL representing this object. Populated by the system.
     Read-only.

   uid    <string>
     UID is the unique in time and space value for this object. It is typically
     generated by the server on successful creation of a resource and is not
     allowed to change on PUT operations. Populated by the system. Read-only.
     More info: http://kubernetes.io/docs/user-guide/identifiers#uids

[root@shkf6-243 ~]# kubectl explain service
KIND:     Service
VERSION:  v1

DESCRIPTION:
     Service is a named abstraction of software service (for example, mysql)
     consisting of local port (for example 3306) that the proxy listens on, and
     the selector that determines which pods will answer requests sent through
     the proxy.

FIELDS:
   apiVersion    <string>
     APIVersion defines the versioned schema of this representation of an
     object. Servers should convert recognized schemas to the latest internal
     value, and may reject unrecognized values. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#resources

   kind    <string>
     Kind is a string value representing the REST resource this object
     represents. Servers may infer this from the endpoint the client submits
     requests to. Cannot be updated. In CamelCase. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds

   metadata    <Object>
     Standard object's metadata. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata

   spec    <Object>
     Spec defines the behavior of a service.
     https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status

   status    <Object>
     Most recently observed status of the service. Populated by the system.
     Read-only. More info:
     https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status
```

### 4.创建资源配置清单

```shell
[root@shkf6-243 ~]# vi /root/nginx-ds-svc.yaml 
[root@shkf6-243 ~]# cat /root/nginx-ds-svc.yaml
apiVersion: v1
kind: Service
metadata:
  labels:
    app: nginx-ds
  name: nginx-ds
  namespace: default
spec:
  ports:
  - port: 80
    protocol: TCP
    targetPort: 80
  selector:
    app: nginx-ds
  sessionAffinity: None
  type: ClusterIP
```

### 5.应用资源配置清单

```shell
[root@shkf6-243 ~]# kubectl apply -f /root/nginx-ds-svc.yaml 
service/nginx-ds created
```

### 6.修改资源配置清单并应用

- 离线修改
  修改nginx-ds-svc.yaml文件，并用`kubelet apply -f nginx-ds-svc.yaml`文件使之生效。
- 在线修改
  直接使用kubectl edit service nginx-ds在线编辑资源配置清单并保存生效

### 7.删除资源配置清单

- 陈述式删除

```
kubectl delete service nginx-ds -n kube-public
```

- 声明式删除

```
kubectl delete -f nginx-ds-svc.yaml
```

### 8.查看并使用Service资源

```shell
[root@shkf6-244 ~]# kubectl get svc -o wide
NAME         TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE   SELECTOR
kubernetes   ClusterIP   10.96.0.1    <none>        443/TCP   45h   <none>
nginx-ds     ClusterIP   10.96.0.42   <none>        80/TCP    99m   app=nginx-ds
[root@shkf6-244 ~]# curl 10.96.0.42
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
[root@shkf6-244 ~]# kubectl describe svc nginx-ds
Name:              nginx-ds
Namespace:         default
Labels:            app=nginx-ds
Annotations:       kubectl.kubernetes.io/last-applied-configuration:
                     {"apiVersion":"v1","kind":"Service","metadata":{"annotations":{},"labels":{"app":"nginx-ds"},"name":"nginx-ds","namespace":"default"},"spe...
Selector:          app=nginx-ds
Type:              ClusterIP
IP:                10.96.0.42
Port:              <unset>  80/TCP
TargetPort:        80/TCP
Endpoints:         172.6.243.2:80,172.6.244.2:80
Session Affinity:  None
Events:            <none>
```

# 第二章：kubenetes的CNI网络插件-flannel

kubernetes设计了网络模型，但却将它的实现交给了网络插件，CNI网络插件最主要的功能就是实现POD资源能够跨宿主机进行通信

常见的网络插件：

- flannel
- calico
- contiv
- opencontrail
- NSX-T
- Kube-router

### 1.集群规划：

| 主机名             | 角色    | ip            |
| :----------------- | :------ | :------------ |
| shkf6-243.host.com | flannel | 192.168.6.243 |
| shkf6-244.host.com | flannel | 192.168.6.244 |

注意：这里部署文档以`shkf6-243.host.com`主机为例，另外一台运算节点安装部署方法类似

### 2.下载软件，解压，做软连接

[fiannel官方下载地址](https://github.com/coreos/flannel/releases)

shkf6-243.host.com上：

```shell
[root@shkf6-243 ~]# cd /opt/src/
[root@shkf6-243 src]# wget https://github.com/coreos/flannel/releases/download/v0.11.0/flannel-v0.11.0-linux-amd64.tar.gz

[root@shkf6-243 src]# mkdir /opt/flannel-v0.11.0
[root@shkf6-243 src]# tar xf flannel-v0.11.0-linux-amd64.tar.gz -C /opt/flannel-v0.11.0/
[root@shkf6-243 src]# ln -s /opt/flannel-v0.11.0/ /opt/flannel
```

### 3.最终目录结构

```shell
[root@shkf6-243 flannel]# tree /opt -L 2
/opt
├── containerd
│   ├── bin
│   └── lib
├── etcd -> /opt/etcd-v3.1.20/
├── etcd-v3.1.20
│   ├── certs
│   ├── Documentation
│   ├── etcd
│   ├── etcdctl
│   ├── etcd-server-startup.sh
│   ├── README-etcdctl.md
│   ├── README.md
│   └── READMEv2-etcdctl.md
├── flannel -> /opt/flannel-v0.11.0/
├── flannel-v0.11.0
│   ├── cert
│   ├── flanneld
│   ├── mk-docker-opts.sh
│   └── README.md
├── kubernetes -> /opt/kubernetes-v1.15.2/
├── kubernetes-v1.15.2
│   ├── addons
│   ├── LICENSES
│   └── server
├── rh
└── src
    ├── etcd-v3.1.20-linux-amd64.tar.gz
    ├── flannel-v0.11.0-linux-amd64.tar.gz
    └── kubernetes-server-linux-amd64-v1.15.2.tar.gz
```

### 4.拷贝证书

```shell
[root@shkf6-243 src]# mkdir /opt/flannel/cert/
[root@shkf6-243 src]# cd /opt/flannel/cert/
[root@shkf6-243 flannel]# scp -P52113 shkf6-245:/opt/certs/ca.pem /opt/flannel/cert/

[root@shkf6-243 flannel]# scp -P52113 shkf6-245:/opt/certs/client.pem /opt/flannel/cert/

[root@shkf6-243 flannel]# scp -P52113 shkf6-245:/opt/certs/client-key.pem /opt/flannel/cert/
```

### 5.创建配置

```shell
[root@shkf6-243 flannel]# vi subnet.env
[root@shkf6-243 flannel]# cat subnet.env
FLANNEL_NETWORK=172.6.0.0/16
FLANNEL_SUBNET=172.6.243.1/24
FLANNEL_MTU=1500
FLANNEL_IPMASQ=false
```

注意：flannel集群各主机的配置略有不同，部署其他节点时之一修改。

### 6.创建启动脚本

```shell
[root@shkf6-243 flannel]# vi flanneld.sh 
[root@shkf6-243 flannel]# cat flanneld.sh
#!/bin/sh
./flanneld \
  --public-ip=192.168.6.243 \
  --etcd-endpoints=https://192.168.6.242:2379,https://192.168.6.243:2379,https://192.168.6.244:2379 \
  --etcd-keyfile=./cert/client-key.pem \
  --etcd-certfile=./cert/client.pem \
  --etcd-cafile=./cert/ca.pem \
  --iface=eth0 \
  --subnet-file=./subnet.env \
  --healthz-port=2401
```

注意：flannel集群各主机的启动脚本略有不同，部署其他节点时注意修改

### 7.检查配置，权限，创建日志目录

```shell
[root@shkf6-243 flannel]# chmod +x /opt/flannel/flanneld.sh 
[root@shkf6-243 flannel]# mkdir -p /data/logs/flanneld
```

### 8.创建supervisor配置

```shell
[root@shkf6-243 flannel]# cat /etc/supervisord.d/flannel.ini
[program:flanneld-6-243]
command=/opt/flannel/flanneld.sh                             ; the program (relative uses PATH, can take args)
numprocs=1                                                   ; number of processes copies to start (def 1)
directory=/opt/flannel                                       ; directory to cwd to before exec (def no cwd)
autostart=true                                               ; start at supervisord start (default: true)
autorestart=true                                             ; retstart at unexpected quit (default: true)
startsecs=30                                                 ; number of secs prog must stay running (def. 1)
startretries=3                                               ; max # of serial start failures (default 3)
exitcodes=0,2                                                ; 'expected' exit codes for process (default 0,2)
stopsignal=QUIT                                              ; signal used to kill process (default TERM)
stopwaitsecs=10                                              ; max num secs to wait b4 SIGKILL (default 10)
user=root                                                    ; setuid to this UNIX account to run the program
redirect_stderr=true                                         ; redirect proc stderr to stdout (default false)
killasgroup=true                                             ; kill all process in a group
stopasgroup=true                                             ; stop all process in a group
stdout_logfile=/data/logs/flanneld/flanneld.stdout.log       ; stderr log path, NONE for none; default AUTO
stdout_logfile_maxbytes=64MB                                 ; max # logfile bytes b4 rotation (default 50MB)
stdout_logfile_backups=4                                     ; # of stdout logfile backups (default 10)
stdout_capture_maxbytes=1MB                                  ; number of bytes in 'capturemode' (default 0)
stdout_events_enabled=false                                  ; emit events on stdout writes (default false)
killasgroup=true
stopasgroup=true
```

注意：flannel集群各主机的启动脚本略有不同，部署其他节点时注意修改

supervisord管理进程的时候，默认是不kill子进程的，需要在对应的服务.ini配置文件中加以下两个配置：

```shell
killasgroup=true
stopasgroup=true
```

### 9.操作etcd，增加host-gw

```shell
[root@shkf6-243 etcd]# ./etcdctl set /coreos.com/network/config '{"Network": "172.6.0.0/16", "Backend": {"Type": "host-gw"}}'
{"Network": "172.6.0.0/16", "Backend": {"Type": "host-gw"}}

[root@shkf6-243 etcd]# ./etcdctl get /coreos.com/network/config
{"Network": "172.6.0.0/16", "Backend": {"Type": "host-gw"}}
```

Flannel的host-gw模型

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_b12d9d33897632ec412af7644b40298a_r.png)

本质就是增加静态路由：

```shell
[root@shkf6-243 flannel]# route add -net 172.6.244.0/24 gw 192.168.6.244 dev eth0

[root@shkf6-244 flannel]# route add -net 172.6.243.0/24 gw 192.168.6.244 dev eth0

[root@shkf6-243 ~]# iptables -t filter -I FORWARD -d 172.5.243.0/24 -j ACCEPT

[root@shkf6-244 ~]# iptables -t filter -I FORWARD -d 172.5.244.0/24 -j ACCEPT
```

附：flannel的其他网络模型

VxLAN模型

```shell
'{"Network": "172.6.0.0/16", "Backend": {"Type": "VxLAN"}}'
```

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_f125be65a578dc5128c832c86f3edb9e_r.png)

```shell
1、更改为VxLAN模型

1.1 拆除现有的host-gw模式

1.1.1查看当前flanneld工作模式
[root@shkf6-243 flannel]# cd /opt/etcd
[root@shkf6-243 etcd]# ./etcdctl get /coreos.com/network/config
{"Network": "172.6.0.0/16", "Backend": {"Type": "host-gw"}}

1.1.2关闭flanneld
[root@shkf6-243 flannel]# supervisorctl stop flanneld-6-243
flanneld-6-243: stopped

[root@shkf6-244 flannel]# supervisorctl stop flanneld-6-244
flanneld-6-244: stopped


1.1.3检查关闭情况
[root@shkf6-243 flannel]# ps -aux | grep flanneld
root     12784  0.0  0.0 112660   964 pts/0    S+   10:30   0:00 grep --color=auto flanneld

[root@shkf6-244 flannel]# ps -aux | grep flanneld
root     12379  0.0  0.0 112660   972 pts/0    S+   10:31   0:00 grep --color=auto flanneld

1.2拆除静态规则

1.2.1查看当前路由

[root@shkf6-243 flannel]# route -n 
Kernel IP routing table
Destination     Gateway         Genmask         Flags Metric Ref    Use Iface
0.0.0.0         192.168.6.254   0.0.0.0         UG    100    0        0 eth0
172.6.243.0     0.0.0.0         255.255.255.0   U     0      0        0 docker0
172.6.244.0     192.168.6.244   255.255.255.0   UG    0      0        0 eth0
192.168.6.0     0.0.0.0         255.255.255.0   U     100    0        0 eth0


[root@shkf6-244 flannel]# route -n
Kernel IP routing table
Destination     Gateway         Genmask         Flags Metric Ref    Use Iface
0.0.0.0         192.168.6.254   0.0.0.0         UG    100    0        0 eth0
172.6.243.0     192.168.6.243   255.255.255.0   UG    0      0        0 eth0
172.6.244.0     0.0.0.0         255.255.255.0   U     0      0        0 docker0
192.168.6.0     0.0.0.0         255.255.255.0   U     100    0        0 eth0



1.2.2拆除host-gw静态路由规则：
[root@shkf6-243 flannel]# route del -net  172.6.244.0/24 gw 192.168.6.244 dev eth0     # 删除规则的方法
[root@shkf6-244 flannel]# route del -net  172.6.243.0/24 gw 192.168.6.243 dev eth0        # 删除规则的方法


1.2.3查看拆除后的路由信息
[root@shkf6-243 flannel]# route -n
Kernel IP routing table
Destination     Gateway         Genmask         Flags Metric Ref    Use Iface
0.0.0.0         192.168.6.254   0.0.0.0         UG    100    0        0 eth0
172.6.243.0     0.0.0.0         255.255.255.0   U     0      0        0 docker0
192.168.6.0     0.0.0.0         255.255.255.0   U     100    0        0 eth0

[root@shkf6-244 flannel]# route -n
Kernel IP routing table
Destination     Gateway         Genmask         Flags Metric Ref    Use Iface
0.0.0.0         192.168.6.254   0.0.0.0         UG    100    0        0 eth0
172.6.244.0     0.0.0.0         255.255.255.0   U     0      0        0 docker0
192.168.6.0     0.0.0.0         255.255.255.0   U     100    0        0 eth0


1.3查看后端模式，并删除host-gw模型
[root@shkf6-243 flannel]# cd /opt/etcd
[root@shkf6-243 etcd]# ./etcdctl get /coreos.com/network/config
{"Network": "172.6.0.0/16", "Backend": {"Type": "host-gw"}}

[root@shkf6-243 etcd]# ./etcdctl rm /coreos.com/network/config
Error:  x509: certificate signed by unknown authority
[root@shkf6-243 etcd]# ./etcdctl get /coreos.com/network/config
{"Network": "172.6.0.0/16", "Backend": {"Type": "host-gw"}}
[root@shkf6-243 etcd]# ./etcdctl rm /coreos.com/network/config
PrevNode.Value: {"Network": "172.6.0.0/16", "Backend": {"Type": "host-gw"}}
[root@shkf6-243 etcd]# ./etcdctl get /coreos.com/network/config
Error:  100: Key not found (/coreos.com/network/config) [13]


2、更改后端模式为VxLAN

2.1后端模式添加为VxLAN
[root@shkf6-243 etcd]# ./etcdctl set /coreos.com/network/config '{"Network": "172.6.0.0/16", "Backend": {"Type": "VxLAN"}}'
{"Network": "172.6.0.0/16", "Backend": {"Type": "VxLAN"}}

2.2查看模式
[root@shkf6-243 etcd]# ./etcdctl get /coreos.com/network/config
{"Network": "172.6.0.0/16", "Backend": {"Type": "VxLAN"}}

2.3查看flanneld（VxLan）应用前网卡信息
[root@shkf6-243 etcd]# ifconfig 
docker0: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
        inet 172.6.243.1  netmask 255.255.255.0  broadcast 172.6.243.255

eth0: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
        inet 192.168.6.243  netmask 255.255.255.0  broadcast 192.168.6.255

lo: flags=73<UP,LOOPBACK,RUNNING>  mtu 65536
        inet 127.0.0.1  netmask 255.0.0.0

vethba8da49: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500


2.4启动flanneld
[root@shkf6-243 etcd]# supervisorctl start flanneld-6-243
[root@shkf6-244 flannel]# supervisorctl start flanneld-6-244


2.5查看flanneld（VxLan）应用生效后网卡信息
[root@shkf6-243 etcd]# ifconfig 
docker0: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
        inet 172.6.243.1  netmask 255.255.255.0  broadcast 172.6.243.255

eth0: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
        inet 192.168.6.243  netmask 255.255.255.0  broadcast 192.168.6.255

flannel.1: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1450
        inet 172.6.243.0  netmask 255.255.255.255  broadcast 0.0.0.0

lo: flainet 127.0.0.1  netmask 255.0.0.0

vethba8da49: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500



2.6查看静态路由信息    
[root@shkf6-243 flannel]# route -n
Kernel IP routing table
Destination     Gateway         Genmask         Flags Metric Ref    Use Iface
0.0.0.0         192.168.6.254   0.0.0.0         UG    100    0        0 eth0
172.6.243.0     0.0.0.0         255.255.255.0   U     0      0        0 docker0
172.6.244.0     172.6.244.0     255.255.255.0   UG    0      0        0 flannel.1
192.168.6.0     0.0.0.0         255.255.255.0   U     100    0        0 eth0

[root@shkf6-244 flannel]# route -n
Kernel IP routing table
Destination     Gateway         Genmask         Flags Metric Ref    Use Iface
0.0.0.0         192.168.6.254   0.0.0.0         UG    100    0        0 eth0
172.6.243.0     172.6.243.0     255.255.255.0   UG    0      0        0 flannel.1
172.6.244.0     0.0.0.0         255.255.255.0   U     0      0        0 docker0
192.168.6.0     0.0.0.0         255.255.255.0   U     100    0        0 eth0
3.恢复成host-gw模型：
3.1查看当前flanneld工作模式
[root@shkf6-243 flannel]# cd /opt/etcd
[root@shkf6-243 etcd]# ./etcdctl get /coreos.com/network/config
{"Network": "172.6.0.0/16", "Backend": {"Type": "VxLAN"}}

3.2关闭flanneld
[root@shkf6-243 flannel]# supervisorctl stop flanneld-6-243
flanneld-6-243: stopped

[root@shkf6-244 flannel]# supervisorctl stop flanneld-6-244
flanneld-6-244: stopped

3.3删除后端flanneld（VxLAN）模型
[root@shkf6-243 etcd]# ./etcdctl rm /coreos.com/network/config
PrevNode.Value: {"Network": "172.6.0.0/16", "Backend": {"Type": "VxLAN"}}

3.4查看后端当前模型
[root@shkf6-243 etcd]# ./etcdctl get /coreos.com/network/config
Error:  100: Key not found (/coreos.com/network/config) [17]

3.5后端更改host-gw模型
[root@shkf6-243 etcd]#  ./etcdctl set /coreos.com/network/config '{"Network": "172.6.0.0/16", "Backend": {"Type": "host-gw"}}'
{"Network": "172.6.0.0/16", "Backend": {"Type": "host-gw"}}
[root@shkf6-243 etcd]# ./etcdctl get /coreos.com/network/config
{"Network": "172.6.0.0/16", "Backend": {"Type": "host-gw"}}

3.6查看静态路由
[root@shkf6-243 etcd]# route -n
Kernel IP routing table
Destination     Gateway         Genmask         Flags Metric Ref    Use Iface
0.0.0.0         192.168.6.254   0.0.0.0         UG    100    0        0 eth0
172.6.243.0     0.0.0.0         255.255.255.0   U     0      0        0 docker0
172.6.244.0     192.168.6.244   255.255.255.0   UG    0      0        0 eth0
192.168.6.0     0.0.0.0         255.255.255.0   U     100    0        0 eth0

[root@shkf6-244 flannel]# route -n
Kernel IP routing table
Destination     Gateway         Genmask         Flags Metric Ref    Use Iface
0.0.0.0         192.168.6.254   0.0.0.0         UG    100    0        0 eth0
172.6.243.0     192.168.6.243   255.255.255.0   UG    0      0        0 eth0
172.6.244.0     0.0.0.0         255.255.255.0   U     0      0        0 docker0
192.168.6.0     0.0.0.0         255.255.255.0   U     100    0        0 eth0
```

Directrouting模型(直接路由)

```shell
'{"Network": "172.6.0.0/16", "Backend": {"Type": "VxLAN","Directrouting": true}}'
```

查找主：

```shell
[root@shkf6-243 etcd]# ./etcdctl member list
4244d625c76d5482: name=etcd-server-6-242 peerURLs=https://192.168.6.242:2380 clientURLs=http://127.0.0.1:2379,https://192.168.6.242:2379 isLeader=true
aa911af67b8285a2: name=etcd-server-6-243 peerURLs=https://192.168.6.243:2380 clientURLs=http://127.0.0.1:2379,https://192.168.6.243:2379 isLeader=false
c751958d48e7e127: name=etcd-server-6-244 peerURLs=https://192.168.6.244:2380 clientURLs=http://127.0.0.1:2379,https://192.168.6.244:2379 isLeader=false
```

### 10.启动服务并检查

```shell
[root@shkf6-243 flannel]# supervisorctl update
flanneld-6-243: added process group

[root@shkf6-243 flannel]# supervisorctl status
etcd-server-6-243                RUNNING   pid 28429, uptime 2 days, 0:10:55
flanneld-6-243                   STARTING  
kube-apiserver-6-243             RUNNING   pid 17808, uptime 18:50:14
kube-controller-manager-6.243    RUNNING   pid 17999, uptime 18:49:47
kube-kubelet-6-243               RUNNING   pid 28717, uptime 1 day, 23:25:50
kube-proxy-6-243                 RUNNING   pid 31546, uptime 1 day, 23:14:58
kube-scheduler-6-243             RUNNING   pid 28574, uptime 1 day, 23:39:57
```

### 11.安装部署启动检查所有集群规划的机器

略

### 12.再次验证集群，pod之间网络互通

```shell
[root@shkf6-243 flannel]# kubectl get pods -o wide
NAME             READY   STATUS    RESTARTS   AGE   IP            NODE                 NOMINATED NODE   READINESS GATES
nginx-ds-jrbdg   1/1     Running   0          9h    172.6.243.2   shkf6-243.host.com   <none>           <none>
nginx-ds-ttlx9   1/1     Running   0          9h    172.6.244.2   shkf6-244.host.com   <none>           <none>

[root@shkf6-243 flannel]# kubectl exec -it nginx-ds-jrbdg bash 
root@nginx-ds-jrbdg:/# ping 172.6.244.2
PING 172.6.244.2 (172.6.244.2): 48 data bytes
56 bytes from 172.6.244.2: icmp_seq=0 ttl=62 time=0.446 ms
56 bytes from 172.6.244.2: icmp_seq=1 ttl=62 time=0.449 ms
56 bytes from 172.6.244.2: icmp_seq=2 ttl=62 time=0.344 ms

[root@shkf6-244 flannel]# kubectl exec -it nginx-ds-ttlx9 bash
root@nginx-ds-ttlx9:/# ping 172.6.243.2
PING 172.6.243.2 (172.6.243.2): 48 data bytes
56 bytes from 172.6.243.2: icmp_seq=0 ttl=62 time=0.324 ms
56 bytes from 172.6.243.2: icmp_seq=1 ttl=62 time=0.286 ms
56 bytes from 172.6.243.2: icmp_seq=2 ttl=62 time=0.345 ms
[root@shkf6-243 flannel]# ping 172.6.244.2
PING 172.6.244.2 (172.6.244.2) 56(84) bytes of data.
64 bytes from 172.6.244.2: icmp_seq=1 ttl=63 time=0.878 ms
64 bytes from 172.6.244.2: icmp_seq=2 ttl=63 time=0.337 ms

[root@shkf6-243 flannel]# ping 172.6.243.2
PING 172.6.243.2 (172.6.243.2) 56(84) bytes of data.
64 bytes from 172.6.243.2: icmp_seq=1 ttl=64 time=0.062 ms
64 bytes from 172.6.243.2: icmp_seq=2 ttl=64 time=0.072 ms
64 bytes from 172.6.243.2: icmp_seq=3 ttl=64 time=0.071 ms


[root@shkf6-244 flannel]# ping 172.6.243.2
PING 172.6.243.2 (172.6.243.2) 56(84) bytes of data.
64 bytes from 172.6.243.2: icmp_seq=1 ttl=63 time=0.248 ms
64 bytes from 172.6.243.2: icmp_seq=2 ttl=63 time=0.125 ms

[root@shkf6-244 flannel]# ping 172.6.244.2
PING 172.6.244.2 (172.6.244.2) 56(84) bytes of data.
64 bytes from 172.6.244.2: icmp_seq=1 ttl=64 time=0.091 ms
64 bytes from 172.6.244.2: icmp_seq=2 ttl=64 time=0.064 ms
```

为甚么172.6.243.2和172.6.244.2容器能通信呢？

```shell
[root@shkf6-244 flannel]# route -n
Kernel IP routing table
Destination     Gateway         Genmask         Flags Metric Ref    Use Iface
0.0.0.0         192.168.6.254   0.0.0.0         UG    100    0        0 eth0
172.6.243.0     192.168.6.243   255.255.255.0   UG    0      0        0 eth0
172.6.244.0     0.0.0.0         255.255.255.0   U     0      0        0 docker0
192.168.6.0     0.0.0.0         255.255.255.0   U     100    0        0 eth0


[root@shkf6-243 flannel]# route -n
Kernel IP routing table
Destination     Gateway         Genmask         Flags Metric Ref    Use Iface
0.0.0.0         192.168.6.254   0.0.0.0         UG    100    0        0 eth0
172.6.243.0     0.0.0.0         255.255.255.0   U     0      0        0 docker0
172.6.244.0     192.168.6.244   255.255.255.0   UG    0      0        0 eth0
192.168.6.0     0.0.0.0         255.255.255.0   U     100    0        0 eth0
```

本质就是

```shell
[root@shkf6-243 flannel]# route add -net 172.6.244.0/24 gw 192.168.6.244 dev eth0

[root@shkf6-244 flannel]# route add -net 172.6.243.0/24 gw 192.168.6.244 dev eth0
```

### 13.在各个节点上优化iptables规则

为什么要优化？一起看下面的案例：

1.创建nginx-ds.yaml

```shell
[root@shkf6-243 ~]# cat nginx-ds.yaml 
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
        image: harbor.od.com/public/nginx:curl
        command: ["nginx","-g","daemon off;"]
        ports:
        - containerPort: 80
```

------

提示：直接用nginx:curl起不来原因是，当时打包nginx:curl 用的是bash：

```shell
docker run --rm -it  sunrisenan/nginx:v1.12.2 bash
```

1、重做nginx:curl
启动nginx，docker run -d –rm nginx:1.7.9
然后exec进去安装curl，再commit

2、仍然用之前做的nginx:curl
然后在k8s的yaml里，加入cmd指令的配置
需要显示的指明，这个镜像的CMD指令是：nginx -g “daemon off;”

```shell
command: ["nginx","-g","daemon off;"]
```

------

2.应用nginx-ds.yaml

```shell
[root@shkf6-243 ~]# kubectl apply -f nginx-ds.yaml
daemonset.extensions/nginx-ds created
```

3.可以发现，pod之间访问nginx显示的是代理的IP地址

```shell
[root@shkf6-243 ~]# kubectl get pods -o wide
NAME             READY   STATUS    RESTARTS   AGE   IP            NODE                 NOMINATED NODE   READINESS GATES
nginx-ds-86f7k   1/1     Running   0          21m   172.6.244.2   shkf6-244.host.com   <none>           <none>
nginx-ds-twfgq   1/1     Running   0          21m   172.6.243.2   shkf6-243.host.com   <none>           <none>

[root@shkf6-243 ~]# kubectl exec -it nginx-ds-86f7k /bin/sh
# curl 172.6.243.2

[root@shkf6-244 ~]# kubectl log -f nginx-ds-twfgq
log is DEPRECATED and will be removed in a future version. Use logs instead.
192.168.6.244 - - [21/Nov/2019:06:56:35 +0000] "GET / HTTP/1.1" 200 612 "-" "curl/7.38.0" "-"
```

4.优化pod之间不走原地址nat

在所有需要优化的机器上安装iptables

```shell
 ~]# yum install iptables-services -y

 ~]# systemctl start iptables.service
 ~]# systemctl enable iptables.service
```

优化shkf6-243

```shell
[root@shkf6-243 ~]# iptables-save|grep -i postrouting
:POSTROUTING ACCEPT [15:912]
:KUBE-POSTROUTING - [0:0]
-A POSTROUTING -m comment --comment "kubernetes postrouting rules" -j KUBE-POSTROUTING
-A POSTROUTING -s 172.6.243.0/24 ! -o docker0 -j MASQUERADE
-A KUBE-POSTROUTING -m comment --comment "kubernetes service traffic requiring SNAT" -m mark --mark 0x4000/0x4000 -j MASQUERADE

[root@shkf6-243 ~]# iptables -t nat -D POSTROUTING -s 172.6.243.0/24 ! -o docker0 -j MASQUERADE
[root@shkf6-243 ~]# iptables -t nat -I POSTROUTING -s 172.6.243.0/24 ! -d 172.6.0.0/16 ! -o docker0 -j MASQUERADE

[root@shkf6-243 ~]# iptables-save|grep -i postrouting
:POSTROUTING ACCEPT [8:488]
:KUBE-POSTROUTING - [0:0]
-A POSTROUTING -s 172.6.243.0/24 ! -d 172.6.0.0/16 ! -o docker0 -j MASQUERADE
-A POSTROUTING -m comment --comment "kubernetes postrouting rules" -j KUBE-POSTROUTING
-A KUBE-POSTROUTING -m comment --comment "kubernetes service traffic requiring SNAT" -m mark --mark 0x4000/0x4000 -j MASQUERADE
```

优化shkf6-244

```shell
[root@shkf6-244 ~]# iptables-save|grep -i postrouting
:POSTROUTING ACCEPT [7:424]
:KUBE-POSTROUTING - [0:0]
-A POSTROUTING -m comment --comment "kubernetes postrouting rules" -j KUBE-POSTROUTING
-A POSTROUTING -s 172.6.244.0/24 ! -o docker0 -j MASQUERADE
-A KUBE-POSTROUTING -m comment --comment "kubernetes service traffic requiring SNAT" -m mark --mark 0x4000/0x4000 -j MASQUERADE

[root@shkf6-244 ~]# iptables -t nat -D POSTROUTING -s 172.6.244.0/24 ! -o docker0 -j MASQUERADE
[root@shkf6-244 ~]# iptables -t nat -I POSTROUTING -s 172.6.244.0/24 ! -d 172.6.0.0/16 ! -o docker0 -j MASQUERADE

[root@shkf6-244 ~]# iptables-save|grep -i postrouting
:POSTROUTING ACCEPT [2:120]
:KUBE-POSTROUTING - [0:0]
-A POSTROUTING -s 172.6.244.0/24 ! -d 172.6.0.0/16 ! -o docker0 -j MASQUERADE
-A POSTROUTING -m comment --comment "kubernetes postrouting rules" -j KUBE-POSTROUTING
-A KUBE-POSTROUTING -m comment --comment "kubernetes service traffic requiring SNAT" -m mark --mark 0x4000/0x4000 -j MASQUERADE
```

> 192.168.6.243主机上的，来源是172.6.243.0/24段的docker的ip，目标ip不是172.6.0.0/16段，网络发包不从docker0桥设备出站的，才进行转换
>
> 192.168.6.244主机上的，来源是172.6.244.0/24段的docker的ip，目标ip不是172.6.0.0/16段，网络发包不从docker0桥设备出站的，才进行转换

5.把默认禁止规则删掉

```shell
root@shkf6-243 ~]# iptables-save | grep -i reject
-A INPUT -j REJECT --reject-with icmp-host-prohibited
-A FORWARD -j REJECT --reject-with icmp-host-prohibited
[root@shkf6-243 ~]# iptables -t filter -D INPUT -j REJECT --reject-with icmp-host-prohibited
[root@shkf6-243 ~]# iptables -t filter -D FORWARD -j REJECT --reject-with icmp-host-prohibited
[root@shkf6-243 ~]# iptables-save > /etc/sysconfig/iptables


[root@shkf6-244 ~]# iptables-save | grep -i reject
-A INPUT -j REJECT --reject-with icmp-host-prohibited
-A FORWARD -j REJECT --reject-with icmp-host-prohibited
[root@shkf6-244 ~]# iptables -t filter -D INPUT -j REJECT --reject-with icmp-host-prohibited
[root@shkf6-244 ~]# iptables -t filter -D FORWARD -j REJECT --reject-with icmp-host-prohibited
[root@shkf6-244 ~]# iptables-save > /etc/sysconfig/iptables
```

6.优化SNAT规则，各运算节点之间的各pod之间的网络通信不再出网

测试6.244

```shell
[root@shkf6-243 ~]# kubectl get pods -o wide
NAME             READY   STATUS    RESTARTS   AGE   IP            NODE                 NOMINATED NODE   READINESS GATES
nginx-ds-86f7k   1/1     Running   0          70m   172.6.244.2   shkf6-244.host.com   <none>           <none>
nginx-ds-twfgq   1/1     Running   0          70m   172.6.243.2   shkf6-243.host.com   <none>           <none>

[root@shkf6-243 ~]# curl 172.6.243.2

[root@shkf6-244 ~]# kubectl log -f nginx-ds-86f7k
log is DEPRECATED and will be removed in a future version. Use logs instead.
192.168.6.243 - - [21/Nov/2019:07:56:22 +0000] "GET / HTTP/1.1" 200 612 "-" "curl/7.29.0" "-"


[root@shkf6-243 ~]# kubectl exec -it  nginx-ds-twfgq /bin/sh
# curl 172.6.244.2

[root@shkf6-244 ~]# kubectl log -f nginx-ds-86f7k
log is DEPRECATED and will be removed in a future version. Use logs instead.
172.6.243.2 - - [21/Nov/2019:07:57:17 +0000] "GET / HTTP/1.1" 200 612 "-" "curl/7.38.0" "-"
```

测试6.243

```shell
[root@shkf6-244 ~]# kubectl get pods -o wide
NAME             READY   STATUS    RESTARTS   AGE   IP            NODE                 NOMINATED NODE   READINESS GATES
nginx-ds-86f7k   1/1     Running   0          81m   172.6.244.2   shkf6-244.host.com   <none>           <none>
nginx-ds-twfgq   1/1     Running   0          81m   172.6.243.2   shkf6-243.host.com   <none>           <none>

[root@shkf6-244 ~]# curl 172.6.243.2

[root@shkf6-243 ~]# kubectl log -f nginx-ds-twfgq
192.168.6.244 - - [21/Nov/2019:08:01:38 +0000] "GET / HTTP/1.1" 200 612 "-" "curl/7.29.0" "-"


[root@shkf6-244 ~]# kubectl exec -it nginx-ds-86f7k /bin/sh
# curl 172.6.243.2

[root@shkf6-243 ~]# kubectl log -f nginx-ds-twfgq
172.6.244.2 - - [21/Nov/2019:08:02:52 +0000] "GET / HTTP/1.1" 200 612 "-" "curl/7.38.0" "-"
```

### 14.在运算节点保存iptables规则

```shell
 [root@shkf6-243 ~]# iptables-save > /etc/sysconfig/iptables
 [root@shkf6-244 ~]# iptables-save > /etc/sysconfig/iptables
 # 上面这中保存重启系统失效，建议用下面的方式

 [root@shkf6-243 ~]# service iptables save
 [root@shkf6-244 ~]# service iptables save
```

# 第三章：kubernetes的服务发现插件–coredns

- 简单来说，服务发现就是服务(应用)之间相互定位的过程。
- 服务发现并非云计算时代独有的，传统的单体架构时代也会用到。以下应用场景下，更需要服务发现
  - 服务(应用)的动态性强
  - 服务(应用)更新发布频繁
  - 服务(应用)支持自动伸缩

- 在K8S集群里，POD的IP是不断变化的，如何“以不变应万变“呢？
  - 抽象出了Service资源，通过标签选择器，关联一组POD
  - 抽象出了集群网络，通过相对固定的“集群IP”，使服务接入点固定

- 那么如何自动关联Service资源的“名称”和“集群网络IP”，从而达到服务被集群自动发现的目的呢？
  - 考虑传统DNS的模型：shkf6-243.host.com –> 192.168.6.243
  - 能否在K8S里建立这样的模型：nginx-ds –> 10.96.0.5
- K8S里服务发现的方式–DNS
- 实现K8S里DNS功能的插件（软件）
  - Kube-dns—kubernetes-v1.2至kubernetes-v1.10
  - Coredns—kubernetes-v1.11至今
- 注意：
  - K8S里的DNS不是万能的！它应该只负责自动维护“服务名”—>“集群网络IP”之间的关系

## 1.部署K8S的内网资源配置清单http服务

> 在运维主机shkf6-245上，配置一个nginx想虚拟主机，用以提供k8s统一资源配置清单访问入口

- 配置nginx

```shell
[root@shkf6-245 ~]# cat /etc/nginx/conf.d/k8s-yaml.od.com.conf
server {
    listen 80;
    server_name k8s-yaml.od.com;

    location / {
        autoindex on;
        default_type text/plain;
        root /data/k8s-yaml;
   }
}

[root@shkf6-245 ~]# mkdir -p /data/k8s-yaml/coredns
[root@shkf6-245 ~]# nginx -t
nginx: the configuration file /etc/nginx/nginx.conf syntax is ok
nginx: configuration file /etc/nginx/nginx.conf test is successful
[root@shkf6-245 ~]# nginx -s reload
```

- 配置内网DNS解析

在shkf6-241机器上：

```shell
[root@shkf6-241 ~]# vim /var/named/od.com.zone 

k8s-yaml    A    192.168.6.245
```

注意： `2019111204 ; serial` #序列号前滚

```shell
[root@shkf6-241 ~]# named-checkconf 
[root@shkf6-241 ~]# systemctl restart named
[root@shkf6-241 ~]# dig -t A k8s-yaml.od.com @192.168.6.241 +short
192.168.6.245
```

以后多有的资源配置清单统一放置在运维主机上的`/data/k8s-yaml`目录即可

## 2.部署coredns

[coredns官方Guthub](https://github.com/coredns/coredns)

[coredns官方DockerHub](https://hub.docker.com/r/coredns/coredns/tags)

### 1.准备coredns-v1.6.1镜像

在运维主机shkf6-245上：

```shell
[root@shkf6-245 ~]# docker pull coredns/coredns:1.6.1
1.6.1: Pulling from coredns/coredns
c6568d217a00: Pull complete 
d7ef34146932: Pull complete 
Digest: sha256:9ae3b6fcac4ee821362277de6bd8fd2236fa7d3e19af2ef0406d80b595620a7a
Status: Downloaded newer image for coredns/coredns:1.6.1
docker.io/coredns/coredns:1.6.1

[root@shkf6-245 ~]# docker images|grep coredns
coredns/coredns                 1.6.1                      c0f6e815079e        3 months ago        42.2MB

[root@shkf6-245 ~]# docker tag c0f6e815079e harbor.od.com/public/coredns:v1.6.1
[root@shkf6-245 ~]# docker push !$
docker push harbor.od.com/public/coredns:v1.6.1
The push refers to repository [harbor.od.com/public/coredns]
da1ec456edc8: Pushed 
225df95e717c: Pushed 
v1.6.1: digest: sha256:c7bf0ce4123212c87db74050d4cbab77d8f7e0b49c041e894a35ef15827cf938 size: 739
```

### 2.准备资源配置清单

在运维主机shkf6-245.host.com上：

```shell
[root@shkf6-245 ~]# mkdir -p /data/k8s-yaml/coredns && cd /data/k8s-yaml/coredns/
```

RBAC

```shell
[root@shkf6-245 coredns]# cat rbac.yaml 
apiVersion: v1
kind: ServiceAccount
metadata:
  name: coredns
  namespace: kube-system
  labels:
      kubernetes.io/cluster-service: "true"
      addonmanager.kubernetes.io/mode: Reconcile
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  labels:
    kubernetes.io/bootstrapping: rbac-defaults
    addonmanager.kubernetes.io/mode: Reconcile
  name: system:coredns
rules:
- apiGroups:
  - ""
  resources:
  - endpoints
  - services
  - pods
  - namespaces
  verbs:
  - list
  - watch
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  annotations:
    rbac.authorization.kubernetes.io/autoupdate: "true"
  labels:
    kubernetes.io/bootstrapping: rbac-defaults
    addonmanager.kubernetes.io/mode: EnsureExists
  name: system:coredns
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: system:coredns
subjects:
- kind: ServiceAccount
  name: coredns
  namespace: kube-system
```

ConfigMap

```shell
[root@shkf6-245 coredns]# cat cm.yaml 
apiVersion: v1
kind: ConfigMap
metadata:
  name: coredns
  namespace: kube-system
data:
  Corefile: |
    .:53 {
        errors
        log
        health
        ready
        kubernetes cluster.local 10.96.0.0/22
        forward . 192.168.6.241
        cache 30
        loop
        reload
        loadbalance
       }
```

Deployment

```shell
[root@shkf6-245 coredns]# cat dp.yaml 
apiVersion: apps/v1
kind: Deployment
metadata:
  name: coredns
  namespace: kube-system
  labels:
    k8s-app: coredns
    kubernetes.io/name: "CoreDNS"
spec:
  replicas: 1
  selector:
    matchLabels:
      k8s-app: coredns
  template:
    metadata:
      labels:
        k8s-app: coredns
    spec:
      priorityClassName: system-cluster-critical
      serviceAccountName: coredns
      containers:
      - name: coredns
        image: harbor.od.com/public/coredns:v1.6.1
        args:
        - -conf
        - /etc/coredns/Corefile
        volumeMounts:
        - name: config-volume
          mountPath: /etc/coredns
        ports:
        - containerPort: 53
          name: dns
          protocol: UDP
        - containerPort: 53
          name: dns-tcp
          protocol: TCP
        - containerPort: 9153
          name: metrics
          protocol: TCP
        livenessProbe:
          httpGet:
            path: /health
            port: 8080
            scheme: HTTP
          initialDelaySeconds: 60
          timeoutSeconds: 5
          successThreshold: 1
          failureThreshold: 5
      dnsPolicy: Default
      volumes:
        - name: config-volume
          configMap:
            name: coredns
            items:
            - key: Corefile
              path: Corefile
```

Service

```shell
[root@shkf6-245 coredns]# cat svc.yaml 
apiVersion: v1
kind: Service
metadata:
  name: coredns
  namespace: kube-system
  labels:
    k8s-app: coredns
    kubernetes.io/cluster-service: "true"
    kubernetes.io/name: "CoreDNS"
spec:
  selector:
    k8s-app: coredns
  clusterIP: 10.96.0.2
  ports:
  - name: dns
    port: 53
    protocol: UDP
  - name: dns-tcp
    port: 53
  - name: metrics
    port: 9153
    protocol: TCP
```

### 3.依次执行创建

```shell
[root@shkf6-243 ~]# kubectl apply -f http://k8s-yaml.od.com/coredns/rbac.yaml
[root@shkf6-243 ~]# kubectl apply -f http://k8s-yaml.od.com/coredns/cm.yaml
[root@shkf6-243 ~]# kubectl apply -f http://k8s-yaml.od.com/coredns/dp.yaml
[root@shkf6-243 ~]# kubectl apply -f http://k8s-yaml.od.com/coredns/svc.yaml
```

### 4.检查

```shell
[root@shkf6-243 ~]# kubectl get all -n kube-system
NAME                           READY   STATUS    RESTARTS   AGE
pod/coredns-6b6c4f9648-x5zvz   1/1     Running   0          7m37s


NAME              TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)                  AGE
service/coredns   ClusterIP   10.96.0.2    <none>        53/UDP,53/TCP,9153/TCP   7m27s


NAME                      READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/coredns   1/1     1            1           7m37s

NAME                                 DESIRED   CURRENT   READY   AGE
replicaset.apps/coredns-6b6c4f9648   1         1         1       7m37s
[root@shkf6-243 ~]# dig -t A www.baidu.com @10.96.0.2 +short
www.a.shifen.com.
180.101.49.12
180.101.49.11

[root@shkf6-243 ~]# dig -t A shkf6-245.host.com @10.96.0.2 +short
192.168.6.245
[root@shkf6-243 ~]# kubectl create deployment nginx-dp --image=harbor.od.com/public/nginx:v1.7.9 -n kube-public
deployment.apps/nginx-dp created

[root@shkf6-243 ~]# kubectl expose deployment nginx-dp --port=80 -n kube-public
service/nginx-dp exposed

[root@shkf6-243 ~]# kubectl get service -n kube-public
NAME       TYPE        CLUSTER-IP    EXTERNAL-IP   PORT(S)   AGE
nginx-dp   ClusterIP   10.96.3.154   <none>        80/TCP    38s
[root@shkf6-243 ~]# dig -t A nginx-dp.kube-public.svc.cluster.local. @10.96.0.2 +short
10.96.3.154
```

# 第四章：kubernetes的服务暴露插件–traefik

- K8S的DNS实现了服务在集群“内”被自动发现，那如何是的服务在K8S集群“外”被使用和访问呢？
  - 使用NodePort型的Service
    - 注意：无法使用kube-proxy的ipvs模型，只能使用iptables模型
  - 使用Ingress资源
    - 注意：Ingress只能调度并暴露7层应用，特指http和https协议
- Ingress是K8S API的标准资源类型之一，也是一种核心资源，它其实就是一组基于域名和URL路径，把用户的请求转发至指定Service资源的规则
- 可以将集群外部的请求流量，转发至集群内部，从而实现“服务暴露”
- Ingress控制器是能够为Ingress资源监听某套接字，然后根据Ingress规则匹配机制路由调度流量的一个插件
- 说白了，Ingress没啥神秘的，就是个nginx+一段go脚本而已
- 常用的Ingress控制器的实现软件
  - Ingress-nginx
  - HAProxy
  - Traefik
  - …

## 1.使用NodePort型Service暴露服务

注意：使用这种方法暴露服务，要求kube-proxy的代理类型为：iptables

```shell
1、第一步更改为proxy-mode更改为iptables，调度方式为RR
[root@shkf6-243 ~]# cat /opt/kubernetes/server/bin/kube-proxy.sh
#!/bin/sh
./kube-proxy \
  --cluster-cidr 172.6.0.0/16 \
  --hostname-override shkf6-243.host.com \
  --proxy-mode=iptables \
  --ipvs-scheduler=rr \
  --kubeconfig ./conf/kube-proxy.kubeconfig

[root@shkf6-244 ~]# cat /opt/kubernetes/server/bin/kube-proxy.sh
#!/bin/sh
./kube-proxy \
  --cluster-cidr 172.6.0.0/16 \
  --hostname-override shkf6-243.host.com \
  --proxy-mode=iptables \
  --ipvs-scheduler=rr \
  --kubeconfig ./conf/kube-proxy.kubeconfig


2.使iptables模式生效
[root@shkf6-243 ~]# supervisorctl restart kube-proxy-6-243
kube-proxy-6-243: stopped
kube-proxy-6-243: started
[root@shkf6-243 ~]# ps -ef|grep kube-proxy
root     26694 12008  0 10:25 ?        00:00:00 /bin/sh /opt/kubernetes/server/bin/kube-proxy.sh
root     26695 26694  0 10:25 ?        00:00:00 ./kube-proxy --cluster-cidr 172.6.0.0/16 --hostname-override shkf6-243.host.com --proxy-mode=iptables --ipvs-scheduler=rr --kubeconfig ./conf/kube-proxy.kubeconfig
root     26905 13466  0 10:26 pts/0    00:00:00 grep --color=auto kube-proxy


[root@shkf6-244 ~]# supervisorctl restart kube-proxy-6-244
kube-proxy-6-244: stopped
kube-proxy-6-244kube-proxy-6-244: started
[root@shkf6-244 ~]# ps -ef|grep kube-proxy
root     25998 11030  0 10:22 ?        00:00:00 /bin/sh /opt/kubernetes/server/bin/kube-proxy.sh
root     25999 25998  0 10:22 ?        00:00:00 ./kube-proxy --cluster-cidr 172.6.0.0/16 --hostname-override shkf6-243.host.com --proxy-mode=iptables --ipvs-scheduler=rr --kubeconfig ./conf/kube-proxy.kubeconfig


[root@shkf6-243 ~]# tail -fn 11 /data/logs/kubernetes/kube-proxy/proxy.stdout.log 

[root@shkf6-244 ~]# tail -fn 11 /data/logs/kubernetes/kube-proxy/proxy.stdout.log


3.清理现有的ipvs规则
[root@shkf6-243 ~]# ipvsadm -Ln
IP Virtual Server version 1.2.1 (size=4096)
Prot LocalAddress:Port Scheduler Flags
  -> RemoteAddress:Port           Forward Weight ActiveConn InActConn
TCP  10.96.0.1:443 nq
  -> 192.168.6.243:6443           Masq    1      0          0         
  -> 192.168.6.244:6443           Masq    1      0          0         
TCP  10.96.0.2:53 nq
  -> 172.6.244.3:53               Masq    1      0          0         
TCP  10.96.0.2:9153 nq
  -> 172.6.244.3:9153             Masq    1      0          0         
TCP  10.96.3.154:80 nq
  -> 172.6.243.3:80               Masq    1      0          0         
UDP  10.96.0.2:53 nq
  -> 172.6.244.3:53               Masq    1      0          0         
[root@shkf6-243 ~]# ipvsadm -D -t 10.96.0.1:443
[root@shkf6-243 ~]# ipvsadm -D -t 10.96.0.2:53
[root@shkf6-243 ~]# ipvsadm -D -t 10.96.0.2:9153
[root@shkf6-243 ~]# ipvsadm -D -t 10.96.3.154:80
[root@shkf6-243 ~]# ipvsadm -D -u 10.96.0.2:53
[root@shkf6-243 ~]# ipvsadm -Ln
IP Virtual Server version 1.2.1 (size=4096)
Prot LocalAddress:Port Scheduler Flags
  -> RemoteAddress:Port           Forward Weight ActiveConn InActConn


[root@shkf6-244 ~]# ipvsadm -Ln
IP Virtual Server version 1.2.1 (size=4096)
Prot LocalAddress:Port Scheduler Flags
  -> RemoteAddress:Port           Forward Weight ActiveConn InActConn
TCP  10.96.0.1:443 nq
  -> 192.168.6.243:6443           Masq    1      0          0         
  -> 192.168.6.244:6443           Masq    1      1          0         
TCP  10.96.0.2:53 nq
  -> 172.6.244.3:53               Masq    1      0          0         
TCP  10.96.0.2:9153 nq
  -> 172.6.244.3:9153             Masq    1      0          0         
TCP  10.96.3.154:80 nq
  -> 172.6.243.3:80               Masq    1      0          0         
UDP  10.96.0.2:53 nq
  -> 172.6.244.3:53               Masq    1      0          0         
[root@shkf6-244 ~]# ipvsadm -D -t 10.96.0.1:443
[root@shkf6-244 ~]# ipvsadm -D -t 10.96.0.2:53
[root@shkf6-244 ~]# ipvsadm -D -t 10.96.0.2:9153
[root@shkf6-244 ~]# ipvsadm -D -t 10.96.3.154:80
[root@shkf6-244 ~]# ipvsadm -D -u 10.96.0.2:53
[root@shkf6-244 ~]# ipvsadm -Ln
IP Virtual Server version 1.2.1 (size=4096)
Prot LocalAddress:Port Scheduler Flags
  -> RemoteAddress:Port           Forward Weight ActiveConn InActConn
```

### 1.修改nginx-ds的service资源配置清单

```shell
[root@shkf6-243 ~]# cat /root/nginx-ds-svc.yaml
apiVersion: v1
kind: Service
metadata:
  labels:
    app: nginx-ds
  name: nginx-ds
  namespace: default
spec:
  ports:
  - port: 80
    protocol: TCP
    nodePort: 8000
  selector:
    app: nginx-ds
  sessionAffinity: None
  type: NodePort
[root@shkf6-243 ~]# kubectl apply -f nginx-ds-svc.yaml 
service/nginx-ds created
```

### 2.重建nginx-ds的service资源

```shell
[root@shkf6-243 ~]# cat nginx-ds.yaml 
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
        image: harbor.od.com/public/nginx:curl
        command: ["nginx","-g","daemon off;"]
        ports:
        - containerPort: 80
[root@shkf6-243 ~]# kubectl apply -f nginx-ds.yaml 
daemonset.extensions/nginx-ds created
```

### 3.查看service

```shell
[root@shkf6-243 ~]# kubectl get svc nginx-ds
NAME       TYPE       CLUSTER-IP    EXTERNAL-IP   PORT(S)       AGE
nginx-ds   NodePort   10.96.1.170   <none>        80:8000/TCP   2m20s

[root@shkf6-243 ~]# netstat -lntup|grep 8000
tcp6       0      0 :::8000                 :::*                    LISTEN      26695/./kube-proxy

[root@shkf6-244 ~]# netstat -lntup|grep 8000
tcp6       0      0 :::8000                 :::*                    LISTEN      25999/./kube-proxy 
```

- nodePort本质

```shell
[root@shkf6-244 ~]# iptables-save | grep 8000
-A KUBE-FIREWALL -m comment --comment "kubernetes firewall for dropping marked packets" -m mark --mark 0x8000/0x8000 -j DROP
-A KUBE-MARK-DROP -j MARK --set-xmark 0x8000/0x8000
-A KUBE-NODEPORTS -p tcp -m comment --comment "default/nginx-ds:" -m tcp --dport 8000 -j KUBE-MARK-MASQ
-A KUBE-NODEPORTS -p tcp -m comment --comment "default/nginx-ds:" -m tcp --dport 8000 -j KUBE-SVC-T4RQBNWQFKKBCRET
```

### 4.浏览器访问

访问：http://192.168.6.243:8000/

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_219520c675b6c8faf33df7c9ebaa07bf_r.png)

访问：http://192.168.6.244:8000/

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_7103dd03286a23073e5055cd87a272c1_r.png)

### 5.还原成ipvs

删除service和pod

```shell
[root@shkf6-243 ~]# kubectl delete -f nginx-ds.yaml 
daemonset.extensions "nginx-ds" deleted

[root@shkf6-243 ~]# kubectl delete -f nginx-ds-svc.yaml 
service "nginx-ds" deleted

[root@shkf6-243 ~]# netstat -lntup|grep 8000
[root@shkf6-243 ~]# 
```

在运算节点上：

```shell
[root@shkf6-243 ~]# cat /opt/kubernetes/server/bin/kube-proxy.sh
#!/bin/sh
./kube-proxy \
  --cluster-cidr 172.6.0.0/16 \
  --hostname-override shkf6-243.host.com \
  --proxy-mode=ipvs \
  --ipvs-scheduler=nq \
  --kubeconfig ./conf/kube-proxy.kubeconfig

[root@shkf6-243 ~]# supervisorctl restart kube-proxy-6-243
kube-proxy-6-243: stopped
kube-proxy-6-243: started

[root@shkf6-243 ~]# ipvsadm -Ln
IP Virtual Server version 1.2.1 (size=4096)
Prot LocalAddress:Port Scheduler Flags
  -> RemoteAddress:Port           Forward Weight ActiveConn InActConn
TCP  10.96.0.1:443 nq
  -> 192.168.6.243:6443           Masq    1      0          0         
  -> 192.168.6.244:6443           Masq    1      0          0         
TCP  10.96.0.2:53 nq
  -> 172.6.244.3:53               Masq    1      0          0         
TCP  10.96.0.2:9153 nq
  -> 172.6.244.3:9153             Masq    1      0          0         
TCP  10.96.3.154:80 nq
  -> 172.6.243.3:80               Masq    1      0          0         
UDP  10.96.0.2:53 nq
  -> 172.6.244.3:53               Masq    1      0          0   
[root@shkf6-244 ~]# cat /opt/kubernetes/server/bin/kube-proxy.sh
#!/bin/sh
./kube-proxy \
  --cluster-cidr 172.6.0.0/16 \
  --hostname-override shkf6-243.host.com \
  --proxy-mode=ipvs \
  --ipvs-scheduler=nq \
  --kubeconfig ./conf/kube-proxy.kubeconfig

[root@shkf6-244 ~]# supervisorctl restart kube-proxy-6-244
kube-proxy-6-244: stopped
kube-proxy-6-244: started

[root@shkf6-244 ~]# ipvsadm -Ln
IP Virtual Server version 1.2.1 (size=4096)
Prot LocalAddress:Port Scheduler Flags
  -> RemoteAddress:Port           Forward Weight ActiveConn InActConn
TCP  10.96.0.1:443 nq
  -> 192.168.6.243:6443           Masq    1      0          0         
  -> 192.168.6.244:6443           Masq    1      0          0         
TCP  10.96.0.2:53 nq
  -> 172.6.244.3:53               Masq    1      0          0         
TCP  10.96.0.2:9153 nq
  -> 172.6.244.3:9153             Masq    1      0          0         
TCP  10.96.3.154:80 nq
  -> 172.6.243.3:80               Masq    1      0          0         
UDP  10.96.0.2:53 nq
  -> 172.6.244.3:53               Masq    1      0          0   
```

## 2.部署traefik（ingress控制器）

[traefik官方GitHub](https://github.com/containous/traefik)

[traefik官方DockerHub](https://hub.docker.com/_/traefik)

### 1.准备traefik镜像

在运维主机上shkf6-245.host.com上：

```shell
[root@shkf6-245 traefik]# docker pull traefik:v1.7.2-alpine
v1.7.2-alpine: Pulling from library/traefik
4fe2ade4980c: Pull complete 
8d9593d002f4: Pull complete 
5d09ab10efbd: Pull complete 
37b796c58adc: Pull complete 
Digest: sha256:cf30141936f73599e1a46355592d08c88d74bd291f05104fe11a8bcce447c044
Status: Downloaded newer image for traefik:v1.7.2-alpine
docker.io/library/traefik:v1.7.2-alpine

[root@shkf6-245 traefik]# docker images|grep traefik
traefik                         v1.7.2-alpine              add5fac61ae5        13 months ago       72.4MB

[root@shkf6-245 traefik]# docker tag add5fac61ae5 harbor.od.com/public/traefik:v1.7.2
[root@shkf6-245 traefik]# docker push !$
docker push harbor.od.com/public/traefik:v1.7.2
The push refers to repository [harbor.od.com/public/traefik]
a02beb48577f: Pushed 
ca22117205f4: Pushed 
3563c211d861: Pushed 
df64d3292fd6: Pushed 
v1.7.2: digest: sha256:6115155b261707b642341b065cd3fac2b546559ba035d0262650b3b3bbdd10ea size: 1157
```

### 2.准备资源配置清单

```shell
[root@shkf6-245 traefik]# cat /data/k8s-yaml/traefik/rbac.yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: traefik-ingress-controller
  namespace: kube-system
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRole
metadata:
  name: traefik-ingress-controller
rules:
  - apiGroups:
      - ""
    resources:
      - services
      - endpoints
      - secrets
    verbs:
      - get
      - list
      - watch
  - apiGroups:
      - extensions
    resources:
      - ingresses
    verbs:
      - get
      - list
      - watch
---
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1beta1
metadata:
  name: traefik-ingress-controller
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: traefik-ingress-controller
subjects:
- kind: ServiceAccount
  name: traefik-ingress-controller
  namespace: kube-system
[root@shkf6-245 traefik]# cat /data/k8s-yaml/traefik/ds.yaml
apiVersion: extensions/v1beta1
kind: DaemonSet
metadata:
  name: traefik-ingress
  namespace: kube-system
  labels:
    k8s-app: traefik-ingress
spec:
  template:
    metadata:
      labels:
        k8s-app: traefik-ingress
        name: traefik-ingress
    spec:
      serviceAccountName: traefik-ingress-controller
      terminationGracePeriodSeconds: 60
      containers:
      - image: harbor.od.com/public/traefik:v1.7.2
        name: traefik-ingress
        ports:
        - name: controller
          containerPort: 80
          hostPort: 81
        - name: admin-web
          containerPort: 8080
        securityContext:
          capabilities:
            drop:
            - ALL
            add:
            - NET_BIND_SERVICE
        args:
        - --api
        - --kubernetes
        - --logLevel=INFO
        - --insecureskipverify=true
        - --kubernetes.endpoint=https://192.168.6.66:7443
        - --accesslog
        - --accesslog.filepath=/var/log/traefik_access.log
        - --traefiklog
        - --traefiklog.filepath=/var/log/traefik.log
        - --metrics.prometheus
[root@shkf6-245 traefik]# cat /data/k8s-yaml/traefik/svc.yaml
kind: Service
apiVersion: v1
metadata:
  name: traefik-ingress-service
  namespace: kube-system
spec:
  selector:
    k8s-app: traefik-ingress
  ports:
    - protocol: TCP
      port: 80
      name: controller
    - protocol: TCP
      port: 8080
      name: admin-web
[root@shkf6-245 traefik]# cat /data/k8s-yaml/traefik/ingress.yaml 
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: traefik-web-ui
  namespace: kube-system
  annotations:
    kubernetes.io/ingress.class: traefik
spec:
  rules:
  - host: traefik.od.com
    http:
      paths:
      - path: /
        backend:
          serviceName: traefik-ingress-service
          servicePort: 8080
```

### 3.依次执行创建

```shell
[root@shkf6-243 ~]# kubectl apply -f http://k8s-yaml.od.com/traefik/rbac.yaml
[root@shkf6-243 ~]# kubectl apply -f http://k8s-yaml.od.com/traefik/ds.yaml
[root@shkf6-243 ~]# kubectl apply -f http://k8s-yaml.od.com/traefik/svc.yaml
[root@shkf6-243 ~]# kubectl apply -f http://k8s-yaml.od.com/traefik/ingress.yaml
```

## 3.解析域名

```shell
[root@shkf6-241 ~]# cat /var/named/od.com.zone 
$ORIGIN od.com.
$TTL 600    ; 10 minutes
@           IN SOA    dns.od.com. dnsadmin.od.com. (
                2019111205 ; serial
                10800      ; refresh (3 hours)
                900        ; retry (15 minutes)
                604800     ; expire (1 week)
                86400      ; minimum (1 day)
                )
                NS   dns.od.com.
$TTL 60    ; 1 minute
dns                A    192.168.6.241
harbor             A    192.168.6.245
k8s-yaml           A    192.168.6.245
traefik            A    192.168.6.66

[root@shkf6-241 ~]# named-checkconf 
[root@shkf6-241 ~]# systemctl restart named
```

## 4.配置反向代理

```shell
[root@shkf6-241 ~]# cat /etc/nginx/conf.d/od.com.conf
upstream default_backend_traefik {
    server 192.168.6.243:81    max_fails=3 fail_timeout=10s;
    server 192.168.6.244:81    max_fails=3 fail_timeout=10s;
}
server {
    server_name *.od.com;

    location / {
        proxy_pass http://default_backend_traefik;
        proxy_set_header Host       $http_host;
        proxy_set_header x-forwarded-for $proxy_add_x_forwarded_for;
    }
}
[root@shkf6-241 ~]# nginx -t
nginx: the configuration file /etc/nginx/nginx.conf syntax is ok
nginx: configuration file /etc/nginx/nginx.conf test is successful
[root@shkf6-241 ~]# nginx -s reload
[root@shkf6-242 ~]# vim /etc/nginx/conf.d/od.com.conf
[root@shkf6-242 ~]# cat /etc/nginx/conf.d/od.com.conf
upstream default_backend_traefik {
    server 192.168.6.243:81    max_fails=3 fail_timeout=10s;
    server 192.168.6.244:81    max_fails=3 fail_timeout=10s;
}
server {
    server_name *.od.com;

    location / {
        proxy_pass http://default_backend_traefik;
        proxy_set_header Host       $http_host;
        proxy_set_header x-forwarded-for $proxy_add_x_forwarded_for;
    }
}
[root@shkf6-242 ~]# nginx -t
nginx: the configuration file /etc/nginx/nginx.conf syntax is ok
nginx: configuration file /etc/nginx/nginx.conf test is successful
[root@shkf6-242 ~]# nginx -s reload
```

## 5.检查

```shell
[root@shkf6-244 ~]# kubectl get all -n kube-system 
NAME                           READY   STATUS    RESTARTS   AGE
pod/coredns-6b6c4f9648-x5zvz   1/1     Running   0          18h
pod/traefik-ingress-bhhkv      1/1     Running   0          17m
pod/traefik-ingress-mm2ds      1/1     Running   0          17m


NAME                              TYPE        CLUSTER-IP    EXTERNAL-IP   PORT(S)                  AGE
service/coredns                   ClusterIP   10.96.0.2     <none>        53/UDP,53/TCP,9153/TCP   18h
service/traefik-ingress-service   ClusterIP   10.96.3.175   <none>        80/TCP,8080/TCP          17m

NAME                             DESIRED   CURRENT   READY   UP-TO-DATE   AVAILABLE   NODE SELECTOR   AGE
daemonset.apps/traefik-ingress   2         2         2       2            2           <none>          17m

NAME                      READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/coredns   1/1     1            1           18h

NAME                                 DESIRED   CURRENT   READY   AGE
replicaset.apps/coredns-6b6c4f9648   1         1         1       18h
```

------

如果pod没有起来没有起来请重启docker，原因是我上面测试了nodeport，防火墙规则改变了

```shell
[root@shkf6-243 ~]# systemctl restart docker
[root@shkf6-244 ~]# systemctl restart docker
```

------

## 6.浏览器访问

**访问 http://traefik.od.com/**
![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_81f6fbe8147a8cbae19f37dcaacc63b9_r.png)

# 第五章：K8S的GUI资源管理插件-仪表篇

## 1.部署kubenetes-dashborad

[dashboard官方Github](https://github.com/kubernetes/kubernetes/tree/master/cluster/addons/dashboard)

[dashboard下载地址](https://github.com/kubernetes/dashboard/releases)

### 1.准备dashboard镜像

运维主机SHKF6-245.host.com上：

```shell
[root@shkf6-245 ~]# docker pull sunrisenan/kubernetes-dashboard-amd64:v1.10.1

[root@shkf6-245 ~]# docker pull sunrisenan/kubernetes-dashboard-amd64:v1.8.3

[root@shkf6-245 ~]# docker images |grep dash
sunrisenan/kubernetes-dashboard-amd64   v1.10.1                    f9aed6605b81        11 months ago       122MB
sunrisenan/kubernetes-dashboard-amd64   v1.8.3                     0c60bcf89900        21 months ago       102MB

[root@shkf6-245 ~]# docker tag f9aed6605b81  harbor.od.com/public/kubernetes-dashboard-amd64:v1.10.1
[root@shkf6-245 ~]# docker push !$

[root@shkf6-245 ~]# docker tag 0c60bcf89900 harbor.od.com/public/kubernetes-dashboard-amd64:v1.8.3
[root@shkf6-245 ~]# docker push !$
```

### 2.准备配置清单

运维主机SHKF6-245.host.com上：

- 创建目录

  ```shell
  [root@shkf6-245 ~]# mkdir -p /data/k8s-yaml/dashboard && cd /data/k8s-yaml/dashboard
  ```

- rbac

```shell
[root@shkf6-245 dashboard]# cat rbac.yaml 
apiVersion: v1
kind: ServiceAccount
metadata:
  labels:
    k8s-app: kubernetes-dashboard
    addonmanager.kubernetes.io/mode: Reconcile
  name: kubernetes-dashboard-admin
  namespace: kube-system
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: kubernetes-dashboard-admin
  namespace: kube-system
  labels:
    k8s-app: kubernetes-dashboard
    addonmanager.kubernetes.io/mode: Reconcile
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-admin
subjects:
- kind: ServiceAccount
  name: kubernetes-dashboard-admin
  namespace: kube-system
```

- Deployment

```shell
[root@shkf6-245 dashboard]# cat dp.yaml 
apiVersion: apps/v1
kind: Deployment
metadata:
  name: kubernetes-dashboard
  namespace: kube-system
  labels:
    k8s-app: kubernetes-dashboard
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  selector:
    matchLabels:
      k8s-app: kubernetes-dashboard
  template:
    metadata:
      labels:
        k8s-app: kubernetes-dashboard
      annotations:
        scheduler.alpha.kubernetes.io/critical-pod: ''
    spec:
      priorityClassName: system-cluster-critical
      containers:
      - name: kubernetes-dashboard
        image: harbor.od.com/public/kubernetes-dashboard-amd64:v1.8.3
        resources:
          limits:
            cpu: 100m
            memory: 300Mi
          requests:
            cpu: 50m
            memory: 100Mi
        ports:
        - containerPort: 8443
          protocol: TCP
        args:
          # PLATFORM-SPECIFIC ARGS HERE
          - --auto-generate-certificates
        volumeMounts:
        - name: tmp-volume
          mountPath: /tmp
        livenessProbe:
          httpGet:
            scheme: HTTPS
            path: /
            port: 8443
          initialDelaySeconds: 30
          timeoutSeconds: 30
      volumes:
      - name: tmp-volume
        emptyDir: {}
      serviceAccountName: kubernetes-dashboard-admin
      tolerations:
      - key: "CriticalAddonsOnly"
        operator: "Exists"
```

- Service

```shell
[root@shkf6-245 dashboard]# cat svc.yaml 
apiVersion: v1
kind: Service
metadata:
  name: kubernetes-dashboard
  namespace: kube-system
  labels:
    k8s-app: kubernetes-dashboard
    kubernetes.io/cluster-service: "true"
    addonmanager.kubernetes.io/mode: Reconcile
spec:
  selector:
    k8s-app: kubernetes-dashboard
  ports:
  - port: 443
    targetPort: 8443
```

- ingress

```shell
[root@shkf6-245 dashboard]# cat ingress.yaml 
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: kubernetes-dashboard
  namespace: kube-system
  annotations:
    kubernetes.io/ingress.class: traefik
spec:
  rules:
  - host: dashboard.od.com
    http:
      paths:
      - backend:
          serviceName: kubernetes-dashboard
          servicePort: 443
```

### 3.依次执行创建

浏览器打开：`http://k8s-yaml.od.com/dashboard/`检查资源配置清单文件是否正确创建

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_738286ba6d1370be5fa09d275f5ac09d_r.png)

在SHKF6-243.host.com机器上：

```shell
[root@shkf6-243 ~]# kubectl apply -f http://k8s-yaml.od.com/dashboard/rbac.yaml
[root@shkf6-243 ~]# kubectl apply -f http://k8s-yaml.od.com/dashboard/dp.yaml
[root@shkf6-243 ~]# kubectl apply -f http://k8s-yaml.od.com/dashboard/svc.yaml
[root@shkf6-243 ~]# kubectl apply -f http://k8s-yaml.od.com/dashboard/ingress.yaml
```

## 2.解析域名

- 添加解析记录

```shell
[root@shkf6-241 ~]# cat /var/named/od.com.zone
$ORIGIN od.com.
$TTL 600    ; 10 minutes
@           IN SOA    dns.od.com. dnsadmin.od.com. (
                2019111208 ; serial    # 向后滚动+1
                10800      ; refresh (3 hours)
                900        ; retry (15 minutes)
                604800     ; expire (1 week)
                86400      ; minimum (1 day)
                )
                NS   dns.od.com.
$TTL 60    ; 1 minute
dns                A    192.168.6.241
harbor             A    192.168.6.245
k8s-yaml           A    192.168.6.245
traefik            A    192.168.6.66
dashboard          A    192.168.6.66   # 添加这条解析
```

- 重启named并检查

```shell
[root@shkf6-241 ~]# systemctl restart named

[root@shkf6-243 ~]# dig dashboard.od.com @10.96.0.2 +short
192.168.6.66
[root@shkf6-243 ~]# dig dashboard.od.com @192.168.6.241 +short
192.168.6.66
```

## 3.浏览器访问

浏览器访问：`http://dashboard.od.com`

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_56d7bba51541a6e914564efc19910251_r.png)

## 4.配置认证

### 1.签发证书

```shell
[root@shkf6-245 certs]# (umask 077; openssl genrsa -out dashboard.od.com.key 2048)
Generating RSA private key, 2048 bit long modulus
............................+++
........+++
e is 65537 (0x10001)
[root@shkf6-245 certs]# openssl req -new -key dashboard.od.com.key -out dashboard.od.com.csr -subj "/CN=dashboard.od.com/C=CN/ST=BJ/L=Beijing/O=OldboyEdu/OU=ops"
[root@shkf6-245 certs]# openssl x509 -req -in dashboard.od.com.csr -CA ca.pem -CAkey ca-key.pem -CAcreateserial -out dashboard.od.com.crt -days 3650
Signature ok
subject=/CN=dashboard.od.com/C=CN/ST=BJ/L=Beijing/O=OldboyEdu/OU=ops
Getting CA Private Key

[root@shkf6-245 certs]# ll dash*
-rw-r--r-- 1 root root 1196 Nov 27 12:52 dashboard.od.com.crt
-rw-r--r-- 1 root root 1005 Nov 27 12:52 dashboard.od.com.csr
-rw------- 1 root root 1679 Nov 27 12:52 dashboard.od.com.key
```

### 2.检查证书

```shell
[root@shkf6-245 certs]# cfssl-certinfo -cert dashboard.od.com.crt 
{
  "subject": {
    "common_name": "dashboard.od.com",
    "country": "CN",
    "organization": "OldboyEdu",
    "organizational_unit": "ops",
    "locality": "Beijing",
    "province": "BJ",
    "names": [
      "dashboard.od.com",
      "CN",
      "BJ",
      "Beijing",
      "OldboyEdu",
      "ops"
    ]
  },
  "issuer": {
    "common_name": "OldboyEdu",
    "country": "CN",
    "organization": "od",
    "organizational_unit": "ops",
    "locality": "beijing",
    "province": "beijing",
    "names": [
      "CN",
      "beijing",
      "beijing",
      "od",
      "ops",
      "OldboyEdu"
    ]
  },
  "serial_number": "11427294234507397728",
  "not_before": "2019-11-27T04:52:30Z",
  "not_after": "2029-11-24T04:52:30Z",
  "sigalg": "SHA256WithRSA",
  "authority_key_id": "",
  "subject_key_id": "",
  "pem": "-----BEGIN CERTIFICATE-----\nMIIDRTCCAi0CCQCeleeP167KYDANBgkqhkiG9w0BAQsFADBgMQswCQYDVQQGEwJD\nTjEQMA4GA1UECBMHYmVpamluZzEQMA4GA1UEBxMHYmVpamluZzELMAkGA1UEChMC\nb2QxDDAKBgNVBAsTA29wczESMBAGA1UEAxMJT2xkYm95RWR1MB4XDTE5MTEyNzA0\nNTIzMFoXDTI5MTEyNDA0NTIzMFowaTEZMBcGA1UEAwwQZGFzaGJvYXJkLm9kLmNv\nbTELMAkGA1UEBhMCQ04xCzAJBgNVBAgMAkJKMRAwDgYDVQQHDAdCZWlqaW5nMRIw\nEAYDVQQKDAlPbGRib3lFZHUxDDAKBgNVBAsMA29wczCCASIwDQYJKoZIhvcNAQEB\nBQADggEPADCCAQoCggEBALeeL9z8V3ysUqrAuT7lEKcF2bi0pSuwoWfFgfBtGmQa\nQtyNaOlyemEexeUOKaIRsNlw0fgcK6HyyLkaMFsVa7q+bpYBPKp4d7lTGU7mKJNG\nNcCU21G8WZYS4jVtd5IYSmmfNkCnzY7l71p1P+sAZNZ7ht3ocNh6jPcHLMpETLUU\nDaKHmT/5iAhxmgcE/V3DUnTawU9PXF2WnICL9xJtmyErBKF5KDIEjC1GVjC/ZLtT\nvEgbH57TYgrp4PeCEAQTtgNbVJuri4awaLpHkNz2iCTNlWpLaLmV1jT1NtChz6iw\n4lDfEgS6YgDh9ZhlB2YvnFSG2eq4tGm3MKorbuMq9S0CAwEAATANBgkqhkiG9w0B\nAQsFAAOCAQEAG6szkJDIvb0ge2R61eMBVe+CWHHSE6X4EOkiQCaCi3cs8h85ES63\nEdQ8/FZorqyZH6nJ/64OjiW1IosXRTFDGMRunqkJezzj9grYzUKfkdGxTv+55IxM\ngtH3P9wM1EeNwdJCpBq9xYaPzZdu0mzmd47BP7nuyrXzkMSecC/d+vrKngEnUXaZ\n9WK3eGnrGPmeW7z5j9uVsimzNlri8i8FNBTGCDx2sgJc16MtYfGhORwN4oVXCHiS\n4A/HVSYMUeR4kGxoX9RUbf8vylRsdEbKQ20M5cbWQAAH5LNig6jERRsuylEh4uJE\nubhEbfhePgZv+mkFQ6tsuIH/5ETSV4v/bg==\n-----END CERTIFICATE-----\n"
}
```

### 3.配置nginx

在shkf6-241和shkf6-242上：

- 拷贝证书

```shell
~]# mkdir /etc/nginx/conf.d/certs

~]# scp shkf6-245:/opt/certs/dashboard.od.com.crt /etc/nginx/certs/
~]# scp shkf6-245:/opt/certs/dashboard.od.com.key /etc/nginx/certs/
```

- 配置虚拟主机dashboard.od.com.conf，走https

```shell
[root@shkf6-241 ~]# cat /etc/nginx/conf.d/dashboard.od.com.conf
server {
    listen       80;
    server_name  dashboard.od.com;

    rewrite ^(.*)$ https://${server_name}$1 permanent;
}
server {
    listen       443 ssl;
    server_name  dashboard.od.com;

    ssl_certificate "certs/dashboard.od.com.crt";
    ssl_certificate_key "certs/dashboard.od.com.key";
    ssl_session_cache shared:SSL:1m;
    ssl_session_timeout  10m;
    ssl_ciphers HIGH:!aNULL:!MD5;
    ssl_prefer_server_ciphers on;

    location / {
        proxy_pass http://default_backend_traefik;
        proxy_set_header Host       $http_host;
        proxy_set_header x-forwarded-for $proxy_add_x_forwarded_for;
    }
}
```

- 重载nginx配置

```shell
 ~]# nginx -s reload
```

- 刷新页面检查

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_cc2c885819a8a0b49994178e10d055ba_r.png)

### 4.获取kubernetes-dashboard-admin-token

```shell
[root@shkf6-243 ~]# kubectl get secrets -n kube-system 
NAME                                     TYPE                                  DATA   AGE
coredns-token-r5s8r                      kubernetes.io/service-account-token   3      5d19h
default-token-689cg                      kubernetes.io/service-account-token   3      6d14h
kubernetes-dashboard-admin-token-w46s2   kubernetes.io/service-account-token   3      16h
kubernetes-dashboard-key-holder          Opaque                                2      14h
traefik-ingress-controller-token-nkfb8   kubernetes.io/service-account-token   3      5d1h
[root@shkf6-243 ~]# kubectl describe secret kubernetes-dashboard-admin-token-w46s2 -n kube-system |tail
Annotations:  kubernetes.io/service-account.name: kubernetes-dashboard-admin
              kubernetes.io/service-account.uid: 11fedd46-3591-4c15-b32d-5818e5aca7d8

Type:  kubernetes.io/service-account-token

Data
====
ca.crt:     1346 bytes
namespace:  11 bytes
token:      eyJhbGciOiJSUzI1NiIsImtpZCI6IiJ9.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJrdWJlLXN5c3RlbSIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VjcmV0Lm5hbWUiOiJrdWJlcm5ldGVzLWRhc2hib2FyZC1hZG1pbi10b2tlbi13NDZzMiIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50Lm5hbWUiOiJrdWJlcm5ldGVzLWRhc2hib2FyZC1hZG1pbiIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50LnVpZCI6IjExZmVkZDQ2LTM1OTEtNGMxNS1iMzJkLTU4MThlNWFjYTdkOCIsInN1YiI6InN5c3RlbTpzZXJ2aWNlYWNjb3VudDprdWJlLXN5c3RlbTprdWJlcm5ldGVzLWRhc2hib2FyZC1hZG1pbiJ9.HkVak9znUafeh4JTkzGRiH3uXVjcuHMTOsmz58xJy1intMn25ouC04KK7uplkAtd_IsA6FFo-Kkdqc3VKZ5u5xeymL2ccLaLiCXnlxAcVta5CuwyyO4AXeS8ss-BMKCAfeIldnqwJRPX2nzORJap3CTLU0Cswln8x8iXisA_gBuNVjiWzJ6tszMRi7vX1BM6rp6bompWfNR1xzBWifjsq8J4zhRYG9sVi9Ec3_BZUEfIc0ozFF91Jc5qCk2L04y8tHBauVuJo_ecgMdJfCDk7VKVnyF3Z-Fb8cELNugmeDlKYvv06YHPyvdxfdt99l6QpvuEetbMGAhh5hPOd9roVw
```

### 5.验证toke登录

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_ca78f8eda1b9f2313e4db8f6b8a5a94c_r.png)

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_14d5d69f3353aa8df4299059efd10383_r.png)

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_f7fded22ee708f59a805ef8051417953_r.png)

### 6.升级dashboard为v1.10.1

在shkf6-245.host.com上：

- 更改镜像地址：

  ```shell
  [root@shkf6-245 dashboard]# grep image dp.yaml 
        image: harbor.od.com/public/kubernetes-dashboard-amd64:v1.10.1
  ```

在shkf6-243.host.com上：

- 应用配置：

  ```shell
  [root@shkf6-243 ~]# kubectl apply -f http://k8s-yaml.od.com/dashboard/dp.yaml
  ```

### 7.dashboard 官方给的rbac-minimal

```shell
dashboard]# cat rbac-minimal.yaml 
apiVersion: v1
kind: ServiceAccount
metadata:
  labels:
    k8s-app: kubernetes-dashboard
    addonmanager.kubernetes.io/mode: Reconcile
  name: kubernetes-dashboard
  namespace: kube-system
---
kind: Role
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  labels:
    k8s-app: kubernetes-dashboard
    addonmanager.kubernetes.io/mode: Reconcile
  name: kubernetes-dashboard-minimal
  namespace: kube-system
rules:
  # Allow Dashboard to get, update and delete Dashboard exclusive secrets.
- apiGroups: [""]
  resources: ["secrets"]
  resourceNames: ["kubernetes-dashboard-key-holder", "kubernetes-dashboard-certs"]
  verbs: ["get", "update", "delete"]
  # Allow Dashboard to get and update 'kubernetes-dashboard-settings' config map.
- apiGroups: [""]
  resources: ["configmaps"]
  resourceNames: ["kubernetes-dashboard-settings"]
  verbs: ["get", "update"]
  # Allow Dashboard to get metrics from heapster.
- apiGroups: [""]
  resources: ["services"]
  resourceNames: ["heapster"]
  verbs: ["proxy"]
- apiGroups: [""]
  resources: ["services/proxy"]
  resourceNames: ["heapster", "http:heapster:", "https:heapster:"]
  verbs: ["get"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: kubernetes-dashboard-minimal
  namespace: kube-system
  labels:
    k8s-app: kubernetes-dashboard
    addonmanager.kubernetes.io/mode: Reconcile
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: kubernetes-dashboard-minimal
subjects:
- kind: ServiceAccount
  name: kubernetes-dashboard
  namespace: kube-system
```

## 5.部署heapster

[heapster官方github地址](https://github.com/kubernetes-retired/heapster)

### 1.准备heapster镜像

```shell
[root@shkf6-245 ~]# docker pull sunrisenan/heapster:v1.5.4

[root@shkf6-245 ~]# docker images|grep heapster
sunrisenan/heapster                               v1.5.4                     c359b95ad38b        9 months ago        136MB

[root@shkf6-245 ~]# docker tag c359b95ad38b harbor.od.com/public/heapster:v1.5.4
[root@shkf6-245 ~]# docker push !$
```

### 2.准备资源配置清单

- rbac.yaml

```shell
[root@shkf6-245 ~]# cat /data/k8s-yaml/heapster/rbac.yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: heapster
  namespace: kube-system
---
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1beta1
metadata:
  name: heapster
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: system:heapster
subjects:
- kind: ServiceAccount
  name: heapster
  namespace: kube-system
```

- Deployment

```shell
[root@shkf6-245 ~]# cat /data/k8s-yaml/heapster/dp.yaml 
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: heapster
  namespace: kube-system
spec:
  replicas: 1
  template:
    metadata:
      labels:
        task: monitoring
        k8s-app: heapster
    spec:
      serviceAccountName: heapster
      containers:
      - name: heapster
        image: harbor.od.com/public/heapster:v1.5.4
        imagePullPolicy: IfNotPresent
        command:
        - /opt/bitnami/heapster/bin/heapster
        - --source=kubernetes:https://kubernetes.default
```

- service

```shell
[root@shkf6-245 ~]# cat /data/k8s-yaml/heapster/svc.yaml 
apiVersion: v1
kind: Service
metadata:
  labels:
    task: monitoring
    # For use as a Cluster add-on (https://github.com/kubernetes/kubernetes/tree/master/cluster/addons)
    # If you are NOT using this as an addon, you should comment out this line.
    kubernetes.io/cluster-service: 'true'
    kubernetes.io/name: Heapster
  name: heapster
  namespace: kube-system
spec:
  ports:
  - port: 80
    targetPort: 8082
  selector:
    k8s-app: heapster
```

### 3.应用资源配置清单

```shell
[root@shkf6-243 ~]# kubectl apply -f http://k8s-yaml.od.com/heapster/rbac.yaml

[root@shkf6-243 ~]# kubectl apply -f http://k8s-yaml.od.com/heapster/dp.yaml

[root@shkf6-243 ~]# kubectl apply -f http://k8s-yaml.od.com/heapster/svc.yaml
```

### 4.重启dashboard(可以不重启)

```shell
[root@shkf6-243 ~]# kubectl delete -f http://k8s-yaml.od.com/dashboard/dp.yaml

[root@shkf6-243 ~]# kubectl apply -f http://k8s-yaml.od.com/dashboard/dp.yaml
```

### 5.检查

- 主机检查

```shell
[root@shkf6-243 ~]# kubectl top node
NAME                 CPU(cores)   CPU%   MEMORY(bytes)   MEMORY%   
shkf6-243.host.com   149m         3%     3643Mi          47%       
shkf6-244.host.com   130m         3%     3300Mi          42%       

[root@shkf6-243 ~]# kubectl top pod -n kube-public 
NAME                        CPU(cores)   MEMORY(bytes)   
nginx-dp-5dfc689474-dm555   0m           10Mi 
```

- 浏览器检查

![null](http://www.sunrisenan.com/uploads/kubernetes/images/m_6e4ce39904dc61713b70c334143ec1b4_r.png)

# 第六章：k8s平滑升级

第一步：观察哪台机器pod少

```shell
[root@shkf6-243 src]# kubectl get nodes
NAME                 STATUS   ROLES         AGE     VERSION
shkf6-243.host.com   Ready    master,node   6d16h   v1.15.2
shkf6-244.host.com   Ready    master,node   6d16h   v1.15.2

[root@shkf6-243 src]# kubectl get pods -n kube-system -o wide
NAME                                    READY   STATUS    RESTARTS   AGE     IP            NODE                 NOMINATED NODE   READINESS GATES
coredns-6b6c4f9648-x5zvz                1/1     Running   0          5d22h   172.6.244.3   shkf6-244.host.com   <none>           <none>
heapster-b5b9f794-s2vj9                 1/1     Running   0          76m     172.6.243.4   shkf6-243.host.com   <none>           <none>
kubernetes-dashboard-5dbdd9bdd7-qlk52   1/1     Running   0          62m     172.6.244.4   shkf6-244.host.com   <none>           <none>
traefik-ingress-bhhkv                   1/1     Running   0          5d4h    172.6.244.2   shkf6-244.host.com   <none>           <none>
traefik-ingress-mm2ds                   1/1     Running   0          5d4h    172.6.243.2   shkf6-243.host.com   <none>           <none>
```

第二步：在负载均衡了禁用7层和4层

```shell
略
```

第三步：摘除node节点

```shell
[root@shkf6-243 src]# kubectl delete node shkf6-243.host.com
node "shkf6-243.host.com" deleted
[root@shkf6-243 src]# kubectl get node
NAME                 STATUS   ROLES         AGE     VERSION
shkf6-244.host.com   Ready    master,node   6d17h   v1.15.2
```

第四步：观察运行POD情况

```shell
[root@shkf6-243 src]# kubectl get pods -n kube-system -o wide
NAME                                    READY   STATUS    RESTARTS   AGE     IP            NODE                 NOMINATED NODE   READINESS GATES
coredns-6b6c4f9648-x5zvz                1/1     Running   0          5d22h   172.6.244.3   shkf6-244.host.com   <none>           <none>
heapster-b5b9f794-dlt2z                 1/1     Running   0          15s     172.6.244.6   shkf6-244.host.com   <none>           <none>
kubernetes-dashboard-5dbdd9bdd7-qlk52   1/1     Running   0          64m     172.6.244.4   shkf6-244.host.com   <none>           <none>
traefik-ingress-bhhkv                   1/1     Running   0          5d4h    172.6.244.2   shkf6-244.host.com   <none>           <none>
```

第五步：检查dns

```shell
[root@shkf6-243 src]# dig -t A kubernetes.default.svc.cluster.local @10.96.0.2 +short
10.96.0.1
```

第六步：开始升级

```shell
[root@shkf6-243 ~]# cd /opt/src/
[root@shkf6-243 src]# wget http://down.sunrisenan.com/k8s/kubernetes/kubernetes-server-linux-amd64-v1.15.4.tar.gz
[root@shkf6-243 src]# tar xf kubernetes-server-linux-amd64-v1.15.4.tar.gz
[root@shkf6-243 src]# mv kubernetes /opt/kubernetes-v1.15.4
[root@shkf6-243 src]# cd /opt/
[root@shkf6-243 opt]# rm -f kubernetes
[root@shkf6-243 opt]# ln -s /opt/kubernetes-v1.15.4 kubernetes
[root@shkf6-243 opt]# cd /opt/kubernetes
[root@shkf6-243 kubernetes]# rm -f kubernetes-src.tar.gz 
[root@shkf6-243 kubernetes]# cd server/bin/
[root@shkf6-243 bin]# rm -f *.tar
[root@shkf6-243 bin]# rm -f *tag

[root@shkf6-243 bin]# cp -r /opt/kubernetes-v1.15.2/server/bin/cert .
[root@shkf6-243 bin]# cp -r /opt/kubernetes-v1.15.2/server/bin/conf .
[root@shkf6-243 bin]# cp -r /opt/kubernetes-v1.15.2/server/bin/*.sh .

[root@shkf6-243 bin]# systemctl restart supervisord.service 

[root@shkf6-243 bin]# kubectl get node
NAME                 STATUS   ROLES         AGE     VERSION
shkf6-243.host.com   Ready    <none>        16s     v1.15.4
shkf6-244.host.com   Ready    master,node   6d17h   v1.15.2
```

升级另一台：

```shell
[root@shkf6-244 src]# kubectl get node
NAME                 STATUS   ROLES         AGE     VERSION
shkf6-243.host.com   Ready    <none>        95s     v1.15.4
shkf6-244.host.com   Ready    master,node   6d17h   v1.15.2

[root@shkf6-244 src]# kubectl delete node shkf6-244.host.com
node "shkf6-244.host.com" deleted

[root@shkf6-244 src]# kubectl get node
NAME                 STATUS   ROLES    AGE    VERSION
shkf6-243.host.com   Ready    <none>   3m2s   v1.15.4

[root@shkf6-244 src]# kubectl get pods -n kube-system -o wide
NAME                                    READY   STATUS    RESTARTS   AGE     IP            NODE                 NOMINATED NODE   READINESS GATES
coredns-6b6c4f9648-bxqcp                1/1     Running   0          20s     172.6.243.3   shkf6-243.host.com   <none>           <none>
heapster-b5b9f794-hjx74                 1/1     Running   0          20s     172.6.243.4   shkf6-243.host.com   <none>           <none>
kubernetes-dashboard-5dbdd9bdd7-gj6vc   1/1     Running   0          20s     172.6.243.5   shkf6-243.host.com   <none>           <none>
traefik-ingress-4hl97                   1/1     Running   0          3m22s   172.6.243.2   shkf6-243.host.com   <none>           <none>

[root@shkf6-244 src]# dig -t A kubernetes.default.svc.cluster.local @10.96.0.2 +short
10.96.0.1
```