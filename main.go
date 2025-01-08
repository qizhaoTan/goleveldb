package main

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
)

func main() {
	db, err := leveldb.OpenFile("./data/leveldb/", nil)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	db.Put([]byte("name"), []byte("tan"), nil)

	value, _ := db.Get([]byte("name"), nil)

	fmt.Println(string(value))
}
