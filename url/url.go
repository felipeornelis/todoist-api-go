package url

import "fmt"

const baseURL string = "https://api.todoist.com/rest/v2"

var (
	Tasks    = fmt.Sprintf("%s/tasks", baseURL)
	Projects = fmt.Sprintf("%s/projects", baseURL)
)
