// * 
// * (c) Copyright 2013, Scott Cagno. All rights reserved.
// * This source code is free under the BSD style license.
// * 
// * ns :: node store - a distributed data store
// * 

package ns

// node
type Node struct {
    st Store    
    recv int 
    send string
}

// init node
func InitNode(st Store, recv) *Node {
    return &Node{
        st : st,
        recv : recv,
    }
}

// run node
func (self *Node) Run() {
    if self.st == nil {
        log.Fatal("err: missing node store!")
    }
    if self.recv == nil {
        
    }
}

// handle recv 
func (self *Node) handleRecv(conn net.Conn) {

}

// handle send
func (self *Node) handleSend(data <-chan interface{}) {

}

