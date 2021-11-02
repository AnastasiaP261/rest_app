package store_test

import (
	"github.com/stretchr/testify/assert"
	"rest_app/internal/app/model"
	"rest_app/internal/app/store"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u, err := s.User().Create(&model.User{
		Email: "user@example.org",
	})
	assert.NoError(t, err) // нет ошибки
	assert.NotNil(t, u)    // юзер не nil
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	email := "users@examples.org"
	// кейс: пытаемся найти юзера не существующего в бд и получаем ошибку
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	// кейс: польз-ль существует в бд и все хорошо
	s.User().Create(&model.User{
		Email: email,
	})
	u, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
