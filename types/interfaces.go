package types

// Service is an individual protocol that can be registered into a node. Having a beacon
// node maintain a service registry allows for easy, shared-dependencies.
type Service interface {
	// Start is called after all services have been constructed to
	// spawn any goroutines required by the service.
	Start()
	// Stop terminates all goroutines belonging to the service,
	// blocking until they are all terminated.
	Stop() error
}
