package utils

import (
	"bytes"
	"encoding/gob"
)

func Copy(old interface{}) (new interface{}) {
	var err error
	var buf bytes.Buffer

	enc := gob.NewEncoder(&buf)
	err = enc.Encode(old)
	if err != nil {
		panic(err)
	}

	dec := gob.NewDecoder(&buf)
	err = dec.Decode(&new)
	if err != nil {
		panic(err)
	}
	
	return new
}