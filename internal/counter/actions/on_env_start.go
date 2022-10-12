package actions

import (
	"context"

	"github.com/lenra-io/counter/internal/counter/services"
	"github.com/lenra-io/counter/pkg/lenra"
)

type OnEnvStartAction struct{ lenra.BaseAction }

var _ lenra.Action = &OnEnvStartAction{}

// Implementation
func (*OnEnvStartAction) Name() string {
	return "onEnvStart"
}

func (a *OnEnvStartAction) Perform(ctx context.Context) error {
	service := services.NewCounterService(a.ApiData)
	return service.CreateGlobalUserCounter(ctx)
}
