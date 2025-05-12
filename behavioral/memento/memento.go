package memento

import "fmt"

type Originator struct {
	State string
}

// CreateMemento returns state storage.
func (o *Originator) CreateMemento() *Memento {
	return &Memento{state: o.State}
}

// SetMemento sets old state.
func (o *Originator) SetMemento(memento *Memento) {
	o.State = memento.GetState()
}

// Memento implements storage for the state of Originator
type Memento struct {
	state string
}

// GetState returns state.
func (m *Memento) GetState() string {
	return m.state
}

// Caretaker keeps Memento until it is needed by Originator.
type Caretaker struct {
	Memento *Memento
}

func Handler() {
	//memento := &Caretaker{
	//	Memento: &Memento{
	//		state: "qwe",
	//	},
	//}
	//
	//orig := Originator{}
	////fmt.Println(memento.CreateMemento())
	//orig.CreateMemento()
	//orig.SetMemento(memento.Memento)
	//fmt.Println(memento.Memento.GetState())
	//orig := &Originator{}
	//
	//memento := Caretaker{
	//	Memento: &Memento{
	//		state: orig.State,
	//	},
	//}
	//
	//orig.SetMemento(memento.Memento)
	//orig.CreateMemento()
	//
	//fmt.Println(memento.Memento.GetState())

	originator := new(Originator)
	caretaker := new(Caretaker)

	originator.State = "1"

	caretaker.Memento = originator.CreateMemento()

	fmt.Println(originator.State)

	originator.State = "2"

	//caretaker.Memento = originator.CreateMemento()

	originator.SetMemento(caretaker.Memento)

	fmt.Println(originator.State)

	originator.State = "3"

	//caretaker.Memento = originator.CreateMemento()

	originator.SetMemento(caretaker.Memento)

	fmt.Println(originator.State)

	fmt.Println()

	if originator.State != "1" {
		fmt.Println(fmt.Errorf("Expect State to %s, but %s", originator.State, "1"))
	}
}
