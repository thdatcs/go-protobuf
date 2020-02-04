// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	proto "github.com/golang/protobuf/proto"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// Cache is an autogenerated mock type for the Cache type
type Cache struct {
	mock.Mock
}

// Del provides a mock function with given fields: ctx, key
func (_m *Cache) Del(ctx context.Context, key string) error {
	ret := _m.Called(ctx, key)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DelTx provides a mock function with given fields: ctx, keys
func (_m *Cache) DelTx(ctx context.Context, keys ...string) error {
	_va := make([]interface{}, len(keys))
	for _i := range keys {
		_va[_i] = keys[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, ...string) error); ok {
		r0 = rf(ctx, keys...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Exists provides a mock function with given fields: ctx, key
func (_m *Cache) Exists(ctx context.Context, key string) error {
	ret := _m.Called(ctx, key)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetBytes provides a mock function with given fields: ctx, key, del
func (_m *Cache) GetBytes(ctx context.Context, key string, del bool) ([]byte, error) {
	ret := _m.Called(ctx, key, del)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(context.Context, string, bool) []byte); ok {
		r0 = rf(ctx, key, del)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, bool) error); ok {
		r1 = rf(ctx, key, del)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetInt provides a mock function with given fields: ctx, key, del
func (_m *Cache) GetInt(ctx context.Context, key string, del bool) (int, error) {
	ret := _m.Called(ctx, key, del)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, string, bool) int); ok {
		r0 = rf(ctx, key, del)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, bool) error); ok {
		r1 = rf(ctx, key, del)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProto provides a mock function with given fields: ctx, key, value, del
func (_m *Cache) GetProto(ctx context.Context, key string, value proto.Message, del bool) error {
	ret := _m.Called(ctx, key, value, del)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, proto.Message, bool) error); ok {
		r0 = rf(ctx, key, value, del)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetString provides a mock function with given fields: ctx, key, del
func (_m *Cache) GetString(ctx context.Context, key string, del bool) (string, error) {
	ret := _m.Called(ctx, key, del)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string, bool) string); ok {
		r0 = rf(ctx, key, del)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, bool) error); ok {
		r1 = rf(ctx, key, del)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Incr provides a mock function with given fields: ctx, key
func (_m *Cache) Incr(ctx context.Context, key string) (int64, error) {
	ret := _m.Called(ctx, key)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string) int64); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Set provides a mock function with given fields: ctx, key, value, expiration
func (_m *Cache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	ret := _m.Called(ctx, key, value, expiration)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}, time.Duration) error); ok {
		r0 = rf(ctx, key, value, expiration)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetProto provides a mock function with given fields: ctx, key, value, expiration
func (_m *Cache) SetProto(ctx context.Context, key string, value proto.Message, expiration time.Duration) error {
	ret := _m.Called(ctx, key, value, expiration)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, proto.Message, time.Duration) error); ok {
		r0 = rf(ctx, key, value, expiration)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetProtoTx provides a mock function with given fields: ctx, data, expiration
func (_m *Cache) SetProtoTx(ctx context.Context, data map[string]proto.Message, expiration time.Duration) error {
	ret := _m.Called(ctx, data, expiration)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, map[string]proto.Message, time.Duration) error); ok {
		r0 = rf(ctx, data, expiration)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetTx provides a mock function with given fields: ctx, data, expiration
func (_m *Cache) SetTx(ctx context.Context, data map[string]interface{}, expiration time.Duration) error {
	ret := _m.Called(ctx, data, expiration)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, map[string]interface{}, time.Duration) error); ok {
		r0 = rf(ctx, data, expiration)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
