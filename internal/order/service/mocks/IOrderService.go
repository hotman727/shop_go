// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	context "context"
	dto "goshop/internal/order/dto"

	mock "github.com/stretchr/testify/mock"

	model "goshop/internal/order/model"

	paging "goshop/pkg/paging"
)

// IOrderService is an autogenerated mock type for the IOrderService type
type IOrderService struct {
	mock.Mock
}

// CancelOrder provides a mock function with given fields: ctx, orderID, userID
func (_m *IOrderService) CancelOrder(ctx context.Context, orderID string, userID string) (*model.Order, error) {
	ret := _m.Called(ctx, orderID, userID)

	var r0 *model.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*model.Order, error)); ok {
		return rf(ctx, orderID, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *model.Order); ok {
		r0 = rf(ctx, orderID, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Order)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, orderID, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMyOrders provides a mock function with given fields: ctx, req
func (_m *IOrderService) GetMyOrders(ctx context.Context, req *dto.ListOrderReq) ([]*model.Order, *paging.Pagination, error) {
	ret := _m.Called(ctx, req)

	var r0 []*model.Order
	var r1 *paging.Pagination
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, *dto.ListOrderReq) ([]*model.Order, *paging.Pagination, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dto.ListOrderReq) []*model.Order); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Order)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dto.ListOrderReq) *paging.Pagination); ok {
		r1 = rf(ctx, req)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*paging.Pagination)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, *dto.ListOrderReq) error); ok {
		r2 = rf(ctx, req)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetOrderByID provides a mock function with given fields: ctx, id
func (_m *IOrderService) GetOrderByID(ctx context.Context, id string) (*model.Order, error) {
	ret := _m.Called(ctx, id)

	var r0 *model.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*model.Order, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.Order); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Order)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PlaceOrder provides a mock function with given fields: ctx, req
func (_m *IOrderService) PlaceOrder(ctx context.Context, req *dto.PlaceOrderReq) (*model.Order, error) {
	ret := _m.Called(ctx, req)

	var r0 *model.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dto.PlaceOrderReq) (*model.Order, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dto.PlaceOrderReq) *model.Order); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Order)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dto.PlaceOrderReq) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIOrderService creates a new instance of IOrderService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIOrderService(t interface {
	mock.TestingT
	Cleanup(func())
}) *IOrderService {
	mock := &IOrderService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
