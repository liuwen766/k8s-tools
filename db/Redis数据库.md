# Redis数据库

## 一、服务端与客户端

```c
// 服务端状态结构体
struct redisServer{
    //...
    //保存所有的数据库的数组
    redisDb *db;
    //数据库数量【默认16个】
    int dbnum; 
    //...
}

// 客户端状态结构体
struct redisClient{
    //...
    //记录当前使用的数据库
    redisDb *db;
    //...
}
```

> 注：通过修改redisClient.db指针，让它指向服务器中的某一个数据库，从而实现切换目标数据库的功能——select命令的实现原理。
>
> tips：为了避免对数据库进行误操作，在执行命令前，最好先执行一个select命令。

## 二、数据库键空间

Redis数据库主要是由dict和expires两个字典构成，其中dict字典负责保存键值对，而expires字典则负责保存键的过期时间。

1、键空间的CRUD操作

```c
// 数据库结构体
struct redisDb{
    //...
    //数据库键空间，它保存着数据库中的所有键值对
    dict *dict;
    //...
    //过期字典，它保存着键的过期时间
    dict *expires;
    //...
}


```



2、其它键空间操作



3、键空间的维护操作

当使用Redis命令对数据库进行读写时，服务器不仅会对键空间执行指定的读写操作，还会执行一些额外的维护操作，其中包括：

- hit/miss次数：使用info stats查看。
- lru
- dirty
- 通知：



## 三、数据库过期字典

1、设置过期时间

expire/pexpire

expireat/pexpireat

ttl

time

2、保存过期时间

》》》》》数据库键空间图

3、计算剩余时间

4、过期键判定

5、过期键删除策略



## 四、AOF、RDB和复制功能对过期键的处理

1、