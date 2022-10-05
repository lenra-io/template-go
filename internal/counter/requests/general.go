package requests

import (
	"encoding/json"

	"github.com/lenra-io/counter/internal/counter/util"
	"github.com/sirupsen/logrus"
)

// Root entry point
func HandleRootRequest(request []byte) (interface{}, bool) {
	logrus.Debugf("Handling request: %s", string(request))
	// var f interface{}
	var err error

	// Starting with widgets
	widget := &Widget{}
	err = json.Unmarshal(request, widget)
	if err == nil && widget.Widget != "" {
		widget.Raw = request
		return HandleWidgetRequest(widget), true
	}

	// Moving to actions
	action := &Action{}
	err = json.Unmarshal(request, action)
	if err == nil && action.Action != "" {
		action.Raw = request
		err := HandleActionRequest(action)
		if err != nil {
			logrus.Error("Action finished with error: %s", err)
		}
		return nil, true
	}

	// Moving to resource
	resource := &Resource{}
	err = json.Unmarshal(request, resource)
	if err == nil && resource.Resource != "" {
		resource.Raw = request
		raw_resource, err := HandleResourceRequest(resource)
		if err != nil {
			return util.Manifest(), true
		} else {
			return raw_resource, false
		}
	}

	// No valid
	logrus.Errorf("Request is not valid: %s", string(request))
	return util.Manifest(), true
}
