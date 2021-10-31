package server

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"net/http"
)

type GinWebServer struct {
	addr string
	port int32
	*gin.Engine
	httpServer *http.Server
}

func NewWebServer(addr string, port int32) *GinWebServer {
	return &GinWebServer{
		addr: addr,
		port: port,
	}
}

func (webserver *GinWebServer) Run() {
	webserver.httpServer = &http.Server{
		Addr:    webserver.addr + ":" + fmt.Sprint(webserver.port),
		Handler: webserver,
		// ReadTimeout:    10 * time.Second,
		// WriteTimeout:   10 * time.Second,
		// MaxHeaderBytes: 1 << 20,
	}

	var eg errgroup.Group
	eg.Go(func() error {
		fmt.Printf("Start to listening the incoming requests on http address: %s", webserver.addr+":"+fmt.Sprint(webserver.port))
		if err := webserver.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Println(err)
			return err
		}
		fmt.Printf("Server on %s stopped", webserver.addr+":"+fmt.Sprint(webserver.port))

		return nil
	})

	if err := eg.Wait(); err != nil {
		fmt.Println(err.Error())
	}
}
