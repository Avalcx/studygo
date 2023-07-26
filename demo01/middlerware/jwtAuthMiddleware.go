package middlerware

import (
	"net/http"
	"project01/model/routermodel"
	"project01/utils/jwtutil"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() func(c *gin.Context) {
	returnData := routermodel.NewReturnData()
	return func(c *gin.Context) {
		// 除了login和logout，其他路径全局注册
		requestUrl := c.FullPath()
		if requestUrl == "/api/auth/login" {
			c.Next()
			return
		}

		// 校验请求头中是否包含Authorization
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			returnData.Status = 2000
			returnData.Message = "请求头token为空"
			c.JSON(http.StatusOK, returnData)
			c.Abort()
			return
		}

		// 校验请求头的格式并取出token
		params := strings.SplitN(authHeader, " ", 2)
		if !(len(params) == 2 && params[0] == "Bearer") {
			returnData.Status = 2004
			returnData.Message = "token格式错误"
			c.JSON(http.StatusOK, returnData)
			c.Abort()
			return
		}

		myCustomClaims, err := jwtutil.ParseToken(params[1])
		if err != nil {
			returnData.Status = 2005
			returnData.Message = "无效的Token"
			c.JSON(http.StatusOK, returnData)
			c.Abort()
			return
		}
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("username", myCustomClaims.Username)
		c.Next()
	}
}
