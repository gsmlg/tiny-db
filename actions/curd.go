package actions

import (
	"github.com/gsmlg/tiny-db/types"
)

func Add(db *types.TinyDatabase, item types.TinyDataUnit) *types.TinyDatabase {
	db.Data = append(db.Data, item)
	db.Size = db.Size + 1
	return db
}

func Remove(db *types.TinyDatabase, item types.TinyDataUnit) *types.TinyDatabase {
	idx := FindIndex(db, item)
	if idx == -1 {
		return db
	}
	RemoveIndex(db, idx)
	return db
}

func FindIndex(db *types.TinyDatabase, item types.TinyDataUnit) int {
	for k, v := range db.Data {
		if v.Key == item.Key {
			return k
		}
	}
	return -1
}

func RemoveIndex(db *types.TinyDatabase, idx int) *types.TinyDatabase {
	db.Data[idx] = db.Data[db.Size-1]
	db.Data = db.Data[:db.Size-1]
	db.Size = db.Size - 1
	return db
}
