package main

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`     // e.g., "pending", "in-progress", "completed"
	CreatedAt   string `json:"created_at"` // ISO 8601 format
	UpdatedAt   string `json:"updated_at"` // ISO 8601 format
}
