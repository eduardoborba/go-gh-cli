package logic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// type IssueFetcher struct {
// 	BaseURL string
// }

type Issue struct {
	ID  int64  `json:"id"`
	Url string `json:"url"`
}

func FetchIssues(baseUrl string) ([]Issue, error) {
	url := fmt.Sprintf("%s/repos/rails/rails/issues", baseUrl)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Add("Accept", "application/json")
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var issues []Issue
	json.Unmarshal(body, &issues)

	return issues, nil
}
