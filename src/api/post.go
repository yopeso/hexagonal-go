package api

import (
	"net/http"

	"github.com/andrei-dascalu/shortener/src/metrics"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func GetSaveHandler(handler RedirectHandler) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		contentType := c.Get("content-type")
		metrics.RecordUrlCreatedRequest()

		serializer := serializer(contentType)

		redirect, err := serializer.Decode(c.Request().Body())

		if err != nil {
			log.Error().Err(err).Msg("Failed")
			return c.SendStatus(http.StatusInternalServerError)
		}

		log.Warn().Msg("Preparing to store")
		err = handler.redirectService.Store(redirect)

		if err != nil {
			log.Error().Err(err).Msg("Failed")
			return c.SendStatus(http.StatusInternalServerError)
		}

		responseBody, err := serializer.Encode(redirect)

		if err != nil {
			log.Error().Err(err).Msg("Failed")
			return c.SendStatus(http.StatusInternalServerError)
		}

		c.Response().Header.Set("content-type", contentType)

		metrics.RecordUrlCreatedSuccess()
		return c.Status(http.StatusCreated).Send(responseBody)
	}
}
