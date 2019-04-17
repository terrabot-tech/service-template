package repository

import (
	"errors"

	"github.com/terrabot-tech/service-template/models"
)

// GetNotificationsByUserID Получить уведомления для пользователя
func (rep *Repository) GetNotificationsByUserID(userID uint64) ([]models.Notification, error) {
	notifications := make([]models.Notification, 0)
	db, _ := rep.provider.GetConn()

	stmt, err := db.Prepare("SELECT * FROM public.get_notifications_for_user($1, $2);")
	if err != nil {
		return notifications, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(userID, 0)
	if err != nil {
		return notifications, err
	}
	defer rows.Close()

	for rows.Next() {
		var notification models.Notification
		if err := rows.Scan(
			&notification.ID,
			&notification.Body,
			&notification.CreateDate,
			&notification.TypeID,
			&notification.ModifyDate,
			&notification.Deleted);
			err != nil {
			//Todo: логировать ошибку
		}
		notifications = append(notifications, notification)
	}
	return notifications, nil
}

// GetNotification Получить уведомление по ID
func (rep *Repository) GetNotification(id uint64) (models.Notification, error) {
	var notification models.Notification
	db, _ := rep.provider.GetConn()

	stmt, err := db.Prepare("SELECT * FROM public.get_notifications($1);")
	if err != nil {
		return notification, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return notification, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&notification.ID,
			&notification.Body,
			&notification.CreateDate,
			&notification.TypeID,
			&notification.ModifyDate,
			&notification.Deleted);
			err != nil {
			//Todo: логировать ошибку
		}
		return notification, err
	}
	return notification, errors.New("Неизвестная ошибка.\n\r")
}

// SetNotificationWatched Уведомление просмотрено/не просмотрено
func (rep *Repository) SetNotificationWatched(notificationID uint64, userID uint64, watched int) ([]uint64, error) {
	res := make([]uint64, 0)
	db, _ := rep.provider.GetConn()

	stmt, err := db.Prepare("SELECT * FROM public.set_notification_watched($1, $2, $3);")
	if err != nil {
		return res, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(notificationID, userID, watched)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		var userNotificationID uint64
		if err := rows.Scan(&res);
			err != nil {
			//Todo: логировать ошибку
		}
		res = append(res, userNotificationID)
	}
	return res, nil
}