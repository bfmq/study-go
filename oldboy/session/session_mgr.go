package session

// session管理者接口
type SessionMgr interface {
	Init(addr string, options ...string) error
	Create() (Session, error)
	Get(sessionId string) (Session, error)
}
