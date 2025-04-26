// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package tunnel

import "fmt"

const (
	endpointTypeUnixSocket = "unix"
	endpointTypeTCP        = "tcp"
)

type Endpoint struct {
	host       string
	port       int
	unixSocket string
}

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func (e *Endpoint) String() string {
	if e.unixSocket != "" {
		return e.unixSocket
	}
	return fmt.Sprintf("%s:%d", e.host, e.port)
}

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func (e *Endpoint) Type() string {
	if e.unixSocket != "" {
		return endpointTypeUnixSocket
	}
	return endpointTypeTCP
}

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func NewTCPEndpoint(host string, port int) *Endpoint {
	return &Endpoint{
		host: host,
		port: port,
	}
}

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func NewUnixEndpoint(socket string) *Endpoint {
	return &Endpoint{
		unixSocket: socket,
	}
}
