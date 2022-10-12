package lenra

import "context"

type Api struct {
	Url   string
	Token string
}

type Action interface {
	// Secret function to diverse interfaces
	thisFunctionEnsuresThatActionEmbedsBaseAction()

	// Main functions
	Name() string
	Perform(ctx context.Context) error
}

var _ Action = &BaseAction{}

type BaseAction struct {
	ApiData Api `json:"api"`
}

func (*BaseAction) thisFunctionEnsuresThatActionEmbedsBaseAction() {}

func (a *BaseAction) Name() string { panic("unimplemented") }

// Perform implements Action
func (a *BaseAction) Perform(ctx context.Context) error {
	// Do nothing
	return nil
}
