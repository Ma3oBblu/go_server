package main

import (
	"fmt"
	"go_server/internal/router"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	router.RegisterRoutes(mux)

	server := http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("starting server at :8080")
	fmt.Println("http://localhost:8080")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("error starting server at :8080")
	}
}
