package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type HelloRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type HelloResponse struct {
	Msg string `json:"msg"`
}

func main() {
	var (
		port = flag.String("port", ":7000", "http listen address")
	)

	router := mux.NewRouter()
	router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var p HelloRequest
		err = json.Unmarshal(data, &p)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		log.Println("****************************************")
		log.Println("Received request")

		res := &HelloResponse{
			Msg: fmt.Sprintf("Hello %s %s... welcome to simple-go-app", p.FirstName, p.LastName),
		}
		out, err := json.Marshal(res)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(out)
		w.WriteHeader(http.StatusOK)
		log.Println("Request served...")
		log.Println("****************************************")
	})
	err := http.ListenAndServe(*port, router)
	if err != nil {
		log.Fatal(err)
	}
}
