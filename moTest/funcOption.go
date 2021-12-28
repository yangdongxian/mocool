package moTest

import (
	"crypto/tls"
	"time"
)

type Server struct {
	Addr string
	Port int
	Protocol string
	Timeout time.Duration
	Maxconns int
	TLS *tls.Config
}

type Option func(server *Server)

func Addr(addr string) Option  {
	return func(server *Server) {
		if addr != "" {
			server.Addr = addr
		}
	}
}
func Port(port int) Option  {
	return func(server *Server) {
		if port > 0 {
			server.Port = port
		}
	}
}

func NewServer(addr string, port int, options ...Server) (*Server,error) {
	srv := &Server{
		Addr:addr,
		Port: port,
		Protocol: "TCP",
		Timeout: time.Second * 30,
		Maxconns: 1000,
		TLS: nil,
	}
	//for _,option := range options {
	//	//option(srv)
	//}
	return srv,nil
}

