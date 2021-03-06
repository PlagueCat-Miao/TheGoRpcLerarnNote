# TheGoRpcLerarnNote
主要学习net/rpc和kite框架 \
[流行的rpc框架性能测试对比](https://blog.csdn.net/quuqu/article/details/79304614) :kite 不是开源的 改学 thrift kit

## 下载方式

    go get github.com/PlagueCat-Miao/TheGoRpcLerarnNote  
    git clone https://github.com/PlagueCat-Miao/TheGoRpcLerarnNote.git
    
    cd $GOPATH/src/github.com/PlagueCat-Miao/TheGoRpcLerarnNote;git pull origin master;git diff  

## 上载方式

git add .   \
git commit -m "XX"   \
git statue   \
git diff   \
git push origin master   \
git add . ; git commit -m "快捷上传"; git push origin master

     cd $GOPATH/src/github.com/PlagueCat-Miao/TheGoRpcLerarnNote;git add . ;git commit -m "快速上传"; git push origin master 


## 用户信息
```
#### 用户名
    PlagueCat-Miao
#### 自动输入密码
    git config  credential.helper store
```
## install 
MaC:
    [Mac install Thrift](https://blog.csdn.net/liuxinmingcode/article/details/45567241)
    
liunx:

    sudo apt-get install automake bison flex g++ git libboost-all-dev libevent-dev libssl-dev libtool make pkg-config
    # 下载安装包
    http://archive.apache.org/dist/thrift/ 
    http://www.apache.org/dyn/closer.cgi?path=/thrift/0.9.1/thrift-0.9.1.tar.gz
    curl -# -O http://archive.apache.org/dist/thrift/0.9.2/thrift-0.9.2.tar.gz
    wget http://archive.apache.org/dist/thrift/0.9.2/thrift-0.9.2.tar.gz 
    # 下载完成后开始解压
    tar zxvf thrift-0.9.2.tar.gz
    # 进入文件夹
    cd thrift-0.9.2
    # 初始化配置
    ./configure
    # 开始构建
    make
    # 一定要用sudo，本机编译
    sudo make install
    # 查看thrift是否生效
    thrift --version
    # 显示：Thrift version 0.9.2

windows:
[Apache Thrift 在Windows下的安装与开发](https://blog.csdn.net/colouroo/article/details/38588297)

## get-start
[Thrift RPC 使用指南实战(附golang&PHP代码)](https://blog.csdn.net/liuxinmingcode/article/details/45696237)
[thrift go语言官方例子](https://www.jianshu.com/p/27f721c13c5d)
git clone https://github.com/apache/thrift.git


## 墓碑
### go
#### thift
##### demo教程
-[Go thrift使用举例](https://blog.csdn.net/lanyang123456/article/details/80372977)
##### 参数总结
-[golang thrift 总结一下网络上的一些坑](https://www.cnblogs.com/ka200812/p/5865213.html)

 
## 雷区
### 1.因麻烦而写的民工三连
git add . ; git commit -m "快捷上传"; git push origin master

[使用 Go Modules（模块）进行依赖项迁移](https://studygolang.com/articles/23133?fr=sidebar)
开始使用依赖关系管理
若要转换已使用依赖关系管理工具的项目，请运行以下命令:
go: creating new go.mod: module github.com/my/project
go: copying requirements from Godeps/Godeps.json

### 2. go.mod replace
replace 时如何输入版本号
 - 可以输入 分支名  / tag / commit 号:一般他们为网页的最后一个尾缀；go get 的过程中会自动找到正确的版本编号,并更新 go.mod
 - 输入master 如:`https://github.com/koding/kite`
     - 就可以输入 `xxxx => github.com/koding/kite master`
     - go get 后go.mod更新为  `github.com/koding/kite v0.0.0-20180710021347-baa1a54919e3`
 - 完整commit号 如：`https://github.com/koding/kite/commit/59a699eb5ebad76fecf2b375b4599650eaf936f3`
    - 就可以输入 `xxxx => github.com/koding/kite 59a699eb5ebad76fecf2b375b4599650eaf936f3`
    - go get 后go.mod更新为  `github.com/koding/kite v0.0.0-20180710021347-baa1a54919e3`
 - 简码commit 如 ![github 网站截图](./BasicPractice/pic/2.png)
    - 就可以输入 `xxxx => github.com/koding/kite baa1a54`
    - go get 后go.mod更新为 `github.com/koding/kite v0.0.0-20180710021347-baa1a54919e3`
 - branch  如：`https://github.com/koding/kite/tree/fix-etcd-build`
    -  就可以输入 `xxxx => github.com/koding/kite fix-etcd-build`
    -  也可以commit 简码 `xxxx => github.com/koding/kite 59a699e`
    -  go get 后go.mod更新为 `github.com/koding/kite v0.0.0-20150902091132-59a699eb5eba`
 - version 如 `https://github.com/gorilla/websocket/releases/tag/v1.4.0` 
    -  就可以输入 `xxxx => github.com/gorilla/websocket v1.4.0`
    -  也可以是tag 下的 commit 简码 ![github tag 网站截图](./BasicPractice/pic/3.png) 如`xxxx => github.com/gorilla/websocket 66b9c49`
    -  go get 后go.mod更新为 `github.com/gorilla/websocket v1.4.0`
 ### 3. go.mod replace 特例
 1. 某些项目发布 vx 高级版本如：`https://github.com/igm/sockjs-go/tree/master/v3`
    -  确定发布版本
      - 存在vx文件夹，其下有独立的go.mod，且module含有vx如:`module github.com/igm/sockjs-go/v3`
      - or 项目打有标签/分支 如`v3` 其下go.mod的module含有vx，如:`module github.com/igm/sockjs-go/v3`
    -  如果要 import `github.com/igm/sockjs-go/v3`
    -  如果要 replace(commit简码) `XXX => github.com/igm/sockjs-go/v3 50a6500` 注意这个页面的go.mod 应含有`module github.com/igm/sockjs-go/v3`如：
    ![github tag 网站截图](./BasicPractice/pic/4.png)
    
 2. `code.google.com/` 失败
    - 其中`p/go.net`路径下 已被转移`golang.org/x/net`
    - 如`code.google.com/p/go.net/websocket`
       - go.mod 不可改为`golang.org/x/net/websocket` 因为没有go.mod
       - go get `golang.org/x/net`
       - ~~import 改为`golang.org/x/net/websocket`~~
  3. 参考网站
    - [go mod 使用旧版本 版本号指定](https://blog.csdn.net/gs80140/article/details/95320215)
    - [go1.13 mod 实践和常见问题](https://blog.csdn.net/qq_23109825/article/details/103604685)
 ### 4. go.mod 配置
 对于某个项目，如果需要 replace 、require 只需要在此目录下的go.mod 设计。整个go build过程都生效，
  - 当然你也可以去pkg 里去改。
  - 当然你子项目GET不到，不会在主项目里更新require，但在主go.mod里replace生效
  - //是注释 会因为 `go mod tidy` 而被清除 
 ### 5. git clone 是克隆整个仓库
  你会得到所有分支 直接 `git branch` 转分支
   
 ### 6. thrift 版本与GO不匹配
 - 一句话 官方大法好[thrift 0.13.0](http://thrift.apache.org/tutorial/go)
   - 里面实例中`tutorial` `shared` 是gen-go下的路径
 - [thrift解析](https://www.cnblogs.com/ka200812/p/5865213.html)
  
## 扩展阅读
### rpc
  - [流行的rpc框架性能测试对比](https://blog.csdn.net/quuqu/article/details/79304614)
  - [历史-晁岳攀---基于go的 rpc框架实践](https://blog.csdn.net/RA681t58CJxsgCkJ31/article/details/82455716)
###  redis
  - [谈谈陌陌争霸在数据库方面踩过的坑( Redis 篇)](https://blog.codingnow.com/2014/03/mmzb_redis.html)
  - [Redis与数据库数据同步解决方案](https://blog.csdn.net/tennysonsky/article/details/78205453)
  - [利用redis-sentinel+consul实现redis高可用](https://blog.csdn.net/weixin_33937499/article/details/85087856)
  - [redis如何实现主从数据的同步](https://www.cnblogs.com/lice-blog/p/11616364.html)
  - [Redis和数据库同步问题](https://www.cnblogs.com/George1994/p/10601244.html)
### Consul 
[Consul 快速入门](https://www.jianshu.com/p/7d20dc58c9fc)
### 中间件
  - [什么是中间件？常见中间件有哪些？](http://c.biancheng.net/view/3860.html)
  - [对新手的初级入门，什么是中间件。](https://blog.csdn.net/intermediat/article/details/96178217)
  - [微服务（Microservices）和服务网格（Service Mesh）架构概念整理](https://www.cnblogs.com/xishuai/p/microservices-and-service-mesh.html)