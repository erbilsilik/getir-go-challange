package presenter

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type response struct {
	Code        	 	 int `json:"code"`
	Msg string      	`json:"msg"`
	Data interface{}    `json:"data"`
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	var msg string
	if statusCode == http.StatusOK || statusCode == http.StatusCreated {
		msg = "Success"
	} else {
		msg = "Failed"
	}

	response := response{
		Code: statusCode,
		Msg: msg,
		Data: data,
	}

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JSON(w, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	JSON(w, http.StatusBadRequest, nil)
}
