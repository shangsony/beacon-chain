package node

import (
	"reflect"
	"testing"

	"github.com/urfave/cli"
)

type mockService struct{}
type secondMockService struct{}

func (m *mockService) Start() {
	return
}

func (m *mockService) Stop() error {
	return nil
}

func (s *secondMockService) Start() {
	return
}

func (s *secondMockService) Stop() error {
	return nil
}

func TestRegisterServiceTwice(t *testing.T) {
	b, err := New(&cli.Context{})
	if err != nil {
		t.Fatalf("failed to setup node: %v", err)
	}
	m := &mockService{}
	if err := b.registerService(m); err != nil {
		t.Fatalf("failed to register first service")
	}

	// checks if first service was indeed registered
	if len(b.serviceTypes) != 1 {
		t.Fatalf("service types slice should contain 1 service, contained %v", len(b.serviceTypes))
	}

	if err := b.registerService(m); err == nil {
		t.Errorf("should not be able to register a service twice, got nil error")
	}
}

func TestRegisterDifferentServices(t *testing.T) {
	b, err := New(&cli.Context{})
	if err != nil {
		t.Fatalf("failed to setup node: %v", err)
	}
	m := &mockService{}
	s := &secondMockService{}
	if err := b.registerService(m); err != nil {
		t.Fatalf("failed to register first service")
	}

	if err := b.registerService(s); err != nil {
		t.Fatalf("failed to register second service")
	}

	if len(b.serviceTypes) != 2 {
		t.Errorf("service types slice should contain 2 services, contained %v", len(b.serviceTypes))
	}

	if _, exists := b.services[reflect.TypeOf(m)]; !exists {
		t.Errorf("service of type %v not registered", reflect.TypeOf(m))
	}

	if _, exists := b.services[reflect.TypeOf(s)]; !exists {
		t.Errorf("service of type %v not registered", reflect.TypeOf(s))
	}
}

func TestFetchService(t *testing.T) {
	b, err := New(&cli.Context{})
	if err != nil {
		t.Fatalf("failed to setup node: %v", err)
	}
	m := &mockService{}
	if err := b.registerService(m); err != nil {
		t.Fatalf("failed to register first service")
	}

	if err := b.fetchService(*m); err == nil {
		t.Errorf("passing in a value should throw an error, received nil error")
	}

	var s *secondMockService
	if err := b.fetchService(&s); err == nil {
		t.Errorf("fetching an unregistered service should return an error, got nil")
	}

	var m2 *mockService
	if err := b.fetchService(&m2); err != nil {
		t.Fatalf("failed to fetch service")
	}

	if m2 != m {
		t.Errorf("pointers were not equal, instead got %p, %p", m2, m)
	}
}
