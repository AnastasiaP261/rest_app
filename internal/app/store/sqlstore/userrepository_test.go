package sqlstore_test

import (
	"github.com/stretchr/testify/assert"
	"rest_app/internal/app/model"
	"rest_app/internal/app/store"
	"rest_app/internal/app/store/sqlstore"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")
	s := sqlstore.New(db)

	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u)) // нет ошибки
	assert.NotNil(t, u)                   // юзер не nil
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")
	s := sqlstore.New(db)

	// кейс: пытаемся найти юзера не существующего в бд и получаем ошибку
	u := model.TestUser(t)
	_, err := s.User().FindByEmail(u.Email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	// кейс: польз-ль существует в бд и все хорошо
	s.User().Create(u)
	u, err = s.User().FindByEmail(u.Email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_Find(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)

	// кейс: польз-ль существует в бд и все хорошо
	u := model.TestUser(t)
	s.User().Create(u)
	u, err := s.User().Find(u.Id)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
