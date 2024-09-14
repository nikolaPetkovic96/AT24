package main

import (
	op "at24/operativac2"
	"fmt"
	"time"
)

type PrinterActor struct{}

func (p *PrinterActor) Receive(envelope op.Envelope) {
	switch msg := envelope.Message.(type) {
	case string:
		fmt.Println("PrinterActor received:", msg)
	default:
		fmt.Println("PrinterActor received unknown message")
	}
}

type WorkerActor struct {
	brojac int
}

func (w *WorkerActor) Receive(envelope op.Envelope) {
	w.brojac++
	switch msg := envelope.Message.(type) {
	case string:
		fmt.Println("WorkerActor processed:", msg+"!")
	default:
		fmt.Printf("WorkerActor received unknown message. BROJAC: %d\n", w.brojac)
	}
}

func main() {
	printerProps := op.NewProps("print", func() op.Actor {
		return &PrinterActor{}
	})
	workerProps := op.NewProps("work", func() op.Actor {
		return &WorkerActor{brojac: 0}
	})

	sl := op.NovaSluzba(nil)
	sl.Spawn(printerProps, "p")
	sl.Spawn(workerProps, "w")

	sl.Send("p", *op.NewEnvelope("HELLO P", "unknown", "p"))
	sl.Send("w", *op.NewEnvelope("HELLO W", "unknown", "W"))

	sl.Stop("p")

	sl.Send("p", *op.NewEnvelope("35", "unknown", "p"))
	sl.Send("w", *op.NewEnvelope(23, "unknown", "2"))
	sl.UgasiSluzbu()
	time.Sleep(1 * time.Second)
}
