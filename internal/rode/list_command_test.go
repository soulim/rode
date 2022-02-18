package rode_test

import (
	"bytes"
	"path/filepath"
	"strings"
	"testing"

	"github.com/soulim/rode/internal/rode"
)

// type DummyFeedParser struct {
// 	parseFn func([]byte) error
// }
//
// func (p *DummyFeedParser) Parse(data []byte) error {
// 	return p.parseFn(data)
// }
//
// func NewDummyFeedParser() rode.Parser {
// 	return &DummyFeedParser{
// 		parseFn: func(_ []byte) error {
// 			return nil
// 		},
// 	}
// }

func TestListCommand_Run(t *testing.T) {
	// parser := NewDummyFeedParser()

	t.Run("ok", func(t *testing.T) {
		args := []string{string(filepath.Join("testdata", "feed.xml"))}

		stdin := &bytes.Buffer{}
		stdout := &bytes.Buffer{}
		stderr := &bytes.Buffer{}

		cmd := rode.NewListCommand(stdin, stdout, stderr)
		// cmd.SetParser(parser)

		if err := cmd.Run(args); err != nil {
			t.Errorf("want no errors, got %q", err)
		}

		got := stdout.String()
		want := "320bf466-e8b7-4011-a849-ae932bb0c96d Heroes and mentors\n"
		if !strings.Contains(got, want) {
			t.Errorf("stdout=%q, want %q", got, want)
		}
	})

	t.Run("invalid feed file", func(t *testing.T) {
		args := []string{string(filepath.Join("testdata", "invalid-feed.xml"))}

		stdin := &bytes.Buffer{}
		stdout := &bytes.Buffer{}
		stderr := &bytes.Buffer{}

		cmd := rode.NewListCommand(stdin, stdout, stderr)

		if err := cmd.Run(args); err == nil {
			t.Errorf("want error, got %q", err)
		}
	})

	t.Run("no feed file", func(t *testing.T) {
		args := []string{string(filepath.Join("testdata", "non-existing-file.xml"))}

		stdin := &bytes.Buffer{}
		stdout := &bytes.Buffer{}
		stderr := &bytes.Buffer{}

		cmd := rode.NewListCommand(stdin, stdout, stderr)

		if err := cmd.Run(args); err == nil {
			t.Errorf("want error, got %q", err)
		}
	})

	t.Run("no arguments", func(t *testing.T) {
		args := []string{}

		stdin := &bytes.Buffer{}
		stdout := &bytes.Buffer{}
		stderr := &bytes.Buffer{}

		cmd := rode.NewListCommand(stdin, stdout, stderr)

		if err := cmd.Run(args); err == nil {
			t.Errorf("want error, got %q", err)
		}
	})
}
