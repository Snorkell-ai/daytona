// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package common

import "os"

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func AgentMode() bool {
	_, devEnv := os.LookupEnv("DAYTONA_DEV")
	if devEnv {
		return false
	}
	val, agentMode := os.LookupEnv("DAYTONA_TARGET_ID")
	if agentMode && val != "" {
		return true
	}
	return false
}
