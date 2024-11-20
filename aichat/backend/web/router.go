package web

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"net/http"
	"strings"
)

type Handler interface {
	HandleChat(w http.ResponseWriter, r *http.Request)
	HandleLogin(w http.ResponseWriter, r *http.Request)
	HandlePing(w http.ResponseWriter, r *http.Request)
}

func NewRouter(handler Handler) *chi.Mux {
	r := chi.NewRouter()

	corsOptions := cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},        // Allow GET, POST, and OPTIONS methods
		AllowedHeaders:   []string{"Content-Type", "Authorization"}, // Allow Content-Type header
		AllowCredentials: true,                                      // Disable credentials (cookies, authorization headers, etc.)
		MaxAge:           300,                                       // Set preflight request cache time (seconds)
	}
	r.Use(cors.New(corsOptions).Handler)

	r.With(authorizationMiddleware).Route("/chat", func(r chi.Router) {
		r.Post("/", handler.HandleChat)
		// You can add more routes here like r.Post("/", customerHandler.HandleCreateCustomer)
	})

	// Add the authorization middleware only for routes except "/login"
	r.Post("/login", handler.HandleLogin)
	r.Get("/ping", handler.HandlePing)

	return r
}

func authorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Perform authorization logic here
		if !isAuthorized(r) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// If authorization succeeds, call the next middleware or handler
		next.ServeHTTP(w, r)
	})
}

func isAuthorized(r *http.Request) bool {
	authHeader := r.Header.Get("Authorization")
	if authHeader != "" {
		token := strings.Split(authHeader, "Bearer ")
		if len(token) == 2 {
			// Check if the token is valid
			valid := validateToken(token[1])
			return valid
		}
	}
	return false
}

func validateToken(token string) bool {
	_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("your_secret_key"), nil
	})
	return err == nil
}
