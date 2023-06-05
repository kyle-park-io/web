package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type EmailRequestBody struct {
	Email string
}

type ExampleRequestBody struct {
	Function string     `json:"function"`
	Data     DataStruct `json:"data"`
}

type DataStruct struct {
	User      string `json:"user"`
	Holder    string `json:"holder"`
	Owner     string `json:"owner"`
	Spender   string `json:"spender"`
	From      string `json:"from"`
	To        string `json:"to"`
	Recipient string `json:"recipient"`
	Drop      string `json:"drop"`

	Partition string `json:"partition"`

	Amount string `json:"amount"`

	Bookmark string `json:"bookmark"`
	PageSize string `json:"pageSize"`

	Recipients map[string]string `json:"recipients"`
}

func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestPost(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()

	// test mock
	objectToSerializeToJSON := EmailRequestBody{"kyle.park@themedium.io"}

	var b bytes.Buffer
	err := json.NewEncoder(&b).Encode(objectToSerializeToJSON)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(http.MethodPost, "/examplePost", &b)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "kyle.park@themedium.io", w.Body.String())
}

func TestPostObject(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()

	// test mock
	// objectToSerializeToJSON := ExampleRequestBody{"MintByPartition", DataStruct{User: "Org1User1", Partition: "mediumToken", Amount: "-100"}}

	ex := map[string]string{
		"0x47a7a67edf2e0f1e89d1ab7b547dc67d0ce334df": "50",
		"0x7eddc225c347da6b844b87baeecdfd7be35eb1c0": "30",
		"0x1396c5d0dc6f26dee34ec3a6b33325c22838d38a": "120",
	}

	objectToSerializeToJSON := ExampleRequestBody{"DistributeToken", DataStruct{Recipients: ex, Partition: "mediumToken"}}

	var b bytes.Buffer
	err := json.NewEncoder(&b).Encode(objectToSerializeToJSON)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(http.MethodPost, "/testChaincode", &b)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "checkData", w.Body.String())
}
