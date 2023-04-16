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

func (s Service) Get(id string) (task, error) {
	if id == "" {
		return task{}, errors.New("no ID provided")
	}

	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", url.Tasks, id), nil)
	if err != nil {
		return task{}, err
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.token))

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	response, err := client.Do(request)
	if err != nil {
		return task{}, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return task{}, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return task{}, err
	}

	var output task
	if err := json.Unmarshal(body, &output); err != nil {
		return task{}, err
	}

	return output, nil
}
