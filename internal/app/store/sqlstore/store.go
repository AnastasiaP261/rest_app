package sqlstore

import (
	"database/sql"
	_ "github.com/lib/pq" // ...
	"rest_app/internal/app/store"
)

type Store struct {
	db             *sql.DB
	userRepository *UserRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// User для защиты репозитория от внешнего мира
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		store: s,
	}
	return s.userRepository
}

//
//// Open для подключения к бд
//func (s *Store) Open() error {
//	db, err := sql.Open("postgres", s.db.DatabaseURL) // ленивое создание соединения когда будет совершен первый вызов
//	if err != nil {
//		return err
//	}
//
//	if err := db.Ping(); err != nil {
//		return err
//	}
//	s.db = db
//
//	return nil
//}
//
//// Close для зотключения от бд
//func (s *Store) Close() {
//	s.db.Close()
//}
