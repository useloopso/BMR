package types

type ConfigInfo struct {
	PrivKey string `mapstructure:"PRIVATE_KEY"`
	Mode    string `mapstructure:"MODE"`
}

type ChainInfo struct {
	RpcURL        string   `mapstructure:"rpc_url"`
	FallbackRpcs  []string `mapstructure:"fallback_rpcs"`
	ChainID       int      `mapstructure:"chain_id"`
	BridgeAddress string   `mapstructure:"bridge_address"`
}
