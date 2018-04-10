package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Result(ctx *gin.Context, code int, data interface{}, msg string) {
	ctx.JSON(http.StatusOK, gin.H{"code": code, "data": data, "msg": msg})
}

func ResultOk(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, Response.Instances(data))
}
func ResultOkNoBody(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Response.InstancesNoBody())
}

