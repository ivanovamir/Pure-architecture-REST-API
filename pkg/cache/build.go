package cache

type Option func(cfg *Config)

func WithAddress(address string) Option {
	return func(cfg *Config) {
		cfg.Address = address
	}
}

func WithPassword(password string) Option {
	return func(cfg *Config) {
		cfg.Password = password
	}
}

func WithDB(db int) Option {
	return func(cfg *Config) {
		cfg.DB = db
	}
}
