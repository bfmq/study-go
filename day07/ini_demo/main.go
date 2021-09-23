package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Database int    `ini:"database"`
	Password string `ini:"password"`
	Test     bool   `ini:"test"`
}

func loadIni(fileName string, data interface{}) error {
	t := reflect.TypeOf(data)
	if t.Kind() != reflect.Ptr {
		err := errors.New("data should be a pointer")
		return err
	}
	if t.Elem().Kind() != reflect.Struct {
		err := errors.New("data param should be a struct")
		return err
	}

	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("open file failed,err:%v\n", err)
		return err
	}
	lineSlice := strings.Split(string(b), "\r\n")
	var structName string
	for idx, line := range lineSlice {
		line = strings.TrimSpace(line)

		// 空行跳过
		if len(line) == 0 {
			continue
		}
		// 注释跳过
		if strings.HasPrefix(line, "#") || strings.HasPrefix(line, ";") {
			continue
		}

		// 以[开头为章节
		if strings.HasPrefix(line, "[") {
			// 不以]结尾不符合规则
			if !strings.HasSuffix(line, "]") {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return err
			}
			// 章节长度为0不符合规则
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return err
			}
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if sectionName == field.Tag.Get("ini") {
					structName = field.Name
					fmt.Printf("找到%s对应嵌套结构体%s\n", sectionName, structName)
				}
			}
		} else {
			// 不包含=，以=开头不符合规则
			if strings.Index(line, "=") == -1 || strings.HasPrefix(line, "=") {
				err := fmt.Errorf("line:%d syntax error", idx+1)
				return err
			}
			index := strings.Index(line, "=")
			key, value := strings.TrimSpace(line[:index]), strings.TrimSpace(line[index+1:])
			// 根据结构体名将data里的值取出来
			v := reflect.ValueOf(data)
			// 拿到嵌套结构体值信息
			sValue := v.Elem().FieldByName(structName)
			// 拿到嵌套结构体类型信息
			sType := sValue.Type()

			if sType.Kind() != reflect.Struct {
				err := fmt.Errorf("data中%s字段应该是一个结构体", structName)
				return err
			}

			var fieldName string
			var fieldType reflect.StructField
			// 遍历循环找到key
			for i := 0; i < sValue.NumField(); i++ {
				// tag信息存储在类型信息里
				field := sType.Field(i)
				fieldType = field
				if field.Tag.Get("ini") == key {
					fieldName = field.Name
					break
				}
			}

			if len(fieldName) == 0 {
				continue
			}
			// 赋值
			fieldObj := sValue.FieldByName(fieldName)
			switch fieldType.Type.Kind() {
			case reflect.String:
				fieldObj.SetString(value)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				valueInt, err := strconv.ParseInt(value, 10, 64)
				if err != nil {
					err = fmt.Errorf("line%d value type error", idx+1)
					return err
				}
				fieldObj.SetInt(valueInt)
			case reflect.Bool:
				valueBool, err := strconv.ParseBool(value)
				if err != nil {
					err = fmt.Errorf("line%d value type error", idx+1)
					return err
				}
				fieldObj.SetBool(valueBool)
			case reflect.Float32, reflect.Float64:
				valueFloat, err := strconv.ParseFloat(value, 64)
				if err != nil {
					err = fmt.Errorf("line%d value type error", idx+1)
					return err
				}
				fieldObj.SetFloat(valueFloat)
			}
		}
	}
	return nil
}

func main() {
	var cfg Config
	err := loadIni(`C:\Users\Aone\go\src\gitlab.renrenchongdian.com\studygo\day07\ini_demo\conf.ini`, &cfg)
	if err != nil {
		fmt.Printf("load ini failed,err:%v\n", err)
	}
	fmt.Printf("%#v", cfg)
}
