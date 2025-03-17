package repository

import (
	"gorm.io/gorm"
)

type IRepository interface {
	WithTx(fn func(tx *gorm.DB) error) error
}

type BaseRepository struct {
	DB *gorm.DB
}

func NewBaseRepository(db *gorm.DB) BaseRepository {
	return BaseRepository{DB: db}
}

func (r *BaseRepository) WithTx(fn func(tx *gorm.DB) error) error {
	tx := r.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r) // re-throw panic after rollback
		}
	}()

	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
