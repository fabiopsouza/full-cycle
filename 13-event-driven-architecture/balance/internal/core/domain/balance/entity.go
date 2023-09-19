package balance

type UpdateEvent struct {
	Name    string             `json:"name"`
	Payload UpdateEventPayload `json:"payload"`
}

type UpdateEventPayload struct {
	AccountIDFrom        string `json:"account_id_from"`
	AccountIDTo          string `json:"account_id_to"`
	BalanceAccountIDFrom int    `json:"balance_account_id_from"`
	BalanceAccountIDTo   int    `json:"balance_account_id_to"`
}

type Model struct {
	AccountIDFrom        string `json:"account_id_from"`
	AccountIDTo          string `json:"account_id_to"`
	BalanceAccountIDFrom int    `json:"balance_account_id_from"`
	BalanceAccountIDTo   int    `json:"balance_account_id_to"`
}

func NewModel(accountIDFrom, accountIDTo string, balanceAccountIDFrom, balanceAccountIDTo int) Model {
	return Model{
		AccountIDFrom:        accountIDFrom,
		AccountIDTo:          accountIDTo,
		BalanceAccountIDFrom: balanceAccountIDFrom,
		BalanceAccountIDTo:   balanceAccountIDTo,
	}
}
