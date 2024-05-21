package infrastructure

type DBConfig struct {
	Host     string	`yaml:"host"`
	Port     int 	`yaml:"port"`
	User     string
	Password string
	Database string	`yaml:"database"`
	Driver   string	`yaml:"driver"`
}
