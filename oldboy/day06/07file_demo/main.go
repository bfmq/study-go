package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open(`C:\Users\Aone\go\src\gitlab.renrenchongdian.com\studygo\day06\05mylogger_test\beifang.log`)
	if err != nil {
		fmt.Printf("open file failed,err:%v\n", err)
		return
	}
	fmt.Printf("%T\n", f)

	fileInfo, err := f.Stat()
	if err != nil {
		fmt.Printf("get file failed,err:%v\n", err)
		return
	}
	fmt.Printf("文件大小是%dB\n", fileInfo.Size())
}
