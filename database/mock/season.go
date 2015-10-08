package mock

import (
	"fmt"
	"time"

	"leaguelog.ca/model"
)

func NewMockSeasonRepository() *MockSeasonRepository {
	repo := &MockSeasonRepository{
		lastId: 0,
		mocks:  make(map[int]interface{}),
	}

	return repo
}

func (repo *MockSeasonRepository) Create(season *model.Season) error {
	err := season.Validate(repo)
	if err != nil {
		return err
	}

	id := repo.lastId + 1
	season.Id = id
	repo.lastId = id

	t := time.Now()
	season.Created = t
	season.Modified = t

	repo.mocks[id] = season

	return nil
}

func (repo *MockSeasonRepository) FindById(id int) (*model.Season, error) {
	if season, ok := repo.mocks[id]; ok {
		if s, ok := season.(*model.Season); ok {
			return s, nil
		}

		return nil, fmt.Errorf("Season struct not returned")
	}

	return nil, fmt.Errorf("Season id not found: %d", id)
}

func (repo *MockSeasonRepository) FindMostRecentByLeague(league *model.League) (*model.Season, error) {
	var recentSeason *model.Season

	for _, season := range repo.mocks {
		if s, ok := season.(*model.Season); ok {
			if s.League.Id == league.Id && (recentSeason == nil || s.EndDate.After(recentSeason.EndDate)) {
				recentSeason = s
			}
		}
	}

	return recentSeason, nil
}
