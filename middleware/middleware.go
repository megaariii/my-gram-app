package middleware

import (
	"context"
	"my-gram/helper"
	"my-gram/model/domain"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt"
)

var ctxKey = &contextKey{"user"}

type contextKey struct {
	data string
}

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if !strings.Contains(header, "Bearer") {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(header, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		//validate jwt token
		token, err := helper.ValidateToken(tokenString)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		payload, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		userID := payload["id"].(string)

		id, _ := strconv.Atoi(userID)
		user := domain.User{ID: id}

		ctx := context.WithValue(r.Context(), ctxKey, &user)

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func ForContext(ctx context.Context) *domain.User {
	rawData, _ := ctx.Value(ctxKey).(*domain.User)
	return rawData
}