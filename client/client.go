package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	for {
		time.Sleep(5 * time.Second)
		resp, err := http.Get("http://192.168.2.150:8080/ping")
		if err != nil {
			log.Printf("Error: %v", err)
		} else {
			log.Printf("Got status code: %v", resp.StatusCode)
		}
	}
}
