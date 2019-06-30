package cartV1

import (
	"github.com/labstack/echo/v4"
	"github.com/taufikardiyan28/mataharitest1/controller/carts"
	mid "github.com/taufikardiyan28/mataharitest1/middleware"
)

type (
	Router struct {
		cart.Router
	}
)

func (r *Router) RegisterRouters(g *echo.Group) {
	g.Use(mid.Auth(false))
	g.GET("", r.List)
	g.POST("", r.Insert)
	g.PUT("/:id", r.Update)
	g.DELETE("/:id", r.Delete)
}
