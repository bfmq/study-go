package main

import "fmt"

type BaseDB struct {
	address string
	port    int64
	user    string
	passwd  string
}

//Connect()应该返回的是连接对象,这里做了省略,所以用string来代替
type DBconnect interface {
	Init(address string, port int64, account string, passwd string)
	Connect() string
}

func (this *BaseDB) Init(address string, port int64, user string, passwd string) {
	this.address = address
	this.port = port
	this.user = user
	this.passwd = passwd
}

type MysqlDB struct {
	BaseDB
}

func (mdb *MysqlDB) Connect() string {
	return fmt.Sprintf("连接了mysql,地址是:%s", fmt.Sprintf("%s:%d", mdb.address, mdb.port))
}

type MongoDB struct {
	BaseDB
}

func (mdb *MongoDB) Connect() string {
	return fmt.Sprintf("mongo,地址是:%s", fmt.Sprintf("%s:%d", mdb.address, mdb.port))
}

// ==========================================================================================

type CreateDBFactory struct{}

func (cdb *CreateDBFactory) CreateDBConnect(dbname string) (bdb DBconnect, err error) {
	switch dbname {
	case "mysql":
		bdb = &MysqlDB{}
	case "mongodb":
		bdb = &MongoDB{}
	default:
		bdb = &MysqlDB{}
	}
	return
}

func main() {
	dbFactory := &CreateDBFactory{}
	mysqlObject, err := dbFactory.CreateDBConnect("mysql")
	if err != nil {
		panic("mysql error")
	}
	mysqlObject.Init("mysqladdress", 3306, "kengerukong", "kengerukong")
	mysqlConnect := mysqlObject.Connect()
	fmt.Println("mysql connect is " + mysqlConnect)
}
