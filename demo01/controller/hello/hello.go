package hello

import (
	"net/http"
	"project01/model/routermodel"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	returnData := routermodel.NewReturnData()
	username := c.MustGet("username").(string)
	returnData.Status = 2000
	returnData.Message = "success"
	returnData.Data["username"] = username
	c.JSON(http.StatusOK, returnData)
}
