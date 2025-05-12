package events

type AcceptRequestEvent struct {
}

func NewAcceptRequestEvent() *AcceptRequestEvent {
	return &AcceptRequestEvent{}
}

func (s *AcceptRequestEvent) Visit() error {
	s.VisitBookingTrigger()
	s.VisitSearchTrigger()

	return nil
}

func (*AcceptRequestEvent) VisitBookingTrigger() {

}

func (*AcceptRequestEvent) VisitSearchTrigger() {

}
