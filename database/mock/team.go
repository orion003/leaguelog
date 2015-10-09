package mock

import (
	"fmt"
	"time"

	"leaguelog/model"
)

func NewMockTeamRepository() *MockTeamRepository {
	repo := &MockTeamRepository{
		lastId: 0,
		mocks:  make(map[int]interface{}),
	}

	return repo
}

func (repo *MockTeamRepository) Create(team *model.Team) error {
	err := team.Validate(repo)
	if err != nil {
		return err
	}

	id := repo.lastId + 1
	team.Id = id
	repo.lastId = id

	t := time.Now()
	team.Created = t
	team.Modified = t

	repo.mocks[id] = team

	return nil
}

func (repo *MockTeamRepository) FindById(id int) (*model.Team, error) {
	if team, ok := repo.mocks[id]; ok {
		if t, ok := team.(*model.Team); ok {
			return t, nil
		}

		return nil, fmt.Errorf("Team struct not returned")
	}

	return nil, fmt.Errorf("Team id not found: %d", id)
}
