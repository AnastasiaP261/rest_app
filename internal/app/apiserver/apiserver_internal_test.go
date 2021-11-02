package apiserver

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIServer_HandleHello(t *testing.T) {
	s := New(NewConfig()) // сервер
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil) // объект запроса
	s.handleHello().ServeHTTP(rec, req)                      // вызов метода сервера
	assert.Equal(t, rec.Body.String(), "Hello")
}
