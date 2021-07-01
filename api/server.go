package server

import (
	"github.com/codegangsta/negroni"
	"github.com/erbilsilik/getir-go-challange/api/handler"
	"github.com/erbilsilik/getir-go-challange/api/middleware"
	"github.com/erbilsilik/getir-go-challange/infrastructure/repository"
	"github.com/erbilsilik/getir-go-challange/pkg/metric"
	"github.com/erbilsilik/getir-go-challange/pkg/mongodb"
	"github.com/erbilsilik/getir-go-challange/usecase/configuration"
	"github.com/erbilsilik/getir-go-challange/usecase/record"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
	configurationRepository := repository.NewConfigurationRepository();

	// services
	recordService := record.NewService(recordRepository)
	configurationService := configuration.NewService(configurationRepository)

	// metrics
	metricService, err := metric.NewPrometheusService()
	if err != nil {
		log.Fatal(err.Error())
	}

	// handlers
	r := mux.NewRouter()
	http.Handle("/", r)
	http.Handle("/metrics", promhttp.Handler())
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// middlewares
	n := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
		middleware.Metrics(metricService),
		negroni.HandlerFunc(middleware.EnforceJSONMiddleware),
	)

	// handlers
	handler.MakeRecordHandlers(r, *n, recordService)
	handler.MakeConfigurationHandlers(r, *n, configurationService)

	// logger
	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)

	// server
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":" + os.Getenv("API_PORT"),
		Handler:      context.ClearHandler(http.DefaultServeMux),
		ErrorLog: logger,
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
