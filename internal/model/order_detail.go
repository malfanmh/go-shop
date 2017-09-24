package model

import (
	"github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
)

type OrderDetails struct {
	ID			int 			`db:"id"`
	OrderCode 	string 			`db:"order_code"`
	ProductID	int 			`db:"product_id"`
	Price 		float64 		`db:"price"`
	Quantity	int 			`db:"quantity"`
	Amount 		float64			`db:"amount"`
	CreatedAt   mysql.NullTime 	`db:"created_at"`
	CreatedBy   string        	`db:"created_by"`
	DeletedAt   mysql.NullTime  `db:"created_at,omitempty"`
	DeletedBy   sql.NullString  `db:"created_by,omitempty"`
	Status      string        	`db:"status"`
}

func FindOrderDetailByOrderID(id string) ([]OrderDetails , error){
	return findOrderDetail("order_code" , id)
}
func findOrderDetail(field, value string) ([]OrderDetails , error){
	q := `
		SELECT
			id
			, order_code
			, product_id
			, price
			, quantity
			, amount
			, created_at
			, created_by
			, deleted_at
			, deleted_by
			, status
		FROM
			order_details
		WHERE
			`+field+` = ?
		AND
			status = ?
	`

	var orderdetail []OrderDetails
	rows , err := db.Query(q,value, StatusCreated)
	if err != nil {
		return []OrderDetails{} , err
	}

	defer rows.Close()
	for rows.Next() {
		od := new(OrderDetails)
		err := rows.Scan(&od.ID,&od.OrderCode,&od.ProductID,&od.Price,&od.Quantity,&od.Amount,
			&od.CreatedAt,&od.CreatedBy,&od.DeletedAt,&od.DeletedBy,&od.Status)
		if err != nil {
			return []OrderDetails{} , err
		}
		orderdetail = append(orderdetail, *od)
	}
	if len(orderdetail) < 1  {
		return []OrderDetails{} , ErrResourceNotFound
	}
	return orderdetail, nil
}


func (od *OrderDetails) Insert() error{
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("error", err)
		return err
	}
	q := `
		INSERT INTO
			order_details
				(
					order_code
					, product_id
					, price
					, quantity
					, amount
				)
		VALUES
			(?, ?, ?, ?, ?)
	`
	stmt, err := tx.Prepare(q)
	if err != nil {
		fmt.Println("error1", err)
		return err
	}
	res, err := stmt.Exec(od.OrderCode , od.ProductID, od.Price,od.Quantity,od.Amount)
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

func (od *OrderDetails) SetOrderCode(i string) *OrderDetails {
	od.OrderCode = i
	return od
}
func (od *OrderDetails) SetProductID(i int) *OrderDetails {
	od.ProductID = i
	return od
}
func (od *OrderDetails) SetQuantity(i int) *OrderDetails {
	od.Quantity = i
	return od
}
func (od *OrderDetails) SetPrice(f float64) *OrderDetails {
	od.Price = f
	return od
}
func (od *OrderDetails) SetAmount(f float64) *OrderDetails {
	od.Amount = f
	return od
}

