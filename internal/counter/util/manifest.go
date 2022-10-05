package util

func Manifest() map[string]interface{} {
	return map[string]interface{}{
		"manifest": map[string]interface{}{
			"widgets":    []interface{}{"root"},
			"listeners":  []interface{}{},
			"rootWidget": "root",
		},
	}
}
