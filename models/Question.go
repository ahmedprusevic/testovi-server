package models

import (
	"time"

	pq "github.com/lib/pq"
)

type Question struct {
	ID        uint32         `gorm:"primary_key;auto_increment" json:"id"`
	Name      string         `gorm:"size:1000;not null;" json:"name"`
	Answers   pq.StringArray `gorm:"type:string[];not null;" json:"answers"`
	Correct   pq.StringArray `gorm:"type:string[];not null;" json:"correct"`
	CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	GroupID   uint32         `gorm:"type:bigserial;not null;" json:"group_id"`
	TestID    uint32         `gorm:"type:bigserial;" json:"test_id"`
}
