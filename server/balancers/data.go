package balancers

import (
	"database/sql"
	"fmt"
)

type Balancer struct {
	Id int64 `json:"id"`
	Used []int64 `json:"usedMachines"`
	Total int64 `json:"totalMachinesCount"`
}

type Storage struct {
	Db *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{Db: db}
}

func (s *Storage) GetBalancerInfo(balancer_id int64) (*Balancer, error) {
	quantity, err1 := s.Db.Query("SELECT get_machines_quantity($1)", balancer_id)
	usable, err2 := s.Db.Query("SELECT get_usable_machines($1)", balancer_id)

	if err1 != nil {
		return nil, err1
	}

	if err2 != nil {
		return nil, err2
	}

	defer quantity.Close()
	defer usable.Close()


	res := Balancer{Id: balancer_id}

	for usable.Next() {
		var c int64
		if err := usable.Scan(&c); err != nil {
			return nil, err
		}
		res.Used = append(res.Used, c)
	}

	var q int64
	if err := quantity.Scan(&q); err != nil {
		return nil, err
	}
	res.Total = q

	return &res, nil
}

func (s *Storage) UpdateMachine(machineId int64, state bool) error {
	if machineId <= 0 {
		return fmt.Errorf("error: machine id is invalid")
	}
	_, err := s.Db.Exec("CALL update_machine($1, $2)", machineId, state)
	return err
}