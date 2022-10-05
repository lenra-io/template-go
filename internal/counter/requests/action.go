package requests

import (
	"encoding/json"
	"fmt"

	"github.com/lenra-io/counter/internal/counter/repo"
	"github.com/lenra-io/counter/internal/counter/util"
	"github.com/sirupsen/logrus"
)

type Action struct {
	Action string
	Api    Api
	Raw    []byte
}

type IncrementAction struct {
	Action string
	Api    Api
	Props  IncrementProps
}

type IncrementProps struct {
	Id string
}

type Api struct {
	Url   string
	Token string
}

type Counter struct {
	Count uint32 `json:"count"`
	User  string `json:"user"`
}

// Entry Point
func HandleActionRequest(action *Action) error {
	switch action.Action {
	case "onEnvStart":
		return HandleOnEnvStartAction(action)
	case "onSessionStart":
		// Do nothing
		return nil
	case "onUserFirstJoin":
		return HandleOnUserFirstJoinAction(action)
	case "increment":
		return HandleIncrementAction(action)
	default:
		logrus.Errorf("Action request is malformed: %s", string(action.Raw))
		return fmt.Errorf("Action is malformed!")
	}
}

func HandleOnEnvStartAction(action *Action) error {
	return createCounter(action.Api, util.GLOBAL_USER)
}

func HandleOnUserFirstJoinAction(action *Action) error {
	return createCounter(action.Api, util.CURRENT_USER)
}

func HandleIncrementAction(action *Action) error {
	incrementAction := &IncrementAction{}
	err := json.Unmarshal(action.Raw, incrementAction)
	if err != nil {
		logrus.Errorf("IncrementAction request is malformed: %s, with error: %s", string(action.Raw), err)
		return err
	}

	repo := repo.NewDataRepo(incrementAction.Api.Url, incrementAction.Api.Token)
	data, err := repo.GetDoc(util.COUNTER_COLLECTION, incrementAction.Props.Id)

	counter := &Counter{}
	err = json.Unmarshal(data, counter)
	if err != nil {
		return err
	}
	counter.Count += 1
	_, err = repo.UpdateDoc(util.COUNTER_COLLECTION, incrementAction.Props.Id, counter)
	if err != nil {
		return err
	}
	return nil
}

func createCounter(api Api, user string) error {
	repo := repo.NewDataRepo(api.Url, api.Token)
	_, err := repo.CreateDoc(util.COUNTER_COLLECTION, Counter{Count: 0, User: user})
	return err
}
