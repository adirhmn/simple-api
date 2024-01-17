package dbpool

import (
	"context"
	"database/sql"
	"github.com/sirupsen/logrus"
)

func Transaction(ctx context.Context, db *sql.DB, exec func(trx *sql.Tx) error) error {

	tx, errTx := db.BeginTx(ctx, nil)
	if errTx != nil {
		// send error log
		return errTx
	}

	if err := exec(tx); err != nil {
		if errRb := tx.Rollback(); errRb != nil {
			logrus.Error("Failed rollback DB")
		}
		return err
	}

	if errCommit := tx.Commit(); errCommit != nil {
		// send error log
		return errCommit
	}

	return nil
}
