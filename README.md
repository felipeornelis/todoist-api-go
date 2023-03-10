# Todoist API Go Client (wip)
This is an Unofficial Go API Client for the Todoist REST API.

## Installation
```
go get github.com/felipeornelis/todoist-api-go
```

## Usage
An example of initialising the API client and fetching a user's tasks:
```go
package main

import (
    "fmt"

    "github.com/felipeornelis/todoist-api-go"
)

func main() {
    t := todoist.New("<YOUR TOKEN HERE>")

    tasks, err := t.Tasks()
    if err != nil {
        panic(err)
    }

    fmt.Println(tasks)
}
```

## Documentation

## Development and Testing

## Releases

## Feedback

## Contributions
