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
	children    []string
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
	mu          sync.Mutex
	wg          sync.WaitGroup
}

func (o *Operativac) Start() {
	fmt.Printf("Startovanje aktora sa id= %s\n", o.info.id)
	if o.info.parent != "" {
		o.parent.mu.Lock()
		o.parent.info.children = append(o.parent.info.children, o.info.id)
		o.parent.mu.Unlock()
		o.parent.wg.Add(1)
	} else {
		o.sluzba.wg.Add(1)
	}

	go func() {
		//defer o.sluzba.wg.Done()
		defer o.oslobodi()
		for {

			if o.obustavljen && len(o.mailbox) == 0 {
				//time.Sleep(10 * time.Second)
				o.mu.Lock()
				for _, s1 := range o.info.children {
					fmt.Printf("Zaustavi child : %s\n", s1)
					go func() {
						o.sluzba.Stop(s1)
					}()
				}
				o.mu.Unlock()
				o.wg.Wait()
				//time.Sleep(5 * time.Second)
				fmt.Println("Operativac " + o.info.id + " uspesno PENIZIONISAN")
				o.penzionisan = true
				//defer o.sluzba.wg.Done()
				break
			}
			select {

			case dopis := <-o.mailbox:

				//fmt.Println("Operativac " + o.info.id + " primio poruku")

				o.actor.Receive(dopis)
			case <-o.stopSignal:
				//time.Sleep(1 * time.Second)
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

func (o *Operativac) SpawnChild(props *Props, id string) string {
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
			info:        Info{nazivSluzbe: o.info.nazivSluzbe, id: o.info.id + "_" + id, cntFail: 0, kind: props.naziv, parent: o.info.id},
		}
		o.sluzba.operativci[op.info.id] = op
		o.sluzba.mu.Unlock()
		fmt.Printf("SPAWN child :  %s\n", op.info.id)

		//go sl.AktivirajOperativca(op)
		op.Start()
		return op.info.id
	} else {
		o.sluzba.mu.Unlock()
		fmt.Println("Posotji operativac sa zadatim id")
		return ""
	}
}

func (o *Operativac) oslobodi() {
	if o.info.parent != "" {
		o.parent.mu.Lock()
		idx := -1
		for index, s := range o.parent.info.children {
			if s == o.info.id {
				idx = index
				break
			}

		}
		if idx > -1 && len(o.parent.info.children) > 1 {
			o.parent.info.children = append(o.parent.info.children[:idx], o.parent.info.children[idx+1])
		} else if idx > -1 {
			o.parent.info.children = o.parent.info.children[:0]
		}
		o.parent.mu.Unlock()
		o.parent.wg.Done()
	} else {
		o.sluzba.mu.Lock()
		delete(o.sluzba.operativci, o.info.id)
		o.sluzba.mu.Unlock()
		o.sluzba.wg.Done()
	}
}

func (o *Operativac) SendToChildren(msg Message) bool {
	//env := &Envelope{senderId: o.info.id, Message: msg, receiverId: ""}
	o.mu.Lock()
	for _, childId := range o.info.children {
		//env.receiverId = childId
		o.sluzba.Send(childId, Envelope{senderId: o.info.id, Message: msg, receiverId: childId})
	}
	o.mu.Unlock()
	return true
}
