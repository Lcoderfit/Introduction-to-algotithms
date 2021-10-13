# 一、Go环境变量相关
## 1.1 Go env命令
* 查看env命令的帮助信息
go help env

* 设置go环境变量(-w表示写的意思，go1.13及以上版本支持)
go env -w GOPROXY=https://goproxy.cn,direct

* 将go环境变量还原为默认设置
go env -u GOPROXY

* 将go环境变量按照json的格式输出
go env -json
go env -json GOPROXY

* 显示所有环境变量
go env

* 显示单个环境变量
go env GOPROXY

## 1.2 GOPROXY
* GOPROXY用于设置拉取第三方库时的代理，一般设置为 https://goproxy.cn,direct
注意：GOPROXY可以设置多个代理，拉取代码时候直接从设置的代理获取，例如：proxy1,proxy2,...,direct, 
当拉取第三方库时会按从左到右的顺序逐个尝试代理，直到成功获取为止，
direct表示在所有设置的其他代理无效时，直接从第三方库的源地址拉取代码，例如：
设置GOPROXY为 https://goproxy.cn,direct， go get -v github.com/Lcoderfit/GopherUtils
会先从https://goproxy.cn这个代理拉取代码，如果获取不到，则直接从github.com/Lcoderfit/GopherUtils链接获取

## 1.3 go modules \[ˈmɑdʒulz\]
* go1.11开始引入go module
* go1.12开始支持go module
* go1.14宣布go module可用于生产环境
* go1.16默认开启go module

* go module设置
go env GO111MODULE=on
on表示开启go module
off表示关闭go module
auto是1.12到1.15版本时的默认值，当遇到以下两种情况时会启用go module
    1.当前目录不在GOPATH/src中且该目录包含go.mod文件
    2.当前文件在包含go.mod的目录下面

当启用go module时，会把第三方依赖下载到GOPATH/pkg/mod目录中，go install的结果会存放到GOPATH/bin目录中

## 1.4 GOROOT/GOPATH/GOBIN
* GOROOT 表示go的安装路径
* GOPATH 表示go项目路径，开启go module之后项目不用存放到该目录下的src目录中
bin/ go项目的执行文件 (go install命令会将生成的.exe存放到GOPATH/bin目录下)
pkg/ go第三方库存放的路径(开启go module之后会存放到 GOPATH/pkg/mod目录下)
src/ go项目存放路径，开启go module后可以不放在该目录下

* GOBIN
默认为 GOPATH/bin

# 二、go命令
## 2.1 go build
* 将main.go编译成当前目录下的m.exe可执行文件
go build -o m.exe main.go

* go help build
???

## 2.2 go run
??

## 2.3 go install
* 先生成可执行文件，然后移动到GOBIN目录下
?????

# 三、更新go版本
## 3.1 Linux
* 使用uname命令查看Ubuntu的系统架构(x86/x86-64/ARM64)
可以得知当前的Linux系统是x86_64的架构
```text
root@iZuf6a99qngm2w94o1aqz3Z:/home/SoftWareForCoding# uname -a
Linux iZuf6a99qngm2w94o1aqz3Z 4.4.0-117-generic #141-Ubuntu SMP Tue Mar 13 11:58:07 UTC 2018 x86_64 x86_64 x86_64 GNU/Linux
```

* 到[Go安装包下载页面](https://studygolang.com/dl)找到对应的Linux版本，不需要在这个页面，下载，直接复制安装包名称即可，后面通过wget命令下载

  ![1634093587774](D:\PrivateProject\Introduction-to-algotithms\十五、Go语言程序设计\doc\img\Go安装包下载.png)

* 如果系统中之前已经安装过go，则先通过go env GOROOT查看之前的安装位置，然后切换到该目录下
```text
root@iZuf6a99qngm2w94o1aqz3Z:/home/SoftWareForCoding# go env GOROOT
/home/SoftWareForCoding/go
```

* 直接在Linux系统上通过wget命令安装
后面的go1.17.2.linux-amd64.tar.gz改成与你的Linux版本对应的go安装包文件名
```text
wget https://dl.google.com/go/go1.17.2.linux-amd64.tar.gz
```

* 先删除之前的安装目录，然后再解压新的安装包
```text
rm -rf /home/SoftWareForCoding/go && tar -C /home/SoftWareForCoding/ -xzf go1.17.2.linux-amd64.tar.gz
```

* 由于之前已经安装过go，所以环境遍历已经啥的之前已经配置好了，不需要再配置，直接用go version查看是否安装成功
可以看到go版本已经更新
```text
root@iZuf6a99qngm2w94o1aqz3Z:/home/SoftWareForCoding# go version
go version go1.17.2 linux/amd64
```

* 如果之前没有安装，则解压后需要添加环境变量
```text
vim /etc/profile

# 然后添加以下三行(将GOROOT/GOPATH添加到环境变量中)，保存退出
export GOROOT=/home/SoftWareForCoding/go
export GOPATH=/home/GoProject
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

# 执行使配置文件生效
source /etc/profile
```