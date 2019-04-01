package types

import (
	// "encoding/json"
)

type TinyDataUnit struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type TinyDatabase struct {
	Data []TinyDataUnit `json:"data"`
	Size int            `json:"size"`
}

const Port = 21985
