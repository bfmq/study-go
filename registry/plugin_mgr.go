package registry

import (
	"context"
	"fmt"
	"sync"
)

// 管理者结构体
type PluginMgr struct {
	// map维护所有插件
	plugins map[string]Registry
	lock    sync.RWMutex
}

var (
	pluginMgr = &PluginMgr{
		plugins: make(map[string]Registry, 8),
	}
)

// 插件注册外部暴露方法
func RegisterPlugin(registry Registry) (err error) {
	return pluginMgr.registerPlugin(registry)
}

// 插件注册内部实现方法
func (p *PluginMgr) registerPlugin(plugin Registry) (err error) {
	p.lock.Lock()
	defer p.lock.Unlock()

	// 先看有没有
	_, ok := p.plugins[plugin.Name()]
	if ok {
		err = fmt.Errorf("register plugin exist")
		return
	}

	p.plugins[plugin.Name()] = plugin
	return
}

// 插件初始化
func InitRegister(ctx context.Context, name string, opts ...Option) (register Registry, err error) {
	return pluginMgr.initRegister(ctx, name, opts...)
}

func (p *PluginMgr) initRegister(ctx context.Context, name string, opts ...Option) (register Registry, err error) {
	p.lock.Lock()
	defer p.lock.Unlock()

	// 先看有没有，服务不存在不可以初始化
	plugin, ok := p.plugins[name]
	if !ok {
		err = fmt.Errorf("plugin %s not exist", name)
		return
	}
	register = plugin
	// 进行插件初始化
	err = register.Init(ctx, opts...)
	return
}
