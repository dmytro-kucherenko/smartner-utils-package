package server

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server/errors"
)

func ConnectSQL(connection string, timeout time.Duration) (*sql.DB, error) {
	dbChan := make(chan *sql.DB, 1)
	errChan := make(chan error, 1)

	go func() {
		db, err := sql.Open("postgres", connection)
		if err != nil {
			errChan <- err

			return
		}

		err = db.Ping()
		if err != nil {
			errChan <- err

			return
		}

		dbChan <- db
	}()

	select {
	case db := <-dbChan:
		return db, nil
	case err := <-errChan:
		return nil, err
	case <-time.After(timeout):
		return nil, errors.NewHttpError(http.StatusGatewayTimeout, "database connection timeout reached")
	}
}
