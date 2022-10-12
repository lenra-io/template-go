package resources

import (
	"context"

	"github.com/lenra-io/counter/assets"
	"github.com/lenra-io/counter/pkg/lenra"
)

type LogoResource struct{ lenra.BaseResource }

var _ lenra.Resource = &LogoResource{}

// Data implements lenra.Resource
func (r *LogoResource) Data(ctx context.Context) ([]byte, error) {
	return assets.GetByName(r.Name())
}

// Name implements lenra.Resource
func (r *LogoResource) Name() string {
	return "logo.png"
}
