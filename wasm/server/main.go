package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("Serving on port :8080")
	err := http.ListenAndServe(":8080", http.FileServer(http.Dir("wasm/assets")))
	if err != nil {
		fmt.Println("Failed to start server", err)
		return
	}
}
