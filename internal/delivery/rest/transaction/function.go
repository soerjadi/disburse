package transaction

import (
	"encoding/json"
	"net/http"

	"github.com/soerjadi/brick/internal/model"
	"github.com/soerjadi/brick/internal/pkg/log"
)

func (h *Handler) CheckAccount(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	dec := json.NewDecoder(r.Body)
	req := model.CheckAccountRequest{}

	if err := dec.Decode(&req); err != nil {
		log.ErrorWithFields("handler.transaction.CheckAccount fail to decode request body", log.KV{
			"err": err,
		})
		return nil, err
	}

	account, err := h.usecase.CheckAccount(r.Context(), req)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (h *Handler) Transaction(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	dec := json.NewDecoder(r.Body)
	req := model.DisbursementRequest{}

	if err := dec.Decode(&req); err != nil {
		log.ErrorWithFields("handler.transaction.Transaction fail to decode request body", log.KV{
			"err": err,
		})
		return nil, err
	}

	trxID, err := h.usecase.Disbursement(r.Context(), req)
	if err != nil {
		return "", err
	}

	return map[string]interface{}{
		"transaction_id": trxID,
	}, nil
}

func (h *Handler) Callback(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	dec := json.NewDecoder(r.Body)
	req := model.CallbackRequest{}

	if err := dec.Decode(&req); err != nil {
		log.ErrorWithFields("handler.transaction.Callback fail to decode request body", log.KV{
			"err": err,
		})

		return nil, err
	}
	err := h.usecase.Callback(r.Context(), req)
	if err != nil {
		log.ErrorWithFields("handler.transaction.Callback fail to send callback", log.KV{
			"err": err,
		})
		return nil, err
	}

	return "", nil
}
