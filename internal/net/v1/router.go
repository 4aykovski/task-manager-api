package v1

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/4aykovski/task-manager-api/internal/net/v1/middleware"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

func NewMux(log *slog.Logger) *chi.Mux {
	var (
		mux = chi.NewMux()
		mw  = middleware.New()
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
