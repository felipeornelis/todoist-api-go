package todoist

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/felipeornelis/todoist-api-go/dto"
	"github.com/google/uuid"
)

var (
	ErrTaskContentNotFilled           = errors.New("new task: content field is required")
	ErrTaskUnableToMarshalRequestBody = errors.New("failed to marshal request body as a JSON")
	ErrTaskUnableToCompleteRequest    = errors.New("failed to make the HTTP request")
)

type Task struct {
	ID           string   `json:"id"`
	ProjectID    string   `json:"project_id"`
	SectionID    string   `json:"section_id"`
	Content      string   `json:"content"`
	Description  string   `json:"description"`
	IsCompleted  bool     `json:"is_completed"`
	Labels       []string `json:"labels"`
	ParentID     string   `json:"parent_id"`
	Order        int      `json:"order"`
	Priority     uint8    `json:"priority"`
	Due          Due      `json:"due"`
	Url          string   `json:"url"`
	CommentCount int      `json:"comment_count"`
	CreatedAt    string   `json:"created_at"`
	CreatorID    string   `json:"creator_id"`
	AssigneeID   string   `json:"assignee_id"`
	AssignerID   string   `json:"assigner_id"`
}

type Due struct {
	String      string `json:"string"`
	Date        string `json:"date"`
	IsRecurring bool   `json:"is_recurring"`
	Datetime    string `json:"datetime,omitempty"`
	Timezone    string `json:"timezone,omitempty"`
}

func (t Todoist) NewTask(payload dto.CreateTaskRequest) (Task, error) {
	if payload.Content == "" {
		return Task{}, ErrTaskContentNotFilled
	}

	bodyRequest, err := json.Marshal(payload)
	if err != nil {
		return Task{}, fmt.Errorf("%s: %v", ErrTaskUnableToMarshalRequestBody.Error(), err)
	}

	req, err := http.NewRequest(http.MethodPost, taskURL, bytes.NewBuffer(bodyRequest))
	if err != nil {
		return Task{}, fmt.Errorf("%s: %v", ErrTaskUnableToCompleteRequest.Error(), err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", t.token))
	req.Header.Set("X-Request-Id", uuid.New().String())

	c := &http.Client{}

	resp, err := c.Do(req)
	if err != nil {
		return Task{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Task{}, err
	}

	var task Task

	if err := json.Unmarshal(body, &task); err != nil {
		return Task{}, err
	}

	return task, nil
}
