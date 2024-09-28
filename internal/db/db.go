package db

import (
	"log"

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
func (db *DB) Register(prof models.Profile) error {
	res := db.Create(&prof)
	if res.Error != nil {
		return res.Error
	}
	return nil
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
		return 0, "", err
	}
	if prof.Money < count {
		return prof.Money, "litle", nil
	}
	return prof.Money, "succes", nil
}
func (db *DB) AcrivateTransate(trans models.Transate) {
	var User_to models.Profile
	var User_from models.Profile

	res := db.Where("id_from = ?", trans.Id_From).Find(&User_from)
	res2 := db.Where("id_to = ?", trans.Id_to).Find(&User_to)
	res3 := db.Model(&User_from).Update("count", User_from.Money-trans.Count)
	res4 := db.Model(&User_to).Update("count", User_to.Money + trans.Count)

	
	if res.Error != nil && res2.Error != nil && res3.Error != nil && res4 != nil{
		log.Fatal(res.Error, res2.Error, res3.Error, res4.Error)
		return
	}
	db.SaveTransateToDB(trans)
}
