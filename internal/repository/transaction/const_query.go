package transaction

import "github.com/jmoiron/sqlx"

type prepareQuery struct {
	disbursement    *sqlx.Stmt
	updateTrxStatus *sqlx.Stmt
	getAccount      *sqlx.Stmt
	insertAccount   *sqlx.Stmt
	updateTrxID     *sqlx.Stmt
}

const (
	disbursement = `
	INSERT INTO transaction (
		destination_id,
		amount,
		unique_number,
		type,
		status,
		created_at
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		NOW()
	) RETURNING 
		id,
		destination_id,
		amount,
		unique_number,
		type,
		status,
		created_at
	`
	getAccount = `
	SELECT
		id,
		name,
		number,
		origin_bank
	FROM
		account
	WHERE
		number = $1 AND origin_bank = $2
	`

	insertAccount = `
	INSERT INTO account (
		id,
		name,
		number,
		origin_bank
	) VALUES (
		$1,
		$2,
		$3,
		$4
	) RETURNING
		id,
		name,
		number,
		origin_bank
	`

	updateTrxStatus = `
	UPDATE transaction 
	SET
		status=$1
	WHERE
		trx_id=$2
	`

	updateTrxID = `
	UPDATE transaction
	SET
		trx_id=$1
	WHERE
		id=$2
	`
)
