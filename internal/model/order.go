package model

import (
	"github.com/go-sql-driver/mysql"
	"database/sql"
	"strconv"
)

type(
	Order struct {
		ID          	int				`db:"id"`
		Code        	sql.NullString	`db:"code"`
		CouponCode		sql.NullString	`db:"coupon"`
		PaymentType		sql.NullString	`db:"payment_type"`
		PaymentProof	bool			`db:"payment_proof"`
		CustomerID		int				`db:"customer_id"`
		ShippingCode	sql.NullString	`db:"shipping_code"`
		SubTotal		float64			`db:"sub_total"`
		ShipmentCost	float64			`db:"shipment_cost"`
		DiscountValue	float64			`db:"discount_value"`
		TotalAmount		float64			`db:"total_amount"`
		State 			string			`db:"state"`
		CreatedAt   	mysql.NullTime 	`db:"created_at"`
		CreatedBy   	string        	`db:"created_by"`
		DeletedAt   	mysql.NullTime  `db:"created_at,omitempty"`
		DeletedBy   	sql.NullString  `db:"created_by,omitempty"`
		Status      	string        	`db:"status"`
	}
)

func FindOrderByID(id string) ([]Order , error){
	return findOrder("code",id)
}
func findOrder(field, value string) ([]Order , error){
	q := `
		SELECT
			id
			, code
			, coupon
			, payment_type
			, customer_id
			, shipping_code
			, sub_total
			, shipment_cost
			, discount_value
			, total_amount
			, state
			, created_at
			, created_by
			, deleted_at
			, deleted_by
			, status
		FROM
			orders
		WHERE
			`+field+` = ?
		AND
			status = ?
		ORDER BY
			id desc
	`

	var orders []Order
	rows , err := db.Query(q, value, StatusCreated)
	if err != nil {
		return []Order{} , err

	}

	defer rows.Close()
	for rows.Next() {
		o := new(Order)
		err := rows.Scan(&o.ID,&o.Code,&o.CouponCode,&o.PaymentType,&o.CustomerID,&o.ShippingCode,&o.SubTotal,&o.ShipmentCost,&o.DiscountValue,&o.TotalAmount,&o.State,
			&o.CreatedAt,&o.CreatedBy,&o.DeletedAt,&o.DeletedBy,&o.Status)
		if err != nil {
			return []Order{} , err
		}
		orders = append(orders, *o)
	}
	if len(orders) < 1  {
		return []Order{} , err
	}
	return orders , nil
}

func ListOrder(f *Filter)([]Order, bool , error){
	if f.Page < 1 {
		f.Page = 1
	}
	if f.Count < 1 {
		f.Count = 1
	}

	q := `
		SELECT
			id
			, code
			, coupon
			, payment_type
			, customer_id
			, shipping_code
			, sub_total
			, shipment_cost
			, discount_value
			, total_amount
			, state
			, created_at
			, created_by
			, deleted_at
			, deleted_by
			, status
		FROM
			orders
		WHERE
			status = ?

		ORDER BY
			id desc
		LIMIT ? OFFSET ?

	`

	var orders []Order
	rows , err := db.Query(q,  StatusCreated, f.Count+1, (f.Page-1)*f.Count)
	if err != nil {
		return []Order{} ,false, err

	}

	defer rows.Close()
	for rows.Next() {
		o := new(Order)
		err := rows.Scan(&o.ID,&o.Code,&o.CouponCode,&o.PaymentType,&o.CustomerID,&o.ShippingCode,&o.SubTotal,&o.ShipmentCost,&o.DiscountValue,&o.TotalAmount,&o.State,
			&o.CreatedAt,&o.CreatedBy,&o.DeletedAt,&o.DeletedBy,&o.Status)
		if err != nil {
			return []Order{} ,false, err
		}
		orders = append(orders, *o)
	}
	if len(orders) < 1  {
		return []Order{} ,false, ErrResourceNotFound
	}

	next := false
	if len(orders) > f.Count {
		next = true
	}
	if len(orders) < f.Count {
		f.Count = len(orders)
	}

	return orders,next, nil
}

