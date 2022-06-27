package models

import "time"

// User schema of the user table
type User struct {
	ID        int64     `json:"id"`
	UserName  string    `json:"username"`
	Password  byte      `json:"password"`
	FullName  string    `json:"fullname"`
	Dob       time.Time `json:"dob"`
	Mob       int       `json:"mob"`
	Active    bool      `json:"isActive"`
	CreatedAt time.Time `json:"created"`
}
