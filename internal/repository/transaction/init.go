package transaction

import (
	"github.com/jmoiron/sqlx"
	"github.com/soerjadi/brick/internal/config"
)

func prepareQueries(db *sqlx.DB) (prepareQuery, error) {
	var (
		err error
		q   prepareQuery
	)

	q.disbursement, err = db.Preparex(disbursement)
	if err != nil {
		return q, err
	}

	q.updateTrxStatus, err = db.Preparex(updateTrxStatus)
	if err != nil {
		return q, err
	}

	q.getAccount, err = db.Preparex(getAccount)
	if err != nil {
		return q, err
	}

	q.insertAccount, err = db.Preparex(insertAccount)
	if err != nil {
		return q, err
	}

	q.updateTrxID, err = db.Preparex(updateTrxID)
	if err != nil {
		return q, err
	}

	return q, nil

}

func GetRepository(db *sqlx.DB, cfg *config.Config) (Repository, error) {
	query, err := prepareQueries(db)
	if err != nil {
		return nil, err
	}

	return &trxRepository{
		query:  query,
		DB:     db,
		config: cfg,
	}, nil
}
