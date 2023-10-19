package main

import (
	"crypto/tls"
	"net/http"
)

func hello(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("Hello, world\n"))
}

func main() {
	config := &tls.Config{
		MinVersion: tls.VersionTLS13,
	}

	http.HandleFunc("/", hello)

	server := &http.Server{
		Addr:      ":8080",
		TLSConfig: config,
	}

	cert := "cert.pem"
	key := "key.pem"
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		panic(err)
	}
}
