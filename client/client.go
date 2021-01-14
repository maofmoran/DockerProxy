package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"time"
)

func main() {

	if err := exec.Command("route", "add", "default", "gw", "192.168.2.150", "eth0").Run(); err != nil {
		log.Fatal(err)
	} else {
		log.Print("added 150 as dg")
	}

	if err := exec.Command("route", "del", "default", "gw", "192.168.2.1").Run(); err != nil {
		log.Fatal(err)
	} else {
		log.Print("removed 1 as dg")
	}

	fmt.Println("AVI")

	for {
		time.Sleep(5 * time.Second)
		resp, err := http.Get("http://192.168.1.10:8080/ping")
		if err != nil {
			log.Printf("Error: %v", err)
		} else {
			log.Printf("Got status code: %v", resp.StatusCode)
		}
	}
}
