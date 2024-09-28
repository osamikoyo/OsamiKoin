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

	log.Println("Input password to ", prof.Username + "\n--->")
	fmt.Scanln(&prof.Password)
	
	err := db.Register(prof)
	if err != nil {
		return err
	}

	return nil
}
