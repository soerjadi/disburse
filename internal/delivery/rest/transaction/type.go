package transaction

import "github.com/soerjadi/brick/internal/usecase/transaction"

type Handler struct {
	usecase transaction.Usecase
}
