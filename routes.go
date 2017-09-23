package main

import (
	"strconv"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	c "github.com/malfanmh/go-shop/internal/controller"
	"fmt"
)

func newApp() *iris.Application {
	r := iris.New()

	r.OnErrorCode(iris.StatusNotFound, notFoundHandler)

	product := r.Party("/v1/api/products",PrintUrl)
	{
		product.Post("/",c.NewProduct)
		product.Put("/{id}",c.UpdateProduct)
		product.Get("/",c.GetProduct)
		product.Get("/{id}",c.GetProductByID)
	}

	coupon := r.Party("/v1/api/coupons",)
	{
		coupon.Post("/",ping)
		coupon.Put("/{id}",ping)
		coupon.Get("/",ping)
	}

	order := r.Party("/v1/api/order",)
	{
		order.Post("/",ping)
		order.Get("/",ping)
		order.Get("/{id}",ping)
	}

	r.Get("/ping", ping)

	return r
}

func ping(ctx context.Context) {
	ctx.JSON("ping")
}

func notFoundHandler(ctx context.Context) {
	r := c.NewResponse(nil)
	r.AddError(strconv.Itoa(iris.StatusNotFound),"page not found !" ,"we are looking for your page...but we can't find it")
	c.RenderJSON(ctx,r,iris.StatusNotFound)
}

func PrintUrl(ctx context.Context) {
	fmt.Println(ctx.Request().Method,ctx.Request().URL)
	ctx.Next()
}