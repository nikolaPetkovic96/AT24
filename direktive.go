package main

import (
	"at24/poruke"
	"fmt"
	"math/rand"
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
		PosaljiDrugojSluzbi2(sluzba, &poruke.Poruka{Posiljalac: "STOP", CntFail: 0, Msg: &poruke.Poruka_Stop{Stop: &poruke.Stop{}}}, conn)
	case "force_fail":
		for i := 1; i <= 30; i++ {
			failure := randomBool()
			PosaljiDrugojSluzbi2(sluzba, &poruke.Poruka{Posiljalac: "FAIL", CntFail: 0, Msg: &poruke.Poruka_Fail{Fail: &poruke.Fail{Fail: failure}}}, conn)
		}

	case "mini_novi":
		for i := 1; i <= 30; i++ {
			sendByte("novi", sluzba, conn)
		}
	case "mini_stop":
		sendByte("stop", sluzba, conn)
	case "mini_check":
		sendByte("check", sluzba, conn)
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

func randomBool() bool {
	randomNumber := rand.Intn(2)
	return randomNumber == 1
}

func sendByte(msg string, adr string, conn *net.UDPConn) {
	marsh := []byte(msg)

	addr, err := net.ResolveUDPAddr("udp", adr)
	if err != nil {
		fmt.Printf("NEISPRAVAN FORMAT ADRESE : %s \n", adr)
		return
	}
	conn.WriteToUDP(marsh, addr)
}
