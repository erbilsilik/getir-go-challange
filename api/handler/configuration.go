package handler

import (
	"encoding/json"
	"github.com/codegangsta/negroni"
	"github.com/erbilsilik/getir-go-challange/api/presenter"
	"github.com/erbilsilik/getir-go-challange/entity"
	"github.com/erbilsilik/getir-go-challange/usecase/configuration"
	"github.com/gorilla/mux"
	"net/http"
)

func createConfig(service configuration.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var config *entity.Config
		var err error

		err = json.NewDecoder(r.Body).Decode(&config)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = service.Create(config.Key, config.Value)

		presenter.JSON(
			w,
			http.StatusCreated,
			config,
		)
	})
}

func getConfig(service configuration.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")

		val := service.FindByKey(key)

		presenter.JSON(
			w,
			http.StatusOK,
			val,
		)
	})
}

func MakeConfigurationHandlers(r *mux.Router, n negroni.Negroni, service configuration.UseCase) {
	r.Handle("/v1/configurations", n.With(
		negroni.Wrap(createConfig(service)),
	)).Methods("POST", "OPTIONS").Name("createConfig");
	r.Handle("/v1/configurations", n.With(
		negroni.Wrap(getConfig(service)),
	)).Methods("GET", "OPTIONS").Name("getConfig");
}
