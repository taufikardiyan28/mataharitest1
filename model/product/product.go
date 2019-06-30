package ProductModel

import (
	db "github.com/taufikardiyan28/mataharitest1/model"
)

type (
	Product     struct{ db.DB }
	ProductData struct {
		Id          int64  `db:"id" json:"id"`
		ProductName string `db:"productName" json:"product_name"`
		Stock       int    `db:"stock" json:"stock"`
	}
)

func (p *Product) List() ([]ProductData, error) {
	strSQL := `SELECT id, productName, stock FROM products`
	data := []ProductData{}
	pool := p.GetPool()
	err := pool.Select(&data, strSQL)
	return data, err
}

func (p *Product) Get(productId int64) (ProductData, error) {
	strSQL := `SELECT id, productName, stock FROM products WHERE id=?`
	data := ProductData{}
	pool := p.GetPool()
	err := pool.Get(&data, strSQL)
	return data, err
}
