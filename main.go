package main

import (
	"WowjoyProject/WowjoyQueueCall/global"
	"WowjoyProject/WowjoyQueueCall/internal/model"
	"WowjoyProject/WowjoyQueueCall/internal/routers"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
)

// @title 叫号系统服务
// @version 1.0.0.1
// @description 叫号系统服务
// @termsOfService https://github.com/jianghuxiaoloulou/ObjectCloudService_Down.git
func main() {
	global.Logger.Info("***开始运行叫号系统服务***")
	global.ScreenRoomTotalData = make(map[string]map[string]global.CallData)
	// 定时获取数据
	// go timedTask()
	// 启动web服务
	web()
}

func web() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()

	ser := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	ser.ListenAndServe()
}

func timedTask() {
	// 定时任务
	MyCron := cron.New()
	// 获取患者数据
	MyCron.AddFunc(global.GeneralSetting.GetDataCronSpec, func() {
		global.Logger.Info("开始执行获取患者数据的定时任务")
		getData()
	})
	// 删除处理的数据
	// MyCron.AddFunc(global.GeneralSetting.DelDataCronSpec, func() {
	// 	global.Logger.Info("开始执行删除患者数据的定时任务")
	// 	delData()
	// })
	MyCron.Start()
	defer MyCron.Stop()
	select {}
}

func getData() {
	global.Logger.Debug("开始获取患者数据")
	model.GetPatientData()
}

func delData() {
	global.Logger.Debug("开始删除患者数据")
}
