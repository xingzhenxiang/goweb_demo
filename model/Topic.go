package model

import "time"

type Topic struct {
	ID            uint      `json:"id"`
	Title         string    `json:"title"`
	TopicsPreview string    `json:"topics_preview"`
	TopicsInfo    string    `json:"topics_info"`
	CreateTime    time.Time `json:"create_time"`
	ModifyTime    time.Time `json:"modify_time"`
}
