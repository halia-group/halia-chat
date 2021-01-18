package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"halia-chat/server"
	"os"
)

var (
	addr string
)

func init() {
	log.SetLevel(log.DebugLevel)
	log.SetOutput(os.Stdout)
	flag.StringVar(&addr, "addr", ":8080", "listen address")
}

func main() {
	flag.Parse()

	cs, err := server.NewServer()
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(cs.Run("tcp", addr))
}
