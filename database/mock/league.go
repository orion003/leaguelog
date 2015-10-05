package mock

import (
	"fmt"
	"time"

	"recleague/model"
)

func NewMockLeagueRepository() *MockLeagueRepository {
	repo := &MockLeagueRepository{
		lastId: 0,
		mocks:  make(map[int]interface{}),
	}

	return repo
}

func (repo *MockLeagueRepository) Create(league *model.League) error {
	err := league.Validate(repo)
	if err != nil {
		return err
	}

	id := repo.lastId + 1
	league.Id = id
	repo.lastId = id

	t := time.Now()
	league.Created = t
	league.Modified = t

	repo.mocks[id] = league

	return nil
}

func (repo *MockLeagueRepository) FindAll() ([]model.League, error) {
	leagues := make([]model.League, 0, len(repo.mocks))

	for _, league := range repo.mocks {
		if l, ok := league.(*model.League); ok {
			leagues = append(leagues, *l)
		}
	}

	return leagues, nil
}

func (repo *MockLeagueRepository) FindById(id int) (*model.League, error) {
	if league, ok := repo.mocks[id]; ok {
		if l, ok := league.(*model.League); ok {
			return l, nil
		}

		return nil, fmt.Errorf("League struct not returned")
	}

	return nil, fmt.Errorf("League id not found: %d", id)
}
