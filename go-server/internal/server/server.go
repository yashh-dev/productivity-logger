package server

import (
	"github.com/google/wire"
)

// ProviderSet is server providers.
var ServerSet = wire.NewSet(NewGRPCServer, NewHTTPServer)
