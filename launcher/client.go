package main

import (
	log "github.com/sirupsen/logrus"
	"halia-chat/client"
	"os"
)

func init() {
	log.SetLevel(log.DebugLevel)
	log.SetOutput(os.Stdout)
}

func main() {
	c := client.NewChatClient()
	log.Fatal(c.Dial("tcp", "127.0.0.1:8080"))
}
