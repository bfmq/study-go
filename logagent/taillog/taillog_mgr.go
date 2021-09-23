package taillog

import (
	"fmt"
	"time"

	"gitlab.renrenchongdian.com/studygo/logagent/etcd"
)

var tskMgr *tailMgr

type tailMgr struct {
	logEntry    []*etcd.LogEntry
	tskMap      map[string]*TailTask
	newConfChan chan []*etcd.LogEntry
}

func Init(logConf []*etcd.LogEntry) {
	tskMgr = &tailMgr{
		logEntry:    logConf,
		tskMap:      make(map[string]*TailTask, 16),
		newConfChan: make(chan []*etcd.LogEntry),
	}
	for _, logEntry := range tskMgr.logEntry {
		tailObj := NewTailTask(logEntry.Path, logEntry.Topic)
		mk := fmt.Sprintf("%s_%s", logEntry.Path, logEntry.Topic)
		tskMgr.tskMap[mk] = tailObj
	}
	go tskMgr.run()
}

func (t *tailMgr) run() {
	for {
		select {
		case newConf := <-t.newConfChan:
			for _, conf := range newConf {
				mk := fmt.Sprintf("%s_%s", conf.Path, conf.Topic)
				_, ok := t.tskMap[mk]
				if ok {
					// 已有不操作
					continue
				} else {
					tailObj := NewTailTask(conf.Path, conf.Topic)
					t.tskMap[mk] = tailObj
				}
			}

			for _, c1 := range t.logEntry {
				isDelete := true
				for _, c2 := range newConf {
					if c1.Path == c2.Path && c1.Topic == c2.Topic {
						isDelete = false
						continue
					}
				}
				if isDelete {
					// 把c1删掉
					mk := fmt.Sprintf("%s_%s", c1.Path, c1.Topic)
					t.tskMap[mk].cancelFunc()
				}
			}

		default:
			time.Sleep(time.Second)
		}
	}
}

func NewConfChan() chan []*etcd.LogEntry {
	return tskMgr.newConfChan
}
