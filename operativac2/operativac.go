package operativac2

import (
	"fmt"
	"sync"
)

type Message interface{} //omogucava da bilo sta bude konkretna poruka
type Envelope struct {   //omotac oko poruke sa dodatnim informacijama
	Message
	senderId   string
	receiverId string
	cntFail    int
}

func NewEnvelope(msg interface{}, senderId string, receiverId string) *Envelope {
	return &Envelope{
		Message:    msg,
		senderId:   senderId,
		receiverId: receiverId,
		cntFail:    0,
	}
}

type Actor interface {
	//Receive(msg interface{})
	Receive(envelope Envelope)
}
type Info struct {
	nazivSluzbe string
	id          string
	kind        string
	cntFail     int
	parent      string //roditeljski actor unutar sistema
}

type Operativac struct {
	info        Info
	mailbox     chan Envelope
	actor       Actor
	stopSignal  chan struct{}
	sluzba      *Sluzba
	parent      *Operativac
	obustavljen bool
	penzionisan bool
	mu          sync.WaitGroup
}

func (o *Operativac) Start() {
	o.sluzba.wg.Add(1)
	go func() {
		defer o.sluzba.wg.Done()
		for {

			if o.obustavljen && len(o.mailbox) == 0 {
				//time.Sleep(10 * time.Second)
				fmt.Println("Operativac " + o.info.id + " uspesno PENIZIONISAN")
				o.penzionisan = true
				//defer o.sluzba.wg.Done()
				return
			}
			select {

			case dopis := <-o.mailbox:

				//fmt.Println("Operativac " + o.info.id + " primio poruku")

				o.actor.Receive(dopis)
			case <-o.stopSignal:
				if !o.obustavljen {
					o.obustavljen = true
					fmt.Println("Operativac " + o.info.id + " vise ne prima poruke")
					close(o.mailbox)
					o.obustavljen = true
				}
				//kad je stigao signal za obustavu i sanduce prazno operativac se moze ukloniti iz sluzbe
			}

		}
	}()
}

func (o *Operativac) SpawnChild(props *Props, id string) {
	o.sluzba.mu.Lock()
	_, exProps := o.sluzba.kinds[props.naziv]
	if !exProps {
		o.sluzba.kinds[props.naziv] = props
	}
	o.sluzba.mu.Unlock()
	o.sluzba.mu.Lock()
	//defer sl.mu.Unlock()
	_, exists := o.sluzba.operativci[id]
	if !exists {
		actor := props.actorFunc()
		op := &Operativac{
			mailbox:     make(chan Envelope, props.mailboxSize),
			actor:       actor,
			stopSignal:  make(chan struct{}),
			sluzba:      o.sluzba,
			parent:      o,
			obustavljen: false,
			penzionisan: false,
			info:        Info{nazivSluzbe: o.info.nazivSluzbe, id: id, cntFail: 0, kind: props.naziv, parent: o.info.id},
		}
		o.sluzba.operativci[id] = op
		o.sluzba.mu.Unlock()
		//go sl.AktivirajOperativca(op)
		op.Start()
	} else {
		o.sluzba.mu.Unlock()
		fmt.Println("Posotji operativac sa zadatim id")
	}
}
