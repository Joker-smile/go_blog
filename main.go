package main

import (
	"blog/pkg/setting"
	"blog/pkg/util"
	"blog/routers"
	"fmt"
	"net/http"
)

func main() {
	go util.ScheduleTask() //启动另外一个goroutine去执行定时任务
	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
