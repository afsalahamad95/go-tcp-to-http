package main

import (
	"log"
	"os"
)

func main() {
	content, err := os.ReadFile("messages.txt")
	if err != nil {
		log.Println("error reading the file", err)
		return
	}

	log.Println(string(content))
	x := 0
	for x < len(content)-8 {
		log.Println(string(content[x : x+8]))
		x += 8
	}

	// another approach:
	newContent, err := os.Open("messages.txt")
	if err != nil {
		log.Println("there is an error", err)
		return
	}

	for {
		currentBuffer := make([]byte, 8)
		read, err := newContent.Read(currentBuffer)
		if err != nil {
			log.Println("there was an error", err)
			return
		}
		log.Println("read:", string(currentBuffer[:read]))
	}
}
