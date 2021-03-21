package apiserver

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

//APIServer ...
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

//New ...
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

//Start ...
func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	//if err := s.configureStore(); err != nil {
	//	return err
	//}

	s.logger.Info("starting server")

	return http.ListenAndServe(s.config.BindAdress, s.router)
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)
	return nil
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/ping", s.handlePing())
}

//func (s *APIServer) configureStore() error {
//	st := store.New(s.config.Store)
//	if err := st.Open(); err != nil {
//		return err
//	}
//
//	s.store = st
//
//	return nil
//}

func (s *APIServer) handlePing() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "pong")
	}
}
