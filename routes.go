package main

import (
	"github.com/malfanmh/go-shop/internal/controller"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"strconv"
)

func newApp() *iris.Application {
	r := iris.New()

	r.OnErrorCode(iris.StatusNotFound, notFoundHandler)

	product := r.Party("/v1/api/products",)
	{
		product.Post("/",ping)
		product.Put("/{id:integer}",ping)
		product.Get("/",ping)
	}

	coupon := r.Party("/v1/api/coupons",)
	{
		coupon.Post("/",ping)
		coupon.Put("/{id:integer}",ping)
		coupon.Get("/",ping)
	}

	order := r.Party("/v1/api/order",)
	{
		order.Post("/",ping)
		order.Get("/",ping)
		order.Get("/{id:integer}",ping)
		order.Get("/{code:string}",ping)
		order.Get("/{code:string}",ping)
	}

	r.Get("/ping", ping)

	return r
}

func ping(ctx context.Context) {
	ctx.JSON("ping")
}

func notFoundHandler(ctx context.Context) {
	r := controller.NewResponse(nil)
	r.AddError(strconv.Itoa(iris.StatusNotFound),"page not found !" ,"we are looking for your page...but we can't find it")
	controller.RenderJSON(ctx,r,iris.StatusNotFound)
}
