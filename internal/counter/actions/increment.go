package actions

import (
	"context"

	"github.com/lenra-io/counter/internal/counter/services"
	"github.com/lenra-io/counter/pkg/lenra"
)

type IncrementProps struct {
	Id string
}

type IncrementAction struct {
	lenra.BaseAction
	Props IncrementProps
}

var _ lenra.Action = &IncrementAction{}

// Implementation
func (*IncrementAction) Name() string {
	return "increment"
}

func (a *IncrementAction) Perform(ctx context.Context) error {
	service := services.NewCounterService(a.ApiData)
	return service.Increment(ctx, a.Props.Id)
}
