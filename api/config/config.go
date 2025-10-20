package config

import (
	"log"
	"workplace_api/utils"// importa suas funções utilitárias
)

// DBConfig representa a estrutura de configuração do banco de dados.
type DBConfig struct {
	Host             string
	Port             int
	User             string
	Password         string
	Database         string
	WaitConnections  bool
	ConnectionLimit  int
	QueueLimit       int
	ConnectTimeoutMs int
}

// LoadDBConfig retorna a configuração do banco de dados baseada no ambiente.
func LoadDBConfig() DBConfig {
	env := utils.GetEnvOrDefault("NODE_ENV", "development")

	cfg := DBConfig{
		Host:             utils.GetEnvOrDefault("DB_HOST", "localhost"),
		Port:             utils.ParseInt(utils.GetEnvOrDefault("DB_PORT", "3306"), 3306),
		User:             utils.GetEnvOrDefault("DB_USER", "root"),
		Password:         utils.GetEnvOrDefault("DB_PASSWORD", ""),
		Database:         utils.GetEnvOrDefault("DB_NAME", "sistema_WP"),
		WaitConnections:  utils.ParseBool(utils.GetEnvOrDefault("DB_WAIT_CONNECTIONS", "true"), true),
		ConnectionLimit:  utils.ParseInt(utils.GetEnvOrDefault("DB_CONN_LIMIT", "10"), 10),
		QueueLimit:       utils.ParseInt(utils.GetEnvOrDefault("DB_QUEUE_LIMIT", "50"), 50),
		ConnectTimeoutMs: utils.ParseInt(utils.GetEnvOrDefault("DB_CONN_TIMEOUT", "5000"), 5000),
	}

	if env != "production" {
		log.Printf("Configuração do banco carregada (%s): %s@%s:%d/%s",
			env, cfg.User, cfg.Host, cfg.Port, cfg.Database)
	}

	return cfg
}
