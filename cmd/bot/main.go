package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"telegram-weather-bot/pkg/bot"
	"telegram-weather-bot/pkg/config"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cmd = &cobra.Command{
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			_ = viper.BindEnv("google.geocoding.token")
			_ = viper.BindEnv("owm.token")
			_ = viper.BindEnv("telegram.token")
			_ = viper.BindEnv("telegram.error.admin")
			_ = viper.BindEnv("dsn")

			if err := viper.ReadInConfig(); err != nil {
				log.Fatal(err)
			}

			if err := viper.Unmarshal(config.Cfg); err != nil {
				log.Fatal(err)
			}
		},

		Run: func(cmd *cobra.Command, args []string) {
			bot := bot.Bot{}

			go func() {
				if err := bot.Run(); err != nil {
					log.Fatal(err)
				}
			}()

			interrupt := make(chan os.Signal, 1)
			signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
			<-interrupt

			if err := bot.Stop(); err != nil {
				log.Fatal(err)
			}
		},
	}
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
