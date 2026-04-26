package main

import (
	"log"
	"os"

	"github.com/emersion/go-imap/v2/imapclient"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	host := os.Getenv("IMAP_HOST")
	user := os.Getenv("IMAP_USER")
	pass := os.Getenv("IMAP_PASS")

	c, err := imapclient.DialTLS(host, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer c.Close()

	if err := c.Login(user, pass).Wait(); err != nil {
		log.Fatal(err)
	}

	mbox, err := c.Select("INBOX", nil).Wait()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("There are %d messages in the inbox", mbox.NumMessages)

	c.Logout().Wait()
}