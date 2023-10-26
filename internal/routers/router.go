package routers

import (
	v1 "WowjoyProject/WowjoyQueueCall/internal/routers/api/v1"
	"WowjoyProject/WowjoyQueueCall/internal/routers/api/ws"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// 注册中间件
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := r.Group("/api/v1")
	{
		// 呼叫返回文件形式stream
		apiv1.POST("/CallFile", v1.CallFile)
		// 呼叫返回数据流形式
		apiv1.POST("/CallStream", v1.CallStream)
		// 插入患者信息
		apiv1.POST("/InsPatientData", v1.InsPatientData)
		// 更新获取患者信息
		apiv1.GET("/HandGetPatientData", v1.HandGetPatientData)
	}
	// websocket
	apiws := r.Group("/api/ws")
	{
		apiws.GET("", ws.HandleWebSocket)

	}
	return r
}
