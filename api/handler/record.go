package handler

import (
	"encoding/json"
	"github.com/codegangsta/negroni"
	recordMiddleware "github.com/erbilsilik/getir-go-challange/api/middleware/record"
	"github.com/erbilsilik/getir-go-challange/api/presenter"
	"github.com/erbilsilik/getir-go-challange/entity"
	"github.com/erbilsilik/getir-go-challange/pkg/utilities"
	"github.com/erbilsilik/getir-go-challange/usecase/record"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func getRecordsFilteredByTimeAndTotalCountInGivenNumberRange(service record.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading records"
		var records []*entity.Record
		var err error

		// convert
		minCount, _ := strconv.Atoi(r.URL.Query().Get("minCount"))
		maxCount, _ := strconv.Atoi(r.URL.Query().Get("maxCount"))
		layout := "2006-01-02"
		startDateParsed := utilities.ParseDate(layout, r.URL.Query().Get("startDate"))
		endDateParsed := utilities.ParseDate(layout, r.URL.Query().Get("endDate"))

		// create query struct
		q := record.RecordsFilteredByTimeAndTotalCountInGivenNumberRangeQuery{
			StartDate: startDateParsed,
			EndDate:   endDateParsed,
			MinCount:  minCount,
			MaxCount:  maxCount,
		}
		records, err = service.GetRecordsFilteredByTimeAndTotalCountInGivenNumberRange(&q)

		// server error
		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			_, err := w.Write([]byte(errorMessage))
			if err != nil {
				return
			}
			return
		}

		// no data found
		if records == nil {
			w.WriteHeader(http.StatusNotFound)
			_, err := w.Write([]byte(errorMessage))
			if err != nil {
				return
			}
			return
		}

		// serialize records data and format for response
		var recordPresenter []*presenter.Record
		for _, r := range records {
			recordPresenter = append(recordPresenter, &presenter.Record{
				Key:        r.Key,
				TotalCount: r.TotalCount,
				CreatedAt: r.CreatedAt,
			})
		}
		response := presenter.Response{
			Code: http.StatusOK,
			Msg: "Success",
			Records: recordPresenter,
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
		negroni.HandlerFunc(
			recordMiddleware.ValidateGetRecordsFilteredByTimeAndTotalCountInGivenNumberRange,
		),
		negroni.Wrap(getRecordsFilteredByTimeAndTotalCountInGivenNumberRange(service)),
	)).Methods("GET", "OPTIONS").Name("getRecordsFilteredByTimeAndTotalCountInGivenNumberRange")
}
