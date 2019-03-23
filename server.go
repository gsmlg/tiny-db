package main

import (
	"encoding/json"
	"fmt"
	"github.com/gsmlg/tiny-db/actions"
	"github.com/gsmlg/tiny-db/types"
	"io/ioutil"
	"log"
	"net/http"
)

func listHandler(db *types.TinyDatabase) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		list, _ := json.Marshal(db.Data)
		w.Write([]byte(list))
		log.Println("List data...")
	}
	return http.HandlerFunc(fn)
}

func addHandler(db *types.TinyDatabase) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var item types.TinyDataUnit
		body, _ := ioutil.ReadAll(r.Body)
		err := json.Unmarshal(body, &item)
		if err != nil {
			w.Write([]byte("errors"))
		}
		actions.Add(db, item)
		w.Write([]byte(body))
		log.Println("Add data to list...")
	}
	return http.HandlerFunc(fn)
}

func removeHandler(db *types.TinyDatabase) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var item types.TinyDataUnit
		body, _ := ioutil.ReadAll(r.Body)
		err := json.Unmarshal(body, &item)
		if err != nil {
			w.Write([]byte("errors"))
		}
		actions.Remove(db, item)
		w.Write([]byte(body))
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

	actions.Add(&db, data)

	fmt.Println("%v", db)

	actions.Remove(&db, data)

	fmt.Println("%v", db)

	server := http.NewServeMux()

	server.Handle("/", listHandler(&db))
	server.Handle("/add", addHandler(&db))
	server.Handle("/remove", removeHandler(&db))

	log.Println("Listening...")
	http.ListenAndServe(":3000", server)
}
