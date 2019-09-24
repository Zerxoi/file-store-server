package route

import (
	"file-store-server/service/upload/api"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

// Router 路由配置表
func Router() *gin.Engine {
	// gin framework, 包括Logger, Recovery
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Range", "x-requested-with", "Content-Disposition"},
		ExposeHeaders: []string{"Content-Length", "Accept-Range", "Cotent-Range", "Content-Disposition"},
	}))
	router.Static("/static/", "./static")

	// 文件存取接口
	router.GET("/file/upload", api.GETUploadHandler)
	router.POST("/file/upload", api.POSTUploadHandler)
	return router
}
