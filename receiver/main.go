package main

import (
	"flag"
	"log"
	"net"
)

func main() {
	source := flag.String("source", "localhost:2003", "Source for the UDP Graphite Data")
	flag.Parse()

	log.Printf("Listening for Graphite Data on %s", *source)
	s, err := net.ResolveUDPAddr("udp4", *source)
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.ListenUDP("udp4", s)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		n, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Printf("ERROR: %s", err)
			continue
		}
		log.Printf("Received %d bytes from %s: %s", n, addr, string(buf))
	}
}
