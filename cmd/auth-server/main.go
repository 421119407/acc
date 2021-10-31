package main

import auth_server "acc/internal/auth-server"

func main() {
	auth_server.NewServer("auth-server").Start()
}
