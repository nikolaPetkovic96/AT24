package main

import (
	"at24/poruke"
	"fmt"
	"net"
	"os"

	"google.golang.org/protobuf/proto"
)

func main() {
	direktiva := os.Args[1]
	sluzba := os.Args[2]

	//pokretanje servera
	connIP, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		fmt.Printf("err")
	}
	defer connIP.Close()
	address := connIP.LocalAddr().(*net.UDPAddr)

	// Create UDP listener
	conn, err := net.ListenUDP("udp", address)
	if err != nil {
		fmt.Printf("Error listening: %v\n", err)
		return
	}
	fmt.Println("Adresa udp servera : ", address.String())

	//prepoznavanje direktive
	switch direktiva {
	case "stop":
		PosaljiDrugojSluzbi2(sluzba, &poruke.Poruka{Posiljalac: "STOP", Msg: &poruke.Poruka_Stop{Stop: &poruke.Stop{}}}, conn)
	default:
		fmt.Println("Direktiva nije prepoznata!")

	}

}

func PosaljiDrugojSluzbi2(adr string, p *poruke.Poruka, conn *net.UDPConn) {
	marsh, err := proto.Marshal(p)
	if err != nil {
		fmt.Println("Error marshaling protobuf:", err)
		return
	}
	addr, err := net.ResolveUDPAddr("udp", adr)
	if err != nil {
		fmt.Printf("NEISPRAVAN FORMAT ADRESE : %s \n", adr)
		return
	}
	conn.WriteToUDP(marsh, addr)
}
