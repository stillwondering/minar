package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	AppName       = "minar"
	ListenAddress = ":8080"
)

func main() {
	fmt.Println(os.Args)
	if err := run(os.Args[1:]); err != nil {
		//fmt.Println(err)
		os.Exit(1)
	}
}

func run(args []string) error {
	_, err := configFromCmdline(args)
	if err != nil {
		return err
	}

	return nil
}

type config struct {
	listenAddress string
}

func configFromCmdline(cmdline []string) (*config, error) {
	fs := flag.NewFlagSet(AppName, flag.ContinueOnError)

	listenAddress := fs.String("address", ListenAddress, "listen on this address")

	if err := fs.Parse(cmdline); err != nil {
		return nil, err
	}

	return &config{
		listenAddress: *listenAddress,
	}, nil
}
