package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/codegangsta/negroni"
	"github.com/erbilsilik/getir-go-challange/api/presenter"
	"github.com/erbilsilik/getir-go-challange/entity"
	"github.com/erbilsilik/getir-go-challange/usecase/record"
	"github.com/gorilla/mux"
)

func getFilteredRecords(service record.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading records"
		var data []*entity.Record
		var err error
		minCount, _ := strconv.Atoi(r.URL.Query().Get("minCount"))
		maxCount, _ := strconv.Atoi(r.URL.Query().Get("minCount"))
		q := record.FindAvailableRecordsQuery{
			StartDate: r.URL.Query().Get("startDate"),
			EndDate:   r.URL.Query().Get("endDate"),
			MinCount:  minCount,
			MaxCount:  maxCount,
		}
		data, err = service.ListRecords(&q)
		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			_, err := w.Write([]byte(errorMessage))
			if err != nil {
				return
			}
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			_, err := w.Write([]byte(errorMessage))
			if err != nil {
				return
			}
			return
		}
		var toJ []*presenter.Record
		for _, d := range data {
			toJ = append(toJ, &presenter.Record{
				Key:        d.Key,
				TotalCount: d.TotalCount,
				// CreatedAt:  d.CreatedAt,
			})
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, err := w.Write([]byte(errorMessage))
			if err != nil {
				return
			}
		}
	})
}

func createRecord(service record.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding record"
		var input struct {
			Key        string `json:"key"`
			TotalCount int    `json:"totalCount"`
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err := w.Write([]byte(errorMessage))
			if err != nil {
				return
			}
			return
		}
		err = service.CreateRecord(input.Key, input.TotalCount)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, err := w.Write([]byte(errorMessage))
			if err != nil {
				return
			}
			return
		}
		toJ := &presenter.Record{
			Key:        input.Key,
			TotalCount: input.TotalCount,
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, err := w.Write([]byte(errorMessage))
			if err != nil {
				return
			}
			return
		}
	})
}

func MakeBookHandlers(r *mux.Router, n negroni.Negroni, service record.UseCase) {
	r.Handle("/v1/records", n.With(
		negroni.Wrap(getFilteredRecords(service)),
	)).Methods("GET", "OPTIONS").Name("listRecords")

	r.Handle("/v1/records", n.With(
		negroni.Wrap(createRecord(service)),
	)).Methods("POST", "OPTIONS").Name("createRecord")
}
