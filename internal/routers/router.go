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
	r.Use(Cors())
	apiv1 := r.Group("/api/v1")
	{
		// 呼叫返回文件形式stream
		apiv1.POST("/CallFile", v1.CallFile)
		// 呼叫返回数据流形式
		apiv1.POST("/CallStream", v1.CallStream)
		// 插入患者信息
		apiv1.POST("/InsPatientData", v1.InsPatientData)
		// 手动获取患者信息
		apiv1.GET("/HandGetPatientData", v1.HandGetPatientData)
		// 前端log保存接口
		apiv1.POST("/weblog", v1.InsWebLog)
	}
	// websocket
	apiws := r.Group("/api/ws")
	{
		apiws.GET("", ws.HandleWebSocket)

	}
	return r
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "*")
			c.Header("Access-Control-Allow-Headers", "*")
			c.Header("Access-Control-Expose-Headers", "*")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(204)
		}
		c.Next()
	}
}
