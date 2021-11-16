package model

import "time"

type LogRecord struct {
	Id         uint32    `db:"id"`
	FirstName  string    `db:"first_name"`
	LastName   string    `db:"last_name"`
	UserId     uint32    `db:"user_id"`
	ActionDate time.Time `db:"action_day"`
	ActionTime time.Time `db:"action_time"`
}
