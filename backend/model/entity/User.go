package entity

import "time"

type User struct {
	ID        int64
	Username  string
	Password  string
	Gender    int64
	Age       int64
	Weight    int64
	BMI       int64
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}
