package project

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/felipeornelis/todoist-api-go/url"
)

func (s Service) GetAll() ([]Project, error) {
	request, err := http.NewRequest(http.MethodGet, url.Projects, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.token))

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code is not 200: %v", response.StatusCode)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	print(string(body))

	var output []Project

	if err := json.Unmarshal(body, &output); err != nil {
		return nil, err
	}

	return output, nil
}
