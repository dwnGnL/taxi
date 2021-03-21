package store

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Store struct {
	config *Config
	db     *gorm.DB
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

	s.db = db
	return nil
}

// Close ...
func (s *Store) Close() {
	sqlDB, _ := s.db.DB()
	sqlDB.Close()
}
