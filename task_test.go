package todoist

import "testing"

var td = New("aaf9decbe52a5a9a3b519d92ff5a9efff5938970")

func TestNewTask(t *testing.T) {
	tests := []struct {
		input   NewTaskInput
		wantErr bool
	}{
		{
			NewTaskInput{Content: "___Todoist SDK___ - Testing", Labels: []string{"TodoistSDK"}},
			false,
		},
		{
			NewTaskInput{},
			true,
		},
		{
			NewTaskInput{Content: ""},
			true,
		},
	}

	for _, test := range tests {
		if _, err := td.NewTask(test.input); !test.wantErr && err != nil {
			t.Errorf("todoist.NewTask(\"%s\") = %v. Got: %s", test.input.Content, test.wantErr, err)
		}
	}
}
