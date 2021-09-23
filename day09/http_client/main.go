package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	data := url.Values{}
	urObj, _ := url.Parse("http://127.0.0.1:9000/world/")
	data.Set("name", "bfmq")
	data.Set("age", "9000")
	dataStr := data.Encode()
	urObj.RawQuery = dataStr
	req, err := http.NewRequest("GET", urObj.String(), nil)
	if err != nil {
		fmt.Printf("NewRequest failed,err:%s\n", err)
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("http failed,err:%s\n", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp)

	// response, err := http.Get("http://127.0.0.1:9000/world/")
	// if err != nil {
	// 	fmt.Printf("get url failed,err:%s\n", err)
	// 	return
	// }
	// b, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	fmt.Printf("read response.Body failed,err:%s\n", err)
	// 	return
	// }
	// fmt.Println(string(b))
}
