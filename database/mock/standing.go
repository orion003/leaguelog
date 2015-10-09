package mock

import (
	"time"

	"leaguelog/model"
)

func NewMockStandingRepository() *MockStandingRepository {
	repo := &MockStandingRepository{
		lastId: 0,
		mocks:  make(map[int]interface{}),
	}

	return repo
}

func (repo *MockStandingRepository) Create(standing *model.Standing) error {
	err := standing.Validate(repo)
	if err != nil {
		return err
	}

	id := repo.lastId + 1
	standing.Id = id
	repo.lastId = id

	t := time.Now()
	standing.Created = t
	standing.Modified = t

	repo.mocks[id] = standing

	return nil
}

func (repo *MockStandingRepository) FindAllBySeason(season *model.Season) ([]model.Standing, error) {
	standings := make([]model.Standing, 0, len(repo.mocks))

	for _, standing := range repo.mocks {
		if s, ok := standing.(*model.Standing); ok {
			if s.Season.Id == season.Id {
				standings = append(standings, *s)
			}
		}
	}

	return standings, nil
}
