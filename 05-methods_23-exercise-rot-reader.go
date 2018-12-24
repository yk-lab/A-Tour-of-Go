package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13 rot13Reader) Read(b []byte) (int, error) {
	n, err := rot13.r.Read(b)
	if err == nil {
		for i := range b {
			if b[i] >= 'A' && b[i] < 'A'+13 || b[i] >= 'a' && b[i] < 'a'+13 {
				b[i] = b[i] + 13
			} else if b[i] >= 'A'+13 && b[i] <= 'Z' || b[i] >= 'a'+13 && b[i] <= 'z' {
				b[i] = b[i] - 13
			}
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
