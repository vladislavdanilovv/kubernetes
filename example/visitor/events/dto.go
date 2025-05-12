package events

import "time"

type EventDto struct {
	Event      string    `json:"event"`
	FromCityID int       `json:"from_city_id"`
	ToCityID   int       `json:"to_city_id"`
	Date       time.Time `json:"date"`
	UserID     int       `json:"user_id"`
	OwnUserID  int
	OrderID    int `json:"order_id"`
	OfferID    int `json:"offer_id"`
}
