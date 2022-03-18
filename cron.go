package main

import (
	"github.com/robfig/cron"
	"log"
)

func main0(){
	log.Println("starting task")
	c := cron.New()
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllTag...")
	})
	c.Start()
}
