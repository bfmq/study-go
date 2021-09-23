package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func processConn(conn net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	for {
		var tmp [128]byte
		n, err := conn.Read(tmp[:])
		if err != nil {
			fmt.Printf("read from conn failed,err:%s\n", err)
			return
		}
		fmt.Println(string(tmp[:n]))
		fmt.Print("请回复：")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == "exit" {
			break
		}
		conn.Write([]byte(text))
	}
}

func main() {
	// 1.启动本地端口服务
	listener, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Printf("start tcp server on 127.0.0.1:8888 failed,err:%s\n", err)
		return
	}

	// 2.等别人建立链接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("accept failed,err:%s\n", err)
			continue
		}
		// 3.通信交互
		go processConn(conn)
	}

}
