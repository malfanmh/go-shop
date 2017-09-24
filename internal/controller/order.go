package controller

import (
	"strconv"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/malfanmh/go-shop/internal/model"
	"time"
	"fmt"
)

type (
	ResponseOrderList struct {
		ID          	int				`json:"id"`
		Code        	string 			`json:"code"`
		Coupon 			string			`json:"coupon"`
		PaymentType		string			`json:"payment_type"`
		CustomerID		int				`json:"customer_id,omitempty"`
		CustomerProfile	CustomerProfile `json:"customer,omitempty"`
		ShippingCode	string			`json:"shipping_code"`
		SubTotal		float64			`json:"sub_total"`
		ShipmentCost	float64			`json:"shipment_cost"`
		DiscountValue	float64			`json:"discount_value"`
		TotalAmount		float64			`json:"total_amount"`
		State 			string			`json:"state"`
		Details			OrderDetailList	`json:"details,omitempty"`
		CreatedAt   	time.Time 		`json:"created_at"`
		CreatedBy   	string        	`json:"created_by"`
		DeletedAt   	time.Time  		`json:"deleted_at,omitempty"`
		DeletedBy   	string  		`json:"deleted_by,omitempty"`
		Status      	string        	`json:"status"`
	}
	OrderDetailList []OrderDetailData
	OrderDetailData struct {
		ID			int 			`json:"id"`
		OrderID 	int 			`json:"order_id"`
		Product		ProductData		`json:"product_id"`
		Price 		float64 		`json:"price"`
		Quantity	int 			`json:"quantity"`
		Amount 		float64			`json:"amount"`
	}
	ProductData struct {
		Name        string    `json:"name"`
		Description string    `json:"description"`
	}
)
func GetOrder(ctx context.Context){
	r := NewResponse(nil)
	q := ctx.URLParam("q")
	page, _ := ctx.URLParamInt("page")
	count, _ := ctx.URLParamInt("count")
	f := new(model.Filter).
		SetQ(q).
		SetPage(page).
		SetCount(count)

	order, next, err := model.ListOrder(f)
	if err != nil {
		r.AddError(strconv.Itoa(iris.StatusInternalServerError), "Internal Server Error !", err.Error())
		RenderJSON(ctx, r, iris.StatusInternalServerError)
		return
	}
	respOrder := make([]ResponseOrderList, len(order))
	for k , v := range order{
		respOrder[k].ID = v.ID
		respOrder[k].Code = v.Code.String
		respOrder[k].Coupon = v.CouponCode.String
		respOrder[k].CustomerID = v.CustomerID
		respOrder[k].PaymentType = v.PaymentType.String
		respOrder[k].ShippingCode =v.ShippingCode.String
		respOrder[k].SubTotal = v.SubTotal
		respOrder[k].ShipmentCost = v.ShipmentCost
		respOrder[k].DiscountValue = v.DiscountValue
		respOrder[k].TotalAmount = v.TotalAmount
		respOrder[k].State = v.State
		respOrder[k].CreatedAt = v.CreatedAt.Time
		respOrder[k].CreatedBy = v.CreatedBy
		respOrder[k].State = v.State
		respOrder[k].Status = v.Status
	}


	r = NewResponse(respOrder)
	r.SetPagination(ctx, f.Page, next)
	RenderJSON(ctx, r, iris.StatusOK)
}

