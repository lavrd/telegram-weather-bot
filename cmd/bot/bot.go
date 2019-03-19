package main

import (
	"log"
	"path/filepath"
	"strings"

	"telegram-weather-bot/pkg/bot"
	"telegram-weather-bot/pkg/config"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfg string

	CLI = &cobra.Command{
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			abs, err := filepath.Abs(cfg)
			if err != nil {
				log.Fatal(err)
			}

			base := filepath.Base(abs)
			path := filepath.Dir(abs)

			viper.SetConfigName(strings.Split(base, ".")[0])
			viper.AddConfigPath(path)

			if err := viper.ReadInConfig(); err != nil {
				log.Fatal(err)
			}

			if err := viper.Unmarshal(config.Viper); err != nil {
				log.Fatal(err)
			}
		},

		Run: func(cmd *cobra.Command, args []string) {
			bot.Daemon()
		},
	}
)

func init() {
	CLI.Flags().StringVarP(&cfg, "config", "c", "./contrib/config.yml", "/path/to/config.yml")

	viper.BindEnv("google.geocoding.token")
	viper.BindEnv("darksky.token")
	viper.BindEnv("telegram.token")
	viper.BindEnv("telegram.error.admin")
}

func main() {
	if err := CLI.Execute(); err != nil {
		log.Fatal(err)
	}
}
