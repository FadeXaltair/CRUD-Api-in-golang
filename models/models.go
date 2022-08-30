package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	TaskName        string
	TaskDescription string
}
