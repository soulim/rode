package rode_test

import (
	"bytes"
	"path/filepath"
	"strings"
	"testing"

	"github.com/soulim/rode/internal/rode"
)

func TestRun(t *testing.T) {
	t.Run("rode list", func(t *testing.T) {
		args := []string{"list", string(filepath.Join("testdata", "feed.xml"))}

		stdin := &bytes.Buffer{}
		stdout := &bytes.Buffer{}
		stderr := &bytes.Buffer{}

		if err := rode.Run(stdin, stdout, stderr, args); err != nil {
			t.Errorf("want no errors, got %q", err)
		}

		got := stdout.String()
		want := "320bf466-e8b7-4011-a849-ae932bb0c96d Heroes and mentors\n"
		if !strings.Contains(got, want) {
			t.Errorf("stdout=%q, want %q", got, want)
		}
	})

	t.Run("unknown command", func(t *testing.T) {
		args := []string{"foo"}

		stdin := &bytes.Buffer{}
		stdout := &bytes.Buffer{}
		stderr := &bytes.Buffer{}

		if err := rode.Run(stdin, stdout, stderr, args); err == nil {
			t.Errorf("want error, got %q", err)
		}

		got := stderr.String()
		want := "error: \"foo\" command not found\n"
		if !strings.Contains(got, want) {
			t.Errorf("stderr=%q, want %q", got, want)
		}
	})
}
