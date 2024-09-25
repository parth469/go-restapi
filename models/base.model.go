package models

import "time"

type Basc struct {
	ID       string    `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
}
