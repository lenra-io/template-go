package actions

import (
	"context"

	"github.com/lenra-io/counter/internal/counter/services"
	"github.com/lenra-io/counter/pkg/lenra"
)

type OnUserFirstJoinAction struct{ lenra.BaseAction }

var _ lenra.Action = &OnUserFirstJoinAction{}

// Implementation
func (*OnUserFirstJoinAction) Name() string {
	return "onUserFirstJoin"
}

func (a *OnUserFirstJoinAction) Perform(ctx context.Context) error {
	service := services.NewCounterService(a.ApiData)
	return service.CreateCurrentUserCounter(ctx)
}
