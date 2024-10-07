package config

type ServicesConfig struct {
	CompaniesEndpoint string `env:"GRPC_COMPANIES_ENDPOINT,required" envDefault:"127.0.0.1:51052"`
}
