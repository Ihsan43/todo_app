package model

import "time"

type Todo struct {
	Id         string    `json:"id"`
	Name       string    `json:"name"`
	Status     string    `json:"status"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"updated_at"`
}
