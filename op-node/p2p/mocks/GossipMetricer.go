// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	peer "github.com/libp2p/go-libp2p/core/peer"
)

// GossipMetricer is an autogenerated mock type for the GossipMetricer type
type GossipMetricer struct {
	mock.Mock
}

// RecordGossipEvent provides a mock function with given fields: evType
func (_m *GossipMetricer) RecordGossipEvent(evType int32) {
	_m.Called(evType)
}

// RecordPeerScoring provides a mock function with given fields: peerID, score
func (_m *GossipMetricer) RecordPeerScoring(peerID peer.ID, score float64) {
	_m.Called(peerID, score)
}

type mockConstructorTestingTNewGossipMetricer interface {
	mock.TestingT
	Cleanup(func())
}

// NewGossipMetricer creates a new instance of GossipMetricer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGossipMetricer(t mockConstructorTestingTNewGossipMetricer) *GossipMetricer {
	mock := &GossipMetricer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
