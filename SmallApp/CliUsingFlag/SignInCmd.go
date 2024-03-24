package main

import (
	"flag"
	"fmt"
)

type SignInCmd struct {
	fs       *flag.FlagSet
	username string
	password string
}

func (g *SignInCmd) Name() string {
	return g.fs.Name()
}
func (g *SignInCmd) Init(args []string) error {
	return g.fs.Parse(args)
}
func (g *SignInCmd) Run() error {
	fmt.Println("Username: ", g.username)
	fmt.Println("Pwd : ", g.password)
	return nil
}
func SigninUserCmd() *SignInCmd {
	si := &SignInCmd{
		fs: flag.NewFlagSet("signin", flag.ContinueOnError),
	}
	si.fs.StringVar(&si.username, "u", "root", "User name for this app")
	si.fs.StringVar(&si.password, "p", "", "password")
	return si
}
