package apiserver

import (
	"io"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/sotanodroid/gophiato/internal/app/store"
)

// APIServer ...
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

// NewAPIServer returns new instance of api server
func NewAPIServer(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start starts new server
func (s *APIServer) Start() error {
	s.configureLogger()
	s.configureStore()
	s.configureRouter()

	s.logger.Info("Sarting service")

	return http.ListenAndServe(
		net.JoinHostPort("", s.config.Bindport),
		s.router,
	)
}

func (s *APIServer) configureLogger() {
	level, err := logrus.ParseLevel(s.config.Loglevel)
	if err != nil {
		s.logger.Error(err)
	}

	s.logger.SetLevel(level)
}

func (s *APIServer) configureStore() {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		s.logger.Error(err)
	}

	s.store = st
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handlehello())
}

func (s *APIServer) handlehello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}
