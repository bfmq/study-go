package main

import "errors"

type ResourcePoolConfig struct {
	Name     string
	MaxTotal int
	MaxIdle  int
	MinIdle  int
}

type Builder struct {
	Name     string
	MaxTotal int
	MaxIdle  int
	MinIdle  int
}

func InitResourcePoolConfig(b *Builder) (rpc *ResourcePoolConfig, err error) {
	rpc.Name = b.Name
	rpc.MaxTotal = b.MaxTotal
	rpc.MaxIdle = b.MaxIdle
	rpc.MinIdle = b.MinIdle
	if !rpc.isValidate() {
		err = errors.New("Validate failed!")
	}
	return
}

func (rpc *ResourcePoolConfig) isValidate() bool {
	if rpc.MaxIdle >= rpc.MaxTotal || rpc.MinIdle >= rpc.MaxTotal {
		return false
	}
	return true
}

func InitBuilder(name string) *Builder {
	return &Builder{
		Name:     name,
		MaxTotal: 8,
		MaxIdle:  8,
		MinIdle:  0,
	}
}

func (b *Builder) SetMaxTotal(maxTotal int) {
	if maxTotal < 0 {
		panic("maxTotal")
	}
	b.MaxTotal = maxTotal
}

func (b *Builder) SetMaxIdle(maxIdle int) {
	if maxIdle < 0 {
		panic("maxIdle")
	}
	b.MaxIdle = maxIdle
}

func (b *Builder) SetMinIdle(minIdle int) {
	if minIdle < 0 {
		panic("minIdle")
	}
	b.MinIdle = minIdle
}
