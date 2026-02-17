package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	"github.com/beyondbrewing/octo/common"
	"github.com/beyondbrewing/octo/logger"
	flag "github.com/spf13/pflag"
)

var configFile string

func init() {
	flag.StringVarP(&configFile, "config", "c", "", "path to config file")
	flag.Parse()
	common.LoadConfig(configFile)
}

func main() {
	logger.SetDefault(logger.MustProduction())
	defer logger.SyncDefault()

	_, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	fmt.Println(common.ENV_BASE_CHAIN)
}
