package todoist

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var (
	CloseOperation  = "close"
	ReopenOperation = "reopen"
)

type (
	Task struct {
		ID          string   `json:"id"`
		ProjectID   string   `json:"project_id"`
		SectionID   string   `json:"section_id"`
		Content     string   `json:"content"`
		Description string   `json:"description"`
		IsCompleted bool     `json:"is_completed"`
		Labels      []string `json:"labels"`
		ParentID    string   `json:"parent_id"`
		Priority    uint8    `json:"priority"`
		Due         TaskDue  `json:"due"`
		AssigneeID  string   `json:"assignee_id"`
		AssignerID  string   `json:"assigner_id"`

		// Read only
		Order        int    `json:"order"`
		Url          string `json:"url"`
		CommentCount int    `json:"comment_count"`
		CreatedAt    string `json:"created_at"`
		CreatorID    string `json:"creator_id"`
	}

	TaskDue struct {
		String      string `json:"string"`
		Date        string `json:"date"`
		IsRecurring bool   `json:"is_recurring"`
		Datetime    string `json:"datetime,omitempty"`
		Timezone    string `json:"timezone,omitempty"`
	}

	NewTaskInput struct {
		Content     string   `json:"content"`
		Description string   `json:"description,omitempty"`
		ProjectID   string   `json:"project_id,omitempty"`
		SectionID   string   `json:"section_id,omitempty"`
		Order       int      `json:"order,omitempty"`
		Labels      []string `json:"labels,omitempty"`
		Priority    uint8    `json:"priority,omitempty"`
		DueString   string   `json:"due_string,omitempty"`
		DueDate     string   `json:"due_date,omitempty"`
		DueDatetime string   `json:"due_datetime,omitempty"`
		DueLang     string   `json:"due_lang,omitempty"`
		AssigneeID  string   `json:"assignee_id,omitempty"`
	}

	TaskUpdateInput struct {
		Content     string   `json:"content,omitempty"`
		Description string   `json:"description,omitempty"`
		Labels      []string `json:"labels,omitempty"`
		DueString   string   `json:"due_string,omitempty"`
		DueDate     string   `json:"due_date,omitempty"`
		DueDatetime string   `json:"due_datetime,omitempty"`
		DueLang     string   `json:"due_lang,omitempty"`
		AssigneeID  string   `json:"assignee_id,omitempty"`
	}
)

func (t Todoist) NewTask(payload NewTaskInput) (Task, error) {
	input, err := Json(payload)
	if err != nil {
		return Task{}, err
	}

	res, err := Request(http.MethodPost, tasksURL, t.token, input)
	if err != nil {
		return Task{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Task{}, err
	}

	var task Task

	if err := json.Unmarshal(body, &task); err != nil {
		return Task{}, err
	}

	return task, nil
}

func (t Todoist) Tasks() ([]Task, error) {
	res, err := Request(http.MethodGet, tasksURL, t.token, nil)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var tasks []Task

	if err := json.Unmarshal(body, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (t Todoist) Task(id string) (Task, error) {
	res, err := Request(http.MethodGet, fmt.Sprintf("%s/%s", tasksURL, id), t.token, nil)
	if err != nil {
		return Task{}, nil
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Task{}, err
	}

	var task Task

	if err := json.Unmarshal(body, &task); err != nil {
		return Task{}, err
	}

	return task, nil
}

func (t Todoist) TaskUpdate(id string, payload TaskUpdateInput) (Task, error) {
	buf, err := Json(payload)
	if err != nil {
		return Task{}, nil
	}

	res, err := Request(http.MethodPost, fmt.Sprintf("%s/%s", tasksURL, id), t.token, buf)
	if err != nil {
		return Task{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Task{}, err
	}

	var task Task

	if err := json.Unmarshal(body, &task); err != nil {
		return Task{}, nil
	}

	return task, nil
}

func (t Todoist) TaskClose(id string) error {
	res, err := Request(http.MethodPost, fmt.Sprintf("%s/%s/%s", tasksURL, id, CloseOperation), t.token, nil)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusNoContent {
		return fmt.Errorf("todoist.TaskClose: expected status code 204. Got %v", res.StatusCode)
	}

	return nil
}

func (t Todoist) TaskReopen(id string) error {
	res, err := Request(http.MethodPost, fmt.Sprintf("%s/%s/%s", tasksURL, id, ReopenOperation), t.token, nil)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusNoContent {
		return fmt.Errorf("todoist.TaskReopen: expected status code 204. Got %v", res.StatusCode)
	}

	return nil
}

func (t Todoist) TaskDelete(id string) error {
	res, err := Request(http.MethodDelete, fmt.Sprintf("%s/%s", tasksURL, id), t.token, nil)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusNoContent {
		return fmt.Errorf("todoist.TaskDelete: expected status code 204. Got %v", res.StatusCode)
	}

	return nil
}
