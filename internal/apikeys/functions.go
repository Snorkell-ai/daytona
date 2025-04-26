// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package apikeys

import (
	"encoding/base64"

	"github.com/daytonaio/daytona/internal/util"
	"github.com/google/uuid"
)

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func HashKey(key string) string {
	return util.Hash(key)
}

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func GenerateRandomKey() string {
	uuid := uuid.NewString()
	return base64.RawStdEncoding.EncodeToString([]byte(uuid))
}
