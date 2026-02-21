package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	fmt.Println("server: hello handle started...")
	defer fmt.Println("Server : hello handle ended")

	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, " hello \n")
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("server error", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8070", nil)
}
