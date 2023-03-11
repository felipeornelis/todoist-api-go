package todoist

import (
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

var todoist *Todoist

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	todoist = New(os.Getenv("TODOIST_TOKEN"), &http.Client{
		Timeout: 12 * time.Second,
	})
}

func TestNewTask(t *testing.T) {
	tests := []struct {
		scenario string
		input    NewTaskInputDTO
		err      error
	}{
		{
			scenario: "Create a new task successfully",
			input:    NewTaskInputDTO{Content: "I gotta do dishes after finishing it"},
			err:      nil,
		},
		{
			scenario: "Prevent from creating a task with a blank content",
			input:    NewTaskInputDTO{},
			err:      errTaskContentNotFilled,
		},
	}

	t.Log("Validate the NewTask method")
	{
		for _, test := range tests {
			t.Logf("\tTest scenario: %s", test.scenario)
			{
				_, err := todoist.NewTask(test.input)
				if err != test.err {
					t.Errorf("\tNewTask(%v) = %v; Got %v", test.input, test.err, err)
				}
			}
		}
	}
}

func TestTasks(t *testing.T) {
	tests := []struct {
		scenario string
		err      error
	}{
		{
			scenario: "Retrieve an array of tasks",
			err:      nil,
		},
	}

	t.Log("Getting all active tasks")
	{
		for _, test := range tests {
			t.Logf("\tScenario: %s", test.scenario)
			{
				_, err := todoist.Tasks()
				if err != test.err {
					t.Errorf("Tasks() = %v; Got = %v", test.err, err)
				}
			}
		}
	}
}

func TestTask(t *testing.T) {
	tests := []struct {
		scenario string
		input    string
		err      error
	}{
		{
			scenario: "Get a valid active task",
			input:    "6685702092",
			err:      nil,
		},
	}

	t.Log("Validate the Task(id) method")
	{
		for _, test := range tests {
			t.Logf("\tScenario: %s", test.scenario)
			{
				_, err := todoist.Task(test.input)
				if err != test.err {
					t.Errorf("Task(%s) = %v; Got = %v", test.input, test.err, err)
				}
			}
		}
	}
}

func TestUpdateTask(t *testing.T) {
	tests := []struct {
		scenario string
		input    UpdateTaskInputDTO
		err      error
	}{
		{
			scenario: "Update an active task",
			input:    UpdateTaskInputDTO{ID: "6690254681", Content: "Customise my CV based on the Turing mail content"},
			err:      nil,
		},
	}

	t.Log("Validate UpdateTask(UpdateTaskInputDTO) method")
	{
		for _, test := range tests {
			t.Logf("\tScenario: %s", test.scenario)
			{
				_, err := todoist.UpdateTask(test.input)
				if err != test.err {
					t.Errorf("UpdateTask(%v) = %v; Got = %v", test.input, test.err, err)
				}
			}
		}
	}
}

func TestCloseTask(t *testing.T) {
	tests := []struct {
		scenario string
		input    string
		err      error
	}{
		{
			scenario: "Close an active task",
			input:    "6690348288",
			err:      nil,
		},
	}

	t.Log("Validate CloseTask(id) method")
	{
		for _, test := range tests {
			t.Logf("\tScenario: %s", test.scenario)
			{
				err := todoist.CloseTask(test.input)
				if err != test.err {
					t.Errorf("CloseTask(%s) = %v; Got = %v", test.input, test.err, err)
				}
			}
		}
	}
}

func TestReopenTask(t *testing.T) {
	tests := []struct {
		scenario string
		input    string
		err      error
	}{
		{
			scenario: "Reopen a closed task",
			input:    "6690348288",
			err:      nil,
		},
	}

	t.Log("Validate ReopenTask(id) method")
	{
		for _, test := range tests {
			t.Logf("\tScenario: %s", test.scenario)
			{
				err := todoist.ReopenTask(test.input)
				if err != test.err {
					t.Errorf("ReopenTask(%s) = %v; Got = %v", test.input, test.err, err)
				}
			}
		}
	}
}

func TestDeleteTask(t *testing.T) {
	tests := []struct {
		scenario string
		input    string
		err      error
	}{
		{
			scenario: "Delete a task",
			input:    "6691656280",
			err:      nil,
		},
	}

	t.Log("Validate DeleteTask(id) method")
	{
		for _, test := range tests {
			t.Logf("\tScenario: %s", test.scenario)
			{
				err := todoist.DeleteTask(test.input)
				if err != test.err {
					t.Errorf("DeleteTask(%s) = %v; Got = %v", test.input, test.err, err)
				}
			}
		}
	}
}
