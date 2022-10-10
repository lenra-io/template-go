package lenra

import "context"

type Widget interface {
	// Secret function to diverse interfaces
	thisFunctionEnsuresThatWidgetEmbedsBaseWidget()

	// Main functions
	Render(ctx context.Context) (interface{}, error)
	Name() string
}

var _ Widget = &BaseWidget{}

// Default widget
type BaseWidget struct{}

func (w BaseWidget) Name() string { return "#_defaultWidget" }
func (w BaseWidget) Render(ctx context.Context) (interface{}, error) {
	return map[string]interface{}{}, nil
}
func (w BaseWidget) thisFunctionEnsuresThatWidgetEmbedsBaseWidget() {}
