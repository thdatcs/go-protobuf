package servers

import (
	"context"
	"go-protobuf/api"

	"github.com/golang/protobuf/ptypes/empty"
)

type pingServer struct {
}

// NewPingServer creates a new instance
func NewPingServer() api.PingServer {
	return &pingServer{}
}

func (s *pingServer) Ping(ctx context.Context, request *empty.Empty) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}
