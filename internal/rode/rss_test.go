package rode_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/soulim/rode/internal/rode"
)

func TestRSSParser_Parse(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		parser := rode.RSSParser{}
		feed := &rode.Feed{}
		data, err := os.ReadFile(string(filepath.Join("testdata", "feed.xml")))
		if err != nil {
			t.Errorf("want no errors, got %q", err)
		}

		if err := parser.Parse(data, feed); err != nil {
			t.Errorf("want no errors, got %q", err)
		}

		got := len(feed.Items)
		want := 25
		if want != got {
			t.Errorf("want %v feed items, got %v", want, got)
		}
	})

	t.Run("invalid feed file", func(t *testing.T) {
		parser := rode.RSSParser{}
		feed := &rode.Feed{}
		data, err := os.ReadFile(string(filepath.Join("testdata", "invalid-feed.xml")))
		if err != nil {
			t.Errorf("want no errors, got %q", err)
		}

		if err := parser.Parse(data, feed); err == nil {
			t.Errorf("want error, got %q", err)
		}
	})
}
