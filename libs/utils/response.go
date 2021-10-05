package utils

import (
	"io/ioutil"
	"net/http"
)

func GetResponseBody(res *http.Response) []byte {
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return body
}
