package route

import (
	"file-store-server/service/apigw/handler"

	"github.com/gin-gonic/gin"
)

// Router 网关api路由
func Router() *gin.Engine {
	router := gin.Default()

	router.Static("/static/", "./static")

	router.GET("/user/signup", handler.GETSignupHandler)
	router.POST("/user/signup", handler.POSTSignupHandler)

	return router
}
