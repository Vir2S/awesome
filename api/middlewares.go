package api

import (
	"awesome/database"
	"net/http"
)

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ApiKey := r.Header.Get("Api-key")
		validAPIKeys, err := database.GetAllApiKeysFromDatabase()
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		for _, key := range validAPIKeys {
			if ApiKey == key.ApiKey {
				next.ServeHTTP(w, r)
				return
			}
		}
		http.Error(w, "Forbidden", http.StatusForbidden)
	})
}
