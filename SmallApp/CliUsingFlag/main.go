package main

import (
	"errors"
	"fmt"
	"os"
)

func Root(args []string) error {
	if len(args) < 1 {
		return errors.New("check out the argument length. ")
	}

	cmds := []CliRunner{
		NewGreetCmd(),
		SigninUserCmd(),
	}

	subcmd := os.Args[1]
	for _, cmd := range cmds {
		if cmd.Name() == subcmd {
			cmd.Init(os.Args[2:])
			return cmd.Run()
		}
	}
	return fmt.Errorf("unknown subcmd : %s", subcmd)
}

func main() {
	if err := Root(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
