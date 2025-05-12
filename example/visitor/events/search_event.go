package events

type SearchEvent struct {
}

func NewSearchEvent() *SearchEvent {
	return &SearchEvent{}
}

func (s *SearchEvent) Visit() error {
	s.VisitSearchTrigger()

	return nil
}

func (*SearchEvent) VisitSearchTrigger() {

}
