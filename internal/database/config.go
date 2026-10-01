package database

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type DBConfig struct {
	Path               string
	Timeout            time.Duration
	MaxOpenConnections int
	MaxIdleConnections int
	ConnectionMaxLife  time.Duration
}

// NewConfig is for my local SQLite instance, for Postgresql, should be reviewed.
func NewConfig() (*DBConfig, error) {
	cfg := DBConfig{
		Path: getEnv("DB_PATH", "test.db"),
	}

	var err error
	if cfg.Timeout, err = getDuration("DB_TIMEOUT", 5*time.Second); err != nil {
		return nil, err
	}

	if cfg.MaxOpenConnections, err = getInt("DB_MAX_OPEN_CONNECTIONS", 2); err != nil {
		return nil, err
	}

	if cfg.MaxIdleConnections, err = getInt("DB_MAX_IDLE_CONNECTIONS", 2); err != nil {
		return nil, err
	}

	if cfg.ConnectionMaxLife, err = getDuration("DB_CONN_MAX_LIFE", 0); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (c *DBConfig) DSN() string {
	return fmt.Sprintf("file:%s?_busy_timeout=%d&_foreign_keys=on", c.Path, c.Timeout.Milliseconds())
}

func getEnv(key, value string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}

	return value
}

func getDuration(key string, value time.Duration) (time.Duration, error) {
	v := os.Getenv(key)
	if v == "" {
		return value, nil
	}

	duration, err := time.ParseDuration(v)
	if err != nil {
		return 0, fmt.Errorf("database variable %s failed: %v", key, err)
	}

	return duration, nil
}

func getInt(key string, value int) (int, error) {
	v := os.Getenv(key)
	if v == "" {
		return value, nil
	}

	i, err := strconv.Atoi(v)
	if err != nil {
		return 0, fmt.Errorf("database variable %s failed: %v", key, err)
	}

	return i, nil
}
