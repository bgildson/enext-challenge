package serializer

import (
	"encoding/json"
	"fmt"

	"github.com/bgildson/enext-challenge/parser"
)

// GameSerializer indicates how to implement GameSerializer
type GameSerializer interface {
	Serialize(game *parser.Game) ([]byte, error)
}

type jsonGameSerializer struct{}

// NewJSONGameSerializer creates a new instance of GameSerializer
func NewJSONGameSerializer() GameSerializer {
	return &jsonGameSerializer{}
}

func (s *jsonGameSerializer) Serialize(game *parser.Game) ([]byte, error) {
	b, err := json.Marshal(game)
	if err != nil {
		return nil, fmt.Errorf("could not serialize game: %v", err)
	}
	return b, nil
}

// GamesSerializer indicates how to implement GamesSerializer
type GamesSerializer interface {
	Serialize(games []*parser.Game) ([]byte, error)
}

type jsonGamesSerializer struct{}

// NewJSONGamesSerializer creates a new instance of GamesSerializer
func NewJSONGamesSerializer() GamesSerializer {
	return &jsonGamesSerializer{}
}

func (s *jsonGamesSerializer) Serialize(games []*parser.Game) ([]byte, error) {
	b, err := json.Marshal(games)
	if err != nil {
		return nil, fmt.Errorf("could not serialize games: %v", err)
	}
	return b, nil
}
