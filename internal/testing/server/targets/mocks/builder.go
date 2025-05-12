//go:build testing

// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package mocks

import (
	"github.com/daytonaio/daytona/internal/util"
	"github.com/daytonaio/daytona/pkg/build"
	"github.com/daytonaio/daytona/pkg/gitprovider"
	"github.com/daytonaio/daytona/pkg/models"
	"github.com/stretchr/testify/mock"
)

var MockBuild = &models.Build{
	Id:    "1",
	Image: util.Pointer("image"),
	User:  util.Pointer("user"),
	ContainerConfig: models.ContainerConfig{
		Image: "test",
		User:  "test",
	},
	BuildConfig: &models.BuildConfig{
		Devcontainer: MockWorkspaceTemplate.BuildConfig.Devcontainer,
	},
	Repository: &gitprovider.GitRepository{
		Url: MockWorkspaceTemplate.RepositoryUrl,
	},
	EnvVars: map[string]string{},
}

type MockBuilderFactory struct {
	mock.Mock
}

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func (f *MockBuilderFactory) Create(build models.Build, workspaceDir string) (build.IBuilder, error) {
	args := f.Called(build, workspaceDir)
	return args.Get(0).(*MockBuilder), args.Error(1)
}

type MockBuilder struct {
	mock.Mock
}

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func (b *MockBuilder) Build(build models.Build) (string, string, error) {
	args := b.Called(build)
	return args.String(0), args.String(1), args.Error(2)
}

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func (b *MockBuilder) CleanUp() error {
	args := b.Called()
	return args.Error(0)
}

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func (b *MockBuilder) Publish(build models.Build) error {
	args := b.Called(build)
	return args.Error(0)
}

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func (b *MockBuilder) SaveBuild(r models.Build) error {
	args := b.Called(r)
	return args.Error(0)
}

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func (b *MockBuilder) GetImageName(build models.Build) (string, error) {
	args := b.Called(build)
	return args.String(0), args.Error(1)
}
