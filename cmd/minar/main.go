package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
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

	return nil
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
