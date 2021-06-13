package middleware

import (
	"net/http"
	"regexp"
	"strconv"
)

func ValidateGetRecordsFilteredByTimeAndTotalCountInGivenNumberRange(
	rw http.ResponseWriter,
	r *http.Request,
	next http.HandlerFunc,
) {
	q := r.URL.Query()

	var validCount = regexp.MustCompile(`^\d+$`)
	var validDateString = regexp.MustCompile(`^\d{4}\-(0[1-9]|1[012])\-(0[1-9]|[12][0-9]|3[01])$`)

	minCount := q.Get("minCount")
	maxCount := q.Get("maxCount")
	startDate := q.Get("startDate")
	endDate := q.Get("endDate")

	minCountConverted, _ := strconv.Atoi(minCount)
	maxCountConverted, _ := strconv.Atoi(maxCount)

	if !validCount.MatchString(minCount) {
		panic("invalid minCount value")
	} else if !validCount.MatchString(maxCount) {
		panic("invalid maxCount value")
	} else if minCountConverted > maxCountConverted {
		panic("minCount should be less than maxCount")
	} else if !validDateString.MatchString(startDate) {
		panic("invalid startDate value")
	} else if !validDateString.MatchString(endDate) {
		panic("invalid endDate value")
	}
	next(rw, r)
}
