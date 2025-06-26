package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	port := 5000

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[%s] Received /hello - 200 OK\n", time.Now().Format("2006-01-02T15:04:05Z07:00"))
		io.WriteString(w, "Hello, World!")
	})

	fmt.Printf("Starting server at %d\n", port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed gracefully\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
