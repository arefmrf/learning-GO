package config

type Config struct {
	Server Server
	DB     DB
}

type Server struct {
	Host      string
	Port      string
	JwtSecret string
}

type DB struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
}
