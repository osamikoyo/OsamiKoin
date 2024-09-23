package hashing

import (
	"crypto/sha256"
	"fmt"

	"osamikoin/internal/models"
)

func Hash(trans models.Transate) (string) {
	var str string
	resid := sha256.Sum256([]byte(trans.Identify + trans.Id_From + trans.Id_to + trans.Date + trans.Pre_hash))
	for i := 0; i < len(resid); i++ {
		str = fmt.Sprintf("%s%s", str, string(resid[i]))
	}
	return string(str)
}
