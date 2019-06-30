package productV1

import (
	"github.com/labstack/echo/v4"
	"github.com/taufikardiyan28/mataharitest1/controller/products"
	mid "github.com/taufikardiyan28/mataharitest1/middleware"
)

type (
	Router struct {
		product.Router
	}
)

func (r *Router) RegisterRouters(g *echo.Group) {
	g.Use(mid.Auth(false))
	g.GET("", r.List)
}
