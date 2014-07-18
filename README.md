omi-blog
==========

omi's blog use beego framework

thanks ulricqin

## install 
centos安装使用如下命令：

```
sudo yum install go  
sudo yum install git  
sudo yum install mariadb 
sudo yum install modify
mkdir ~/go
export GOPATH=~/go
``` 


Ubuntu安装使用如下命令：

```
sudo apt-get install python-software-properties  
sudo add-apt-repository ppa:gophers/go  
sudo apt-get update  
sudo apt-get install golang-stable git-core mercurial  
sudo apt-get install mariadb
mkdir ~/go
export GOPATH=~/go
```

执行完上面的命令之后环境变量会设置好，这样就可以直接使用了。

## make

```
mkdir -p $GOPATH/src/github.com/darknessomi
cd $GOPATH/src/github.com/darknessomi
git clone https://github.com/darknessomi/omi-blog.git
go get github.com/darknessomi/omi-blog/...
cd omi-blog && modify conf/app.conf
bee run
```
## mysql

```
create database omi_blog;
use omi_blog;
source $GOPATH/src/github.com/darknessomi/omi-Blog/db.sql;
```




## admin 

```
url: /me
username: omi
password: md5
```