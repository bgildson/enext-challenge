package repository

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/bgildson/enext-challenge/api/database"
	"github.com/bgildson/enext-challenge/parser"
)

// GamesRepository indicates how to implements a new GamesRepository
type GamesRepository interface {
	GetAll() ([]*parser.Game, error)
	GetByID(id string) (*parser.Game, error)
}

// Reusable errors
var (
	ErrMalformedDatabaseResult = errors.New("could not handle database result")
)

type jsonGamesRepository struct {
	db database.Database
}

// NewJSONGamesRepository creates a new GamesRepository using a Database from parsed logs
func NewJSONGamesRepository(db database.Database) GamesRepository {
	return &jsonGamesRepository{db}
}

func (r *jsonGamesRepository) GetAll() ([]*parser.Game, error) {
	gs, err := r.db.Get()
	if err != nil {
		return nil, fmt.Errorf("could not load games from database: %v", err)
	}

	b, err := json.Marshal(gs)
	if err != nil {
		return nil, ErrMalformedDatabaseResult
	}

	var games []*parser.Game
	if err := json.Unmarshal(b, &games); err != nil {
		return nil, ErrMalformedDatabaseResult
	}

	return games, nil
}

func (r *jsonGamesRepository) GetByID(id string) (*parser.Game, error) {
	g, err := r.db.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("could not get game %v: %v", id, err)
	}

	b, err := json.Marshal(g)
	if err != nil {
		return nil, ErrMalformedDatabaseResult
	}

	var game *parser.Game
	if err := json.Unmarshal(b, &game); err != nil {
		return nil, ErrMalformedDatabaseResult
	}

	return game, nil
}
