package transaction

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/soerjadi/brick/internal/config"
	"github.com/soerjadi/brick/internal/model"
)

//go:generate `mockgen -package=mocks -mock_names=Repository=MockTransactionRepository -destination=../../mocks/transaction_repo_mock.go -source=type.go`
type Repository interface {
	CheckAccount(ctx context.Context, req model.CheckAccountRequest) (*model.Account, error)
	Disbursement(ctx context.Context, req model.DisbursementRequest) (trxId string, err error)
	GetAccountByNumberOrigin(ctx context.Context, Number, OriginBank string) (model.Account, error)
	InsertAccount(ctx context.Context, req model.Account) (model.Account, error)
	InsertTrx(ctx context.Context, req model.Transaction) (model.Transaction, error)
	Callback(ctx context.Context, req model.CallbackRequest) error
	UpdateTrxID(ctx context.Context, id int64, trxID string) error
}

type trxRepository struct {
	query  prepareQuery
	DB     *sqlx.DB
	config *config.Config
}
