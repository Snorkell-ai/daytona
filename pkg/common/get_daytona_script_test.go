//go:build testing

// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func TestGetDaytonaScript(t *testing.T) {
	baseUrl := "http://localhost:8080/daytona"
	expectedString := "http://localhost:8080/daytona"
	script := GetDaytonaScript(baseUrl)
	assert.Contains(t, script, expectedString, "the script should contain the correct base URL")

}
