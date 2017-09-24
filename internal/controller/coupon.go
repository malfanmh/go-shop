package controller

import (
	"time"
	"strconv"
	"math/rand"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/malfanmh/go-shop/internal/model"
	"github.com/kataras/iris/core/errors"
	"fmt"
)

type (
	ReqNewCoupon struct {
		Description string	`json:"description"`
		Type		string	`json:"type"`
		Value		float64	`json:"value"`
		Quantity	int		`json:"quantity"`
		ValidAt     string 	`json:"valid_at"`
		ValidUntil  string  `json:"valid_until"`
		TnC         string 	`json:"tnc"`
	}
	ResponseCoupon struct {
		ID          string 		`json:"id"`
		Code        string 		`json:"code"`
		Description string 		`json:"description"`
		Type        string 		`json:"type"`
		Value       float64 	`json:"value"`
		ValidAt     time.Time 	`json:"valid_at"`
		ValidUntil  time.Time   `json:"valid_until"`
		TnC         string 		`json:"tnc,omitempty"`
		Quantity	int			`json:"quantity"`
		Used		int			`json:"used,omitempty"`
		CreatedAt   time.Time   `json:"created_at"`
		CreatedBy   string      `json:"created_by"`
		UpdatedAt   time.Time   `json:"updated_at,omitempty"`
		UpdatedBy   string      `json:"updated_by,omitempty"`
		DeletedAt   time.Time   `json:"deleted_at,omitempty"`
		DeletedBy   string      `json:"deleted_by,omitempty"`
		Status      string      `json:"status"`
	}
)

func NewCoupon(ctx context.Context) {
	r := NewResponse(nil)
	reqNewCoupon := ReqNewCoupon{}
	if err := ctx.ReadJSON(&reqNewCoupon); err != nil {
		r.AddError(strconv.Itoa(iris.StatusBadRequest), "Bad Request !", err.Error())
		RenderJSON(ctx, r, iris.StatusBadRequest)
		return
	}

	c := new(model.Coupon)
	c.SetDescription(reqNewCoupon.Description).
		SetType(reqNewCoupon.Type).
		SetValue(reqNewCoupon.Value).
		SetQuantity(reqNewCoupon.Quantity).
		SetCode(randStr(6,"Alphanumeric")).
		SetTnc(reqNewCoupon.TnC).
		SetValidAt(reqNewCoupon.ValidAt).SetValidUntil(reqNewCoupon.ValidUntil)


	err := c.Insert()
	if err != nil {
		r.AddError(strconv.Itoa(iris.StatusInternalServerError), "Internal Server Error !", err.Error())
		RenderJSON(ctx, r, iris.StatusInternalServerError)
		return
	}

	r = NewResponse(nil)

	RenderJSON(ctx, r, iris.StatusCreated)
}

type ReqUpdateCoupon struct {
	Description string	`json:"description"`
	Type		string	`json:"type"`
	Value		float64	`json:"value"`
	Quantity	int		`json:"quantity"`
	ValidAt     string 	`json:"valid_at"`
	ValidUntil  string  `json:"valid_until"`
	TnC         string 	`json:"tnc"`
}

func UpdateCoupon(ctx context.Context) {
	r := NewResponse(nil)
	id, _ := strconv.Atoi(ctx.Params().Get("id"))
	reqUpdCoupon := ReqUpdateCoupon{}
	if err := ctx.ReadJSON(&reqUpdCoupon); err != nil {
		r.AddError(strconv.Itoa(iris.StatusBadRequest), "Bad Request !", err.Error())
		RenderJSON(ctx, r, iris.StatusBadRequest)
		return
	}

	c := new(model.Coupon)
	c.SetDescription(reqUpdCoupon.Description).
		SetType(reqUpdCoupon.Type).
		SetValue(reqUpdCoupon.Value).
		SetQuantity(reqUpdCoupon.Quantity).
		SetTnc(reqUpdCoupon.TnC).
		SetValidAt(reqUpdCoupon.ValidAt).SetValidUntil(reqUpdCoupon.ValidUntil).
		SetID(id)

	err := c.Update()
	if err != nil {
		r.AddError(strconv.Itoa(iris.StatusInternalServerError), "Internal Server Error !", err.Error())
		RenderJSON(ctx, r, iris.StatusInternalServerError)
		return
	}

	r = NewResponse(nil)
	RenderJSON(ctx, r, iris.StatusCreated)
}


