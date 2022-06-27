package main

import (
	"blog/models"
	"blog/pkg/setting"
	"blog/pkg/util"
	"blog/pkg/util/logging"
	"blog/routers"
	"fmt"
	"net/http"
)

func init() {
	setting.Setup()
	models.Setup()
	logging.Setup()
}

func main() {
	go util.ScheduleTask() //启动另外一个goroutine去执行定时任务
	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
