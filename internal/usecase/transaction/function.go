package transaction

import (
	"context"
	"database/sql"

	"github.com/soerjadi/brick/internal/model"
	"github.com/soerjadi/brick/internal/pkg/log"
)

func (t *trxUsecase) CheckAccount(ctx context.Context, req model.CheckAccountRequest) (*model.Account, error) {

	res, err := t.repository.CheckAccount(ctx, req)
	if err != nil {
		log.ErrorWithFields("usecase.transaction.CheckAccount failed check account", log.KV{
			"err":     err,
			"request": req,
		})
		return nil, err
	}

	return res, nil
}

func (t *trxUsecase) Disbursement(ctx context.Context, req model.DisbursementRequest) (trxID string, err error) {

	account, err := t.repository.GetAccountByNumberOrigin(ctx, req.Account.Number, req.Account.OriginBank)
	if err != nil && err != sql.ErrNoRows {
		log.ErrorWithFields("usecase.transaction.Disbursement failed get account by number and origin", log.KV{
			"err":     err,
			"request": req,
		})
		return
	}

	if err == sql.ErrNoRows {
		account, err = t.repository.InsertAccount(ctx, model.Account{
			ID:         req.Account.ID,
			Name:       req.Account.Name,
			Number:     req.Account.Number,
			OriginBank: req.Account.OriginBank,
		})

		if err != nil {
			log.ErrorWithFields("usecase.transaction.Disbursement failed insert account", log.KV{
				"err":     err,
				"request": req,
			})

			return
		}
	}

	req.Account = account
	trx, err := t.repository.InsertTrx(ctx, model.Transaction{
		Destination:  req.Account,
		Amount:       req.Amount,
		UniqueNumber: req.UniqueNumber,
		Type:         model.Disbursement,
		Status:       model.StatusPending,
	})
	if err != nil {
		log.ErrorWithFields("usecase.transaction.Disbursement failed insert transaction", log.KV{
			"err":     err,
			"request": req,
		})
		return
	}

	trxID, err = t.repository.Disbursement(ctx, model.DisbursementRequest{
		Account:      req.Account,
		Amount:       req.Amount,
		UniqueNumber: req.UniqueNumber,
	})
	if err != nil {
		log.ErrorWithFields("usecase.transaction.Disbursement failed send disbursement", log.KV{
			"err":     err,
			"request": req,
		})
		return
	}

	err = t.repository.UpdateTrxID(ctx, trx.ID, trxID)
	if err != nil {
		log.ErrorWithFields("usecase.transaction.Disbursement failed update transaction ID", log.KV{
			"err": err,
		})
		return
	}

	return
}

func (t *trxUsecase) Callback(ctx context.Context, req model.CallbackRequest) error {

	err := t.repository.Callback(ctx, req)
	if err != nil {
		log.ErrorWithFields("usecase.transaction.Callbak failed get callback", log.KV{
			"err":     err,
			"request": req,
		})
		return err
	}

	return nil
}
