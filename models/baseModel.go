package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"goweb/conf"
	"log"
)

type BaseModel struct {
	dbConnect *gorm.DB
}

func (baseModel BaseModel) getDbConnect() *gorm.DB {
	var err error
	mysqlConfig := conf.GetMysqlConf("default")

	fmt.Println(mysqlConfig["user"] + ":" + mysqlConfig["password"] + "@(" + mysqlConfig["host"] + ":" + mysqlConfig["port"] + ")/" + mysqlConfig["dbname"] + "?charset=utf8&parseTime=True&loc=Local")

	baseModel.dbConnect, err = gorm.Open("mysql", mysqlConfig["user"]+":"+mysqlConfig["password"]+"@("+mysqlConfig["host"]+":"+mysqlConfig["port"]+")/"+mysqlConfig["dbname"]+"?charset=utf8&parseTime=True&loc=Local")
	//defer baseModel.dbConnect.Close()
	if err != nil {
		// 打日志
		log.Println("数据库连接错误----", err)
		return nil
	}
	return baseModel.dbConnect
}

func (baseModel BaseModel) close() {
	if baseModel.dbConnect != nil {
		baseModel.dbConnect.Close()
	}
}
