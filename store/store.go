package store

import (
	"database/sql"
	_ "github.com/lib/pq" //
	"http-rest-api/store/repositories"
)

type Store struct {
	config         *Config
	Db             *sql.DB
	userRepository *repositories.UserRepository
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	s.Db = db
	return nil
}

func (s *Store) Close() {
	s.Db.Close()
}

func (s *Store) User() *repositories.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &repositories.UserRepository{
		Store: s,
	}
	return s.userRepository
}
