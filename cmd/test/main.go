package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "127.0.0.1", "请输入host地址")
	port := flag.Int("port", 3306, "请输入端口号")
	flag.Parse() // 解析参数
	fmt.Printf("%s:%d\n", *host, *port)
}
