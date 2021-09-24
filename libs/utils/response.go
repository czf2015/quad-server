package utils

import (
	"net/http"
	"io/ioutil"
)

func GetResponseBody(resp *http.Response) []byte {
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}