package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)


var (
	conf       *ini.File
)

type App struct {
	JwtSecret string
	PageSize int

	ImagePrefixUrl string
	ImageSavePath string
	ImageMaxSize int
	ImageAllowExists []string

	LogSavePath string
	LogSaveName string
	TimeFormat string
}

type Server struct {
	RunMode string `ini:"RUN_MODE"`
	ServerPort int `ini:"HTTP_PORT"`
	ReadTimeOut time.Duration `ini:"READ_TIMEOUT"`
	WriteTimeOut time.Duration `ini:"WRITE_TIMEOUT"`
}

var ServerSetting  = &Server{}

type DataBase struct {
	Type string
	User string
	Password string
	Host string
	Name string
	TablePrefix string
}

var DataBaseSetting  = &DataBase{}

var AppSetting = &App{}


func SetUp(){
	var err error

	conf, err = ini.Load("conf/app.ini")
	if err!=nil{
		log.Fatalf("load app.ini error :%v",err)
	}
	err = conf.Section("app").MapTo(AppSetting)
	if err!=nil{
		log.Fatalf("load app.ini error :%v",err)
	}
	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024

	err = conf.Section("server").MapTo(ServerSetting)
	if err!=nil{
		log.Fatalf("load app.ini error :%v",err)
	}
	ServerSetting.ReadTimeOut = ServerSetting.ReadTimeOut * time.Second
	ServerSetting.WriteTimeOut = ServerSetting.WriteTimeOut * time.Second

	err = conf.Section("database").MapTo(DataBaseSetting)
	if err!=nil{
		log.Fatalf("load app.ini error :%v",err)
	}
}



