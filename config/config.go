package config

type Core struct {
	Port int
	Bind string
}

type Security struct {
	PasswordSalt  string
	JWTSecret     string
	JWTExpiration int
}

type Config struct {
	Core     `mapstructure:"core"`
	Security `mapstructure:"security"`
}

var C Config
