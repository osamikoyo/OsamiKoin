package handler

import (
	"fmt"
	"log"

	"osamikoin/internal/db"
	"osamikoin/internal/models"
)

func RegisterCLI() error {
	db := db.New()
	var Profiles []models.Profile
	var prof models.Profile

	log.Println("Input your username\n--->")
	fmt.Scanln(&prof.Username)
	res := db.Find(&Profiles)
	if res.Error != nil {
		return res.Error
	}

	log.Println("Input password to ", prof.Username+"\n--->")
	fmt.Scanln(&prof.Password)

	err := db.Register(prof)
	if err != nil {
		return err
	}

	return nil
}
func GetSending(username string) error {
	var Transate []models.Transate

	db := db.New()
	if err := db.Where("id_from = ? OR id_to = ?", username, username).Find(&Transate).Error; err != nil {
		return err
	}
	for i := 0; i < len(Transate); i++ {
		fmt.Println("ID - ", Transate[i].ID)
		fmt.Println("id_to - ", Transate[i].Id_to)
		fmt.Println("id_from - ", Transate[i].Id_From)
		fmt.Println("count - ", Transate[i].Count)
		fmt.Println("date - ", Transate[i].Date)
	}
	return nil
}
