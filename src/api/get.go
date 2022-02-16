package api

import (
	"net/http"

	"github.com/andrei-dascalu/shortener/src/metrics"
	"github.com/gofiber/fiber/v2"
)

func GetRedirectHandler(handler RedirectHandler) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		code := c.Params("code")
		metrics.RecordUrlUsedRequest()

		if code == "" {
			return c.SendStatus(http.StatusBadRequest)
		}

		redirect, err := handler.redirectService.Find(code)

		if err != nil {
			return c.SendStatus(http.StatusNotFound)
		}

		metrics.RecordUrlUseSuccess()
		return c.Redirect(redirect.URL, http.StatusMovedPermanently)
	}
}
