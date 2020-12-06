package balancers

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Balancer struct {
	Id int64 `json:"id"`
	Used []int64 `json:"usedMachines"`
	Total int64 `json:"totalMachinesCount"`
}

type Storage struct {
	Db *pgxpool.Pool
}

func NewStorage(db *pgxpool.Pool) *Storage {
	return &Storage{Db: db}
}

func (s *Storage) GetBalancersInfo() ([]*Balancer, error) {
	balancers, err := s.Db.Query(context.Background(), "SELECT get_balancers_id()")
	if err != nil {
		return nil, err
	}
	defer balancers.Close()

	/*
		I initialize empty slice here instead of nil declaration(via var)
		because I want to return an empty array, not a null, if there are no balancers provided
	*/
	res := make([]*Balancer, 0, 0)

	for balancers.Next() {
		var balancerId int64
		if err := balancers.Scan(&balancerId); err != nil {
			return nil, err
		}
		el, err := s.getBalancerInfo(balancerId)
		if err != nil {
			return nil, err
		}
		res = append(res, el)
	}
	return res, nil
}

func (s *Storage) getBalancerInfo(balancerId int64) (*Balancer, error) {
	res := Balancer{Id: balancerId, Used: make([]int64, 0)}

	err1 := s.Db.QueryRow(
		context.Background(),
		"SELECT get_machines_quantity($1)",
		balancerId,
	).Scan(&res.Total)

	if err1 != nil {
		return nil, err1
	}

	usable, err2 := s.Db.Query(context.Background(),"SELECT get_usable_machines($1)", balancerId)

	if err2 != nil {
		return nil, err2
	}
	defer usable.Close()

	for usable.Next() {
		var c int64
		if err := usable.Scan(&c); err != nil {
			return nil, err
		}
		res.Used = append(res.Used, c)
	}

	return &res, nil
}

func (s *Storage) UpdateMachine(machineId int64, state bool) error {
	_, err := s.Db.Exec(context.Background(),"CALL update_machine($1, $2)", machineId, state)
	return err
}
