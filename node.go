// *
// * (c) Copyright 2013, Scott Cagno. All rights reserved.
// * This source code is free under the BSD style license.
// *
// * ns :: node store - a distributed data store
// *

package ns

import (
	"encoding/json"
	"io"
	"log"
	"net"
	"time"
)

// data
type Dat struct {
	Cmd string      `json:"cmd,omitempty"`
	Key string      `json:"key,omitempty"`
	Val interface{} `json:"val,omitempty"`
}

// store
type Store interface {
	Has(dat *Dat) Dat
	Set(dat *Dat) Dat
	Get(dat *Dat) Dat
	Del(dat *Dat) Dat
}

// node
type Node struct {
	st         Store
	recv, send string
}

// init node
func InitNode(st Store, recv string) *Node {
	return &Node{
		st:   st,
		recv: recv,
	}
}

// run node
func (self *Node) Run() {
	ln, err := net.Listen("tcp", self.recv)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go self.recvConn(conn)
	}
}

// handle recv'd connection
func (self *Node) recvConn(conn net.Conn) {
	dec, enc := json.NewDecoder(conn), json.NewEncoder(conn)
	for {
		var dat Dat
		if err := dec.Decode(&dat); err == io.EOF {
			break
		} else if err != nil {
			conn.Close()
			return
		} else {
			conn.SetDeadline(time.Now().Add(time.Duration(900) * time.Second))
		}
		switch dat.Cmd {
		case "has":
			enc.Encode(self.st.Has(&dat))
		case "set":
			enc.Encode(self.st.Set(&dat))
		case "get":
			enc.Encode(self.st.Get(&dat))
		case "del":
			enc.Encode(self.st.Del(&dat))
		case "bye":
			enc.Encode(Dat{Val: 1})
			conn.SetDeadline(time.Now())
		default:
			enc.Encode(Dat{Val: 0})
		}
	}
}