func GetOrderListPending(ctx context.Context){
	r := NewResponse(nil)
	q := ctx.URLParam("q")
	page, _ := ctx.URLParamInt("page")
	count, _ := ctx.URLParamInt("count")
	f := new(model.Filter).
		SetQ(q).
		SetPage(page).
		SetCount(count)

	order, next, err := model.ListOrder(f)
	if err != nil {
		r.AddError(strconv.Itoa(iris.StatusInternalServerError), "Internal Server Error !", err.Error())
		RenderJSON(ctx, r, iris.StatusInternalServerError)
		return
	}
	respOrder := make([]ResponseOrderList, len(order))
	for k , v := range order{
		respOrder[k].ID = v.ID
		respOrder[k].Code = v.Code.String
		respOrder[k].Coupon = v.CouponCode.String
		respOrder[k].CustomerID = v.CustomerID
		respOrder[k].PaymentType = v.PaymentType.String
		respOrder[k].ShippingCode =v.ShippingCode.String
		respOrder[k].SubTotal = v.SubTotal
		respOrder[k].ShipmentCost = v.ShipmentCost
		respOrder[k].DiscountValue = v.DiscountValue
		respOrder[k].TotalAmount = v.TotalAmount
		respOrder[k].State = v.State
		respOrder[k].CreatedAt = v.CreatedAt.Time
		respOrder[k].CreatedBy = v.CreatedBy
		respOrder[k].State = v.State
		respOrder[k].Status = v.Status
	}


	r = NewResponse(respOrder)
	r.SetPagination(ctx, f.Page, next)
	RenderJSON(ctx, r, iris.StatusOK)
}

func GetOrderDetail(ctx context.Context){
	r := NewResponse(nil)
	id := ctx.Params().Get("id")

	o, err := model.FindOrderByID(id)
	if err != nil {
		r.AddError(strconv.Itoa(iris.StatusInternalServerError), "Internal Server Error !", err.Error())
		RenderJSON(ctx, r, iris.StatusInternalServerError)
		return
	}

	respOrder := make([]ResponseOrderList, len(o))
	for k,v := range o {
		respOrder[k].ID = v.ID
		respOrder[k].Code = v.Code.String
		respOrder[k].Coupon = v.CouponCode.String
		respOrder[k].CustomerID = v.CustomerID
		respOrder[k].PaymentType = v.PaymentType.String
		respOrder[k].ShippingCode =v.ShippingCode.String
		respOrder[k].SubTotal = v.SubTotal
		respOrder[k].ShipmentCost = v.ShipmentCost
		respOrder[k].DiscountValue = v.DiscountValue
		respOrder[k].TotalAmount = v.TotalAmount
		respOrder[k].State = v.State
		respOrder[k].CreatedAt = v.CreatedAt.Time
		respOrder[k].CreatedBy = v.CreatedBy
		respOrder[k].State = v.State
		respOrder[k].Status = v.Status

		cust , _ := model.FindCustomerByID(v.CustomerID)
		respOrder[k].CustomerProfile = CustomerProfile{Name:cust.Name,Phone:cust.Phone,Email:cust.Email,Address:cust.Address}

		orderDetails , _ := model.FindOrderDetailByOrderID(strconv.Itoa(o[0].ID))
		respOrder[k].Details = make(OrderDetailList, len(orderDetails))
		for k1,v1 := range orderDetails{
			respOrder[k].Details[k1].ID = v1.ID
			respOrder[k].Details[k1].Price = v1.Price
			respOrder[k].Details[k1].Quantity = v1.Quantity
			respOrder[k].Details[k1].Amount = v1.Amount

			p , _ := model.FindProductByID(strconv.Itoa(v1.ProductID))
			respOrder[k].Details[k1].Product.Name = p.Name
			respOrder[k].Details[k1].Product.Description = p.Description.String
		}

	}


	r = NewResponse(respOrder)
	RenderJSON(ctx, r, iris.StatusOK)
}

