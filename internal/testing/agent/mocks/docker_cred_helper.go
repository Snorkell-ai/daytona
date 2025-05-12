//go:build testing

// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package mocks

import "github.com/stretchr/testify/mock"

type mockDockerCredHelper struct {
	mock.Mock
}

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func (m *mockDockerCredHelper) SetDockerConfig() error {
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
func NewMockDockerCredHelper() *mockDockerCredHelper {
	mockCredHelper := new(mockDockerCredHelper)
	mockCredHelper.On("SetDockerConfig").Return(nil)

	return mockCredHelper
}
