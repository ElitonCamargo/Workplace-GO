package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"workplace_api/config"

	_ "github.com/go-sql-driver/mysql" // Importa o driver do MySQL anonimamente
)

// DB é a instância global do pool de conexões.
var DB *sql.DB

// ConnectMySQL inicializa o pool de conexões MySQL.
func ConnectMySQL() {
	cfg := config.LoadDBConfig()

	// Monta a string de conexão (DSN)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)

	// Abre a conexão com o banco de dados (sem se conectar ainda)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Erro ao abrir conexão com o banco: %v", err)
	}

	// Aplica as configurações de pool de conexões
	db.SetMaxOpenConns(cfg.ConnectionLimit) // Limita conexões abertas
	db.SetMaxIdleConns(cfg.QueueLimit) // Limita conexões ociosas
	db.SetConnMaxLifetime(time.Duration(cfg.ConnectTimeoutMs) * time.Millisecond) // Tempo máximo de vida da conexão

	// Testa a conexão real
	if err := db.Ping(); err != nil {
		log.Fatalf("Erro ao conectar ao banco: %v", err)
	}

	log.Println("✅ Conexão com o banco de dados estabelecida com sucesso!")
	DB = db
}
// CloseDB fecha a conexão com o banco de dados.
func CloseDB() {
	if DB != nil {
		if err := DB.Close(); err != nil {
			log.Printf("Erro ao fechar a conexão com o banco: %v", err)
		} else {
			log.Println("Conexão com o banco de dados fechada.")
		}
	}
}