type (
	OrderItems struct {
		ID 			int `json:"id"`
		Quantity	int `json:"quantity"`
	}
	CustomerProfile struct {
		Address string `json:"address"`
		Email   string `json:"email"`
		Name    string `json:"name"`
		Phone   string `json:"phone"`
	}
	RequestOrder struct {
		CustomerProfile 	CustomerProfile `json:"profile"`
		OrderItems 			[]OrderItems 	`json:"order_items"`
		PaymentType 		string 			`json:"payment_type"`
		Coupon				string			`json:"coupon"`
		ShippingCode		string			`json:"shipping_code"`
		Code				string			`json:"order_id"`
		ImgURL				string			`json:"img_url"`
	}
	ResponseOrder struct {
		Code string `json:"order_id"`
	}
)
func AddToCart(ctx context.Context){
	r := NewResponse(nil)
	order := RequestOrder{}
	if err := ctx.ReadJSON(&order); err != nil {
		r.AddError(strconv.Itoa(iris.StatusBadRequest), "Bad Request !", err.Error())
		RenderJSON(ctx, r, iris.StatusBadRequest)
		return
	}
	//initial order
	o := new(model.Order).
		SetCode(randStr(16,"Numerals"))
	err := o.Insert()
	if err != nil {
		r.AddError(strconv.Itoa(iris.StatusInternalServerError), "Internal Server Error !", err.Error() )
		RenderJSON(ctx, r, iris.StatusInternalServerError)
		return
	}

	for _, v := range order.OrderItems{

		product , err := model.FindProductByID(strconv.Itoa(v.ID))
		fmt.Println("a :",v.ID, err)
		if err != nil {
			r.AddError(strconv.Itoa(iris.StatusInternalServerError), "Internal Server Error !", err.Error())
			RenderJSON(ctx, r, iris.StatusInternalServerError)
			return
		}
		amount := float64(v.Quantity) * product.Price
		od := new(model.OrderDetails).
			SetPrice(product.Price).
			SetProductID(product.ID).
			SetQuantity(v.Quantity).
			SetOrderCode(o.Code.String).
			SetAmount(amount)
		err = od.Insert()
		if err != nil {
			r.AddError(strconv.Itoa(iris.StatusInternalServerError), "Internal Server Error !", err.Error())
			RenderJSON(ctx, r, iris.StatusInternalServerError)
			return
		}
	}

	r = NewResponse(ResponseOrder{Code:o.Code.String})
	RenderJSON(ctx, r, iris.StatusCreated)
}

func UpdateCart(ctx context.Context){
	r := NewResponse(nil)
	id := ctx.Params().Get("id")
	order := RequestOrder{}
	if err := ctx.ReadJSON(&order); err != nil {
		r.AddError(strconv.Itoa(iris.StatusBadRequest), "Bad Request !", err.Error())
		RenderJSON(ctx, r, iris.StatusBadRequest)
		return
	}
	//initial order
	for _, v := range order.OrderItems{
		product , err := model.FindProductByID(strconv.Itoa(v.ID))
		if err != nil {
			r.AddError(strconv.Itoa(iris.StatusInternalServerError), "Internal Server Error !", err.Error())
			RenderJSON(ctx, r, iris.StatusInternalServerError)
			return
		}
		amount := float64(v.Quantity) * product.Price
		od := new(model.OrderDetails).
			SetPrice(product.Price).
			SetProductID(product.ID).
			SetQuantity(v.Quantity).
			SetOrderCode(id).
			SetAmount(amount)
		err = od.Insert()
		if err != nil {
			r.AddError(strconv.Itoa(iris.StatusInternalServerError), "Internal Server Error !", err.Error())
			RenderJSON(ctx, r, iris.StatusInternalServerError)
			return
		}
	}

	r = NewResponse(ResponseOrder{Code:id})
	RenderJSON(ctx, r, iris.StatusCreated)
}

func AddCoupon(ctx context.Context){
	r := NewResponse(nil)
	id := ctx.Params().Get("id")
	order := RequestOrder{}
	if err := ctx.ReadJSON(&order); err != nil {
		r.AddError(strconv.Itoa(iris.StatusBadRequest), "Bad Request !", err.Error())
		RenderJSON(ctx, r, iris.StatusBadRequest)
		return
	}
	co , _ := model.FindOrderByID(id)
	o := new(model.Order).
		SetID(co[0].ID)
	param := make(map[string]interface{})
	param["coupon"] = order.Coupon
	err := o.Update(param)
	if err != nil {
		r.AddError(strconv.Itoa(iris.StatusInternalServerError), "Internal Server Error !", err.Error())
		RenderJSON(ctx, r, iris.StatusInternalServerError)
		return
	}
	r = NewResponse(ResponseOrder{Code:id})
	RenderJSON(ctx, r, iris.StatusCreated)
}
func AddPaymentMethod(ctx context.Context){
	r := NewResponse(nil)
	id := ctx.Params().Get("id")
	order := RequestOrder{}
	if err := ctx.ReadJSON(&order); err != nil {
		r.AddError(strconv.Itoa(iris.StatusBadRequest), "Bad Request !", err.Error())
		RenderJSON(ctx, r, iris.StatusBadRequest)
		return
	}
	co , _ := model.FindOrderByID(id)
	o := new(model.Order).
		SetID(co[0].ID)
	param := make(map[string]interface{})
	param["payment_type"] = model.OrderPaymentTypeBankTf
	err := o.Update(param)
	if err != nil {
		r.AddError(strconv.Itoa(iris.StatusInternalServerError), "Internal Server Error !", err.Error())
		RenderJSON(ctx, r, iris.StatusInternalServerError)
		return
	}
	r = NewResponse(ResponseOrder{Code:id})
	RenderJSON(ctx, r, iris.StatusCreated)
}

