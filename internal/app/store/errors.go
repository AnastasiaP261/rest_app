package store

import "errors"

var (
	// ErrRecordNotFound общая ошибка для всех хранилищ
	ErrRecordNotFound = errors.New("record not found")
)
