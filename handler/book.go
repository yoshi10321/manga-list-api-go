package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

type HelloHandler struct {
}

func NewHelloHandler() HelloHandler {
	return HelloHandler{}
}

func (handler HelloHandler) Hello(c echo.Context) (err error) {
	return c.String(http.StatusOK, "Hello world handler")
}
