package task

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/felipeornelis/todoist-api-go/url"
)

func (s Service) Reopen(task *Task) error {
	if task.ID == "" {
		return errors.New("no ID provided")
	}

	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/%s/reopen", url.Tasks, task.ID), nil)
	if err != nil {
		return err
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.Token))

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	response, err := client.Do(request)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusNoContent {
		return err
	}

	return nil
}
