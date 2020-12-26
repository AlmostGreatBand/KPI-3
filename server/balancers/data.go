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
	balancers, err := s.Db.Query(context.Background(), "SELECT * FROM get_balancers_info()")
	if err != nil {
		return nil, err
	}
	defer balancers.Close()

	/*
		I initialize empty slice here instead of nil declaration(via var)
		because I want to return an empty array, not a nil, if there are no balancers provided
	*/
	res := make([]*Balancer, 0, 0)

	usedMachines, err := s.getUsedMachines()
	if err != nil {
		return nil, err
	}

	for balancers.Next() {
		var balancerId int64
		var count int64
		if err := balancers.Scan(&balancerId, &count); err != nil {
			return nil, err
		}

		/*
			as with balancers - return empty slice instead of returning nil
		*/
		used := usedMachines[balancerId]
		if used == nil {
			used = make([]int64, 0)
		}

		el := &Balancer{Id: balancerId, Used: used, Total: count}
		res = append(res, el)
	}
	return res, nil
}

func (s *Storage) getUsedMachines() (map[int64][]int64, error) {
	usedMachines, err1 := s.Db.Query(context.Background(), "SELECT * FROM get_usable_machines()")
	dividedUsedMachines := map[int64][]int64{}
	if err1 != nil {
		return nil, err1
	}
	defer usedMachines.Close()

	for usedMachines.Next() {
		var balancerId int64
		var machineId int64
		if err := usedMachines.Scan(&balancerId, &machineId); err != nil {
			return nil, err
		}
		dividedUsedMachines[balancerId] = append(dividedUsedMachines[balancerId], machineId)
	}

	return dividedUsedMachines, nil
}

func (s *Storage) UpdateMachine(machineId int64, state bool) error {
	_, err := s.Db.Exec(context.Background(),"CALL update_machine($1, $2)", machineId, state)
	return err
}
