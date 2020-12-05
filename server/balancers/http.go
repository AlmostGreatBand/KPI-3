package balancers

import (
	"encoding/json"
	"github.com/AlmostGreatBand/KPI-3/server/utils"
	"log"
	"net/http"
)

type MachineRequest struct {
	Id int64 `json:"id"`
	State bool `json:"state"`
}

type BalancerRequest struct {
	Id int64 `json:"id"`
}

func HttpHandler(storage *Storage) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handleGetBalancerInfo(r, rw, storage)
		} else if r.Method == "PUT" {
			handleUpdateMachine(r, rw, storage)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func handleGetBalancerInfo(r *http.Request, rw http.ResponseWriter, storage *Storage) {
	var b BalancerRequest
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		log.Printf("error: something went wrong while decoding json - %s", err)
		utils.ResponseBadRequest(rw, "inappropriate json body")
		return
	}

	res, err := storage.GetBalancerInfo(b.Id)
	if err == nil {
		log.Println(res)
		utils.ResponseOk(rw, res)
	} else {
		log.Printf("error: something went wrong while trying to get balancersInfo, %s", err)
		utils.ResponseInternalError(rw, "cannot get balancer info")
	}
}

func handleUpdateMachine(r *http.Request, rw http.ResponseWriter, storage *Storage) {
	var m MachineRequest
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		log.Printf("error: something went wrong while decoding json - %s", err)
		utils.ResponseBadRequest(rw, "inappropriate json body")
		return
	}

	err := storage.UpdateMachine(m.Id, m.State)
	if err == nil {
		utils.ResponseOkNoBody(rw)
	} else {
		log.Printf("error: something went wrong while trying to update machine %s", err)
		utils.ResponseInternalError(rw, "cannot update machine status")
		return
	}
}
