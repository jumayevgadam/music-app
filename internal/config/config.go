package config

// Config struct keeps all needed configurations for application
type Config struct {
	Postgres Postgres
	Server   struct {
		HttpPort string `envconfig:"HTTP_PORT" validate:"required"`
	}
}

// Postgres struct is
type Postgres struct {
	Host     string `envconfig:"DB_HOST" validate:"required"`
	Port     string `envconfig:"DB_PORT" validate:"required"`
	User     string `envconfig:"DB_USER" validate:"required"`
	Password string `envconfig:"DB_PASSWORD" validate:"required"`
	Name     string `envconfig:"DB_NAME" validate:"required"`
	SslMode  string `envconfig:"DB_SSLMODE" validate:"required"`
}
