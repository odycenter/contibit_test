package pkg

type Transfer struct {
	UserId       int64   `json:"user_id"`
	TargetId     int64   `json:"target_id"`
	TransferType string  `json:"type"`
	Coin         string  `json:"coin"`
	Amount       float64 `json:"amount"`
	ApiTimestamp int64   `json:"api_timestamp"`
	ApiDatetime  string  `json:"api_datetime"`
	ApiKey       string  `json:"api_key"`
}

type TransferChannel struct {
	ApiTimestamp int64
	TransferInfo string
}
