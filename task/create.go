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
	DueDate     string   `json:"due_date,omitempty"`
	DueDatetime string   `json:"due_datetime,omitempty"`
	DueLang     string   `json:"due_lang,omitempty"`
	AssigneeID  string   `json:"assignee_id,omitempty"`
}

func (s Service) Create(input CreateTaskInput) (task, error) {
	if input.Content == "" {
		return task{}, errors.New("content is a required field")
	}

	if input.DueDate == "" {
		input.DueDate = time.Now().Format("2006-01-02")
	}

	payload, err := json.Marshal(input)
	if err != nil {
		return task{}, err
	}

	request, err := http.NewRequest(http.MethodPost, url.Tasks, bytes.NewBuffer(payload))
	if err != nil {
		return task{}, err
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.token))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("x-request-id", uuid.New().String())

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	response, err := client.Do(request)
	if err != nil {
		return task{}, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return task{}, err
	}

	var output task
	if err := json.Unmarshal(body, &output); err != nil {
		return task{}, err
	}

	return task{
		ID:          output.ID,
		ProjectID:   output.ProjectID,
		SectionID:   output.SectionID,
		Content:     output.Content,
		Description: output.Description,
		IsCompleted: output.IsCompleted,
		Labels:      output.Labels,
		ParentID:    output.ParentID,
		Order:       output.Order,
		Priority:    output.Priority,
		Due: due{
			String:      output.Due.String,
			Date:        output.Due.Date,
			IsRecurring: output.Due.IsRecurring,
			Datetime:    output.Due.Datetime,
			Timezone:    output.Due.Timezone,
		},
		URL:          output.URL,
		CommentCount: output.CommentCount,
		CreatedAt:    output.CreatedAt,
		CreatorID:    output.CreatorID,
		AssigneeID:   output.AssigneeID,
		AssignerID:   output.AssignerID,
	}, nil
}
