package lenra

import "context"

type Resource interface {
	// Secret function to diverse interfaces
	thisFunctionEnsuresThatResourceEmbedsBaseResource()

	// Main functions
	Name() string
	Data(ctx context.Context) ([]byte, error)
}

var _ Resource = &BaseResource{}

type BaseResource struct{}

// Data implements Resource
func (*BaseResource) Data(ctx context.Context) ([]byte, error) {
	panic("unimplemented")
}

// Name implements Resource
func (*BaseResource) Name() string {
	panic("unimplemented")
}

// thisFunctionEnsuresThatResourceEmbedsBaseResource implements Resource
func (*BaseResource) thisFunctionEnsuresThatResourceEmbedsBaseResource() {}
