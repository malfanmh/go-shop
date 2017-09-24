package model

import (
	"errors"
)

var (
	ErrResourceNotFound = errors.New("resource not found.")
	ErrNotModified      = errors.New("data not modified.")
)

const (
	StatusCreated string = "created"
	StatusDeleted string = "deleted"

	// "Orders.State"
	OrderStateCreated 	string = "created"
	OrderStatePending 	string = "pending"
	OrderStatePaid 		string = "paid"
	OrderStateShipped 	string = "shipped"
	OrderStateSuccess 	string = "success"
	OrderStateCanceled 	string = "canceled"
	//Orders.PaymmentType
	OrderPaymentTypeBankTf string = "bank_transfer"
	//"Coupons.Type"
	CouponTypeAmount string = "disc_amount"
	CouponTypePercendtage string = "disc_percentage"

	ALPHABET     = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	NUMERALS     = "1234567890"
	ALPHANUMERIC = ALPHABET + NUMERALS

	//hard code for "orders.shipment_cost" default value
	ShipmentCost float64 = 10000.00
	//hard code Warehouse Address for "shipments.sender_address" default address
	WarehouseAddress string = "Jl. Mayjen DI Panjaitan No. 1C, RT 001 / RW 006 (Samping Komplek Kemhan), Kelurahan Kebon Pala, Kecamatan Makasar, Jakarta Timur 13650"
)