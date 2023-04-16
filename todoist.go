package todoist

import (
	"github.com/felipeornelis/todoist-api-go/project"
	"github.com/felipeornelis/todoist-api-go/task"
)

type todoist struct {
	Task    task.Service
	Project project.Service
}

func New(token string) *todoist {
	return &todoist{
		Task:    task.NewService(token),
		Project: project.NewService(token),
	}
}
