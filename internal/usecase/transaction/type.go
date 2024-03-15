package transaction

import (
	"context"

	"github.com/soerjadi/brick/internal/model"
	"github.com/soerjadi/brick/internal/repository/transaction"
)

//go:generate mockgen -package=mocks -mock_names=Usecase=MockTransactionUsecase -destination=../../mocks/transaction_usecase_mock.go -source=type.go
type Usecase interface {
	CheckAccount(ctx context.Context, req model.CheckAccountRequest) (*model.Account, error)
	Disbursement(ctx context.Context, req model.DisbursementRequest) (trxID string, err error)
	Callback(ctx context.Context, req model.CallbackRequest) error
}

type trxUsecase struct {
	repository transaction.Repository
}
