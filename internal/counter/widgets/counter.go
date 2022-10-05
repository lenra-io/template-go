package widgets

import "fmt"

func Counter(id string, count int32, text string) map[string]interface{} {
	return map[string]interface{}{
		"type":               "flex",
		"spacing":            2,
		"mainAxisAlignment":  "spaceEvenly",
		"crossAxisAlignment": "center",
		"children": []interface{}{
			map[string]interface{}{
				"type":  "text",
				"value": fmt.Sprintf("%s: %d", text, count),
			},
			map[string]interface{}{
				"type": "button",
				"text": "+",
				"onPressed": map[string]interface{}{
					"action": "increment",
					"props": map[string]interface{}{
						"id": id,
					},
				},
			},
		},
	}
}
