package main

import (
	"flag"
	"github.com/gsmlg/tiny-db/handlers"
	"net/http"
)

func main() {
	flag.Parse()
	cmd := flag.Arg(0)
	key := flag.Arg(1)
	value := flag.Arg(2)

	if cmd == "list" {
		handlers.List()
	}

	if cmd == "add" {
		handlers.Add(key, value)
	}

	if cmd == "remove" {

	}

}
