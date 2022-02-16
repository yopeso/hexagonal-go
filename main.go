package main

import (
	"errors"
	"os"

	"github.com/rs/zerolog/log"

	"github.com/andrei-dascalu/shortener/src/api"
	"github.com/andrei-dascalu/shortener/src/repository/mongodb"
	"github.com/andrei-dascalu/shortener/src/repository/redis"
	"github.com/andrei-dascalu/shortener/src/shortener"
	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
)

func main() {
	repo, err := chooseRepo()

	if err != nil {
		log.Panic().Err(err).Msg("error")
	}

	service := shortener.NewRedirectService(repo)

	if service == nil {
		log.Panic().Err(err).Msg("error")
	}

	app := fiber.New()

	prometheus := fiberprometheus.New("my-service-name")
	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Middleware)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "ok",
		})
	})

	app.Post("/url", api.GetSaveHandler(api.NewHandler(service)))
	app.Get("/:code", api.GetRedirectHandler(api.NewHandler(service)))

	app.Listen(":8080")
}

func chooseRepo() (shortener.LinkRepository, error) {
	repoType := os.Getenv("DB_TYPE")
	log.Warn().Str("repoType", repoType).Msg("Selected repo type")
	switch repoType {
	case "redis":
		redisURL := os.Getenv("REDIS_URL")
		repo, _ := redis.NewRedisRepository(redisURL)

		return repo, nil
	case "mongo":
		mongoURL := os.Getenv("MONGO_URL")
		mongoDb := os.Getenv("MONGO_DB")

		repo, err := mongodb.NewMongoRepository(mongoURL, mongoDb)

		if err != nil {
			return nil, err
		}

		return repo, nil
	}

	return nil, errors.New("missing configuration")
}
