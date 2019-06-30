package controller

import (

	"github.com/labstack/echo/v4"	
	productV1 "github.com/taufikardiyan28/mataharitest1/controller/products/v1"
)

func RegisterRouters(r *echo.Group) {
	registerV1(r.Group("/v1"))
}

func registerV1(r *echo.Group) {
	_productV1 := productV1.Router{}
	_productV1.RegisterRouters(r.Group("/products"))
}
