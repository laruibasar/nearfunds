// Package database will handle the interaction between application logic and database service.
package database

import (
	"context"

	"github.com/google/uuid"
	"github.com/laruibasar/nearfunds/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database interface {
	CreateOrder(order models.Order) (uuid.UUID, error)
	Ping(ctx context.Context) error
	Close() error
}

type database struct {
	db *gorm.DB
}

func New(cfg DBConfig) (Database, error) {
	db, err := gorm.Open(sqlite.Open(cfg.DSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		return nil, err
	}

	return &database{
		db: db,
	}, nil
}

func (d *database) Ping(ctx context.Context) error {
	db, err := d.db.DB()
	if err != nil {
		return err
	}

	return db.PingContext(ctx)
}

func (d *database) Close() error {
	db, err := d.db.DB()
	if err != nil {
		return err
	}

	return db.Close()
}

func (d *database) CreateOrder(order models.Order) (uuid.UUID, error) {
	return uuid.New(), nil
}
