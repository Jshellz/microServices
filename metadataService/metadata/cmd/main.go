package main

import (
	"books_microservices/metadataService/metadata/internal/controller"
	"books_microservices/metadataService/metadata/internal/handler"
	"books_microservices/metadataService/metadata/internal/repository"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting the movie metadata service")
	repo := repository.New()
	ctrl := controller.New(repo)
	h := handler.New(ctrl)
	http.Handle("/metadata", http.HandlerFunc(h.GetMetadata))
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		panic(err)
	}
}
