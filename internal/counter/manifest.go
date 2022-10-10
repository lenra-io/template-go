package counter

import (
	"github.com/lenra-io/counter/internal/counter/actions"
	"github.com/lenra-io/counter/internal/counter/resources"
	"github.com/lenra-io/counter/internal/counter/widgets"
	"github.com/lenra-io/counter/pkg/lenra"
)

type Manifest struct {
	lenra.BaseManifest
}

var _ lenra.Manifest = &Manifest{}

func (m *Manifest) Widgets() []lenra.Widget {
	return []lenra.Widget{
		&widgets.RootWidget{},
		&widgets.HomeWidget{},
		&widgets.MenuWidget{},
		&widgets.CounterWidget{},
	}
}

func (m *Manifest) Actions() []lenra.Action {
	return []lenra.Action{
		&actions.OnEnvStartAction{},
		&actions.OnSessionStartAction{},
		&actions.OnUserFirstJoinAction{},
		&actions.IncrementAction{},
	}
}

func (m *Manifest) Resources() []lenra.Resource {
	return []lenra.Resource{
		&resources.LogoResource{},
	}
}
