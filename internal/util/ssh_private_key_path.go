// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package util

import (
	"fmt"
	"os"

	"golang.org/x/crypto/ssh"
	"golang.org/x/term"
)

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func GetSshPrivateKeyPath(privateKeyPath string) (string, *string, error) {
	keyContent, err := os.ReadFile(privateKeyPath)
	if err != nil {
		return "", nil, err
	}

	_, err = ssh.ParsePrivateKey(keyContent)
	if err == nil {
		return privateKeyPath, nil, err
	}

	if err.Error() == (&ssh.PassphraseMissingError{}).Error() {
		fmt.Print("Enter password for key: ")
		password, err := term.ReadPassword(0)
		fmt.Println()
		if err != nil {
			return "", nil, err
		}

		stringPassword := string(password)

		return privateKeyPath, &stringPassword, nil
	}

	return "", nil, err
}
