package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Calling Service B\n")

	req, err := http.NewRequest("GET", "http://service_a_envoy:8788/", nil)
	if err != nil {
		fmt.Printf("%s", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Printf("%s", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Fprintf(w, string(body)+"\n")
	fmt.Fprintf(w, "Calling Service C\n")

	req, err = http.NewRequest("GET", "http://service_a_envoy:8791/", nil)
	if err != nil {
		fmt.Printf("%s", err)
	}

	client = &http.Client{}
	resp, err = client.Do(req)

	if err != nil {
		fmt.Printf("%s", err)
	}

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
	}

	fmt.Fprintf(w, string(body))
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
