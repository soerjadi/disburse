package transaction

import (
	"context"
	"fmt"
	"net/http"

	"github.com/soerjadi/brick/internal/model"
	"github.com/soerjadi/brick/internal/pkg/log"
)

type BankAccountResponse struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	AccountNumber string `json:"account_number"`
}

type TrxResponse struct {
	TransactionID string `json:"transaction_id"`
}

func (t *trxRepository) CheckAccount(ctx context.Context, req model.CheckAccountRequest) (*model.Account, error) {
	checkAccountURL := fmt.Sprintf("%s/v8/vbank/search/account", t.config.Bank.URL)

	t.config.HttpClient.SetHeaderFn(func(req *http.Request) error {
		req.Header.Set("Authorization", "Bearer token")
		return nil
	})

	checkReq := t.config.HttpClient.Request(http.MethodPost, checkAccountURL, map[string]any{
		"account_number": req.AccountNumber,
	})
	var resp BankAccountResponse

	if err := checkReq.Decode(&resp); err != nil {
		return nil, err
	}

	return &model.Account{
		ID:         resp.ID,
		Name:       resp.Name,
		Number:     resp.AccountNumber,
		OriginBank: req.BankName,
	}, nil
}

func (t *trxRepository) Disbursement(ctx context.Context, req model.DisbursementRequest) (trxId string, err error) {
	checkAccountURL := fmt.Sprintf("%s/v8/vbank/transfers", t.config.Bank.URL)

	t.config.HttpClient.SetHeaderFn(func(req *http.Request) error {
		req.Header.Set("Authorization", "Bearer token")
		return nil
	})

	checkReq := t.config.HttpClient.Request(http.MethodPost, checkAccountURL, map[string]any{
		"account": req.Account,
		"amount":  req.Amount,
	})
	var resp TrxResponse

	if err = checkReq.Decode(&resp); err != nil {
		log.ErrorWithFields("repository.transaction.Disbursement failed disburse", log.KV{
			"err":     err,
			"request": req,
		})
		return
	}

	trxId = resp.TransactionID

	return
}

func (t *trxRepository) GetAccountByNumberOrigin(ctx context.Context, number, originBank string) (model.Account, error) {
	var res model.Account

	err := t.query.getAccount.GetContext(ctx, &res, number, originBank)
	if err != nil {
		log.ErrorWithFields("repository.transaction.GetAccountByNumberOrigin failed get account by number and origin", log.KV{
			"err":    err,
			"number": number,
			"origin": originBank,
		})
		return model.Account{}, err
	}

	return res, nil
}

func (t *trxRepository) InsertAccount(ctx context.Context, req model.Account) (model.Account, error) {
	var (
		err error
		res model.Account
	)

	if err = t.query.insertAccount.GetContext(
		ctx,
		&res,
		req.ID,
		req.Name,
		req.Number,
		req.OriginBank,
	); err != nil {
		log.ErrorWithFields("repository.transaction.InsertAccount failed insert account.", log.KV{
			"err":     err,
			"request": req,
		})
		return model.Account{}, err
	}

	return res, nil
}

func (t *trxRepository) Callback(ctx context.Context, req model.CallbackRequest) error {
	var err error

	if _, err = t.query.updateTrxStatus.ExecContext(
		ctx,
		req.Status,
		req.TransactionID,
	); err != nil {
		log.ErrorWithFields("repository.transaction.Callback failed update transaction", log.KV{
			"err":     err,
			"request": req,
		})
		return err
	}

	return nil
}

func (t *trxRepository) InsertTrx(ctx context.Context, req model.Transaction) (model.Transaction, error) {
	var (
		err error
		res model.Transaction
	)

	if err = t.query.disbursement.GetContext(
		ctx,
		&res,
		req.Destination.ID,
		req.Amount,
		req.UniqueNumber,
		req.Type,
		req.Status,
	); err != nil {
		log.ErrorWithFields("repository.transaction.InsertTrx failed insert transaction", log.KV{
			"err":     err,
			"request": req,
		})
		return model.Transaction{}, err
	}

	return res, nil
}

func (t *trxRepository) UpdateTrxID(ctx context.Context, id int64, trxID string) error {
	var err error

	if _, err = t.query.updateTrxID.ExecContext(
		ctx,
		trxID,
		id,
	); err != nil {
		log.ErrorWithFields("repository.transation.updateTrxID failed update trxID", log.KV{
			"err":   err,
			"id":    id,
			"trxID": trxID,
		})

		return err
	}

	return nil
}
