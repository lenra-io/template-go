package requests

import (
	"github.com/lenra-io/counter/assets"
	"github.com/sirupsen/logrus"
)

type Resource struct {
	Resource string
	Raw      []byte
}

func HandleResourceRequest(resource *Resource) ([]byte, error) {
	res, err := assets.GetByName(resource.Resource)
	if err != nil {
		logrus.Error("Resource not found: %s", resource.Resource)
	}
	return res, err
}
