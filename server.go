package main

import (
	"fmt"
	"github.com/gsmlg/tiny-db/handlers"
	"github.com/gsmlg/tiny-db/types"
	"log"
	"net/http"
)

func main() {
	db := types.TinyDatabase{
		Data: []types.TinyDataUnit{},
		Size: 0,
	}

	server := http.NewServeMux()

	server.Handle("/", handlers.ListHandler(&db))
	server.Handle("/add", handlers.AddHandler(&db))
	server.Handle("/remove", handlers.RemoveHandler(&db))
	server.Handle("/save", handlers.SaveHandler(&db))
	server.Handle("/load", handlers.LoadHandler(&db))

	port := types.Port
	log.Println(fmt.Sprintf("Listening at port %d ...", port))
	http.ListenAndServe(fmt.Sprintf(":%d", port), server)
}
