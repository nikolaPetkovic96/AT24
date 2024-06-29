package main

import (
	op "at24/operativac"
	"fmt"
	"net"
	"os"
	"strings"
)

var address string

func main() {
	port := os.Args[1]
	connIP, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		fmt.Printf("err")
	}
	defer connIP.Close()
	address := connIP.LocalAddr().(*net.UDPAddr)
	add := strings.Split(address.String(), ":")
	novaA := add[0]

	novaA += ":"
	novaA += port
	udpAddr, err := net.ResolveUDPAddr("udp", novaA)
	if err != nil {
		fmt.Printf("Error resolving address: %v\n", err)
		return
	}

	// Create UDP listener
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Printf("Error listening: %v\n", err)
		return
	}
	fmt.Println("Adresa udp servera : ", address.String())
	defer conn.Close()

	sl1 := op.NovaSluzba("domacaSluzba")

	id1 := sl1.DodajOperativca()
	id2 := sl1.DodajOperativca()
	sl1.PosaljiPoruku(id1, "Dobrodosao u sluzbu operativcu 0\n")
	sl1.PosaljiPoruku(id2, "Dobrodosao u sluzbu operativcu 1\n")
	for {
		buffer := make([]byte, 1024)
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Printf("Error reading UDP message: %v\n", err)
			continue
		}
		msg := buffer[:n]

		fmt.Printf("Received message from %s: %s\n", clientAddr, string(msg))
		sl1.PosaljiPoruku(len(msg)%2, msg)
	}
}
