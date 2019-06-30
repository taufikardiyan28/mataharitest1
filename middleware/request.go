package middleware

import (
	"github.com/labstack/echo/v4"
)

type (
	Request  struct{}
	Response struct {
		Status int
		Msg    string
		ErrMsg string
		Data   interface{}
	}
)

func (r *Request) BindAndValidate(c echo.Context, i interface{}) error {
	if err := c.Bind(i); err != nil {
		return err
	}
	err := c.Validate(i)
	return err
}

func (r *Request) Send(c echo.Context, status int, msg string, errmsg string, data interface{}) error {
	if data == nil {
		data = []string{}
	}

	outType := c.QueryParam("output")
	if outType == "xml" {
		return r.SendXML(c, status, msg, errmsg, data)
	}

	return r.SendJSON(c, status, msg, errmsg, data)
}

func (r *Request) SendJSON(c echo.Context, status int, msg string, errmsg string, data interface{}) error {
	return c.JSON(status, map[string]interface{}{"status": status, "msg": msg, "err_msg": errmsg, "data": data})
}

func (r *Request) SendXML(c echo.Context, status int, msg string, errmsg string, data interface{}) error {
	res := &Response{
		Status: status,
		Msg:    msg,
		ErrMsg: errmsg,
		Data:   data,
	}
	return c.XML(status, res)
}

func (r *Request) GetCookie(c echo.Context, key string) string {
	cookie, err := c.Cookie(key)
	if err != nil {
		return ""
	}
	return cookie.Value
}
