package storage

import (
	"gorm.io/gorm"
	"initCake/pkg/ormx"
)

func New(cfg ormx.DBConfig) (*gorm.DB, error) {
	db, err := ormx.New(cfg)
	if err != nil {
		return nil, err
	}

	return db, nil
}
