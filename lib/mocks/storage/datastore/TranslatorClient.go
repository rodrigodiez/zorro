// Code generated by mockery v1.0.0. DO NOT EDIT.
package datastore

import context "context"
import datastore "cloud.google.com/go/datastore"
import mock "github.com/stretchr/testify/mock"
import storagedatastore "github.com/rodrigodiez/zorro/pkg/storage/datastore"

// TranslatorClient is an autogenerated mock type for the TranslatorClient type
type TranslatorClient struct {
	mock.Mock
}

// Get provides a mock function with given fields: _a0, _a1, _a2
func (_m *TranslatorClient) Get(_a0 context.Context, _a1 *datastore.Key, _a2 interface{}) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *datastore.Key, interface{}) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RunInTransaction provides a mock function with given fields: _a0, _a1, _a2
func (_m *TranslatorClient) RunInTransaction(_a0 context.Context, _a1 func(storagedatastore.Transaction) error, _a2 ...datastore.TransactionOption) (*datastore.Commit, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *datastore.Commit
	if rf, ok := ret.Get(0).(func(context.Context, func(storagedatastore.Transaction) error, ...datastore.TransactionOption) *datastore.Commit); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datastore.Commit)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, func(storagedatastore.Transaction) error, ...datastore.TransactionOption) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
