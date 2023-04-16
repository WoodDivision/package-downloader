package service

import (
	"io"
	"log"
	"net/http"
)

func GetRequest(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Print("Wrong request.Check that you enter correct url")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Print("There is error in body pars")
	}
	return body
}
