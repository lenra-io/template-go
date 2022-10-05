package widgets

import "github.com/lenra-io/counter/internal/counter/util"

func Home() map[string]interface{} {
	return map[string]interface{}{
		"type":               "flex",
		"direction":          "vertical",
		"spacing":            4,
		"mainAxisAlignment":  "spaceEvenly",
		"crossAxisAlignment": "center",
		"children": []interface{}{
			map[string]interface{}{
				"type": "widget",
				"name": "counter",
				"coll": util.COUNTER_COLLECTION,
				"query": map[string]interface{}{
					"user": util.CURRENT_USER,
				},
				"props": CounterWidgetProps{Text: "My personal counter"},
			},
			map[string]interface{}{
				"type": "widget",
				"name": "counter",
				"coll": util.COUNTER_COLLECTION,
				"query": map[string]interface{}{
					"user": util.GLOBAL_USER,
				},
				"props": CounterWidgetProps{Text: "The common counter"},
			},
		},
	}
}
