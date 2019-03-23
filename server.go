package main

import (
	"fmt"
	"github.com/gsmlg/tiny-db/actions"
	"github.com/gsmlg/tiny-db/handlers"
	"github.com/gsmlg/tiny-db/types"
	"log"
	"net/http"
)

func main() {
	data := types.TinyDataUnit{
		Key:   "j",
		Value: "josh",
	}
	fmt.Println("%v", data)

	db := types.TinyDatabase{
		Data: []types.TinyDataUnit{},
		Size: 0,
	}

	actions.Add(&db, data)

	fmt.Println("%v", db)

	actions.Remove(&db, data)

	fmt.Println("%v", db)

	server := http.NewServeMux()

	server.Handle("/", handlers.ListHandler(&db))
	server.Handle("/add", handlers.AddHandler(&db))
	server.Handle("/remove", handlers.RemoveHandler(&db))
	server.Handle("/save", handlers.SaveHandler(&db))
	server.Handle("/load", handlers.LoadHandler(&db))

	log.Println("Listening...")
	http.ListenAndServe(":3000", server)
}
