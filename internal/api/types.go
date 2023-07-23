package api

import "Codenotary/internal/model"

type VacRecordResponse struct {
	Records []model.VacRecordModel `json:"records"`
}
