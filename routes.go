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
		coupon.Post("/",c.NewCoupon)
		coupon.Put("/{id}",c.UpdateCoupon)
		coupon.Get("/",c.GetCoupon)
		coupon.Get("/{code}",c.GetCouponByCode)
	}

	order := r.Party("/v1/api/order",)
	{
		order.Get("/",c.GetOrder)
		order.Get("/{id}",c.GetOrderDetail)
		order.Post("/cart",c.AddToCart)
		order.Get("/paid",c.GetOrderListPending)
		order.Put("/{id}/cart",c.UpdateCart)
		order.Post("/{id}/coupon",c.AddCoupon)
		order.Post("/{id}/payment",c.AddPaymentMethod)
		order.Post("/{id}/payment/profile",c.UpdatePaymentProfile)
		order.Post("/{id}/payment/proof",c.AddPaymentProofed)
		order.Post("/{id}/shipped",c.UpdateShipped)
		order.Post("/{id}/delivered",c.OrderSucceeded)
		order.Post("/{id}/void",c.OrderVoided)
	}

	r.Get("/ping", ping)

	return r
}

func ping(ctx context.Context) {
	ctx.JSON("ping")
}

func notFoundHandler(ctx context.Context) {
	r := c.NewResponse(nil)
	r.AddError(strconv.Itoa(iris.StatusNotFound),"page not found !" ,"we are looking for your API...but we can't find it")
	c.RenderJSON(ctx,r,iris.StatusNotFound)
}

func PrintUrl(ctx context.Context) {
	fmt.Println(ctx.Request().Method,ctx.Request().URL)
	ctx.Next()
}