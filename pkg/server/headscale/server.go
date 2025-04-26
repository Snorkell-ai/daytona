// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package headscale

import (
	"context"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/daytonaio/daytona/pkg/frpc"
	"github.com/daytonaio/daytona/pkg/server"
	"github.com/juanfont/headscale/hscontrol"

	log "github.com/sirupsen/logrus"
)

type HeadscaleServerConfig struct {
	ServerId      string
	FrpsDomain    string
	FrpsProtocol  string
	HeadscalePort uint32
	ConfigDir     string
	Frps          *server.FRPSConfig
}

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func NewHeadscaleServer(config *HeadscaleServerConfig) *HeadscaleServer {
	return &HeadscaleServer{
		serverId:      config.ServerId,
		frpsDomain:    config.FrpsDomain,
		frpsProtocol:  config.FrpsProtocol,
		headscalePort: config.HeadscalePort,
		configDir:     config.ConfigDir,
		frps:          config.Frps,
	}
}

type HeadscaleServer struct {
	serverId      string
	frpsDomain    string
	frpsProtocol  string
	headscalePort uint32
	configDir     string
	frps          *server.FRPSConfig

	stopChan       chan struct{}
	disconnectChan chan struct{}
}

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func (s *HeadscaleServer) Init() error {
	return os.MkdirAll(s.configDir, 0700)
}

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func (s *HeadscaleServer) Start(errChan chan error) error {
	// Check if port is already in use
	_, err := net.Dial("tcp", fmt.Sprintf(":%d", s.headscalePort))
	if err == nil {
		return fmt.Errorf("cannot start Headscale server, port %d is already in use", s.headscalePort)
	}

	cfg, err := s.getHeadscaleConfig()
	if err != nil {
		return err
	}

	app, err := hscontrol.NewHeadscale(cfg)
	if err != nil {
		return err
	}

	s.stopChan = make(chan struct{})

	go func() {
		select {
		case <-s.stopChan:
			s.disconnectChan <- struct{}{}
			errChan <- nil
			return
		case errChan <- app.Serve():
			return
		}
	}()

	if s.frps == nil {
		return err
	}

	healthCheck, frpcService, err := frpc.GetService(frpc.FrpcConnectParams{
		ServerDomain: s.frps.Domain,
		ServerPort:   int(s.frps.Port),
		Name:         fmt.Sprintf("daytona-server-%s", s.serverId),
		Port:         int(s.headscalePort),
		SubDomain:    s.serverId,
	})
	if err != nil {
		return err
	}

	go func() {
		err := frpcService.Run(context.Background())
		if err != nil {
			errChan <- err
		}
	}()

	for i := 0; i < 5; i++ {
		if err = healthCheck(); err != nil {
			log.Debugf("Failed to connect to headscale frpc: %s", err)
			time.Sleep(2 * time.Second)
		} else {
			break
		}
	}

	return err
}

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func (s *HeadscaleServer) Stop() error {
	go func() {
		s.stopChan <- struct{}{}
	}()

	return nil
}

// Sort sorts the input slice of integers using the QuickSort algorithm.
//
// Parameters:
//   arr []int: The slice of integers to be sorted.
//
// Returns:
//   []int: A new sorted slice containing the elements of arr in ascending order.
func (s *HeadscaleServer) Purge() error {
	return os.RemoveAll(s.configDir)
}
