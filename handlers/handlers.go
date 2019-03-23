package handlers

import (
	"encoding/json"
	"github.com/gsmlg/tiny-db/actions"
	"github.com/gsmlg/tiny-db/types"
	"io/ioutil"
	"log"
	"net/http"
)

func ListHandler(db *types.TinyDatabase) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		list, _ := json.Marshal(db)
		w.Write([]byte(list))
		log.Println("List data...")
	}
	return http.HandlerFunc(fn)
}

func AddHandler(db *types.TinyDatabase) http.Handler {
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

func RemoveHandler(db *types.TinyDatabase) http.Handler {
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

func SaveHandler(db *types.TinyDatabase) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		list, _ := json.Marshal(db)
		err := ioutil.WriteFile("./db.data", []byte(list), 0600)
		if err != nil {
			w.Write([]byte("errors"))
		}
		w.Write([]byte("Data saved..."))
		log.Println("Data saved...")
	}
	return http.HandlerFunc(fn)
}

func LoadHandler(db *types.TinyDatabase) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		c, err := ioutil.ReadFile("./db.data")
		if err != nil {
			w.Write([]byte("errors"))
		}
		json.Unmarshal(c, db)
		w.Write([]byte("Data loaded..."))
		log.Println("Data loaded...")
	}
	return http.HandlerFunc(fn)
}
