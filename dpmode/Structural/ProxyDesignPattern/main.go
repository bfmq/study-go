package main

import (
	"errors"
	"sync"
)

type server interface {
	handleRequest(string, string) (int, error)
}

type RealRequest struct{}

func (request *RealRequest) handleRequest(url string, method string) (int, error) {
	if url == "zhihu.com" && method == "post" {
		return 200, nil
	}
	return 404, errors.New("not find, 404")
}

type ngx_http_proxy_module struct {
	request          *RealRequest
	maxRequestNum    int
	handleRequestMap map[string]int
}

var instance *ngx_http_proxy_module
var once sync.Once

func GetNgxProxyInstance(maxRequestNum int) *ngx_http_proxy_module {
	once.Do(func() {
		instance = &ngx_http_proxy_module{
			request:          &RealRequest{},
			maxRequestNum:    maxRequestNum,
			handleRequestMap: make(map[string]int, 0),
		}
	})
	return instance
}

func (nginx *ngx_http_proxy_module) check(url string) error {
	if _, ok := nginx.handleRequestMap[url]; !ok {
		return nil
	}
	if nginx.handleRequestMap[url] >= nginx.maxRequestNum {
		return errors.New("the num of request more than limit")
	}
	return nil
}
func (nginx *ngx_http_proxy_module) addRequstCount(url string) {
	nginx.handleRequestMap[url] = nginx.handleRequestMap[url] + 1
}

func (nginx *ngx_http_proxy_module) HandleRequest(url string, method string) (int, error) {
	if err := nginx.check(url); err != nil {
		return 403, err
	}
	nginx.addRequstCount(url)
	return nginx.request.handleRequest(url, method)
}
