package config

// ** Config struct keeps all needed configurations for application
type Config struct {
	Postgres Postgres
}

// * Postgres struct is
type Postgres struct {
	Host     string `envconfig:"db_host"`
	Port     string `envconfig:"db_port"`
	User     string `envconfig:"db_user"`
	Password string `envconfig:"db_password"`
	Name     string `envconfig:"db_name"`
	SslMode  string `envconfig:"db_sslmode"`
}
