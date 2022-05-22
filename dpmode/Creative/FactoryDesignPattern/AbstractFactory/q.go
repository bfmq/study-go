package main

import "fmt"

type BaseDBV3 struct {
	address string
	port    int64
	account string
	passwd  string
}

//Connect()应该返回的是连接对象,这里做了省略,所以用string来代替
type DBconnectV3 interface {
	InitV3(address string, port int64, account string, passwd string)
	ConnectV3() string
}

type DBOpt interface {
	Create() error
	Update() error
	Delete() error
	Select() error
}

func (this *BaseDBV3) InitV3(address string, port int64, account string, passwd string) {
	this.address = address
	this.port = port
	this.account = account
	this.passwd = passwd
}

type MysqlDBV3 struct {
	BaseDBV3
}

func (mdb *MysqlDBV3) ConnectV3() string {
	return fmt.Sprintf("连接了mysql,地址是:%s", fmt.Sprintf("%s:%d", mdb.address, mdb.port))
}

func (mdb *MysqlDBV3) Create() error {
	return nil
}
func (mdb *MysqlDBV3) Update() error {
	return nil
}
func (mdb *MysqlDBV3) Delete() error {
	return nil
}
func (mdb *MysqlDBV3) Select() error {
	return nil
}

type CreateMysqlDBInterV3 interface {
	CreateMysqlDBConnectV3() DBconnectV3
}

type CreateMysqlDBFactoryV3 struct{}

func (cmysql *CreateMysqlDBFactoryV3) CreateMysqlDBConnectV3() DBconnectV3 {
	return &MysqlDBV3{}
}

type MongoDBV3 struct {
	BaseDBV3
}

func (mdb *MongoDBV3) ConnectV3() string {
	return fmt.Sprintf("mongo,地址是:%s", fmt.Sprintf("%s:%d", mdb.address, mdb.port))
}

func (mdb *MongoDBV3) Create() error {
	return nil
}
func (mdb *MongoDBV3) Update() error {
	return nil
}
func (mdb *MongoDBV3) Delete() error {
	return nil
}
func (mdb *MongoDBV3) Select() error {
	return nil
}

type CreateMongodbDBInterV3 interface {
	CreateMongodbDBConnectV3() DBconnectV3
}

type CreateMongodbDBFactoryV3 struct {
}

func (cmongo *CreateMysqlDBFactoryV3) CreateMongodbDBConnectV3() DBconnectV3 {
	return &MongoDBV3{}
}
