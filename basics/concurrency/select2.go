package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const url_test string = "https://api.restful-api.dev/objects"

// const url_test string = "http://matt.aimonetti.net/"

func main() {
	response := make(chan *http.Response, 1)
	errors := make(chan *error)

	go func() {
		resp, err := http.Get(url_test)
		if err != nil {
			errors <- &err
		}
		response <- resp
	}()

	for {
		select {
		case r := <-response:
			fmt.Printf("%s", r.Body)
			return
		case err := <-errors:
			log.Fatal(*err)
		case <-time.After(20000 * time.Millisecond):
			fmt.Println("TimedOut")
			return
		}
	}

}
