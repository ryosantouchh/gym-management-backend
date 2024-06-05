package models

type Credit struct {
	ID             uint   `json:"id"`
	Type           string `json:"type"`
	IsMembership   bool   `json:"isMembership"`
	ExpireDuration uint   `json:"classDuration"`
	Price          int    `json:"price"`
	SpecialPrice   int    `json:"specialPrice"`
}
