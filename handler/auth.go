package handler

import (
	"github.com/gin-gonic/gin"
)

// HTTPInterceptor HTTP 请求拦截器
func HTTPInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		// username := c.Request.FormValue("username")
		// token := c.Request.FormValue("token")
		// ok := db.IsTokenValid(username, token)
		// if !ok {
		// 	c.Abort()
		// 	log.Println("Token is not validated")
		// 	resp := util.RespMsg{
		// 		Code: -3,
		// 		Msg:  "Token 无效",
		// 		Data: nil,
		// 	}
		// 	c.JSON(http.StatusOK, resp.JSONBytes())
		// 	return
		// }
		c.Next()
	}
}
