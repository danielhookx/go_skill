package net

import (
	"net"
	"testing"
)

// Below init function
func TestTCPServer_Run(t *testing.T) {
	// Simply check that the server is up and can
	// accept connections.
	conn, err := net.Dial("tcp", ":11123")
	if err != nil {
		t.Error("could not connect to server: ", err)
	}
	defer conn.Close()
}
