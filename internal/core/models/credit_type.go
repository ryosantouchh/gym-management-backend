package models

type CreditType struct {
	ID            uint   `json:"id"`
	Type          string `json:"type"`
	IsMembership  bool   `json:"isMembership"`
	ClassDuration uint   `json:"classDuration"`
}
