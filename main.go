package main

import (
	"context"
	"go-gin-example/models"
	"go-gin-example/pkg/logging"
	"go-gin-example/pkg/setting"
	"go-gin-example/routers"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
)

func main(){
	router := routers.InitRouter()
	setting.SetUp()
	models.SetUp()
	logging.SetUp()
	s:=&http.Server{
		Addr:           strconv.Itoa(setting.ServerSetting.ServerPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeOut,
		WriteTimeout:   setting.ServerSetting.WriteTimeOut,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		 if err := s.ListenAndServe();err==nil{
			 logging.Info("Listen: %s\n", err)
		 }
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	logging.Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		logging.Fatal("Server Shutdown:", err)
	}

	logging.Info("Server exiting")

}
