package main

import (
	"io"
	"log"
	"os"
)

// this function hypothetically accepts a TCP connection, and we're gonna read a stream of messages.
// Note: io.ReadCloser is an interface with two methods - Read and Close.
// So anything that implements that can be used here, we'll be using files here.
func getLinesFromChannel(f io.ReadCloser) <-chan string {
	ch := make(chan string, 1)
	go func() {
		data := make([]byte, 1)
		offset, err := f.Read(data)
		if err != nil {
			log.Println("There was an error", err)
			return
		}
		log.Printf("Read %v bytes", offset)
		ch <- string(data)

	}()
	return ch
}

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal("there was an error", err)
		return
	}
	defer file.Close()

	for {
		dataChannel := getLinesFromChannel(file)
		log.Println("Received:", <-dataChannel)
	}
}
