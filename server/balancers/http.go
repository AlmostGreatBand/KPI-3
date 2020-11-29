package balancers

import (
	"encoding/json"
	"log"
	"net/http"
)

type Machine struct {
	Id int64 `json:"id"`
	State bool `json:"state"`
}

type Balancer struct {
	Id int64 `json:"id"`
}

func HttpHandler(storage *Storage) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {

		} else if r.Method == "PUT" {

		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func handleGetBalancerInfo(r *http.Request, rw http.ResponseWriter, storage *Storage) {
	var b Balancer
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		log.Printf("error: something went wrong while decoding json - %s", err)
		// tools write bad
		return
	}

	res, err := storage.GetBalancerInfo(b.Id)
	if err == nil {
		// tools write ok
		log.Println(res)
	} else {
		log.Printf("error: something went wrong while trying to get balancersInfo, %s", err)
		// tools write bad
	}
}

func handleUpdateMachine(r *http.Request, rw http.ResponseWriter, storage *Storage) {
	var m Machine
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		log.Printf("error: something went wrong while decoding json - %s", err)
		// tools write bad
		return
	}

	err := storage.UpdateMachine(m.Id, m.State)
	if err == nil {
		// tools write ok

	} else {
		log.Printf("error: something went wrong while trying to update machine %s", err)
		// tools write bad
		return
	}

}
