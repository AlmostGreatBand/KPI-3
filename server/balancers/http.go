package balancers

import (
	"encoding/json"
	"github.com/AlmostGreatBand/KPI-3/server/utils"
	"log"
	"net/http"
)

type Machine struct {
	Id int64 `json:"id"`
	State bool `json:"state"`
}

func HttpHandler(storage *Storage) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handleGetBalancersInfo(rw, storage)
		} else if r.Method == "PUT" {
			handleUpdateMachine(r, rw, storage)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func handleGetBalancersInfo(rw http.ResponseWriter, storage *Storage) {
	res, err := storage.GetBalancersInfo()
	if err == nil {
		utils.ResponseOk(rw, res)
	} else {
		log.Printf("error: something went wrong while trying to get balancersInfo, %s", err)
		utils.ResponseInternalError(rw, "cannot get balancer info")
	}
}

func handleUpdateMachine(r *http.Request, rw http.ResponseWriter, storage *Storage) {
	var m Machine
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		log.Printf("error: something went wrong while decoding json - %s", err)
		utils.ResponseBadRequest(rw, "inappropriate json body")
		return
	}

	if m.Id <= 0 {
		utils.ResponseBadRequest(rw, "error: machine id is invalid")
		return
	}

	if err := storage.UpdateMachine(m.Id, m.State); err == nil {
		utils.ResponseOkNoBody(rw)
	} else {
		log.Printf("error: something went wrong while trying to update machine %s", err)
		utils.ResponseInternalError(rw, "cannot update machine status")
		return
	}
}
