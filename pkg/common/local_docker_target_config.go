// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package common

import (
	"strings"
)

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func IsLocalDockerTarget(providerName, options, runnerId string) bool {
	if providerName != "docker-provider" {
		return false
	}

	return !strings.Contains(options, "Remote Hostname") && runnerId == LOCAL_RUNNER_ID
}
