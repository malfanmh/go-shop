package controller

import (
	"github.com/kataras/iris/context"
	"strconv"
	"github.com/kataras/iris"
	"time"
)

type (
	RequestNewProduct struct {
		Name        string 	`json:"name"`
		Description string 	`json:"description"`
		ImgURL      string 	`json:"img_url"`
		Price       float64 `json:"price"`
		Quantity    int 	`json:"quantity"`
	}
	ResponseNewProduct struct {
		ID          int    		`json:"id"`
		Name        string 		`json:"name"`
		Description string 		`json:"description"`
		ImgURL      string 		`json:"img_url"`
		Price       float64 	`json:"price"`
		Quantity    int			`json:"quantity"`
		CreatedAt   time.Time   `json:"created_at"`
		CreatedBy   string    	`json:"created_by"`
		UpdatedAt   time.Time   `json:"created_at,omitempty"`
		UpdatedBy   string    	`json:"created_by,omitempty"`
		DeletedAt   time.Time   `json:"created_at,omitempty"`
		DeletedBy   string    	`json:"created_by,omitempty"`
		Status      string    	`json:"status"`
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

	r = NewResponse(nil)
	RenderJSON(ctx, r, iris.StatusCreated)
}
