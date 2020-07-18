package service

import (
	"github.com/bgildson/enext-challenge/api/repository"
	"github.com/bgildson/enext-challenge/parser"
)

// GamesService indicates how to implements a new GamesService
type GamesService interface {
	List() ([]*parser.Game, error)
	Find(id string) (*parser.Game, error)
}

type gamesService struct {
	repo repository.GamesRepository
}

// NewGamesService creates a new instance of GamesService
func NewGamesService(repo repository.GamesRepository) GamesService {
	return &gamesService{
		repo: repo,
	}
}

func (s *gamesService) List() ([]*parser.Game, error) {
	return s.repo.GetAll()
}

func (s *gamesService) Find(id string) (*parser.Game, error) {
	return s.repo.GetByID(id)
}
