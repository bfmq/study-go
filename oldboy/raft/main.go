package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/rpc"
	"sync"
	"time"
)

const (
	// 三节点常量
	raftCount = 3
)

// 声明leader对象
type Leader struct {
	// 任期
	Term int
	// LeaderId编号
	LeaderId int
}

// 声明raft
type Raft struct {
	// 锁
	mu sync.Mutex
	// 节点编号
	me int
	// 当前任期
	currentTerm int
	// 为哪个节点投票
	voteFor int
	// 3个状态
	// 0 follower	1 candidate		2 leader
	state int
	// 发送最后一条数据时间
	lastMessageTime int64
	// 设置当前节点领导
	currentLeader int
	// 节点间信息通道
	message chan bool
	// 选举通道
	electCh chan bool
	// 心跳通道
	heartBeat chan bool
	// 返回心跳通道
	heartbeatRe chan bool
	// 超时时间
	timeout int
}

var leader = Leader{0, -1}

func Make(me int) *Raft {
	rf := &Raft{}
	rf.me = me
	// -1代表谁都不投
	rf.voteFor = -1
	// 0 follower
	rf.state = 0
	rf.timeout = 0
	rf.currentLeader = -1
	// 节点任期
	rf.setTerm(0)
	// 初始化通道
	rf.message = make(chan bool)
	rf.electCh = make(chan bool)
	rf.heartBeat = make(chan bool)
	rf.heartbeatRe = make(chan bool)
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())

	// 选举协程
	go rf.election()
	// 心跳协程
	go rf.sendLeaderHeartBeat()
	return rf
}

func (rf *Raft) setTerm(term int) {
	rf.currentTerm = term
}

func (rf *Raft) election() {
	var result bool
	for {
		// 设置超时
		timeout := randRange(150, 300)
		rf.lastMessageTime = millisecond()
		select {
		case <-time.After(time.Duration(timeout) * time.Millisecond):
			fmt.Println("当前状态为:", rf.state)
		}
		result = false
		for !result {
			// 选主逻辑
			result = rf.election_one_round(&leader)
		}

	}
}

// 实现选主逻辑
func (rf *Raft) election_one_round(leader *Leader) bool {
	// 定义超时
	var timeout int64
	timeout = 100
	// 投票数量
	var vote int
	// 定义是否心跳信号产生
	var triggerHeartbeat bool
	// 时间
	last := millisecond()
	// 设置返回值
	success := false

	// 当前节点变candidate
	rf.mu.Lock()
	rf.becomeCandidate()
	rf.mu.Unlock()
	fmt.Println("start electing leader")

	for {
		// 遍历所有节点拉选票
		for i := 0; i < raftCount; i++ {
			if i != rf.me {
				go func() {
					if leader.LeaderId < 0 {
						// 设置投票
						rf.electCh <- true
					}
				}()
			}
		}
		// 设置投票数量
		vote = 1
		// 遍历
		for i := 0; i < raftCount; i++ {
			// 计算投票数量
			select {
			case ok := <-rf.electCh:
				if ok {
					// 投票数量加一
					vote++
					// 选票大于节点一半则当选
					success = vote > raftCount/2
					if success && !triggerHeartbeat {
						// 变成主节点，选主成功了
						// 开始触发心跳信号检测
						triggerHeartbeat = true
						// 变成主
						rf.mu.Lock()
						rf.becomeLeader()
						rf.mu.Unlock()
						// 此时由leader给其他节点发心跳信号
						rf.heartBeat <- true
						fmt.Println(rf.me, "号节点成为了leader")
						fmt.Println("leader开始发送心跳信号了！")
					}
				}
			}
		}
		// 做最后校验工作
		// 不超时，票数大于一半，证明成功
		if timeout+last < millisecond() && vote > raftCount/2 && rf.currentLeader > -1 {
			break
		} else {
			// 等待操作
			select {
			case <-time.After(time.Duration(10) * time.Millisecond):
			}
		}
	}
	return success
}

// leader节点发送心跳信号
// 顺便完成数据同步
// 看小弟挂没挂
func (rf *Raft) sendLeaderHeartBeat() {
	// 死循环
	for {
		select {
		case <-rf.heartBeat:
			rf.sendAppendEntriesImpl()
		}
	}
}

// 用于返给leader的确认信号
func (rf *Raft) sendAppendEntriesImpl() {
	// leader的逻辑
	if rf.currentLeader == rf.me {
		var success_count = 0
		// 设置确认信号
		for i := 0; i < raftCount; i++ {
			if i != rf.me {
				go func() {
					// rf.heartbeatRe <- true
					// 相当于客户端
					rp, err := rpc.DialHTTP("tcp", "127.0.0.1:8080")
					if err != nil {
						log.Fatal(err)
					}
					// 接受服务端返回的信息
					// 接受服务端返回的变量
					var ok = false
					err = rp.Call("Raft.Communication", Param{"hello"}, &ok)
					if err != nil {
						log.Fatal(err)
					}
					if ok {
						rf.heartbeatRe <- true
					}
				}()
			}
		}
		// 计算返回信号个数
		for i := 0; i < raftCount; i++ {
			select {
			case ok := <-rf.heartbeatRe:
				if ok {
					success_count++
					if success_count > raftCount/2 {
						fmt.Println("投票选举成功，心跳ok")
						fmt.Println("程序结束")
					}
				}
			}
		}
	}
}

// 修改状态leader
func (rf *Raft) becomeLeader() {
	rf.state = 2
	rf.currentLeader = rf.me
}

// 修改状态candidate
func (rf *Raft) becomeCandidate() {
	rf.state = 1
	rf.setTerm(rf.currentTerm + 1)
	rf.voteFor = rf.me
	rf.currentLeader = -1
}

// 获取当前时间，发送最后一条数据时间
func millisecond() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// 随机值
func randRange(min, max int64) int64 {
	return rand.Int63n(max-min) + min
}

func main() {
	// 过程：3个节点，最初都是follower
	// 若有candidate状态，进行投票拉票
	// 产生leader

	for i := 0; i < raftCount; i++ {
		// 创建3个节点
		Make(i)
	}

	// 加入服务端监听
	rpc.Register(new(Raft))
	rpc.HandleHTTP()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

	for {
	}
}

// 首字母大写，RPC规范
// 分布式通讯
type Param struct {
	Msg string
}

// 通讯方法
func (rf *Raft) Communication(p *Param, a *bool) error {
	fmt.Println(p.Msg)
	*a = true
	return nil
}
