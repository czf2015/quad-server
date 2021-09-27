package utils

import (
	"net/http"
	"io/ioutil"
)

func GetResponseBody(res *http.Response) []byte {
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return body
}