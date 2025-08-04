package entity

type UserEntity struct {
	ID         int64
	Name       string
	Email      string
	Password   string
	Phone      string
	Photo      string
	RoleName   string
	Address    string
	Lat        float64
	Lng        float64
	IsVerified bool
}
