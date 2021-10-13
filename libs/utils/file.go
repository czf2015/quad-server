package utils

import (
	"io/ioutil"
	"log"
)

func ReadFile(paths... string) []byte {
	file, err := ioutil.ReadFile(JoinPath(paths...))
	if err != nil {
		log.Printf(err.Error())
	}
	return file
}