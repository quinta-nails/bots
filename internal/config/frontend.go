package config

type FrontendConfig struct {
	URL string `env:"FRONTEND_URL,required" envDefault:"localhost:8080"`
}
