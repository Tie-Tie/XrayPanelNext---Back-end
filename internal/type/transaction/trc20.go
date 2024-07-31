package transaction

// Trc20TokenInfo 表示代币的信息
type Trc20TokenInfo struct {
	Symbol   string `json:"symbol"`
	Address  string `json:"address"`
	Decimals int    `json:"decimals"`
	Name     string `json:"name"`
}

// Trc20Transaction 表示单个交易信息
type Trc20Transaction struct {
	TransactionID  string         `json:"transaction_id"`
	TokenInfo      Trc20TokenInfo `json:"token_info"`
	BlockTimestamp int64          `json:"block_timestamp"`
	From           string         `json:"from"`
	To             string         `json:"to"`
	Type           string         `json:"type"`
	Value          string         `json:"value"`
}

// Trc20Meta 表示元数据
type Trc20Meta struct {
	At       int64 `json:"at"`
	PageSize int   `json:"page_size"`
}

// Trc20Res 表示API返回的整体结构
type Trc20Res struct {
	Data    []Trc20Transaction `json:"data"`
	Success bool               `json:"success"`
	Meta    Trc20Meta          `json:"meta"`
}
