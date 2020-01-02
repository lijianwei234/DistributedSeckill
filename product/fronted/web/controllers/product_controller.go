package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"html/template"
	"os"
	"path/filepath"
	"product/datamodels"
	"product/services"
	"strconv"
)

type ProductController struct {
	Ctx iris.Context
	ProductService services.IProductService
	Session *sessions.Session
}

var (
	htmlOutPath = "./fronted/web/htmlProductShow/"
	templatePath = "./fronted/web/template/"
)

func (p *ProductController) GetGenerateHtml() {
	productString := p.Ctx.URLParam("productID")
	productID,err:=strconv.Atoi(productString)
	//获取模板文件地址
	contenstTmp, err := template.ParseFiles(filepath.Join(templatePath), "product.html")
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	//获取html生成路径
	fileName := filepath.Join(htmlOutPath, "htmlProduct.html")
	//获取模板渲染数据
	product,err:=p.ProductService.GetProductByID(int64(productID))
	if err !=nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	//4.生成静态文件
	generateStaticHtml(p.Ctx,contenstTmp,fileName,product)
}

//生成html静态文件
func generateStaticHtml(ctx iris.Context,template *template.Template,fileName string,product *datamodels.Product)  {
	//1.判断静态文件是否存在
	if exist(fileName) {
		err:=os.Remove(fileName)
		if err !=nil {
			ctx.Application().Logger().Error(err)
		}
	}
	//2.生成静态文件
	file,err := os.OpenFile(fileName,os.O_CREATE|os.O_WRONLY,os.ModePerm)
	if err !=nil {
		ctx.Application().Logger().Error(err)
	}
	defer file.Close()
	template.Execute(file,&product)
}

func exist (fileName string) bool {
	_,err := os.Stat(fileName)
	return err == nil || os.IsExist(err)
}

func (p *ProductController) GetDetail() mvc.View {
	product, err := p.ProductService.GetProductByID(1)
	if err != nil {
		p.Ctx.Application().Logger().Error(err)
	}
	return mvc.View{
		Layout: "shared/productLayout.html",
		Name:   "product/view.html",
		Data: iris.Map{
			"product": product,
		},
	}
}
