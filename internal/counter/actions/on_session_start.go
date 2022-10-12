package actions

import (
	"github.com/lenra-io/counter/pkg/lenra"
)

type OnSessionStartAction struct{ lenra.BaseAction }

var _ lenra.Action = &OnSessionStartAction{}

// Implementation
func (*OnSessionStartAction) Name() string {
	return "onSessionStart"
}
