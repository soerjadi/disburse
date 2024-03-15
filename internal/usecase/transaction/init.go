package transaction

import (
	"github.com/soerjadi/brick/internal/repository/transaction"
)

func GetUsecase(repo transaction.Repository) Usecase {
	return &trxUsecase{
		repository: repo,
	}
}
