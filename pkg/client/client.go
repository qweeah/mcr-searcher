package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const endpoint = "https://mcr.microsoft.com"

func Catalog() ([]string, error) {
	resp, err := http.Get(fmt.Sprintf("%s/v2/_catalog", endpoint))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Repositories []string `json:"repositories"`
	}
	decoder := json.NewDecoder(resp.Body)
	if err = decoder.Decode(&result); err != nil {
		return nil, err
	}

	return result.Repositories, nil
}

func TagList(repo string) ([]string, error) {
	resp, err := http.Get(fmt.Sprintf("%s/v2/%s/tags/list", endpoint, repo))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Tags []string `json:"tags"`
	}
	decoder := json.NewDecoder(resp.Body)
	if err = decoder.Decode(&result); err != nil {
		return nil, err
	}

	return result.Tags, nil
}
