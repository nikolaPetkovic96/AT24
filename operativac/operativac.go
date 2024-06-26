package operativac

import "fmt"

type Message interface{}

type Actor interface {
	Receive(msg Message)
}

type Operativac struct {
	pid int
}

func (o *Operativac) Receive(msg Message) {
	switch m := msg.(type) {
	case string:
		fmt.Printf("Opertaivac sa id = %d je primio poruku : %s\n", o.pid, m)
	default:
		fmt.Printf("Opertaivac sa id = %d ne prepoznaje tip poruke\n", o.pid)
	}

}
