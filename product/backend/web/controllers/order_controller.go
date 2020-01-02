//Small_HG
package controllers

import (
	"product/services"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type OrderController struct {
	Ctx iris.Context
	OrderService services.IOrderService
}

func (o *OrderController) Get() mvc.View {

	orderArray,err:=o.OrderService.GetAllOrderInfo()
	if err !=nil {
		o.Ctx.Application().Logger().Debug("查询订单信息失败")
	}

	return mvc.View{
		Name:"order/view.html",
		Data:iris.Map{
			"order":orderArray,
		},
	}

}

