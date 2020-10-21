package repository

import "testing"

func TestGetHistory(t *testing.T) {
    cases := []struct {
		DBPath string
		expected error
    }{
		{
			"/home/punkrou404/.chc/History",
			nil,
		},
    }

    for _, c := range cases {
		actual := GetHistory(c.DBPath)
        if actual != c.expected {
            t.Error("failed.\n")
        }
    }
}