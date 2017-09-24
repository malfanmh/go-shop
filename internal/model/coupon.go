package model

import (
	"time"
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

type (
	Coupon struct {
		ID          int				`db:"id"`
		Code        string 			`db:"code"`
		Description sql.NullString 	`db:"description"`
		Type        string 			`db:"type"`
		Value       float64 		`db:"value"`
		Quantity	int				`db:"quantity"`
		Used		int				`db:"used"`
		ValidAt     string 			`db:"valid_at"`
		ValidUntil  string  		`db:"valid_until"`
		TnC         sql.NullString 	`db:"tnc"`
		CreatedAt   mysql.NullTime 	`db:"created_at"`
		CreatedBy   string        	`db:"created_by"`
		UpdatedAt   mysql.NullTime  `db:"created_at,omitempty"`
		UpdatedBy   sql.NullString  `db:"created_by,omitempty"`
		DeletedAt   mysql.NullTime  `db:"created_at,omitempty"`
		DeletedBy   sql.NullString  `db:"created_by,omitempty"`
		Status      string        	`db:"status"`
	}
)

func FindCouponByCode(code string) (Coupon,error){
	return findCoupon("code" ,code)
}

func findCoupon(field, value string) (Coupon , error){
	q := `
		SELECT
			id
			, code
			, description
			, type
			, value
			, quantity
			, used
			, valid_at
			, valid_until
			, tnc
			, created_at
			, created_by
			, updated_at
			, updated_by
			, deleted_at
			, deleted_by
			, status
		FROM
			coupons
		WHERE
			`+field+` = ?
		AND
			status = ?
	`

	c := new(Coupon)
	err := db.QueryRow(q,value, StatusCreated).Scan(&c.ID,&c.Code,&c.Description,&c.Type,&c.Value,&c.Quantity,&c.Used,&c.ValidAt,&c.ValidUntil,&c.TnC,
		 				&c.CreatedAt,&c.CreatedBy,&c.UpdatedAt,&c.UpdatedBy,&c.DeletedAt,&c.DeletedBy,&c.Status)
	if	c.Code == "" {
		return Coupon{} , ErrResourceNotFound
	}else if err != nil {
		return Coupon{} , err
	}
	return *c, nil
}

func ListCoupon(f *Filter)([]Coupon, bool , error){
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
			, description
			, type
			, value
			, quantity
			, used
			, valid_at
			, valid_until
			, tnc
			, created_at
			, created_by
			, updated_at
			, updated_by
			, deleted_at
			, deleted_by
			, status
		FROM
			coupons
		WHERE
			status = ?
		LIMIT ? OFFSET ?

	`

	var coupons []Coupon
	rows , err := db.Query(q,  StatusCreated, f.Count+1, (f.Page-1)*f.Count)
	if err != nil {
		return []Coupon{} ,false, err

	}

	defer rows.Close()
	for rows.Next() {
		c := new(Coupon)
		err := rows.Scan(&c.ID,&c.Code,&c.Description,&c.Type,&c.Value,&c.Quantity,&c.Used,&c.ValidAt,&c.ValidUntil,&c.TnC,
			&c.CreatedAt,&c.CreatedBy,&c.UpdatedAt,&c.UpdatedBy,&c.DeletedAt,&c.DeletedBy,&c.Status)
		if err != nil {
			return []Coupon{} ,false, err
		}
		coupons = append(coupons, *c)
	}
	if len(coupons) < 1  {
		return []Coupon{} ,false, ErrResourceNotFound
	}

	next := false
	if len(coupons) > f.Count {
		next = true
	}
	if len(coupons) < f.Count {
		f.Count = len(coupons)
	}

	return coupons,next, nil
}

func (c *Coupon) Insert() error{
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	q := `
		INSERT INTO
			coupons
				(
					code
					, description
					, type
					, value
					, quantity
					, valid_at
					, valid_until
					, tnc
				)
		VALUES
			(?, ?, ?, ?, ?, ? , ?, ?)
	`
	stmt, err := tx.Prepare(q)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(c.Code, c.Description,c.Type,c.Value,c.Quantity,c.ValidAt,c.ValidUntil,c.TnC)
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

func (c *Coupon) Update() error{
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	q := `
		UPDATE
			coupons
		SET
			description = ?
			, type = ?
			, value = ?
			, quantity = ?
			, valid_at = ?
			, valid_until = ?
			, tnc = ?
		WHERE
			id = ?
	`
	stmt, err := tx.Prepare(q)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(c.Description,c.Type,c.Value,c.Quantity,c.ValidAt,c.ValidUntil,c.TnC,c.ID)
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


func (c *Coupon) SetID(i int) *Coupon {
	c.ID = i
	return c
}

func (c *Coupon) SetCode(s string) *Coupon {
	c.Code = s
	return c
}
func (c *Coupon) SetDescription(d string) *Coupon {
	c.Description = sql.NullString{String:d,Valid:true}
	return c
}
func (c *Coupon) SetType(s string) *Coupon {
	c.Type = s
	return c
}
func (c *Coupon) SetValue(f float64) *Coupon {
	c.Value = f
	return c
}
func (c *Coupon) SetQuantity(i int) *Coupon {
	c.Quantity = i
	return c
}
func (c *Coupon) SetValidAt(s string) *Coupon {
	c.ValidAt = s
	return c
}
func (c *Coupon) SetValidUntil(s string) *Coupon {
	c.ValidUntil = s
	return c
}

func (c *Coupon) SetTnc(s string) *Coupon {
	c.TnC = sql.NullString{String:s,Valid:true}
	return c
}

func (c *Coupon) SetUpdatedBy(s string) *Coupon {
	c.UpdatedBy =  sql.NullString{String:s,Valid:true}
	return c
}
func (c *Coupon) SetUpdatedAt(t time.Time) *Coupon {
	c.UpdatedAt =  mysql.NullTime{Time:t,Valid:true}
	return c
}

func (c *Coupon) SetDeletedBy(s string) *Coupon {
	c.UpdatedBy =  sql.NullString{String:s,Valid:true}
	return c
}
func (c *Coupon) SetDeletedAt(t time.Time) *Coupon {
	c.UpdatedAt =  mysql.NullTime{Time:t,Valid:true}
	return c
}
