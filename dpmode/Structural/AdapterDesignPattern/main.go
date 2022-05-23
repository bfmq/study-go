package main

import (
	"fmt"
)

type OldInterface interface {
	InsertToDatabase(Data interface{}) (bool, error)
}

type AddCustomInfoToMysql struct {
	DbName string
}

func (pA *AddCustomInfoToMysql) InsertToDatabase(info interface{}) (bool, error) {
	switch info.(type) {
	case string:
		fmt.Println("add ", info.(string), " to ", pA.DbName, " successful!")
	}
	return true, nil
}

type NewInterface interface {
	SaveData(Data interface{}) (bool, error)
}

type Adapter struct {
	OldInterface
}

func (pA *Adapter) SaveData(Data interface{}) (bool, error) {
	fmt.Println("In Adapter")
	return pA.InsertToDatabase(Data)
}

func main() {

	var iNew NewInterface
	iNew = &Adapter{OldInterface: &AddCustomInfoToMysql{DbName: "mysql"}}
	iNew.SaveData("helloworld")
	return
}
