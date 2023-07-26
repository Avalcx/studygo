package hello

import (
	"net/http"
	"project01/model/routermodel"

	"github.com/gin-gonic/gin"
)

var returnData = routermodel.NewReturnData()

func Hello(c *gin.Context) {
	username := c.MustGet("username").(string)
	returnData.Status = 2000
	returnData.Message = "success"
	returnData.Data["username"] = username
	c.JSON(http.StatusOK, returnData)
}
