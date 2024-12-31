package middleware

import (
	"log"
	"net/http"
	"time"
)

type WrappedRW struct {
	http.ResponseWriter
	StatusCode int
	Flusher    http.Flusher
}

func (w *WrappedRW) Flush() {
	if w.Flusher != nil {
		w.Flusher.Flush()
	}
}

func (w *WrappedRW) WriteHeader(statusCode int) {
	w.StatusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrappedRW := &WrappedRW{ResponseWriter: w}

		if flusher, ok := w.(http.Flusher); ok {
			wrappedRW.Flusher = flusher
		}

		next.ServeHTTP(wrappedRW, r)

		duration := time.Since(start)

		// Define ANSI escape codes for colors
		cyan := "\033[36m"
		green := "\033[32m"
		yellow := "\033[33m"
		magenta := "\033[35m"
		reset := "\033[0m"

		// Print the formatted and colored log message using ANSI codes
		log.Printf(
			"%s%s %s%s %s[%d] %s%s%s",
			cyan, r.RemoteAddr, // Remote Address in cyan
			green, r.Method, // HTTP Method in green
			yellow, wrappedRW.StatusCode, // Status code in yellow (wrapped in brackets)
			magenta, duration, // Duration in magenta
			reset, // Reset color
		)
	})
}
