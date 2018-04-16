package server

import (
	"log"
	"net/http"

	"code.ysitd.cloud/component/aviation/sky/pkg/modals/airline"
	"code.ysitd.cloud/component/aviation/sky/pkg/modals/flyer"
	"context"
	"github.com/gorilla/handlers"
	"github.com/sirupsen/logrus"
	"time"
)

const requestTimeout = 30 * time.Second

type Service struct {
	Logger   logrus.FieldLogger
	Hostname *flyer.Store
	Airline  airline.Store
}

func (s *Service) CreateServer(addr string) *http.Server {
	handler := s.createHandler()

	logWriter := s.Logger.WithField("source", "access").Writer()
	errWriter := s.Logger.WithField("source", "http").WriterLevel(logrus.ErrorLevel)
	recoverLogger := s.Logger.WithField("source", "recover")

	handler = handlers.CombinedLoggingHandler(logWriter, handler)
	handler = handlers.RecoveryHandler(
		handlers.RecoveryLogger(recoverLogger),
		handlers.PrintRecoveryStack(true),
	)(handler)

	return &http.Server{
		Addr:     addr,
		Handler:  handler,
		ErrorLog: log.New(errWriter, "", 0),
	}
}

func (s *Service) createHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
		defer cancel()
		f, err := s.Hostname.GetFlyer(ctx, r.Host)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else if f == nil {
			http.NotFound(w, r)
			return
		}

		a, err := s.Airline.GetRevision(ctx, f.Revision)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		handler := a.GetHandler()

		handler.ServeHTTP(w, r.WithContext(ctx))
	})
}
