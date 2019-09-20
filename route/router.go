package route

import (
	"file-store-server/handler"

	"github.com/gin-gonic/gin"
)

// Router 路由配置表
func Router() *gin.Engine {
	// gin framework, 包括Logger, Recovery
	router := gin.Default()

	// 处理静态资源
	router.Static("/static/", "../../static")

	// 不需要任何验证就能访问的接口
	router.GET("/user/signup", handler.GETSignupHandler)
	router.POST("/user/signup", handler.POSTSignupHandler)
	router.GET("/user/signin", handler.GETSigninHandler)
	router.POST("/user/signin", handler.POSTSigninHandler)

	// 加入中间件,用于校验Token的拦截器
	// Use 之后的所有handler都会经过拦截器进行Token校验
	router.Use(handler.HTTPInterceptor())

	// 用户相关接口
	router.POST("/user/info", handler.POSTUserInfoHandler) // OK

	// 文件存取接口
	router.GET("/file/upload", handler.GETUploadHandler)     // OK
	router.POST("/file/upload", handler.POSTUploadHandler)   // OK
	router.GET("/file/download", handler.GETDownloadHandler) // 有问题
	router.POST("/file/meta", handler.POSTGetFileMetaHandler)
	router.POST("/file/downloadurl", handler.POSTDownloadURLHandler) // 有问题
	router.POST("/file/query", handler.POSTFileQueryHandler)
	// router.POST("/file/update", handler.POSTUpdateFileMetaHandler) // 更新文件元信息
	// router.POST("/file/delete", handler.FileDeleteHandler) // 文件删除

	// 秒传接口
	router.POST("/file/fastupload", handler.POSTFastUploadHandler) // 有待修改

	// TODO:分块上传接口

	return router
}
