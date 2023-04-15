package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/felipeornelis/todoist-api-go/url"
)

func (s Service) Get(id string) (*Task, error) {
	if id == "" {
		return nil, errors.New("no ID provided")
	}

	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", url.Tasks, id), nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.Token))

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var task *Task
	if err := json.Unmarshal(body, &task); err != nil {
		return nil, err
	}

	return task, nil
}
