package models

import "github.com/manicar2093/web-monitor/entities"

type Notification struct {
	PageID string         `json:"page_id,omitempty"`
	Error  string         `json:"error,omitempty"`
	Cause  string         `json:"cause,omitempty"`
	Page   *entities.Page `json:"page,inline,omitempty"`
}
