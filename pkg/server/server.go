package server

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os"
)

type Server struct {
	name string
	desc string

	childCommand []*cobra.Command
	args         cobra.PositionalArgs
	rootCmd      *cobra.Command
	runFunc      RunFunc
}

type Option func(server *Server)

type RunFunc func(server *Server) error

func (server *Server) AddCommand(command *cobra.Command) *Server {
	server.childCommand = append(server.childCommand, command)
	fmt.Println("")
	return server
}

func (server *Server) Name() string {
	return server.name
}

func WithRunFunc(runFunc RunFunc) Option {
	return func(server *Server) {
		server.runFunc = runFunc
	}
}

func NewServer(name string, options ...Option) *Server {
	server := &Server{
		name: name,
	}

	for _, o := range options {
		o(server)
	}

	server.buildCommand()

	return server
}

func (server *Server) buildCommand() {
	cmd := cobra.Command{
		Use:           server.name,
		Short:         server.name,
		Long:          server.desc,
		SilenceUsage:  true,
		SilenceErrors: true,
		Args:          server.args,
	}

	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)
	cmd.Flags().SortFlags = true

	if len(server.childCommand) > 0 {
		for _, command := range server.childCommand {
			cmd.AddCommand(command)
		}
	}

	cmd.RunE = server.Run
	server.rootCmd = &cmd
}

func (server *Server) Start() {
	if err := server.rootCmd.Execute(); err != nil {
		fmt.Printf("%v %v\n", color.RedString("Error:"), err)
		os.Exit(1)
	}
}

func (server *Server) Run(cmd *cobra.Command, args []string) error {
	if server.runFunc == nil {
		return errors.New("No command to start function")
	}

	server.runFunc(server)
	return nil
}
