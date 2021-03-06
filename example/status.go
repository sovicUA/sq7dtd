package main

import (
	"flag"
	"github.com/sovicUA/sq7dtd"
	"log"
)

var host = flag.String("host", "localhost", "The game server hostname or IP address")
var port = flag.Int("port", 26900, "Listener port of game server")

func main() {

	flag.Parse()

	err := sq7dtd.QueryInfo(*host, *port)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println(sq7dtd.Json())
	log.Println("\n" + sq7dtd.String())
}
