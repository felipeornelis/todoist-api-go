package todoist

import "errors"

var (
	ErrEmptyToken = errors.New("todoist: no token provided")
)

const (
	baseURL string = "https://api.todoist.com/rest/v2/"

	taskURL = baseURL + "tasks"
)

type Todoist struct {
	token string
}

func New(token string) (Todoist, error) {
	if token == "" {
		return Todoist{}, ErrEmptyToken
	}

	return Todoist{
		token: token,
	}, nil
}
