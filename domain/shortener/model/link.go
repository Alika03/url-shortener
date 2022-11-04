package model

import "time"

type Link struct {
	Id        string    `db:"id" json:"id"`
	Code      string    `db:"code" json:"code"`
	FullUrl   string    `db:"full_url" json:"full_url"`
	ExpiredAt time.Time `db:"created_at" json:"expired_at"`
}
