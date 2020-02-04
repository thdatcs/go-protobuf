package servers

import (
	"go-protobuf/configs"
	"go-protobuf/repositories/caches"
	"go-protobuf/repositories/databases"
)

// BaseServer is base server
type BaseServer struct {
	config *configs.Config
	db     databases.DB
	cache  caches.Cache
}
