package config

type FrontendConfig struct {
	URL string `env:"FRONTEND_URL,required" envDefault:"http://127.0.0.1:8080"`
}
