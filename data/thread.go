package data

import (
	"time"

	"github.com/prometheus/common/log"
)

type Thread struct {
	ID        int       `json:"id"`
	Uuid      string    `json:"uuid"`
	Topic     string    `json:"topic"`
	UserID    int       `json:"userID"`
	CreatedAt time.Time `json:"createdAt"`
}

func Threads() (threads []Thread, err error) {
	rows, err := DB.Query("SELECT id, uuid, topic, user_id, created_at FROM threads ORDER BY created_at DESC")
	if err != nil {
		return
	}
	for rows.Next() {
		conv := Thread{}
		if err = rows.Scan(&conv.ID, &conv.Uuid, &conv.Topic, &conv.UserID, &conv.CreatedAt); err != nil {
			log.Info(err)
			return
		}
		threads = append(threads, conv)
		log.Info(conv)
	}
	rows.Close()
	return
}
