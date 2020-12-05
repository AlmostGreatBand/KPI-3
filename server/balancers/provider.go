package balancers

import "github.com/google/wire"

var Providers = wire.NewSet(NewStorage, HttpHandler)
