package daemon

import (
	"log"
	"net"
	"os"
)

type DaemonServer struct {
	scheduler Scheduler
	listener  net.Listener
	sockPath  string
}

func NewServer() (*DaemonServer, error) {
	sockPath := "/tmp/taskd.sock"

	if err := os.Remove(sockPath); err != nil {
		return nil, err
	}

	return &DaemonServer{
		scheduler: NewScheduler(),
		sockPath:  sockPath,
	}, nil
}

func (s *DaemonServer) Start() error {
	listener, err := net.Listen("unix", s.sockPath)
	if err != nil {
		return err
	}
	s.listener = listener

	log.Printf("Server listening on %s", s.sockPath)

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}

		go s.handleConnection(conn)
	}
}

func (s *DaemonServer) Shutdown() error {
	if err := s.listener.Close(); err != nil {
		return err
	}

	if err := os.Remove(s.sockPath); err != nil {
		return err
	}

	return nil
}

func (s *DaemonServer) handleConnection(conn net.Conn) {
	defer conn.Close()

	// TODO: Parse message
}
