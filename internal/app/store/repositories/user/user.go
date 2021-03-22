package user

import "github.com/dwnGnL/taxi/internal/app/store"

type UserRepository struct {
	Store *store.Store
}

// Create
func (r *UserRepository) Create(u *User) (*User, error) {

	if err := u.Validate(); err != nil {
		return nil, err
	}

	if err := u.BeforeCreate(); err != nil {
		return nil, err
	}

	if err := r.Store.DB.Create(&u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

// FindByEmail
func (r *UserRepository) FindByEmail(email string) (*User, error) {
	u := User{}
	if err := r.Store.DB.Where("email = ?", email).Find(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}
