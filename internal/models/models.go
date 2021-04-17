package models

type Task struct {
	Id          uint64 `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
}
