package util

import (
	"blog/models"
	"github.com/robfig/cron/v3"
	"log"
)

func ScheduleTask() {
	log.Println("Starting...")

	c := cron.New(cron.WithSeconds())
	c.AddFunc("*/15 * * * * *", func() {
		log.Println("Run models.CleanAllTag...")
		models.CleanAllTag()
	})

	c.AddFunc("*/15 * * * * *", func() {
		log.Println("Run models.CleanAllArticle...")
		models.CleanAllArticle()
	})

	c.Start()
}
