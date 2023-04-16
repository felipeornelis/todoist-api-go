package task

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/felipeornelis/todoist-api-go/url"
)

func (s Service) GetAll() ([]task, error) {
	request, err := http.NewRequest(http.MethodGet, url.Tasks, nil)
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

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var output []task

	if err := json.Unmarshal(body, &output); err != nil {
		return nil, err
	}

	return output, nil
}
