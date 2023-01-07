package main

import "fmt"

type Node struct {
}

func (n Node) HandlePing(p Ping) {
	fmt.Printf("ping: %+v", p)
}

func (n Node) HandleSync(s Sync) {
	fmt.Printf("sync: %+v", s)

}

func (n Node) HandleCall(c Call) {
	fmt.Printf("call: %+v", c)
}

func (n Node) HandleGetProgram(g GetProgram) {
	fmt.Printf("get program: %+v", g)
}
