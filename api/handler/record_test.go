package handler

import (
	"encoding/json"
	"github.com/codegangsta/negroni"
	"github.com/erbilsilik/getir-go-challange/api/presenter"
	"github.com/erbilsilik/getir-go-challange/entity"
	"github.com/erbilsilik/getir-go-challange/pkg/utilities"
	recordUsecase "github.com/erbilsilik/getir-go-challange/usecase/record"
	recordMock "github.com/erbilsilik/getir-go-challange/usecase/record/mock"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetRecordsFilteredByTimeAndTotalCountInGivenNumberRange(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := recordMock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeRecordHandlers(r, *n, service)
	path, err := r.GetRoute("getRecordsFilteredByTimeAndTotalCountInGivenNumberRange").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/records", path)

	layout := "2006-01-02"
	q := recordUsecase.RecordsFilteredByTimeAndTotalCountInGivenNumberRangeQuery{
		StartDate: utilities.ParseDate(layout, "2016-01-26"),
		EndDate: utilities.ParseDate(layout, "2018-02-02"),
		MinCount: 2700,
		MaxCount: 3000,
	}

	RFC3339 := "2006-01-02T15:04:05Z07:00"
	recordResponse := &entity.Record{
		Key:        "ibfRLaFT",
		TotalCount: 2892,
		CreatedAt: utilities.ParseDate(RFC3339, "2016-12-25T16:43:27.909Z"),
	}

	service.EXPECT().
		GetRecordsFilteredByTimeAndTotalCountInGivenNumberRange(&q).
		Return([]*entity.Record{recordResponse}, nil)
	ts := httptest.NewServer(getRecordsFilteredByTimeAndTotalCountInGivenNumberRange(service))
	defer ts.Close()
	res, err := http.Get(ts.URL + "?startDate=2016-01-26&endDate=2018-02-02&minCount=2700&maxCount=3000")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	var response presenter.Response

	json.NewDecoder(res.Body).Decode(&response)

	resp := response.Data.(map[string]interface{})

	for _, item:= range resp["records"].([]interface{}) {
		assert.Equal(t, item.(map[string]interface{})["key"], recordResponse.Key)
	}
}