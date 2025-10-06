package main

import (
	"autosuggest/internal/api"
	"autosuggest/internal/app"
	"fmt"
	"net/http"
)

func main() {
	appService := app.NewService()
	handler := api.NewHandler(appService)

	http.HandleFunc("/suggest", handler.Suggest)
	http.HandleFunc("/phrase", handler.AddPhrase)

	fmt.Println("server running...")
	http.ListenAndServe(":8080", nil)
}
