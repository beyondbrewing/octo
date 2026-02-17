package common

import (
	"log"

	"github.com/beyondbrewing/octo/utils"
	"github.com/spf13/viper"
)

var (
	APP_NAME    string
	APP_VERSION string
)

var (
	PEBBLE_DATABASE_DIR string

	ENV_BASE_CHAIN string
	ENV_MAX_PEERS  string
	ENV_ENODE      []string
)

func LoadConfig(filepath string) {
	path, err := utils.ReturnAbsolutePath(filepath)
	if err != nil {
		log.Fatalf("failed to find a path : %v", err)
	}
	if err := utils.ReadConfigutableVariables(path); err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	PEBBLE_DATABASE_DIR = viper.GetString("pebble.database_dir")
	ENV_BASE_CHAIN = viper.GetString("chain.chain_param")
	ENV_MAX_PEERS = viper.GetString("chain.max_peers")
	ENV_ENODE = viper.GetStringSlice("chain.enode")
}
