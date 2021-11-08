package teststore_test

import (
	"github.com/stretchr/testify/assert"
	"rest_app/internal/app/model"
	"rest_app/internal/app/store"
	"rest_app/internal/app/store/teststore"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s := teststore.New()

	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u)) // нет ошибки
	assert.NotNil(t, u)                   // юзер не nil
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s := teststore.New()
	email := "users@examples.org"

	// кейс: пытаемся найти юзера не существующего в бд и получаем ошибку
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	// кейс: польз-ль существует в бд и все хорошо
	u1 := model.TestUser(t)
	u1.Email = email
	s.User().Create(u1)
	u2, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}

func TestUserRepository_Find(t *testing.T) {
	s := teststore.New()
	u1 := model.TestUser(t)
	s.User().Create(u1)
	u2, err := s.User().Find(u1.Id)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}
