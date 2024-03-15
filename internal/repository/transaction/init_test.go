package transaction

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/soerjadi/brick/internal/config"
)

type prepareQueryMock struct {
	disbursement    *sqlmock.ExpectedPrepare
	updateTrxStatus *sqlmock.ExpectedPrepare
	getAccount      *sqlmock.ExpectedPrepare
	insertAccount   *sqlmock.ExpectedPrepare
	updateTrxID     *sqlmock.ExpectedPrepare
}

func expectPrepareMock(mock sqlmock.Sqlmock) prepareQueryMock {
	prepareQueryMock := prepareQueryMock{}

	prepareQueryMock.disbursement = mock.ExpectPrepare(`
	INSERT INTO transaction \(
		destination_id,
		amount,
		unique_number,
		type,
		status,
		created_at
	\) VALUES \(
		(.*),
		(.*),
		(.*),
		(.*),
		(.*),
		NOW\(\)
	\) RETURNING 
		id,
		destination_id,
		amount,
		unique_number,
		type,
		status,
		created_at
	`)

	prepareQueryMock.updateTrxStatus = mock.ExpectPrepare(`
	UPDATE transaction 
	SET
		status=(.*)
	WHERE
		trx_id=(.*)
	`)

	prepareQueryMock.getAccount = mock.ExpectPrepare(`
	SELECT
		id,
		name,
		number,
		origin_bank
	FROM
		account
	WHERE
		number = (.*) AND origin_bank = (.*)
	`)

	prepareQueryMock.insertAccount = mock.ExpectPrepare(`
	INSERT INTO account \(
		id,
		name,
		number,
		origin_bank
	\) VALUES \(
		(.*),
		(.*),
		(.*),
		(.*)
	\) RETURNING
		id,
		name,
		number,
		origin_bank
	`)

	prepareQueryMock.updateTrxID = mock.ExpectPrepare(`
	UPDATE transaction
	SET
		trx_id=(.*)
	WHERE
		id=(.*)
	`)
	return prepareQueryMock
}

func TestGetRepository(t *testing.T) {
	tests := []struct {
		name     string
		initMock func() (*sqlx.DB, *sql.DB, sqlmock.Sqlmock)
		want     func(db *sqlx.DB, cfg *config.Config) Repository
		wantErr  bool
	}{
		{
			name: "success",
			initMock: func() (*sqlx.DB, *sql.DB, sqlmock.Sqlmock) {
				db, mock, _ := sqlmock.New()
				expectPrepareMock(mock)
				expectPrepareMock(mock)
				return sqlx.NewDb(db, "postgres"), db, mock
			},
			want: func(db *sqlx.DB, cfg *config.Config) Repository {
				q, _ := prepareQueries(db)

				return &trxRepository{
					query:  q,
					DB:     db,
					config: cfg,
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, dbMock, mock := tt.initMock()
			defer dbMock.Close()

			cfg := config.Config{}
			got, err := GetRepository(db, &cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRepository() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			want := tt.want(db, &cfg)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("GetRepository() = %v, want %v", got, want)
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err.Error())
			}
		})
	}

}
