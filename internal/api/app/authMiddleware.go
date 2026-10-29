package app

import (
	"net/http"
	"strings"

	"github.com/codacube/go-microservice-banking/internal/api/domain"
	"github.com/codacube/go-microservice-banking/internal/errs"
	"github.com/gorilla/mux"
)

type AuthMiddleware struct {
	repo domain.AuthRepository
}

func (a AuthMiddleware) authorisationHandler() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			currentRoute := mux.CurrentRoute(r)
			currentRouteVars := mux.Vars(r)
			authHeader := r.Header.Get("Authorization")

			token := getTokenFromHeader(authHeader)
			if token == "" {
				appError := errs.AppError{Code: http.StatusForbidden, Message: "Missing token"}
				writeResponse(w, appError.Code, appError.AsMessage())
				return
			}

			isAuthorised := a.repo.IsAuthorised(token, currentRoute.GetName(), currentRouteVars)
			if !isAuthorised {
				appError := errs.AppError{Code: http.StatusForbidden, Message: "Unauthorized"}
				writeResponse(w, appError.Code, appError.AsMessage())
				return
			}

			// Next middleware or router
			next.ServeHTTP(w, r)
		})
	}
}

func getTokenFromHeader(header string) string {
	/*
	   token is coming in the format as below
	   "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50cyI6W.yI5NTQ3MCIsIjk1NDcyIiw"
	*/

	// Check for blank header
	if header == "" {
		return ""
	}

	splitToken := strings.Split(header, "Bearer")
	if len(splitToken) == 2 {
		return strings.TrimSpace(splitToken[1])
	}

	// No token
	return ""
}