func (o *Order) Insert() error{
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	q := `
		INSERT INTO
			orders
				(
					code
					, shipment_cost
					, state
				)
		VALUES
			(?, ?, ?)
	`
	stmt, err := tx.Prepare(q)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(o.Code,ShipmentCost, OrderStateCreated)
	if err != nil {
		return err
	}
	defer stmt.Close()

	rowCnt, err := res.RowsAffected()
	if err != nil {
		return err
	}else if rowCnt < 1 {
		return ErrNotModified
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	//o.SetID(int(id))
	return nil
}
func (o *Order) InsertProfile() error{
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	q := `
		INSERT INTO
			orders
				(
					code
					, shipment_cost
					, state
					, customer_id
					, sub_total
					, discount_value
					, total_amount
					, coupon
				)
		VALUES
			(?, ?, ?, ?, ?, ?, ?,? )
	`
	stmt, err := tx.Prepare(q)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(o.Code,ShipmentCost, o.State,o.CustomerID,o.SubTotal,o.DiscountValue,o.TotalAmount , o.CouponCode.String)
	if err != nil {
		return err
	}
	defer stmt.Close()

	rowCnt, err := res.RowsAffected()
	if err != nil {
		return err
	}else if rowCnt < 1 {
		return ErrNotModified
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
func (o *Order) InsertProof() error{
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	q := `
		INSERT INTO
			orders
				(
					code
					, shipment_cost
					, state
					, customer_id
					, sub_total
					, discount_value
					, total_amount
					, payment_proof
					, coupon
				)
		VALUES
			(?, ?, ?, ?, ?, ?, ?,? ,?)
	`
	stmt, err := tx.Prepare(q)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(o.Code,ShipmentCost, o.State,o.CustomerID,o.SubTotal,o.DiscountValue,o.TotalAmount ,o.PaymentProof,o.CouponCode.String)
	if err != nil {
		return err
	}
	defer stmt.Close()

	rowCnt, err := res.RowsAffected()
	if err != nil {
		return err
	}else if rowCnt < 1 {
		return ErrNotModified
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
func (o *Order) InsertShipped() error{
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	q := `
		INSERT INTO
			orders
				(
					code
					, shipment_cost
					, state
					, customer_id
					, sub_total
					, discount_value
					, total_amount
					, payment_proof
					, shipping_code
					, coupon
				)
		VALUES
			(?, ?, ?, ?, ?, ?, ?, ?, ?,?)
	`
	stmt, err := tx.Prepare(q)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(o.Code,ShipmentCost, o.State,o.CustomerID,o.SubTotal,o.DiscountValue,o.TotalAmount ,o.PaymentProof , o.ShippingCode, o.CouponCode.String)
	if err != nil {
		return err
	}
	defer stmt.Close()

	rowCnt, err := res.RowsAffected()
	if err != nil {
		return err
	}else if rowCnt < 1 {
		return ErrNotModified
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
func (o *Order) InsertSucceeded() error{
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	q := `
		INSERT INTO
			orders
				(
					code
					, shipment_cost
					, state
					, customer_id
					, sub_total
					, discount_value
					, total_amount
					, payment_proof
					, shipping_code
					, coupon
				)
		VALUES
			(?, ?, ?, ?, ?, ?, ?, ?, ?,?)
	`
	stmt, err := tx.Prepare(q)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(o.Code,ShipmentCost, o.State,o.CustomerID,o.SubTotal,o.DiscountValue,o.TotalAmount ,o.PaymentProof , o.ShippingCode, o.CouponCode.String)
	if err != nil {
		return err
	}
	defer stmt.Close()

	rowCnt, err := res.RowsAffected()
	if err != nil {
		return err
	}else if rowCnt < 1 {
		return ErrNotModified
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
func (o *Order) InsertVoided() error{
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	q := `
		INSERT INTO
			orders
				(
					code
					, shipment_cost
					, state
					, customer_id
					, sub_total
					, discount_value
					, total_amount
					, payment_proof
					, shipping_code
				)
		VALUES
			(?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	stmt, err := tx.Prepare(q)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(o.Code,ShipmentCost, o.State,o.CustomerID,o.SubTotal,o.DiscountValue,o.TotalAmount ,o.PaymentProof , o.ShippingCode)
	if err != nil {
		return err
	}
	defer stmt.Close()

	rowCnt, err := res.RowsAffected()
	if err != nil {
		return err
	}else if rowCnt < 1 {
		return ErrNotModified
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
func (o *Order) Update(param map[string]interface{}) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	q := `
		UPDATE
			orders
		SET
			 `
		a := 0
		for key, value := range param {
			if a >= 1{
				q += `,`
			}

			fl, ok := value.(float64)
			if ok {
				q +=  key + ` = '` + strconv.FormatFloat(fl, 'g', 16, 64) + `'`
			}else{
				str, ok := value.(string)
				if ok {
					q +=  key + ` = '` + str + `'`
				}else if value.(bool){
					q +=  key + ` = true`
				}else {
					q +=  key + ` = false`
				}
			}

			a+=1
		}
	q +=`
		WHERE
			id = ?
	`

	stmt, err := tx.Prepare(q)
	if err != nil {
		return err
	}
	res, err := stmt.Exec( o.ID)
	if err != nil {
		return err
	}
	defer stmt.Close()

	rowCnt, err := res.RowsAffected()
	if err != nil {
		return err
	}else if rowCnt < 1 {
		return ErrNotModified
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func NewOrder(o Order) *Order{
	return &Order{
		ID          	:o.ID          	,
		Code        	:o.Code        	,
		CouponCode		:o.CouponCode	,
		PaymentType		:o.PaymentType	,
		PaymentProof	:o.PaymentProof	,
		CustomerID		:o.CustomerID	,
		ShippingCode	:o.ShippingCode	,
		SubTotal		:o.SubTotal		,
		ShipmentCost	:o.ShipmentCost	,
		DiscountValue	:o.DiscountValue,
		TotalAmount		:o.TotalAmount	,
		State 			:o.State 		,
		CreatedAt   	:o.CreatedAt   	,
		CreatedBy   	:o.CreatedBy   	,
		DeletedAt   	:o.DeletedAt   	,
		DeletedBy   	:o.DeletedBy   	,
		Status       	:o.Status       ,
	}
}
func (o *Order) SetID(i int) *Order {
	o.ID = int(i)
	return o
}
func (o *Order) SetCustomerID(i int) *Order {
	o.CustomerID = int(i)
	return o
}
func (o *Order) SetCode(s string) *Order {
	o.Code = sql.NullString{String:s,Valid:true}
	return o
}
func (o *Order) SetState(s string) *Order {
	o.State = s
	return o
}
func (o *Order) SetShippingCode(s string) *Order {
	o.ShippingCode = sql.NullString{String:s,Valid:true}
	return o
}
func (o *Order) SetDiscountValue(f float64) *Order {
	o.DiscountValue = f
	return o
}
func (o *Order) SetSubTotal(f float64) *Order {
	o.SubTotal = f
	return o
}
func (o *Order) SetTotalAmount(f float64) *Order {
	o.TotalAmount = f
	return o
}
func (o *Order) SetShipmentCost(f float64) *Order {
	o.ShipmentCost = f
	return o
}
func (o *Order) SetPaymentType(s string) *Order {
	o.PaymentType = sql.NullString{String:s,Valid:true}
	return o
}
func (o *Order) SetPaymentProof(b bool) *Order {
	o.PaymentProof = b
	return o
}
func (o *Order) SetCoupon(s string) *Order {
	o.CouponCode = sql.NullString{String:s,Valid:true}
	return o
}