package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile(`C:\Users\Aone\go\src\gitlab.renrenchongdian.com\studygo\day09\http_server\f1.txt`)
	if err != nil {
		fmt.Printf("read file failed,err:%s\n", err)
		return
	}
	w.Write(b)
}

func f2(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	fmt.Println(r.Method)
	fmt.Println(ioutil.ReadAll(r.Body))
	w.Write([]byte("ok"))
}

func main() {
	http.HandleFunc("/hello/", f1)
	http.HandleFunc("/world/", f2)
	http.ListenAndServe("127.0.0.1:9000", nil)
}
