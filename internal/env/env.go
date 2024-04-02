package env

import (
	"fmt"
	"os"
	"sms_portal/internal/ui"
	"strings"

	"github.com/joho/godotenv"
)

var AppConfig struct {
	FrontendDomain string
	DatabaseName   string
	DatabaseUser   string
	DatabasePass   string
	DatabaseHost   string
	DatabasePort   string
	ServerPort     string
}

func Init() {
	fmt.Println("Initializing environment variables")
	if err := godotenv.Load(); err != nil {
		ui.Error("No .env file found")
	}

	AppConfig.FrontendDomain = getVarOrDefault("FRONTEND_DOMAIN", "http://localhost:5173")
	AppConfig.DatabaseName = getVarOrDefault("DB_NAME", "postgres")
	AppConfig.DatabaseUser = getVarOrDefault("DB_USER", "postgres")
	AppConfig.DatabasePass = getVarOrDefault("DB_PASS", "postgres")
	AppConfig.DatabaseHost = getVarOrDefault("DB_HOST", "localhost")
	AppConfig.DatabasePort = getVarOrDefault("DB_PORT", "5432")
	AppConfig.ServerPort = getVarOrDefault("SERVER_PORT", "8080")

}

func getVarOrDefault(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func GetConnectionString() string {

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", AppConfig.DatabaseUser, AppConfig.DatabasePass, AppConfig.DatabaseHost, AppConfig.DatabasePort, AppConfig.DatabaseName)
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

func (e Env) Get(key, defaultValue string) string {
	godotenv.Load()
	return getVarOrDefault(key, defaultValue)
}
