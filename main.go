package main

import (
	op "at24/operativac"
	"fmt"
)

func main() {
	fmt.Println("Hello world !")

	sl1 := op.NovaSluzba("domacaSluzba")

	op1 := &op.Operativac{}
	op2 := &op.Operativac{}

	id1 := sl1.DodajOperativca(op1)
	id2 := sl1.DodajOperativca(op2)

	sl1.PosaljiPoruku(id1, "Dobrodosao u sluzbu operativcu 0\n")
	sl1.PosaljiPoruku(id2, "Dobrodosao u sluzbu operativcu 1\n")
	sl1.PosaljiPoruku(id2, 2)

	sl1.ObustaviSluzbu()
}
