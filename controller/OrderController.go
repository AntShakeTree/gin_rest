package controller

import (
	"github.com/gin-gonic/gin"
	"gin_rest/repositories"
	"gin_rest/entity"

	"gin_rest/common"

	"strconv"
	"gin_rest/service"
)

type OrderController struct {
}

func (ctrl *OrderController) Router(router *gin.Engine) {

	r := router.Group("broker")
	r.POST("/subscribe/order_path", ctrl.SubscribeOrderPath)
	r.GET("/get/order_path", ctrl.GetOrderPath)
	r.GET("/search/order_path", ctrl.GetOrderPathAll)
	//r.POST("findOne", ctrl.findOne)
}

func GetOrderPathService() service.LbsService {
	return service.CreateInstance(&repositories.OrderPathRepositoryFactory{})
}
func GetPassengerLbsResService() service.LbsService {
	return service.CreateInstance(&repositories.PassengerLBSRepositoryFactory{})
}

//同步OrderPath
func (ctrl *OrderController) SubscribeOrderPath(ctx *gin.Context) {
	var req entity.OrderPathReq
	ctx.ShouldBindJSON(&req)
	var buf = common.GetStringBuffer()
	buf.Join("orders/").Join(strconv.Itoa(req.Area)).Join("/1/").Join(strconv.Itoa(req.OrderId))
	//实例化
	lss := GetOrderPathService()
	lss.CreateOrderPath(req)
	common.ResultOk(ctx, req)
}

//查询
func (ctrl *OrderController) GetOrderPath(ctx *gin.Context) {
	lss := GetOrderPathService()
	common.ResultOk(ctx, lss.GetOrderPath(1))
}

//查询全部路径
func (ctrl *OrderController) GetOrderPathAll(ctx *gin.Context) {
	lss := GetPassengerLbsResService()
	common.ResultOk(ctx, lss.FindOrderPathAll())
}
