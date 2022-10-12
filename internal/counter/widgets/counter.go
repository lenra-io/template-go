package widgets

import (
	"context"
	"fmt"

	"github.com/lenra-io/counter/pkg/lenra"
)

type CounterWidget struct {
	lenra.BaseWidget
	Data  []CounterWidgetData
	Props CounterWidgetProps
}

type CounterWidgetData struct {
	Id    string `json:"_id"`
	Count int32
	User  string
}

type CounterWidgetProps struct {
	Text string `json:"text"`
}

// Ensure that widget follows the interface
var _ lenra.Widget = CounterWidget{}

// Implementation
func (w CounterWidget) Name() string { return "counter" }

func (w CounterWidget) Render(ctx context.Context) (interface{}, error) {
	if len(w.Data) > 0 {
		return map[string]interface{}{
			"type":               "flex",
			"spacing":            2,
			"mainAxisAlignment":  "spaceEvenly",
			"crossAxisAlignment": "center",
			"children": []interface{}{
				map[string]interface{}{
					"type":  "text",
					"value": fmt.Sprintf("%s: %d", w.Props.Text, w.Data[0].Count),
				},
				map[string]interface{}{
					"type": "button",
					"text": "+",
					"onPressed": map[string]interface{}{
						"action": "increment",
						"props": map[string]interface{}{
							"id": w.Data[0].Id,
						},
					},
				},
			},
		}, nil
	} else {
		return map[string]interface{}{
			"type": "widget",
			"name": "loading",
		}, nil
	}
}
