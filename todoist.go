package todoist

import (
	"net/http"
)

const (
	TodoistEnpoint string = "https://api.todoist.com/rest/v2"

	TaskURL = TodoistEnpoint + "/tasks"
)

type Todoist struct {
	token  string
	client *http.Client
}

func NewTodoist(token string, client *http.Client) *Todoist {
	return &Todoist{
		token:  token,
		client: client,
	}
}
