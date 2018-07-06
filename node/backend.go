// Package node defines a backend for a beacon chain node and all its
// associated services. It defines a struct which handles each service
// lifecycle in the node.
package node

import (
	"fmt"
	"os"
	"os/signal"
	"reflect"
	"sync"
	"syscall"

	"github.com/ethereum/go-ethereum/log"
	"github.com/prysmaticlabs/beacon-chain/types"
	"github.com/urfave/cli"
)

// BeaconNode is a service that is registered and started when the system is launched.
// it contains APIs and fields that handle the different components of the Ethereum 2.0
// protocol.
type BeaconNode struct {
	services     map[reflect.Type]types.Service // Service registry.
	serviceTypes []reflect.Type                 // Keeps an ordered slice of registered service types.
	lock         sync.RWMutex
	stop         chan struct{} // Channel to wait for termination notifications
}

// New creates a new beacon chain node instance. This is called in the main
// geth sharding entrypoint.
func New(ctx *cli.Context) (*BeaconNode, error) {
	beaconNode := &BeaconNode{
		services: make(map[reflect.Type]types.Service),
		stop:     make(chan struct{}),
	}
	return beaconNode, nil
}

// Start the BeaconNode service and kicks off each service's main loop.
func (s *BeaconNode) Start() {
	s.lock.Lock()

	log.Info("Starting sharding node")

	for _, kind := range s.serviceTypes {
		// Start each service in order of registration.
		s.services[kind].Start()
	}

	stop := s.stop
	s.lock.Unlock()

	go func() {
		sigc := make(chan os.Signal, 1)
		signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
		defer signal.Stop(sigc)
		<-sigc
		log.Info("Got interrupt, shutting down...")
		go s.Close()
		for i := 10; i > 0; i-- {
			<-sigc
			if i > 1 {
				log.Warn("Already shutting down, interrupt more to panic.", "times", i-1)
			}
		}
		panic("Panic closing the beacon chain node")
	}()

	// Wait for stop channel to be closed
	<-stop
}

// Close handles graceful shutdown of the system.
func (s *BeaconNode) Close() {
	s.lock.Lock()
	defer s.lock.Unlock()

	for kind, service := range s.services {
		if err := service.Stop(); err != nil {
			log.Crit(fmt.Sprintf("Could not stop the following service: %v, %v", kind, err))
		}
	}
	log.Info("Stopping beacon chain node")

	// unblock n.Wait
	close(s.stop)
}

// registerService appends a service constructor function to the service registry of the
// sharding node.
func (s *BeaconNode) registerService(service types.Service) error {
	kind := reflect.TypeOf(service)
	if _, exists := s.services[kind]; exists {
		return fmt.Errorf("service already exists: %v", kind)
	}
	s.services[kind] = service
	s.serviceTypes = append(s.serviceTypes, kind)
	return nil
}

// fetchService takes in a struct pointer and sets the value of that pointer
// to a service currently stored in the service registry. This ensures the input argument is
// set to the right pointer that refers to the originally registered service.
func (s *BeaconNode) fetchService(service interface{}) error {
	if reflect.TypeOf(service).Kind() != reflect.Ptr {
		return fmt.Errorf("input must be of pointer type, received value type instead: %T", service)
	}
	element := reflect.ValueOf(service).Elem()
	if running, ok := s.services[element.Type()]; ok {
		element.Set(reflect.ValueOf(running))
		return nil
	}
	return fmt.Errorf("unknown service: %T", service)
}
