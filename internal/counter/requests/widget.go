package requests

import (
	"encoding/json"

	"github.com/lenra-io/counter/internal/counter/util"
	"github.com/lenra-io/counter/internal/counter/widgets"
	"github.com/sirupsen/logrus"
)

type Widget struct {
	Widget string
	Raw    []byte
}

type CounterWidget struct {
	Widget string
	Data   []CounterWidgetData
	Props  CounterWidgetProps
}

type CounterWidgetData struct {
	Id    string `json:"_id"`
	Count int32
	User  string
}

type CounterWidgetProps struct {
	Text string
}

// Entry point
func HandleWidgetRequest(widget *Widget) interface{} {
	switch widget.Widget {
	case "root":
		return widgets.Root()
	case "menu":
		return widgets.Menu()
	case "home":
		return widgets.Home()
	case "counter":
		return HandleCounterWidgetRequest(widget)
	default:
		logrus.Errorf("Widget request is malformed: %s", string(widget.Raw))
		return util.Manifest()
	}
}

func HandleCounterWidgetRequest(widget *Widget) interface{} {
	counterWidget := &CounterWidget{}
	err := json.Unmarshal(widget.Raw, counterWidget)
	if err != nil {
		logrus.Errorf("CounterWidget request is malformed: %s, with error: %s", string(widget.Raw), err)
		return util.Manifest()
	}
	// Data exists
	if len(counterWidget.Data) > 0 {
		return widgets.Counter(counterWidget.Data[0].Id, counterWidget.Data[0].Count, counterWidget.Props.Text)
	}
	// Not exists - showing loading instead
	return widgets.Loading()
}
