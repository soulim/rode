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

	cmd.BoolVar(&opts.showVersion, "version", false, "print version")

	cmd.Usage = usage
}

func usage() {
	fmt.Fprintf(cmd.Output(), "rode is a tool to parse RSS feeds and generate episode files for Hugo.\n\n")
	fmt.Fprintf(cmd.Output(), "USAGE\n\n")
	fmt.Fprintf(cmd.Output(), "  $ %s [OPTIONS]\n\n", os.Args[0])
	fmt.Fprintf(cmd.Output(), "OPTIONS\n\n")
	cmd.PrintDefaults()
	fmt.Fprintf(cmd.Output(), "\n")
	fmt.Fprintf(cmd.Output(), "EXAMPLES\n\n")
	fmt.Fprintf(cmd.Output(), "  # with default options\n")
	fmt.Fprintf(cmd.Output(), "  $ %s\n\n", os.Args[0])
}

func main() {
	cmd.Parse(os.Args[1:])

	if opts.showVersion {
		fmt.Fprintf(cmd.Output(), "%s %s (runtime: %s)\n", os.Args[0], version, runtime.Version())
		os.Exit(0)
	}
}
