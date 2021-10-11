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

