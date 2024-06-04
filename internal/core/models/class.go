package models

type Class struct {
	ID        uint   `json:"id"`
	ClassType string `json:"classType"`
	Duration  uint   `json:"duration"`
}
