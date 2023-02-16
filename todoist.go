package todoist

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/google/uuid"
)

type Todoist struct {
	token string
}

const endpoint string = "https://api.todoist.com/rest/v2/"

var (
	tasksURL = endpoint + "tasks"
)

func New(authenticationToken string) *Todoist {
	return &Todoist{
		token: authenticationToken,
	}
}

func Request(method, url, token string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")

	if method != http.MethodGet {
		req.Header.Add("X-Request-Id", uuid.New().String())
	}

	c := &http.Client{}

	resp, err := c.Do(req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func Json(v any) (*bytes.Buffer, error) {
	input, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(input), nil
}
