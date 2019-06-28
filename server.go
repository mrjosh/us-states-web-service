package main

import (
	"fmt"
	"os"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/gorilla/mux"
)

var port = ":8001"

func main() {

	r := mux.NewRouter()

	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	r.HandleFunc("/{state}/cities.json", func (w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		state := vars["state"]

		jsonFile, err := os.Open("./US_States_and_Cities.json")

		// if we os.Open returns an error then handle it
		if err != nil {
			fmt.Println(err)
		}

		var states map[string] []string

		byteValue, _ := ioutil.ReadAll(jsonFile)
		json.Unmarshal(byteValue, &states)

		defer jsonFile.Close()

		if states[state] == nil {

			http.NotFound(w, r)
		} else {

			json.NewEncoder(w).Encode(states[state])
		}
	})

	http.ListenAndServe(port, r)
}
