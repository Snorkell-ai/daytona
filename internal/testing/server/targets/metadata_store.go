//go:build testing

// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package targets

import (
	"context"

	"github.com/daytonaio/daytona/internal/testing/common"
	"github.com/daytonaio/daytona/pkg/models"
	"github.com/daytonaio/daytona/pkg/stores"
)

type InMemoryTargetMetadataStore struct {
	common.InMemoryStore
	targetMetadataEntries map[string]*models.TargetMetadata
}

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func NewInMemoryTargetMetadataStore() stores.TargetMetadataStore {
	return &InMemoryTargetMetadataStore{
		targetMetadataEntries: make(map[string]*models.TargetMetadata),
	}
}

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func (s *InMemoryTargetMetadataStore) Find(ctx context.Context, targetId string) (*models.TargetMetadata, error) {
	metadata, ok := s.targetMetadataEntries[targetId]
	if !ok {
		return nil, stores.ErrTargetMetadataNotFound
	}

	return metadata, nil
}

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func (s *InMemoryTargetMetadataStore) Save(ctx context.Context, targetMetadata *models.TargetMetadata) error {
	s.targetMetadataEntries[targetMetadata.TargetId] = targetMetadata
	return nil
}

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func (s *InMemoryTargetMetadataStore) Delete(ctx context.Context, targetMetadata *models.TargetMetadata) error {
	delete(s.targetMetadataEntries, targetMetadata.TargetId)
	return nil
}
