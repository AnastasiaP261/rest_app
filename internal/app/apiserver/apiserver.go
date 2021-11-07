package apiserver

import (
	"database/sql"
	"github.com/gorilla/sessions"
	"net/http"
	"rest_app/internal/app/store/sqlstore"
)

func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}

	defer db.Close()
	store := sqlstore.New(db)
	sessionStore := sessions.NewCookieStore([]byte(config.SessionKey))
	srv := newServer(store, sessionStore)

	return http.ListenAndServe(config.BindAddr, srv)
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

//import (
//	"github.com/gorilla/mux"
//	"github.com/sirupsen/logrus"
//	"io"
//	"net/http"
//	"rest_app/internal/app/store/sqlstore"
//)
//
//type APIServer struct {
//	config *Config
//	logger *logrus.Logger
//	router *mux.Router
//	store  *sqlstore.Store
//}
//
//func New(config *Config) *APIServer {
//	return &APIServer{
//		config: config,
//		logger: logrus.New(),
//		router: mux.NewRouter(),
//	}
//}
//
//func (s *APIServer) Start() error {
//	if err := s.configureLogger(); err != nil {
//		return err
//	}
//
//	s.configureRouter()
//
//	if err := s.configureStore(); err != nil {
//		return err
//	}
//
//	s.logger.Info("Starting API server. ")
//
//	return http.ListenAndServe(s.config.BindAddr, s.router)
//}
//
//// функция для создания подключения к бд
//func (s *APIServer) configureStore() error {
//	st := sqlstore.New(s.config.Store)
//	if err := st.Open(); err != nil {
//		return err
//	}
//	s.store = st
//	return nil
//}
//
//// функция конфигурации роутера
//func (s *APIServer) configureRouter() {
//	s.router.HandleFunc("/hello", s.handleHello())
//}
//
//func (s *APIServer) handleHello() http.HandlerFunc {
//	// код здесь выполнится один раз
//
//	return func(w http.ResponseWriter, r *http.Request) {
//		io.WriteString(w, "Hello")
//	}
//}
//
//// функция конфигурации логгера
//func (s *APIServer) configureLogger() error {
//	level, err := logrus.ParseLevel(s.config.LogLevel)
//	if err != nil {
//		return err
//	}
//	s.logger.SetLevel(level)
//
//	return nil
//}
