package cartV2

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
	g.GET("", func(c echo.Context) error {

		return r.SendJSON(c, 200, "OK", "", "tes")
	})
}
