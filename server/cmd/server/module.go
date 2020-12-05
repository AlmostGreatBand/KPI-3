//+build wireinject

package main

import (
	"github.com/AlmostGreatBand/KPI-3/server/balancers"
	"github.com/google/wire"
)

func ComposeApiServer(port HttpPortNumber) (*BalancersApiServer, error) {
	wire.Build(
		DatabaseConnection,
		balancers.Providers,
		wire.Struct(new(BalancersApiServer), "Port", "BalancersHandler"), 
	)
	return nil, nil
}
