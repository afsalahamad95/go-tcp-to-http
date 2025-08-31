package main

import (
	"bytes"
	"io"
	"log"
	"net"
)

// this function hypothetically accepts a TCP connection, and we're gonna read a stream of messages.
// Note: io.ReadCloser is an interface with two methods - Read and Close.
// So anything that implements that can be used here, we'll be using files here.
func getLinesFromChannel(f io.ReadCloser) <-chan string {
	channel := make(chan string, 1)
	go func() {
		defer f.Close()
		defer close(channel)

		str := ""
		for {
			buffer := make([]byte, 8)
			offset, err := f.Read(buffer)
			if err != nil {
				log.Println("There was an error", err)
				break
			}
			// we are looking for '/n' as the delimiter. If we see it we print the current sentence and print the next in the
			// next line.

			// note: this is a mini state machine below
			if i := bytes.IndexByte(buffer, '\n'); i != -1 {
				str += string(buffer[:i])
				channel <- str
				buffer = buffer[i+1:]     // this goes to the next line
				log.Println("Read:", str) // print the current string in the current line
				str = ""                  // reset the string so that we can print the next line
			}
			log.Printf("Read %v bytes", offset)
			str += string(buffer) // so if we had a previous string with '\n', it would've been printed already and cleared
			// from the buffer. Now we just print the string, which will be the current string.
		}
	}()
	return channel
}

func main() {
	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Println("there was an error", err)
		return
	}
	defer listener.Close()
	for {
		connection, err := listener.Accept() // this is a TCP connection now
		if err != nil {
			log.Println("there was an error with the TCP connection", err)
		}
		log.Println("Connection accepted.")
		lines := getLinesFromChannel(connection)
		log.Println("The result:")
		// Note: The for-range block below will block till we get a value from the channel or till it is closed.
		for line := range lines {
			log.Println("received from chan:", line)
		}
	}
}
