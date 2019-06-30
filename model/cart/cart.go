package CartModel

import (
	"database/sql"

	db "github.com/taufikardiyan28/mataharitest1/model"
)

type (
	Cart     struct{ db.DB }
	CartData struct {
		Id           int64         `db:"id" json:"id"`
		ProductId    int64         `db:"productId" json:"product_id" validate:"required"`
		ProductName  string        `db:"productName" json:"product_name"`
		Qty          int           `db:"qty" json:"qty" validate:"required"`
		AID          string        `db:"AID" json:"-"`
		UserId       sql.NullInt64 `db:"userId" json:"-"`
		IsOrdered    int           `db:"isOrdered"`
		CurrentStock int           `db:"currentStock" json:"current_stock"`
	}
)

func (c *Cart) List(param CartData) ([]CartData, error) {
	if !param.UserId.Valid {
		return c.listByAID(param.AID)
	}
	return c.listByUser(param.UserId.Int64)
}

func (c *Cart) listByAID(AID string) ([]CartData, error) {
	strSQL := `SELECT a.id, b.productName, a.qty, b.stock as currentStock FROM carts 
				a INNER JOIN products b ON a.productId=b.id WHERE AID=? AND isOrdered=0`
	data := []CartData{}
	pool := c.GetPool()
	err := pool.Select(&data, strSQL, AID)
	return data, err
}

func (c *Cart) listByUser(userId int64) ([]CartData, error) {
	strSQL := `SELECT a.id, b.productName, a.qty, b.stock as currentStock FROM carts 
				a INNER JOIN products b ON a.productId=b.id WHERE userId=? AND isOrdered=0`
	data := []CartData{}
	pool := c.GetPool()
	err := pool.Select(&data, strSQL, userId)
	return data, err
}

func (c *Cart) Insert(param *CartData) (int64, error) {
	args := []interface{}{}
	strSQL := `SELECT id, qty FROM carts WHERE `
	if param.UserId.Valid {
		strSQL += `userId=?`
		args = append(args, param.UserId)
	} else {
		strSQL += `AID=?`
		args = append(args, param.AID)
	}
	strSQL += ` AND productId=?`
	args = append(args, param.ProductId)

	pool := c.GetPool()

	cart := CartData{}
	if err := pool.Get(&cart, strSQL, args...); err != nil && err != sql.ErrNoRows {
		return 0, err
	}

	if !((CartData{}) == cart) {
		param.Id = cart.Id
		param.Qty += cart.Qty
		return param.Id, c.Update(param)
	}

	strSQL = `INSERT INTO carts (productId, qty, AID, userId) VALUES (?, ?, ?, ?)`
	args = []interface{}{param.ProductId, param.Qty, param.AID, param.UserId}
	res, err := pool.Exec(strSQL, args...)
	if err != nil {
		return 0, err
	}
	insertId, _ := res.LastInsertId()

	return insertId, err
}

func (c *Cart) Update(param *CartData) error {
	pool := c.GetPool()
	strSQL := `UPDATE carts SET qty=? WHERE id=?`
	_, err := pool.Exec(strSQL, param.Qty, param.Id)
	return err
}

func (c *Cart) Delete(id int64) error {
	pool := c.GetPool()
	strSQL := `DELETE FROM carts WHERE id=?`
	_, err := pool.Exec(strSQL, id)
	return err
}
