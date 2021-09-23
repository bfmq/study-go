package main

import (
	"fmt"
	"os"
)

// func readFromFile() {
// 	f, err := os.Open("C:\\Users\\Aone\\Desktop\\new 12.txt")
// 	if err != nil {
// 		fmt.Println("文件打开失败")
// 		return
// 	}
// 	defer f.Close()

// 	var tmp [128]byte
// 	for {
// 		n, err := f.Read(tmp[:])
// 		if err != nil {
// 			fmt.Println("文件读取失败")
// 			return
// 		}
// 		fmt.Println(n)
// 		fmt.Println(string(tmp[:n]))
// 		if n < 128 {
// 			return
// 		}
// 	}
// }

// func readFromFileByBufio() {
// 	f, err := os.Open("C:\\Users\\Aone\\Desktop\\new 12.txt")
// 	if err != nil {
// 		fmt.Println("文件打开失败")
// 		return
// 	}
// 	defer f.Close()

// 	reader := bufio.NewReader(f)
// 	for {
// 		line, err := reader.ReadString('\n')
// 		if err == io.EOF {
// 			return
// 		}
// 		if err != nil {
// 			fmt.Println("文件读取失败")
// 			return
// 		}
// 		fmt.Print(line)
// 	}
// }

// func readFromFileByIoutil() {
// 	ret, err := ioutil.ReadFile("C:\\Users\\Aone\\Desktop\\new 12.txt")
// 	if err != nil {
// 		fmt.Println("文件读取失败")
// 		return
// 	}
// 	fmt.Println(string(ret))
// }
func main() {
	f, err := os.OpenFile("./new.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("文件打开失败")
		return
	}
	defer f.Close()

}
