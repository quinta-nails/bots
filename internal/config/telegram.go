package config

type TelegramConfig struct {
	IsTestEnvironment bool `env:"TELEGRAM_IS_TEST_ENVIRONMENT,required" envDefault:"true"`
}
