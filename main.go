package main

import (
	"fmt"

	"github.com/emersion/go-imap/v2/imapclient"
)

func main() {
	
	opts := &imapclient.Options{
		Username: "[EMAIL_ADDRESS]",
		Password: "[PASSWORD]",
		TLS:      imapclient.TLSAuto,
	}

	client, err := imapclient.DialWithTLS("imap.gmail.com:993", opts)
	if err != nil {
		panic(err)
	}

	defer client.Close()

	fmt.Println("Connected to server!")
}