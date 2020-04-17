package bot

type Bot struct{}

func (b *Bot) Run() error {
	// bot, err := tgbotapi.NewBotAPI(config.Viper.Telegram.Token)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// bot.Debug = true

	// u := tgbotapi.NewUpdate(0)
	// u.Timeout = 60

	// updates, err := bot.GetUpdatesChan(u)

	// db.Init()

	// for update := range updates {
	// 	msg.Updates(bot, update)
	// }
	panic("not implemented")
}

func (b *Bot) Stop() error {
	panic("not implemented")
}
