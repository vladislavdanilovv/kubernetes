package events

type IVisitor interface {
	Visit() error
}
