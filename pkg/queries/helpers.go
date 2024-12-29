package queries

import "database/sql"

type CombineResult struct {
	Transaction *sql.Tx
	Commit      func() error
	Close       func(rollback bool) error
}

func combine[T TransactionCaller[Prepared]](transaction *sql.Tx, managers ...T) CombineResult {
	for _, manager := range managers {
		manager.SetTransaction(transaction)
	}

	commit := func() error { return transaction.Commit() }
	close := func(rollback bool) error {
		var err error
		if rollback {
			err = transaction.Rollback()
		}

		for _, manager := range managers {
			manager.Close()
		}

		return err
	}

	return CombineResult{transaction, commit, close}
}

func Combine[T Manager[Prepared]](opener T, managers ...T) (result CombineResult, err error) {
	transaction, err := opener.Begin()
	if err != nil {
		return
	}

	return combine(transaction, managers...), nil
}
