package main

import (
	"flag"
	"fmt"
)

type GreetCmd struct {
	fs   *flag.FlagSet
	name string
}

func (g *GreetCmd) Name() string {
	return g.fs.Name()
}
func (g *GreetCmd) Init(args []string) error {
	return g.fs.Parse(args)
}
func (g *GreetCmd) Run() error {
	fmt.Println("Hello", g.name, "!")
	return nil
}
func NewGreetCmd() *GreetCmd {
	gc := &GreetCmd{
		fs: flag.NewFlagSet("greet", flag.ContinueOnError),
	}
	gc.fs.StringVar(&gc.name, "name", "world", "name of the person to be greeted")
	return gc
}
