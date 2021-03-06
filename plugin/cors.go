package plugin

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type (
	CORS struct {
		Base                  `json:",squash"`
		middleware.CORSConfig `json:",squash"`
	}
)

func (c *CORS) Initialize() (err error) {
	c.Middleware = middleware.CORSWithConfig(c.CORSConfig)
	return
}

func (*CORS) Priority() int {
	return 1
}

func (c *CORS) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return c.Middleware(next)
}

func (*CORS) Terminate() {
}
