package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func get() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	log.Println(string(body))
}

func post() {
	data := map[string]interface{}{
		"title":  "foo",
		"body":   "bar",
		"userId": 1,
	}

	client := &http.Client{}

	jsonByte, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	buff := bytes.NewBuffer(jsonByte)

	req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", buff)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatal(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(body))

}

func simpleAPI() {
	http.HandleFunc("/student")

	server := new(http.Server)
	server.Addr = ":8080"

	log.Fatal(server.ListenAndServe())
}

func main() {
	//get()
	//post()
	simpleAPI()
}
