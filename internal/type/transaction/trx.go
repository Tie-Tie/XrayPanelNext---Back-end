package transaction

// TrxTransactionInfo ContractMap TokenInfo Data Trc20Res --------------------------------->
// TrxTransactionInfo represents the structure of the response for a single transaction.
type TrxTransactionInfo struct {
	Ret                  []TrxRet                 `json:"ret"`
	Signature            []string                 `json:"signature"`
	TxID                 string                   `json:"txID"`
	NetUsage             int64                    `json:"net_usage"`
	RawDataHex           string                   `json:"raw_data_hex"`
	NetFee               int64                    `json:"net_fee"`
	EnergyUsage          int64                    `json:"energy_usage"`
	BlockNumber          int64                    `json:"blockNumber"`
	BlockTimestamp       int64                    `json:"block_timestamp"`
	EnergyFee            int64                    `json:"energy_fee"`
	EnergyUsageTotal     int64                    `json:"energy_usage_total"`
	RawData              TrxRawData               `json:"raw_data"`
	InternalTransactions []TrxInternalTransaction `json:"internal_transactions"`
}

// TrxRet represents the structure of the `ret` field in the transaction response.
type TrxRet struct {
	ContractRet string `json:"contractRet"`
	Fee         int64  `json:"fee"`
}

// TrxRawData represents the structure of the `raw_data` field in the transaction response.
type TrxRawData struct {
	Contract      []TrxContract `json:"contract"`
	RefBlockBytes string        `json:"ref_block_bytes"`
	RefBlockHash  string        `json:"ref_block_hash"`
	Expiration    int64         `json:"expiration"`
	Timestamp     int64         `json:"timestamp"`
}

// TrxContract represents the structure of the contract field within raw_data.
type TrxContract struct {
	Parameter TrxParameter `json:"parameter"`
	Type      string       `json:"type"`
}

// TrxParameter represents the structure of the parameter field within a contract.
type TrxParameter struct {
	Value   TrxContractValue `json:"value"`
	TypeURL string           `json:"type_url"`
}

// TrxContractValue represents the structure of the value field within parameter.
type TrxContractValue struct {
	Amount       int64  `json:"amount"`
	OwnerAddress string `json:"owner_address"`
	ToAddress    string `json:"to_address"`
}

// TrxInternalTransaction represents the structure of internal transactions in the response.
type TrxInternalTransaction struct {
	// Define fields based on the internal transaction response if needed
}

// TrxRes represents the structure of the entire API response.
type TrxRes struct {
	Data    []TrxTransactionInfo `json:"data"`
	Success bool                 `json:"success"`
	Meta    TrxMeta              `json:"meta"`
}

// TrxMeta represents the metadata section of the API response.
type TrxMeta struct {
	At       int64 `json:"at"`
	PageSize int32 `json:"page_size"`
}
