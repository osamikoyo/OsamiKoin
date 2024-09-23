package db

import (
	"gorm.io/gorm"

	"osamikoin/internal/hashing"
	"osamikoin/internal/models"
)

type DB struct{
	*gorm.DB
}

func (db *DB) SaveTransateToDB(t models.Transate) error {
	var PreHash models.Transate

	rem := db.Order("id desc").Limit(1).Find(&PreHash)
	if rem.Error != nil {
		return rem.Error
	}

	t.Hash = hashing.Hash(t)
	t.Pre_hash = PreHash.Pre_hash

	res := db.Create(&t)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
