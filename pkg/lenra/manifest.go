package lenra

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
)

type Manifest interface {
	Widgets() []Widget
	Actions() []Action
	Resources() []Resource
}
type BaseManifest struct{}

var _ Manifest = &BaseManifest{}

// Base realisation
func (*BaseManifest) Widgets() []Widget     { return []Widget{} }
func (*BaseManifest) Actions() []Action     { return []Action{} }
func (*BaseManifest) Resources() []Resource { return []Resource{} }

func Serve(ctx context.Context, m Manifest, request []byte) {
	logrus.Debugf("Handling request: %s", string(request))
	// var f interface{}
	var err error

	// Starting with widgets
	widget := &struct {
		Widget string
	}{}
	err = json.Unmarshal(request, widget)
	if err == nil && widget.Widget != "" {
		renderedWidget, err := serveWidget(ctx, m, widget.Widget, request)
		if err != nil {
			render(renderManifest(), true)
			return
		} else {
			render(renderedWidget, true)
			return
		}
	}

	// Moving to actions
	action := &struct {
		Action string
	}{}
	err = json.Unmarshal(request, action)
	if err == nil && action.Action != "" {
		err := serveAction(ctx, m, action.Action, request)
		if err != nil {
			logrus.Error("Action finished with error: %s", err)
		}
		render(nil, true)
		return
	}

	// Moving to resource
	resource := &struct {
		Resource string
	}{}
	err = json.Unmarshal(request, resource)
	if err == nil && resource.Resource != "" {
		rawResource, err := serveResource(ctx, m, resource.Resource, request)
		if err != nil {
			render(renderManifest(), true)
			return
		} else {
			render(rawResource, false)
			return
		}
	}

	// No valid
	logrus.Errorf("Request is not valid: %s", string(request))
	render(renderManifest(), true)
}

func serveWidget(ctx context.Context, m Manifest, widgetName string, request []byte) (interface{}, error) {
	var err error
	for _, widget := range m.Widgets() {
		if widget.Name() == widgetName {
			err = json.Unmarshal(request, widget)
			if err != nil {
				logrus.Error("\"%s\" widget request is malformed: %s, with error %s", widgetName, string(request), err)
				return nil, err
			}
			return widget.Render(ctx)
		}
	}
	logrus.Debug("Widget %s is not found", widgetName)
	return nil, fmt.Errorf("widget %s is not found", widgetName)
}

func serveAction(ctx context.Context, m Manifest, actionName string, request []byte) error {
	var err error
	for _, action := range m.Actions() {
		if action.Name() == actionName {
			err = json.Unmarshal(request, action)
			if err != nil {
				logrus.Error("\"%s\" action request is malformed: %s, with error %s", actionName, string(request), err)
				return err
			}
			return action.Perform(ctx)
		}
	}
	logrus.Debug("Action %s is not found", actionName)
	return fmt.Errorf("action %s is not found", actionName)
}

func serveResource(ctx context.Context, m Manifest, resourceName string, request []byte) ([]byte, error) {
	for _, resource := range m.Resources() {
		if resource.Name() == resourceName {
			return resource.Data(ctx)
		}
	}
	logrus.Debug("Resource %s is not found", resourceName)
	return nil, fmt.Errorf("widget %s is not found", resourceName)
}

func renderManifest() interface{} {
	return map[string]interface{}{
		"manifest": map[string]interface{}{
			"rootWidget": "root",
		},
	}
}

func render(output interface{}, marshal bool) {
	if marshal {
		json_output, err := json.Marshal(output)
		if err != nil {
			logrus.Errorf("Internal response return is malformed: %v", output)
		}
		fmt.Printf("%s", string(json_output))
	} else {
		fmt.Printf("%s", output)
	}
}
