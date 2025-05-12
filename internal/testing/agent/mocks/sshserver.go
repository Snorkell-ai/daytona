//go:build testing

// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package mocks

import (
	"errors"

	"github.com/stretchr/testify/mock"
)

var SshServerStartError = errors.New("start error")

type mockSshServer struct {
	mock.Mock
}

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func (m *mockSshServer) Start() error {
	args := m.Called()
	return args.Error(0)
}

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func NewMockSshServer() *mockSshServer {
	mockSshServer := new(mockSshServer)
	mockSshServer.On("Start").Return(SshServerStartError)

	return mockSshServer
}
