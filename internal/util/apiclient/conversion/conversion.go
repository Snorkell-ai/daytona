// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package conversion

import (
	"encoding/json"
)

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func Convert[F any, T any](from *F) (*T, error) {
	if from == nil {
		return nil, nil
	}

	fromBytes, err := json.Marshal(from)
	if err != nil {
		return nil, err
	}

	var to T
	err = json.Unmarshal(fromBytes, &to)
	if err != nil {
		return nil, err
	}
	return &to, nil
}
