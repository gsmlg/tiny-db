package types

type TinyDataUnit struct {
	Key   string
	Value string
}

type TinyDatabase struct {
	Data []TinyDataUnit
	Size int
}
