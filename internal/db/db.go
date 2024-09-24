package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"osamikoin/internal/hashing"
	"osamikoin/internal/models"
)

type DB struct {
	*gorm.DB
}

func New() DB {
	db, _ := gorm.Open(sqlite.Open("storage/main.db"))
	return DB{db}
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
func (db *DB) GetProf(username string) (models.Profile, error) {
	var prof models.Profile
	res := db.Where("username = ?", username).Find(&prof)
	if res.Error != nil {
		return prof, res.Error
	}
	return prof, nil
}
func ChekProfile(prof models.Profile, count int) (int, string, error) {
	db := New()
	var err error
	prof, err = db.GetProf(prof.Username)
	if err != nil {
		return 0,"",err
	}
	if prof.Money < count {
		return prof.Money,"litle", nil
	}
	return prof.Money, "succes", nil
}
