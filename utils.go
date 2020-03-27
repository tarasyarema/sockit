package main

import (
	"net"
	"regexp"

	log "github.com/sirupsen/logrus"
)

func closeConn(c net.Conn) {
	log.Infoln(c.RemoteAddr().String(), "closed")
	c.Close()
}

func toString(b []byte, n int) string {
	// Removes all trailing newline and carriage char (Windows)
	return regexp.MustCompile("\r?\n").ReplaceAllString(string(b[0:n]), "")
}
