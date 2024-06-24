package main

import (
	"net/http"

	"github.com/oliver-day/rss-feed-aggregator/internal/auth"
	"github.com/oliver-day/rss-feed-aggregator/internal/database"
)

type authedHander func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHander) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, "Failed to provide API key")
			return
		}

		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, "Failed find authenticated user")
			return
		}
		handler(w, r, user)
	}
}
