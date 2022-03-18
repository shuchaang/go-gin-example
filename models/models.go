package models

import (
	"fmt"
	"go-gin-example/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

var gormDB *gorm.DB

type Model struct {
	ID int `gorm:"primary_key" json:"id"`
	CreatedOn int64 `json:"created_on"`
	ModifiedOn int64 `json:"modified_on"`
}

func SetUp() {
	dbName := setting.DataBaseSetting.Name
	user := setting.DataBaseSetting.User
	password := setting.DataBaseSetting.Password
	host := setting.DataBaseSetting.Host
	tablePrefix := setting.DataBaseSetting.TablePrefix
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			user,
			password,
			host,
			dbName),
	}), &gorm.Config{
		SkipDefaultTransaction: true,
		DryRun:false,
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: tablePrefix,
			SingularTable: true,
		},
	})
	if err != nil {
		log.Println(err)
	}
	s, err := gormDB.DB()
	if err != nil {
		log.Println(err)
	}
	s.SetMaxIdleConns(10)
	s.SetMaxOpenConns(100)
	s.SetConnMaxLifetime(time.Duration(1*time.Hour))
	gormDB.Callback().Create().Replace("gorm:before_create", setTimeStampForCreateCallback)
	gormDB.Callback().Update().Replace("gorm:before_update", updateTimeStampForBeforeUpdateCallback)
}


func setTimeStampForCreateCallback(db *gorm.DB){
	var nowTime = time.Now().Unix()
	db.Statement.SetColumn("CreatedOn", nowTime)
}

func updateTimeStampForBeforeUpdateCallback(db *gorm.DB){
	var nowTime = time.Now().Unix()
	db.Statement.SetColumn("ModifiedOn", nowTime)
}