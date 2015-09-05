package postgres

import (
    "time"
    
    "recleague/database/marshal"
	"recleague/model"
)

func (repo *PgUserRepository) Create(user *model.User) error {
    err := user.Validate()
	if err != nil {
		return err
	}

	t := time.Now()

	var id int
	err = repo.manager.db.QueryRow(`INSERT INTO user0(email, created, modified) 
	    VALUES($1, $2, $3) RETURNING id`,
		user.Email, t, t).Scan(&id)

	if err != nil {
		return err
	}

	user.Id = id
	user.Created = t
	user.Modified = t

	return nil
}

func (repo *PgUserRepository) FindAll() ([]model.User, error) {
	rows, err := repo.manager.db.Query(`SELECT id, email, created, modified
        FROM user0`)

	if err != nil {
		return []model.User{}, err
	}

	var users []model.User
	for rows.Next() {
		user, err := marshal.User(rows)
		if err != nil {
			return []model.User{}, err
		}

		if users == nil {
			users = make([]model.User, 1, 10)
		}

		users = append(users, *user)
	}

	err = rows.Err()
	if err != nil {
		return []model.User{}, err
	}

	return users, nil
}

func (repo *PgUserRepository) FindById(id int) (*model.User, error) {
	row := repo.manager.db.QueryRow(`SELECT id, email, created, modified
        FROM user0
        WHERE id = $1`, id)

	user, err := marshal.User(row)
	if err != nil {
		return &model.User{}, err
	}

	return user, nil
}
