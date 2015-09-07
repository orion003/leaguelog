package mock

import (
	"fmt"
	"time"

	"recleague/model"
)

func NewMockUserRepository() *MockUserRepository {
	repo := &MockUserRepository{
		lastId: 0,
		mocks:  make(map[int]interface{}),
	}

	return repo
}

func (repo *MockUserRepository) Create(user *model.User) error {
	err := user.Validate(repo)
	if err != nil {
		return err
	}

	id := repo.lastId + 1
	user.Id = id
	repo.lastId = id

	t := time.Now()
	user.Created = t
	user.Modified = t

	repo.mocks[id] = user

	return nil
}

func (repo *MockUserRepository) FindAll() ([]model.User, error) {
	users := make([]model.User, 0, len(repo.mocks))

	for _, user := range repo.mocks {
		if u, ok := user.(*model.User); ok {
			users = append(users, *u)
		}
	}

	return users, nil
}

func (repo *MockUserRepository) FindById(id int) (*model.User, error) {
	if user, ok := repo.mocks[id]; ok {
		if u, ok := user.(*model.User); ok {
			return u, nil
		}

		return nil, fmt.Errorf("User struct not returned")
	}

	return nil, fmt.Errorf("User id not found: %d", id)
}

func (repo *MockUserRepository) FindByEmail(email string) (*model.User, error) {
	for _, user := range repo.mocks {
		if u, ok := user.(*model.User); ok {
			if u.Email == email {
				return u, nil
			}
		}
	}

	return nil, fmt.Errorf("User with email not found: %s", email)
}
