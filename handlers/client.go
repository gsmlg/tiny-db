package handlers

import (
	"fmt"
	"encoding/json"
	"github.com/gsmlg/tiny-db/types"
	"io/ioutil"
	"net/http"
	"bytes"
)

var port int
var baseurl string


func List() {
	port = types.Port
	baseurl = fmt.Sprintf("http://localhost:%d", port)

	resp, err := http.Get(baseurl + "/list")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	db := &types.TinyDatabase{}
	json.Unmarshal(body, db)
	fmt.Println("Database size: ", db.Size)
	fmt.Println("Database data: ")
	for k, v := range db.Data {
		fmt.Printf("Row %d: %s => %s\n", k, v.Key, v.Value)
	}
}

func Add(key string, value string) {
	port = types.Port
	baseurl = fmt.Sprintf("http://localhost:%d", port)

	item := types.TinyDataUnit{Key: key, Value: value}
	buf, err := json.Marshal(item)
	if err != nil {
		// handle error
	}
	r := bytes.NewReader(buf)
	resp, err := http.Post(baseurl + "/add", "application/json", r)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	db := &types.TinyDatabase{}
	json.Unmarshal(body, db)
	fmt.Println("Database size: ", db.Size)
	fmt.Println("Database data: ")
	for k, v := range db.Data {
		fmt.Printf("Row %d: %s => %s\n", k, v.Key, v.Value)
	}
}
