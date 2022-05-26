package main

import (
	"fmt"
	"strconv"
)

type Handler interface {
	Handler(handlerID int) string
}
type handler struct {
	name         string
	next         Handler
	RequestLevel int
}

func NewHandler(name string, next Handler, requestLevel int) *handler {
	return &handler{
		name:         name,
		next:         next,
		RequestLevel: requestLevel,
	}
}

func (h *handler) Handler(requestLevel int) string {
	if h.RequestLevel == requestLevel {
		return h.name + " handled " + strconv.Itoa(requestLevel)
	}
	if h.next == nil {
		return ""
	}
	return h.next.Handler(requestLevel)
}

func main() {
	wang := NewHandler("wang", nil, 1)
	zhang := NewHandler("zhang", wang, 2)
	wu := NewHandler("wu", zhang, 3)
	r := wu.Handler(2)
	fmt.Println(r)
}
