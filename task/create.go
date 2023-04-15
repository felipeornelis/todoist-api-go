package task

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/felipeornelis/todoist-api-go/url"
	"github.com/google/uuid"
)

type CreateTaskInput struct {
	Content     string   `json:"content"`
	Description string   `json:"description,omitempty"`
	ProjectID   string   `json:"project_id,omitempty"`
	SectionID   string   `json:"section_id,omitempty"`
	ParentID    string   `json:"parent_id,omitempty"`
	Order       uint     `json:"order,omitempty"`
	Labels      []string `json:"labels,omitempty"`
	Priority    uint8    `json:"priority,omitempty"`
	DueString   string   `json:"due_string,omitempty"`
	DueDatetime string   `json:"due_datetime,omitempty"`
	DueLang     string   `json:"due_lang,omitempty"`
	AssigneeID  string   `json:"assignee_id,omitempty"`
}

func (s Service) Create(input CreateTaskInput) (*Task, error) {
	if input.Content == "" {
		return nil, errors.New("content is a required field")
	}

	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, url.Tasks, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.Token))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("x-request-id", uuid.New().String())

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

	var task Task
	if err := json.Unmarshal(body, &task); err != nil {
		return nil, err
	}

	return &task, nil
}
