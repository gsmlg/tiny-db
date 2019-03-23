package main

import (
	"encoding/json"
	"fmt"
	"github.com/gsmlg/tiny-db/actions"
	"github.com/gsmlg/tiny-db/types"
	"ioutil"
	"log"
	"net/http"
)

func listHandler(db types.TinyDatabase) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		list, _ := json.Marshal(db.Data)
		w.Write([]byte(list))
		log.Println("List data...")
	}
	return http.HandlerFunc(fn)
}

func addHandler(db types.TinyDatabase) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var item types.TinyDataUnit
		body := ioutil.ReadAll(r.Body)
		err := json.Unmarshal(body, &item)
		db = actions.Add(db, item)
		w.Write([]byte(r.Body))
		log.Println("Add data to list...")
	}
	return http.HandlerFunc(fn)
}

func removeHandler(db types.TinyDatabase) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var item types.TinyDataUnit
		err := json.Unmarshal(r.Body, &item)
		db = actions.Remove(db, item)
		w.Write([]byte(r.Body))
		log.Println("Remove data from list...")
	}
	return http.HandlerFunc(fn)
}

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

	db = actions.Add(db, data)

	fmt.Println("%v", db)

	db = actions.Remove(db, data)

	fmt.Println("%v", db)

	server := http.NewServeMux()

	server.Handle("/", listHandler(&db))

	log.Println("Listening...")
	http.ListenAndServe(":3000", server)
}
