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
	router.GET("/user/signin", handler.GETSigninHandler)
	router.POST("/user/signin", handler.POSTSigninHandler)

	router.Use(handler.HTTPInterceptor)

	router.POST("/user/info", handler.POSTUserInfoHandler)

	router.POST("/file/query", handler.POSTFileQueryHandler)
	return router
}
