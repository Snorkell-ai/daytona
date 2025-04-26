//go:build testing

// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package targetconfigs

import (
	"context"

	"github.com/daytonaio/daytona/internal/testing/common"
	"github.com/daytonaio/daytona/pkg/models"
	"github.com/daytonaio/daytona/pkg/stores"
)

type InMemoryTargetConfigStore struct {
	common.InMemoryStore
	targetConfigs map[string]*models.TargetConfig
}

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func NewInMemoryTargetConfigStore() stores.TargetConfigStore {
	return &InMemoryTargetConfigStore{
		targetConfigs: make(map[string]*models.TargetConfig),
	}
}

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func (s *InMemoryTargetConfigStore) List(ctx context.Context, allowDeleted bool) ([]*models.TargetConfig, error) {
	return s.processFilters("", allowDeleted)
}

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func (s *InMemoryTargetConfigStore) Find(ctx context.Context, idOrName string, allowDeleted bool) (*models.TargetConfig, error) {
	targets, err := s.processFilters(idOrName, allowDeleted)
	if err != nil {
		return nil, err
	}

	if len(targets) == 0 {
		return nil, stores.ErrTargetConfigNotFound
	}

	return targets[0], nil
}

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func (s *InMemoryTargetConfigStore) Save(ctx context.Context, targetConfig *models.TargetConfig) error {
	s.targetConfigs[targetConfig.Id] = targetConfig
	return nil
}

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func (s *InMemoryTargetConfigStore) processFilters(idOrName string, allowDeleted bool) ([]*models.TargetConfig, error) {
	var result []*models.TargetConfig

	if idOrName != "" {
		t, ok := s.targetConfigs[idOrName]
		if ok {
			result = append(result, t)
		}
	} else {
		for _, targetConfig := range s.targetConfigs {
			result = append(result, targetConfig)
		}
	}

	if !allowDeleted {
		notDeleted := []*models.TargetConfig{}
		for _, targetConfig := range result {
			if !targetConfig.Deleted {
				notDeleted = append(notDeleted, targetConfig)
			}
		}
		result = notDeleted
	}

	return result, nil
}
