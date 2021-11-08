package apiserver

import "net/http"

type responseWriter struct {
	http.ResponseWriter // анонимное поле для того чтобы иметь доступ к методам ResponseWriter
	code                int
}

func (w *responseWriter) WriteHeader(statusCode int) {
	w.code = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
