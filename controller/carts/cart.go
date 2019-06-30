package cart

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	mid "github.com/taufikardiyan28/mataharitest1/middleware"
	CartModel "github.com/taufikardiyan28/mataharitest1/model/cart"
)

type (
	Router struct {
		mid.Request
	}
)

func (r *Router) RegisterRouters(g *echo.Group) {}

func (r *Router) List(c echo.Context) error {
	cc := c.(*mid.AuthContext)

	mod := CartModel.Cart{}
	param := CartModel.CartData{
		AID:    r.GetCookie(c, "AID"),
		UserId: sql.NullInt64{Int64: cc.UserId, Valid: cc.UserId != 0},
	}
	data, err := mod.List(param)
	if err != nil {
		return r.Send(c, http.StatusInternalServerError, "Cannot get data", err.Error(), nil)
	}

	return r.Send(c, http.StatusOK, "OK", "", data)
}

func (r *Router) Insert(c echo.Context) error {
	cc := c.(*mid.AuthContext)

	body := new(CartModel.CartData)
	if err := r.BindAndValidate(c, body); err != nil {
		return r.Send(c, http.StatusBadRequest, "Invalid Data", err.Error(), "")
	}

	mod := CartModel.Cart{}
	body.AID = r.GetCookie(c, "AID")
	body.UserId = sql.NullInt64{Int64: cc.UserId, Valid: cc.UserId != 0}

	insertId, err := mod.Insert(body)
	if err != nil {
		return r.Send(c, http.StatusInternalServerError, "Cannot insert data", err.Error(), nil)
	}

	return r.Send(c, http.StatusOK, "OK", "", map[string]interface{}{"insertId": insertId})
}

func (r *Router) Update(c echo.Context) error {
	body := new(CartModel.CartData)
	if err := r.BindAndValidate(c, body); err != nil {
		return r.Send(c, http.StatusBadRequest, "Invalid Data", err.Error(), "")
	}

	id, _ := strconv.Atoi(c.Param("id"))
	mod := CartModel.Cart{}
	body.Id = int64(id)
	err := mod.Update(body)
	if err != nil {
		return r.Send(c, http.StatusInternalServerError, "Cannot insert data", err.Error(), nil)
	}

	return r.Send(c, http.StatusOK, "OK", "", nil)
}

func (r *Router) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	mod := CartModel.Cart{}
	err := mod.Delete(int64(id))
	if err != nil {
		return r.Send(c, http.StatusInternalServerError, "Cannot delete data", err.Error(), nil)
	}

	return r.Send(c, http.StatusOK, "OK", "", nil)
}
