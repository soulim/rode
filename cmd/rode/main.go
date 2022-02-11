package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
)

type options struct {
	showVersion bool
}

var (
	opts options
	cmd *flag.FlagSet

	version = "devel"
)

func init() {
	opts = options{}
	cmd = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	cmd.BoolVar(&opts.showVersion, "version", false, "")
	cmd.BoolVar(&opts.showVersion, "v", false, "")

	cmd.Usage = usage
}

func usage() {
	fmt.Fprintf(cmd.Output(), "rode is a tool to parse RSS feeds and generate episode files for Hugo.\n\n")
	fmt.Fprintf(cmd.Output(), "USAGE\n\n")
	fmt.Fprintf(cmd.Output(), "  $ %s [OPTIONS]\n\n", os.Args[0])
	fmt.Fprintf(cmd.Output(), "OPTIONS\n\n")
	fmt.Fprintf(cmd.Output(), "  -v, --version  Print version information\n")
	fmt.Fprintf(cmd.Output(), "  -h, --help     Print this help\n")
}

func main() {
	cmd.Parse(os.Args[1:])

	if opts.showVersion {
		fmt.Fprintf(cmd.Output(), "%s %s (runtime: %s)\n", os.Args[0], version, runtime.Version())
		os.Exit(0)
	}
}
