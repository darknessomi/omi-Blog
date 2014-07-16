omi-blog
==========

omi's blog use beego framework

thanks ulricqin

## install

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
source /Users/omi/go/src/omi-Blog/db.sql;
```




## admin 

```
url: /me
username: omi
password: a
```