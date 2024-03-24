package app

import (
	"context"
	"fmt"
	"log"

	"github.com/AxiomSamarth/gcm/example-06/pkg/server"
	"github.com/AxiomSamarth/gcm/example-06/pkg/watcher"

	v1 "k8s.io/api/core/v1"
)

const (
	configMapName      = "server-config"
	configMapNamespace = "default"
)

// Service represents the main application service which coordinates server and watcher.
type Service struct {
	server           *server.Server
	configMapWatcher *watcher.ConfigMapWatcher
}

// NewService creates a new instance of the application service.
func NewService() (*Service, error) {
	// create a new configmap watcher
	configMapWatcher, err := watcher.NewConfigMapWatcher(configMapName, configMapNamespace)
	if err != nil {
		return nil, err
	}

	// create a new instance of server
	server := server.NewServer(context.Background(), configMapWatcher.ConfigMap.Data["host"], configMapWatcher.ConfigMap.Data["port"])
	return &Service{
		configMapWatcher: configMapWatcher,
		server:           server,
	}, nil
}

// Run starts the application service.
func (s *Service) Run() {
	go func() {
		s.server.Start()
	}()

	s.processWatcherEvent()
}

// processWatcherEvent watches upon the K8s resources (configmap in this case) & captures
// the events on those resources. Based on change in the object, the respective components
// are modified to reflect the changes.
//
// Example: In this example program, if there is a change in the server's address in the configMap,
// this processWatcherEvent() method restarts the HTTP Server with the updated address without
// any manual intervention of the user.
func (s *Service) processWatcherEvent() {
	var err error

	for {
		select {
		case object, ok := <-s.configMapWatcher.Watcher:
			if !ok {
				log.Println("resetting the closed configmap watcher....")
				s.configMapWatcher, err = watcher.NewConfigMapWatcher(configMapName, configMapNamespace)
				if err != nil {
					log.Printf("error creating new configmap watcher for the closed one: %s", err.Error())
				}
			} else {
				newCM := object.Object.(*v1.ConfigMap)
				newAddr := fmt.Sprintf("%s:%s", newCM.Data["host"], newCM.Data["port"])

				if s.server.Addr != newAddr {
					log.Printf("server config changes observed. Restarting the server with new address %s", newAddr)
					s.server.Addr = newAddr
					if err := s.server.Stop(); err != nil {
						log.Printf("error shutting down the server %s", err.Error())
					}
					if err := s.server.Start(); err != nil {
						log.Printf("error starting the server: %s", err.Error())
					}
				}
			}
		// Add more watchers here if there are any
		default:
			continue
		}
	}
}
