package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/mathhaug/is105sem03_REP03/mycrypt"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: tcp_client <message>")
		os.Exit(1)
	}

	conn, err := net.Dial("tcp", "172.17.0.2:5002")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Din melding: ", os.Args[1])

	kryptertMelding := mycrypt.Krypter([]rune(os.Args[1]), mycrypt.ALF_SEM03, 4)
	log.Println("Kryptert sendt melding: ", string(kryptertMelding))
	_, err = conn.Write([]byte(string(kryptertMelding)))
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	response := string(buf[:n])
	switch response {
	case "":
		log.Println("response from proxy: ", response)
	default:
		response := mycrypt.Krypter(([]rune(string(buf[:n]))), mycrypt.ALF_SEM03, len(mycrypt.ALF_SEM03)-4)
		log.Println("response from proxy:  ", string(response))
	}
}

