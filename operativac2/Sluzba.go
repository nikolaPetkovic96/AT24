package operativac2

import (
	"fmt"
	"net"
	"sync"
)

type Sluzba struct {
	operativci    map[string]*Operativac
	kinds         map[string]*Props
	mu            sync.Mutex
	wg            sync.WaitGroup
	conn          *net.UDPConn
	poznateSluzbe map[string]*net.UDPAddr
}

func NovaSluzba(conn *net.UDPConn) *Sluzba {
	return &Sluzba{
		operativci:    make(map[string]*Operativac),
		poznateSluzbe: make(map[string]*net.UDPAddr),
		conn:          conn,
		kinds:         make(map[string]*Props),
	}
}

func (sl *Sluzba) Spawn(props *Props, id string) {
	sl.mu.Lock()
	_, exProps := sl.kinds[props.naziv]
	if !exProps {
		sl.kinds[props.naziv] = props
	}
	sl.mu.Unlock()
	sl.mu.Lock()
	//defer sl.mu.Unlock()
	_, exists := sl.operativci[id]
	if !exists {
		actor := props.actorFunc()
		op := &Operativac{
			mailbox:     make(chan Envelope, props.mailboxSize),
			actor:       actor,
			stopSignal:  make(chan struct{}),
			sluzba:      sl,
			parent:      nil,
			obustavljen: false,
			penzionisan: false,
			info:        Info{nazivSluzbe: "", id: id, cntFail: 0, kind: props.naziv, parent: ""},
		}
		sl.operativci[id] = op
		sl.mu.Unlock()
		//go sl.AktivirajOperativca(op)
		op.Start()
	} else {
		sl.mu.Unlock()
		fmt.Println("Posotji operativac sa zadatim id")
	}
}

// func (sl *Sluzba) AktivirajOperativca(op *Operativac) {
// 	for {
// 		select {
// 		case msg := <-op.mailbox:
// 			op.actor.Receive(msg)
// 		case <-op.stopSignal:
// 			return
// 		}
// 	}
//}

func (sl *Sluzba) Send(id string, msg Envelope) {
	sl.mu.Lock()
	op, exists := sl.operativci[id]
	sl.mu.Unlock()
	if exists { //TODO razradi ako je puno sanduce
		op.mailbox <- msg
	} else {
		fmt.Printf("Ne postoji operativac sa zadatim id = %s , msg:[%v]\n", id, msg.Message)
	}
}

func (sl *Sluzba) Stop(id string) {
	sl.mu.Lock()
	op, exists := sl.operativci[id]
	sl.mu.Unlock()
	if exists && !op.obustavljen {
		//fmt.Printf("Stop signal : %s\n", id)
		fmt.Printf(" Poslata poruka za gasenje, %s\n", op.info.id)
		op.stopSignal <- struct{}{} // Signal actor to stop
		//delete(sl.operativci, id)
	}
}

func (sl *Sluzba) Remove(id string) {
	sl.mu.Lock()
	op, exists := sl.operativci[id]

	if exists && op.penzionisan {
		delete(sl.operativci, id)
	}
	sl.mu.Unlock()
	//sl.wg.Done()
}

func (sl *Sluzba) UgasiSluzbu() {
	sl.mu.Lock()
	for _, op := range sl.operativci {
		//sl.wg.Add(1)
		if op.info.parent == "" {
			go func() {
				//defer sl.wg.Done()
				sl.Stop(op.info.id)
			}()
		}
	}
	sl.mu.Unlock()
	sl.wg.Wait()
	//fmt.Printf("Sluzba < %s> je uspesno UGASENA\n", sl.naziv)
	fmt.Printf("Sluzba je uspesno UGASENA\n")
}

func (sl *Sluzba) dodajPoznateSluzbu(poznate []string) {
	for _, s := range poznate {
		addr, err := net.ResolveUDPAddr("udp", s)
		if err != nil {
			fmt.Printf("NEISPRAVAN FORMAT ADRESE : %s \n", s)
			continue
		}
		sl.poznateSluzbe[s] = addr
	}
}

func (sl *Sluzba) PosaljiDrugojSluzbi(adr string, s string) {
	sl.conn.WriteToUDP([]byte(s), sl.poznateSluzbe[adr])
}
