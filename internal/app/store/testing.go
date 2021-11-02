package store

import (
	"fmt"
	"strings"
	"testing"
)

// TestStore создание бд для тестов
func TestStore(t *testing.T, databaseURL string) (*Store, func(...string)) {
	t.Helper() // значит что это тестовый метод, его не надо тестировать

	config := NewConfig()
	config.DatabaseURL = databaseURL
	s := New(config)
	if err := s.Open(); err != nil { // попытка подключения
		t.Fatal(err)
	}

	return s, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := s.db.Exec(fmt.Sprintf(
				"TRUNCATE %s CASCADE", strings.Join(tables, ", "))); err != nil {
				t.Fatal(err)
			}
		}

		s.Close()
	}
}
