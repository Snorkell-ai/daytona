// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package tailscale

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/daytonaio/daytona/pkg/common"
	"tailscale.com/tsnet"
)

var conn *tsnet.Server = nil

type TsnetConnConfig struct {
	AuthKey    string
	ControlURL string
	Dir        string
	Logf       func(format string, args ...any)
	Hostname   string
}

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func GetConnection(config *TsnetConnConfig) (*tsnet.Server, error) {
	if conn != nil {
		return conn, nil
	}

	if config == nil {
		return nil, errors.New("connection not initialized")
	}

	conn = &tsnet.Server{
		AuthKey:    config.AuthKey,
		ControlURL: config.ControlURL,
		Dir:        config.Dir,
		Logf:       config.Logf,
		Hostname:   config.Hostname,
		UserLogf:   config.Logf,
		Ephemeral:  true,
	}

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := conn.Up(timeoutCtx)
	if err != nil {
		return nil, fmt.Errorf("%w. %w", err, common.ErrConnection)
	}

	return conn, nil
}
