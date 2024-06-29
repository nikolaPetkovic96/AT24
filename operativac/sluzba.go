package operativac

import (
	"fmt"
	"sync"
	"time"
)

type Sluzba struct {
	naziv      string
	operativci map[int]*Operativac
	naredniId  int
	mu         sync.Mutex
	wg         sync.WaitGroup
}

func NovaSluzba(ime string) *Sluzba {
	return &Sluzba{
		naziv:      ime,
		operativci: make(map[int]*Operativac),
		naredniId:  0,
	}
}

func (sl *Sluzba) DodajOperativca() int {
	id := sl.naredniId

	op := NoviOperativac(id, sl)
	sl.naredniId += 1
	sl.mu.Lock()
	sl.operativci[id] = op
	sl.mu.Unlock()
	op.Start()
	//fmt.Printf("id: %d\n", id)
	return id
}

func (sl *Sluzba) ProveriOperativca(id int) {
	sl.mu.Lock()
	op, exists := sl.operativci[id]
	sl.mu.Unlock()
	if exists {
		op.Receive("Provera")
	} else {
		fmt.Printf("U Sluzbi %s ne postoji operativac sa Id : %d .", sl.naziv, id)
	}
}

func (sl *Sluzba) PosaljiPoruku(id int, msg Message) {
	t := time.Now()
	sl.mu.Lock()
	op, exists := sl.operativci[id]
	sl.mu.Unlock()
	if exists {
		uspeh := op.Obradi(msg)
		switch uspeh {
		case true:
			fmt.Printf("%s : Poruka uspesno poslata operativcu : %d\n", t, id)
		default:
			fmt.Printf("%s : Poruka nije isporucena operativcu : %d\n", t, id)
		}
	} else {
		fmt.Printf("%s : U Sluzbi %s ne postoji operativac sa Id : %d\n", t, sl.naziv, id)
	}
}

func (sl *Sluzba) ObustaviSluzbu(t time.Time) {
	sl.mu.Lock()
	for _, op := range sl.operativci {
		go func() { op.Stop(sl) }()
		fmt.Printf("%s : Poslata poruka za obustavljanje, %d\n", t, op.pid)
	}
	sl.mu.Unlock()
	time.Sleep(500 * time.Millisecond)
	sl.PosaljiPoruku(1, "NE BI TREBALO DA SE PRIKAZE U KONZOLI!")
	sl.wg.Wait()

	fmt.Printf("%s : Sluzba < %s > je obustavljena\n", t, sl.naziv)
}

// gasi sluzbu koja je prethodno otkazana
func (sl *Sluzba) UgasiSluzbu(t time.Time) {

	for _, op := range sl.operativci {
		sl.wg.Add(1)
		go func() {
			defer sl.wg.Done()
			sl.UkloniOperativca(op.pid)
		}()
		fmt.Printf("%s : Poslata poruka za obustavljanje, %d\n", t, op.pid)
	}
	sl.wg.Wait()
	fmt.Printf("Sluzba < %s> je uspesno UGASENA\n", sl.naziv)
}

// uklanja operativca koji je vec dobio signal za obustavu i ispraznio sanduce
func (sl *Sluzba) UkloniOperativca(i int) {
	sl.mu.Lock()
	o, exists := sl.operativci[i]
	if exists && o.isPenzionisan() {
		delete(sl.operativci, i)
		fmt.Printf("Uspesno uklonjen operativac sa id : %d\n", i)
	}
	sl.mu.Unlock()
}
