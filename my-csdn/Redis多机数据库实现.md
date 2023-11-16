# Redis多机数据库

## 一、复制



## 二、哨兵

![](D:\MyGitHub\my-tools\my-csdn\Redis sentinel结构.jpg)

```shell
# 放行所有 IP 限制
bind 0.0.0.0
# 进程端口号
port 26379
# 后台启动
daemonize yes
# 日志记录文件
logfile "/usr/local/redis/log/sentinel.log"
# 进程编号记录文件
pidfile /var/run/sentinel.pid
# 指示 Sentinel 去监视一个名为 mymaster 的主服务器 2为制裁权重值
sentinel monitor mymaster 192.168.10.101 6379 2
# 访问主节点的密码【如果设置密码，有必要统一密码的设置】
sentinel auth-pass mymaster 123456
# Sentinel 认为服务器已经断线所需的毫秒数
sentinel down-after-milliseconds mymaster 10000
# 若 Sentinel 在该配置值内未能完成 failover 操作，则认为本次 failover 失败
sentinel failover-timeout mymaster 180000
```

## 三、集群
