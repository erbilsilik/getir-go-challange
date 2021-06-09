package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/erbilsilik/getir-go-challange/pkg/mongodb"
	"github.com/gorilla/mux"
)

var collectionName = "records"

func init() {
	client, _ := mongodb.New(
		os.Getenv("MONGODB_URI"),
		os.Getenv("MONGODB_DB"),
		collectionName,
	)
	client.FindAll(context.TODO())
}

func main() {
	r := mux.NewRouter()

	http.Handle("/", r)
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)

	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":" + "8081",
		// Handler:      context.ClearHandler(http.DefaultServeMux),
		ErrorLog: logger,
	}

	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err.Error())
	}
}
