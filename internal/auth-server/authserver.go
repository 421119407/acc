package auth_server

import (
	"acc/pkg/server"
	"fmt"
)

type authServer struct {
	webServer *server.GinWebServer
}

func (authServer *authServer) Run() {
	initController(authServer.webServer.Engine)
	authServer.webServer.Run()
}

func NewServer(name string) *server.Server {
	server := server.NewServer(
		name,
		server.WithRunFunc(run()),
	)
	return server
}

func run() server.RunFunc {
	return func(s *server.Server) error {
		ginServer := server.NewWebServer("localhost", 8080)
		authserver := &authServer{
			webServer: ginServer,
		}
		authserver.Run()
		fmt.Printf("服务启动: %s", s.Name())
		return nil
	}
}
