// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package apiclient

import (
	"fmt"
	"strings"
)

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func ErrHealthCheckFailed(healthUrl string) error {
	return fmt.Errorf("failed to check server health at: %s. Make sure Daytona is running on the appropriate port", healthUrl)
}

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func IsHealthCheckFailed(err error) bool {
	return strings.HasPrefix(err.Error(), "failed to check server health at:")
}
