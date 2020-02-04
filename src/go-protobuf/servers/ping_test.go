package servers

import (
	"context"
	"testing"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/stretchr/testify/suite"
)

type PingServerTestSuite struct {
	suite.Suite
}

func (s *PingServerTestSuite) TestPing() {
	tests := []struct {
		name string
		args struct {
			ctx     context.Context
			request *empty.Empty
		}
		want    *empty.Empty
		wantErr error
	}{
		{
			name: "success",
			args: struct {
				ctx     context.Context
				request *empty.Empty
			}{
				ctx:     context.Background(),
				request: &empty.Empty{},
			},
			want:    &empty.Empty{},
			wantErr: nil,
		},
	}
	for _, test := range tests {
		s.T().Run(test.name, func(t *testing.T) {
			pingServer := NewPingServer()
			if response, err := pingServer.Ping(test.args.ctx, test.args.request); !s.Assert().Equal(response, test.want) && !s.Assert().Equal(err, test.wantErr) {
				t.Errorf("pingServer.Ping() = (%v, %v), want (%v, %v)", response, err, test.want, test.wantErr)
			}
		})
	}
}

func TestPingServerTestSuite(t *testing.T) {
	suite.Run(t, new(PingServerTestSuite))
}
