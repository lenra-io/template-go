package widgets

func Root() map[string]interface{} {
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
	}
}
