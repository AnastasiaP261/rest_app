package model_test

import (
	"github.com/stretchr/testify/assert"
	"rest_app/internal/app/model"
	"testing"
)

func TestUserValidate(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *model.User
		isValid bool
	}{
		{
			name: "valid",
			u: func() *model.User {
				return model.TestUser(t)
			},
			isValid: true,
		},
		{
			name: "empty_email",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = ""
				return u
			},
			isValid: false,
		},
		{
			name: "invalid_email",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = "invalid@"
				return u
			},
			isValid: false,
		},
		{
			name: "with_encrypt_password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = ""
				u.EncryptedPassword = "encryptedpasswordencryptedpasswordencryptedpassword"
				return u
			},
			isValid: true,
		},
		{
			name: "empty_password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = ""
				return u
			},
			isValid: false,
		},
		{
			name: "short_password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = "12 a"
				return u
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().Validate())
			} else {
				assert.Error(t, tc.u().Validate())
			}
		})
	}
}

func TestUser_BeforeCreate(t *testing.T) {
	u := model.TestUser(t)

	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.EncryptedPassword)
}
