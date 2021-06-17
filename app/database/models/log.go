package models

import "time"

// LogModel struct
type LogModel struct {
	ID          int         `json:"id"`
	Modules     string      `json:"modules"`
	RefID       int         `json:"ref_id"`
	Actions     string      `json:"actions"`
	LogActivity interface{} `json:"log_activity"`
	CreatedAt   time.Time   `json:"created_at"`
	CreatedBy   string      `json:"created_by"`
}

// ListLogModel struct
type ListLogModel struct {
	Data  []LogModel `json:"data"`
	Total int        `json:"total"`
}
