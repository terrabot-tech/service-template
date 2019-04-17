package models

import (
	"time"
)

// Notification struct
type Notification struct {
	ID          uint64     `json:"id"`
	Body        string     `json:"body"`
	TypeID      uint64     `json:"typeId"`
	CreateDate  *time.Time `json:"createDate"`
	ModifyDate  *time.Time `json:"modifyDate"`
	Deleted     uint       `json:"deleted"`
}