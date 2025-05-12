//go:build testing

// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package workspaces

import (
	"context"

	"github.com/daytonaio/daytona/internal/testing/common"
	"github.com/daytonaio/daytona/pkg/models"
	"github.com/daytonaio/daytona/pkg/stores"
)

type InMemoryWorkspaceMetadataStore struct {
	common.InMemoryStore
	workspaceMetadataEntries map[string]*models.WorkspaceMetadata
}

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func NewInMemoryWorkspaceMetadataStore() stores.WorkspaceMetadataStore {
	return &InMemoryWorkspaceMetadataStore{
		workspaceMetadataEntries: make(map[string]*models.WorkspaceMetadata),
	}
}

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func (s *InMemoryWorkspaceMetadataStore) Find(ctx context.Context, workspaceId string) (*models.WorkspaceMetadata, error) {
	if _, ok := s.workspaceMetadataEntries[workspaceId]; !ok {
		return nil, stores.ErrWorkspaceMetadataNotFound
	}

	return s.workspaceMetadataEntries[workspaceId], nil
}

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func (s *InMemoryWorkspaceMetadataStore) Save(ctx context.Context, workspaceMetadata *models.WorkspaceMetadata) error {
	s.workspaceMetadataEntries[workspaceMetadata.WorkspaceId] = workspaceMetadata
	return nil
}

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func (s *InMemoryWorkspaceMetadataStore) Delete(ctx context.Context, workspaceMetadata *models.WorkspaceMetadata) error {
	delete(s.workspaceMetadataEntries, workspaceMetadata.WorkspaceId)
	return nil
}
