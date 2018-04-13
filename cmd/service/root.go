package main

import (
	"log"
	"os"

	"github.com/spf13/viper"

	"github.com/chespinoza/test-zenroom-go/pkg/globals"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	configFile string
)

var RootCmd = &cobra.Command{
	Use:   "prototype",
	Short: "Prototype ...",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	hostName, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	logger, err := zap.NewProduction(
		zap.Fields(zap.String("host", hostName)),
	)
	if err != nil {
		log.Fatal(err)
	}

	globals.Logger = *logger

	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().String("myval", "", "My val")

	viper.BindPFlag("myval", RootCmd.PersistentFlags().Lookup("myval"))

}

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigName("config.yml")
	}
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		globals.Logger.Info("Config file loaded", zap.String("file", viper.ConfigFileUsed()))
	}

}
