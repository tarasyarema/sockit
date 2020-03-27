package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		log.Infoln(conn.RemoteAddr().String(), "accepted")

		go func(c net.Conn) {
			defer closeConn(c)

			c.Write([]byte("give me a name: "))

			b := make([]byte, 64)
			n, err := c.Read(b)
			if err != nil {
				log.Errorf("could not read name bytes: %v", err.Error())
				return
			}

			name := toString(b, n)
			log.Infof("name: %v %d", name, len(name))

			if len(name) == 0 {
				msg := "okok, privacy is nice but next time give me a name pls :3\n"

				if _, err = c.Write([]byte(msg)); err != nil {
					log.Errorf("could not write err name len bytes: %v", err.Error())
				}

				return
			}

			greet := fmt.Sprintf("Hi %s, who are you?\n", name)

			n, err = c.Write([]byte(greet))
			if err != nil {
				log.Errorf("could not write greeting bytes: %v", err.Error())
				return
			}

			log.Infoln("wrote", n, "bytes")

			game, err := newGame(name)
			if err != nil {
				msg := fmt.Sprintf("%s\n", err.Error())

				if _, err = c.Write([]byte(msg)); err != nil {
					log.Error(err)
				}
				return
			}

			log.Infof("the game began")

			for {
				if err := game.print(c); err != nil {
					log.Error(err)
					return
				}

				c.Write([]byte("give me a position (comma separated): "))

				b := make([]byte, 64)
				n, err := c.Read(b)
				if err != nil {
					log.Errorf("could not read name bytes: %v", err.Error())
					return
				}

				raw := toString(b, n)
				pos := strings.Split(raw, ",")

				if len(pos) != 2 {
					if _, err = c.Write([]byte("wrong position format\n")); err != nil {
						log.Error(err)
					}

					continue
				}

				x, errX := strconv.ParseInt(pos[0], 10, 64)
				y, errY := strconv.ParseInt(pos[1], 10, 64)

				if errX != nil || errY != nil {
					if _, err = c.Write([]byte("wrong position format\n")); err != nil {
						log.Error(err)
					}

					continue
				}

				msg, err := game.Move(int(x), int(y))

				if err != nil {
					if _, err = c.Write([]byte(msg)); err != nil {
						log.Error(err)
					}
					continue
				}
			}
		}(conn)
	}
}
