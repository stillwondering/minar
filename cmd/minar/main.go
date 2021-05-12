package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/stillwondering/minar/cmd/minar/templates"
)

const (
	AppName       = "minar"
	Version       = "0.1.0"
	ListenAddress = ":8080"
)

func main() {
	if err := run(os.Args[1:], os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(args []string, out io.Writer) error {
	log.SetOutput(out)

	cfg, err := configFromCmdline(args)
	if err != nil {
		return err
	}

	if cfg.printVersion {
		fmt.Fprintln(out, Version)
		return nil
	}

	fs := http.FileServer(http.Dir("assets"))

	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/", index())

	return http.ListenAndServe(cfg.listenAddress, mux)
}

type config struct {
	printVersion  bool
	listenAddress string
}

func configFromCmdline(cmdline []string) (*config, error) {
	fs := flag.NewFlagSet(AppName, flag.ExitOnError)

	printVersion := fs.Bool("version", false, "print version information")
	listenAddress := fs.String("address", ListenAddress, "listen on this address")

	if err := fs.Parse(cmdline); err != nil {
		return nil, err
	}

	return &config{
		printVersion:  *printVersion,
		listenAddress: *listenAddress,
	}, nil
}

func index() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		buf := bytes.Buffer{}

		if err := templates.Index(&buf); err != nil {
			http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		io.Copy(rw, &buf)
	}
}
