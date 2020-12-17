package http

import (
	"net/http"
	"time"
)

func ListenAndServe(router http.Handler) error {
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	return srv.ListenAndServe()
}
