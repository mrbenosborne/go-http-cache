package server

import (
	"context"
	"net"
	"net/http"

	"github.com/mrbenosborne/go-http-cache/store"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Server a HTTP server for interacting
// with cache stores.
type Server struct {
	Ready  chan struct{}
	router *chi.Mux
	l      net.Listener
	store  *store.Store
}

// New initialize a new server and
// setup routes/middleware.
func New() *Server {
	router := setMiddleware(chi.NewRouter())

	s := &Server{
		Ready:  make(chan struct{}, 1),
		router: router,
		store:  store.New(),
	}
	s.routes()
	return s
}

// routes set routes against the router
// and return.
func (s *Server) routes() {
	s.router.Get("/", StatusHandler())
	s.router.Get("/{key}", GetHandler(s.store))
	s.router.Put("/{key}", SetHandler(s.store).Func())
}

// setMiddleware set middleware to use on the
// router.
func setMiddleware(mux *chi.Mux) *chi.Mux {
	mux.Use(
		middleware.DefaultLogger,
	)
	return mux
}

// Serve start accepting requests, an empty
// struct will be sent to channel "Ready",
// this signals that the server is ready to accept
// requests.
//
// This method should be called via a go routine
// to avoid blocking.
func (s *Server) Serve(ctx context.Context) error {
	var err error
	s.l, err = net.Listen("tcp", ":8901")
	if err != nil {
		return err
	}

	s.Ready <- struct{}{}
	if err := http.Serve(s.l, s.router); err != nil {
		return err
	}
	return nil
}

// Close close the listener.
func (s *Server) Close() error {
	return s.l.Close()
}

// GetRouter return the instance of chi router.
func (s *Server) GetRouter() chi.Router {
	return s.router
}