func GetCouponByCode(ctx context.Context) {
	r := NewResponse(nil)
	code := ctx.Params().Get("code")

	coupon, err := model.FindCouponByCode(code)
	if err != nil {
		r.AddError(strconv.Itoa(iris.StatusInternalServerError), "Internal Server Error !", err.Error())
		RenderJSON(ctx, r, iris.StatusInternalServerError)
		return
	}
	va , _ := time.Parse("2006-01-02 15:04:05", coupon.ValidAt)
	vu , _ := time.Parse("2006-01-02 15:04:05", coupon.ValidUntil)

	resCoupon := ResponseCoupon{
		ID:          strconv.Itoa(coupon.ID),
		Code:        coupon.Code,
		Description: coupon.Description.String,
		Type:      	 coupon.Type,
		Value:       coupon.Value,
		Quantity:	 coupon.Quantity,
		Used:		 coupon.Used,
		ValidAt:	 va,
		ValidUntil:	 vu,
		TnC:		 coupon.TnC.String,
		CreatedAt:   coupon.CreatedAt.Time,
		CreatedBy:   coupon.CreatedBy,
		Status:      coupon.Status,
	}

	r = NewResponse(resCoupon)
	RenderJSON(ctx, r, iris.StatusOK)
}

func GetCoupon(ctx context.Context) {
	r := NewResponse(nil)
	q := ctx.URLParam("q")
	page, _ := ctx.URLParamInt("page")
	count, _ := ctx.URLParamInt("count")
	f := new(model.Filter).
		SetQ(q).
		SetPage(page).
		SetCount(count)

	coupon, next, err := model.ListCoupon(f)
	if err != nil {
		r.AddError(strconv.Itoa(iris.StatusInternalServerError), "Internal Server Error !", err.Error())
		RenderJSON(ctx, r, iris.StatusInternalServerError)
		return
	}
	rs := make([]ResponseCoupon, len(coupon))
	for k, v := range coupon {
		va , _ := time.Parse(time.RFC3339, v.ValidAt)
		vu , _ := time.Parse(time.RFC3339Nano, v.ValidUntil)

		rs[k].ID = strconv.Itoa(v.ID)
		rs[k].Code = v.Code
		rs[k].Description = v.Description.String
		rs[k].Type = v.Type
		rs[k].Value = v.Value
		rs[k].Quantity = v.Quantity
		rs[k].ValidAt = va
		rs[k].ValidUntil = vu
		rs[k].CreatedAt = v.CreatedAt.Time
		rs[k].CreatedBy = v.CreatedBy
		rs[k].Status = v.Status
	}

	r = NewResponse(rs)
	r.SetPagination(ctx, f.Page, next)
	RenderJSON(ctx, r, iris.StatusOK)
}


func randStr(ln int, fm string) string {
	CharsType := map[string]string{
		"Alphabet":     model.ALPHABET,
		"Numerals":     model.NUMERALS,
		"Alphanumeric": model.ALPHANUMERIC,
	}

	rand.Seed(time.Now().UTC().UnixNano())
	chars := CharsType[fm]
	result := make([]byte, ln)
	for i := 0; i < ln; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

func CheckCoupon(code string) (model.Coupon,error){
	cp , err := model.FindCouponByCode(code)
	if err != nil {
		return model.Coupon{}, err
	}
	if cp.Quantity -1 < 1 {
		return model.Coupon{}, errors.New("out of stock")
	}

	va , _ := time.Parse("2006-01-02 15:04:05", cp.ValidAt)
	vu , _ := time.Parse("2006-01-02 15:04:05", cp.ValidUntil)
	fmt.Println("periode",va,time.Now(),vu)
	if !va.Before(time.Now()) || !vu.After(time.Now()) {
		return model.Coupon{}, errors.New("out of date")
	}

	return cp ,nil
}