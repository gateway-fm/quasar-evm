package daemon

type Request struct {
	Id      any    `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  *Params
}

type Params struct {
	Coins       []string `json:"coins,omitempty"`
	Exchanges   []string `json:"exchanges,omitempty"`
	FrequencyMs int      `json:"frequency_ms,omitempty"`
}
