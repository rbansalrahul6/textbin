package main

import "net/http"

/* template for creating a middleware */
// func myMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// TODO: your middleware logic
// 		next.ServeHTTP(w, r)
// 	})
// }

// middleware to add secuirty headres to each request
func secureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("X-Frame-Options", "deny")

		next.ServeHTTP(w, r)
	})
}

// middleware to log requests
func (app *application) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.infoLog.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())

		next.ServeHTTP(w, r)
	})
}