func UpdatePaymentProfile(ctx context.Context){
	r := NewResponse(nil)
	id := ctx.Params().Get("id")
	order := RequestOrder{}
	if err := ctx.ReadJSON(&order); err != nil {
		r.AddError(strconv.Itoa(iris.StatusBadRequest), "Bad Request !", err.Error())
		RenderJSON(ctx, r, iris.StatusBadRequest)
		return
	}

	//initial profile
	cp := order.CustomerProfile
	cust := new(model.Customers).
		SetName(cp.Name).
		SetPhone(cp.Phone).
		SetEmail(cp.Email).
		SetAddress(cp.Address)
	custID , err := cust.Insert()
	if err != nil {
		r.AddError(strconv.Itoa(iris.StatusInternalServerError), "Internal Server Error !", err.Error())
		RenderJSON(ctx, r, iris.StatusInternalServerError)
		return
	}
	orderData , err := model.FindOrderByID(id)
	if err != nil {
		r.AddError(strconv.Itoa(iris.StatusInternalServerError), "Internal Server Error !", err.Error())
		RenderJSON(ctx, r, iris.StatusInternalServerError)
		return
	}

	o := model.NewOrder(orderData[0])
	subTotal := float64(0)
	od , err := model.FindOrderDetailByOrderID(id)
	if err != nil {
		r.AddError(strconv.Itoa(iris.StatusInternalServerError), "Internal Server Error !", err.Error())
		RenderJSON(ctx, r, iris.StatusInternalServerError)
		return
	}else {
		for _, v := range od{
			subTotal += v.Amount
		}

	}

	discountValue := float64(0)
	//check coupon return "discount_value"
	if !o.CouponCode.Valid{
		c , err := CheckCoupon(o.CouponCode.String)
		if err != nil {
			r.AddError(strconv.Itoa(iris.StatusBadRequest), "Bad Request !", err.Error())
			RenderJSON(ctx, r, iris.StatusBadRequest)
			return
		}else {
			switch c.Type {
			case model.CouponTypeAmount:
				discountValue = c.Value
			case model.CouponTypePercendtage:
				discountValue = (subTotal / float64(100)) *  c.Value
			}
		}
	}
	fmt.Println("coupon :",orderData[0].CouponCode)
	totalAmount := subTotal - (model.ShipmentCost + discountValue)
	o.SetCustomerID(int(custID)).
		SetSubTotal(subTotal).
		SetShipmentCost(model.ShipmentCost).
		SetDiscountValue(discountValue).
		SetTotalAmount(totalAmount).
		SetState(model.OrderStatePending).
		SetCode(id)
		//SetCoupon()
	err = o.InsertProfile()
	if err != nil {
		r.AddError(strconv.Itoa(iris.StatusInternalServerError), "Internal Server Error !", err.Error())
		RenderJSON(ctx, r, iris.StatusInternalServerError)
		return
	}

	r = NewResponse(ResponseOrder{Code:id})
	RenderJSON(ctx, r, iris.StatusCreated)
}

