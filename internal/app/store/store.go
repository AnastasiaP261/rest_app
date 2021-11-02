package store

import (
	"database/sql"
	_ "github.com/lib/pq" // ...
)

type Store struct {
	config *Config
	db     *sql.DB
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// Open для подключения к бд
func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL) // ленивое создание соединения когда будет совершен первый вызов
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}
	s.db = db

	return nil
}

// Close для зотключения от бд
func (s *Store) Close() {
	s.db.Close()
}
