package controller

import (
	"time"
	"strconv"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/malfanmh/go-shop/internal/model"
)

type (
	RequestNewProduct struct {
		Name        string    `json:"name"`
		Description string    `json:"description"`
		ImgURL      string    `json:"img_url"`
		Price       float64   `json:"price"`
		Quantity    int    	  `json:"quantity"`
	}
	ResponseProduct struct {
		ID          int           `json:"id"`
		Name        string        `json:"name"`
		Description string        `json:"description"`
		ImgURL      string        `json:"img_url"`
		Price       float64    	  `json:"price"`
		Quantity    int           `json:"quantity"`
		CreatedAt   time.Time   `json:"created_at"`
		CreatedBy   string        `json:"created_by"`
		UpdatedAt   time.Time   `json:"created_at,omitempty"`
		UpdatedBy   string        `json:"created_by,omitempty"`
		DeletedAt   time.Time   `json:"created_at,omitempty"`
		DeletedBy   string        `json:"created_by,omitempty"`
		Status      string        `json:"status"`
	}
)

func NewProduct(ctx context.Context) {
	r := NewResponse(nil)
	reqNewProduct := RequestNewProduct{}
	if err := ctx.ReadJSON(&reqNewProduct); err != nil {
		r.AddError(strconv.Itoa(iris.StatusBadRequest), "Bad Request !", err.Error())
		RenderJSON(ctx, r, iris.StatusBadRequest)
		return
	}

	p := new(model.Products)
	p.SetName(reqNewProduct.Name).
		SetDescription(reqNewProduct.Description).
		SetImgURL(reqNewProduct.ImgURL).
		SetPrice(reqNewProduct.Price).
		SetQuantity(reqNewProduct.Quantity)

	err := p.Insert()
	if err != nil {
		r.AddError(strconv.Itoa(iris.StatusInternalServerError), "Internal Server Error !", err.Error())
		RenderJSON(ctx, r, iris.StatusBadRequest)
		return
	}

	r = NewResponse(nil)
	RenderJSON(ctx, r, iris.StatusCreated)
}

type RequestUpdateProduct struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ImgURL      string    `json:"img_url"`
	Price       float64 `json:"price"`
	Quantity    int    `json:"quantity"`
}

func UpdateProduct(ctx context.Context) {
	r := NewResponse(nil)
	reqNewProduct := RequestUpdateProduct{}
	id, _ := strconv.Atoi(ctx.Params().Get("id"))
	if err := ctx.ReadJSON(&reqNewProduct); err != nil {
		r.AddError(strconv.Itoa(iris.StatusBadRequest), "Bad Request !", err.Error())
		RenderJSON(ctx, r, iris.StatusBadRequest)
		return
	}

	p := new(model.Products)
	p.SetID(id).
		SetName(reqNewProduct.Name).
		SetDescription(reqNewProduct.Description).
		SetImgURL(reqNewProduct.ImgURL).
		SetPrice(reqNewProduct.Price).
		SetQuantity(reqNewProduct.Quantity).
		SetUpdatedBy("system").
		SetUpdatedAt(time.Now())

	err := p.Update()
	if err != nil {
		r.AddError(strconv.Itoa(iris.StatusInternalServerError), "Internal Server Error !", err.Error())
		RenderJSON(ctx, r, iris.StatusInternalServerError)
		return
	}

	r = NewResponse(nil)
	RenderJSON(ctx, r, iris.StatusCreated)
}

func GetProductByID(ctx context.Context) {
	r := NewResponse(nil)
	id := ctx.Params().Get("id")

	product, err := model.FindProductByID(id)
	if err != nil {
		r.AddError(strconv.Itoa(iris.StatusInternalServerError), "Internal Server Error !", err.Error())
		RenderJSON(ctx, r, iris.StatusInternalServerError)
		return
	}

	resProduct := ResponseProduct{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description.String,
		ImgURL:      product.ImgURL.String,
		Price:       product.Price,
		CreatedAt:   product.CreatedAt.Time,
		CreatedBy:   product.CreatedBy,
		Status:      product.Status,
	}

	r = NewResponse(resProduct)
	RenderJSON(ctx, r, iris.StatusOK)
}

func GetProduct(ctx context.Context) {
	r := NewResponse(nil)
	q := ctx.URLParam("q")
	page, _ := ctx.URLParamInt("page")
	count, _ := ctx.URLParamInt("count")
	f := new(model.Filter).
		SetQ(q).
		SetPage(page).
		SetCount(count)

	product, next, err := model.ListProduct(f)
	if err != nil {
		r.AddError(strconv.Itoa(iris.StatusInternalServerError), "Internal Server Error !", err.Error())
		RenderJSON(ctx, r, iris.StatusInternalServerError)
		return
	}

	rs := make([]ResponseProduct, len(product))
	for k, v := range product {
		rs[k].ID = v.ID
		rs[k].Name = v.Name
		rs[k].Description = v.Description.String
		rs[k].ImgURL = v.ImgURL.String
		rs[k].Price = v.Price
		rs[k].Quantity = v.Quantity
		rs[k].CreatedAt = v.CreatedAt.Time
		rs[k].CreatedBy = v.CreatedBy
		rs[k].Status = v.Status
	}

	r = NewResponse(rs)
	r.SetPagination(ctx, f.Page, next)
	RenderJSON(ctx, r, iris.StatusOK)
}
