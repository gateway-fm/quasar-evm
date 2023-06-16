package daemon

type Response struct {
	Id      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
}

type Result struct {
	Coin      string  `json:"coin"`
	Method    string  `json:"method"`
	Timestamp int     `json:"timestamp"`
	Value     float64 `json:"value"`
}
