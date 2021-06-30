package handler

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/erbilsilik/getir-go-challange/api/presenter"
	"github.com/erbilsilik/getir-go-challange/entity"
	configurationMock "github.com/erbilsilik/getir-go-challange/usecase/configuration/mock"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateConfig(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := configurationMock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeConfigurationHandlers(r, *n, service)
	path, err := r.GetRoute("createConfig").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/configurations", path)

	c := &entity.Config{
		Key:    "active-tabs",
		Value: 	"getir",
	}

	service.EXPECT().
		Create(c.Key, c.Value).
		Return(nil)
	h := createConfig(service)

	ts := httptest.NewServer(h)
	defer ts.Close()
	payload := fmt.Sprintf(`{
		"key": "active-tabs",
		"value": "getir"
      }`,
	)
	resp, _ := http.Post(ts.URL + "/v1/configurations", "application/json", strings.NewReader(payload))
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}

func TestGetConfig(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := configurationMock.NewMockUseCase(controller)
	r := mux.NewRouter()
	n := negroni.New()
	MakeConfigurationHandlers(r, *n, service)
	path, err := r.GetRoute("getConfig").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/v1/configurations", path)
	key := "active-tabs"
	c := &entity.Config{
		Key:    key,
		Value: 	"getir",
	}
	service.EXPECT().
		FindByKey(key).
		Return(c)
	ts := httptest.NewServer(getConfig(service))
	defer ts.Close()
	res, err := http.Get(ts.URL + "?key=" + key)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	var response presenter.Response
	json.NewDecoder(res.Body).Decode(&response)
	resp := response.Data.(map[string]interface{})
	assert.Equal(t, resp["Key"], c.Key)
	assert.Equal(t, resp["Value"], c.Value)
}