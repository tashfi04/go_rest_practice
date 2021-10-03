package models

type Class struct {
	Id        	int    `json:"id"`
	Name 		string `json:"name"`
	GroupName	string `json:"group_name"`
	ClassOrder	int    `json:"class_order"`
}
