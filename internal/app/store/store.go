package store

import (
	"github.com/dwnGnL/taxi/internal/app/store/repositories/user"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Store struct {
	config         *Config
	db             *gorm.DB
	userRepository *user.UserRepository
	logger         *logrus.Logger
}

func New(config *Config, logger *logrus.Logger) *Store {
	return &Store{
		config: config,
		logger: logger,
	}
}

// Open ...
func (s *Store) Open() error {
	dsn := s.config.DatabaseURL
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}
	if err = db.AutoMigrate(&user.User{}); err != nil {
		s.logger.Warnln("migrate error: ", err)
	}
	s.db = db
	return nil
}

func (s *Store) User() *user.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = user.New(s.db)
	return s.userRepository
}

// Close ...
func (s *Store) Close() {
	sqlDB, _ := s.db.DB()
	_ = sqlDB.Close()
}
