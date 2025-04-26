// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package util

import "os"

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func DirectoryValidator(path *string) error {
	_, err := os.Stat(*path)
	if os.IsNotExist(err) {
		return os.MkdirAll(*path, 0700)
	}
	return err
}
