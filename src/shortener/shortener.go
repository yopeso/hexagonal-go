package shortener

import (
	"errors"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/teris-io/shortid"
)

type redirectService struct {
	repo LinkRepository
}

func NewRedirectService(repo LinkRepository) RedirectService {
	return &redirectService{
		repo: repo,
	}
}

func (r *redirectService) Find(code string) (*LinkRedirect, error) {
	return r.repo.Find(code)
}

func (r *redirectService) Store(redirect *LinkRedirect) error {
	log.Warn().Msg("Service preparing to store")

	if redirect == nil {
		return errors.New("null link cannot be saved")
	}

	if r.repo == nil {
		return errors.New("repository was not created")
	}

	redirect.Code, _ = shortid.Generate()
	redirect.CreatedAt = time.Now().Unix()

	return r.repo.Store(redirect)
}
