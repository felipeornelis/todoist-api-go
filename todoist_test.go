package todoist

import "testing"

func TestTodoist(t *testing.T) {
	tests := []struct {
		name  string
		token string
		err   error
	}{
		{
			name:  "Provide a token and return Todoist{}",
			token: "0x394382342",
			err:   nil,
		},
		{
			name:  "Do not provide a token and returns an error",
			token: "",
			err:   ErrEmptyToken,
		},
	}

	t.Log("Testing the Todoist factory")
	{
		for _, test := range tests {
			t.Logf("\tTest: %s", test.name)
			{
				_, err := New(test.token)

				if err != test.err {
					t.Errorf("todoist.New(\"%s\") = %v. Got = %v", test.token, test.err, err)
				}
			}
		}
	}
}
