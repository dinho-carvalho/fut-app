package database

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm/logger"
)

type Config struct {
	Host            string
	User            string
	Password        string
	DBName          string
	Port            string
	SSLMode         string
	TimeZone        string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
	LogLevel        logger.LogLevel
}

func NewConfig() *Config {
	return &Config{
		Host:            getEnv("DB_HOST", "localhost"),
		User:            getEnv("DB_USER", "admin"),
		Password:        getEnv("DB_PASSWORD", "admin"),
		DBName:          getEnv("DB_NAME", "futebol_stats"),
		Port:            getEnv("DB_PORT", "5432"),
		SSLMode:         getEnv("DB_SSLMODE", "disable"),
		TimeZone:        getEnv("DB_TIMEZONE", "America/Sao_Paulo"),
		MaxIdleConns:    getEnvAsInt("DB_MAX_IDLE_CONNS", 10),
		MaxOpenConns:    getEnvAsInt("DB_MAX_OPEN_CONNS", 100),
		ConnMaxLifetime: getEnvAsDuration("DB_CONN_MAX_LIFETIME", time.Hour),
		LogLevel:        getEnvAsLogLevel("DB_LOG_LEVEL", logger.Error),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if i, err := strconv.Atoi(value); err == nil {
			return i
		}
	}
	return defaultValue
}

func getEnvAsDuration(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if d, err := time.ParseDuration(value); err == nil {
			return d
		}
	}
	return defaultValue
}

func getEnvAsLogLevel(key string, defaultValue logger.LogLevel) logger.LogLevel {
	levels := map[string]logger.LogLevel{
		"silent": logger.Silent,
		"error":  logger.Error,
		"warn":   logger.Warn,
		"info":   logger.Info,
	}
	if value := os.Getenv(key); value != "" {
		if level, exists := levels[strings.ToLower(value)]; exists {
			return level
		}
	}
	return defaultValue
}

func (c *Config) GetDSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		c.Host, c.User, c.Password, c.DBName, c.Port, c.SSLMode, c.TimeZone,
	)
}
