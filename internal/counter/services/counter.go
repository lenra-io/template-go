package services

import (
	"context"
	"encoding/json"

	"github.com/lenra-io/counter/internal/counter/repo"
	"github.com/lenra-io/counter/pkg/lenra"
)

const (
	currentUser       = "@me"
	globalUser        = "global"
	counterCollection = "counter"
)

type counter struct {
	Count uint32 `json:"count"`
	User  string `json:"user"`
}

type CounterService struct {
	apiData lenra.Api
}

func NewCounterService(apiData lenra.Api) *CounterService {
	return &CounterService{apiData: apiData}
}

func (s *CounterService) CreateGlobalUserCounter(ctx context.Context) error {
	return s.createCounter(ctx, globalUser)
}
func (s *CounterService) CreateCurrentUserCounter(ctx context.Context) error {
	return s.createCounter(ctx, currentUser)
}

func (s *CounterService) createCounter(ctx context.Context, user string) error {
	repo := repo.NewDataRepo(s.apiData.Url, s.apiData.Token)
	_, err := repo.CreateDoc(ctx, counterCollection, counter{Count: 0, User: user})
	return err
}

func (s *CounterService) Increment(ctx context.Context, userId string) error {
	repo := repo.NewDataRepo(s.apiData.Url, s.apiData.Token)
	data, err := repo.GetDoc(ctx, counterCollection, userId)

	counter := &counter{}
	err = json.Unmarshal(data, counter)
	if err != nil {
		return err
	}
	counter.Count += 1
	_, err = repo.UpdateDoc(ctx, counterCollection, userId, counter)
	if err != nil {
		return err
	}
	return nil
}

// Widgets part
func CurrentUserWidgetQuery() map[string]interface{} {
	return map[string]interface{}{"user": currentUser}
}

func GlobalUserWidgetQuery() map[string]interface{} {
	return map[string]interface{}{"user": globalUser}
}
func Collection() string {
	return counterCollection
}
