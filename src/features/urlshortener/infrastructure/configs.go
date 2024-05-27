package infrastructure

type DBConfig struct {
	Host     string	`yaml:"host"`
	Port     int 	`yaml:"port"`
	User     string
	Password string
	Database string	`yaml:"database"`
	Driver   string	`yaml:"driver"`
}

type PoolConfig struct {
	MaxIdleConns int `yaml:"maxIdleConns"`
	MaxOpenConns int `yaml:"maxOpenConns"`
	ConnMaxIdleTime int64 `yaml:"connMaxIdleTime"`
	ConnMaxLifetime int64 `yaml:"connMaxLifetime"`
}