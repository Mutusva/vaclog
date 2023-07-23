package immuDB

import (
	"Codenotary/config"
	"Codenotary/internal/model"
	"context"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetAllRecords_success(t *testing.T) {

	payload := `
     {
	  "page": 1,
	  "perPage": 100,
	  "revisions": [
		{
		  "document": {
			"_id": "64b310630000000000000004db10e4e8",
			"_vault_md": {
			  "creator": "a:651f0d64-5241-47fe-8915-2a2046d16020",
			  "ts": 1689456739
			},
			"age": 8,
			"animal_id": "cattle123",
			"date_administered": "2023-01-02T15:04:05Z",
			"notes": "Prevention of skin disease",
			"species": "cattle",
			"vaccine_name": "werderd"
		  },
		  "revision": "",
		  "transactionId": ""
		},
		{
		  "document": {
			"_id": "64b315640000000000000005db10e4e9",
			"_vault_md": {
			  "creator": "a:651f0d64-5241-47fe-8915-2a2046d16020",
			  "ts": 1689458020
			},
			"age": 5,
			"animal_id": "sheep1234",
			"date_administered": "2023-01-02T15:04:05Z",
			"notes": "used early in the life of a sheep",
			"species": "sheep",
			"vaccine_name": "shepokj"
		  },
		  "revision": "",
		  "transactionId": ""
		}
	  ],
	  "searchId": ""
	}`

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(payload))
		assert.Nil(t, err)
	}))

	config := config.AppConfig{
		ImmuDbBaseUrl: testServer.URL,
		ImmuDbAPIKey:  "Test_key",
	}

	client := New(config, &http.Client{})

	records, err := client.GetAllRecords(context.Background())

	assert.Nil(t, err)
	assert.Equal(t, 2, len(records))
	assert.Equal(t, records[0].AnimalID, "cattle123")

}

func TestGetAllRecords_NoReturnRecord(t *testing.T) {
	// Arrange
	payload := `{
      "page": 1,
	  "perPage": 100,
	  "revisions": []
      }
     `

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(payload))
		assert.Nil(t, err)
	}))

	config := config.AppConfig{
		ImmuDbBaseUrl: testServer.URL,
		ImmuDbAPIKey:  "Test_key",
	}

	client := New(config, &http.Client{})

	// Act
	response, err := client.GetAllRecords(context.Background())

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, 0, len(response))

}

func TestGetAllRecords_ErrorResponse(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))

	config := config.AppConfig{
		ImmuDbBaseUrl: testServer.URL,
		ImmuDbAPIKey:  "Test_key",
	}

	client := New(config, &http.Client{})

	// Act
	_, err := client.GetAllRecords(context.Background())

	// Assert
	assert.Error(t, err)

}

func TestCreateRecord_Success(t *testing.T) {
	// Arrange
	payload := `
         {
           "transactionId": "abcg12334",
           "documentId": "123445"
         } 
      `

	record := model.CreateRecordModel{
		AnimalID:         "animal123",
		Species:          "Cat",
		Age:              12,
		VaccineName:      "test_vaccine",
		DateAdministered: model.JSONDate(time.Now()),
		Notes:            "",
	}

	exp := &model.RecordCreateResponse{
		TransactionId: "abcg12334",
		DocumentId:    "123445",
	}

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(payload))
		assert.Nil(t, err)
	}))

	config := config.AppConfig{
		ImmuDbBaseUrl: testServer.URL,
		ImmuDbAPIKey:  "Test_key",
	}

	client := New(config, &http.Client{})

	// Act
	response, err := client.CreateRecord(context.Background(), record)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, exp, response)
}

func TestCreateRecord_ErrorResponse(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))

	config := config.AppConfig{
		ImmuDbBaseUrl: testServer.URL,
		ImmuDbAPIKey:  "Test_key",
	}

	client := New(config, &http.Client{})

	// Act
	_, err := client.CreateRecord(context.Background(), model.CreateRecordModel{})

	// Assert
	assert.Error(t, err)

}
