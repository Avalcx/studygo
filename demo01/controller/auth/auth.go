package auth

import (
	"net/http"
	"project01/config"
	"project01/model/routermodel"
	"project01/utils/jwtutil"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user UserInfo
	returnData := routermodel.NewReturnData()
	if err := c.ShouldBind(&user); err != nil {
		returnData.Status = 2001
		returnData.Message = "无效的参数"
		c.JSON(http.StatusOK, returnData)
		return
	}

	if user.Username == config.Username && user.Password == config.Password {
		tokenString, _ := jwtutil.CreateToken(user.Username)
		returnData.Status = 2000
		returnData.Message = "登录成功"
		returnData.Data["token"] = tokenString
		c.JSON(http.StatusOK, returnData)
		return
	}

	returnData.Status = 2002
	returnData.Message = "账户或密码错误"
	c.JSON(http.StatusOK, returnData)
}

func Logout(c *gin.Context) {
	returnData := routermodel.NewReturnData()
	returnData.Status = 2000
	returnData.Message = "登出成功"
	c.JSON(http.StatusOK, returnData)
}

type UserInfo struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}
