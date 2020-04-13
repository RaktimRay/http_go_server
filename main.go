package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type request_ping struct {
	Request string
}

type response_pong struct {
	Response string
}

func ping_handler(w http.ResponseWriter, r *http.Request) {

	var r_ping request_ping
	var r_pong response_pong

	err := json.NewDecoder(r.Body).Decode(&r_ping)
	if err != nil {
		panic(err)
	}

	log.Println(r_ping.Request)

	if r_ping.Request == "ping" {
		r_pong.Response = "pong"
		b, _ := json.Marshal(r_pong)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	} else {
		w.Write([]byte("Wrong Input"))
	}

}

func main() {
	http.HandleFunc("/", ping_handler)

	log.Println("Listening on localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
