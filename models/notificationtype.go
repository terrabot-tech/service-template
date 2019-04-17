package models

import (
	"time"
)

// NotificationType struct
type NotificationType struct {
	ID          uint64     `json:"id"`
	Name        string     `json:"body"`
	CreateDate  *time.Time `json:"createDate"`
	ModifyDate  *time.Time `json:"modifyDate"`
	Deleted     uint       `json:"deleted"`
}
