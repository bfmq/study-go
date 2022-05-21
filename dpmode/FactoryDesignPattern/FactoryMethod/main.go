package main

import "fmt"

type BaseDBV2 struct {
	address string
	port    int64
	account string
	passwd  string
}

//Connect()应该返回的是连接对象,这里做了省略,所以用string来代替
type DBconnectV2 interface {
	InitV2(address string, port int64, account string, passwd string)
	ConnectV2() string
}

func (this *BaseDBV2) InitV2(address string, port int64, account string, passwd string) {
	this.address = address
	this.port = port
	this.account = account
	this.passwd = passwd
}

type MysqlDBV2 struct {
	BaseDBV2
}

func (mdb *MysqlDBV2) ConnectV2() string {
	return fmt.Sprintf("连接了mysql,地址是:%s", fmt.Sprintf("%s:%d", mdb.address, mdb.port))
}

type MongoDBV2 struct {
	BaseDBV2
}

func (mdb *MongoDBV2) ConnectV2() string {
	return fmt.Sprintf("mongo,地址是:%s", fmt.Sprintf("%s:%d", mdb.address, mdb.port))
}

// ==========================================================================================

type CreateMysqlDBFactoryV2 struct{}

func (cdb *CreateMysqlDBFactoryV2) CreateDBConnectV2() (bdb *MysqlDBV2, err error) {
	bdb = &MysqlDBV2{}
	return bdb, nil
}

type CreateMongodbDBFactoryV2 struct{}

func (cdb *CreateMongodbDBFactoryV2) CreateDBConnectV2() (bdb *MongoDBV2, err error) {
	bdb = &MongoDBV2{}
	return bdb, nil
}

func main() {
	mysqlFactory := &CreateMysqlDBFactoryV2{}
	mysqlObject, err := mysqlFactory.CreateDBConnectV2()
	if err != nil {
		panic("mysql error")
	}
	mysqlObject.InitV2("mysqladdress", 3306, "kengerukong", "kengerukong")
	mysqlConnect := mysqlObject.ConnectV2()
	fmt.Println("mysql connect is " + mysqlConnect)

	mongodbFactory := &CreateMongodbDBFactoryV2{}
	mongodbObject, err := mongodbFactory.CreateDBConnectV2()
	if err != nil {
		panic("mongodb error")
	}
	mongodbObject.InitV2("mongodbaddress", 3307, "kengerukong", "kengerukong")
	mongodbConnect := mongodbObject.ConnectV2()
	fmt.Println("mongodb connect is " + mongodbConnect)
}
