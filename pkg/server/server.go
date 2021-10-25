package server

import (
	"github.com/spf13/cobra"
	"os"
)

type Server struct {
	name    string
	runfunc RunFunc
	desc    string

	args cobra.PositionalArgs
	cmd  *cobra.Command
}

func NewServer(name string) *Server {
	return &Server{
		name: name,
	}
}

func (server *Server) Start() {
	cmd := cobra.Command{
		Use:   server.name,
		Short: server.name,
		Long:  server.desc,
		// stop printing usage when the command errors
		SilenceUsage:  true,
		SilenceErrors: true,
		Args:          server.args,
	}

	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)

	// 启动
	server.runfunc()
}

type RunFunc func() error
