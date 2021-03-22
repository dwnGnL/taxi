package store

import (
	"github.com/dwnGnL/taxi/internal/app/store/repositories/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Store struct {
	config         *Config
	DB             *gorm.DB
	userRepository *user.UserRepository
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// Open ...
func (s *Store) Open() error {
	dsn := s.config.DatabaseURL
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	s.DB = db
	return nil
}

func (s *Store) User() *user.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &user.UserRepository{
		Store: s,
	}
	return s.userRepository
}

// Close ...
func (s *Store) Close() {
	sqlDB, _ := s.DB.DB()
	sqlDB.Close()
}
