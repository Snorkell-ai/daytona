// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package util

import (
	"os/exec"

	"github.com/daytonaio/daytona/pkg/os"
)

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func GetRemoteOS(remote string) (*os.OperatingSystem, error) {
	unameCmd := exec.Command("ssh", remote, "uname -a")

	output, err := unameCmd.Output()
	if err != nil {
		return nil, err
	}

	return os.OSFromUnameA(string(output))
}
