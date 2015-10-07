package postgres

import (
	"database/sql"

	_ "recleague/Godeps/_workspace/src/github.com/lib/pq"
)

type PgManager struct {
	db *sql.DB
}

type PgRepository struct {
	manager *PgManager
}

type PgLeagueRepository PgRepository
type PgSeasonRepository PgRepository
type PgGameRepository PgRepository
type PgTeamRepository PgRepository
type PgStandingRepository PgRepository
type PgUserRepository PgRepository

func NewPgManager(url string) (*PgManager, error) {
	manager := &PgManager{}
	err := manager.open(url)
	if err != nil {
		return nil, err
	}

	return manager, nil
}

func (manager *PgManager) open(url string) error {
	var err error

	if manager.db == nil {
		manager.db, err = sql.Open("postgres", url)

		return err
	}

	return err
}

func (manager *PgManager) close() error {
	err := manager.db.Close()
	manager.db = nil
	if err != nil {
		return err
	}

	return nil
}
