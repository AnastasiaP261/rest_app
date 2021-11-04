package sqlstore

import (
	"database/sql"
	"fmt"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"
	"strings"
	"testing"
)

// TestStore создание бд для тестов
func TestDB(t *testing.T, databaseURL string) (*sql.DB, func(...string)) {
	t.Helper() // значит что это тестовый метод, его не надо тестировать

	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}

	return db, func(tables ...string) {
		if len(tables) > 0 {
			res, err := db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", ")))
			if err != nil {
				fmt.Println("Таблицы не очистились после теста. Ошибка: ", res)
			} else {
				fmt.Println("Таблицы после теста успешно очищены.")
			}
		}
		db.Close()
	}
}

//	config := NewConfig()
//	config.DatabaseURL = databaseURL
//	s := New(config)
//	if err := s.Open(); err != nil { // попытка подключения
//		t.Fatal(err)
//	}
//
//	return s, func(tables ...string) {
//		if len(tables) > 0 {
//			if _, err := s.db.Exec(fmt.Sprintf(
//				"TRUNCATE %s CASCADE", strings.Join(tables, ", "))); err != nil {
//				t.Fatal(err)
//			}
//		}
//
//		s.Close()
//	}
//}
