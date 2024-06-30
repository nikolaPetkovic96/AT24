package operativac

import (
	poruke "at24/poruke"
	"fmt"
	"math/rand"
	"reflect"
	"sync"
	"time"
)

type Message interface{}

type Actor interface {
	Receive(msg Message)
	Start()
	Stop()
}

type Operativac struct {
	pid         int
	sanduce     chan poruke.Poruka
	obustava    chan Message
	wg          sync.WaitGroup
	sluzba      *Sluzba
	obustavljen bool
	cntReceive  int
	cntFailure  int
}

func NoviOperativac(id int, sl *Sluzba) *Operativac {
	return &Operativac{
		pid:         id,
		sanduce:     make(chan poruke.Poruka, 10),
		obustava:    make(chan Message, 1),
		sluzba:      sl,
		obustavljen: false,
		cntReceive:  0,
		cntFailure:  0,
	}
}

// func(tip) nazivFunc(param[]) retTip
func (o *Operativac) Receive(msg poruke.Poruka) {
	lok := o.sluzba.conn.LocalAddr().String()
	o.cntReceive += 1
	switch m := msg.Msg.(type) {
	//case string:
	//	fmt.Printf("Opertaivac sa id = %d je primio poruku : %s\n", o.pid, m)
	//case *poruke.Poruka:
	//	fmt.Println("PRIMLJEN PING \n")

	//_, err := conn.writeToUDP(&poruke.Pong{Id: "null"})
	//if err != nil {
	//fmt.Printf("Error sending PONG to  %s\n", )
	//}
	//case *poruke.Poruka:
	//	fmt.Println("PRIMLJEN PING :%s\n", m.Posiljalac)
	case *poruke.Poruka_Ping:
		fmt.Printf("PRIMLJEN PING OD:%s\n", msg.Posiljalac)
		o.sluzba.PosaljiDrugojSluzbi2(msg.Posiljalac, &poruke.Poruka{Posiljalac: lok, Msg: &poruke.Poruka_Pong{Pong: &poruke.Pong{Id: "TEST_PONG"}}})
	case *poruke.Poruka_Pong:
		fmt.Printf("PRIMLJEN PONG OD:%s\n", msg.Posiljalac)
		o.sluzba.PosaljiDrugojSluzbi2(msg.Posiljalac, &poruke.Poruka{Posiljalac: lok, Msg: &poruke.Poruka_Ping{Ping: &poruke.Ping{Id: "TEST_PING"}}})
	case *poruke.Poruka_Fail:
		fail := randomBool(0, 3)
		if !fail {
			o.cntFailure += 1
			msg.CntFail += 1
			fmt.Printf("NEUSPELA OBRADA, POKUSAJI %d, ", msg.GetCntFail())
			fmt.Printf("BROJ NEUSPEHA AGENTA  %d : %d\n ", o.pid, o.cntFailure)
			o.ProveriSe()
			go o.sluzba.PosaljiPorukuRand(msg)

		} else {
			fmt.Printf("FAIL PORUKA USPESNO OBRADJENA, BROJ POKUSAJA : %d\n", msg.GetCntFail())

		}
	default:
		fmt.Printf("Opertaivac sa id = %d ne prepoznaje tip poruke\n", o.pid)
		fmt.Printf("NEPREPOZNATA PORUKA : %s\n", m)

		fmt.Printf("TIP:%s\n", reflect.TypeOf(msg))
		o.cntFailure += 1
		//o.sluzba.PosaljiDrugojSluzbi("192.168.1.102:9091")
	}

}

func (o *Operativac) Start() {
	o.sluzba.wg.Add(1)
	go func() {
		//defer o.wg.Done()
		for {
			select {
			case dopis := <-o.sanduce:
				if o.pid == 0 {
					time.Sleep(1 * time.Second)
				}
				o.Receive(dopis)
			case <-o.obustava:
				o.obustavljen = true
				fmt.Printf("Operativac %d vise ne prima poruke\n", o.pid)

			} //kad je stigao signal za obustavu i sanduce prazno operativac se moze ukloniti iz sluzbe
			if o.obustavljen && len(o.sanduce) == 0 {
				time.Sleep(10 * time.Second)
				fmt.Printf("Operativac %d je uspesno PENZIONISAN\n", o.pid)
				defer o.sluzba.wg.Done()
				return
			}
		}
	}()
}

// https://go101.org/article/channel-closing.html
func (o *Operativac) Stop(sl *Sluzba) {
	//fmt.Printf("Operativac %d je primio naredbu o zaustavljanju\n", o.pid)

	close(o.obustava)
}

func (o *Operativac) Obradi(msg poruke.Poruka) bool {
	if o.obustavljen || o.sanducePuno() {
		return false
	} else {
		o.sanduce <- msg
		return true
	}
}

// provera da li je kanal zatvoren
// https://gist.github.com/iamatypeofwalrus/84b6c7d946a6a4143a1d
func (o *Operativac) isPenzionisan() bool {
	if o.obustavljen && len(o.sanduce) == 0 {
		return true
	} else {
		return false
	}
}

// https://www.golinuxcloud.com/check-if-golang-channel-buffer-is-full/
func (o *Operativac) sanducePuno() bool {
	return len(o.sanduce) == cap(o.sanduce)
}

func randomBool(min int, max int) bool {
	randomNumber := rand.Intn(max - min)
	return randomNumber == 1
}

func (o *Operativac) ProveriSe() {
	neuspesi := o.cntFailure
	if neuspesi >= 5 {
		go o.sluzba.RestartujeOp(o.pid)
	}
}

func KopirajSe(op *Operativac) *Operativac {
	close(op.sanduce)
	kopy := &Operativac{
		pid:         op.pid,
		sanduce:     make(chan poruke.Poruka, 10),
		obustava:    make(chan Message, 1),
		sluzba:      op.sluzba,
		obustavljen: false,
		cntReceive:  op.cntReceive,
		cntFailure:  0,
	}
	//defer close(op.sanduce)

	return kopy
}
