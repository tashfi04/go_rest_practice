package models

type Student struct {
	Id        		int    `json:"id"`
	UserId   		int    `json:"user_id"`
	FullName 		string `json:"full_name"`
	Gender      	string `json:"gender"`
	DateOfBirth  	string `json:"date_of_birth"`
	CurrentClass	int    `json:"current_class"`
}
