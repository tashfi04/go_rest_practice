package models

type User struct {
	Id        int		`json:"id"`
	UserName string		`json:"user_name"`
	UserType string		`json:"user_type"`
}
