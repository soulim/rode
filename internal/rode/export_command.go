package rode

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/mattn/godown"
)

type ExportCommand struct {
	stdin  io.Reader
	stdout io.Writer
	stderr io.Writer

	flagset *flag.FlagSet

	parser FeedParser
}

func (cmd *ExportCommand) Run(args []string) error {
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

	item, ok := feed.Items[cmd.flagset.Arg(1)]

	if !ok {
		fmt.Fprintf(cmd.stderr, "error: %q feed item not found\n", args[1])
		return fmt.Errorf("cannot find feed item %q", args[1])
	}

	markdown := new(bytes.Buffer)
	html := bytes.NewBufferString(item.Description)

	if err := godown.Convert(markdown, html, nil); err != nil {
		return err
	}

	fmt.Fprintln(cmd.stdout, "+++")
	fmt.Fprintf(cmd.stdout, "title = %q\n", item.Title)
	fmt.Fprintf(cmd.stdout, "date = %q\n", item.PubDate.UTC().Format(time.RFC3339))
	fmt.Fprintln(cmd.stdout, "draft = false")
	fmt.Fprintln(cmd.stdout, "description = \"TODO: Add description\"")
	fmt.Fprintln(cmd.stdout, "cover = \"TODO: Add cover\"")
	fmt.Fprintln(cmd.stdout, "+++")
	fmt.Fprintln(cmd.stdout, "")
	fmt.Fprintln(cmd.stdout, markdown.String())

	return nil
}

func NewExportCommand(stdin io.Reader, stdout io.Writer, stderr io.Writer) *ExportCommand {
	fs := flag.NewFlagSet("rode-export", flag.ContinueOnError)

	fs.SetOutput(stderr)
	// TODO: Change default usage message to a custom one.
	// fs.Usage =

	return &ExportCommand{
		stdin:   stdin,
		stdout:  stdout,
		stderr:  stderr,
		flagset: fs,
		parser:  &RSSParser{},
	}
}
