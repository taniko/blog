package stack

type Event interface {
	internal()
}

type unimplemented struct{}

func (unimplemented) internal() {}
