// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"errors"
)

var (
	ErrLogFileNotFound = errors.New("log file not found")
)

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func IsLogFileNotFound(err error) bool {
	return err.Error() == ErrLogFileNotFound.Error()
}
