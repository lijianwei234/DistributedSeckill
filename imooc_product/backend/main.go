package main

import (
	"context"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/opentracing/opentracing-go/log"
	"imooc_product/backend/web/controllers"
	"imooc_product/common"
	"imooc_product/repositories"
	"imooc_product/services"
)

func main()  {
	//1.创建iris 实例
	app:=iris.New()
	//2.设置错误模式，在mvc模式下提示错误
	app.Logger().SetLevel("debug" )
	//3.注册模板
	tmplate := iris.HTML("./backend/web/views",".html").Layout("shared/layout.html").Reload(true)
	app.RegisterView(tmplate)
	//4.设置模板目标
	app.HandleDir("/assets","./backend/web/assets")
	//出现异常跳转到指定页面
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message",ctx.Values().GetStringDefault("message","访问的页面出错！"))
		ctx.ViewLayout("")
		ctx.View("shared/error.html")
	})
	//连接数据库
	db,err :=common.NewMysqlConn()
	if err !=nil {
		log.Error(err)
	}
	ctx,cancel := context.WithCancel(context.Background())
	defer cancel()

	//5.注册控制器
	productRepository := repositories.NewProductManager("product",db)
	productSerivce :=services.NewProductService(productRepository)
	productParty := app.Party("/imooc_product")
	product := mvc.New(productParty)
	product.Register(ctx,productSerivce)
	product.Handle(new(controllers.ProductController))


	//6.启动服务
	app.Run(
		iris.Addr("localhost:8080"),
		//iris.WithoutVersionChecker,
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}