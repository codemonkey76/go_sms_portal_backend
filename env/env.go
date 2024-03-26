package env

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func getVarOrDefault(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func GetConnectionString() string {
	godotenv.Load()
	db_name := getVarOrDefault("DB_NAME", "postgres")
	db_user := getVarOrDefault("DB_USER", "postgres")
	db_pass := getVarOrDefault("DB_PASS", "postgres")
	db_host := getVarOrDefault("DB_HOST", "localhost")
	db_port := getVarOrDefault("DB_PORT", "5432")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", db_user, db_pass, db_host, db_port, db_name)
	return connStr
}

type Env string

func (e Env) Set(key, value string) error {
	// Read content of .env file
	envContent, err := os.ReadFile(string(e))
	if err != nil {
		return fmt.Errorf("Error reading .env file: %v", err)
	}

	var updatedLines []string
	keyFound := false
	for _, line := range strings.Split(string(envContent), "\n") {
		if strings.HasPrefix(line, key+"=") {
			line = key + "=" + value
			keyFound = true
		}
		updatedLines = append(updatedLines, line)
	}

	if !keyFound {
		updatedLines = append(updatedLines, key+"="+value)
	}

	if err := os.WriteFile(string(e), []byte(strings.Join(updatedLines, "\n")), 0644); err != nil {
		return fmt.Errorf("Error writing .env file: %v", err)
	}

	return nil
}
