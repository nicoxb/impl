package main

import (
	"fmt"
	"os"
	"flag"
	"io/ioutil"

	"github.com/zjx20/impl"
)

const (
	usage = `Usage: impl [options...] <recv> <iface>

  impl generates method stubs for recv to implement iface.

Examples:

  impl 'f *File' io.Reader
  impl Murmur hash.Hash

  Don't forget the single quotes around the receiver type
  to prevent shell globbing.
`

	defaultTemplate = "func ({{.Recv}}) {{.Name}}" +
		"({{range .Params}}{{.Name}} {{.Type}}, {{end}})" +
		"({{range .Res}}{{.Name}} {{.Type}}, {{end}})" +
		"{\n" + "panic(\"not implemented\")" + "}\n\n"
)

var (
	recv       string
	iface      string
	tmplString = defaultTemplate
)

func parseCmd() {
	flag.CommandLine.Usage = func() {
		fmt.Fprint(os.Stderr, usage)
		fmt.Fprintln(os.Stderr, "\nOptions:\n")
		flag.CommandLine.PrintDefaults()
		fmt.Fprintln(os.Stderr, "")
	}
	var tmplFile string
	flag.StringVar(&tmplFile, "t", "", "`template file` for generating stub methods")
	flag.Parse()

	if flag.NArg() < 2 {
		flag.CommandLine.Usage()
		os.Exit(2)
	}

	recv = flag.Arg(0)
	iface = flag.Arg(1)

	if tmplFile != "" {
		var err error
		data, err := ioutil.ReadFile(tmplFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to load template file: %s", err.Error())
			os.Exit(2)
		}
		tmplString = string(data)
	}
}

func main() {
	parseCmd()
	src, err := impl.Generate(recv, iface, tmplString)
	if err != nil {
		fatal(err)
	}
	fmt.Print(string(src))
}

func fatal(msg interface{}) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}
