package main

import (
	op "at24/operativac"
	"at24/poruke"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"google.golang.org/protobuf/proto"
)

var address string

func main() {
	port := os.Args[1]
	poznateSluzbe := strings.Split(os.Args[2], ",")
	for _, s := range poznateSluzbe {
		fmt.Printf("PoznataSluzba : %s\n", s)
	}

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
	fmt.Println("Adresa udp servera : ", novaA)
	defer conn.Close()

	sl1 := op.NovaSluzba("domacaSluzba", conn, poznateSluzbe)

	sl1.DodajOperativca()
	sl1.DodajOperativca()
	sl1.DodajOperativca()
	sl1.DodajOperativca()
	sl1.DodajOperativca()

	miniSys := op.NoviSistem(conn)
	miniSys.Start()
	//sl1.PosaljiPoruku(id1, "Dobrodosao u sluzbu operativcu 0\n")
	//sl1.PosaljiPoruku(id2, "Dobrodosao u sluzbu operativcu 1\n")
	sl1.PingAll()
	for {
		buffer := make([]byte, 1024)
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Printf("Error reading UDP message: %v\n", err)
			continue
		}
		msg := buffer[:n]
		var poruka poruke.Poruka
		err = proto.Unmarshal(msg, &poruka)
		if err != nil {
			//fmt.Println("Error unmarshaling protobuf:", err)
			miniSys.Receive(string(msg))
			continue
		}

		fmt.Printf("Received message from %s: %s\n", clientAddr, poruka.String())
		nastavi := sl1.PosaljiPorukuRand(poruka)
		if !nastavi {
			break
		} else {
			continue
		}
	}
	sl1.UgasiSluzbu(time.Now())
}

func NoviSistem(conn *net.UDPConn) {
	panic("unimplemented")
}
