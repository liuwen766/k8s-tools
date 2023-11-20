本文地址：https://blog.csdn.net/hancoder/article/details/113821646

其他系列地址：[https://blog.csdn.net/hancoder/article/details/113821646](https://blog.csdn.net/hancoder/article/details/113821646)

主要是node-sass版本问题，他的版本看和node版本对应

### 0 安装

> 可以去这里找到v12的版本。（不要用12.0，可以用12.1）
>
> https://npm.taobao.org/mirrors/node/

`NPM`是随同`NodeJS`一起安装的包管理工具。JavaScript-NPM类似于java-Maven。

命令行输入`node -v` 检查配置好了，配置npm的镜像仓库地址，再执

```bash
node -v
npm config set registry http://registry.npm.taobao.org/
```



### 1 版本信息

去https://github.com/sass/npnode-sass   或者  https://github.com/sass/node-sass/releases  都可以看到node和node-sass对应的版本信息

| NodeJS  | Supported node-sass version | Node Module |
| ------- | --------------------------- | ----------- |
| Node 15 | 5.0+                        | 88          |
| Node 14 | 4.14+                       | 83          |
| Node 13 | 4.13+, <5.0                 | 79          |
| Node 12 | 4.12+                       | 72          |
| Node 11 | 4.10+, <5.0                 | 67          |
| Node 10 | 4.9+                        | 64          |
| Node 8  | 4.5.3+, <5.0                | 57          |
| Node <8 | <5.0                        | <57         |

从https://github.com/sass/node-sass/releases?after=v4.12.0     也能看出4.9.2最多只支持到node10。

### 2 正常命令

- 镜像加速：
  - npm config set registry http://registry.npm.taobao.org/

在项目项目下执行

```bash
# 指定node-sass版本，这句等价于修改package.json文件了。
# 注意不要指定4.9.2了
npm install  node-sass@4.14
# 等价于npm i node-sass --sass_binary_path=
# 等价于修改package.json

#安装其他依赖：
npm install

# 启动项目：
npm run dev
```

然后直接就可以执行起来了。如果没验证码什么的，是因为java项目没有启动

#### 2.1 卸载残留问题

你可能按着评论区或视频的东西安装试过了，要注意执行我上面代码时要先清依赖残留，否则安装不上

```bash
npm rebuild node-sass
npm uninstall node-sass
```

上面步骤等价于直接删除node_modules，只不过node_modules是删了全部下载的依赖包

运行成功案例：

```bash
PS F:\renren-fast-vue> npm install
npm WARN optional SKIPPING OPTIONAL DEPENDENCY: fsevents@1.2.9 (node_modules\fsevents):
npm WARN notsup SKIPPING OPTIONAL DEPENDENCY: Unsupported platform for fsevents@1.2.9: wanted {"os":"darwin","arch":"any"} (current: {"os":"win32","arch":"x64"})

up to date in 10.975s

2 packages are looking for funding
  run `npm fund` for details
  
 # warn先不用管
 # 运行
 npm run dev
 #  I  Your application is running here: http://localhost:8001
 

浏览器输入localhost:8001 就可以看到内容了，登录admin、 admin

我们还可以看到VS和IDEA联动了
```

### 3 视频评论区说法的错误

从上表看，如果你是node12，你得安装node-sass4.12版本以上的，我不知道评论区为什么4.9.2就可以了，应该是碰巧兼容了。

另外评论区写的镜像地址https://npm.taobao.org/mirrors/node-sass/  ，你在浏览器中输入后，你会发现他有4.9.4和5.0.0，却没有4.14+。

还有一种可能是，`npm i `这种语法在镜像中==自动帮你检测与node版本最匹配的npm依赖包版本号==（有这种功能），所以我个人认为你改了4.9.2也没有用。他根本不按照你写的来，自动匹配到了你指定的镜像里的版本号，他想要4.12+，但是你指定的镜像地址里没有，所以他就升级到使用5.0.0，这点是可以验证的，你此时看你项目里写的版本号，你明明写的4.9.2，他自动变成了^5.0.0。所以你写9.0还是9.2都无所谓，因为根本不按你的来（注意安装过的话先执行后文卸载操作）

综上，其实我不建议你写npm -i加链接的方式，原因如下：

安装了5.0版本后，你可能按此方法也能正常使用了，但安装时还是有个警告：

> npm install
> npm WARN `sass-loader`@6.0.6 requires a peer of node-sass@^4.0.0 but none is installed. You must install peer dependencies yourself.
>
> 这个应该说`sass-loader`@6.0.6需要低版本的sass，我除以为是升级`sass-loader`@6.0.6即可，但是升级后报错其他，可能是代码跟版本有关系，所以这是个很矛盾的配置
>
> 所以我们别用i的方式让他升级到版本5了，我们直接@4.14，这个版本兼容node12和node14

#### 3.1 python问题？

视频评论区有人说得先安装python，因为我本身已经安装了anaconda，所以不确定影响与否。如果失败可以尝试安装python3.0以上版本，并配置全局变量。（根据人人开发文档介绍:  node使用8.x版本,无需做任何改动,如果报错,执行npm rebuild node-sass之后,然后重新install即可,无需安装python）

#### 3.2 npm i方式安装？



评论区提供了这么一种解决方式，这个我上面说过了，镜像里并没有4.14版本，你使用了版本5的又造成sass-loader不匹配，升级sass-loader代码又报错。所以这里不适合用这个。

- 单独安装node-sass： npm i node-sass --sass_binary_site=https://npm.taobao.org/mirrors/node-sass/



- node和node-sass有版本对应关系。node12对应4.14的node-sass
- i是install的缩写



#### 3.3 chromedriver问题？

按理说没有这个问题，不知道谁会遇到这个问题，但究其原因还是package.json中有这个起来，没有正确导入，自己提前导入一下即可

- npm install chromedriver --chromedriver_cdnurl=http://cdn.npm.taobao.org/dist/chromedriver

### 知识点：

#### 1、npm i是什么

```sh
npm -h

C:\Users\HAN>npm i -h

npm install (with no args, in package dir)
npm install [<@scope>/]<pkg>
npm install [<@scope>/]<pkg>@<tag>
npm install [<@scope>/]<pkg>@<version>
npm install [<@scope>/]<pkg>@<version range>
npm install <alias>@npm:<name>
npm install <folder>
npm install <tarball file>
npm install <tarball url>
npm install <git:// url>
npm install <github username>/<github project>

aliases: i, isntall, add
common options: [--save-prod|--save-dev|--save-optional] [--save-exact] [--no-save]
```

你会发现npm是npm install的简写，但也有区别：

\1. 用npm i安装的模块无法用npm uninstall删除，用npm uninstall i才卸载掉 （这个说法我比较怀疑）
\2. npm i会帮助检测与当前node版本最匹配的npm包版本号，并匹配出来相互依赖的npm包应该提升的版本号 
\3. 部分npm包在当前node版本下无法使用，必须使用建议版本 
\4. 安装报错时intall肯定会出现npm-debug.log 文件，npm i不一定

> 我试过用npm i代替npm install，结果是4.9.0的下载下来了，但是报错环境不匹配





#### package.json文件的作用

https://blog.csdn.net/csm0912/article/details/90264026

#### 阿里云镜像网站

npm：https://developer.aliyun.com/mirror/NPM

https://developer.aliyun.com/mirror/

### 关系

```
npm install sass-loader node-sass webpack --save-dev
```

必须安装node-sass才能安装sass-loader

F12有如下预警

```
DevTools failed to load SourceMap: Could not load content for chrome-extension://kjacjjdnoddnpbbcjilcajfhhbdhkpgk/js/browser-polyfill.js.map: HTTP error: status code 404, net::ERR_UNKNOWN_URL_SCHEME
```

### 笔记不易：

离线笔记均为markdown格式，图片也是云图，10多篇笔记20W字，压缩包仅500k，推荐使用typora阅读。也可以自己导入有道云笔记等软件中

阿里云图床现在**每周得几十元充值**，都要自己往里搭了，麻烦不要散播与转发

![](https://i0.hdslb.com/bfs/album/ff3fb7e24f05c6a850ede4b1f3acc54312c3b0c6.png)

打赏后请主动发支付信息到邮箱  553736044@qq.com  ，上班期间很容易忽略收账信息，邮箱回邮基本秒回

禁止转载发布，禁止散播，若发现大量散播，将对本系统文章图床进行重置处理。

技术人就该干点技术人该干的事



如果帮到了你，留下赞吧，谢谢支持