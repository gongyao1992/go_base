package main

import (
	"github.com/robfig/cron"
	"log"
	"time"
)

func main()  {
	log.Println("Starting...")

	c := cron.New()
	c.AddFunc("@every 10s", func() {
		log.Println("Run models.CleanAllTag...")
	})
	c.AddFunc("@every 1m", func() {
		log.Println("Run models.CleanAllArticle...")
	})

	c.Start()

	// 防止结束
	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}
