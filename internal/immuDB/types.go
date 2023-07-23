package immuDB

import (
	"Codenotary/internal/model"
	"context"
)

type ImmuDBRepository interface {
	GetAllRecords(ctx context.Context) ([]model.VacRecordModel, error)
	CreateRecord(ctx context.Context, record model.CreateRecordModel) (*model.RecordCreateResponse, error)
	SearchRecord(ctx context.Context, docID string) ([]model.VacRecordModel, error)
}
