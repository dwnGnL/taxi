package user

import (
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func New(p *gorm.DB) *UserRepository {
	return &UserRepository{
		db: p,
	}
}

// Create
func (r *UserRepository) Create(u *User) (*User, error) {

	if err := u.Validate(); err != nil {
		return nil, err
	}

	if err := u.BeforeCreate(); err != nil {
		return nil, err
	}

	if err := r.db.Create(&u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

// FindByEmail
func (r *UserRepository) FindByEmail(email string) (*User, error) {
	u := User{}
	if err := r.db.Where("email = ?", email).Find(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}
