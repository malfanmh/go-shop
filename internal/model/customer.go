package model

import (
	"github.com/go-sql-driver/mysql"
	"database/sql"
)

type (
	Customers struct {
		ID int `db:"id"`
		Name string `db:"name"`
		Phone string `db:"phone"`
		Email string `db:"email"`
		Address string `db:"address"`
		CreatedAt   mysql.NullTime 	`db:"created_at"`
		CreatedBy   string        	`db:"created_by"`
		DeletedAt   mysql.NullTime  `db:"created_at,omitempty"`
		DeletedBy   sql.NullString  `db:"created_by,omitempty"`
		Status      string        	`db:"status"`
	}
)

func FindCustomerByID (id int)(Customers, error){
	q := `
		SELECT
			id
			, name
			, phone
			, email
			, address
			, created_at
			, created_by
			, deleted_at
			, deleted_by
			, status
		FROM
			customers
		WHERE
			id = ?
		AND
			status = ?
	`

	c := new(Customers)
	err := db.QueryRow(q,id, StatusCreated).Scan(&c.ID,&c.Name,&c.Phone,&c.Email,&c.Address,
		  										 &c.CreatedAt,&c.CreatedBy,&c.DeletedAt,&c.DeletedBy,&c.Status)
	if	c.Name == "" {
		return Customers{} , ErrResourceNotFound
	}else if err != nil {
		return Customers{} , err
	}
	return *c, nil
}

func (c *Customers) Insert() (int64 ,error){
	tx, err := db.Begin()
	if err != nil {
		return 0,err
	}
	q := `
		INSERT INTO
			customers
				(
					name
					, phone
					, email
					, address
				)
		VALUES
			(?, ?, ?, ?)
	`
	stmt, err := tx.Prepare(q)
	if err != nil {
		return 0,err
	}

	res, err := stmt.Exec(c.Name, c.Phone,c.Email,c.Address)
	if err != nil {
		return 0,err
	}
	defer stmt.Close()

	id, err := res.LastInsertId()
	if err != nil {
		return 0,err
	}

	rowCnt, err := res.RowsAffected()
	if err != nil {
		return 0,err
	}else if rowCnt < 1 {
		return 0,ErrNotModified
	}

	err = tx.Commit()
	if err != nil {
		return 0,err
	}

	return id ,nil
}

func (c *Customers) SetName(s string) *Customers {
	c.Name = s
	return c
}
func (c *Customers) SetPhone(s string) *Customers {
	c.Phone = s
	return c
}
func (c *Customers) SetEmail(s string) *Customers {
	c.Email = s
	return c
}
func (c *Customers) SetAddress(s string) *Customers {
	c.Address = s
	return c
}