package negroni

import (
	"log"
	"net/http"
	"os"
	"time"
)

// Logger is a middleware handler that logs the request as it goes in and the response as it goes out.
type Logger struct {
	// Logger inherits from log.Logger used to log messages with the Logger middleware
	*log.Logger
}

// NewLogger returns a new Logger instance
func NewLogger() *Logger {
	return &Logger{log.New(os.Stdout, "[Have] ", 0)}
}

func (l *Logger) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	start := time.Now()

	next(rw, r)

	res := rw.(ResponseWriter)
	l.Printf("Now is %s, %s, %s, %s, %v, %s, %v", time.Now().Format("2006-01-02:15:04:05"), r.Header.Get("token"), r.Method, r.RequestURI, res.Status(), http.StatusText(res.Status()), time.Since(start))
}
