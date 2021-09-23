package taillog

import (
	"context"
	"fmt"
	"time"

	"github.com/hpcloud/tail"
	"gitlab.renrenchongdian.com/studygo/logagent/kafka"
)

// 从日志文件收集日志模块

// 一个日志收集任务
type TailTask struct {
	path     string
	topic    string
	instance *tail.Tail
	// 为了实现退出t.run
	ctx        context.Context
	cancelFunc context.CancelFunc
}

func NewTailTask(path, topic string) (tailObj *TailTask) {
	ctx, cancel := context.WithCancel(context.Background())
	tailObj = &TailTask{
		path:       path,
		topic:      topic,
		ctx:        ctx,
		cancelFunc: cancel,
	}
	err := tailObj.init()
	if err != nil {
		fmt.Printf("new tail task failed,err:%s\n", err)
		return
	}
	return
}

func (t *TailTask) init() (err error) {
	config := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}
	t.instance, err = tail.TailFile(t.path, config)
	if err != nil {
		fmt.Printf("tail file %s failed,err:%s\n", t.path, err)
		return
	}

	go t.run()
	return
}

func (t *TailTask) run() {
	for {
		select {
		case <-t.ctx.Done():
			fmt.Printf("tail task:%s_%s 结束了...\n", t.path, t.topic)
			return
		case line := <-t.instance.Lines:
			kafka.SendToChan(t.topic, line.Text)
		default:
			time.Sleep(time.Second)
		}
	}

}
