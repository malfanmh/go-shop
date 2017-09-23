package model

import (
	"github.com/go-sql-driver/mysql"
	"database/sql"
	"time"
)

type (
	Product struct {
		ID          int    			`db:"id"`
		Name        string 			`db:"name"`
		Description sql.NullString	`db:"description"`
		ImgURL      sql.NullString 	`db:"img_url"`
		Price       float64 		`db:"price"`
		Quantity    int				`db:"quantity"`
		CreatedAt   mysql.NullTime  `db:"created_at"`
		CreatedBy   string    		`db:"created_by"`
		UpdatedAt   mysql.NullTime  `db:"updated_at"`
		UpdatedBy   sql.NullString	`db:"updated_by"`
		DeletedAt   mysql.NullTime  `db:"deleted_at"`
		DeletedBy   sql.NullString	`db:"deleted_by"`
		Status      string    		`db:"status"`
	}
)

func FindProductByID(id string)(Product,error){
	return findProduct("id" , id)
}

func findProduct(field, value string) (Product, error){
	q := `
		SELECT
			id
			, name
			, description
			, img_url
			, price
			, quantity
			, created_at
			, created_by
			, updated_at
			, updated_by
			, deleted_at
			, deleted_by
			, status
		FROM
			products
		WHERE
			`+field+` = ?
		AND
			status = ?
	`

	p := new(Product)
	err := db.QueryRow(q,value, StatusCreated).Scan(&p.ID,&p.Name,&p.Description,&p.ImgURL,&p.Price,&p.Quantity,&p.CreatedAt,
													&p.CreatedBy,&p.UpdatedAt,&p.UpdatedBy,&p.DeletedAt,&p.DeletedBy,&p.Status)
	if	err != nil {
		return Product{} , err
	}
	return *p, nil
}

func ListProduct(f *Filter)([]Product, bool , error){
	if f.Page < 1 {
		f.Page = 1
	}
	if f.Count < 1 {
		f.Count = 1
	}
	q := `
		SELECT
			id
			, name
			, description
			, img_url
			, price
			, quantity
			, created_at
			, created_by
			, updated_at
			, updated_by
			, deleted_at
			, deleted_by
			, status
		FROM
			products
		WHERE
			status = ?
		LIMIT ? OFFSET ?

	`

	var products []Product
	rows , err := db.Query(q,  StatusCreated, f.Count+1, (f.Page-1)*f.Count)
	if err != nil {
		return []Product{} ,false, err

	}

	defer rows.Close()
	for rows.Next() {
		p := new(Product)
		err := rows.Scan(&p.ID,&p.Name,&p.Description,&p.ImgURL,&p.Price,&p.Quantity,&p.CreatedAt,
			&p.CreatedBy,&p.UpdatedAt,&p.UpdatedBy,&p.DeletedAt,&p.DeletedBy,&p.Status)
		if err != nil {
			return []Product{} ,false, err
		}
		products = append(products, *p)
	}

	if len(products) < 1  {
		return []Product{} ,false, ErrResourceNotFound
	}

	next := false
	if len(products) > f.Count {
		next = true
	}
	if len(products) < f.Count {
		f.Count = len(products)
	}

	return products,next, nil
}

func (p *Product) Insert() error{
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	q := `
		INSERT INTO
			products
				(
					name
					, description
					, img_url
					, price
					, quantity
					, status
				)
		VALUES
			(?, ?, ?, ?, ?, ?)
	`
	stmt, err := tx.Prepare(q)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(p.Name, p.Description,p.ImgURL,p.Price,p.Quantity,StatusCreated)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = res.LastInsertId()
	if err != nil {
		return err
	}

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

func (p *Product) Update() error {
	tx, err := db.Begin()
	if err != nil {
		return  err
	}
	q := `
		UPDATE
			products
		SET
			name = ?
			, Description = ?
			, img_url = ?
			, price = ?
			, quantity = ?
		WHERE
			id = ?
	`

	stmt, err := tx.Prepare(q)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(p.Name,p.Description,p.ImgURL,p.Price,p.Quantity, p.ID)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = res.LastInsertId()
	if err != nil {
		return err
	}

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

func (p *Product) SetID(i int) *Product {
	p.ID = i
	return p
}

func (p *Product) SetName(s string) *Product {
	p.Name = s
	return p
}
func (p *Product) SetDescription(d string) *Product {
	p.Description = sql.NullString{String:d,Valid:true}
	return p
}
func (p *Product) SetImgURL(u string) *Product {
	p.ImgURL = sql.NullString{String:u,Valid:true}
	return p
}

func (p *Product) SetPrice(f float64) *Product {
	p.Price = f
	return p
}
func (p *Product) SetQuantity(q int) *Product {
	p.Quantity = q
	return p
}
func (p *Product) SetUpdatedBy(s string) *Product {
	p.UpdatedBy =  sql.NullString{String:s,Valid:true}
	return p
}
func (p *Product) SetUpdatedAt(t time.Time) *Product {
	p.UpdatedAt =  mysql.NullTime{Time:t,Valid:true}
	return p
}

func (p *Product) SetDeletedBy(s string) *Product {
	p.UpdatedBy =  sql.NullString{String:s,Valid:true}
	return p
}
func (p *Product) SetDeletedAt(t time.Time) *Product {
	p.UpdatedAt =  mysql.NullTime{Time:t,Valid:true}
	return p
}
