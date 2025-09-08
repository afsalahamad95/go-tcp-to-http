package request

import (
	"fmt"
	"io"
	"log"
	"strings"
)

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

func (r *RequestLine) IsValid() bool {
	return r.HttpVersion == "HTTP/1.1"
}

type Request struct {
	RequestLine RequestLine
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		log.Println("there was an error", err)
		return nil, err
	}
	requestLines := strings.Split(string(data), "\r\n")
	fmt.Println(requestLines)
	return &Request{
		RequestLine: RequestLine{
			HttpVersion:   requestLines[0],
			RequestTarget: requestLines[1],
			Method:        requestLines[2],
		},
	}, nil
}
