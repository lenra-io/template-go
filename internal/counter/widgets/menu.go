package widgets

import (
	"context"

	"github.com/lenra-io/counter/pkg/lenra"
	"github.com/lenra-io/counter/pkg/lenra/layout"
)

type MenuWidget struct {
	lenra.BaseWidget
}

// Ensure that widget follows the interface
var _ lenra.Widget = MenuWidget{}

// Implementation
func (w MenuWidget) Name() string { return "menu" }

func (w MenuWidget) Render(ctx context.Context) (interface{}, error) {
	menu_content := map[string]interface{}{
		"type":               "flex",
		"fillParent":         true,
		"mainAxisAlignment":  "spaceBetween",
		"crossAxisAlignment": "center",
		"padding":            layout.Padding{Right: 4},
		"children": []interface{}{
			map[string]interface{}{
				"type": "container",
				"constraints": map[string]interface{}{
					"minWidth":  32,
					"minHeight": 32,
					"maxWidth":  32,
					"maxHeight": 32,
				},
				"child": map[string]interface{}{
					"type": "image",
					"src":  "logo.png",
				},
			},
			map[string]interface{}{
				"type": "flexible",
				"child": map[string]interface{}{
					"type": "container",
					"child": map[string]interface{}{
						"type":      "text",
						"value":     "Hello World",
						"textAlign": "center",
						"style": map[string]interface{}{
							"fontWeight": "bold",
							"fontSize":   24,
						},
					},
				},
			},
		},
	}

	return map[string]interface{}{
		"type": "container",
		"decoration": layout.Decoration{
			Color: 0xFFFFFFFF,
			BoxShadow: layout.BoxShadow{
				BlurRadius: 8,
				Color:      0x1A000000,
				Offset: layout.Offset{
					Dx: 0,
					Dy: 1,
				},
			},
		},
		"padding": layout.Padding{Top: 2, Bottom: 2, Left: 4, Right: 4},
		"child":   menu_content,
	}, nil
}
