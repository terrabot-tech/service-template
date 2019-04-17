package models

import (
	"time"
)

// UserNotification struct
type UserNotification struct {
	ID             uint64     `json:"id"`
	UserID         uint64     `json:"userId"`
	NotificationID uint64     `json:"notificationId"`
	CreateDate     *time.Time `json:"createDate"`
	ModifyDate     *time.Time `json:"modifyDate"`
	Watched        uint       `json:"watched"`
	Deleted        uint       `json:"deleted"`
}
