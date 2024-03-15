package transaction

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/soerjadi/brick/internal/delivery/rest"
	"github.com/soerjadi/brick/internal/usecase/transaction"
)

func NewHandler(usecase transaction.Usecase) rest.API {
	return &Handler{
		usecase: usecase,
	}
}

func (h *Handler) RegisterRoutes(r *mux.Router) {
	transaction := r.PathPrefix("/transaction").Subrouter()

	transaction.HandleFunc("/account/check", rest.HandlerFunc(h.CheckAccount).Serve).Methods(http.MethodPost)
	transaction.HandleFunc("/disburse", rest.HandlerFunc(h.Transaction).Serve).Methods(http.MethodPost)
	transaction.HandleFunc("/callback", rest.HandlerFunc(h.Callback).Serve).Methods(http.MethodPost)
}
