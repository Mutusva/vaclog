package immuDB

import (
	"Codenotary/config"
	"Codenotary/internal/model"
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
)

var (
	RecordTemplate  = "/default/document"
	RecordsTemplate = "/default/documents"
	contentType     = "application/json"
)

type immudbClient struct {
	Config config.AppConfig
	Client *http.Client
}

func (r *immudbClient) GetAllRecords(ctx context.Context) ([]model.VacRecordModel, error) {
	immudbUrl := r.Config.ImmuDbBaseUrl + RecordsTemplate + "/search"
	data := []byte(`{"page":1,"perPage":100}`)

	var docs model.SearchRecordResponse
	var vacRecords []model.VacRecordModel

	headers := getImmuDBHeaders(r.Config.ImmuDbAPIKey)
	queryParameters := url.Values{}
	results, err := MakeHTTPRequest(ctx, r.Client, immudbUrl, http.MethodPost, headers, queryParameters, bytes.NewBuffer(data), docs)
	if err != nil {
		return vacRecords, err
	}

	return ConvertDocToVacModel(results.Revisions, vacRecords)
}

func ConvertDocToVacModel(docs []model.VacDocument, vacRecords []model.VacRecordModel) ([]model.VacRecordModel, error) {
	for _, doc := range docs {
		vacRecords = append(vacRecords, model.VacRecordModel{
			DocumentId:       doc.Document.ID,
			AnimalID:         doc.Document.AnimalId,
			Species:          doc.Document.Species,
			Age:              doc.Document.Age,
			VaccineName:      doc.Document.VaccineName,
			DateAdministered: doc.Document.DateAdministered,
			Notes:            doc.Document.Notes,
		})
	}

	return vacRecords, nil
}

func (r *immudbClient) CreateRecord(ctx context.Context, record model.CreateRecordModel) (*model.RecordCreateResponse, error) {
	immudbUrl := r.Config.ImmuDbBaseUrl + RecordTemplate

	var response model.RecordCreateResponse
	body, err := json.Marshal(record)
	if err != nil {
		return &response, err
	}

	headers := getImmuDBHeaders(r.Config.ImmuDbAPIKey)
	queryParameters := url.Values{}
	result, err := MakeHTTPRequest(ctx, r.Client, immudbUrl, http.MethodPut, headers, queryParameters, bytes.NewBuffer(body), response)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *immudbClient) SearchRecord(ctx context.Context, docID string) ([]model.VacRecordModel, error) {
	immudbUrl := r.Config.ImmuDbBaseUrl + RecordsTemplate + "/search"

	var docs model.SearchRecordResponse
	query := model.SearchQuery{
		Query: model.Query{
			Expressions: []model.Expression{
				{
					FieldComparisons: []model.FieldComparison{
						{
							Field:    "_id",
							Operator: "EQ",
							Value:    docID,
						},
					},
				},
			},
		},
		Page:    1,
		PerPage: 10,
	}

	var vacRecords []model.VacRecordModel
	body, err := json.Marshal(query)
	if err != nil {
		return vacRecords, err
	}

	headers := getImmuDBHeaders(r.Config.ImmuDbAPIKey)
	queryParameters := url.Values{}
	results, err := MakeHTTPRequest(ctx, r.Client, immudbUrl, http.MethodPost, headers, queryParameters, bytes.NewBuffer(body), docs)

	if err != nil {
		return vacRecords, err
	}

	return ConvertDocToVacModel(results.Revisions, vacRecords)
}

func getImmuDBHeaders(APIKey string) map[string]string {
	return map[string]string{
		"Content-Type": contentType,
		"accept":       contentType,
		"X-API-Key":    APIKey,
	}
}

func New(config config.AppConfig, client *http.Client) ImmuDBRepository {
	return &immudbClient{Config: config, Client: client}
}
