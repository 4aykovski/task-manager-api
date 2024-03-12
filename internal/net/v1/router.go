package v1

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/4aykovski/task-manager-api/internal/net/v1/middleware"
	"github.com/4aykovski/task-manager-api/pkg/types"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

type tokenManager interface {
	Parse(token string) (string, error)
	CreateTokensPair(userId string, ttl time.Duration) (types.Tokens, error)
}

func NewMux(
	log *slog.Logger,
	tokenManager tokenManager,
) *chi.Mux {
	var (
		mux = chi.NewMux()
		mw  = middleware.New(tokenManager)
	)

	mux.Use(chiMiddleware.RealIP)
	mux.Use(chiMiddleware.RequestID)
	mux.Use(chiMiddleware.Recoverer)
	mux.Use(mw.Logger(log))

	mux.Route("/api/v1", func(r chi.Router) {
		r.Get("/test", func(writer http.ResponseWriter, request *http.Request) {
			fmt.Fprint(writer, "asd")
		})
	})

	return mux
}
