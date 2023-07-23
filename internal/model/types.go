package model

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

type VacRecordModel struct {
	DocumentId       string    `json:"document_id"`
	AnimalID         string    `json:"animal_id"`
	Species          string    `json:"species"`
	Age              int       `json:"age"`
	VaccineName      string    `json:"vaccine_name"`
	DateAdministered time.Time `json:"date_administered"`
	Notes            string    `json:"notes"`
}

type CreateRecordModel struct {
	DocumentId       string    `json:"document_id"`
	AnimalID         string    `json:"animal_id"`
	Species          string    `json:"species"`
	Age              StringInt `json:"age"`
	VaccineName      string    `json:"vaccine_name"`
	DateAdministered JSONDate  `json:"date_administered"`
	Notes            string    `json:"notes"`
}

type VaultMd struct {
	Creator string `json:"creator"`
	Ts      int    `json:"ts"`
}

type Document struct {
	ID               string    `json:"_id"`
	VaultMd          VaultMd   `json:"_vault_md"`
	Age              int       `json:"age"`
	AnimalId         string    `json:"animal_id"`
	DateAdministered time.Time `json:"date_administered"`
	Notes            string    `json:"notes"`
	Species          string    `json:"species"`
	VaccineName      string    `json:"vaccine_name"`
}

type SearchRecordResponse struct {
	Page      int           `json:"page"`
	PerPage   int           `json:"perPage"`
	Revisions []VacDocument `json:"revisions"`
	SearchID  string        `json:"searchId"`
}

type VacDocument struct {
	Document      Document `json:"document"`
	Revision      string   `json:"revision"`
	TransactionId string   `json:"transactionId"`
}

type RecordCreateResponse struct {
	TransactionId string `json:"transactionId"`
	DocumentId    string `json:"documentId"`
}

type ErrorResponse struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type OrderProperty struct {
	Field string `json:"field"`
	Desc  bool   `json:"desc"`
}

type Revision struct {
	TransactionID string   `json:"transactionId"`
	Revision      string   `json:"revision"`
	Document      struct{} `json:"document"`
}

type SearchResponse struct {
	SearchID  string     `json:"searchId"`
	Revisions []Revision `json:"revisions"`
	Page      int        `json:"page"`
	PerPage   int        `json:"perPage"`
}

//type ImmuDBSearchModel struct {
//	SearchID string `json:"searchId"`
//	KeepOpen bool   `json:"keepOpen"`
//	Query    RQuery `json:"query"`
//	Page     int    `json:"page"`
//	PerPage  int    `json:"perPage"`
//}

type SearchQuery struct {
	Query   Query `json:"query"`
	Page    int   `json:"page"`
	PerPage int   `json:"perPage"`
}

type Query struct {
	Expressions []Expression `json:"expressions"`
}

type Expression struct {
	FieldComparisons []FieldComparison `json:"fieldComparisons"`
}

type FieldComparison struct {
	Field    string `json:"field"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

type SearchModel struct {
	DocumentID string `json:"document_id"`
}

type JSONDate time.Time

func (j *JSONDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*j = JSONDate(t)
	return nil
}

func (j JSONDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(j))
}

// Maybe a Format function for printing your date
func (j JSONDate) Format(s string) string {
	t := time.Time(j)
	return t.Format(s)
}

// StringInt create a type alias for type int
type StringInt int

// UnmarshalJSON create a custom unmarshal for the StringInt
/// this helps us check the type of our value before unmarshalling it

func (st *StringInt) UnmarshalJSON(b []byte) error {
	//convert the bytes into an interface
	//this will help us check the type of our value
	//if it is a string that can be converted into an int we convert it
	///otherwise we return an error
	var item interface{}
	if err := json.Unmarshal(b, &item); err != nil {
		return err
	}
	switch v := item.(type) {
	case int:
		*st = StringInt(v)
	case float64:
		*st = StringInt(int(v))
	case string:
		///here convert the string into
		///an integer
		i, err := strconv.Atoi(v)
		if err != nil {
			///the string might not be of integer type
			///so return an error
			return err

		}
		*st = StringInt(i)

	}
	return nil
}
