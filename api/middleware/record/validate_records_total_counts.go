package record

import (
	"encoding/json"
	"github.com/erbilsilik/getir-go-challange/api/presenter"
	"github.com/erbilsilik/getir-go-challange/pkg/utilities"
	"net/http"
	"regexp"
	"strconv"
)

func writeErrorMessage(rw http.ResponseWriter, message string) {
	response := presenter.Response{
		Code: http.StatusBadRequest,
		Msg: message,
	}
	rw.WriteHeader(http.StatusBadRequest)
	_ = json.NewEncoder(rw).Encode(response)
}

func ValidateGetRecordsFilteredByTimeAndTotalCountInGivenNumberRange(
	rw http.ResponseWriter,
	r *http.Request,
	next http.HandlerFunc,
) {
	q := r.URL.Query()

	var validCount = regexp.MustCompile(`^\d+$`)
	var validDateString = regexp.MustCompile(`^\d{4}\-(0[1-9]|1[012])\-(0[1-9]|[12][0-9]|3[01])$`)

	minCount := q.Get("minCount")
	if !validCount.MatchString(minCount) {
		writeErrorMessage(rw, "invalid minCount value")
		return
	}

	maxCount := q.Get("maxCount")
	if !validCount.MatchString(maxCount) {
		writeErrorMessage(rw,"invalid maxCount value")
		return
	}

	minCountConverted, _ := strconv.Atoi(minCount)
	maxCountConverted, _ := strconv.Atoi(maxCount)

	if minCountConverted > maxCountConverted {
		writeErrorMessage(rw, "minCount should be less than maxCount")
		return
	}


	startDate := q.Get("startDate")
	if !validDateString.MatchString(startDate) {
		writeErrorMessage(rw,"invalid startDate value")
		return
	}

	endDate := q.Get("endDate")
	if !validDateString.MatchString(endDate) {
		writeErrorMessage(rw, "invalid endDate value")
		return
	}

	layout := "2006-01-02"
	startDateParsed := utilities.ParseDate(layout, startDate)
	endDateParsed := utilities.ParseDate(layout, endDate)
	if !endDateParsed.After(startDateParsed) {
		writeErrorMessage(rw, "the start date must be before the end date")
		return
	}

	next(rw, r)
}
