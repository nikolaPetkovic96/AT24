package operativac

// import (
// 	"fmt"
// 	"math/rand"
// 	"net"
// 	"sync"
// 	"time"
// )

// type OperativacMini struct {
// 	naziv       string
// 	id          int
// 	otac        *OperativacMini
// 	deca        map[int]*OperativacMini
// 	naredniId   int
// 	mu          sync.Mutex
// 	wg          sync.WaitGroup
// 	conn        *net.UDPConn
// 	sanduce     chan Message
// 	obustava    chan Message
// 	cntMsg      int
// 	cntErr      int
// 	obustavljen bool
// }

// func (o *OperativacMini) dodajMini() {
// 	idCur := o.naredniId
// 	novi := NoviMini(idCur, o)
// 	o.naredniId += 1
// 	o.mu.Lock()
// 	o.deca[idCur] = novi
// 	o.mu.Unlock()
// 	novi.Start()
// 	fmt.Printf("Rod: <%s> je dobio dete <%s>\n", o.naziv, o.deca[idCur].naziv)

// }
// func NoviMini(id1 int, rod *OperativacMini) *OperativacMini {
// 	naziv := rod.naziv + "_"
// 	nazivConcat := fmt.Sprint(naziv, id1)
// 	return &OperativacMini{
// 		naziv:       nazivConcat,
// 		id:          id1,
// 		otac:        rod,
// 		deca:        make(map[int]*OperativacMini),
// 		naredniId:   1,
// 		conn:        rod.conn,
// 		sanduce:     make(chan Message, 10),
// 		obustava:    make(chan Message, 1),
// 		cntMsg:      0,
// 		cntErr:      0,
// 		obustavljen: false,
// 	}
// }

// func NoviSistem(conn1 *net.UDPConn) *OperativacMini {
// 	return &OperativacMini{
// 		naziv:       "root",
// 		id:          0,
// 		otac:        nil,
// 		deca:        make(map[int]*OperativacMini),
// 		naredniId:   1,
// 		conn:        conn1,
// 		sanduce:     make(chan Message, 100),
// 		obustava:    make(chan Message, 1),
// 		cntMsg:      0,
// 		cntErr:      0,
// 		obustavljen: false,
// 	}
// }

// func (o *OperativacMini) Start() {
// 	if o.otac != nil {
// 		o.otac.wg.Add(1)
// 	}
// 	//defer o.otac.wg.Done()
// 	if o.obustavljen {
// 		return
// 	}
// 	go func() {
// 		for {
// 			select {
// 			case dopis := <-o.sanduce:
// 				o.Receive(dopis)

// 			case <-o.obustava:
// 				o.obustavljen = true
// 				fmt.Printf("Operativac <%s> vise ne prima poruke\n", o.naziv)
// 			}
// 			if o.obustavljen {
// 				//time.Sleep(10 * time.Second)

// 				fmt.Printf("Operativac <%s> je uspesno PENZIONISAN\n", o.naziv)
// 				if o.otac != nil {
// 					o.otac.wg.Done()
// 				}
// 				return
// 			}
// 		}
// 	}()
// }

// func (o *OperativacMini) Receive(msg Message) {
// 	o.cntMsg += 1
// 	rand := randomBool(1, 3)
// 	switch m := msg.(type) {
// 	case string:
// 		if m == "novi" && rand {
// 			o.dodajMini()
// 			go o.sendToRandom("novi")
// 		} else if m == "check" {
// 			fmt.Printf("<%s> Number of Children : %d\n", o.naziv, len(o.deca))
// 			o.mu.Lock()

// 			for _, dete := range o.deca {
// 				go func() { dete.sanduce <- "check" }()
// 			}
// 			o.mu.Unlock()

// 		} else if m == "stop" {
// 			o.mu.Lock()
// 			for _, dete := range o.deca {
// 				go func() { dete.sanduce <- "stop" }()
// 				fmt.Printf("<%s> naredio obustavljenje za %s\n", o.naziv, dete.naziv)
// 			}
// 			o.mu.Unlock()
// 			o.wg.Wait()
// 			fmt.Printf("<%s> sva deca zaustavljena !\n", o.naziv)
// 			//todo pojedinacno brisanje dece kroz go rutine, kao u izbacivanju agenata iz sluzbe
// 			o.deca = make(map[int]*OperativacMini)

// 			//close(o.sanduce)
// 			time.Sleep(2 * time.Second)
// 			if o.otac != nil {
// 				defer o.otac.wg.Done()
// 			}
// 			return
// 		}
// 	default:
// 		fmt.Printf("<%s> ne prepoznaje poruku\n", o.naziv)
// 	}
// }

// func (o *OperativacMini) sendToRandom(msg Message) {
// 	o.mu.Lock()
// 	id := rand.Intn(len(o.deca))
// 	op, exists := o.deca[id]
// 	o.mu.Unlock()
// 	if exists {
// 		op.sanduce <- msg
// 	}
// }
