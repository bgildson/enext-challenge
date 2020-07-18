package repository

import "github.com/bgildson/enext-challenge/parser"

type mockGamesRepository struct {
	getAll  func() ([]*parser.Game, error)
	getByID func(id string) (*parser.Game, error)
}

// NewMockGamesRepository generates a new GamesRepository instance for mock data
func NewMockGamesRepository(
	getAll func() ([]*parser.Game, error),
	getByID func(id string) (*parser.Game, error),
) GamesRepository {
	return &mockGamesRepository{
		getAll:  getAll,
		getByID: getByID,
	}
}

func (r *mockGamesRepository) GetAll() ([]*parser.Game, error) {
	return r.getAll()
}

func (r *mockGamesRepository) GetByID(id string) (*parser.Game, error) {
	return r.getByID(id)
}
