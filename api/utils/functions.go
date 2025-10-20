package utils

import (
	"os"
	"strconv"
)

// GetEnvOrDefault retorna o valor da variável de ambiente ou um padrão se ela não existir.
func GetEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// ParseInt converte string para int com valor padrão em caso de erro.
func ParseInt(value string, fallback int) int {
	num, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return num
}

// ParseBool converte string ("true"/"false") para bool com valor padrão.
func ParseBool(value string, fallback bool) bool {
	if value == "" {
		return fallback
	}
	b, err := strconv.ParseBool(value)
	if err != nil {
		return fallback
	}
	return b
}
