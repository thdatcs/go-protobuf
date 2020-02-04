// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	proto "github.com/golang/protobuf/proto"
	mock "github.com/stretchr/testify/mock"
)

// Queue is an autogenerated mock type for the Queue type
type Queue struct {
	mock.Mock
}

// SendProtoAsync provides a mock function with given fields: ctx, producer, topic, key, value
func (_m *Queue) SendProtoAsync(ctx context.Context, producer string, topic string, key string, value proto.Message) error {
	ret := _m.Called(ctx, producer, topic, key, value)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, proto.Message) error); ok {
		r0 = rf(ctx, producer, topic, key, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}