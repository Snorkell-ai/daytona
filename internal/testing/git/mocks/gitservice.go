//go:build testing

// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package mocks

import (
	"github.com/daytonaio/daytona/pkg/gitprovider"
	"github.com/daytonaio/daytona/pkg/models"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/stretchr/testify/mock"
)

type MockGitService struct {
	mock.Mock
}

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func (m *MockGitService) CloneRepository(repo *gitprovider.GitRepository, auth *http.BasicAuth) error {
	args := m.Called(repo, auth)
	return args.Error(0)
}

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func (m *MockGitService) CloneRepositoryCmd(repo *gitprovider.GitRepository, auth *http.BasicAuth) []string {
	args := m.Called(repo, auth)
	return args.Get(0).([]string)
}

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func (m *MockGitService) RepositoryExists() (bool, error) {
	args := m.Called()
	return args.Bool(0), args.Error(1)
}

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func (m *MockGitService) SetGitConfig(userData *gitprovider.GitUser, providerConfig *models.GitProviderConfig) error {
	args := m.Called(userData, providerConfig)
	return args.Error(0)
}

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func (m *MockGitService) GetGitStatus() (*models.GitStatus, error) {
	args := m.Called()
	return args.Get(0).(*models.GitStatus), args.Error(1)
}

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func NewMockGitService() *MockGitService {
	gitService := new(MockGitService)
	return gitService
}
