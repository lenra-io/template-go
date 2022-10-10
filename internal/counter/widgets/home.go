package widgets

import (
	"context"

	"github.com/lenra-io/counter/internal/counter/services"
	"github.com/lenra-io/counter/pkg/lenra"
)

type HomeWidget struct {
	lenra.BaseWidget
}

// Ensure that widget follows the interface
var _ lenra.Widget = HomeWidget{}

// Implementation
func (w HomeWidget) Name() string { return "home" }

func (w HomeWidget) Render(ctx context.Context) (interface{}, error) {
	return map[string]interface{}{
		"type":               "flex",
		"direction":          "vertical",
		"spacing":            4,
		"mainAxisAlignment":  "spaceEvenly",
		"crossAxisAlignment": "center",
		"children": []interface{}{
			map[string]interface{}{
				"type":  "widget",
				"name":  "counter",
				"coll":  services.Collection(),
				"query": services.CurrentUserWidgetQuery(),
				"props": CounterWidgetProps{Text: "My personal counter"},
			},
			map[string]interface{}{
				"type":  "widget",
				"name":  "counter",
				"coll":  services.Collection(),
				"query": services.GlobalUserWidgetQuery(),
				"props": CounterWidgetProps{Text: "The common counter"},
			},
		},
	}, nil
}
