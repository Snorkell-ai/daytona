// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package util

import (
	"strings"
	"unicode"
)

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func GenerateIdFromName(name string) string {
	var result strings.Builder

	for _, char := range name {
		if unicode.IsLetter(char) || unicode.IsNumber(char) || char == '-' || char == '_' {
			result.WriteRune(char)
		} else if char == ' ' {
			result.WriteRune('_')
		}
	}

	return strings.ToLower(result.String())
}
