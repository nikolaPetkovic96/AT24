package main

import (
	op "at24/operativac"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("Hello world !")

	sl1 := op.NovaSluzba("domacaSluzba")

	id1 := sl1.DodajOperativca()
	id2 := sl1.DodajOperativca()
	sl1.PosaljiPoruku(id1, "Dobrodosao u sluzbu operativcu 0\n")
	sl1.PosaljiPoruku(id2, "Dobrodosao u sluzbu operativcu 1\n")
	sl1.PosaljiPoruku(id2, 2)

	sl1.ObustaviSluzbu(start)
	duration := time.Since(start)
	fmt.Println("DUZINA IZVRSAVANJA : ", duration)
	fmt.Println("\nSada su svi operativci zavrsili sve obaveze\n")

	sl1.UgasiSluzbu(start)

}
