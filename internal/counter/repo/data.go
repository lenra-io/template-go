package repo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/sirupsen/logrus"
)

type DataRepo struct {
	baseUrl   string
	authToken string
	client    *http.Client
}

func NewDataRepo(baseUrl, authToken string) *DataRepo {
	return &DataRepo{
		baseUrl:   baseUrl,
		authToken: authToken,
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// API
func (dr *DataRepo) GetDoc(ctx context.Context, collection string, id string) ([]byte, error) {
	logrus.Debugf("Getting document: %s from collection: %s", id, collection)

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		dr.getUrl(fmt.Sprintf("/app/colls/%s/docs/%s", collection, id)),
		nil,
	)
	if err != nil {
		return nil, err
	}
	res, err := dr.signAndSend(req)
	if err != nil {
		return nil, err
	}
	doc, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	res.Body.Close()
	return doc, nil
}

func (dr *DataRepo) CreateDoc(ctx context.Context, collection string, document interface{}) (interface{}, error) {
	data, err := json.Marshal(document)
	if err != nil {
		return nil, err
	}
	logrus.Debugf("Creating document: %s in collection: %s", string(data), collection)

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		dr.getUrl(fmt.Sprintf("/app/colls/%s/docs", collection)),
		bytes.NewBuffer(data),
	)
	if err != nil {
		return nil, err
	}

	res, err := dr.signAndSend(req)
	if err != nil {
		return nil, err
	}

	var body interface{}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return nil, err
	}
	logrus.Debug("Creating document response body", body)
	return body, nil
}

func (dr *DataRepo) UpdateDoc(ctx context.Context, collection string, id string, document interface{}) (interface{}, error) {
	data, err := json.Marshal(document)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPut,
		dr.getUrl(fmt.Sprintf("/app/colls/%s/docs/%s", collection, id)),
		bytes.NewBuffer(data),
	)
	if err != nil {
		return nil, err
	}

	res, err := dr.signAndSend(req)
	if err != nil {
		return nil, err
	}

	var body interface{}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// Internals

func (dr *DataRepo) getUrl(route string) string {
	url, err := url.JoinPath(dr.baseUrl, route)
	if err != nil {
		logrus.Panicf("Internal error with url concat: %s, %s", dr.baseUrl, route)
	}
	return url
}

func (dr *DataRepo) signAndSend(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", dr.authToken))
	req.Header.Add("Content-Type", "application/json")

	return dr.client.Do(req)
}
