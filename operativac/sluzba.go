package operativac

import (
	"fmt"
	"sync"
)

type Sluzba struct {
	naziv      string
	operativci map[int]Actor
	sanducad   map[int]chan Message
	wg         sync.WaitGroup
	naredniId  int
}

func NovaSluzba(ime string) *Sluzba {
	return &Sluzba{
		naziv:      ime,
		operativci: map[int]Actor{},
		sanducad:   make(map[int]chan Message),
		wg:         sync.WaitGroup{},
		naredniId:  0,
	}
}

func (sl *Sluzba) DodajOperativca(op Actor) int {
	id := sl.naredniId
	sl.naredniId += 1
	op.(*Operativac).pid = id

	sl.operativci[id] = op
	sl.sanducad[id] = make(chan Message, 10)
	sl.wg.Add(1)
	go sl.ProveriOperativca(id)

	return id
}

func (sl *Sluzba) ProveriOperativca(id int) {
	defer sl.wg.Done()
	sanduce := sl.sanducad[id]
	for msg := range sanduce {
		sl.operativci[id].Receive(msg)
	}
}

func (sl *Sluzba) PosaljiPoruku(id int, msg Message) {
	sanduce, exists := sl.sanducad[id]
	if exists {
		sanduce <- msg
	} else {
		fmt.Printf("U Sluzbi %s ne postoji operativac sa Id : %d .", sl.naziv, id)
	}
}

func (sl *Sluzba) ObustaviSluzbu() {
	for _, sanduce := range sl.sanducad {
		close(sanduce)
	}
	sl.wg.Wait()
	fmt.Printf("Sluzba < %s > je obustavljena", sl.naziv)
}
