package cmd

import (
	"flag"
	"fmt"
	"io"
)

type httpConfig struct {
	url  string
	verb string
}

func HandleHttp(w io.Writer, args []string) error {
	var v string
	fs := flag.NewFlagSet("http", flag.ExitOnError)
	fs.SetOutput(w)
	fs.StringVar(&v, "verb", "GET", "HTTP method")

	fs.Usage = func() {
		var usageString = `
			http: A HTTP client.
			http: <options> server`
		fmt.Fprintf(w, usageString)
		fmt.Fprintln(w)
		fmt.Fprintln(w)
		fmt.Fprintln(w, "Options: ")
		fmt.Println(">>> HTTP usage")
		fs.PrintDefaults()
	}
	err := fs.Parse(args)
	if err != nil {
		return err
	}
	if fs.NArg() != 1 {
		return ErrNoServerSpecified
	}

	c := httpConfig{verb: v}
	c.url = fs.Arg(0)
	fmt.Fprintf(w, "Executing HTTP request to %s with verb %s\n", c.url, c.verb)
	return nil
}
