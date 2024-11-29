package structs

import "time"

type Model struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Url         string    `json:"connection_string"`
	DateCreated time.Time `json:"created_at"`
}
