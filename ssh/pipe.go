package ssh

import (
	"fmt"
	"io"
	"net"
	"os"
)

func Pipe(bastion, target, otp string) error {
	user, addr, err := parseDestWithFile(bastion)
	if err != nil {
		return fmt.Errorf("parsing bastion destination failed: %v", err)
	}

	_, targetAddr, err := parseDestWithFile(target)
	if err != nil {
		return fmt.Errorf("parsing target destination failed: %v", err)
	}

	client, err := connect(user, addr, otp)
	if err != nil {
		return fmt.Errorf("connection failed: %v", err)
	}

	conn, err := client.Dial("tcp", targetAddr)
	if err != nil {
		return fmt.Errorf("client dial failed: %v", err)
	}

	return pipe(conn)
}

func pipe(conn net.Conn) error {
	ch := make(chan error)

	go func() {
		buf := make([]byte, 8192)
		for {
			if _, err := io.CopyBuffer(os.Stdout, conn, buf); err != nil {
				ch <- fmt.Errorf("Rx error: %v", err)
				conn.Close()
				return
			}
		}
	}()

	go func() {
		buf := make([]byte, 8192)
		for {
			if _, err := io.CopyBuffer(conn, os.Stdin, buf); err != nil {
				ch <- fmt.Errorf("Tx error: %v", err)
				conn.Close()
				return
			}
		}
	}()

	return <-ch
}
