package task

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/felipeornelis/todoist-api-go/url"
)

func (s Service) Close(task *task) error {
	if task.ID == "" {
		return errors.New("no ID provided")
	}

	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/%s/close", url.Tasks, task.ID), nil)
	if err != nil {
		return err
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.token))

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
