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

func (c *MySqlClient) List(accountID string) ([]balance.Model, error) {
	query := "SELECT account_id_from, account_id_to, balance_account_id_from, balance_account_id_to FROM balances where account_id_from = ? OR account_id_to = ?"
	stmt, err := c.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(accountID, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	balances := make([]balance.Model, 0)
	for rows.Next() {
		var (
			accountIDFrom        string
			accountIDTo          string
			balanceAccountIDFrom int
			balanceAccountIDTo   int
		)
		err = rows.Scan(&accountIDFrom, &accountIDTo, &balanceAccountIDFrom, &balanceAccountIDTo)
		if err != nil {
			return nil, err
		}
		model := balance.NewModel(accountIDFrom, accountIDTo, balanceAccountIDFrom, balanceAccountIDTo)
		balances = append(balances, model)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return balances, nil
}

func (c *MySqlClient) Save(payload balance.UpdateEventPayload) error {
	query := "INSERT INTO balances (account_id_from, account_id_to, balance_account_id_from, balance_account_id_to) VALUES (?, ?, ?, ?)"
	stmt, err := c.DB.Prepare(query)
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
