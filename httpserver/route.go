package httpserver

import (
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func ping(c *fiber.Ctx) (err error) {

	return c.Send([]byte{'o', 'k'})
}
