package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config 定義應用設定
type Config struct {
	GRPCPort string
	DBSource string
}

// LoadConfig 載入 .env 設定
func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default values")
	}

	return &Config{
		GRPCPort: getEnv("GRPC_PORT", ":50051"),
		DBSource: getEnv("DB_SOURCE", "postgres://user:password@localhost:5432/users"),
	}
}

// getEnv 讀取環境變數，若無則使用預設值
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
