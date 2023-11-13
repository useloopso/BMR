package config

import (
	"github.com/spf13/viper"
	"github.com/useloopso/BMR/types"
)

var Config types.ConfigInfo
var Chains []types.ChainInfo

func LoadChainInfo() error {
	viper.SetConfigFile("chain_config.json")
	viper.SetConfigType("json")

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	if err := viper.UnmarshalKey("chains", &Chains); err != nil {
		return err
	}
	return nil
}

func Load() error {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		return err
	}

	return nil
}
