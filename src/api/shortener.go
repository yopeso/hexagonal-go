package api

import (
	"github.com/andrei-dascalu/shortener/src/serializer/json"
	"github.com/andrei-dascalu/shortener/src/serializer/messagepack"
	"github.com/andrei-dascalu/shortener/src/shortener"
	"github.com/rs/zerolog/log"
)

type RedirectHandler struct {
	redirectService shortener.RedirectService
}

func NewHandler(svc shortener.RedirectService) RedirectHandler {
	log.Warn().Msg("Creating handler")
	return RedirectHandler{
		redirectService: svc,
	}
}

func serializer(contentType string) shortener.LinkSerializer {
	if contentType == "application/x-msgpack" {
		return &messagepack.Serializer{}
	}

	return &json.Serializer{}
}
