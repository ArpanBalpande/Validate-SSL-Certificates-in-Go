package main

import (
	"crypto/tls"
)

func main() {
	conn, err := tls.Dial("tcp", "freecodecamp.org:443", nil)
	if err != nil {
		panic("Server doesn't support SSL certificate err: " + err.Error())
	}
}