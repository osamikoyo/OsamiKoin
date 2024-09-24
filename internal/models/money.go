package models

type Transate struct{
	ID    uint   `gorm:"primaryKey"`
	Date string
	Identify string
	Id_From string
	Id_to string
	Count int
	Hash string
	Pre_hash string
}
type Profile struct{
	Username string
	Password string
	Money int
}