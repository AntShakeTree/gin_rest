package main

import (
	"github.com/gin-gonic/gin"
	//"net/http"
	//
	//"github.com/go-xorm/xorm"
	//
	//"github.com/tommy351/gin-sessions"

	"gin_rest/controller"
	"gin_rest/common"
	//"strconv"
	//
	//"broker_order/entity"
	"net/http"
)

func registerRouter(router *gin.Engine) {
	new(controller.OrderController).Router(router)
}

func main() {
	router := gin.Default()
	registerRouter(router)
	http.ListenAndServe(common.GetCfg().App["addr"]+":"+common.GetCfg().App["port"], router)
}
