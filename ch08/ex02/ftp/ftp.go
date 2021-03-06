// Package ftp は、FTP サーバの機能を提供します。
package ftp

import (
	"bufio"
	"fmt"
	"log"
	"strings"
)

// Serve は、FTP 接続をハンドルして、FTP サーバとして振る舞います。
func Serve(c Conn) {
	c.respond("220 Service ready for new user.")

	s := bufio.NewScanner(c.conn)
	for s.Scan() {
		input := strings.Fields(s.Text())
		if len(input) == 0 {
			continue
		}
		command, args := input[0], input[1:]
		log.Printf("<< %s %v %s", command, args, c.dataport.toAddress())
		switch command {
		case "CWD":
			c.cwd(args)
		case "EPSV":
			c.epsv(args)
		case "LIST":
			c.list(args)
		case "MODE":
			c.mode(args)
		case "NOOP":
			c.noop(args)
		case "PORT":
			c.port(args)
		case "PWD":
			c.pwd(args)
		case "RETR":
			c.retr(args)
		case "SIZE":
			c.size(args)
		case "STOR":
			c.stor(args)
		case "STRU":
			c.stru(args)
		case "SYST":
			c.syst(args)
		case "TYPE":
			c.typeCommand(args)
		case "USER":
			c.user(args)
		case "QUIT":
			c.respond("221 Service closing control connection.")
			return
		default:
			c.respond("502 Command not implemented.")
		}
	}
	if s.Err() != nil {
		log.Print(s.Err())
	}
}

func (c *Conn) respond(s string) {
	log.Print(">> ", s)
	_, err := fmt.Fprint(c.conn, s, c.eol())
	if err != nil {
		log.Print(err)
	}
}
