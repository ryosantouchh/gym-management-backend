package models

type Class struct {
	ID        uint   `json:"id"`
	ClassType string `json:"classType"`
	Duration  uint   `json:"duration"`
}

type UpdateClassRequest struct {
	ClassType string `json:"classType,omitempty"`
	Duration  uint   `json:"duration,omitempty"`
}
