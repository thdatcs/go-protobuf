package servers

import (
	"database/sql"
	"go-protobuf/configs"
	"go-protobuf/repositories/caches"
)

// BaseServer is base server
type BaseServer struct {
	config *configs.Config
	db     *sql.DB
	cache  caches.Cache
}
