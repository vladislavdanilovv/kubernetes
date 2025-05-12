package events

type SendRequestEvent struct {
	Event      string `json:"event"`
	FromCityID int    `json:"from_city_id"`
	ToCityID   int    `json:"to_city_id"`
}

func NewSendRequestEvent(dto EventDto) *SendRequestEvent {
	return &SendRequestEvent{
		Event:      dto.Event,
		FromCityID: dto.ToCityID,
		ToCityID:   dto.ToCityID,
	}
}

func (s *SendRequestEvent) Visit() error {
	s.VisitBookingTrigger()
	s.VisitSearchTrigger()
	return nil
}

func (s *SendRequestEvent) VisitBookingTrigger() {

}

func (s *SendRequestEvent) VisitSearchTrigger() {

}
