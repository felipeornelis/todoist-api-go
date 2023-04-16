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

type UpdateTaskInput struct {
	ID          string   `json:"id"`
	Content     string   `json:"content,omitempty"`
	Description string   `json:"description,omitempty"`
	Labels      []string `json:"labels,omitempty"`
	Priority    uint8    `json:"priority,omitempty"`
	DueString   string   `json:"due_string,omitempty"`
	DueDate     string   `json:"due_date,omitempty"`
	DueDatetime string   `json:"due_datetime,omitempty"`
	// DueLang     string   `json:"due_lang,omitempty"`
	AssigneeID string `json:"assignee_id,omitempty"`
}

func (s Service) Update(input *task) error {
	if input.ID == "" {
		return errors.New("no ID provided")
	}

	payload, err := json.Marshal(UpdateTaskInput{
		ID:          input.ID,
		Content:     input.Content,
		Labels:      input.Labels,
		Priority:    input.Priority,
		DueString:   input.Due.String,
		DueDate:     input.Due.Date,
		DueDatetime: input.Due.Datetime,
		AssigneeID:  input.AssigneeID,
	})
	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/%s", url.Tasks, input.ID), bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.token))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("x-request-id", uuid.New().String())

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	response, err := client.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, input); err != nil {
		return err
	}

	return nil
}
