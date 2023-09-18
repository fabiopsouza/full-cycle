package balance

import (
	"database/sql"
	"github.com/fabiopsouza/balance/internal/core/domain/balance"
	balancePort "github.com/fabiopsouza/balance/internal/core/port/balance"
)

type MySqlClient struct {
	DB *sql.DB
}

func NewMySqlClient(db *sql.DB) balancePort.Repository {
	return &MySqlClient{
		DB: db,
	}
}

func (t *MySqlClient) Save(payload balance.UpdateEventPayload) error {
	query := `INSERT INTO balances (account_id_from, account_id_to, balance_account_id_from, balance_account_id_to) VALUES (?, ?, ?, ?)`
	stmt, err := t.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(payload.AccountIDFrom, payload.AccountIDTo, payload.BalanceAccountIDFrom, payload.BalanceAccountIDTo)
	if err != nil {
		return err
	}
	return nil
}
