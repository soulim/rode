package rode

import (
	"flag"
	"fmt"
	"io"
	"os"
)

type ListCommand struct {
	stdin  io.Reader
	stdout io.Writer
	stderr io.Writer

	flagset *flag.FlagSet

	parser FeedParser
}

func (cmd *ListCommand) SetParser(parser FeedParser) {
	cmd.parser = parser
}

func (cmd *ListCommand) Run(args []string) error {
	if err := cmd.flagset.Parse(args); err != nil {
		return err
	}

	data, err := os.ReadFile(cmd.flagset.Arg(0))
	if err != nil {
		return err
	}

	feed := &Feed{}

	if err := cmd.parser.Parse(data, feed); err != nil {
		return err
	}

	for _, item := range feed.Items {
		fmt.Fprintf(cmd.stdout, "%s %s\n", item.GUID, item.Title)
	}

	return nil
}

func NewListCommand(stdin io.Reader, stdout io.Writer, stderr io.Writer) *ListCommand {
	fs := flag.NewFlagSet("rode-list", flag.ContinueOnError)

	fs.SetOutput(stderr)
	// TODO: Change default usage message to a custom one.
	// fs.Usage =

	return &ListCommand{
		stdin:   stdin,
		stdout:  stdout,
		stderr:  stderr,
		flagset: fs,
		parser:  &RSSParser{},
	}
}