func AddPaymentProofed(ctx context.Context){
	r := NewResponse(nil)
	id := ctx.Params().Get("id")
	order := RequestOrder{}
	if err := ctx.ReadJSON(&order); err != nil {
		r.AddError(strconv.Itoa(iris.StatusBadRequest), "Bad Request !", err.Error())
		RenderJSON(ctx, r, iris.StatusBadRequest)
		return
	}
	orderData , err := model.FindOrderByID(id)
	if err != nil {
		r.AddError(strconv.Itoa(iris.StatusInternalServerError), "Internal Server Error !", err.Error())
		RenderJSON(ctx, r, iris.StatusInternalServerError)
		return
	}
	o := model.NewOrder(orderData[0])
	o.SetPaymentProof(true).
		SetState(model.OrderStatePaid)
	err = o.InsertProof()
	if err != nil {
		r.AddError(strconv.Itoa(iris.StatusInternalServerError), "Internal Server Error !", err.Error())
		RenderJSON(ctx, r, iris.StatusInternalServerError)
		return
	}
	r = NewResponse(ResponseOrder{Code:id})
	RenderJSON(ctx, r, iris.StatusCreated)
}

func UpdateShipped(ctx context.Context){
	r := NewResponse(nil)
	id := ctx.Params().Get("id")
	order := RequestOrder{}
	if err := ctx.ReadJSON(&order); err != nil {
		r.AddError(strconv.Itoa(iris.StatusBadRequest), "Bad Request !", err.Error())
		RenderJSON(ctx, r, iris.StatusBadRequest)
		return
	}
	orderData , err := model.FindOrderByID(id)
	if err != nil {
		r.AddError(strconv.Itoa(iris.StatusInternalServerError), "Internal Server Error !", err.Error())
		RenderJSON(ctx, r, iris.StatusInternalServerError)
		return
	}
	o := model.NewOrder(orderData[0])
	o.SetShippingCode(order.ShippingCode).
		SetState(model.OrderStateShipped)
	err = o.InsertShipped()
	if err != nil {
		r.AddError(strconv.Itoa(iris.StatusInternalServerError), "Internal Server Error !", err.Error())
		RenderJSON(ctx, r, iris.StatusInternalServerError)
		return
	}
	r = NewResponse(ResponseOrder{Code:id})
	RenderJSON(ctx, r, iris.StatusCreated)
}

func OrderSucceeded(ctx context.Context){
	r := NewResponse(nil)
	id := ctx.Params().Get("id")
	order := RequestOrder{}
	if err := ctx.ReadJSON(&order); err != nil {
		r.AddError(strconv.Itoa(iris.StatusBadRequest), "Bad Request !", err.Error())
		RenderJSON(ctx, r, iris.StatusBadRequest)
		return
	}
	orderData , err := model.FindOrderByID(id)
	if err != nil {
		r.AddError(strconv.Itoa(iris.StatusInternalServerError), "Internal Server Error !", err.Error())
		RenderJSON(ctx, r, iris.StatusInternalServerError)
		return
	}
	o := model.NewOrder(orderData[0])
	o.SetState(model.OrderStateSuccess)
	err = o.InsertSucceeded()
	if err != nil {
		r.AddError(strconv.Itoa(iris.StatusInternalServerError), "Internal Server Error !", err.Error())
		RenderJSON(ctx, r, iris.StatusInternalServerError)
		return
	}
	r = NewResponse(ResponseOrder{Code:id})
	RenderJSON(ctx, r, iris.StatusCreated)
}
func OrderVoided(ctx context.Context){
	r := NewResponse(nil)
	id := ctx.Params().Get("id")
	order := RequestOrder{}
	if err := ctx.ReadJSON(&order); err != nil {
		r.AddError(strconv.Itoa(iris.StatusBadRequest), "Bad Request !", err.Error())
		RenderJSON(ctx, r, iris.StatusBadRequest)
		return
	}
	orderData , err := model.FindOrderByID(id)
	if err != nil {
		r.AddError(strconv.Itoa(iris.StatusInternalServerError), "Internal Server Error !", err.Error())
		RenderJSON(ctx, r, iris.StatusInternalServerError)
		return
	}
	o := model.NewOrder(orderData[0])
	o.SetState(model.OrderStateCanceled)
	err = o.InsertVoided()
	if err != nil {
		r.AddError(strconv.Itoa(iris.StatusInternalServerError), "Internal Server Error !", err.Error())
		RenderJSON(ctx, r, iris.StatusInternalServerError)
		return
	}
	r = NewResponse(ResponseOrder{Code:id})
	RenderJSON(ctx, r, iris.StatusCreated)
}

