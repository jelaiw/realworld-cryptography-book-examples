package main

import "crypto/tls"

func main() {
	destination := "google.com:443"
	TLSconfig := &tls.Config{}
	conn, err := tls.Dial("tcp", destination, TLSconfig)
	if err != nil {
		panic("failed to connect: " + err.Error())
	}
	conn.Close()
}
