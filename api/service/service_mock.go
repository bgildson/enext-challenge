package service

import "github.com/bgildson/enext-challenge/parser"

type mockGamesService struct {
	list func() ([]*parser.Game, error)
	find func(id string) (*parser.Game, error)
}

// NewMockGamesService generates a new GamesService instance for mock data
func NewMockGamesService(
	list func() ([]*parser.Game, error),
	find func(id string) (*parser.Game, error),
) GamesService {
	return &mockGamesService{
		list: list,
		find: find,
	}
}

func (s *mockGamesService) List() ([]*parser.Game, error) {
	return s.list()
}

func (s *mockGamesService) Find(id string) (*parser.Game, error) {
	return s.find(id)
}
