package command

import "testing"

func TestCopyHistoryDB(t *testing.T) {
    cases := []struct {
		copiedDBPath string
		configPath string
		expected error
    }{
		{
			"/home/punkrou404/.chc/History",
			"/home/punkrou404/.chc/chc.config",
			nil,
		},
    }

    for _, c := range cases {
		actual := CopyHistoryDB(c.copiedDBPath, c.configPath)
        if actual != c.expected {
            t.Error("failed.\n")
        }
    }
}