package operativac2

import (
	"fmt"
	"net"
	"strings"
)

type udpServer struct {
	address *net.UDPAddr
	sluzba  *Sluzba
	nastavi bool
}

func NoviServer(port string, s *Sluzba) *udpServer {
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
		return nil
	}
	return &udpServer{
		address: udpAddr,
		sluzba:  s,
		nastavi: true,
	}
}
func (u *udpServer) Pokreni(port string) {

	// Create UDP listener
	conn, err := net.ListenUDP("udp", u.address)
	if err != nil {
		fmt.Printf("Error listening: %v\n", err)
		return
	}
	fmt.Println("Adresa udp servera : ", u.address)
	defer conn.Close()

	for {
		buffer := make([]byte, 1024)
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Printf("Error reading UDP message: %v\n", err)
			continue
		}
		msg := buffer[:n]

		fmt.Printf("Received message from %s: %s\n", clientAddr, msg)
		if !u.nastavi {
			break
		} else {
			continue
		}
	}
}

func (u *udpServer) Zaustavi() { u.nastavi = false }
