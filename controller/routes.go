package controller

import (
	"github.com/labstack/echo/v4"
	cartV1 "github.com/taufikardiyan28/mataharitest1/controller/carts/v1"
	productV1 "github.com/taufikardiyan28/mataharitest1/controller/products/v1"
)

func RegisterRouters(r *echo.Group) {
	registerV1(r.Group("/v1"))
}

func registerV1(r *echo.Group) {
	//api/v1/products
	_productV1 := productV1.Router{}
	_productV1.RegisterRouters(r.Group("/products"))

	//api/v1/carts
	_cartV1 := cartV1.Router{}
	_cartV1.RegisterRouters(r.Group("/carts"))
}
