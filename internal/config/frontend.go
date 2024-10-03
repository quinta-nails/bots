package config

type FrontendConfig struct {
	Url uint16 `env:"FRONTEND_URL,required" envDefault:"localhost:8080"`
}
