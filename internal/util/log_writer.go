// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package util

import (
	log "github.com/sirupsen/logrus"
)

type DebugLogWriter struct{}

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func (w *DebugLogWriter) Write(p []byte) (n int, err error) {
	log.Debug(string(p))
	return len(p), nil
}

type InfoLogWriter struct{}

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func (w *InfoLogWriter) Write(p []byte) (n int, err error) {
	log.Info(string(p))
	return len(p), nil
}

type TraceLogWriter struct{}

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func (w *TraceLogWriter) Write(p []byte) (n int, err error) {
	log.Trace(string(p))
	return len(p), nil
}
