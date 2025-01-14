package middleware

import (
	"log"
	"net/http"
	"os"
	"time"
)

// AccessLogMiddleware logs API access details to a file
func AccessLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Open or create the log file
		file, err := os.OpenFile("/var/log/access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Printf("Error opening log file: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		// Capture start time
		startTime := time.Now()

		// Use a ResponseWriter wrapper to capture the status code
		logWriter := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(logWriter, r)

		// Log the request details
		log.SetOutput(file)
		log.Printf("[%s] %s %s %d %s %dms\n",
			r.Method,
			r.RemoteAddr,
			r.URL.Path,
			logWriter.statusCode,
			r.UserAgent(),
			time.Since(startTime).Milliseconds(),
		)
	})
}

// responseWriter wraps http.ResponseWriter to capture status codes
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
