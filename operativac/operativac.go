package operativac

import (
	"fmt"
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
	sanduce     chan Message
	obustava    chan Message
	wg          sync.WaitGroup
	sluzba      *Sluzba
	obustavljen bool
}

func NoviOperativac(id int, sl *Sluzba) *Operativac {
	return &Operativac{
		pid:         id,
		sanduce:     make(chan Message, 10),
		obustava:    make(chan Message, 1),
		sluzba:      sl,
		obustavljen: false,
	}
}

// func(tip) nazivFunc(param[]) retTip
func (o *Operativac) Receive(msg Message) {
	switch m := msg.(type) {

	case string:
		fmt.Printf("Opertaivac sa id = %d je primio poruku : %s\n", o.pid, m)

	default:
		fmt.Printf("Opertaivac sa id = %d ne prepoznaje tip poruke\n", o.pid)
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

func (o *Operativac) Obradi(msg Message) bool {
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
