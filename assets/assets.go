package assets

import (
	_ "embed"
	"fmt"
)

//go:embed logo.png
var logo []byte

func GetByName(name string) ([]byte, error) {
	switch name {
	case "logo.png":
		return logo, nil
	default:
		return nil, fmt.Errorf("Resource not found!")
	}
}
