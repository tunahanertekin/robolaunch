package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type Chart struct {
	Name        string   `json:"name"`
	Version     string   `json:"version"`
	Description string   `json:"description"`
	ApiVersion  string   `json:"apiVersion"`
	AppVersion  string   `json:"appVersion"`
	Type        string   `json:"type"`
	Urls        []string `json:"urls"`
	Created     string   `json:"created"`
	Digest      string   `json:"digest"`
}

const ChartMuseumHost = "HOST"

func GetLaunches() (map[string][]Chart, error) {

	resp, err := http.Get(ChartMuseumHost + "/api/charts")
	if err != nil {
		return nil, err
	}
	var launches map[string][]Chart
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &launches)
	if err != nil {
		return nil, err
	}

	return launches, nil
}
