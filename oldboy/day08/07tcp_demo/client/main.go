package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// 1.与server建立链接
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Printf("dial 127.0.0.1:8888 failed,err:%s\n", err)
		return
	}

	// 2.发数据通信
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("请输入：")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == "exit" {
			break
		}
		conn.Write([]byte(text))
	}
	conn.Close()
}
