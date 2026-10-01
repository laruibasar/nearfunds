// Package database will handle the interaction between application logic and database service.
package database

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/laruibasar/nearfunds/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database interface {
	FindOrderByCorrelation(id uuid.UUID) (*models.Order, error)
	CreateOrder(order models.Order) (*uint, error)
	Ping(ctx context.Context) error
	Close() error
}

type database struct {
	db *gorm.DB
}

func New(cfg *DBConfig) (Database, error) {
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

// CreateOrder handles the store of the order into database.
// Uses a transaction.
func (d *database) CreateOrder(order models.Order) (*uint, error) {
	tx := d.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	dbOrder := &Order{
		Account: order.Account,
	}

	if err := tx.Create(&dbOrder).Error; err != nil {
		tx.Rollback()

		return nil, err
	}

	return &dbOrder.ID, tx.Commit().Error
}

// FindOrderByCorrelation allows to search the database for the correlation id that is associated with the order.
func (d *database) FindOrderByCorrelation(id uuid.UUID) (*models.Order, error) {
	var dbOrder Order
	find := d.db.First(&dbOrder, "correlation_id = ?", id)

	if find.Error != nil {
		if errors.Is(find.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("not found")
		}
		return nil, find.Error
	}

	return &models.Order{
		Account: dbOrder.Account,
	}, nil
}
