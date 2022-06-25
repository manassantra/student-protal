package entity

type UserEntity struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	// Password  byte   `json:"password"`
	Password string `json:"password"`
	Active   bool   `json:"isActive"`
	// CreatedAt string `json:"createdAt"`
}
