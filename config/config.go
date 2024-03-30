package config

type DbConfig struct {
	DbName   string
	DbUser   string
	DbPort   string
	DbHost   string
	DbPass   string
	DbDriver string
}

type Server struct {
	ApiPort string
	ApiHost string
}

type Config struct {
	DbConfig
	Server
}
