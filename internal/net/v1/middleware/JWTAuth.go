package middleware

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/4aykovski/task-manager-api/pkg/libs/response"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

const authorizationHeader = "Authorization"

var ErrInvalidAuthHeader = errors.New("invalid auth header")

func (m *Middleware) JWTAuth(log *slog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log := log.With("request_id", middleware.GetReqID(r.Context()))

			id, err := m.parseAuthToken(r)
			if err != nil {
				log.Info("Authorization error")

				render.Status(r, http.StatusUnauthorized)
				render.JSON(w, r, response.UnauthorizedError())
				return
			}

			ctx := context.WithValue(r.Context(), "user_id", id)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func (m *Middleware) parseAuthToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get(authorizationHeader)
	if authHeader == "" {
		return "", ErrInvalidAuthHeader
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", ErrInvalidAuthHeader
	}

	if len(headerParts[1]) == 0 {
		return "", ErrInvalidAuthHeader
	}

	claims, err := m.tokenManager.Parse(headerParts[1])
	if err != nil {
		return "", fmt.Errorf("%w: %s", ErrInvalidAuthHeader, err)
	}

	return claims, nil
}
