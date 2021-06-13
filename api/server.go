package server

import (
	"github.com/codegangsta/negroni"
	"github.com/erbilsilik/getir-go-challange/api/handler"
	"github.com/erbilsilik/getir-go-challange/api/middleware"
	"github.com/erbilsilik/getir-go-challange/infrastructure/repository"
	"github.com/erbilsilik/getir-go-challange/pkg/mongodb"
	"github.com/erbilsilik/getir-go-challange/usecase/record"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

func Run() {
	// db
	mongodb.New(
		os.Getenv("MONGODB_URI"),
		os.Getenv("MONGODB_DB"),
	)

	// repositories
	recordRepository := repository.NewRecordRepositoryMongoDB()
	recordService := record.NewService(recordRepository)

	// handlers
	r := mux.NewRouter()
	http.Handle("/", r)
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	n := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
		negroni.HandlerFunc(middleware.EnforceJSONMiddleware),
	)

	// record
	handler.MakeRecordHandlers(r, *n, recordService)

	// logger
	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)

	// server
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":" + os.Getenv("API_PORT"),
		//Handler:      context.ClearHandler(http.DefaultServeMux),
		ErrorLog: logger,
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
