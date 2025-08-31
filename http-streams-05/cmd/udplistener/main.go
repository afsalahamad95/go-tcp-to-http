package main

import (
	"bufio"
	"log"
	"net"
	"os"
)

func main() {
	address, err := net.ResolveUDPAddr("udp", "localhost:42069")
	if err != nil {
		log.Println("there was an error", err)
		return
	}
	// laddr here -> local address - the one you bind to
	// raddr here -> remote address - where the packets go to
	// note: to send UDP packets to remote address, we don't need to bind to it. The above is a source(laddr)-destination(raddr) relation.
	// If you bind to a port X(laddr) and want to send your packets to port X(raddr) where laddr == raddr,
	// then you are listening to the UDP packets using nc -u -l X, your go program will not be able to bind to the port X
	// since nc will have already occupied it. So even if you did connect, nc will not receive anything, since you're sending the
	// packets to yourself.
	// The port mapping should be different so that the packets are actually sent.
	// So leaving the laddr as nil will allow the OS to pick a free random port automatically

	connection, err := net.DialUDP("udp", nil, address)
	if err != nil {
		log.Println("an error occurred", err)
		return
	}
	defer connection.Close()
	reader := bufio.NewReader(os.Stdin)
	for {
		log.Println(">")
		input, err := reader.ReadString('.')
		if err != nil {
			log.Println("there was an error", err)
			return
		}
		_, err = connection.Write([]byte(input))
		if err != nil {
			log.Println("there was an error", err)
			return
		}
	}
}
