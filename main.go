package main

import (
	"fmt"
	"log"
	"net/http"
	"school-api/config"
)

func main() {
	db := config.InitDB()
	router := config.InitRouter(db)

	fmt.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
