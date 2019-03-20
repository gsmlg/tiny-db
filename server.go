package main

import (
	"fmt"
)

type TinyDataUnit struct {
	Key   string
	Value string
}

type TinyDatabase struct {
	data []TinyDataUnit
	size int
}

func main() {
	data := TinyDataUnit{
		Key:   "j",
		Value: "josh",
	}
	fmt.Println("%v", data)
}
