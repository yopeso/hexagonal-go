package api

import (
	"errors"

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

func serializer(contentType string) (shortener.LinkSerializer, error) {
	if contentType == "application/x-msgpack" {
		return &messagepack.Serializer{}, nil
	}

	if contentType == "application/json" {
		return &json.Serializer{}, nil
	}

	return nil, errors.New("unknown content type")
}
