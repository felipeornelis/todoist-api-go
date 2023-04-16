package todoist

import (
	"github.com/felipeornelis/todoist-api-go/project"
	"github.com/felipeornelis/todoist-api-go/task"
)

type Todoist struct {
	Task    task.Service
	Project project.Service
}

func New(token string) *Todoist {
	return &Todoist{
		Task:    task.NewService(token),
		Project: project.NewService(token),
	}
}
