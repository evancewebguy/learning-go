package main

import (
	"fmt"
	"io"
	"strings"
)

/*
	NOTE:
		io(input/output) -  It needs to write in and write out data.
		some useful interfaces are - io.Reader and io.Writer
*/

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

func countLetters(r io.Reader) (map[string]int, error) {
	buf := make([]byte, 2048)
	out := map[string]int{}

	for {
		n, err := r.Read(buf)

		for _, b := range buf[:n] {
			if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
				out[string(b)]++
			}
		}

		if err == io.EOF {
			return out, nil
		}

		if err != nil {
			return nil, err
		}
	}
}

func main() {
	s := "this is what i am talking about when learning, sometimes you know nothing you are learning and still do it without understanding a thing"
	sreader := strings.NewReader(s)

	counts, err := countLetters(sreader)

	if err != nil {
		return
	}
	fmt.Println(counts)
}
