package handler

import (
	"encoding/json"
	"github.com/codegangsta/negroni"
	"github.com/erbilsilik/getir-go-challange/api/presenter"
	"github.com/erbilsilik/getir-go-challange/entity"
	"github.com/erbilsilik/getir-go-challange/pkg/utilities"
	"github.com/erbilsilik/getir-go-challange/usecase/record"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func getFilteredRecords(service record.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading records"
		var records []*entity.Record
		var err error

		minCount, _ := strconv.Atoi(r.URL.Query().Get("minCount"))
		maxCount, _ := strconv.Atoi(r.URL.Query().Get("maxCount"))
		layout := "2006-01-02"
		startDateParsed := utilities.ParseDate(layout, r.URL.Query().Get("startDate"))
		endDateParsed := utilities.ParseDate(layout, r.URL.Query().Get("endDate"))

		q := record.CalculateRecordsTotalCountQuery{
			StartDate: startDateParsed,
			EndDate:   endDateParsed,
			MinCount:  minCount,
			MaxCount:  maxCount,
		}

		records, err = service.CalculateRecordsTotalCount(&q)
		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			_, err := w.Write([]byte(errorMessage))
			if err != nil {
				return
			}
			return
		}

		if records == nil {
			w.WriteHeader(http.StatusNotFound)
			_, err := w.Write([]byte(errorMessage))
			if err != nil {
				return
			}
			return
		}
		var toJ []*presenter.Record
		for _, r := range records {
			toJ = append(toJ, &presenter.Record{
				Key:        r.Key,
				TotalCount: r.TotalCount,
				CreatedAt: r.CreatedAt,
			})
		}
		response := presenter.Response{
			Code: 0,
			Msg: "Success",
			Records: toJ,
		}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, err := w.Write([]byte(errorMessage))
			if err != nil {
				return
			}
		}
	})
}

func MakeRecordHandlers(r *mux.Router, n negroni.Negroni, service record.UseCase) {
	r.Handle("/v1/records", n.With(
		negroni.Wrap(getFilteredRecords(service)),
	)).Methods("GET", "OPTIONS").Name("listRecords")
}
