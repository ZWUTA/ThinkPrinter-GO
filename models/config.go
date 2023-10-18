package models

type Config struct {
	Core struct {
		Port     int
		Bind     string
		LogLevel string
	}
	Print struct {
		WordExePath string
	}
	Security struct {
		PasswordSalt  string
		JWTSecret     string
		JWTExpiration int
	}
}
