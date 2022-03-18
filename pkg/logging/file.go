package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	LogSavePath = "logs/"
	LogSaveName = "app.log"
	TimeFormat = "20060102"
)

func init(){
	getLogFileFullPath()
}

func getLogFileFullPath() string{
	return fmt.Sprintf("%s%s%s", LogSavePath,LogSaveName, time.Now().Format(TimeFormat))
}

func mkdir(){
	pwd, _ := os.Getwd()
	err:= os.MkdirAll(pwd+"/"+LogSavePath, os.ModePerm)
	if err!=nil{
		panic(err)
	}
}

func openFile(path string)*os.File{
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		mkdir()
	}else if os.IsPermission(err) {
		log.Fatalf("Permission :%v", err)
	}
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to OpenFile :%v", err)
	}
	return file
}