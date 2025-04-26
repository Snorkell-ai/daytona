// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package common

import (
	"errors"
	"strings"
)

var (
	ErrCtrlCAbort = errors.New("ctrl-c exit")
)

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func IsCtrlCAbort(err error) bool {
	return err.Error() == ErrCtrlCAbort.Error()
}

var (
	ErrConnection = errors.New("If you are using a VPN or firewall, please read our troubleshooting guide at https://daytona.io/docs/misc/troubleshooting#connectivity-issues")
)

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func IsConnectionError(err error) bool {
	return strings.Contains(err.Error(), ErrConnection.Error())
}
