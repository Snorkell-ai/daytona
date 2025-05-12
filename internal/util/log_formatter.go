// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package util

import (
	"io"

	log "github.com/sirupsen/logrus"
)

type LogFormatter struct {
	TextFormatter    *log.TextFormatter
	ProcessLogWriter io.Writer
}

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func (f *LogFormatter) Format(entry *log.Entry) ([]byte, error) {
	formatted, err := f.TextFormatter.Format(entry)
	if err != nil {
		return nil, err
	}

	if f.ProcessLogWriter != nil {
		_, err = f.ProcessLogWriter.Write(formatted)
		if err != nil {
			return nil, err
		}
	}

	// Return the original message without log decoration
	// We don't want decoration to show up in the target creation logs
	return []byte(entry.Message + "\n"), nil
}
