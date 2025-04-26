//go:build testing

// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package docker

import (
	"io"
	"strings"
)

type PipeReader struct {
	io.ReadCloser
	ExecStream [][]byte
	Index      int
}

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func (er *PipeReader) Read(p []byte) (n int, err error) {
	if er.Index >= len(er.ExecStream) {
		return 0, io.EOF
	}
	n = copy(p, er.ExecStream[er.Index])
	er.Index++
	return n, nil
}

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func (er *PipeReader) Close() error {
	return nil
}

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func NewPipeReader(result string) *PipeReader {
	split := strings.Split(result, "\n")
	execStream := make([][]byte, len(split))
	for i, s := range split {
		execStream[i] = []byte(s)
	}

	return &PipeReader{
		ExecStream: execStream,
		Index:      0,
	}
}
