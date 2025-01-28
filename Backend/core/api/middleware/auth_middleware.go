package middleware

import "net/http"

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Check if the user is authenticated
        if !isAuthenticated(r) {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        next.ServeHTTP(w, r)
    })
}

func isAuthenticated(r *http.Request) bool {
    // Logic to check if the user is authenticated
    r.Body = nil
    return true
}