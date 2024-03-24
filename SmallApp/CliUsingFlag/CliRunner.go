package main

type CliRunner interface {
	Init([]string) error
	Run() error
	Name() string
}
