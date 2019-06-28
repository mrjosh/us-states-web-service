package main

import (
	"fmt"
	"os"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/{state}/cities.json", func (w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		state := mux.Vars(r)["state"]

		jsonFile, jErr := os.Open("./US_States_and_Cities.json")
		if jErr != nil {
			fmt.Println(jErr)
		}

		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)
		states := new(States)
		json.Unmarshal(byteValue, &states)

		if !states.Has(state) {

			json.NewEncoder(w).Encode(map[string] string {
				"error": "Not found",
			})
		} else {

			json.NewEncoder(w).Encode(states.Get(state))
		}
	})

	http.ListenAndServe(":8001", r)
}
