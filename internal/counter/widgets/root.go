package widgets

import (
	"context"

	"github.com/lenra-io/counter/pkg/lenra"
)

type RootWidget struct {
	lenra.BaseWidget
}

// Ensure that widget follows the interface
var _ lenra.Widget = &RootWidget{}

// Implementation
func (w RootWidget) Name() string { return "root" }

func (w RootWidget) Render(ctx context.Context) (interface{}, error) {
	return map[string]interface{}{
		"type":               "flex",
		"direction":          "vertical",
		"scroll":             true,
		"spacing":            4,
		"crossAxisAlignment": "center",
		"children": []interface{}{
			map[string]interface{}{
				"type": "widget",
				"name": "menu",
			},
			map[string]interface{}{
				"type": "widget",
				"name": "home",
			},
		},
	}, nil
}
