# Todoist Client for Golang
This is an **unofficial** Todoist client written in Golang.

## Installation
It's quite simple as the following:
```bash
go get -u github.com/felipeornelis/todoist-api-go
```

## Usage
Just follow the example then nothing goes wrong
```go
package main

import (
    "fmt"

    "github.com/felipeornelis/todoist-api-go"
)

func main() {
    t := todoist.New("<YOUR TOKEN>")

    tasks, err := t.Task.GetAll()
    if err != nil {
        panic(err)
    }

    fmt.Printf("%+v\n", tasks)
}
```

## Documentation

_under development_