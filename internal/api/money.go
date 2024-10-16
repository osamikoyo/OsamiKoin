package api

import (
	"github.com/labstack/echo/v4"

	"osamikoin/internal/db"
	"osamikoin/internal/models"
)

func Home(c echo.Context) error {
	return c.File("public/src/index.html")
}
func Send(c echo.Context) error {
	var transAction models.Transate
	if err := c.Bind(&transAction); err != nil {
		return err
	}
	db := db.New()
	db.AcrivateTransate(transAction)
	return nil
}
