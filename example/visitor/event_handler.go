package visitor

import "gis/example/visitor/events"

type EventHandler struct {
	Events map[string]events.IVisitor
}

func NewEventHandler(dto events.EventDto) *EventHandler {

	var dict = make(map[string]events.IVisitor)
	dict["send_request"] = events.NewSendRequestEvent(dto)
	dict["accept_request"] = events.NewAcceptRequestEvent()
	dict["search"] = events.NewSearchEvent()

	return &EventHandler{
		Events: dict,
	}
}

func (s *EventHandler) Handle(event string) error {
	err := s.Events[event].Visit()
	if err != nil {
		return err
	}

	return nil
}
