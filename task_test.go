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
	todoist = NewTodoist(os.Getenv("TODOIST_TOKEN"), &http.Client{
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
