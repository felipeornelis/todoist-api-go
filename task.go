package todoist

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
)

var (
	errTaskContentNotFilled          = errors.New("content field is required")
	errTaskInputFailedToMarshal      = errors.New("failed to marshal the input request")
	errTaskFailedToWrapNewRequest    = errors.New("failed to wrap the new request")
	errTaskFailedToRequest           = errors.New("failed to request task data")
	errTaskBodyNotParsedToByte       = errors.New("failed to parse body response to []byte")
	errTaskResponseFailedToUnmarshal = errors.New("failed to unmarshal body response")
	errTaskIDNotProvided             = errors.New("the task ID was not provided as a parameter")
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
	Order        uint     `json:"order"`
	Priority     uint8    `json:"priority"`
	Due          Due      `json:"due"`
	URL          string   `json:"url"`
	CommentCount int      `json:"comment_count"`
	CreatedAt    string   `json:"created_at"`
	CreatorID    string   `json:"created_id"`
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

type NewTaskInputDTO struct {
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

type UpdateTaskInputDTO struct {
	ID          string   `json:"id"`
	Content     string   `json:"content,omitempty"`
	Description string   `json:"description,omitempty"`
	Labels      []string `json:"labels,omitempty"`
	Priority    uint8    `json:"priority,omitempty"`
	DueString   string   `json:"due_string,omitempty"`
	DueDate     string   `json:"due_date,omitempty"`
	DueDatetime string   `json:"due_datetime,omitempty"`
	DueLang     string   `json:"due_lang,omitempty"`
	AssigneeID  string   `json:"assignee_id,omitempty"`
}

func (t *Todoist) NewTask(input NewTaskInputDTO) (Task, error) {
	if input.Content == "" {
		return Task{}, errTaskContentNotFilled
	}

	bodyRequest, err := json.Marshal(input)
	if err != nil {
		return Task{}, errTaskInputFailedToMarshal
	}

	request, err := http.NewRequest(http.MethodPost, TaskURL, bytes.NewBuffer(bodyRequest))
	if err != nil {
		return Task{}, errTaskFailedToWrapNewRequest
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-Request-Id", uuid.New().String())
	request.Header.Set("Authorization", "Bearer "+t.token)

	response, err := t.client.Do(request)
	if err != nil {
		return Task{}, errTaskFailedToRequest
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return Task{}, errTaskBodyNotParsedToByte
	}

	var task Task
	if err := json.Unmarshal(body, &task); err != nil {
		return Task{}, errTaskResponseFailedToUnmarshal
	}

	return task, nil
}

// TODO: Apply filters
func (t *Todoist) Tasks() ([]Task, error) {
	req, err := http.NewRequest(http.MethodGet, TaskURL, nil)
	if err != nil {
		return []Task{}, errTaskFailedToWrapNewRequest
	}

	req.Header.Set("Authorization", "Bearer "+t.token)

	resp, err := t.client.Do(req)
	if err != nil {
		return []Task{}, errTaskFailedToRequest
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []Task{}, errTaskBodyNotParsedToByte
	}

	var tasks []Task
	if err := json.Unmarshal(body, &tasks); err != nil {
		return []Task{}, errTaskResponseFailedToUnmarshal
	}

	return tasks, nil
}

func (t *Todoist) Task(id string) (Task, error) {
	if id == "" {
		return Task{}, errTaskIDNotProvided
	}

	url := fmt.Sprintf("%s/%s", TaskURL, id)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Task{}, errTaskFailedToWrapNewRequest
	}

	req.Header.Set("Authorization", "Bearer "+t.token)

	resp, err := t.client.Do(req)
	if err != nil {
		return Task{}, errTaskFailedToRequest
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Task{}, errTaskBodyNotParsedToByte
	}

	var task Task

	if err := json.Unmarshal(body, &task); err != nil {
		return Task{}, errTaskResponseFailedToUnmarshal
	}

	return task, nil
}

func (t *Todoist) UpdateTask(input UpdateTaskInputDTO) (Task, error) {
	if input.ID == "" {
		return Task{}, errTaskIDNotProvided
	}

	requestBody, err := json.Marshal(input)
	if err != nil {
		return Task{}, errTaskInputFailedToMarshal
	}

	url := fmt.Sprintf("%s/%s", TaskURL, input.ID)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return Task{}, errTaskFailedToWrapNewRequest
	}

	req.Header.Set("Authorization", "Bearer "+t.token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Request-Id", uuid.New().String())

	resp, err := t.client.Do(req)
	if err != nil {
		return Task{}, errTaskFailedToRequest
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Task{}, errTaskBodyNotParsedToByte
	}

	var task Task
	if err := json.Unmarshal(body, &task); err != nil {
		return Task{}, errTaskResponseFailedToUnmarshal
	}

	return task, nil
}

func (t *Todoist) CloseTask(id string) error {
	if id == "" {
		return errTaskIDNotProvided
	}

	url := fmt.Sprintf("%s/%s/close", TaskURL, id)

	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return errTaskFailedToWrapNewRequest
	}

	req.Header.Set("Authorization", "Bearer "+t.token)

	_, err = t.client.Do(req)
	if err != nil {
		return errTaskFailedToRequest
	}

	return nil
}

func (t *Todoist) ReopenTask(id string) error {
	if id == "" {
		return errTaskIDNotProvided
	}

	url := fmt.Sprintf("%s/%s/reopen", TaskURL, id)

	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return errTaskFailedToWrapNewRequest
	}

	req.Header.Set("Authorization", "Bearer "+t.token)

	_, err = t.client.Do(req)
	if err != nil {
		return errTaskFailedToRequest
	}

	return nil
}
