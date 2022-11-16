package pkg

type MysqlUser struct {
	Id          int64   `json:"id"`
	User        string  `json:"user"`
	Passwd      string  `json:"passwd"`
	BalanceUSDT float64 `json:"balance_usdt"`
	UserStatus  int64   `json:"user_status"`
}

type MysqlTransfer struct {
	Id             int64   `json:"id"`
	UserId         int64   `json:"user_id"`
	TargetId       int64   `json:"target_id"`
	TransferType   int64   `json:"transfer_type"`
	Coin           string  `json:"coin"`
	Amount         float64 `json:"amount"`
	Status         int64   `json:"status"`
	ApiTimestamp   string  `json:"api_timestamp"`
	ApiDatetime    string  `json:"api_datetime"`
	ApiKey         string  `json:"api_key"`
	CreateDatetime string  `json:"create_datetime"`
}
