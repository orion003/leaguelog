package mock

import ()

type MockRepository struct {
	lastId int
	mocks  map[int]interface{}
}

type MockLeagueRepository MockRepository
type MockSeasonRepository MockRepository
type MockGameRepository MockRepository
type MockTeamRepository MockRepository
type MockStandingRepository MockRepository
type MockUserRepository MockRepository
