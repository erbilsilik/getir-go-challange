package handler

import (
	"errors"
	"github.com/codegangsta/negroni"
	recordMiddleware "github.com/erbilsilik/getir-go-challange/api/middleware/record"
	"github.com/erbilsilik/getir-go-challange/api/presenter"
	recordPres "github.com/erbilsilik/getir-go-challange/api/presenter/record"
	"github.com/erbilsilik/getir-go-challange/entity"
	"github.com/erbilsilik/getir-go-challange/pkg/utilities"
	"github.com/erbilsilik/getir-go-challange/usecase/record"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func getRecordsFilteredByTimeAndTotalCountInGivenNumberRange(service record.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var records []*entity.Record
		var err error

		minCount, _ := strconv.Atoi(r.URL.Query().Get("minCount"))
		maxCount, _ := strconv.Atoi(r.URL.Query().Get("maxCount"))
		layout := "2006-01-02"
		startDateParsed := utilities.ParseDate(layout, r.URL.Query().Get("startDate"))
		endDateParsed := utilities.ParseDate(layout, r.URL.Query().Get("endDate"))

		q := record.RecordsFilteredByTimeAndTotalCountInGivenNumberRangeQuery{
			StartDate: startDateParsed,
			EndDate:   endDateParsed,
			MinCount:  minCount,
			MaxCount:  maxCount,
		}
		records, err = service.GetRecordsFilteredByTimeAndTotalCountInGivenNumberRange(&q)

		if err != nil {
			presenter.ERROR(w, http.StatusInternalServerError, errors.New("error reading records"))
			return
		}

		var recordPresenter []*recordPres.Record

		if records == nil {
			presenter.JSON(w, http.StatusOK, recordPresenter)
			return
		}

		for _, r := range records {
			recordPresenter = append(recordPresenter, &recordPres.Record{
				Key:        r.Key,
				TotalCount: r.TotalCount,
				CreatedAt:  r.CreatedAt,
			})
		}
		presenter.JSON(w, http.StatusOK, recordPresenter)
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
