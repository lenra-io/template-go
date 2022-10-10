package widgets

import (
	"context"

	"github.com/lenra-io/counter/pkg/lenra"
)

type LoadingWidget struct {
	lenra.BaseWidget
}

// Ensure that widget follows the interface
var _ lenra.Widget = LoadingWidget{}

// Implementation
func (w LoadingWidget) Render(ctx context.Context) (interface{}, error) {
	return map[string]interface{}{
		"type":  "text",
		"value": "Loading...",
	}, nil
}
