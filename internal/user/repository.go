package user

import "go-advance/pkg/db"

type UserRepository struct {
	*db.Db
}

func NewUserRepository(db *db.Db) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Create(user *User) error {
	result := r.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *UserRepository) FindByEmail(email string) (*User, error) {
	var user *User
	result := r.DB.First(user, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
