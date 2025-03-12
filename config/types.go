package config

// DispatcherConfig holds configuration for the dispatcher service
type DispatcherConfig struct {
	DBHost     string `json:"db_host"`
	DBPort     int    `json:"db_port"`
	DBUser     string `json:"db_user"`
	DBPassword string `json:"db_password"`
	DBName     string `json:"db_name"`
	HTTPPort   int    `json:"http_port"`
	TCPPort    int    `json:"tcp_port"`
}

// ValidatorConfig holds configuration for the validator service
type ValidatorConfig struct {
	DBHost     string `json:"db_host"`
	DBPort     int    `json:"db_port"`
	DBUser     string `json:"db_user"`
	DBPassword string `json:"db_password"`
	DBName     string `json:"db_name"`
	Dispatcher string `json:"dispatcher"`
	Port       int    `json:"port"`
}
