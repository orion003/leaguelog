package mock

import (
	"fmt"
	"time"

	"recleague/model"
)

func NewMockGameRepository() *MockGameRepository {
	repo := &MockGameRepository{
		lastId: 0,
		mocks:  make(map[int]interface{}),
	}

	return repo
}

func (repo *MockGameRepository) Create(game *model.Game) error {
	err := game.Validate(repo)
	if err != nil {
		return err
	}

	id := repo.lastId + 1
	game.Id = id
	repo.lastId = id

	t := time.Now()
	game.Created = t
	game.Modified = t

	repo.mocks[id] = game

	return nil
}

func (repo *MockGameRepository) FindById(id int) (*model.Game, error) {
	if game, ok := repo.mocks[id]; ok {
		if g, ok := game.(*model.Game); ok {
			return g, nil
		}

		return nil, fmt.Errorf("Game struct not returned")
	}

	return nil, fmt.Errorf("Game id not found: %d", id)
}
