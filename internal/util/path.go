// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package util

import (
	"os/exec"
	"path"
	"strings"

	"github.com/daytonaio/daytona/cmd/daytona/config"
)

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func GetHomeDir(activeProfile config.Profile, workspaceId string, gpgKey *string) (string, error) {
	err := config.EnsureSshConfigEntryAdded(activeProfile.Id, workspaceId, gpgKey)
	if err != nil {
		return "", err
	}

	workspaceHostname := config.GetHostname(activeProfile.Id, workspaceId)

	homeDir, err := exec.Command("ssh", workspaceHostname, "echo", "$HOME").Output()
	if err != nil {
		return "", err
	}

	return strings.TrimRight(string(homeDir), "\n"), nil
}

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func GetWorkspaceDir(activeProfile config.Profile, workspaceId, repoName string, gpgKey *string) (string, error) {
	err := config.EnsureSshConfigEntryAdded(activeProfile.Id, workspaceId, gpgKey)
	if err != nil {
		return "", err
	}

	workspaceHostname := config.GetHostname(activeProfile.Id, workspaceId)

	daytonaWorkspaceDir, err := exec.Command("ssh", workspaceHostname, "echo", "$DAYTONA_WORKSPACE_DIR").Output()
	if err != nil {
		return "", err
	}

	if strings.TrimRight(string(daytonaWorkspaceDir), "\n") != "" {
		return strings.TrimRight(string(daytonaWorkspaceDir), "\n"), nil
	}

	homeDir, err := GetHomeDir(activeProfile, workspaceId, gpgKey)
	if err != nil {
		return "", err
	}

	return path.Join(homeDir, repoName), nil
}
