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
