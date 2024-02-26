package main

import (
	"books_microservices/ratingService/rating/internal/controller"
	hp "books_microservices/ratingService/rating/internal/handler/http"
	"books_microservices/ratingService/rating/internal/repository"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting the rating service")
	repo := repository.New()
	ctrl := controller.New(repo)
	h := hp.New(ctrl)
	http.Handle("/rating", http.HandlerFunc(h.Handle))
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		panic(err)
	}
}
