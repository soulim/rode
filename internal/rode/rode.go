package rode

import (
	"fmt"
	"io"
)

type Feed struct {
	Items []FeedItem
}

type FeedItem struct {
	GUID  string
	Title string
}

type FeedParser interface {
	Parse([]byte, *Feed) error
}

type Command interface {
	Run([]string) error
}

func Run(stdin io.Reader, stdout io.Writer, stderr io.Writer, args []string) error {
	var cmd Command

	switch args[0] {
	case "list":
		cmd = NewListCommand(stdin, stdout, stderr)
	default:
		fmt.Fprintf(stderr, "error: %q command not found\n", args[0])
		return fmt.Errorf("cannot find command %q", args[0])
	}

	if err := cmd.Run(args[1:]); err != nil {
		return err
	}

	return nil
}
