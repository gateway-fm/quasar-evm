package config

import (
	"github.com/spf13/viper"
)

// init initialize default config params
func init() {
	// environment - could be "local", "prod", "dev"
	viper.SetDefault("env", "prod")

	viper.SetDefault("tokens", []string{
		"ADA", "ALGO", "ARB",
		"AVAX", "BNB", "BTC",
		"DOGE", "ETH", "FIL",
		"FLR", "LTC", "MATIC",
		"SOL", "USDC", "USDT",
		"XDC", "XLM", "XRP",
	})

	viper.SetDefault("ws.url", "wss://oracle.gateway.fm")
	viper.SetDefault("ws.WriteFrequencyMS", 5000)

	viper.SetDefault("web3.SendFrequencyMS", 10000)
	viper.SetDefault("web3.RpcURL", "")
	viper.SetDefault("web3.SignerPK", "")
	viper.SetDefault("web3.ChainID", 0)
	viper.SetDefault("web3.OracleAddress", "")
}
