// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package util

import (
	"fmt"
)

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func GetFrpcApiDomain(serverId, frpsDomain string) string {
	return fmt.Sprintf("api-%s", GetFrpcServerDomain(serverId, frpsDomain))
}

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func GetFrpcServerDomain(serverId, frpsDomain string) string {
	return fmt.Sprintf("%s.%s", serverId, frpsDomain)
}

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func GetFrpcHeadscaleUrl(protocol, serverId, frpsDomain string) string {
	return fmt.Sprintf("%s://%s", protocol, GetFrpcServerDomain(serverId, frpsDomain))
}

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func GetFrpcApiUrl(protocol, serverId, frpsDomain string) string {
	return fmt.Sprintf("%s://%s", protocol, GetFrpcApiDomain(serverId, frpsDomain))
}

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func GetFrpcRegistryDomain(serverId, frpsDomain string) string {
	return fmt.Sprintf("registry-%s", GetFrpcServerDomain(serverId, frpsDomain))
}

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func GetFrpcRegistryUrl(protocol, serverId, frpsDomain string) string {
	return fmt.Sprintf("%s://%s", protocol, GetFrpcRegistryDomain(serverId, frpsDomain))
}
