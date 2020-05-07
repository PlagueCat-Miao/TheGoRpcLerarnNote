# TheGoRpcLerarnNote
主要学习net/rpc和kite框架

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
## get-start

## 墓碑
### 架构

### 日志

## 雷区

### 1.因麻烦而写的民工三连
git add . ; git commit -m "快捷上传"; git push origin master

[使用 Go Modules（模块）进行依赖项迁移](https://studygolang.com/articles/23133?fr=sidebar)
开始使用依赖关系管理
若要转换已使用依赖关系管理工具的项目，请运行以下命令:
go: creating new go.mod: module github.com/my/project
go: copying requirements from Godeps/Godeps.json

### go.mod replace
版本号哪里既可以提供 vx.y.z
也可以输入 分支名  / tag / commit 号 一般他们为网页的最后一个尾缀
go get 的过程中会自动找到正确的 版本编号 并更新 go.mod